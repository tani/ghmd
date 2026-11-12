import { cc } from "bun:ffi";
import { join } from "node:path";
// HTML goes in and a JSON tree comes back; locations are optional. Lexbor's
// native tree and buffers are destroyed before the call returns.
const root = join(import.meta.dir, "../..");
const { symbols } = cc({
  source: join(import.meta.dir, "native.c"),
  include: [join(root, "vendor/lexbor/compat")],
  symbols: {
    gfm_html_parse: {
      args: ["napi_env", "napi_value", "napi_value"],
      returns: "napi_value",
    },
  },
});
export interface SourceRange {
  startOffset: number;
  endOffset: number;
}
export interface ElementLocation {
  startOffset: number;
  startTag: SourceRange;
  endTag?: SourceRange;
  attrs: Record<string, SourceRange>;
}
export type HtmlNode = {
  nodeName: string;
  tagName?: string;
  namespaceURI?: string;
  attrs?: { name: string; value: string; namespace?: string; prefix?: string }[];
  childNodes?: HtmlNode[];
  content?: HtmlNode;
  parentNode?: HtmlNode;
  value?: string;
  data?: string;
  sourceCodeLocation?: ElementLocation;
};
export function parseFragment(
  html: string,
  options: { sourceCodeLocationInfo?: boolean } = {},
): HtmlNode {
  const tree = JSON.parse(
    symbols.gfm_html_parse(undefined, html, Boolean(options.sourceCodeLocationInfo)) as string,
  );
  const parents = (node: HtmlNode) => {
    for (const child of node.childNodes ?? []) {
      Object.defineProperty(child, "parentNode", {
        value: node,
        writable: true,
        configurable: true,
      });
      parents(child);
    }
    if (node.content) parents(node.content);
  };
  parents(tree);
  return tree;
}
const escapeText = (text: string) =>
  text
    .replaceAll("&", "&amp;")
    .replaceAll("\u00a0", "&nbsp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;");
const escapeAttribute = (text: string) =>
  text.replaceAll("&", "&amp;").replaceAll("\u00a0", "&nbsp;").replaceAll('"', "&quot;");
const voids = new Set(
  "area base br col embed hr img input keygen link meta param source track wbr".split(" "),
);
const literal = new Set("style script xmp iframe noembed noframes plaintext noscript".split(" "));
export function serialize(node: HtmlNode): string {
  const visit = (node: HtmlNode, raw = false): string => {
    if (node.nodeName === "#text") return raw ? (node.value ?? "") : escapeText(node.value ?? "");
    if (node.nodeName === "#comment") return `<!--${node.data ?? ""}-->`;
    if (!node.tagName) return (node.childNodes ?? []).map((child) => visit(child)).join("");
    const tag = node.tagName;
    const attrs = (node.attrs ?? [])
      .map((a) => ` ${a.prefix ? a.prefix + ":" : ""}${a.name}="${escapeAttribute(a.value)}"`)
      .join("");
    const open = `<${tag}${attrs}>`;
    if (node.namespaceURI === "http://www.w3.org/1999/xhtml" && voids.has(tag)) return open;
    return (
      open +
      (node.content?.childNodes ?? node.childNodes ?? [])
        .map((child) =>
          visit(child, node.namespaceURI === "http://www.w3.org/1999/xhtml" && literal.has(tag)),
        )
        .join("") +
      `</${tag}>`
    );
  };
  // Match fragment serialization: serialize children rather than the node itself.
  return (node.content?.childNodes ?? node.childNodes ?? [])
    .map((child) =>
      visit(
        child,
        node.namespaceURI === "http://www.w3.org/1999/xhtml" && literal.has(node.tagName ?? ""),
      ),
    )
    .join("");
}
