import assert from "node:assert/strict";
import { spawnSync } from "node:child_process";
import { createHash } from "node:crypto";
import { mkdtempSync, readFileSync, readdirSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fileURLToPath } from "node:url";
import { Worker } from "node:worker_threads";
import { test } from "node:test";
import { cc } from "gfm-bun-render/ffi";
import { render } from "gfm-bun-render";
import { render as renderVSCode } from "gfm-bun-render/vscode";
import { compareMarkdown } from "gfm-bun-render/compare";
import { parseFragment, serialize } from "../src/html/index.ts";
import { htmlAst } from "./support/html-ast.ts";
import { normalizeGithubImages } from "./support/github-images.ts";

const source = new URL("./support/ffi-probe.c", import.meta.url);
const runId = "a".repeat(32);
const root = fileURLToPath(new URL("../", import.meta.url));
const hash = (value: string) => createHash("sha256").update(value).digest("hex");

test("Node cc compiles scalar ABIs, wide integers, strings and buffers", () => {
  const library = cc({
    source,
    define: { OFFSET: "1" },
    flags: ["-DMULTIPLIER=2"],
    symbols: {
      ffi_add: { args: [5, "int"], returns: "int" },
      ffi_half: { args: ["double"], returns: "double" },
      ffi_not: { args: ["bool"], returns: "bool" },
      ffi_signed: { args: ["i64"], returns: "i64" },
      ffi_unsigned: { args: ["u64"], returns: "u64" },
      ffi_greeting: { returns: "cstring" },
      ffi_sum: { args: ["buffer", "u64"], returns: "u64" },
      ffi_fill: { args: ["buffer", "u32", "u8"] },
      ffi_pointer: { args: ["buffer"], returns: "ptr" },
      ffi_byte: { args: ["ptr"], returns: "u8" },
    },
  });
  try {
    const s = library.symbols;
    assert.equal(s.ffi_add(2, 3), 12);
    assert.equal(s.ffi_half(3), 1.5);
    assert.equal(s.ffi_not(true), false);
    for (const value of [1n, -(1n << 63n)]) assert.equal(s.ffi_signed(value), value);
    for (const value of [1n, (1n << 64n) - 1n]) assert.equal(s.ffi_unsigned(value), value);
    assert.equal(s.ffi_greeting(), "Hello 日本語 😀");
    const bytes = Buffer.from([1, 2, 250]);
    assert.equal(s.ffi_sum(bytes, 3n), 253n);
    s.ffi_fill(bytes, 3, 9);
    assert.deepEqual([...bytes], [9, 9, 9]);
    const pointer = s.ffi_pointer(bytes);
    assert.equal(typeof pointer, "number");
    assert.equal(s.ffi_byte(pointer), 9);
    assert.equal(s.ffi_sum(new Uint8Array([4, 5]), 2n), 9n);
  } finally {
    library.close();
  }
  library.close();
  assert.throws(() => library.symbols.ffi_add(1, 2), TypeError);
});

test("compile and symbol errors leave subsequent cc calls usable", () => {
  const directory = mkdtempSync(join(tmpdir(), "gfm-cc-test-"));
  try {
    const broken = join(directory, "broken.c");
    writeFileSync(broken, "int broken( { return 0; }");
    assert.throws(() => cc({ source: broken, symbols: {} }), /compilation failed/);
    assert.throws(
      () => cc({ source, symbols: { absent: { returns: "int" } } }),
      /symbol|function/i,
    );
    assert.throws(
      () => cc({ source, symbols: { ffi_add: { args: ["buffer_length"] } } }),
      /Unsupported.*ABI/,
    );
    assert.throws(
      () => cc({ source, symbols: { bad: { args: ["napi_value"], returns: "napi_value" } } }),
      /Unsupported Node-API signature/,
    );
    const library = cc({ source, symbols: { ffi_add: { args: ["int", "int"], returns: "int" } } });
    assert.equal(library.symbols.ffi_add(1, 2), 3);
    library.close();
  } finally {
    rmSync(directory, { recursive: true, force: true });
  }
});

