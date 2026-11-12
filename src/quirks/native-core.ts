import { cc } from "bun:ffi";
import { join } from "node:path";

const root = join(import.meta.dir, "../..");
const { symbols } = cc({
  source: join(root, "src/quirks/cmark/bun.c"),
  include: [
    join(root, "vendor/lexbor/compat"),
    ...["src", "extensions", "generated"].map((dir) => join(root, "vendor/cmark-gfm", dir)),
  ],
  symbols: {
    gfm_bun_call: {
      args: ["napi_env", "napi_value", "napi_value", "napi_value", "napi_value", "napi_value"],
      returns: "napi_value",
    },
  },
});

export type RawHtmlFilter = (html: string) => string;
// Bun 1.4.4 keeps an argument slot for the automatically supplied napi_env.
const call = (
  operation: string,
  source: string,
  a?: string,
  b?: string,
  filter?: RawHtmlFilter,
): string => symbols.gfm_bun_call(undefined, operation, source, a, b, filter) as string;

export const nativeVersion = call("version", "");
export const renderMarkdown = (source: string, filter?: RawHtmlFilter): string =>
  call("render", source, undefined, undefined, filter);
/** Diagnostic XML is generated only when requested. */
export const renderAstXml = (source: string, filter?: RawHtmlFilter): string =>
  call("ast", source, undefined, undefined, filter);
export const unicodeLower = (source: string): string => call("lowercase", source);
export const renderDocument = (source: string, token: string, boundary: string): string =>
  call("document", source, token, boundary);

export const removeRawInputs = (html: string): string => call("rawInputs", html);
