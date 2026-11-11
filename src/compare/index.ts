import { parseFragment } from "parse5";

type Node = {
  nodeName: string;
  tagName?: string;
  value?: string;
  attrs?: { name: string; value: string }[];
  childNodes?: Node[];
};
export type SemanticKind =
  "text" | "math" | "code" | "links" | "images" | "tasks" | "structure" | "formatting";
export type SemanticSummary = Record<SemanticKind, unknown[]> & {
  mathErrors: string[];
  footnoteReferences: number;
  alerts: number;
};

/** A deliberately lossy projection for triage, not an HTML equivalence test. */
export function summarizeHtml(html: string): SemanticSummary {
  const result: SemanticSummary = {
    text: [],
    math: [],
    code: [],
    links: [],
    images: [],
    tasks: [],
    structure: [],
    formatting: [],
    mathErrors: [],
    footnoteReferences: 0,
    alerts: 0,
  };
  const text: string[] = [];
  const attr = (n: Node, name: string) => n.attrs?.find((a) => a.name === name)?.value;
  const contents = (n: Node): string =>
    n.nodeName === "#text" ? (n.value ?? "") : (n.childNodes ?? []).map(contents).join("");
  const findAnnotation = (n: Node): Node | undefined => {
    if (n.tagName === "annotation" && attr(n, "encoding") === "application/x-tex") return n;
    for (const child of n.childNodes ?? []) {
      const found = findAnnotation(child);
      if (found) return found;
    }
  };
  const visit = (n: Node) => {
    if (n.nodeName === "#text") {
      text.push(n.value ?? "");
      return;
    }
    const tag = n.tagName;
    const classes = (attr(n, "class") ?? "").split(/\s+/);
    if (classes.includes("markdown-alert")) result.alerts++;
    if (tag === "a" && attr(n, "data-footnote-ref") !== undefined) result.footnoteReferences++;
    if (classes.includes("katex-error")) {
      result.mathErrors.push(attr(n, "title") ?? contents(n));
      text.push("[math-error]");
      return;
    }
    if (classes.includes("katex-display") || classes.includes("katex")) {
      const annotation = findAnnotation(n);
      if (annotation) {
        const display = classes.includes("katex-display");
        const expression = display ? contents(annotation).replace(/\n$/, "") : contents(annotation);
        result.math.push({ display, expression });
        text.push("[math]");
        return;
      }
    }
    if (
      tag === "svg" ||
      attr(n, "aria-hidden") === "true" ||
      /(?:^|\s)sr-only(?:\s|$)/.test(attr(n, "class") ?? "")
    )
      return;
    if (tag === "math-renderer") {
      let expression = contents(n);
      const display = /(?:^|\s)js-display-math(?:\s|$)/.test(attr(n, "class") ?? "");
      if (
        tag === "math-renderer" &&
        display &&
        expression.startsWith("$$") &&
        expression.endsWith("$$")
      )
        expression = expression.slice(2, -2).replace(/^\n|\n$/g, "");
      else if (
        tag === "math-renderer" &&
        !display &&
        expression.startsWith("$") &&
        expression.endsWith("$")
      )
        expression = expression.slice(1, -1);
      result.math.push({ display, expression });
      text.push("[math]");
      return;
    }
    if (tag === "pre" || tag === "code") {
      const content = contents(n);
      result.code.push({
        block: tag === "pre",
        text: tag === "pre" ? content.replace(/\n$/, "") : content,
      });
      text.push("[code]");
      return;
    }
    if (tag === "img") {
      result.images.push({
        src: attr(n, "data-canonical-src") ?? attr(n, "src") ?? "",
        alt: attr(n, "alt") ?? "",
      });
      text.push("[image]");
      return;
    }
    if (tag === "input" && attr(n, "type") === "checkbox") {
      result.tasks.push(n.attrs?.some((a) => a.name === "checked") ?? false);
      return;
    }
    if (tag === "a" && attr(n, "data-footnote-backref") !== undefined) return;
    if (tag === "br") result.structure.push("<br>");
    const format = tag === "b" ? "strong" : tag === "i" ? "em" : tag === "s" ? "del" : tag;
    if (
      format &&
      ["strong", "em", "del", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote"].includes(format)
    )
      result.formatting.push(`<${format}>`);
    if (
      tag === "a" &&
      attr(n, "href") &&
      attr(n, "data-footnote-ref") === undefined &&
      !(n.childNodes ?? []).some((c) => c.tagName === "img")
    )
      result.links.push(attr(n, "href"));
    if (tag && ["ul", "ol", "li", "table", "thead", "tbody", "tr", "th", "td"].includes(tag))
      result.structure.push(`<${tag}>`);
    if (
      tag &&
      ["p", "div", "li", "tr", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote", "br"].includes(tag)
    )
      text.push(" ");
    for (const child of n.childNodes ?? []) visit(child);
    if (
      format &&
      ["strong", "em", "del", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote"].includes(format)
    )
      result.formatting.push(`</${format}>`);
    if (tag && ["ul", "ol", "li", "table", "thead", "tbody", "tr", "th", "td"].includes(tag))
      result.structure.push(`</${tag}>`);
    if (
      tag &&
      ["p", "div", "li", "tr", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote"].includes(tag)
    )
      text.push(" ");
  };
  visit(parseFragment(html) as Node);
  result.text = [
    text
      .join("")
      .replace(/[\t\n\f\r ]+/g, " ")
      .trim(),
  ];
  return result;
}

/** Categories that need investigation; differences alone do not establish which renderer is correct. */
export function compareHtml(left: string, right: string): SemanticKind[] {
  const a = summarizeHtml(left),
    b = summarizeHtml(right);
  return (
    [
      "text",
      "math",
      "code",
      "links",
      "images",
      "tasks",
      "structure",
      "formatting",
    ] as SemanticKind[]
  ).filter((key) => JSON.stringify(a[key]) !== JSON.stringify(b[key]));
}

/** Candidate differences in which VS Code preserves a feature that GitHub may not. */
export function detectPotentialBreakage(githubHtml: string, vscodeHtml: string) {
  const github = summarizeHtml(githubHtml),
    vscode = summarizeHtml(vscodeHtml);
  const missing = (expected: unknown[], actual: unknown[]) => {
    const counts = new Map<string, number>();
    for (const item of actual) {
      const key = JSON.stringify(item);
      counts.set(key, (counts.get(key) ?? 0) + 1);
    }
    return expected.filter((item) => {
      const key = JSON.stringify(item),
        count = counts.get(key) ?? 0;
      if (!count) return true;
      counts.set(key, count - 1);
      return false;
    });
  };
  const categories: SemanticKind[] = [];
  for (const kind of ["math", "links", "images", "code"] as const) {
    if (kind === "code" && github.math.length > vscode.math.length) continue;
    if (missing(vscode[kind], github[kind]).length) categories.push(kind);
  }
  if (vscode.tasks.length && JSON.stringify(vscode.tasks) !== JSON.stringify(github.tasks))
    categories.push("tasks");
  const githubStructure = github.structure.map(String),
    vscodeStructure = vscode.structure.map(String);
  let index = 0;
  for (const item of githubStructure) if (item === vscodeStructure[index]) index++;
  if (index < vscodeStructure.length) categories.push("structure");
  const githubFormatting = github.formatting.map(String),
    vscodeFormatting = vscode.formatting.map(String);
  index = 0;
  for (const item of githubFormatting) if (item === vscodeFormatting[index]) index++;
  if (index < vscodeFormatting.length) categories.push("formatting");
  // If VS Code could not render a formula, or GitHub recognizes features that
  // VS Code leaves literal, text differences do not establish a GitHub defect.
  const vscodeText = String(vscode.text[0] ?? "")
    .replace(/\[math\]|\[code\]|\[image\]/g, "")
    .trim();
  const githubText = String(github.text[0] ?? "")
    .replace(/\[math\]|\[code\]|\[image\]/g, "")
    .trim();
  const githubHasExtraFeatures =
    github.math.length > vscode.math.length || github.tasks.length > vscode.tasks.length;
  if (
    !vscode.mathErrors.length &&
    !githubHasExtraFeatures &&
    vscodeText &&
    vscodeText !== githubText &&
    !githubText.includes(vscodeText)
  )
    categories.push("text");
  const vscodeUnsupported: string[] = [];
  if (github.footnoteReferences && !vscode.footnoteReferences) vscodeUnsupported.push("footnotes");
  if (github.alerts && !vscode.alerts) vscodeUnsupported.push("alerts");
  if (github.tasks.length && !vscode.tasks.length) vscodeUnsupported.push("task checkboxes");
  return { categories, vscodeMathErrors: vscode.mathErrors, vscodeUnsupported, github, vscode };
}
