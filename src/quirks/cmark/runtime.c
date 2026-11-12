/* Own the parser, any recovered source, and its AST during one C render. */
typedef struct {
  cmark_parser *parser;
  cmark_node *doc;
  cmark_node **raw_html;
  size_t raw_count, raw_capacity;
  int options;
} gfm_document;

static void gfm_document_free(gfm_document *document) {
  if (!document)
    return;
  if (document->doc)
    cmark_node_free(document->doc);
  if (document->parser)
    cmark_parser_free(document->parser);
  free(document->raw_html);
  free(document);
}

static int gfm_document_parse(gfm_document *document, const unsigned char *source, size_t length) {
  document->parser = gfm_parser_new(document->options);
  if (!document->parser)
    return 0;
  cmark_parser_feed(document->parser, (const char *)source, length);
  document->doc = cmark_parser_finish(document->parser);
  return document->doc != NULL;
}

static gfm_document *gfm_document_new(const unsigned char *source, size_t length, int options) {
  gfm_document *document = calloc(1, sizeof(*document));
  if (!document)
    return NULL;
  document->options = options;
  if (!gfm_document_parse(document, source, length)) {
    gfm_document_free(document);
    return NULL;
  }
  gfm_recovered *recovered = gfm_recover_labels(document->doc, source, length);
  if (!recovered) {
    gfm_document_free(document);
    return NULL;
  }
  const unsigned char *input = source;
  size_t input_length = length;
  if (recovered->changed) {
    cmark_node_free(document->doc);
    document->doc = NULL;
    cmark_parser_free(document->parser);
    document->parser = NULL;
    input = recovered->data;
    input_length = recovered->length;
    if (!gfm_document_parse(document, input, input_length)) {
      free(recovered);
      gfm_document_free(document);
      return NULL;
    }
  }
  int ok = gfm_adapt_nodes(document->doc, (const char *)input, input_length);
  free(recovered);
  if (!ok) {
    gfm_document_free(document);
    return NULL;
  }
  cmark_iter *iter = cmark_iter_new(document->doc);
  if (!iter) {
    gfm_document_free(document);
    return NULL;
  }
  cmark_event_type event;
  while ((event = cmark_iter_next(iter)) != CMARK_EVENT_DONE) {
    if (event != CMARK_EVENT_ENTER)
      continue;
    cmark_node *node = cmark_iter_get_node(iter);
    const char *type = cmark_node_get_type_string(node);
    if (strcmp(type, "html_inline") && strcmp(type, "html_block"))
      continue;
    if (!gfm_grow((void **)&document->raw_html, &document->raw_capacity, document->raw_count + 1,
                  sizeof(*document->raw_html))) {
      ok = 0;
      break;
    }
    document->raw_html[document->raw_count++] = node;
  }
  cmark_iter_free(iter);
  if (!ok) {
    gfm_document_free(document);
    return NULL;
  }
  return document;
}

static char *gfm_document_html(gfm_document *document) {
  return cmark_render_html(document->doc, document->options,
                           cmark_parser_get_syntax_extensions(document->parser));
}
