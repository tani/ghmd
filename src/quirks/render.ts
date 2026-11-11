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
    const output = finalizeHtml(document.body.innerHTML, token).replace(/\n$/, "");
    if (Buffer.byteLength(output) > 32 * 1024 * 1024)
      throw new RangeError("Rendered HTML exceeds 32 MB");
    return output;
  } finally {
    await window.happyDOM.abort();
  }
}
