import { createDOM } from "./dom";
import {
  finalizeHtml,
  flattenTemplates,
  nativeRender,
  prepareDetails,
  prepareMathTables,
  processLinks,
  removeBoundaries,
  renderMathHtml,
} from "./native";
import { parseFragment, serialize } from "parse5";
import createDOMPurify from "dompurify";
import type { RenderOptions } from "./index";

// GitHub assigns a unique suffix to footnote anchors for each rendered fragment.
export async function renderInProcess(source: string, options: RenderOptions): Promise<string> {
  const { html } = nativeRender(source);
  const window = createDOM();
  const document = window.document;
  try {
    // Remember which anchors had a destination before sanitization. A raw
    // anchor without href must not be mistaken for a rejected dangerous URL.
    const mark = crypto.randomUUID();
    let tree: any = parseFragment(html, { sourceCodeLocationInfo: true });
    // GitHub strips SVG markup while retaining allowed descendants and text.
    // Reparse those descendants as HTML, then run the normal sanitizer.
    let hadSvg = false;
    const unwrapSvg = (node: any) => {
      for (let i = 0; i < (node.childNodes?.length ?? 0); i++) {
        const child = node.childNodes[i];
        if (child.namespaceURI === "http://www.w3.org/2000/svg" && child.tagName !== "a") {
          hadSvg = true;
          const children = child.childNodes ?? [];
          for (const element of children) element.parentNode = node;
          node.childNodes.splice(i, 1, ...children);
          i--;
          continue;
        }
        if (child.content) unwrapSvg(child.content);
        unwrapSvg(child);
      }
    };
    // GitHub closes an unclosed link at a block boundary instead of keeping
    // the HTML parser's reconstructed link around later blocks.
    const seenAnchors = new Set<number>();
    const unwrapReconstructedAnchors = (node: any) => {
      for (let i = 0; i < (node.childNodes?.length ?? 0); i++) {
        const child = node.childNodes[i];
        const start = child.sourceCodeLocation?.startOffset;
        if (child.tagName === "a" && start !== undefined && seenAnchors.has(start)) {
          const children = child.childNodes ?? [];
          node.childNodes.splice(i, 1, ...children);
          i--;
          continue;
        }
        if (child.tagName === "a" && start !== undefined) seenAnchors.add(start);
        unwrapReconstructedAnchors(child);
      }
    };
    unwrapReconstructedAnchors(tree);
    // GitHub drops cmark's final newline when the last raw HTML element
    // reaches EOF without an explicit closing tag. Keep newlines before
    // source-authored closing tags and in already closed Markdown blocks.
    const last = tree.childNodes.at(-1);
    if (last?.tagName && last.sourceCodeLocation && !last.sourceCodeLocation.endTag) {
      let leaf = last;
      while ((leaf.content?.childNodes ?? leaf.childNodes)?.length)
        leaf = (leaf.content?.childNodes ?? leaf.childNodes).at(-1);
      if (leaf.nodeName === "#text" && leaf.value?.endsWith("\n"))
        leaf.value = leaf.value.slice(0, -1);
    }
    // HTML parsing can reconstruct an unclosed formatting tag around the
    // final newline. GitHub drops that empty reconstructed fragment.
    const tail = tree.childNodes.at(-1);
    if (
      tail?.tagName &&
      tail.childNodes?.every((n: any) => n.nodeName === "#text" && !n.value.trim())
    ) {
      const start = tail.sourceCodeLocation?.startOffset;
      const appearsEarlier = (node: any): boolean =>
        node !== tail &&
        ((node.tagName === tail.tagName && node.sourceCodeLocation?.startOffset === start) ||
          (node.childNodes ?? []).some(appearsEarlier));
      if (start !== undefined && tree.childNodes.some(appearsEarlier)) tree.childNodes.pop();
    }
    unwrapSvg(tree);
    if (hadSvg) tree = parseFragment(serialize(tree), { sourceCodeLocationInfo: true });
    const markLinks = (node: any) => {
      if (node.tagName === "a" && node.attrs?.some((a: any) => a.name === "href")) {
        node.attrs = node.attrs.filter((a: any) => a.name !== "data-gfm-had-href");
        node.attrs.push({ name: "data-gfm-had-href", value: mark });
      }
      for (const child of node.childNodes ?? []) markLinks(child);
      if (node.content) markLinks(node.content);
    };
    markLinks(tree);
    document.body.innerHTML = createDOMPurify(window as any).sanitize(serialize(tree), {
      FORBID_ATTR: ["style"],
      FORBID_TAGS: ["svg"],
      ADD_ATTR: [
        "data-footnote-ref",
        "data-footnote-backref",
        "data-footnote-backref-idx",
        "data-footnotes",
      ],
      ADD_TAGS: ["input"],
    });
    // Native preparation preserves text-node boundaries left by templates and comments.
    const flattened = flattenTemplates(document.body.innerHTML, mark);
    const withDetails = prepareDetails(flattened);
    let prepared = flattened;
    if (withDetails !== flattened) {
      // The previous DOM mutation reparsed here; preserve that behavior.
      document.body.innerHTML = removeBoundaries(withDetails, mark);
      prepared = document.body.innerHTML;
    }
    const token = options.runId ?? crypto.randomUUID().replaceAll("-", "");
    const tableMarker = `data-gfm-table-${mark.replaceAll("-", "")}`;
    const markedHtml = prepareMathTables(prepared, token, mark, tableMarker);
    const htmlBeforeMath = removeBoundaries(markedHtml, mark);
    const mathHtml = renderMathHtml(markedHtml, token, mark);
    let linkedHtml: string;
    if (mathHtml !== htmlBeforeMath) {
      document.body.innerHTML = mathHtml;
      linkedHtml = processLinks(document.body.innerHTML, mark);
    } else linkedHtml = processLinks(htmlBeforeMath, mark);
    const output = finalizeHtml(linkedHtml, token, tableMarker).replace(/\n$/, "");
    if (Buffer.byteLength(output) > 32 * 1024 * 1024)
      throw new RangeError("Rendered HTML exceeds 32 MB");
    return output;
  } finally {
    await window.happyDOM.abort();
  }
}
