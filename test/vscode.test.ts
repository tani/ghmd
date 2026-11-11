import { describe, expect, test } from "bun:test";
import { render } from "../src/vscode";
import { summarizeHtml } from "../src/compare";

describe("VS Code-style Markdown preview", () => {
  test("supports Markdown tables, strikethrough, URL links and code", async () => {
    const html = await render(
      "# Title\n\n| A | B |\n| --- | --- |\n| ~~old~~ | https://example.com |\n\n```js\nconst x = 1;\n```",
    );
    expect(html).toContain("<h1>Title</h1>");
    expect(html).toContain("<table>");
    expect(html).toContain("<s>old</s>");
    expect(html).toContain('<a href="https://example.com">https://example.com</a>');
    expect(html).toContain("hljs-");
    expect(summarizeHtml(html).code).toEqual([{ block: true, text: "const x = 1;" }]);
  });
  test("uses the VS Code KaTeX plugin for inline, protected and display math", async () => {
    const html = await render("$x$ and $`y`$\n\n$$\nz^2\n$$");
    expect(summarizeHtml(html).math).toEqual([
      { display: false, expression: "x" },
      { display: false, expression: "y" },
      { display: true, expression: "z^2" },
    ]);
    expect(html).toContain('class="katex"');
    expect(summarizeHtml(await render("$x$", { math: false })).math).toEqual([]);
  });
  test("handles html-like characters in math and keeps literal code", async () => {
    const html = await render("$a <b > c$ and `$x$`");
    const summary = summarizeHtml(html);
    expect(summary.math).toEqual([{ display: false, expression: "a <b > c" }]);
    expect(summary.code).toEqual([{ block: false, text: "$x$" }]);
  });
  test("reports math errors and does not treat them as rendered expressions", async () => {
    const summary = summarizeHtml(await render("$\\notacommand$"));
    expect(summary.math).toEqual([]);
    expect(summary.mathErrors).toEqual(["\\notacommand"]);
  });
  test("preserves Mermaid source and follows standard markdown-it task behavior", async () => {
    const html = await render("```mermaid\nflowchart TD\n A --> B\n```\n\n- [x] task");
    const summary = summarizeHtml(html);
    expect(summary.code).toEqual([{ block: true, text: "flowchart TD\n A --> B" }]);
    expect(summary.text).toContain("[x] task");
    expect(summary.nodes.some((node) => node.kind === "code")).toBe(true);
    expect(summary.tasks).toEqual([]);
  });
  test("validates input and offers VS Code preview options", async () => {
    expect(await render("")).toBe("");
    expect(await render("a\nb", { breaks: true })).toContain("<br>");
    expect(await render("https://example.com", { linkify: false })).not.toContain("<a ");
    await expect(render(null as any)).rejects.toBeInstanceOf(TypeError);
  });
});
