import { cc } from "bun:ffi";
import { join } from "node:path";

const root = join(import.meta.dir, "../..");
const { symbols } = cc({
  source: join(root, "src/quirks/cmark/bun.c"),
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
export interface HtmlStage {
  html: string;
  needsReparse: boolean;
}
const stage = (result: string): HtmlStage => {
  const flag = result[0];
  if (flag !== "0" && flag !== "1") throw Error("Invalid native HTML stage result");
  return { html: result.slice(1), needsReparse: flag === "1" };
};
export const prepareHtmlStage = (html: string, boundary: string): HtmlStage =>
  stage(call("prepareStage", html, boundary));
export const renderMathStage = (
  html: string,
  token: string,
  boundary: string,
  tableMarker: string,
): HtmlStage => stage(call("mathStage", html, token, boundary, tableMarker));
export const finishHtml = (
  html: string,
  token: string,
  boundary: string,
  tableMarker: string,
): string => call("finish", html, token, boundary, tableMarker);

/** Experimental pipeline with no intervening DOM reparsing. */
export const postprocessStreamingHtml = (
  html: string,
  token: string,
  boundary: string,
  tableMarker: string,
): string => call("stream", html, token, boundary, tableMarker);
