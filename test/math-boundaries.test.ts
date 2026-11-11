import { describe, expect, test } from "bun:test";
import { render } from "../src/quirks/index";

// These short, independently authored cases exercise behaviors observed in
// the GitHub GFM reference batches. The ID is fixed for stable comparisons.
const runId = "0123456789abcdef0123456789abcdef";
const math = (body: string, display = false) =>
  `<math-renderer class="js-${display ? "display" : "inline"}-math" style="display: ${display ? "block" : "inline-block"}" data-run-id="${runId}">${body}</math-renderer>`;
describe("GitHub math boundary behavior", () => {
  test("leaves a second adjacent bare formula literal", async () => {
    expect(await render("$a$$b$", { runId })).toBe(`<p>${math("$a$")}$b$</p>`);
  });
  test("distinguishes bare and protected formulas adjacent to letters", async () => {
    const html = await render("a$x$ ($x$) $x$a $`x`$a a$`x`$", { runId });
    expect(html).toBe(
      `<p>a$x$ (${math("$x$")}) $x$a ${math("$x$")}a a$<code class="notranslate">x</code>$</p>`,
    );
  });
  test("keeps escaped HTML inside protected math as escaped text", async () => {
    expect(await render("$`a <b > c`$", { runId })).toBe(
      `<p>${math("$a &amp;lt;b &amp;gt; c$")}</p>`,
    );
  });
  test("handles protected math inside nested emphasis", async () => {
    expect(await render("***value $`x`$ here***", { runId })).toBe(
      `<p><em><strong>value ${math("$x$")} here</strong></em></p>`,
    );
  });
  test("does not convert display delimiters inside details", async () => {
    expect(await render("<details>\n\n$$\nx\n$$\n\n</details>", { runId })).toBe(
      "<details><summary>Details</summary>\n<p>$$<br>\nx<br>\n$$</p>\n</details>",
    );
  });
  test("recovers a later expression after rejecting an earlier candidate", async () => {
    expect(await render(String.raw`$\sqrt{\$4}$, $x = \$$`, { runId })).toBe(
      `<p>$\\sqrt{$4}$, ${math("$x = $")}$</p>`,
    );
  });
});

test("preserves multiple display blocks and escapes fenced math", async () => {
  const html = await render("$$\na\n$$\n$$\nb\n$$\n\n```math\nx$y\n```", {
    runId: "00000000000000000000000000000000",
  });
  expect(html.match(/class="js-display-math"/g)).toHaveLength(3);
  expect(html).toContain("$$x\\$y$$");
});

test("recognizes an empty display-math text block", async () => {
  const { displayMathBlocks } = await import("../src/quirks/native");
  expect(displayMathBlocks("$$\n\n$$")).toEqual(["$$\n\n$$"]);
});

test("adds details summaries before reading fenced math text", async () => {
  const html = await render('<div>\n<pre lang="math"><details></details></pre>\n</div>', { runId });
  expect(html).toContain("$$Details$$");
});
