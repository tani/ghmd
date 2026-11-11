import { createDOM } from "./dom";
import {
  displayMathBlocks,
  emphasisDisplayMath,
  fencedMath,
  finalizeHtml,
  nativeRender,
  protectedMath,
  splitInlineMath,
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
    // The API leaves TeX for its client-side math renderer. It does not render
    // math inside footnotes, links, code, or inline formatting elements.
    const blocked = (element: any) => {
      for (let p = element.parentElement; p; p = p.parentElement) {
        if (p.tagName === "DETAILS") return false;
        if (["B", "STRONG", "EM", "A", "CODE", "PRE", "SECTION"].includes(p.tagName)) return true;
      }
      return false;
    };
    const mathElement = (content: string, display: boolean) => {
      const math = document.createElement("math-renderer");
      math.className = display ? "js-display-math" : "js-inline-math";
      math.setAttribute("style", display ? "display: block" : "display: inline-block");
      math.setAttribute("data-run-id", token);
      math.textContent = content;
      return math;
    };
    // Details must be populated before math reads descendant textContent.
    for (const details of document.querySelectorAll("details")) {
      if (!details.querySelector(":scope > summary")) {
        const summary = document.createElement("summary");
        summary.textContent = "Details";
        details.prepend(summary);
      }
    }
    for (const paragraph of [
      ...document.querySelectorAll(
        "p, li, h1, h2, h3, h4, h5, h6, td, th, strong, em, i, del, span, summary",
      ),
    ]) {
      if (
        blocked(paragraph) &&
        !(paragraph.tagName === "STRONG" && !paragraph.closest("b,section,a,pre,code"))
      )
        continue;
      const children = [...paragraph.childNodes];
      // GitHub reconstructs underscore emphasis in display math after Markdown
      // parsing, adding a blank line around the reconstructed expression.
      if (
        !paragraph.closest("b") &&
        children.some((n) => n.nodeName === "EM") &&
        children.every((n) => n.nodeType === 3 || n.nodeName === "BR" || n.nodeName === "EM")
      ) {
        const reconstructed = children
          .map((n) =>
            n.nodeName === "EM" ? "_" + n.textContent + "_" : n.nodeType === 3 ? n.textContent : "",
          )
          .join("");
        const display = emphasisDisplayMath(reconstructed);
        if (display !== undefined) {
          paragraph.replaceChildren(mathElement(display, true));
          continue;
        }
      }
      // Display formulas contain only text and hard-break elements.
      if (children.every((n) => n.nodeType === 3 || n.nodeName === "BR")) {
        const text = paragraph.textContent ?? "";
        if (paragraph.tagName !== "LI" && !paragraph.closest("b,details")) {
          const blocks = displayMathBlocks(text);
          if (blocks) {
            const nodes: any[] = [];
            for (const block of blocks) {
              if (nodes.length) nodes.push(document.createTextNode("\n"));
              nodes.push(mathElement(block, true));
            }
            paragraph.replaceChildren(...nodes);
            continue;
          }
        }
      }
      // Dollar-backtick spans are parsed as a code element surrounded by dollars.
      for (const code of [...paragraph.querySelectorAll(":scope > code")]) {
        const before = code.previousSibling,
          after = code.nextSibling;
        if (before?.nodeType !== 3 || after?.nodeType !== 3) continue;
        const value = protectedMath(
          before.textContent ?? "",
          after.textContent ?? "",
          code.textContent ?? "",
          paragraph.tagName === "EM",
        );
        if (value === undefined) continue;
        before.textContent = before.textContent!.slice(0, -1);
        after.textContent = after.textContent!.slice(1);
        code.replaceWith(mathElement(value, false));
      }
      if (paragraph.tagName === "EM" || (paragraph.tagName === "STRONG" && paragraph.closest("em")))
        continue;
      for (const node of [...paragraph.childNodes]) {
        if (node.nodeType !== 3 || !node.textContent?.includes("$")) continue;
        const segments = splitInlineMath(node.textContent);
        if (segments) {
          const fragment = document.createDocumentFragment();
          for (const segment of segments) {
            fragment.append(
              segment.math
                ? mathElement(segment.text.replaceAll("<", "&lt;").replaceAll(">", "&gt;"), false)
                : document.createTextNode(segment.text),
            );
          }
          node.replaceWith(fragment);
        }
      }
    }
    for (const pre of [...document.querySelectorAll('pre[lang="math"]')]) {
      if (blocked(pre)) continue;
      pre.replaceWith(mathElement(fencedMath(pre.textContent ?? ""), true));
    }
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
    if (options.imageProxy) {
      for (const image of [...document.querySelectorAll("img[src]")]) {
        const canonical = image.getAttribute("src")!;
        if (!/^https?:\/\//i.test(canonical)) continue;
        // Proxy anchors are created here, so adjust their future parent first.
        if (image.parentElement === document.body || image.parentElement?.tagName === "B") {
          const paragraph = document.createElement("p");
          image.replaceWith(paragraph);
          paragraph.append(image);
        }
        const proxy = await options.imageProxy(canonical);
        if (!/^https?:\/\//i.test(proxy))
          throw new TypeError("imageProxy must return an HTTP(S) URL");
        image.setAttribute("src", proxy);
        image.setAttribute("data-canonical-src", canonical);
        image.setAttribute("style", "max-width: 100%;");
        if (!image.closest("a")) {
          const link = document.createElement("a");
          link.setAttribute("target", "_blank");
          link.setAttribute("rel", "noopener noreferrer nofollow");
          link.setAttribute("href", proxy);
          image.replaceWith(link);
          link.append(image);
        }
      }
    }
    // JavaScript retains responsibility for full Unicode lowercase mappings.
    for (const anchor of document.querySelectorAll("a[data-footnote-ref][id]")) {
      const id = anchor.getAttribute("id");
      if (id) anchor.setAttribute("id", id.toLowerCase());
    }
    const output = finalizeHtml(document.body.innerHTML, token).replace(/\n$/, "");
    if (Buffer.byteLength(output) > 32 * 1024 * 1024)
      throw new RangeError("Rendered HTML exceeds 32 MB");
    return output;
  } finally {
    await window.happyDOM.abort();
  }
}
