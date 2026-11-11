/* Own the parser, any recovered source, and its AST as a single FFI object. */
typedef struct {
  cmark_parser *parser;
  cmark_node *doc;
  cmark_node **raw_html;
  size_t raw_count, raw_capacity;
  int options;
} gfm_document;

static void gfm_document_free(gfm_document *document) {
  if (!document) return;
  if (document->doc) cmark_node_free(document->doc);
  if (document->parser) cmark_parser_free(document->parser);
  free(document->raw_html);
  free(document);
}

static int gfm_document_parse(gfm_document *document, const unsigned char *source, size_t length) {
  document->parser = gfm_parser_new(document->options);
  if (!document->parser) return 0;
  cmark_parser_feed(document->parser, (const char *)source, length);
  document->doc = cmark_parser_finish(document->parser);
  return document->doc != NULL;
}

static gfm_document *gfm_document_new(const unsigned char *source, size_t length, int options) {
  gfm_document *document = calloc(1, sizeof(*document));
  if (!document) return NULL;
  document->options = options;
  if (!gfm_document_parse(document, source, length)) { gfm_document_free(document); return NULL; }
  gfm_recovered *recovered = gfm_recover_labels(document->doc, source, length);
  if (!recovered) { gfm_document_free(document); return NULL; }
  const unsigned char *input = source;
  size_t input_length = length;
  if (recovered->changed) {
    cmark_node_free(document->doc); document->doc = NULL;
    cmark_parser_free(document->parser); document->parser = NULL;
    input = recovered->data; input_length = recovered->length;
    if (!gfm_document_parse(document, input, input_length)) {
      free(recovered); gfm_document_free(document); return NULL;
    }
  }
  int ok = gfm_adapt_nodes(document->doc, (const char *)input, input_length);
  free(recovered);
  if (!ok) { gfm_document_free(document); return NULL; }
  cmark_iter *iter = cmark_iter_new(document->doc);
  if (!iter) { gfm_document_free(document); return NULL; }
  cmark_event_type event;
  while ((event = cmark_iter_next(iter)) != CMARK_EVENT_DONE) {
    if (event != CMARK_EVENT_ENTER) continue;
    cmark_node *node = cmark_iter_get_node(iter);
    const char *type = cmark_node_get_type_string(node);
    if (strcmp(type, "html_inline") && strcmp(type, "html_block")) continue;
    if (!gfm_grow((void **)&document->raw_html, &document->raw_capacity,
                  document->raw_count + 1, sizeof(*document->raw_html))) { ok = 0; break; }
    document->raw_html[document->raw_count++] = node;
  }
  cmark_iter_free(iter);
  if (!ok) { gfm_document_free(document); return NULL; }
  return document;
}

static char *gfm_document_html(gfm_document *document) {
  return cmark_render_html(document->doc, document->options,
                           cmark_parser_get_syntax_extensions(document->parser));
}

/* The upstream XML renderer uses <unknown> for footnote node types. */
static char *gfm_document_xml(gfm_document *document) {
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
