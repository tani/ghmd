import { parseFragment } from "parse5";

type Node = {
  nodeName: string;
  tagName?: string;
  namespaceURI?: string;
  value?: string;
  data?: string;
  attrs?: { name: string; value: string; namespace?: string; prefix?: string }[];
  childNodes?: Node[];
  content?: Node;
};
type Canonical = {
  type: string;
  namespace?: string;
  attrs?: [string, string][];
  children?: Canonical[];
  value?: string;
};

const blocks = new Set(
  "address article aside blockquote body caption dd details dialog div dl dt fieldset figcaption figure footer form h1 h2 h3 h4 h5 h6 header hgroup hr li main menu nav ol p pre section summary table tbody td tfoot th thead tr ul".split(
    " ",
  ),
);
const preserve = new Set(["pre", "code", "textarea", "script", "style", "math-renderer"]);
const generated = (name: string, value: string) => {
  if (name === "data-run-id" && /^[0-9a-f]{32}$/.test(value)) return "GENERATED";
  if (["id", "href"].includes(name))
    return value.replace(/(user-content-fn(?:ref)?-[^"#\s]*?-)[0-9a-f]{32}(?=$)/, "$1GENERATED");
  return value;
};

/** Compare HTML structure and attributes; collapse HTML whitespace outside literal contexts. */
export function htmlAst(html: string, options: { ignoreHighlighting?: boolean } = {}): Canonical {
  const contents = (node: Node): string =>
    node.nodeName === "#text" ? (node.value ?? "") : (node.childNodes ?? []).map(contents).join("");
  const codeBlock = (node: Node): Canonical => {
    const attrs = (node.attrs ?? [])
      .filter((a) => a.name !== "lang" && !(a.name === "class" && a.value === "notranslate"))
      .map((a): [string, string] => [a.name, generated(a.name, a.value)])
      .sort(([a], [b]) => a.localeCompare(b));
    return {
      type: "pre",
      attrs,
      children: [{ type: "#text", value: contents(node).replace(/\n$/, "") }],
    };
  };
  const highlightedChildren = (node: Node): boolean =>
    (node.childNodes ?? []).every(
      (child) =>
        child.nodeName === "#text" ||
        (child.tagName === "span" &&
          (child.attrs ?? []).every(
            (a) =>
              a.name === "class" && a.value.split(/\s+/).every((name) => name.startsWith("pl-")),
          ) &&
          highlightedChildren(child)),
    );
  function visit(node: Node, literal = false): Canonical {
    if (node.nodeName === "#text")
      return {
        type: "#text",
        value: literal ? node.value : node.value!.replace(/[\t\n\f\r ]+/g, " "),
      };
    if (node.nodeName === "#comment") return { type: "#comment", value: node.data };
    const tag = node.tagName;
    if (
      options.ignoreHighlighting &&
      tag === "div" &&
      (node.attrs ?? []).some(
        (a) => a.name === "class" && a.value.split(/\s+/).includes("highlight"),
      )
    ) {
      const children = (node.childNodes ?? []).filter(
        (child) => child.nodeName !== "#text" || child.value?.trim(),
      );
      if (
        children.length === 1 &&
        children[0].tagName === "pre" &&
        highlightedChildren(children[0])
      )
        return codeBlock(children[0]);
    }
    if (options.ignoreHighlighting && tag === "pre") {
      const children = node.childNodes ?? [];
      if (
        children.length === 1 &&
        children[0].tagName === "code" &&
        (children[0].childNodes ?? []).every((child) => child.nodeName === "#text")
      )
        return codeBlock(node);
    }
    const keep = literal || (!!tag && preserve.has(tag));
    const rawChildren = node.content?.childNodes ?? node.childNodes ?? [];
    const children = rawChildren.map((child) => visit(child, keep));
    if (!keep) {
      // Whitespace between block elements and at block edges is HTML layout formatting.
      for (let i = children.length - 1; i >= 0; i--) {
        const child = children[i];
        if (child.type !== "#text" || !child.value) continue;
        const previous = children[i - 1],
          next = children[i + 1];
        const leftBlock = !previous ? !tag || blocks.has(tag) : blocks.has(previous.type);
        const rightBlock = !next ? !tag || blocks.has(tag) : blocks.has(next.type);
        if (child.value === " " && (leftBlock || rightBlock)) children.splice(i, 1);
        else {
          if (leftBlock) child.value = child.value.replace(/^ /, "");
          if (rightBlock) child.value = child.value.replace(/ $/, "");
          if (!child.value) children.splice(i, 1);
        }
      }
    }
    const result: Canonical = { type: tag ?? node.nodeName };
    if (node.namespaceURI && node.namespaceURI !== "http://www.w3.org/1999/xhtml")
      result.namespace = node.namespaceURI;
    if (tag)
      result.attrs = (node.attrs ?? [])
        .map((a): [string, string] => [
          (a.namespace ? `${a.namespace}:` : "") + a.name,
          generated(a.name, a.value),
        ])
        .sort(([a], [b]) => a.localeCompare(b));
    if (children.length) result.children = children;
    return result;
  }
  return visit(parseFragment(html) as Node);
}
