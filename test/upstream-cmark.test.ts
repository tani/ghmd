import { describe, expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { readdir } from "node:fs/promises";
import { join } from "node:path";
import manifest from "./fixtures/cmark-gfm-upstream.json";
import { renderMarkdown, renderMathHtml, finalizeHtml } from "../src/quirks/native";

const vendor = join(import.meta.dir, "../vendor/cmark-gfm");
describe("unmodified upstream cmark-gfm", () => {
  test("release source and recorded generated headers match their pinned hashes", async () => {
    expect(manifest.commit).toBe("587a12bb54d95ac37241377e6ddc93ea0e45439b");
    const actual: string[] = [];
    for (const dir of ["src", "extensions", "generated"])
      for (const entry of await readdir(join(vendor, dir), { withFileTypes: true }))
        if (entry.isFile()) actual.push(`${dir}/${entry.name}`);
    actual.push("COPYING");
    expect(actual.sort()).toEqual(Object.keys(manifest.files).sort());
    for (const [path, expected] of Object.entries(manifest.files)) {
      const bytes = await Bun.file(join(vendor, path)).arrayBuffer();
      expect(createHash("sha256").update(Buffer.from(bytes)).digest("hex")).toBe(expected);
    }
  });
  test("adjusts task states without changing literal code", () => {
    const html = renderMarkdown("- [ ] `[x]`\n- [x] yes\n\n`[^a\\]b]`");
    expect(html).toContain('<input type="checkbox" disabled="" /> <code>[x]</code>');
    expect(html).toContain('<input type="checkbox" checked="" disabled="" /> yes');
    expect(html).toContain("<code>[^a\\]b]</code>");
  });
  test("adapts nested task markers after multibyte input and embedded NUL", () => {
    const html = renderMarkdown("日本語\u0000\n\n- [x] 一\n  - [ ] 二\n  - [X] 三");
    expect(html.match(/<input type="checkbox" checked="" disabled="" \/>/g)).toHaveLength(2);
    expect(html.match(/<input type="checkbox" disabled="" \/>/g)).toHaveLength(1);
  });
  test("recovers escaped unresolved labels without modifying code", () => {
    const html = renderMarkdown("[^a\\]b] and `[^a\\]b]`");
    expect(html).toContain("[^a]b] and <code>[^a\\]b]</code>");
  });
  test("retains references whose label precedes an image", () => {
    const html = renderMarkdown("[^a![x](#)]\n\n[^a]: note");
    expect(html).toContain("data-footnote-ref");
    expect(html).toContain("<p>note ");
  });
  test("recovers multiline labels after multibyte text and embedded NUL", () => {
    const html = renderMarkdown("日本語\u0000 [^] [^a\nb]\n\n[^a]: note");
    expect(html).toContain("日本語� [^] <sup");
    expect(html).toContain('href="#fn-a"');
    expect(html).toContain("<p>note ");
  });
  test("matches inline math boundaries around Unicode whitespace", () => {
    const token = "00000000000000000000000000000000";
    const html = renderMathHtml("<p>日本語 $x$&nbsp;$y$ z$x$</p>", token, "boundary");
    expect(html.match(/class="js-inline-math"/g)).toHaveLength(2);
    expect(html).toContain(" z$x$");
    expect(renderMathHtml("<p>a$x$</p>", token, "boundary")).toBe("<p>a$x$</p>");
  });
  test("wraps serialized tables without modifying raw-text content", () => {
    expect(
      finalizeHtml(
        '<style>a{content:"<table>"}</style><table class="x" role="grid"><tr><td>1</td></tr></table>',
      ),
    ).toBe(
      '<style>a{content:"<table>"}</style><markdown-accessiblity-table><table class="x" role="table"><tr><td>1</td></tr></table></markdown-accessiblity-table>',
    );
  });
  test("leaves tables in inert template contents unchanged", () => {
    const html =
      "<template><table><tbody><tr><td>x</td></tr></tbody></table></template><table></table>";
    expect(finalizeHtml(html)).toBe(
      '<template><table><tbody><tr><td>x</td></tr></tbody></table></template><markdown-accessiblity-table><table role="table"></table></markdown-accessiblity-table>',
    );
  });
  test("flattens deep lists and normalizes moved task item classes", () => {
    const deep =
      "<ul><li>parent".repeat(10) +
      '<ul><li class="task-list-item other other">child</li></ul>' +
      "</li></ul>".repeat(10);
    const output = finalizeHtml(deep);
    expect(output.match(/<ul>/g)).toHaveLength(10);
    expect(output).toContain('<li class="other">child</li>');
    expect(output).not.toContain("task-list-item");
  });
  test("leaves deep lists in template contents unchanged", () => {
    const deep = "<ul><li>".repeat(12) + "child" + "</li></ul>".repeat(12);
    const input = `<template>${deep}</template>`;
    expect(finalizeHtml(input)).toBe(input);
  });
  test("normalizes code classes and preserves inert template content", () => {
    expect(
      finalizeHtml(
        '<pre lang="math" class="a a"><code class="x x">z</code></pre><template><code>z</code></template>',
      ),
    ).toBe(
      '<pre class="a notranslate"><code class="x notranslate">z</code></pre><template><code>z</code></template>',
    );
  });
  test("matches list spacing around rendered math", () => {
    expect(
      finalizeHtml(
        "<ul><li><math-renderer>$x$</math-renderer> y</li><li>x <em><math-renderer>$y$</math-renderer></em></li></ul>",
      ),
    ).toBe(
      "<ul><li>\n<math-renderer>$x$</math-renderer> y</li><li>x <em><math-renderer>$y$</math-renderer></em>\n</li></ul>",
    );
  });
  test("formats footnote references and sections while preserving escaped attributes", () => {
    const token = "0123456789abcdef0123456789abcdef";
    const input =
      '<p><sup class="footnote-ref"><a href="#fn-a&amp;b" id="fnref-a" data-footnote-ref="">1</a></sup></p><section class="old" data-footnotes=""><ol><li id="fn-a"><p>note <a href="#fnref-a" data-footnote-backref="" data-footnote-backref-idx="1" class="old">↩</a></p></li></ol></section>';
    const output = finalizeHtml(input, token);
    expect(output).toContain(`href="#user-content-fn-a&amp;b-${token}"`);
    expect(output).toContain(
      `id="user-content-fnref-a-${token}" data-footnote-ref="" aria-describedby="footnote-label"`,
    );
    expect(output).toContain(
      '<section data-footnotes="" class="footnotes"><h2 id="footnote-label" class="sr-only">Footnotes</h2>',
    );
    expect(output).toContain(`id="user-content-fn-a-${token}"`);
    expect(output).toContain(
      `href="#user-content-fnref-a-${token}" data-footnote-backref="" class="data-footnote-backref"`,
    );
  });
  test("applies image and checkbox corrections without changing inert contents", () => {
    const input =
      '<ul><li><input type="checkbox" checked="" disabled=""> <p><img src="a?x=1&amp;y=2"></p></li></ul><template><img src="b"><input type="checkbox"></template>';
    const output = finalizeHtml(input);
    expect(output).toContain(
      '<ul class="contains-task-list"><li class="task-list-item">\n<p><input type="checkbox" id="" disabled="" class="task-list-item-checkbox" aria-label="Completed task" checked=""> ',
    );
    expect(output).toContain(
      '<a target="_blank" rel="noopener noreferrer" href="a?x=1&amp;y=2"><img src="a?x=1&amp;y=2" style="max-width: 100%;"></a>',
    );
    expect(output).toContain('<template><img src="b"><input type="checkbox"></template>');
  });
  test("keeps code containing multiline footnote syntax literal", () => {
    const html = renderMarkdown("[^] and `[^a\nb]`\n\n[^a]: note");
    expect(html).toContain("<code>[^a b]</code>");
  });
});
