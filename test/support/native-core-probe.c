/* This translation unit intentionally includes no Bun or Node-API adapter. */
#include "../../src/quirks/cmark/unity.c"

int gfm_api_probe(void) {
  const unsigned char markdown[] = {'$', 'x', '$'};
  const unsigned char run_id[] = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const unsigned char boundary[] = "01234567-89ab-cdef-0123-456789abcdef";
  gfm_view source = {markdown, sizeof(markdown)};
  gfm_view token = {run_id, sizeof(run_id) - 1}, marker = {boundary, sizeof(boundary) - 1};
  gfm_result result = {0};
  if (gfm_render(source, token, marker, &result) != GFM_OK)
    return 1;
  const char *expected =
      "<p><math-renderer class=\"js-inline-math\" style=\"display: inline-block\" data-run-id=\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\">$x$</math-renderer></p>";
  int ok = result.text.length == strlen(expected) &&
           !memcmp(result.text.data, expected, result.text.length);
  free(result.text.data);
  if (!ok)
    return 2;
  gfm_view bad_token = {(const unsigned char *)"bad", 3};
  if (gfm_render(source, bad_token, marker, &result) != GFM_BAD_ARGUMENT || result.text.data ||
      !result.error)
    return 3;
  source.length = 1000001;
  if (gfm_render(source, token, marker, &result) != GFM_LIMIT_EXCEEDED || result.text.data)
    return 4;
  source = (gfm_view){NULL, 0};
  if (gfm_render(source, token, marker, &result) != GFM_BAD_ARGUMENT || result.text.data)
    return 5;
  if (gfm_render(source, token, marker, NULL) != GFM_BAD_ARGUMENT)
    return 6;
  return 0;
}
