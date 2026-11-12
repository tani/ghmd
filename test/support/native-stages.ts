import { cc } from "bun:ffi";
import { join } from "node:path";
const root = join(import.meta.dir, "../..");
const native = cc({
  source: join(import.meta.dir, "native-stages.c"),
  include: [
    join(root, "vendor/lexbor/compat"),
    ...["src", "extensions", "generated"].map((dir) => join(root, "vendor/cmark-gfm", dir)),
  ],
  symbols: {
    test_gfm_stage: {
      args: ["napi_env", "napi_value", "napi_value", "napi_value", "napi_value", "napi_value"],
      returns: "napi_value",
    },
  },
});
const call = (name: string, source: string, a: string, b?: string, c?: string): string =>
  native.symbols.test_gfm_stage(undefined, name, source, a, b, c) as string;
const stage = (text: string) => ({ html: text.slice(1), needsReparse: text[0] === "1" });
export const prepareHtmlStage = (source: string, boundary: string) =>
  stage(call("prepare", source, boundary));
export const renderMathStage = (source: string, token: string, boundary: string, table: string) =>
  stage(call("math", source, token, boundary, table));
export const finishHtml = (source: string, token: string, boundary: string, table: string) =>
  call("finish", source, token, boundary, table);
