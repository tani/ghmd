import { expect, test } from "bun:test";
import { compareHtml, detectPotentialBreakage } from "../src/compare";
import { render as renderVsCode } from "../src/vscode";
import { render as renderGitHub } from "../src/quirks";

test("ignores serializer formatting and syntax-highlighter wrappers", () => {
  expect(
    compareHtml(
      '<div class="highlight"><pre><span class="pl-k">const</span> x</pre></div>',
      "<pre><code>const x</code></pre>",
    ),
  ).toEqual([]);
  expect(compareHtml("<ul>\n<li>hello</li>\n</ul>", "<ul><li>hello</li></ul>")).toEqual([]);
});
test("reports changed math, code, links, tasks, and list structure", () => {
  expect(compareHtml("<p><math-renderer>$x$</math-renderer></p>", "<p>$x$</p>")).toContain("math");
  expect(compareHtml("<pre>a\nb</pre>", "<pre>a b</pre>")).toContain("code");
  expect(compareHtml('<a href="/a">x</a>', '<a href="/b">x</a>')).toContain("links");
  expect(
    compareHtml('<li><input type="checkbox" checked></li>', '<li><input type="checkbox"></li>'),
  ).toContain("tasks");
  expect(
    compareHtml("<ul><li>x<ul><li>y</li></ul></li></ul>", "<ul><li>x</li><li>y</li></ul>"),
  ).toContain("structure");
});
test("does not classify a generated wrapper around an image as a link", () => {
  expect(
    compareHtml(
      '<a href="https://proxy.example/a"><img src="https://proxy.example/a" data-canonical-src="https://example.com/a" alt="a"></a>',
      '<img src="https://example.com/a" alt="a">',
    ),
  ).toEqual([]);
});

test("normalizes KaTeX and GitHub math without losing TeX", async () => {
  const vscode = await renderVsCode("$x$");
  expect(compareHtml('<math-renderer class="js-inline-math">$x$</math-renderer>', vscode)).toEqual(
    [],
  );
  expect(
    compareHtml('<math-renderer class="js-inline-math">$y$</math-renderer>', vscode),
  ).toContain("math");
  const display = await renderVsCode("$$\nx\n$$");
  expect(
    compareHtml('<math-renderer class="js-display-math">$$\nx\n$$</math-renderer>', display),
  ).toEqual([]);
});
test("reports VS Code math missing in GitHub but not the reverse", async () => {
  const input = "$a <b > c$";
  expect(
    detectPotentialBreakage(await renderGitHub(input), await renderVsCode(input)).categories,
  ).toContain("math");
  const invalid = "$\\notacommand$";
  const result = detectPotentialBreakage(await renderGitHub(invalid), await renderVsCode(invalid));
  expect(result.categories).not.toContain("math");
  expect(result.vscodeMathErrors).toEqual(["\\notacommand"]);
  const task = "- [x] task";
  const taskResult = detectPotentialBreakage(await renderGitHub(task), await renderVsCode(task));
  expect(taskResult.categories).toEqual([]);
  expect(taskResult.vscodeUnsupported).toContain("task checkboxes");
});
test("retains visible hard line breaks but ignores a code block serialization newline", () => {
  expect(compareHtml("<p>a<br>b</p>", "<p>a\nb</p>")).toContain("structure");
  expect(compareHtml("<pre>const x</pre>", "<pre><code>const x\n</code></pre>")).toEqual([]);
  expect(compareHtml("<pre>a\n\n</pre>", "<pre>a\n</pre>")).toContain("code");
});
test("normalizes equivalent markup and detects missing visual emphasis", () => {
  expect(detectPotentialBreakage("<p><del>x</del></p>", "<p><s>x</s></p>").categories).toEqual([]);
  expect(detectPotentialBreakage("<p>x</p>", "<p><strong>x</strong></p>").categories).toContain(
    "formatting",
  );
  expect(detectPotentialBreakage("<p><strong>x</strong></p>", "<p>x</p>").categories).toEqual([]);
});
test("marks unsupported VS Code features as inconclusive", async () => {
  const source = "ref[^a]\n\n[^a]: note";
  const result = detectPotentialBreakage(await renderGitHub(source), await renderVsCode(source));
  expect(result.vscodeUnsupported).toContain("footnotes");
  const alert = "> [!NOTE]\n> text";
  const alertResult = detectPotentialBreakage(await renderGitHub(alert), await renderVsCode(alert));
  expect(alertResult.vscodeUnsupported).toContain("alerts");
});
test("detects deep-list differences from both saved and local GitHub output", async () => {
  const root = `${import.meta.dir}/fixtures/e2e/cases/expanded-size-nesting`;
  const source = await Bun.file(`${root}.md`).text();
  const savedGitHub = await Bun.file(`${root}.html`).text();
  const vscode = await renderVsCode(source);
  expect(detectPotentialBreakage(savedGitHub, vscode).categories).toContain("structure");
  expect(detectPotentialBreakage(await renderGitHub(source), vscode).categories).toContain(
    "structure",
  );
});
