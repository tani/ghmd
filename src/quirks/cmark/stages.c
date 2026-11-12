/* Internal HTML preparation. Reparse decisions remain within the C pipeline;
 * no stage framing or intermediate HTML crosses the runtime boundary. */
#include "api.h"

static gfm_status gfm_complete_html(gfm_result *result) {
  size_t length = result->text.length;
  if (length && result->text.data[length - 1] == '\n') {
    result->text.data[--length] = 0;
    result->text.length = length;
  }
  if (length > GFM_MAX_OUTPUT_BYTES) {
    free(result->text.data);
    result->text = (gfm_string){0};
    result->error = "Rendered HTML exceeds 32 MB";
    return GFM_LIMIT_EXCEEDED;
  }
  return GFM_OK;
}

static char *gfm_prepare_html_stage(const unsigned char *source, size_t length, const char *boundary, int *needs_reparse) {
  *needs_reparse = 0;
  char *flattened = gfm_flatten_templates(source, length, boundary);
  if (!flattened) return NULL;
  size_t length_flattened = strlen(flattened);
  char *details = gfm_prepare_details((const unsigned char *)flattened, length_flattened);
  if (!details) { free(flattened); return NULL; }
  int reparse = strcmp(flattened, details) != 0;
  free(flattened);
  if (reparse) {
    size_t length_details = strlen(details);
    char *clean = gfm_remove_boundaries((const unsigned char *)details, length_details, boundary);
    free(details);
    details = clean;
  }
  *needs_reparse = reparse; return details;
}

static char *gfm_render_math_stage(const unsigned char *source, size_t length, const char *token,
                                   const char *boundary, const char *table_marker, int *needs_reparse) {
  *needs_reparse = 0;
  char *marked = gfm_prepare_math_tables(source, length, token, boundary, table_marker);
  if (!marked) return NULL;
  size_t marked_length = strlen(marked);
  char *plain = gfm_remove_boundaries((const unsigned char *)marked, marked_length, boundary);
  if (!plain) { free(marked); return NULL; }
  char *math = gfm_math_dom((const unsigned char *)marked, marked_length, token, boundary);
  free(marked);
  if (!math) { free(plain); return NULL; }
  int reparse = strcmp(plain, math) != 0;
  free(plain);
  *needs_reparse = reparse; return math;
}

static char *gfm_finish_html(const unsigned char *source, size_t length, const char *token,
                              const char *boundary, const char *table_marker) {
  char *linked = gfm_process_links(source, length, boundary);
  if (!linked) return NULL;
  size_t linked_length = strlen(linked);
  char *result = gfm_finalize_serialized_html((const unsigned char *)linked, linked_length, token, table_marker);
  free(linked);
  return result;
}
