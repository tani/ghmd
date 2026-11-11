import { cc } from "bun:ffi";
import { join } from "node:path";

const root = join(import.meta.dir, "../..");
const { symbols } = cc({
  source: join(root, "src/quirks/cmark-bun.c"),
  include: ["src", "extensions", "generated"].map((dir) => join(root, "vendor/cmark-gfm", dir)),
  symbols: {
    gfm_bun_call: {
      args: [
        "napi_env",
        "napi_value",
        "napi_value",
        "napi_value",
        "napi_value",
        "napi_value",
        "napi_value",
      ],
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
  c?: string,
  filter?: RawHtmlFilter,
): string => symbols.gfm_bun_call(undefined, operation, source, a, b, c, filter) as string;

export const nativeVersion = call("version", "");
export const renderMarkdown = (source: string, filter?: RawHtmlFilter): string =>
  call("render", source, undefined, undefined, undefined, filter);
/** Diagnostic XML is generated only when requested. */
export const renderAstXml = (source: string, filter?: RawHtmlFilter): string =>
  call("ast", source, undefined, undefined, undefined, filter);
export const unicodeLower = (source: string): string => call("lowercase", source);
export const decodeHtml = (source: string): string => call("decode", source);
export const finalizeHtml = (html: string, token?: string, tableMarker?: string): string =>
  call("finalize", html, token, tableMarker);
export const renderMathHtml = (html: string, token: string, boundary: string): string =>
  call("math", html, token, boundary);
export const processLinks = (html: string, marker: string): string => call("links", html, marker);
export const prepareDetails = (html: string): string => call("details", html);
export const flattenTemplates = (html: string, boundary: string): string =>
  call("templates", html, boundary);
export const removeBoundaries = (html: string, boundary: string): string =>
  call("boundaries", html, boundary);
export const prepareMathTables = (
  html: string,
  token: string,
  boundary: string,
  tableMarker: string,
): string => call("mathTables", html, token, boundary, tableMarker);
