/* Test-only access to private C helpers. Production exposes no stage calls. */
#include "../../src/quirks/cmark/bun.c"

napi_value test_gfm_stage(napi_env env, napi_value operation, napi_value source,
                         napi_value a, napi_value b, napi_value c) {
  gfm_string name = {0}, inputs[4] = {{0}};
  napi_value values[4] = {source, a, b, c}, value = NULL;
  gfm_result result = {0}; char *html = NULL; int reparse = 0;
  if (!gfm_bun_read_string(env, operation, &name, "Expected operation")) goto done;
  size_t count = !strcmp((char *)name.data, "prepare") ? 2 : 4;
  for (size_t i = 0; i < count; i++)
    if (!gfm_bun_read_string(env, values[i], &inputs[i], "Expected string modifier")) goto done;
  if (!strcmp((char *)name.data, "prepare"))
    html = gfm_prepare_html_stage(inputs[0].data, inputs[0].length, (char *)inputs[1].data, &reparse);
  else if (!strcmp((char *)name.data, "math"))
    html = gfm_render_math_stage(inputs[0].data, inputs[0].length, (char *)inputs[1].data, (char *)inputs[2].data, (char *)inputs[3].data, &reparse);
  else if (!strcmp((char *)name.data, "finish"))
    html = gfm_finish_html(inputs[0].data, inputs[0].length, (char *)inputs[1].data, (char *)inputs[2].data, (char *)inputs[3].data);
  else { gfm_bun_error(env, GFM_BUN_TYPE_ERROR, "Unknown test helper"); goto done; }
  if (!html) { gfm_bun_error(env, GFM_BUN_ERROR, "Cannot allocate helper output"); goto done; }
  if (!strcmp((char *)name.data, "finish")) {
    result.text = (gfm_string){(unsigned char *)html, strlen(html)}; html = NULL;
    if (gfm_complete_html(&result) != GFM_OK) { gfm_bun_error(env, GFM_BUN_RANGE_ERROR, result.error); goto done; }
  } else {
    size_t size = strlen(html);
    result.text.data = malloc(size + 2); result.text.length = size + 1;
    if (!result.text.data) { gfm_bun_error(env, GFM_BUN_ERROR, "Cannot allocate test framing"); goto done; }
    result.text.data[0] = reparse ? '1' : '0'; memcpy(result.text.data + 1, html, size + 1);
  }
  if (napi_create_string_utf8(env, (char *)result.text.data, result.text.length, &value) != napi_ok) value = NULL;
done:
  free(name.data); for (size_t i = 0; i < 4; i++) free(inputs[i].data);
  free(html); free(result.text.data); return value;
}
