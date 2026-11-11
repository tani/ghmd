import { createDOM } from "./dom";
import { finalizeHtml, nativeRender, renderMathHtml } from "./native";
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
    // GitHub removes template elements while retaining their sanitized contents.
    const unwrapTemplates = (root: any) => {
      for (const child of [...root.childNodes]) {
        if (child.nodeType !== 1) continue;
        if (child.tagName === "TEMPLATE") {
          unwrapTemplates(child.content);
          child.replaceWith(...child.content.childNodes);
        } else unwrapTemplates(child);
      }
    };
    unwrapTemplates(document.body);
    const comments = document.createTreeWalker(document.body, window.NodeFilter.SHOW_COMMENT);
    const removeComments = [];
    while (comments.nextNode()) removeComments.push(comments.currentNode);
    for (const comment of removeComments) comment.remove();
    const token = options.runId ?? crypto.randomUUID().replaceAll("-", "");
    // Details must be populated before math reads descendant textContent.
    for (const details of document.querySelectorAll("details")) {
      if (!details.querySelector(":scope > summary")) {
        const summary = document.createElement("summary");
        summary.textContent = "Details";
        details.prepend(summary);
      }
    }
    // A table moved out of a template can temporarily be inside a paragraph.
    // GitHub processes math within that paragraph before applying its table
    // wrapper; unrelated paragraphs retain their own table wrapper.
    const tableMarker = `data-gfm-table-${mark.replaceAll("-", "")}`;
    for (const paragraph of document.querySelectorAll("p")) {
      const tables = [...paragraph.querySelectorAll("table")];
      if (!tables.length) continue;
      let blocked = false;
      for (let parent = paragraph.parentElement; parent; parent = parent.parentElement) {
        if (parent.tagName === "DETAILS") break;
        if (["B", "STRONG", "EM", "A", "CODE", "PRE", "SECTION"].includes(parent.tagName)) {
          blocked = true;
          break;
        }
      }
      const before = paragraph.outerHTML;
      const hasMath = !blocked && renderMathHtml(before, token, mark) !== before;
      for (const table of tables) {
        table.setAttribute(tableMarker, "");
        if (!hasMath) {
          const wrapper = document.createElement("markdown-accessiblity-table");
          table.replaceWith(wrapper);
          wrapper.append(table);
        }
      }
    }
    const htmlBeforeMath = document.body.innerHTML;
    // Preserve adjacent text-node boundaries left by removed comments.
    const textNodes = document.createTreeWalker(document.body, window.NodeFilter.SHOW_TEXT);
    const markers: any[] = [];
    while (textNodes.nextNode()) {
      const node = textNodes.currentNode;
      if (node.previousSibling?.nodeType === 3) {
        const marker = document.createComment(mark);
        node.parentNode!.insertBefore(marker, node);
        markers.push(marker);
      }
    }
    const markedHtml = markers.length ? document.body.innerHTML : htmlBeforeMath;
    for (const marker of markers) marker.remove();
    const mathHtml = renderMathHtml(markedHtml, token, mark);
    if (mathHtml !== htmlBeforeMath) document.body.innerHTML = mathHtml;
    // GitHub removes links whose destination was removed by sanitization.
    for (const link of [...document.querySelectorAll("a[data-gfm-had-href]")]) {
      const hadHref = link.getAttribute("data-gfm-had-href") === mark;
      link.removeAttribute("data-gfm-had-href");
      if (
        hadHref &&
        !link.hasAttribute("href") &&
        !link.hasAttribute("id") &&
        !link.hasAttribute("name")
      )
        link.replaceWith(...link.childNodes);
    }
    for (const link of document.querySelectorAll("a[href]")) {
      const href = link.getAttribute("href")!;
      if (/^(?:https?:)?\/\//i.test(href) || /^mailto:/i.test(href))
        link.setAttribute("rel", "nofollow");
    }
    const output = finalizeHtml(document.body.innerHTML, token, tableMarker).replace(/\n$/, "");
    if (Buffer.byteLength(output) > 32 * 1024 * 1024)
      throw new RangeError("Rendered HTML exceeds 32 MB");
    return output;
  } finally {
    await window.happyDOM.abort();
  }
}