test("missing optional peers produce an actionable error", () => {
  const result = spawnSync(
    process.execPath,
    [
      "--input-type=module",
      "-e",
      `
    import Module from 'node:module';
    const original = Module._resolveFilename;
    Module._resolveFilename = function(name, ...rest) {
      if (name === 'node-tinycc' || name === 'koffi') throw Error('Missing optional test peer');
      return original.call(this, name, ...rest);
    };
    const {cc} = await import(${JSON.stringify(new URL("../src/ffi/index.ts", import.meta.url).href)});
    try { cc({source:'unused.c',symbols:{}}); process.exitCode=1; }
    catch(error) { if (!error.message.includes('npm install node-tinycc koffi')) throw error; }
  `,
    ],
    { encoding: "utf8" },
  );
  assert.equal(result.status, 0, result.stderr);
});

test("fresh Node rendering needs no compiler executable or subprocess", () => {
  const result = spawnSync(
    process.execPath,
    [
      "--input-type=module",
      "-e",
      `
    import childProcess from 'node:child_process';
    import {syncBuiltinESMExports} from 'node:module';
    for (const name of ['spawn','spawnSync','exec','execSync','execFile','execFileSync','fork'])
      childProcess[name] = () => { throw Error('Unexpected subprocess'); };
    syncBuiltinESMExports();
    const {render} = await import(${JSON.stringify(new URL("../src/quirks/index.ts", import.meta.url).href)});
    const html = await render('日本語');
    if (html !== '<p>日本語</p>') throw Error(html);
  `,
    ],
    { encoding: "utf8", env: { ...process.env, PATH: "/nonexistent", CC: "/nonexistent/cc" } },
  );
  assert.equal(result.status, 0, result.stderr);
});

test("the complete runtime-independent C renderer works through Koffi", () => {
  const library = cc({
    source: new URL("./support/native-core-probe.c", import.meta.url),
    include: [
      join(root, "vendor/lexbor/compat"),
      ...["src", "extensions", "generated"].map((directory) =>
        join(root, "vendor/cmark-gfm", directory),
      ),
    ],
    symbols: { gfm_api_probe: { returns: "int" } },
  });
  try {
    assert.equal(library.symbols.gfm_api_probe(), 0);
  } finally {
    library.close();
  }
});

test("one compiled library can expose both Node-API and scalar functions", () => {
  const directory = mkdtempSync(join(tmpdir(), "gfm-cc-mixed-"));
  try {
    const mixed = join(directory, "mixed.c");
    writeFileSync(
      mixed,
      `#include "${join(root, "src/html/native.c")}"
#include "${fileURLToPath(source)}"
`,
    );
    const library = cc({
      source: mixed,
      include: join(root, "vendor/lexbor/compat"),
      symbols: {
        gfm_html_parse: { args: ["napi_env", "napi_value", "napi_value"], returns: "napi_value" },
        ffi_add: { args: ["int", "int"], returns: "int" },
      },
    });
    try {
      assert.equal(library.symbols.ffi_add(1, 2), 3);
      assert.equal(
        JSON.parse(library.symbols.gfm_html_parse(undefined, "x", false)).childNodes[0].value,
        "x",
      );
    } finally {
      library.close();
    }
    assert.throws(() => library.symbols.ffi_add(1, 2), TypeError);
  } finally {
    rmSync(directory, { recursive: true, force: true });
  }
});

test("Node-API passes strings and exceptions, including missing arguments", async () => {
  assert.equal(await render("日本語\0😀"), "<p>日本語�😀</p>");
  assert.equal(await render(""), "");
  await assert.rejects(render(42 as any), TypeError);
  await assert.rejects(render("x", { runId: "invalid" }), TypeError);
  await assert.rejects(render("x".repeat(1_000_001)), RangeError);
  const library = cc({
    source: new URL("../src/html/native.c", import.meta.url),
    include: join(root, "vendor/lexbor/compat"),
    symbols: {
      gfm_html_parse: { args: ["napi_env", "napi_value", "napi_value"], returns: "napi_value" },
    },
  });
  try {
    const parse = library.symbols.gfm_html_parse;
    assert.throws(() => parse(undefined), TypeError);
    assert.throws(() => parse(undefined, 1, true), TypeError);
    assert.equal(JSON.parse(parse(undefined, "<p>ok</p>", false)).childNodes[0].tagName, "p");
  } finally {
    library.close();
  }
  assert.throws(() => library.symbols.gfm_html_parse(undefined, "", false), TypeError);
});

