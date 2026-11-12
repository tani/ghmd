import { expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { parseFragment, serialize, type HtmlNode } from "../src/html";
import { removeRawInputs } from "./support/native";
import provenance from "../vendor/lexbor/PROVENANCE.json";

const nodes = (root: HtmlNode): HtmlNode[] => [
  root,
  ...(root.content ? nodes(root.content) : []),
  ...(root.childNodes ?? []).flatMap(nodes),
];

test("vendored HTML parser is reproducible from the pinned upstream amalgamation", async () => {
  const generated = await Bun.file(new URL("../vendor/lexbor/lexbor.h", import.meta.url)).text();
  const hash = (value: string) => createHash("sha256").update(value).digest("hex");
  expect(hash(generated)).toBe(provenance.generatedSha256);
  expect(
    generated.match(/gfm_lexbor_create_element_original\(lxb_html_tree_t \*tree,/g),
  ).toHaveLength(1);
  const original = generated.replace(
    "gfm_lexbor_create_element_original(lxb_html_tree_t *tree,",
    "lxb_html_tree_create_element_for_token(lxb_html_tree_t *tree,",
  );
  expect(hash(original)).toBe(provenance.upstreamSha256);
  expect(original).not.toMatch(/\b(?:FLT|DBL|LDBL)_[A-Z0-9_]+\b/);
});

test("preserves HTML5 repair, foster parenting, templates and SVG casing", () => {
  expect(serialize(parseFragment("<b><i>x</b>y</i>"))).toBe("<b><i>x</i></b><i>y</i>");
  expect(serialize(parseFragment("<table>before<tr><td>x</td></tr></table>"))).toBe(
    "before<table><tbody><tr><td>x</td></tr></tbody></table>",
  );
  expect(serialize(parseFragment("<tr><td>x</td></tr>"))).toBe("<tr><td>x</td></tr>");
  const tree = parseFragment(
    '<template><svg viewBox="0 0 2 2"><linearGradient id="g"/></svg></template>',
  );
  expect(nodes(tree).find((n) => n.tagName === "svg")?.attrs).toEqual([
    { name: "viewBox", value: "0 0 2 2" },
  ]);
  expect(nodes(tree).some((n) => n.tagName === "linearGradient")).toBe(true);
  expect(serialize(tree)).toContain('<template><svg viewBox="0 0 2 2"><linearGradient');
});

test("source offsets count UTF-16 and preserve exact quoted and duplicate attributes", () => {
  for (const tag of [
    '<a HREF="x>y" href="z>q">',
    '<a href=x"y title=">">',
    '<a href="日本語" data-x="😀">',
  ]) {
    const source = `😀\r\n${tag}label</a>`;
    const anchor = nodes(parseFragment(source, { sourceCodeLocationInfo: true })).find(
      (n) => n.tagName === "a",
    )!;
    const loc = anchor.sourceCodeLocation!;
    expect(source.slice(loc.startTag.startOffset, loc.startTag.endOffset)).toBe(tag);
    expect(source.slice(loc.endTag.startOffset, loc.endTag.endOffset)).toBe("</a>");
    expect(source.slice(loc.attrs.href.startOffset, loc.attrs.href.endOffset)).toBe(
      tag.match(/HREF="x>y"|href=x"y|href="日本語"/)![0],
    );
  }
});

test("reconstructed formatting keeps provenance while generated wrappers have none", () => {
  const tree = parseFragment("<b><i>x</b>y</i><table><tr><td>z</td></tr></table>", {
    sourceCodeLocationInfo: true,
  });
  const italic = nodes(tree).filter((n) => n.tagName === "i");
  expect(italic).toHaveLength(2);
  expect(italic[0].sourceCodeLocation.startOffset).toBe(italic[1].sourceCodeLocation.startOffset);
  expect(italic[0].sourceCodeLocation.endTag).toBeUndefined();
  expect(italic[1].sourceCodeLocation.endTag).toBeDefined();
  expect(nodes(tree).find((n) => n.tagName === "tbody")?.sourceCodeLocation).toBeUndefined();
});

test("source metadata safely encodes malformed attribute names and raw input spans", () => {
  expect(() =>
    parseFragment('<p a"b="x" a\\b="y">ok</p>', { sourceCodeLocationInfo: true }),
  ).not.toThrow();
  const source = '😀<template><input title="a>b" type=checkbox><b>keep</b></template>';
  expect(removeRawInputs(source)).toBe("😀<template><b>keep</b></template>");
});

test("native parser errors do not leave the next call unusable", () => {
  for (const value of [null, 1, {}, Buffer.from("x")])
    expect(() => parseFragment(value as any)).toThrow(TypeError);
  expect(() => parseFragment("x".repeat(32 * 1024 * 1024 + 1))).toThrow(RangeError);
  expect(() => parseFragment("<div>".repeat(520))).toThrow(RangeError);
  expect(serialize(parseFragment("<p>ok &amp; 日本語</p>"))).toBe("<p>ok &amp; 日本語</p>");
});

test("serializes default XML namespaces and prefixed foreign attributes", () => {
  const source =
    '<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><a xlink:href="#part"><text>link</text></a></svg>';
  expect(serialize(parseFragment(source))).toBe(source);
});
