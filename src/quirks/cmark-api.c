/* Runtime-independent string API. Node-API and JavaScript types stay outside. */
#include "cmark-api.h"

static const char *gfm_empty_token = "00000000000000000000000000000000";

static gfm_status gfm_render_string(gfm_view source, int xml, gfm_html_filter filter,
                                     void *context, gfm_result *result) {
  const int options = CMARK_OPT_UNSAFE | CMARK_OPT_FOOTNOTES |
                      CMARK_OPT_GITHUB_PRE_LANG | CMARK_OPT_HARDBREAKS;
  cmark_gfm_core_extensions_ensure_registered();
  gfm_document *document = gfm_document_new(source.data, source.length, options);
  if (!document) { result->error = "Cannot parse Markdown"; return GFM_NO_MEMORY; }
  gfm_status status = GFM_OK;
  if (filter) {
    for (size_t i = 0; i < document->raw_count; i++) {
      cmark_node *node = document->raw_html[i];
      const char *original = cmark_node_get_literal(node);
      gfm_view input = { (const unsigned char *)original, strlen(original) };
      gfm_string output = {0};
      status = filter(input, &output, context);
      if (status == GFM_OK && !output.data && output.length) {
        result->error = "Raw HTML filter returned invalid string data"; status = GFM_BAD_ARGUMENT;
      }
      if (status == GFM_OK) {
        int changed = input.length != output.length ||
          (input.length && (!output.data || memcmp(input.data, output.data, input.length)));
        if (changed && !cmark_node_set_literal(node, output.data ? (const char *)output.data : "")) {
          result->error = "Cannot replace raw HTML"; status = GFM_NO_MEMORY;
        }
      }
      free(output.data);
      if (status != GFM_OK) {
        if (!result->error) result->error = "Raw HTML filter failed";
        break;
      }
    }
  }
  if (status == GFM_OK) {
    result->text.data = (unsigned char *)(xml ? gfm_document_xml(document) : gfm_document_html(document));
    if (!result->text.data) { result->error = "Cannot render Markdown"; status = GFM_NO_MEMORY; }
    else result->text.length = strlen((const char *)result->text.data);
  }
  gfm_document_free(document);
  return status;
}

gfm_status gfm_call(const char *operation, const gfm_view *arguments, size_t count,
                    gfm_html_filter filter, void *context, gfm_result *result) {
  if (!result) return GFM_BAD_ARGUMENT;
  *result = (gfm_result){0};
  if (!operation || !arguments || count < 1 || count > 4) {
    result->error = "Expected one to four string arguments"; return GFM_BAD_ARGUMENT;
  }
  for (size_t i = 0; i < count; i++) {
    if ((!arguments[i].data && arguments[i].length) || (i == 0 && !arguments[i].data)) {
      result->error = "String data is missing"; return GFM_BAD_ARGUMENT;
    }
  }
  gfm_view source = arguments[0];
  size_t minimum = 1, maximum = 1;
  if (!strcmp(operation, "finalize")) maximum = 3;
  else if (!strcmp(operation, "math")) minimum = maximum = 3;
  else if (!strcmp(operation, "links") || !strcmp(operation, "templates") || !strcmp(operation, "boundaries") || !strcmp(operation, "prepareStage")) minimum = maximum = 2;
  else if (!strcmp(operation, "mathTables") || !strcmp(operation, "mathStage") || !strcmp(operation, "finish")) minimum = maximum = 4;
  else if (strcmp(operation, "render") && strcmp(operation, "ast") && strcmp(operation, "version") &&
           strcmp(operation, "lowercase") && strcmp(operation, "decode") && strcmp(operation, "details")) {
    result->error = "Unknown native operation"; return GFM_BAD_ARGUMENT;
  }
  if (count < minimum || count > maximum) {
    result->error = "Unexpected number of string arguments"; return GFM_BAD_ARGUMENT;
  }
  for (size_t i = 1; i < minimum; i++) {
    if (!arguments[i].data) { result->error = "Required string argument is missing"; return GFM_BAD_ARGUMENT; }
  }
  if (!strcmp(operation, "render") || !strcmp(operation, "ast"))
    return gfm_render_string(source, !strcmp(operation, "ast"), filter, context, result);
  if (!strcmp(operation, "lowercase") || !strcmp(operation, "decode") || !strcmp(operation, "version")) {
    cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    if (!strcmp(operation, "lowercase")) gfm_unicode_lower_into(&out, source.data, source.length);
    else if (!strcmp(operation, "decode")) houdini_unescape_html_f(&out, source.data, source.length);
    else cmark_strbuf_puts(&out, cmark_version_string());
    result->text.length = out.size;
    result->text.data = cmark_strbuf_detach(&out);
  } else {
    /* Views need not be NUL-terminated. Keep C-string modifier adapters here,
     * independently of the runtime's input allocation strategy. */
    char *modifiers[3] = {0};
    for (size_t i = 1; i < count; i++) {
      if (!arguments[i].data) continue;
      if (arguments[i].length != SIZE_MAX) modifiers[i - 1] = malloc(arguments[i].length + 1);
      if (!modifiers[i - 1]) {
        for (size_t n = 0; n < 3; n++) free(modifiers[n]);
        result->error = "Cannot allocate modifier"; return GFM_NO_MEMORY;
      }
      memcpy(modifiers[i - 1], arguments[i].data, arguments[i].length);
      modifiers[i - 1][arguments[i].length] = 0;
    }
    const char *a = modifiers[0] ? modifiers[0] : gfm_empty_token;
    const char *b = modifiers[1] ? modifiers[1] : "";
    const char *c = modifiers[2] ? modifiers[2] : "";
    char *output = NULL;
    if (!strcmp(operation, "finalize")) output = gfm_finalize_serialized_html(source.data, source.length, a, b);
    else if (!strcmp(operation, "math")) output = gfm_math_dom(source.data, source.length, a, b);
    else if (!strcmp(operation, "links")) output = gfm_process_links(source.data, source.length, a);
    else if (!strcmp(operation, "details")) output = gfm_prepare_details(source.data, source.length);
    else if (!strcmp(operation, "templates")) output = gfm_flatten_templates(source.data, source.length, a);
    else if (!strcmp(operation, "boundaries")) output = gfm_remove_boundaries(source.data, source.length, a);
    else if (!strcmp(operation, "mathTables")) output = gfm_prepare_math_tables(source.data, source.length, a, b, c);
    else if (!strcmp(operation, "prepareStage")) output = gfm_prepare_html_stage(source.data, source.length, a);
    else if (!strcmp(operation, "mathStage")) output = gfm_render_math_stage(source.data, source.length, a, b, c);
    else if (!strcmp(operation, "finish")) output = gfm_finish_html(source.data, source.length, a, b, c);
    for (size_t i = 0; i < 3; i++) free(modifiers[i]);
    result->text.data = (unsigned char *)output;
    if (output) result->text.length = strlen(output);
  }
  if (!result->text.data) { result->error = "Native string allocation failed"; return GFM_NO_MEMORY; }
  return GFM_OK;
}
