/* Test-only bindings for private helpers; none are exported by src/. */
#include "../../src/quirks/cmark/bun.c"

/* The upstream XML renderer uses <unknown> for footnote node types. */
static char *test_document_xml(gfm_document *document) {
  char *xml = cmark_render_xml(document->doc, CMARK_OPT_SOURCEPOS);
  if (!xml) return NULL;
  cmark_strbuf output = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  const char *p = xml;
  const char *open = "<<unknown>", *close = "</<unknown>>";
  size_t open_len = strlen(open), close_len = strlen(close);
  while (*p) {
    if (!strncmp(p, close, close_len)) {
      cmark_strbuf_puts(&output, "</footnote_definition>"); p += close_len; continue;
    }
    if (!strncmp(p, open, open_len)) {
      const char *attrs = p + open_len, *end = strchr(attrs, '>');
      if (end) {
        const char *slash = end;
        while (slash > attrs && (slash[-1] == ' ' || slash[-1] == '\t' || slash[-1] == '\r' || slash[-1] == '\n')) slash--;
        if (slash > attrs && slash[-1] == '/') {
          const char *attr_end = slash - 1;
          while (attr_end > attrs && (attr_end[-1] == ' ' || attr_end[-1] == '\t' || attr_end[-1] == '\r' || attr_end[-1] == '\n')) attr_end--;
          cmark_strbuf_puts(&output, "<footnote_reference");
          cmark_strbuf_put(&output, (const unsigned char *)attrs, attr_end - attrs);
          cmark_strbuf_puts(&output, " />"); p = end + 1; continue;
        }
      }
      cmark_strbuf_puts(&output, "<footnote_definition"); p += open_len; continue;
    }
    cmark_strbuf_putc(&output, (unsigned char)*p++);
  }
  free(xml);
  return (char *)cmark_strbuf_detach(&output);
}

napi_value test_gfm_helper(napi_env env, napi_value operation, napi_value source,
                          napi_value a, napi_value b, napi_value c) {
  gfm_string name = {0}, inputs[4] = {{0}};
  napi_value values[4] = {source, a, b, c}, value = NULL;
  gfm_result result = {0}; char *html = NULL; int reparse = 0, staged = 0;
  gfm_status status = GFM_OK;
  if (!gfm_bun_read_string(env, operation, &name, "Expected test operation")) goto done;
  const char *op = (char *)name.data;
  size_t count = (!strcmp(op, "prepare") || !strcmp(op, "sanitize")) ? 2 :
                 (!strcmp(op, "math") || !strcmp(op, "finish")) ? 4 : 1;
  for (size_t i = 0; i < count; i++)
    if (!gfm_bun_read_string(env, values[i], &inputs[i], "Expected string modifier")) goto done;
  gfm_view input = {inputs[0].data, inputs[0].length};
  if (!strcmp(op, "markdown")) status = gfm_render_markdown(input, &result);
  else if (!strcmp(op, "ast")) {
    gfm_document *document = NULL;
    status = gfm_parse_markdown(input, &document, &result);
    if (status == GFM_OK) { html = test_document_xml(document); gfm_document_free(document); }
  } else if (!strcmp(op, "lowercase") || !strcmp(op, "version")) {
    cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    if (!strcmp(op, "lowercase")) gfm_unicode_lower_into(&out, input.data, input.length);
    else cmark_strbuf_puts(&out, cmark_version_string());
    result.text.length = out.size; result.text.data = cmark_strbuf_detach(&out);
  } else if (!strcmp(op, "rawInputs")) status = gfm_filter_raw_inputs(input, &result.text);
  else if (!strcmp(op, "sanitize")) {
    if (input.length > GFM_MAX_OUTPUT_BYTES) { status = GFM_LIMIT_EXCEEDED; goto failed; }
    if (inputs[1].length >= 64) { gfm_bun_error(env, GFM_BUN_TYPE_ERROR, "Invalid sanitizer boundary"); goto done; }
    h_document document = {0}; h_buffer out = {0};
    if (!h_parse(&document, input.data, input.length, 0)) status = GFM_NO_MEMORY;
    else h_sanitize(&out, &document, document.root, (char *)inputs[1].data);
    h_destroy(&document);
    if (out.failed) status = out.failed == H_ALLOC_FAILED ? GFM_NO_MEMORY : GFM_LIMIT_EXCEEDED;
    if (status != GFM_OK) { free(out.data); goto failed; }
    if (!out.data) out.data = (unsigned char *)calloc(1, 1);
    result.text = (gfm_string){out.data, out.length};
  } else if (!strcmp(op, "prepare")) {
    html = gfm_prepare_html_stage(input.data, input.length, (char *)inputs[1].data, &reparse); staged = 1;
  } else if (!strcmp(op, "math")) {
    html = gfm_render_math_stage(input.data, input.length, (char *)inputs[1].data, (char *)inputs[2].data, (char *)inputs[3].data, &reparse); staged = 1;
  } else if (!strcmp(op, "finish"))
    html = gfm_finish_html(input.data, input.length, (char *)inputs[1].data, (char *)inputs[2].data, (char *)inputs[3].data);
  else { gfm_bun_error(env, GFM_BUN_TYPE_ERROR, "Unknown native test helper"); goto done; }
  if (html) {
    size_t length = strlen(html);
    if (staged) {
      result.text.data = malloc(length + 2); result.text.length = length + 1;
      if (result.text.data) { result.text.data[0] = reparse ? '1' : '0'; memcpy(result.text.data + 1, html, length + 1); }
    } else { result.text = (gfm_string){(unsigned char *)html, length}; html = NULL; }
  }
  if (status == GFM_OK && !result.text.data) status = GFM_NO_MEMORY;
  if (status == GFM_OK && !strcmp(op, "finish")) status = gfm_complete_html(&result);
failed:
  if (status != GFM_OK) {
    gfm_bun_error(env, status == GFM_LIMIT_EXCEEDED ? GFM_BUN_RANGE_ERROR : GFM_BUN_ERROR,
                  result.error ? result.error : "Native test helper failed"); goto done;
  }
  if (napi_create_string_utf8(env, (char *)result.text.data, result.text.length, &value) != napi_ok) value = NULL;
done:
  free(name.data); for (size_t i = 0; i < 4; i++) free(inputs[i].data);
  free(html); free(result.text.data); return value;
}
