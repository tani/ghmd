/* This translation unit intentionally includes no Bun or Node-API adapter. */
#include "../../src/quirks/cmark-unity.c"

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

  /* An optional hole and a modifier without a terminating NUL are valid. */
  const unsigned char marker[] = {'d', 'a', 't', 'a', '-', 'x'};
  const char *table = "<table data-x=\"\"></table>";
  args[0] = (gfm_view){(const unsigned char *)table, strlen(table)};
  args[1] = (gfm_view){0};
  args[2] = (gfm_view){marker, sizeof(marker)};
  if (gfm_call("finalize", args, 3, NULL, NULL, &result) != GFM_OK) return 5;
  ok = !strcmp((const char *)result.text.data, "<table role=\"table\"></table>");
  free(result.text.data);
  if (!ok) return 6;
  if (gfm_call("unknown", args, 1, NULL, NULL, &result) != GFM_BAD_ARGUMENT || !result.error || result.text.data) return 7;
  return 0;
}
