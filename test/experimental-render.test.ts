import { expect, test } from "bun:test";
import { join } from "node:path";
import { render } from "../src/experimental/render";
import { htmlAst } from "./support/html-ast";
import { normalizeGithubImages } from "./support/github-images";
import { unsafeHtmlFindings } from "./support/unsafe-html";

test("experimental rendering loads no external HTML parser or sanitizer", async () => {
  const result = await Bun.build({
    entrypoints: [join(import.meta.dir, "../src/experimental/render.ts")],
    target: "bun",
    metafile: true,
  });
  expect(result.success).toBe(true);
  expect(
    Object.keys(result.metafile!.inputs).filter((path) => path.includes("node_modules/")),
  ).toEqual([]);
});

for (const id of [
  "style-alerts-five",
  "style-tasks-markers-nesting",
  "style-table-unicode-empty",
  "style-footnote-case-repeat",
  "quirks-probe-svg-children",
  "expanded-sanitize-entity_url",
  "expanded-html-context-nested_bad",
  "expanded-html-context-link_unclosed",
  "expanded-followup-misnested_anchor",
]) {
  test(`experimental representative GitHub capture: ${id}`, async () => {
    const root = join(import.meta.dir, "fixtures/e2e/cases", id);
    const source = await Bun.file(root + ".md").text();
    const expected = normalizeGithubImages(await Bun.file(root + ".html").text());
    expect(htmlAst(await render(source), { ignoreHighlighting: true })).toEqual(
      htmlAst(expected, { ignoreHighlighting: true }),
    );
  });
}

test("experimental screening removes encoded script URLs and event attributes", async () => {
  for (const source of [
    '<a href="java&#x09;script:alert(1)">x</a>',
    '<a href="javascript&colon;alert(1)">x</a>',
    "<svg><foreignObject><img src=x onerror=alert(1)></foreignObject></svg>",
    '<template><a href="javascript:alert(1)">x</a></template>',
  ])
    expect(unsafeHtmlFindings(await render(source))).toEqual([]);
});

test("experimental API validates input and keeps generated markers private", async () => {
  expect(await render("")).toBe("");
  await expect(render(null as any)).rejects.toBeInstanceOf(TypeError);
  await expect(render("x", { runId: "invalid" })).rejects.toBeInstanceOf(TypeError);
  await expect(render("x".repeat(1_000_001))).rejects.toBeInstanceOf(RangeError);
  const html = await render("$x$", { runId: "a".repeat(32) });
  expect(html).toContain(`data-run-id="${"a".repeat(32)}"`);
  expect(html).not.toContain("data-gfm-");
});

test("implicit scope hints preserve unclosed math-blocking formatting", async () => {
  const { render: baseline } = await import("../src/quirks");
  const options = { runId: "a".repeat(32) };
  for (const source of [
    "<b><em>one</b> two\n\n$x$\n\nlast",
    "<i><strong>one</i> two\n\n$x$\n\nlast",
    "<em><code>one</em> two\n\n$x$\n\nlast",
  ])
    expect(htmlAst(await render(source, options))).toEqual(
      htmlAst(await baseline(source, options)),
    );
});

test("synthetic closing tags never escape native analysis", async () => {
  const html = await render("<b><i>one</b> two</i>\n\n$x$");
  expect(html).not.toContain("data-gfm-scope-");
  expect(html).toContain("<math-renderer");
});
