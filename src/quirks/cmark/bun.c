/* Bun-specific ownership and exception adapter. The included core has no N-API
 * dependency and is usable through api.h from an ordinary C program. */
#include "unity.c"
#include <node/node_api.h>

typedef enum { GFM_BUN_ERROR, GFM_BUN_TYPE_ERROR, GFM_BUN_RANGE_ERROR } gfm_bun_error_kind;

static void gfm_bun_error(napi_env env, gfm_bun_error_kind kind, const char *message) {
  bool pending = false;
  if (napi_is_exception_pending(env, &pending) == napi_ok && pending) return;
  if (kind == GFM_BUN_TYPE_ERROR) napi_throw_type_error(env, NULL, message);
  else if (kind == GFM_BUN_RANGE_ERROR) napi_throw_range_error(env, NULL, message);
  else napi_throw_error(env, NULL, message);
}

static int gfm_bun_read_string(napi_env env, napi_value value, gfm_string *output, const char *message) {
  size_t length = 0;
  if (napi_get_value_string_utf8(env, value, NULL, 0, &length) != napi_ok) {
    gfm_bun_error(env, GFM_BUN_TYPE_ERROR, message); return 0;
  }
  if (length == SIZE_MAX) { gfm_bun_error(env, GFM_BUN_ERROR, "String is too large"); return 0; }
  output->data = malloc(length + 1);
  if (!output->data) { gfm_bun_error(env, GFM_BUN_ERROR, "Cannot allocate input string"); return 0; }
  if (napi_get_value_string_utf8(env, value, (char *)output->data, length + 1, &output->length) != napi_ok) {
    free(output->data); *output = (gfm_string){0};
    gfm_bun_error(env, GFM_BUN_ERROR, "Cannot read input string"); return 0;
  }
  return 1;
}

/* The single Bun entry point accepts strings and returns final HTML. */
napi_value gfm_bun_render(napi_env env, napi_value source, napi_value run_id, napi_value boundary) {
  gfm_string inputs[3] = {{0}};
  napi_value values[3] = {source, run_id, boundary}, result = NULL;
  gfm_result output = {0};
  for (size_t i = 0; i < 3; i++)
    if (!gfm_bun_read_string(env, values[i], &inputs[i], "Render arguments must be strings")) goto done;
  gfm_view markdown = {inputs[0].data, inputs[0].length};
  gfm_view token = {inputs[1].data, inputs[1].length};
  gfm_view marker = {inputs[2].data, inputs[2].length};
  gfm_status status = gfm_render(markdown, token, marker, &output);
  if (status != GFM_OK) {
    gfm_bun_error_kind kind = status == GFM_BAD_ARGUMENT ? GFM_BUN_TYPE_ERROR :
                             status == GFM_LIMIT_EXCEEDED ? GFM_BUN_RANGE_ERROR : GFM_BUN_ERROR;
    gfm_bun_error(env, kind, output.error ? output.error : "Native render failed");
    goto done;
  }
  if (napi_create_string_utf8(env, (const char *)output.text.data, output.text.length, &result) != napi_ok) {
    result = NULL; gfm_bun_error(env, GFM_BUN_ERROR, "Cannot create output string");
  }
done:
  for (size_t i = 0; i < 3; i++) free(inputs[i].data);
  free(output.text.data); return result;
}