test("Node HTML source locations survive Unicode and formatting reconstruction", () => {
  const html = '😀<b><i title="a>b">x</b>y</i>';
  const tree = parseFragment(html, { sourceCodeLocationInfo: true });
  const bold = tree.childNodes!.find((node) => node.tagName === "b")!;
  const first = bold.childNodes![0],
    second = tree.childNodes!.find((node) => node.tagName === "i")!;
  assert.equal(first.sourceCodeLocation!.startOffset, 5);
  assert.equal(second.sourceCodeLocation!.startOffset, 5);
  const attr = second.sourceCodeLocation!.attrs.title;
  assert.equal(html.slice(attr.startOffset, attr.endOffset), 'title="a>b"');
  assert.equal(first.sourceCodeLocation!.endTag, undefined);
  assert.ok(second.sourceCodeLocation!.endTag);
  assert.equal(serialize(tree), '😀<b><i title="a>b">x</i></b><i title="a>b">y</i>');
});

test("Node comparison uses the same math plugin and source mappings", async () => {
  const markdown = "# Heading\n\n$x^2$\n";
  const vscode = await renderVSCode(markdown);
  assert.match(vscode, /katex/);
  assert.doesNotMatch(vscode, /data-source-line/);
  const analysis = compareMarkdown(markdown, await render(markdown), vscode);
  assert.deepEqual(analysis.categories, []);
  assert.equal(analysis.vscode.nodes[0].source!.status, "block");
});

test("Node backend rejects worker compilation before loading the legacy addon", async () => {
  const url = new URL("../src/ffi/index.ts", import.meta.url).href;
  const message = await new Promise((resolve, reject) => {
    const worker = new Worker(
      `
      const {parentPort} = require('node:worker_threads');
      import(${JSON.stringify(url)}).then(({cc}) => {
        try { cc({source: ${JSON.stringify(fileURLToPath(source))}, symbols: {}}); parentPort.postMessage('unexpected success'); }
        catch(error) { parentPort.postMessage(error.message); }
      }).catch(error => { throw error; });
    `,
      { eval: true },
    );
    worker.once("message", resolve);
    worker.once("error", reject);
    worker.once("exit", (code) => {
      if (code !== 0) reject(new Error(`Worker exited ${code}`));
    });
  });
  assert.match(String(message), /main thread/);
});

const manifest = JSON.parse(
  readFileSync(new URL("./fixtures/e2e/manifest.json", import.meta.url), "utf8"),
);
for (const fixture of manifest.fixtures) {
  test(`Node GitHub reference: ${fixture.id}`, async () => {
    const base = join(root, "test/fixtures/e2e/cases", fixture.id);
    const markdown = readFileSync(base + ".md", "utf8");
    const expected = readFileSync(base + ".html", "utf8");
    assert.equal(hash(markdown), fixture.sourceSha256);
    assert.equal(hash(expected), fixture.htmlSha256);
    const referenceId = expected.match(
      /(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/,
    )?.[1];
    const actual = await render(markdown, { runId: referenceId ?? runId });
    assert.deepEqual(
      htmlAst(actual, { ignoreHighlighting: true }),
      htmlAst(normalizeGithubImages(expected), { ignoreHighlighting: true }),
    );
  });
}

const fuzzDirectory = join(root, "test/fixtures/fuzz");
for (const name of readdirSync(fuzzDirectory)
  .filter((name) => /^fuzz-[a-f0-9]{24}\.json$/.test(name))
  .sort()) {
  test(`Node saved fuzz regression: ${name.slice(0, -5)}`, async () => {
    const base = join(fuzzDirectory, name.slice(0, -5));
    const metadata = JSON.parse(readFileSync(base + ".json", "utf8"));
    const markdown = readFileSync(base + ".md", "utf8");
    assert.equal(hash(markdown), metadata.sourceSha256);
    const actual = await render(markdown, { runId });
    if (!metadata.htmlSha256) return;
    const expected = readFileSync(base + ".html", "utf8");
    assert.equal(hash(expected), metadata.htmlSha256);
    const canonical = htmlAst(actual, { ignoreHighlighting: true });
    const reference = htmlAst(normalizeGithubImages(expected), { ignoreHighlighting: true });
    if (metadata.failure === "mismatch") {
      assert.notDeepEqual(canonical, reference);
      assert.equal(hash(JSON.stringify(canonical)), metadata.actualAstSha256);
    } else assert.deepEqual(canonical, reference);
  });
}
