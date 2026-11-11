import { parseFragment } from "parse5";

/** Remove source-authored input elements while preserving other raw HTML bytes. */
export function removeRawInputs(html: string): string {
  if (!/<input(?:\s|\/?>)/i.test(html)) return html;
  const tree: any = parseFragment(html, { sourceCodeLocationInfo: true });
  const spans: { start: number; end: number }[] = [];
  const visit = (node: any) => {
    if (node.tagName === "input" && node.sourceCodeLocation?.startTag) {
      const location = node.sourceCodeLocation.startTag;
      spans.push({ start: location.startOffset, end: location.endOffset });
    }
    for (const child of node.childNodes ?? []) visit(child);
    if (node.content) visit(node.content);
  };
  visit(tree);
  spans.sort((a, b) => b.start - a.start);
  for (const span of spans) html = html.slice(0, span.start) + html.slice(span.end);
  return html;
}
