/* This translation unit intentionally includes no Bun or Node-API adapter. */
#include "../../src/quirks/cmark/unity.c"

static gfm_status test_filter(gfm_view input, gfm_string *output, void *context) {
  (void)input;
  (*(int *)context)++;
  *output = (gfm_string){0};
  return GFM_OK;
}

static gfm_status failing_filter(gfm_view input, gfm_string *output, void *context) {
  (void)input; (void)context;
  output->data = malloc(1);
  if (!output->data) return GFM_NO_MEMORY;
  output->data[0] = 0; output->length = 0;
  return GFM_FILTER_FAILED;
}

int gfm_api_probe(void) {
  gfm_result result = {0};
  const unsigned char text[] = {'A', 0, 'B'};
  gfm_view args[3] = {{text, sizeof(text)}};
  if (gfm_call("lowercase", args, 1, NULL, NULL, &result) != GFM_OK) return 1;
  int ok = result.text.length == 3 && !memcmp(result.text.data, "a\0b", 3);
  free(result.text.data);
  if (!ok) return 2;

  int count = 0;
  args[0] = (gfm_view){(const unsigned char *)"<i>x</i>", 8};
  if (gfm_call("render", args, 1, test_filter, &count, &result) != GFM_OK) return 3;
  ok = count == 2 && result.text.length == 9 && !memcmp(result.text.data, "<p>x</p>\n", 9);
  free(result.text.data);
  if (!ok) return 4;
  if (gfm_call("render", args, 1, failing_filter, NULL, &result) != GFM_FILTER_FAILED || !result.error || result.text.data) return 8;

  /* Full rendering works from non-NUL-terminated views, without Bun, a DOM
   * runtime, callbacks, or an intermediate HTML API. */
  const unsigned char markdown[] = {'$', 'x', '$'};
  const unsigned char run_id[] = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const unsigned char boundary[] = "01234567-89ab-cdef-0123-456789abcdef";
  args[0] = (gfm_view){markdown, sizeof(markdown)};
  args[1] = (gfm_view){run_id, sizeof(run_id) - 1};
  args[2] = (gfm_view){boundary, sizeof(boundary) - 1};
  if (gfm_call("document", args, 3, failing_filter, NULL, &result) != GFM_OK) return 5;
  const char *expected = "<p><math-renderer class=\"js-inline-math\" style=\"display: inline-block\" data-run-id=\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\">$x$</math-renderer></p>";
  ok = result.text.length == strlen(expected) && !memcmp(result.text.data, expected, result.text.length);
  free(result.text.data);
  if (!ok) return 6;
  args[1] = (gfm_view){(const unsigned char *)"bad", 3};
  if (gfm_call("document", args, 3, NULL, NULL, &result) != GFM_BAD_ARGUMENT || result.text.data) return 10;
  args[1] = (gfm_view){run_id, sizeof(run_id) - 1};
  args[0].length = 1000001;
  if (gfm_call("document", args, 3, NULL, NULL, &result) != GFM_LIMIT_EXCEEDED || result.text.data) return 11;
  if (gfm_call("unknown", args, 1, NULL, NULL, &result) != GFM_BAD_ARGUMENT || !result.error || result.text.data) return 7;
  const char *removed[] = {"finalize", "math", "links", "details", "templates", "boundaries", "mathTables", "stream", "decode", "prepareStage", "mathStage", "finish"};
  for (size_t i = 0; i < sizeof(removed) / sizeof(removed[0]); i++) {
    if (gfm_call(removed[i], args, 3, NULL, NULL, &result) != GFM_BAD_ARGUMENT ||
        !result.error || strcmp(result.error, "Unknown native operation") || result.text.data) return 9;
  }
  return 0;
}
