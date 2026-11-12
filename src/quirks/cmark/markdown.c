/* Private Markdown parsing and raw-node filtering. */
static gfm_status gfm_parse_markdown(gfm_view source, gfm_document **output, gfm_result *result) {
  *output = NULL;
  const int options =
      CMARK_OPT_UNSAFE | CMARK_OPT_FOOTNOTES | CMARK_OPT_GITHUB_PRE_LANG | CMARK_OPT_HARDBREAKS;
  cmark_gfm_core_extensions_ensure_registered();
  gfm_document *document = gfm_document_new(source.data, source.length, options);
  if (!document) {
    result->error = "Cannot parse Markdown";
    return GFM_NO_MEMORY;
  }
  gfm_status status = GFM_OK;
  for (size_t i = 0; i < document->raw_count; i++) {
    cmark_node *node = document->raw_html[i];
    const char *original = cmark_node_get_literal(node);
    gfm_view input = {(const unsigned char *)original, strlen(original)};
    gfm_string filtered = {0};
    status = gfm_filter_raw_inputs(input, &filtered);
    if (status == GFM_OK) {
      int changed = input.length != filtered.length ||
                    (input.length && memcmp(input.data, filtered.data, input.length));
      if (changed && !cmark_node_set_literal(node, (const char *)filtered.data)) {
        result->error = "Cannot replace raw HTML";
        status = GFM_NO_MEMORY;
      }
    }
    free(filtered.data);
    if (status != GFM_OK) {
      if (!result->error)
        result->error = "Cannot filter raw HTML";
      gfm_document_free(document);
      return status;
    }
  }
  *output = document;
  return GFM_OK;
}
static gfm_status gfm_render_markdown(gfm_view source, gfm_result *result) {
  gfm_document *document = NULL;
  gfm_status status = gfm_parse_markdown(source, &document, result);
  if (status != GFM_OK)
    return status;
  result->text.data = (unsigned char *)gfm_document_html(document);
  gfm_document_free(document);
  if (!result->text.data) {
    result->error = "Cannot render Markdown";
    return GFM_NO_MEMORY;
  }
  result->text.length = strlen((const char *)result->text.data);
  return GFM_OK;
}
