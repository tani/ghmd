import { expect, test } from "bun:test";
import { parseWithSourceMap } from "../src/compare/source-map";
import { locateDifferences } from "../src/compare/locations";
import { detectPotentialBreakage } from "../src/compare";
import { render } from "../src/vscode";

test("uses the plugin's top-level source lines and falls back for fences", () => {
  const source = "# Title\n\nparagraph $x$\n\n```text\ncode\n```";
  const parsed = parseWithSourceMap(source);
  expect(parsed.blocks.map((block) => block.map)).toEqual([
    [0, 1],
    [2, 3],
  ]);
  expect(parsed.tokens.some((token) => token.type === "fence" && token.map[0] === 4)).toBe(true);
});

test("does not trust source-line attributes supplied by raw HTML", () => {
  const source = '<p data-source-line="3">spoof</p>\n\nreal';
  const parsed = parseWithSourceMap(source);
  expect(parsed.blocks).toHaveLength(1);
  expect(parsed.blocks[0].map).toEqual([2, 3]);
  expect(parsed.blocks[0].html).toBe("real");
});

test("uses rendered scope to exclude hidden duplicate math", async () => {
  const source = '<div aria-hidden="true">\n\n$x$\n\n</div>\n\n$x$';
  const result = detectPotentialBreakage("", await render(source));
  const finding = locateDifferences(source, result).find((x) => x.category === "math")!;
  expect(finding.status).toBe("exact");
  expect(finding.locations).toHaveLength(1);
  expect(finding.locations[0].start.line).toBe(7);
});

test("source map instrumentation does not alter preview output", async () => {
  const source = "# Title\n\n$x$";
  const before = await render(source);
  parseWithSourceMap(source);
  expect(await render(source)).toBe(before);
  expect(before).not.toContain("data-source-line");
});
