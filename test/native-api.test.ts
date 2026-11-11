import { expect, test } from "bun:test";
import {
  renderMarkdown,
  renderAstXml,
  unicodeLower,
  decodeHtml,
  finalizeHtml,
  renderMathHtml,
} from "../src/quirks/native-core";

const source = "before <b>raw</b> after\n";

test("native string API validates arguments without exposing pointers", () => {
  expect(renderMarkdown("")).toBe("");
  expect(renderMarkdown("日本語\0😀")).toBe("<p>日本語�😀</p>\n");
  expect(unicodeLower("A\0ΟΣ😀")).toBe("a\0ος😀");
  expect(decodeHtml("&amp; 日本語 &#x1f600;")).toBe("& 日本語 😀");
  expect(decodeHtml("a\0b")).toBe("a\0b");
  expect(unicodeLower("\ud800")).toBe("�");
  for (const value of [null, undefined, 1, {}, Buffer.from("x")]) {
    expect(() => renderMarkdown(value as any)).toThrow(TypeError);
    expect(() => unicodeLower(value as any)).toThrow(TypeError);
  }
  expect(() => renderMathHtml("<p>x</p>", undefined as any, "boundary")).toThrow(TypeError);
});

test("optional native modifiers retain their defaults even when a later one is supplied", () => {
  const html = '<table data-table=""><tbody><tr><td>x</td></tr></tbody></table>';
  expect(finalizeHtml(html, undefined, "data-table")).toBe(
    '<table role="table"><tbody><tr><td>x</td></tr></tbody></table>',
  );
});

test("raw HTML callbacks exchange strings while C owns the AST", () => {
  const seen: string[] = [];
  const html = renderMarkdown(source, (raw) => {
    seen.push(raw);
    return raw.replace("b", "i");
  });
  expect(seen).toEqual(["<b>", "</b>"]);
  expect(html).toBe("<p>before <i>raw</i> after</p>\n");
  const xml = renderAstXml(source, (raw) => raw.replace("b", "i"));
  expect(Bun.XML.parse(xml, { compact: false }).name).toBe("document");
  expect(xml).toContain("&lt;i&gt;");
});

test("callback exceptions preserve identity and leave later calls usable", () => {
  const marker = new Error("sentinel callback error");
  for (let i = 0; i < 30; i++) {
    let caught: unknown;
    try {
      renderMarkdown(source, () => {
        throw marker;
      });
    } catch (error) {
      caught = error;
    }
    expect(caught).toBe(marker);
    expect(renderMarkdown("ok")).toBe("<p>ok</p>\n");
  }
  expect(() => renderMarkdown(source, (() => 1) as any)).toThrow(TypeError);
  expect(() => renderMarkdown(source, (async (raw) => raw) as any)).toThrow(TypeError);
  expect(() => renderMarkdown(source, "bad" as any)).toThrow(TypeError);
  for (const reason of [undefined, null, 0, "callback failure", { sentinel: true }]) {
    let caught: unknown = Symbol("not thrown");
    try {
      renderMarkdown(source, () => {
        throw reason;
      });
    } catch (error) {
      caught = error;
    }
    expect(caught).toBe(reason);
  }
  expect(renderMarkdown(source)).toBe("<p>before <b>raw</b> after</p>\n");
});

test("callbacks can reenter the native renderer without sharing document state", () => {
  let nested = "";
  const result = renderMarkdown(source, (raw) => {
    nested = renderMarkdown("**nested**", () => "");
    return raw;
  });
  expect(nested).toBe("<p><strong>nested</strong></p>\n");
  expect(result).toBe("<p>before <b>raw</b> after</p>\n");
});
