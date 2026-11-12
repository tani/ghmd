import { describe, expect, test } from "bun:test";
import { prepareHtmlStage, renderMathStage } from "./support/native";

const boundary = "test-boundary";
const token = "a".repeat(32);
const tableMarker = "data-test-table";
const math = `<math-renderer class="js-inline-math" style="display: inline-block" data-run-id="${token}">$x$</math-renderer>`;
const processMath = (html: string) => renderMathStage(html, token, boundary, tableMarker);

describe("native HTML preparation", () => {
  test("preserves text boundaries across nested templates and removed comments", () => {
    const prepared = prepareHtmlStage(
      "<p>$<template>x<template>y</template></template><!-- one --><!-- two -->$</p>",
      boundary,
    );
    expect(prepared.needsReparse).toBe(false);
    expect(prepared.html).not.toContain("<template>");
    expect(prepared.html).toContain("$<!--test-boundary-->x<!--test-boundary-->y");
    // The expression spans separate text nodes, so it must stay literal.
    expect(processMath(prepared.html)).toEqual({ html: "<p>$xy$</p>", needsReparse: false });
  });

  test("retains comment-shaped literal text in raw text elements", () => {
    const prepared = prepareHtmlStage(
      "<textarea><!--literal--><b>text</b></textarea><p>a<!--removed-->b</p>",
      boundary,
    );
    expect(processMath(prepared.html)).toEqual({
      html: "<textarea><!--literal--><b>text</b></textarea><p>ab</p>",
      needsReparse: false,
    });
  });

  test("wraps tables according to math in their paragraph and enclosing elements", () => {
    const table = "<table><tbody><tr><td>v</td></tr></tbody></table>";
    const markedTable = table.replace("<table>", '<table data-test-table="">');
    const wrapper = (value: string) =>
      `<markdown-accessiblity-table>${value}</markdown-accessiblity-table>`;
    expect(processMath(`<p>$x$${table}</p>`).html).toBe(`<p>${math}${markedTable}</p>`);
    expect(processMath(`<p>plain${table}</p>`).html).toBe(`<p>plain${wrapper(markedTable)}</p>`);
    expect(processMath(`<b><p>$x$${table}</p></b>`).html).toBe(
      `<b><p>$x$${wrapper(markedTable)}</p></b>`,
    );
    expect(processMath(`<details><p>$x$${table}</p></details>`).html).toBe(
      `<details><p>${math}${markedTable}</p></details>`,
    );
    expect(processMath(`<b><details><p>$x$${table}</p></details></b>`).html).toBe(
      `<b><details><p>${math}${markedTable}</p></details></b>`,
    );
    expect(processMath(`<p>$x$${table}</p><p>plain${table}</p>`).html).toBe(
      `<p>${math}${markedTable}</p><p>plain${wrapper(markedTable)}</p>`,
    );
  });
});
