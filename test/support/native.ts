import { cc } from "bun:ffi";
import { join } from "node:path";
const root = join(import.meta.dir, "../..");
const native = cc({
  source: join(import.meta.dir, "native.c"),
  include: [
    join(root, "vendor/lexbor/compat"),
    ...["src", "extensions", "generated"].map((dir) => join(root, "vendor/cmark-gfm", dir)),
  ],
  symbols: {
    test_gfm_helper: {
      args: ["napi_env", "napi_value", "napi_value", "napi_value", "napi_value", "napi_value"],
      returns: "napi_value",
    },
  },
});
const call = (name: string, source: string, a?: string, b?: string, c?: string): string =>
  native.symbols.test_gfm_helper(undefined, name, source, a, b, c) as string;
const stage = (text: string) => ({ html: text.slice(1), needsReparse: text[0] === "1" });
export const prepareHtmlStage = (source: string, boundary: string) =>
  stage(call("prepare", source, boundary));
export const renderMathStage = (source: string, token: string, boundary: string, table: string) =>
  stage(call("math", source, token, boundary, table));
export const finishHtml = (source: string, token: string, boundary: string, table: string) =>
  call("finish", source, token, boundary, table);

export const renderMarkdown = (source: string): string => call("markdown", source);
export const renderAstXml = (source: string): string => call("ast", source);
export const unicodeLower = (source: string): string => call("lowercase", source);
export const removeRawInputs = (source: string): string => call("rawInputs", source);
export const sanitizeHtml = (source: string, boundary = ""): string =>
  call("sanitize", source, boundary);
export const nativeVersion = call("version", "");
