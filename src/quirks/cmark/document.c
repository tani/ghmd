/* Complete Markdown-to-HTML operation. No JS callbacks or DOM round trips. */
static gfm_status gfm_render_document(gfm_view source, gfm_view token, gfm_view boundary, gfm_result *result) {
  if (source.length > 1000000) { result->error = "Markdown exceeds 1 MB"; return GFM_LIMIT_EXCEEDED; }
  if (token.length != 32) { result->error = "runId must be 32 lowercase hexadecimal characters"; return GFM_BAD_ARGUMENT; }
  for (size_t i = 0; i < token.length; i++) if (!((token.data[i] >= 'a' && token.data[i] <= 'f') || (token.data[i] >= '0' && token.data[i] <= '9'))) {
    result->error = "runId must be 32 lowercase hexadecimal characters"; return GFM_BAD_ARGUMENT;
  }
  if (boundary.length != 36) { result->error = "Invalid render boundary"; return GFM_BAD_ARGUMENT; }
  for (size_t i = 0; i < 36; i++) {
    unsigned char c = boundary.data[i];
    int dash = i == 8 || i == 13 || i == 18 || i == 23;
    if (dash ? c != '-' : !((c >= 'a' && c <= 'f') || (c >= '0' && c <= '9'))) { result->error = "Invalid render boundary"; return GFM_BAD_ARGUMENT; }
  }
  char *modifiers = malloc(33 + 37 + 48);
  if (!modifiers) { result->error = "Cannot allocate render identifiers"; return GFM_NO_MEMORY; }
  char *run_id = modifiers, *marker = modifiers + 33, *table_marker = modifiers + 70;
  memcpy(run_id, token.data, 32); run_id[32] = 0;
  memcpy(marker, boundary.data, 36); marker[36] = 0;
  memcpy(table_marker, "data-gfm-table-", 15); size_t used = 15;
  for (size_t i = 0; i < 36; i++) if (boundary.data[i] != '-') table_marker[used++] = boundary.data[i];
  table_marker[used] = 0;
  gfm_status status = gfm_render_string(source, 0, gfm_filter_raw_inputs, NULL, result);
  if (status != GFM_OK) { free(modifiers); return status; }
  if (result->text.length > GFM_MAX_OUTPUT_BYTES) {
    free(modifiers); free(result->text.data); result->text = (gfm_string){0};
    result->error = "Rendered HTML exceeds 32 MB"; return GFM_LIMIT_EXCEEDED;
  }
  status = GFM_NO_MEMORY;
  char *html = gfm_dom_sanitize(result->text.data, result->text.length, marker, &status);
  free(result->text.data); result->text = (gfm_string){0};
  if (!html) goto fail;
  /* Cache lengths before calls: Bun 1.4.4's ARM64 C compiler can miscompile
   * nested strlen() argument expressions in calls with several modifiers. */
  size_t length = strlen(html);
  char *parsed = gfm_dom_reparse((const unsigned char *)html, length, &status);
  free(html); html = parsed;
  if (!html) goto fail;
  int reparse = 0;
  length = strlen(html);
  char *prepared = gfm_prepare_html_stage((const unsigned char *)html, length, marker, &reparse);
  free(html); html = prepared;
  if (!html) goto fail;
  if (reparse) {
    length = strlen(html);
    char *next = gfm_dom_reparse((const unsigned char *)html, length, &status);
    free(html); html = next;
    if (!html) goto fail;
  }
  length = strlen(html);
  char *math = gfm_render_math_stage((const unsigned char *)html, length, run_id, marker, table_marker, &reparse);
  free(html); html = math;
  if (!html) goto fail;
  if (reparse) {
    length = strlen(html);
    char *next = gfm_dom_reparse((const unsigned char *)html, length, &status);
    free(html); html = next;
    if (!html) goto fail;
  }
  length = strlen(html);
  char *final = gfm_finish_html((const unsigned char *)html, length, run_id, marker, table_marker);
  free(html);
  if (!final) goto fail;
  result->text = (gfm_string){(unsigned char *)final, strlen(final)};
  free(modifiers); return gfm_complete_html(result);
fail:
  free(modifiers);
  result->error = status == GFM_LIMIT_EXCEEDED ? "HTML tree exceeds output or depth limit" : "Cannot render HTML document"; return status;
}
