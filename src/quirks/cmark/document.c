/* Complete Markdown-to-HTML operation and ownership of intermediate strings. */
enum { GFM_MAX_SOURCE_BYTES = 1000000, GFM_RUN_ID_LENGTH = 32, GFM_BOUNDARY_LENGTH = 36 };
#define GFM_TABLE_MARKER_PREFIX "data-gfm-table-"
typedef struct {
  char run_id[GFM_RUN_ID_LENGTH + 1];
  char boundary[GFM_BOUNDARY_LENGTH + 1];
  char table_marker[sizeof(GFM_TABLE_MARKER_PREFIX) + GFM_RUN_ID_LENGTH];
} gfm_render_ids;

static void gfm_string_clear(gfm_string *string) {
  free(string->data);
  *string = (gfm_string){0};
}
/* Replacement takes a newly allocated string, even when the next step failed.
 * Keep lengths cached: Bun 1.4.4's ARM64 compiler can miscompile nested strlen()
 * expressions in calls that also pass several string modifiers. */
static int gfm_string_replace(gfm_string *string, char *replacement) {
  gfm_string_clear(string);
  string->data = (unsigned char *)replacement;
  if (replacement)
    string->length = strlen(replacement);
  return replacement != NULL;
}
static int gfm_is_hex(unsigned char c) {
  return (c >= 'a' && c <= 'f') || (c >= '0' && c <= '9');
}
static gfm_status gfm_validate_render(gfm_view source, gfm_view token, gfm_view boundary,
                                      gfm_result *result) {
  if (!source.data || !token.data || !boundary.data) {
    result->error = "Required string argument is missing";
    return GFM_BAD_ARGUMENT;
  }
  if (source.length > GFM_MAX_SOURCE_BYTES) {
    result->error = "Markdown exceeds 1 MB";
    return GFM_LIMIT_EXCEEDED;
  }
  if (token.length != GFM_RUN_ID_LENGTH)
    goto invalid_token;
  for (size_t i = 0; i < token.length; i++)
    if (!gfm_is_hex(token.data[i]))
      goto invalid_token;
  if (boundary.length != GFM_BOUNDARY_LENGTH)
    goto invalid_boundary;
  for (size_t i = 0; i < boundary.length; i++) {
    int dash = i == 8 || i == 13 || i == 18 || i == 23;
    if (dash ? boundary.data[i] != '-' : !gfm_is_hex(boundary.data[i]))
      goto invalid_boundary;
  }
  return GFM_OK;
invalid_token:
  result->error = "runId must be 32 lowercase hexadecimal characters";
  return GFM_BAD_ARGUMENT;
invalid_boundary:
  result->error = "Invalid render boundary";
  return GFM_BAD_ARGUMENT;
}
static gfm_render_ids *gfm_render_ids_new(gfm_view token, gfm_view boundary) {
  gfm_render_ids *ids = malloc(sizeof(*ids));
  if (!ids)
    return NULL;
  memcpy(ids->run_id, token.data, GFM_RUN_ID_LENGTH);
  ids->run_id[GFM_RUN_ID_LENGTH] = 0;
  memcpy(ids->boundary, boundary.data, GFM_BOUNDARY_LENGTH);
  ids->boundary[GFM_BOUNDARY_LENGTH] = 0;
  size_t used = sizeof(GFM_TABLE_MARKER_PREFIX) - 1;
  memcpy(ids->table_marker, GFM_TABLE_MARKER_PREFIX, used);
  for (size_t i = 0; i < boundary.length; i++)
    if (boundary.data[i] != '-')
      ids->table_marker[used++] = boundary.data[i];
  ids->table_marker[used] = 0;
  return ids;
}
static int gfm_reparse_html(gfm_string *html, gfm_status *status) {
  char *parsed = gfm_dom_reparse(html->data, html->length, status);
  return gfm_string_replace(html, parsed);
}
static gfm_status gfm_complete_html(gfm_result *result) {
  gfm_string *html = &result->text;
  if (html->length && html->data[html->length - 1] == '\n')
    html->data[--html->length] = 0;
  if (html->length > GFM_MAX_OUTPUT_BYTES) {
    gfm_string_clear(html);
    result->error = "Rendered HTML exceeds 32 MB";
    return GFM_LIMIT_EXCEEDED;
  }
  return GFM_OK;
}

gfm_status gfm_render(gfm_view source, gfm_view token, gfm_view boundary, gfm_result *result) {
  if (!result)
    return GFM_BAD_ARGUMENT;
  *result = (gfm_result){0};
  gfm_status status = gfm_validate_render(source, token, boundary, result);
  if (status != GFM_OK)
    return status;
  gfm_render_ids *ids = gfm_render_ids_new(token, boundary);
  if (!ids) {
    result->error = "Cannot allocate render identifiers";
    return GFM_NO_MEMORY;
  }
  gfm_string html = {0};
  status = gfm_render_markdown(source, result);
  if (status != GFM_OK)
    goto done;
  html = result->text;
  result->text = (gfm_string){0};
  if (html.length > GFM_MAX_OUTPUT_BYTES) {
    result->error = "Rendered HTML exceeds 32 MB";
    status = GFM_LIMIT_EXCEEDED;
    goto done;
  }
  status = GFM_NO_MEMORY;
  char *next = gfm_dom_sanitize(html.data, html.length, ids->boundary, &status);
  if (!gfm_string_replace(&html, next) || !gfm_reparse_html(&html, &status))
    goto failed_html;
  int reparse = 0;
  next = gfm_prepare_html_stage(html.data, html.length, ids->boundary, &reparse);
  if (!gfm_string_replace(&html, next) || (reparse && !gfm_reparse_html(&html, &status)))
    goto failed_html;
  next = gfm_render_math_stage(html.data, html.length, ids->run_id, ids->boundary,
                               ids->table_marker, &reparse);
  if (!gfm_string_replace(&html, next) || (reparse && !gfm_reparse_html(&html, &status)))
    goto failed_html;
  next = gfm_finish_html(html.data, html.length, ids->run_id, ids->boundary, ids->table_marker);
  if (!gfm_string_replace(&html, next))
    goto failed_html;
  result->text = html;
  html = (gfm_string){0};
  status = gfm_complete_html(result);
  goto done;
failed_html:
  result->error = status == GFM_LIMIT_EXCEEDED ? "HTML tree exceeds output or depth limit"
                                               : "Cannot render HTML document";
done:
  gfm_string_clear(&html);
  if (status != GFM_OK)
    gfm_string_clear(&result->text);
  free(ids);
  return status;
}
