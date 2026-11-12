import { parseFragment } from "../html";
import { attachNodeLocations, locateDifferences } from "./locations";
import { parseWithSourceMap } from "./source-map";
import type { SourceRange } from "./locations";

export { locateDifferences } from "./locations";
export type { LocatedFinding, SourceRange } from "./locations";

type Node = {
  nodeName: string;
  tagName?: string;
  value?: string;
  attrs?: { name: string; value: string }[];
  childNodes?: Node[];
};
export type SemanticKind =
  "text" | "math" | "code" | "links" | "images" | "tasks" | "structure" | "formatting";
export type SemanticNode =
  | { kind: "text"; text: string; source?: NodeSource }
  | { kind: "math"; display: boolean; expression: string; source?: NodeSource }
  | { kind: "math-error"; text: string; source?: NodeSource }
  | { kind: "code"; block: boolean; text: string; source?: NodeSource }
  | { kind: "image"; src: string; alt: string; source?: NodeSource };
export type NodeSource = {
  status: "block" | "ambiguous" | "unavailable";
  locations: SourceRange[];
};
export type SemanticSummary = Record<SemanticKind, unknown[]> & {
  nodes: SemanticNode[];
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
    nodes: [],
    mathErrors: [],
    footnoteReferences: 0,
    alerts: 0,
  };
  let text = "";
  const flushText = () => {
    const normalized = text.replace(/[\t\n\f\r ]+/g, " ").trim();
    if (normalized) {
      result.text.push(normalized);
      result.nodes.push({ kind: "text", text: normalized });
    }
    text = "";
  };
  const addNode = (node: SemanticNode) => {
    flushText();
    result.nodes.push(node);
  };
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
      text += n.value ?? "";
      return;
    }
    const tag = n.tagName;
    const classes = (attr(n, "class") ?? "").split(/\s+/);
    if (classes.includes("markdown-alert")) result.alerts++;
    if (tag === "a" && attr(n, "data-footnote-ref") !== undefined) result.footnoteReferences++;
    if (classes.includes("katex-error")) {
      result.mathErrors.push(attr(n, "title") ?? contents(n));
      addNode({ kind: "math-error", text: attr(n, "title") ?? contents(n) });
      return;
    }
    if (classes.includes("katex-display") || classes.includes("katex")) {
      const annotation = findAnnotation(n);
      if (annotation) {
        const display = classes.includes("katex-display");
        const expression = display ? contents(annotation).replace(/\n$/, "") : contents(annotation);
        result.math.push({ display, expression });
        addNode({ kind: "math", display, expression });
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
      addNode({ kind: "math", display, expression });
      return;
    }
    if (tag === "pre" || tag === "code") {
      const content = contents(n);
      result.code.push({
        block: tag === "pre",
        text: tag === "pre" ? content.replace(/\n$/, "") : content,
      });
      addNode({
        kind: "code",
        block: tag === "pre",
        text: tag === "pre" ? content.replace(/\n$/, "") : content,
      });
      return;
    }
    if (tag === "img") {
      result.images.push({
        src: attr(n, "data-canonical-src") ?? attr(n, "src") ?? "",
        alt: attr(n, "alt") ?? "",
      });
      addNode({
        kind: "image",
        src: attr(n, "data-canonical-src") ?? attr(n, "src") ?? "",
        alt: attr(n, "alt") ?? "",
      });
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
      flushText();
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
      flushText();
  };
  visit(parseFragment(html) as Node);
  flushText();
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
export function missingTextFragments(expected: unknown[], actual: unknown[]): string[] {
  const available = actual.map(String).filter(Boolean);
  let index = 0,
    offset = 0;
  const missing: string[] = [];
  for (const item of expected) {
    const value = String(item);
    if (!value) continue;
    let match: { index: number; offset: number } | undefined;
    for (let i = index; i < available.length && !match; i++) {
      for (let start = i === index ? offset : 0; start < available[i].length && !match; start++) {
        if (available[i][start] !== value[0]) continue;
        let node = i,
          position = start,
          char = 0;
        while (char < value.length && node < available.length) {
          const current = position === available[node].length ? " " : available[node][position];
          if (current !== value[char]) break;
          char++;
          if (position === available[node].length) {
            node++;
            position = 0;
          } else position++;
        }
        if (char === value.length) match = { index: node, offset: position };
      }
    }
    if (match) {
      index = match.index;
      offset = match.offset;
    } else missing.push(value);
  }
  return missing;
}

export function detectPotentialBreakage(githubHtml: string, vscodeHtml: string) {
  return detectSummaries(summarizeHtml(githubHtml), summarizeHtml(vscodeHtml));
}

function detectSummaries(github: SemanticSummary, vscode: SemanticSummary) {
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
  const githubHasExtraFeatures =
    github.math.length > vscode.math.length || github.tasks.length > vscode.tasks.length;
  if (
    !vscode.mathErrors.length &&
    !githubHasExtraFeatures &&
    missingTextFragments(vscode.text, github.text).length
  )
    categories.push("text");
  const vscodeUnsupported: string[] = [];
  if (github.footnoteReferences && !vscode.footnoteReferences) vscodeUnsupported.push("footnotes");
  if (github.alerts && !vscode.alerts) vscodeUnsupported.push("alerts");
  if (github.tasks.length && !vscode.tasks.length) vscodeUnsupported.push("task checkboxes");
  return { categories, vscodeMathErrors: vscode.mathErrors, vscodeUnsupported, github, vscode };
}

/** Compare the rendered fragments and locate the resulting candidates in Markdown. */
export function compareMarkdown(source: string, githubHtml: string, vscodeHtml: string) {
  const parsed = parseWithSourceMap(source);
  const vscode = summarizeHtml(vscodeHtml);
  vscode.nodes = attachNodeLocations(source, vscode.nodes, parsed);
  const analysis = detectSummaries(summarizeHtml(githubHtml), vscode);
  return { ...analysis, findings: locateDifferences(source, analysis, parsed) };
}
