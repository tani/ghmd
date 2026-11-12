/** Adapt the env/value ABI used by this project to real Node-API callbacks.
 * Slot zero remains the ignored napi_env placeholder used by Bun 1.4.4. */
export function nodeApiWrappers(symbols: { name: string; args: string[] }[]): string {
  const callbacks = symbols.map(({ name, args }, index) => {
    const values = args.map((type, i) => (type === "napi_env" ? "env" : `argv[${i}]`));
    return `
static napi_value gfm_cc_callback_${index}(napi_env env, napi_callback_info info) {
  napi_value argv[${args.length}];
  size_t argc = ${args.length};
  if (napi_get_cb_info(env, info, &argc, argv, NULL, NULL) != napi_ok) return NULL;
  return ${name}(${values.join(", ")});
}`;
  });
  const properties = symbols.map(
    ({ name }, index) =>
      `{ "${name}", NULL, gfm_cc_callback_${index}, NULL, NULL, NULL, napi_default, NULL }`,
  );
  return `
#include <node/node_api.h>
${callbacks.join("\n")}
NAPI_MODULE_INIT() {
  const napi_property_descriptor properties[] = { ${properties.join(",\n")} };
  if (napi_define_properties(env, exports, ${symbols.length}, properties) != napi_ok) return NULL;
  return exports;
}
`;
}
