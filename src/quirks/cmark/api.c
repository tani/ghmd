/* Runtime-independent string API. Node-API and JavaScript types stay outside. */
#include "api.h"

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

static gfm_status gfm_render_document(gfm_view source, gfm_view token, gfm_view boundary, gfm_result *result);

gfm_status gfm_call(const char *operation, const gfm_view *arguments, size_t count,
                    gfm_html_filter filter, void *context, gfm_result *result) {
  if (!result) return GFM_BAD_ARGUMENT;
  *result = (gfm_result){0};
  if (!operation || !arguments || count < 1 || count > 3) {
    result->error = "Expected one to three string arguments"; return GFM_BAD_ARGUMENT;
  }
  size_t expected = 1;
  if (!strcmp(operation, "document")) expected = 3;
  else if (strcmp(operation, "render") && strcmp(operation, "ast") && strcmp(operation, "version") &&
           strcmp(operation, "lowercase") && strcmp(operation, "rawInputs")) {
    result->error = "Unknown native operation"; return GFM_BAD_ARGUMENT;
  }
  if (count != expected) {
    result->error = "Unexpected number of string arguments"; return GFM_BAD_ARGUMENT;
  }
  for (size_t i = 0; i < count; i++) {
    if (!arguments[i].data) { result->error = "Required string argument is missing"; return GFM_BAD_ARGUMENT; }
  }
  gfm_view source = arguments[0];
  if (!strcmp(operation, "document")) return gfm_render_document(source, arguments[1], arguments[2], result);
  if (!strcmp(operation, "render") || !strcmp(operation, "ast"))
    return gfm_render_string(source, !strcmp(operation, "ast"), filter ? filter : gfm_filter_raw_inputs, context, result);
  if (!strcmp(operation, "rawInputs")) return gfm_filter_raw_inputs(source, &result->text, NULL);
  if (!strcmp(operation, "lowercase") || !strcmp(operation, "version")) {
    cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    if (!strcmp(operation, "lowercase")) gfm_unicode_lower_into(&out, source.data, source.length);
    else cmark_strbuf_puts(&out, cmark_version_string());
    result->text.length = out.size;
    result->text.data = cmark_strbuf_detach(&out);
  }
  if (!result->text.data) { result->error = "Native string allocation failed"; return GFM_NO_MEMORY; }
  return GFM_OK;
}
