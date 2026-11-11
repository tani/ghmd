/* Consecutive native steps are grouped around the DOM reparse boundaries.
 * Stages return an ASCII '0' or '1' followed by HTML; '1' requests a DOM reparse.
 * Keeping this small framing in the string API avoids runtime-specific values. */
#include "cmark-api.h"

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

static char *gfm_stage_result(char *html, int reparse) {
  if (!html) return NULL;
  size_t length = strlen(html);
  char *result = length <= SIZE_MAX - 2 ? malloc(length + 2) : NULL;
  if (result) {
    result[0] = reparse ? '1' : '0';
    memcpy(result + 1, html, length + 1);
  }
  free(html);
  return result;
}

static char *gfm_prepare_html_stage(const unsigned char *source, size_t length, const char *boundary) {
  char *flattened = gfm_flatten_templates(source, length, boundary);
  if (!flattened) return NULL;
  char *details = gfm_prepare_details((const unsigned char *)flattened, strlen(flattened));
  if (!details) { free(flattened); return NULL; }
  int reparse = strcmp(flattened, details) != 0;
  free(flattened);
  if (reparse) {
    char *clean = gfm_remove_boundaries((const unsigned char *)details, strlen(details), boundary);
    free(details);
    details = clean;
  }
  return gfm_stage_result(details, reparse);
}

static char *gfm_render_math_stage(const unsigned char *source, size_t length, const char *token,
                                   const char *boundary, const char *table_marker) {
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
  return gfm_stage_result(math, reparse);
}

static char *gfm_finish_html(const unsigned char *source, size_t length, const char *token,
                              const char *boundary, const char *table_marker) {
  char *linked = gfm_process_links(source, length, boundary);
  if (!linked) return NULL;
  char *result = gfm_finalize_serialized_html((const unsigned char *)linked, strlen(linked), token, table_marker);
  free(linked);
  return result;
}
