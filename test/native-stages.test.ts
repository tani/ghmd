import { expect, test } from "bun:test";
import { resolve } from "node:path";
import { prepareHtmlStage, renderMathStage, finishHtml } from "./support/native-stages";

const token = "a".repeat(32);
const boundary = "test-boundary";
const tableMarker = "data-test-table";

test("HTML preparation preserves template text boundaries unless details require reparsing", () => {
  expect(prepareHtmlStage("<p>日本語</p>", boundary)).toEqual({
    html: "<p>日本語</p>",
    needsReparse: false,
  });
  const templates = prepareHtmlStage("<p>a<template>b</template>c</p>", boundary);
  expect(templates.needsReparse).toBe(false);
  expect(templates.html).toBe("<p>a<!--test-boundary-->b<!--test-boundary-->c</p>");
  const details = prepareHtmlStage("<p>a<template>b</template></p><details>x</details>", boundary);
  expect(details).toEqual({
    html: "<p>ab</p><details><summary>Details</summary>x</details>",
    needsReparse: true,
  });
});

test("table wrapping alone does not request DOM reparsing", () => {
  const table = "<table><tbody><tr><td>v</td></tr></tbody></table>";
  const noMath = renderMathStage(`<p>plain${table}</p>`, token, boundary, tableMarker);
  expect(noMath.needsReparse).toBe(false);
  expect(noMath.html).toContain('<markdown-accessiblity-table><table data-test-table="">');
  const math = renderMathStage(`<p>$x$${table}</p>`, token, boundary, tableMarker);
  expect(math.needsReparse).toBe(true);
  expect(math.html).toContain("<math-renderer");
  expect(math.html).not.toContain("<markdown-accessiblity-table>");
});

test("removed text-boundary markers alone do not request DOM reparsing", () => {
  expect(renderMathStage("<p>$<!--test-boundary-->x$</p>", token, boundary, tableMarker)).toEqual({
    html: "<p>$x$</p>",
    needsReparse: false,
  });
});

test("the final stage removes rejected anchors before deciding image wrapping", () => {
  const html = finishHtml(
    '<a data-gfm-had-href="test-boundary"><img src="local.png"></a>',
    token,
    boundary,
    tableMarker,
  );
  expect(html).toBe(
    '<p><a target="_blank" rel="noopener noreferrer" href="local.png"><img src="local.png" style="max-width: 100%;"></a></p>',
  );
});

test("ordinary rendering makes one native call across every reparse path", () => {
  const root = resolve(import.meta.dir, "..");
  const samples = [
    "simple",
    "<details>body</details>",
    "$x$",
    "<details><p>$x$</p></details>",
    "<template><table><tr><td>x</td></tr></table></template> $x$",
  ];
  const script = `
    import { mock } from "bun:test";
    const ffi = await import("bun:ffi");
    const compile = ffi.cc;
    let calls = 0, compiles = 0;
    mock.module("bun:ffi", () => ({ ...ffi, cc(options) {
      compiles++;
      const library = compile(options);
      const invoke = library.symbols.gfm_bun_call;
      return { ...library, symbols: { ...library.symbols, gfm_bun_call(...args) {
        if (args.some(value => typeof value === "function")) throw Error("Unexpected JS callback");
        calls++; return invoke(...args);
      } } };
    } }));
    const { render } = await import("./src/quirks/index.ts");
    if (compiles !== 1) throw Error("Renderer must compile exactly one C library");
    const counts = [];
    for (const source of ${JSON.stringify(samples)}) {
      calls = 0;
      await render(source, { runId: "${token}" });
      counts.push(calls);
    }
    process.stdout.write(JSON.stringify(counts));
  `;
  const result = Bun.spawnSync([process.execPath, "--no-env-file", "-e", script], {
    cwd: root,
    stdout: "pipe",
    stderr: "pipe",
  });
  expect(result.stderr.toString()).toBe("");
  expect(result.exitCode).toBe(0);
  expect(JSON.parse(result.stdout.toString())).toEqual(samples.map(() => 1));
});

test("the final C stage removes exactly one trailing LF", () => {
  for (const source of [
    "",
    "text",
    "text\n",
    "text\n\n",
    "text\r\n",
    "\n\r\n",
    "text\n\r",
    "text\u2028",
  ]) {
    expect(finishHtml(source, token, boundary, tableMarker)).toBe(source.replace(/\n$/, ""));
  }
});

test("the final C stage enforces the byte limit after trimming and HTML expansion", () => {
  const limit = 32 * 1024 * 1024;
  // Character count is far below the byte limit; UTF-8 bytes fill it exactly.
  const exact = "あ".repeat(Math.floor(limit / 3)) + "x".repeat(limit % 3);
  expect(finishHtml(exact + "\n", token, boundary, tableMarker)).toBe(exact);
  let failure: unknown;
  try {
    finishHtml(exact + "x", token, boundary, tableMarker);
  } catch (error) {
    failure = error;
  }
  expect(failure).toBeInstanceOf(RangeError);
  expect((failure as Error).message).toBe("Rendered HTML exceeds 32 MB");
  // The input fits, but relative-image wrapping pushes the final HTML over.
  const expanded = "x".repeat(limit - 40) + '<img src="x">';
  expect(Buffer.byteLength(expanded)).toBeLessThan(limit);
  expect(() => finishHtml(expanded, token, boundary, tableMarker)).toThrow(RangeError);
  expect(finishHtml("ok\n", token, boundary, tableMarker)).toBe("ok");
}, 60_000);
