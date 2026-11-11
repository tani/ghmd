import { describe, expect, test } from "bun:test";
import {
  flattenTemplates,
  prepareMathTables,
  removeBoundaries,
  renderMathHtml,
} from "../src/quirks/native";

const boundary = "test-boundary";
const token = "a".repeat(32);
const tableMarker = "data-test-table";
const prepare = (html: string) =>
  prepareMathTables(flattenTemplates(html, boundary), token, boundary, tableMarker);

describe("native HTML preparation", () => {
  test("preserves text boundaries across nested templates and removed comments", () => {
    const html = prepare(
      "<p>$<template>x<template>y</template></template><!-- one --><!-- two -->$</p>",
    );
    expect(html).toBe("<p>$<!--test-boundary-->x<!--test-boundary-->y<!--test-boundary-->$</p>");
    expect(removeBoundaries(html, boundary)).toBe("<p>$xy$</p>");
    // No single text node contains the entire expression.
    expect(renderMathHtml(html, token, boundary)).toBe("<p>$xy$</p>");
  });

  test("retains comment-shaped literal text in raw text elements", () => {
    const html = prepare("<textarea><!--literal--><b>text</b></textarea><p>a<!--removed-->b</p>");
    expect(html).toBe(
      "<textarea><!--literal--><b>text</b></textarea><p>a<!--test-boundary-->b</p>",
    );
    expect(removeBoundaries(html, boundary)).toContain(
      "<textarea><!--literal--><b>text</b></textarea>",
    );
  });

  test("wraps tables according to math in their paragraph and enclosing elements", () => {
    const table = "<table><tbody><tr><td>v</td></tr></tbody></table>";
    const markedTable = table.replace("<table>", '<table data-test-table="">');
    const wrapper = (value: string) =>
      `<markdown-accessiblity-table>${value}</markdown-accessiblity-table>`;
    expect(prepare(`<p>$x$${table}</p>`)).toBe(`<p>$x$${markedTable}</p>`);
    expect(prepare(`<p>plain${table}</p>`)).toBe(`<p>plain${wrapper(markedTable)}</p>`);
    expect(prepare(`<b><p>$x$${table}</p></b>`)).toBe(`<b><p>$x$${wrapper(markedTable)}</p></b>`);
    expect(prepare(`<details><p>$x$${table}</p></details>`)).toBe(
      `<details><p>$x$${markedTable}</p></details>`,
    );
    expect(prepare(`<b><details><p>$x$${table}</p></details></b>`)).toBe(
      `<b><details><p>$x$${markedTable}</p></details></b>`,
    );
    expect(prepare(`<p>$x$${table}</p><p>plain${table}</p>`)).toBe(
      `<p>$x$${markedTable}</p><p>plain${wrapper(markedTable)}</p>`,
    );
  });
});
