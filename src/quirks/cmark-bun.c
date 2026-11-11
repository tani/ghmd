/* Bun-specific ownership and exception adapter. The included core has no N-API
 * dependency and is usable through cmark-api.h from an ordinary C program. */
#include "cmark-unity.c"
#include <node/node_api.h>

typedef struct { napi_env env; napi_value function; } gfm_bun_filter;

static void gfm_bun_error(napi_env env, int type_error, const char *message) {
  bool pending = false;
  if (napi_is_exception_pending(env, &pending) == napi_ok && pending) return;
  if (type_error) napi_throw_type_error(env, NULL, message);
  else napi_throw_error(env, NULL, message);
}

static int gfm_bun_read_string(napi_env env, napi_value value, gfm_string *output, const char *message) {
  size_t length = 0;
  if (napi_get_value_string_utf8(env, value, NULL, 0, &length) != napi_ok) {
    gfm_bun_error(env, 1, message); return 0;
  }
  if (length == SIZE_MAX) { gfm_bun_error(env, 0, "String is too large"); return 0; }
  output->data = malloc(length + 1);
  if (!output->data) { gfm_bun_error(env, 0, "Cannot allocate input string"); return 0; }
  if (napi_get_value_string_utf8(env, value, (char *)output->data, length + 1, &output->length) != napi_ok) {
    free(output->data); *output = (gfm_string){0};
    gfm_bun_error(env, 0, "Cannot read input string"); return 0;
  }
  return 1;
}

static gfm_status gfm_bun_filter_html(gfm_view input, gfm_string *output, void *context) {
  gfm_bun_filter *filter = context;
  napi_handle_scope scope;
  if (napi_open_handle_scope(filter->env, &scope) != napi_ok) {
    gfm_bun_error(filter->env, 0, "Cannot open raw HTML handle scope"); return GFM_FILTER_FAILED;
  }
  gfm_status status = GFM_FILTER_FAILED;
  napi_value source, result, receiver;
  if (napi_create_string_utf8(filter->env, (const char *)input.data, input.length, &source) != napi_ok ||
      napi_get_undefined(filter->env, &receiver) != napi_ok ||
      napi_call_function(filter->env, receiver, filter->function, 1, &source, &result) != napi_ok) {
    gfm_bun_error(filter->env, 0, "Raw HTML filter failed");
  } else if (gfm_bun_read_string(filter->env, result, output, "Raw HTML filter must return a string")) status = GFM_OK;
  if (napi_close_handle_scope(filter->env, scope) != napi_ok) {
    gfm_bun_error(filter->env, 0, "Cannot close raw HTML handle scope"); status = GFM_FILTER_FAILED;
  }
  return status;
}

/* Only this function crosses the Bun boundary. env is supplied by cc(); the
 * remaining arguments are strings, except the optional synchronous filter. */
napi_value gfm_bun_call(napi_env env, napi_value operation, napi_value source,
                        napi_value arg1, napi_value arg2, napi_value arg3,
                        napi_value raw_filter) {
  gfm_string name = {0}, strings[4] = {{0}};
  gfm_view arguments[4] = {{0}};
  gfm_result output = {0};
  napi_value result = NULL, values[4] = {source, arg1, arg2, arg3};
  gfm_bun_filter filter = {env, raw_filter};
  gfm_html_filter callback = NULL;
  size_t count = 0;
  if (!gfm_bun_read_string(env, operation, &name, "Operation must be a string")) goto done;
  if (memchr(name.data, 0, name.length)) { gfm_bun_error(env, 1, "Invalid native operation"); goto done; }
  for (size_t i = 0; i < 4; i++) {
    napi_valuetype type;
    if (napi_typeof(env, values[i], &type) != napi_ok) { gfm_bun_error(env, 0, "Cannot inspect argument"); goto done; }
    if (i > 0 && type == napi_undefined) continue;
    if (!gfm_bun_read_string(env, values[i], &strings[i], "Native arguments must be strings")) goto done;
    arguments[i] = (gfm_view){strings[i].data, strings[i].length};
    count = i + 1;
  }
  napi_valuetype filter_type;
  if (napi_typeof(env, raw_filter, &filter_type) != napi_ok) { gfm_bun_error(env, 0, "Cannot inspect raw HTML filter"); goto done; }
  if (filter_type == napi_function) callback = gfm_bun_filter_html;
  else if (filter_type != napi_undefined && filter_type != napi_null) { gfm_bun_error(env, 1, "Raw HTML filter must be a function"); goto done; }
  gfm_status status = gfm_call((const char *)name.data, arguments, count, callback, &filter, &output);
  if (status != GFM_OK) { gfm_bun_error(env, status == GFM_BAD_ARGUMENT, output.error ? output.error : "Native operation failed"); goto done; }
  if (napi_create_string_utf8(env, (const char *)output.text.data, output.text.length, &result) != napi_ok) {
    result = NULL; gfm_bun_error(env, 0, "Cannot create output string");
  }
done:
  free(name.data);
  for (size_t i = 0; i < 4; i++) free(strings[i].data);
  free(output.text.data);
  return result;
}
