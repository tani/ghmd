import { createDOM } from "./dom";
import { nativeRender } from "./native";
import { parseFragment, serialize } from "parse5";
import createDOMPurify from "dompurify";
import { alertTitles } from "./alerts";
import { highlight } from "./highlight";
import { flattenDeepLists } from "./list-depth";
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
        if (/^\$\$\n[\s\S]*\n\$\$$/.test(reconstructed)) {
          const inner = reconstructed.slice(3, -3);
          paragraph.replaceChildren(mathElement("$$\n\n" + inner + "\n\n$$", true));
          continue;
        }
      }
      // Display formulas contain only text and hard-break elements.
      if (children.every((n) => n.nodeType === 3 || n.nodeName === "BR")) {
        const text = paragraph.textContent ?? "";
        if (paragraph.tagName !== "LI" && !paragraph.closest("b,details")) {
          const blocks = [...text.matchAll(/\$\$\n[\s\S]*?\n\$\$/g)];
          if (blocks.length && blocks.map((m) => m[0]).join("\n") === text) {
            const nodes: any[] = [];
            for (const block of blocks) {
              if (nodes.length) nodes.push(document.createTextNode("\n"));
              nodes.push(mathElement(block[0].replace(/\\\n/g, "\\\\\n"), true));
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
        if (
          before?.nodeType !== 3 ||
          after?.nodeType !== 3 ||
          !before.textContent?.endsWith("$") ||
          !after.textContent?.startsWith("$")
        )
          continue;
        // GitHub GFM's current protected-math boundary differs for lowercase
        // a/b; the differential matrix records the tested letters and words.
        if (/[ab]$/.test(before.textContent.slice(0, -1))) continue;
        if (paragraph.tagName === "EM" && after.textContent === "$") continue;
        const value = code.textContent ?? "";
        before.textContent = before.textContent.slice(0, -1);
        after.textContent = after.textContent.slice(1);
        // GitHub's math filter escapes an ampersand already escaped by HTML.
        code.replaceWith(
          mathElement(
            "$" +
              value.replaceAll("&", "&amp;").replaceAll("<", "&lt;").replaceAll(">", "&gt;") +
              "$",
            false,
          ),
        );
      }
      if (paragraph.tagName === "EM" || (paragraph.tagName === "STRONG" && paragraph.closest("em")))
        continue;
      for (const node of [...paragraph.childNodes]) {
        if (node.nodeType !== 3 || !node.textContent?.includes("$")) continue;
        const text = node.textContent,
          re = /\$(?!\$)([^$\n]+?)\$/g;
        let match: RegExpExecArray | null,
          last = 0;
        const fragment = document.createDocumentFragment();
        while ((match = re.exec(text))) {
          const before = text[match.index - 1],
            after = text[match.index + match[0].length];
          if (
            (before && !/[\s(]/.test(before)) ||
            (after && /[A-Za-z0-9_]/.test(after)) ||
            /^\s/.test(match[1]) ||
            (/\s$/.test(match[1]) && after !== "$") ||
            match[1].includes("&")
          ) {
            // A rejected opening dollar can contain the next valid candidate.
            re.lastIndex = match.index + 1;
            continue;
          }
          fragment.append(
            document.createTextNode(text.slice(last, match.index)),
            mathElement(match[0].replaceAll("<", "&lt;").replaceAll(">", "&gt;"), false),
          );
          last = match.index + match[0].length;
        }
        if (last) {
          fragment.append(document.createTextNode(text.slice(last)));
          node.replaceWith(fragment);
        }
      }
    }
    for (const pre of [...document.querySelectorAll('pre[lang="math"]')]) {
      if (blocked(pre)) continue;
      const content = (pre.textContent ?? "")
        .replace(/\n$/, "")
        .replace(/(?<!\\)\$/g, "\\$")
        .replace(/\\\n/g, "\\\\\n");
      pre.replaceWith(mathElement("$$" + content + "$$", true));
    }
    for (const pre of [...document.querySelectorAll("pre[lang]")]) await highlight(document, pre);
    for (const code of document.querySelectorAll("code")) code.classList.add("notranslate");
    for (const pre of document.querySelectorAll("pre")) {
      if (pre.getAttribute("lang") === "math") pre.removeAttribute("lang");
      pre.classList.add("notranslate");
    }
    for (const table of [...document.querySelectorAll("table")]) {
      table.setAttribute("role", "table");
      const wrapper = document.createElement("markdown-accessiblity-table");
      table.replaceWith(wrapper);
      wrapper.append(table);
    }
    for (const quote of [...document.querySelectorAll("blockquote")]) {
      if (quote.parentElement?.closest("blockquote")) continue;
      const paragraph = quote.firstElementChild;
      if (
        !paragraph ||
        paragraph.tagName !== "P" ||
        !paragraph.firstChild ||
        paragraph.firstChild.nodeType !== 3
      )
        continue;
      const marker = paragraph.firstChild.textContent?.match(
        /^\[!(NOTE|TIP|IMPORTANT|WARNING|CAUTION)\]$/,
      );
      if (!marker || paragraph.firstChild.nextSibling?.nodeName !== "BR") continue;
      const kind = marker[1].toLowerCase();
      paragraph.firstChild.nextSibling.remove();
      paragraph.firstChild!.remove();
      if (paragraph.firstChild?.nodeType === 3)
        paragraph.firstChild.textContent = paragraph.firstChild.textContent!.replace(/^\n/, "");
      const alert = document.createElement("div");
      alert.className = `markdown-alert markdown-alert-${kind}`;
      const title = document.createElement("p");
      title.className = "markdown-alert-title";
      title.innerHTML = alertTitles[kind];
      alert.append(title);
      while (quote.firstChild) alert.append(quote.firstChild);
      if (alert.childNodes[1]?.nodeType === 3 && !alert.childNodes[1].textContent?.trim())
        alert.childNodes[1].remove();
      quote.replaceWith(alert);
    }
    for (const input of document.querySelectorAll('input[type="checkbox"]')) {
      if (!input.closest("li")) {
        input.remove();
        continue;
      }
      input.closest("li")?.classList.add("task-list-item");
      input.closest("ul,ol")?.classList.add("contains-task-list");
      const checked = input.hasAttribute("checked");
      input.removeAttribute("checked");
      input.removeAttribute("disabled");
      input.setAttribute("id", "");
      input.setAttribute("disabled", "");
      input.classList.add("task-list-item-checkbox");
      input.setAttribute("aria-label", checked ? "Completed task" : "Incomplete task");
      if (checked) input.setAttribute("checked", "");
    }
    // For loose task items, GitHub places the checkbox inside the first
    // paragraph rather than before it.
    for (const li of document.querySelectorAll("li.task-list-item")) {
      const input = li.firstElementChild;
      if (input?.matches('input[type="checkbox"]')) {
        const paragraph = [...li.children].find((e) => e.tagName === "P");
        if (paragraph && input.nextElementSibling === paragraph) {
          const space = input.nextSibling;
          if (space?.nodeType === 3 && !space.textContent?.trim()) space.remove();
          input.remove();
          paragraph.prepend(document.createTextNode(" "));
          paragraph.prepend(input);
          if (
            !li.firstChild ||
            li.firstChild.nodeType !== 3 ||
            !li.firstChild.textContent?.startsWith("\n")
          )
            li.prepend(document.createTextNode("\n"));
        }
      }
    }
    // GitHub's serializer separates a final math-related inline element from
    // the closing list tag when other inline content precedes it.
    for (const li of document.querySelectorAll("li")) {
      const meaningful = [...li.childNodes].filter(
        (n) => n.nodeType !== 3 || !!n.textContent?.trim(),
      );
      const last = meaningful.at(-1);
      const hasMath =
        last?.nodeType === 1 &&
        (last.nodeName === "MATH-RENDERER" || (last as any).querySelector?.("math-renderer"));
      if (meaningful.length > 1 && hasMath && !li.textContent?.endsWith("\n"))
        li.append(document.createTextNode("\n"));
      // GitHub separates a leading formula when ordinary text follows it.
      if (
        meaningful[0]?.nodeName === "MATH-RENDERER" &&
        meaningful.length > 1 &&
        meaningful[1]?.nodeType === 3 &&
        meaningful[1].textContent?.trim() &&
        [...li.childNodes]
          .slice(0, [...li.childNodes].indexOf(meaningful[0]))
          .every((n) => n.nodeType === 3 && !n.textContent)
      )
        li.insertBefore(document.createTextNode("\n"), meaningful[0]);
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
    // GitHub wraps standalone relative images and adds the responsive style.
    for (const image of [...document.querySelectorAll("img[src]")]) {
      const src = image.getAttribute("src")!;
      if (image.parentElement === document.body || image.parentElement?.tagName === "B") {
        const paragraph = document.createElement("p");
        image.replaceWith(paragraph);
        paragraph.append(image);
      }
      if (/^(?:https?:)?\/\//i.test(src)) continue;
      image.setAttribute("style", "max-width: 100%;");
      const rejected = /^data:/i.test(src);
      if (rejected) image.removeAttribute("src");
      if (!image.closest("a")) {
        const link = document.createElement("a");
        link.setAttribute("target", "_blank");
        link.setAttribute("rel", "noopener noreferrer");
        link.setAttribute("href", rejected ? "" : src);
        image.replaceWith(link);
        link.append(image);
      }
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
    for (const sup of document.querySelectorAll("sup.footnote-ref")) sup.removeAttribute("class");
    for (const a of document.querySelectorAll("a[data-footnote-ref]")) {
      const href = a.getAttribute("href");
      const id = a.getAttribute("id");
      if (href) {
        const name = href.slice(4);
        if (name === ":") a.removeAttribute("href");
        else a.setAttribute("href", `#user-content-fn-${name}-${token}`);
      }
      if (id) a.setAttribute("id", `user-content-${id.toLowerCase()}-${token}`);
      a.setAttribute("aria-describedby", "footnote-label");
    }
    for (const a of document.querySelectorAll("a[data-footnote-backref]")) {
      const href = a.getAttribute("href");
      if (href?.slice(7) === ":") {
        a.replaceWith(document.createTextNode("↩"));
        continue;
      }
      if (href) a.setAttribute("href", `#user-content-${href.slice(1)}-${token}`);
      a.removeAttribute("class");
      a.removeAttribute("data-footnote-backref-idx");
      a.setAttribute("class", "data-footnote-backref");
    }
    const section = document.querySelector("section[data-footnotes]");
    if (section) {
      section.removeAttribute("class");
      section.removeAttribute("data-footnotes");
      section.setAttribute("data-footnotes", "");
      section.setAttribute("class", "footnotes");
      for (const li of section.querySelectorAll("li[id]")) li.id = `user-content-${li.id}-${token}`;
      const title = document.createElement("h2");
      title.id = "footnote-label";
      title.className = "sr-only";
      title.textContent = "Footnotes";
      section.insertBefore(title, section.firstChild);
      const next = section.nextSibling;
      if (
        section.parentElement !== document.body &&
        next?.nodeType === 3 &&
        !next.textContent?.trim() &&
        !next.nextSibling
      )
        next.remove();
    }
    flattenDeepLists(document as unknown as Document);
    const output = document.body.innerHTML.replace(/\n$/, "");
    if (Buffer.byteLength(output) > 32 * 1024 * 1024)
      throw new RangeError("Rendered HTML exceeds 32 MB");
    return output;
  } finally {
    await window.happyDOM.abort();
  }
}
