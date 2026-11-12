#include "core.c"
#include <node/node_api.h>

/* Comparison/source mapping needs parsing only. Sanitization stays in C. */
napi_value gfm_html_parse(napi_env env, napi_value input, napi_value locations) {
  size_t length = 0;
  bool track = false;
  unsigned char *source = NULL;
  h_document document = {0};
  h_buffer out = {0};
  napi_value result = NULL;
  if (napi_get_value_string_utf8(env, input, NULL, 0, &length) != napi_ok ||
      napi_get_value_bool(env, locations, &track) != napi_ok) {
    napi_throw_type_error(env, NULL, "HTML parser expects a string and a boolean");
    goto done;
  }
  if (length > 32 * 1024 * 1024) {
    napi_throw_range_error(env, NULL, "HTML input exceeds 32 MB");
    goto done;
  }
  source = malloc(length + 1);
  if (!source) {
    napi_throw_error(env, NULL, "Cannot allocate HTML input");
    goto done;
  }
  if (napi_get_value_string_utf8(env, input, (char *)source, length + 1, &length) != napi_ok)
    goto done;
  if (!h_parse(&document, source, length, track ? H_TRACK_UTF16 : H_TRACK_NONE)) {
    napi_throw_error(env, NULL, "Cannot parse HTML");
    goto done;
  }
  h_json_node(&out, &document, document.root, 1, 0);
  if (out.failed) {
    if (out.failed == H_ALLOC_FAILED)
      napi_throw_error(env, NULL, "Cannot allocate HTML output");
    else
      napi_throw_range_error(env, NULL, "HTML tree exceeds output or depth limit");
    goto done;
  }
  if (napi_create_string_utf8(env, out.data ? (const char *)out.data : "", out.length, &result) !=
      napi_ok)
    result = NULL;
done:
  h_destroy(&document);
  free(source);
  free(out.data);
  return result;
}
