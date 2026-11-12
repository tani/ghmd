#include "core.c"
#include <node/node_api.h>

static int h_short_string(napi_env env, napi_value value, char *out, size_t capacity) {
  size_t length;
  if (napi_get_value_string_utf8(env, value, NULL, 0, &length) != napi_ok) return 0;
  if (length >= capacity) return 0;
  return napi_get_value_string_utf8(env, value, out, capacity, &length) == napi_ok;
}

napi_value gfm_html_call(napi_env env, napi_value operation, napi_value input, napi_value mode) {
  size_t length = 0; char op[32], option[64];
  unsigned char *source = NULL; h_document document = {0}; h_buffer out = {0}; napi_value result = NULL;
  if (!h_short_string(env, operation, op, sizeof(op)) ||
      napi_get_value_string_utf8(env, input, NULL, 0, &length) != napi_ok ||
      !h_short_string(env, mode, option, sizeof(option))) { napi_throw_type_error(env, NULL, "HTML operation requires a string input and short string modifiers"); goto done; }
  if (length > 32 * 1024 * 1024) { napi_throw_range_error(env, NULL, "HTML input exceeds 32 MB"); goto done; }
  source = malloc(length + 1);
  if (!source) { napi_throw_error(env, NULL, "Cannot allocate HTML input"); goto done; }
  if (napi_get_value_string_utf8(env, input, (char *)source, length + 1, &length) != napi_ok) goto done;
  if (!h_parse(&document, source, length, !strcmp(option, "locations"))) { napi_throw_error(env, NULL, "Cannot parse HTML"); goto done; }
  if (!strcmp(op, "parse")) h_json_node(&out, &document, document.root, 1, 0);
  else if (!strcmp(op, "sanitize")) h_sanitize(&out, &document, document.root, option);
  else { napi_throw_type_error(env, NULL, "Unknown HTML operation"); goto done; }
  if (out.failed) {
    if (out.failed == H_ALLOC_FAILED) napi_throw_error(env, NULL, "Cannot allocate HTML output");
    else napi_throw_range_error(env, NULL, "HTML tree exceeds output or depth limit");
    goto done;
  }
  if (napi_create_string_utf8(env, out.data ? (const char *)out.data : "", out.length, &result) != napi_ok) result = NULL;
done:
  h_destroy(&document); free(source); free(out.data); return result;
}
