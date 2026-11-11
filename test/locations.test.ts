import { expect, test } from "bun:test";
import { render as renderVsCode } from "../src/vscode";
import { render as renderGitHub } from "../src/quirks";
import { detectPotentialBreakage } from "../src/compare";
import { locateDifferences } from "../src/compare/locations";
import { compareMarkdown } from "gfm-bun-render/compare";

async function locate(source: string, github?: string) {
  const [actual, vscode] = await Promise.all([
    github === undefined ? renderGitHub(source) : Promise.resolve(github),
    renderVsCode(source),
  ]);
  return locateDifferences(source, detectPotentialBreakage(actual, vscode));
}

test("maps math to an exact Unicode-aware source range", async () => {
  const source = "😀あ$x$";
  const findings = await locate(source, "<p>😀あ$x$</p>");
  const f = findings.find((x) => x.category === "math")!;
  expect(f.status).toBe("exact");
  expect(f.locations).toEqual([
    {
      start: { line: 1, column: 3, offset: 3 },
      end: { line: 1, column: 6, offset: 6 },
      precision: "exact",
    },
  ]);
  expect(source.slice(f.locations[0].start.offset, f.locations[0].end.offset)).toBe("$x$");
});

test("marks repeated identical equations as ambiguous", async () => {
  const source = "$x$ and $x$";
  const findings = await locate(
    source,
    '<p><math-renderer class="js-inline-math">$x$</math-renderer> and $x$</p>',
  );
  const f = findings.find((x) => x.category === "math")!;
  expect(f.status).toBe("ambiguous");
  expect(f.locations.map((x) => x.start.column)).toEqual([1, 9]);
});

test("bounds the number of repeated location candidates", async () => {
  const source = Array.from({ length: 80 }, () => "$x$").join(" ");
  const f = (await locate(source, "<p>literal</p>")).find((x) => x.category === "math")!;
  expect(f.status).toBe("ambiguous");
  expect(f.locations).toHaveLength(64);
  expect(f.note).toContain("capped");
});

test("maps a missing link to the URL on the correct CRLF line", async () => {
  const source = "one\r\n\r\n[ https://example.com";
  const f = (await locate(source)).find((x) => x.category === "links")!;
  expect(f.status).toBe("exact");
  expect(f.locations[0].start).toEqual({ line: 3, column: 3, offset: 9 });
  expect(source.slice(f.locations[0].start.offset, f.locations[0].end.offset)).toBe(
    "https://example.com",
  );
});

test("reports a block range for a deeply nested list", async () => {
  const source = Array.from({ length: 12 }, (_, i) => `${"  ".repeat(i)}- item ${i}`).join("\n");
  const f = (await locate(source)).find((x) => x.category === "structure")!;
  expect(f.status).toBe("block");
  expect(f.locations[0].start.line).toBe(11);
  expect(f.locations[0].end.line).toBe(12);
});

test("uses block ranges for fenced math and does not invent missing locations", async () => {
  const source = "```math\nx^2\n```";
  const f = (await locate(source, "<pre><code>x^2</code></pre>")).find(
    (x) => x.category === "math",
  )!;
  expect(f.status).toBe("block");
  expect(f.locations[0].start.line).toBe(1);
  expect(f.locations[0].end.line).toBe(3);
  const absent = locateDifferences("source", {
    categories: ["text"],
    github: { text: ["a"] } as any,
    vscode: { text: ["unfindable"] } as any,
  })[0];
  expect(absent.status).toBe("unavailable");
  expect(absent.locations).toEqual([]);
});

test("identifies candidate words lost from rendered text", () => {
  const source = "hello missing world";
  const analysis = detectPotentialBreakage("<p>hello world</p>", "<p>hello missing world</p>");
  const result = locateDifferences(source, analysis).find((x) => x.category === "text")!;
  expect(result.status).toBe("exact");
  expect(source.slice(result.locations[0].start.offset, result.locations[0].end.offset)).toBe(
    "missing",
  );
});

test("locates source-authored checkbox markup", async () => {
  const source = 'before <input type="checkbox" checked> after';
  const f = (await locate(source, "<p>before  after</p>")).find((x) => x.category === "tasks")!;
  expect(f.status).toBe("exact");
  expect(source.slice(f.locations[0].start.offset, f.locations[0].end.offset)).toBe(
    '<input type="checkbox" checked>',
  );
});

test("locates HTML link and image attributes, including encoded URL characters", async () => {
  const source = '<a href="jav&#x61;script:alert(1)">bad</a>';
  const link = (await locate(source, "<p>bad</p>")).find((x) => x.category === "links")!;
  expect(link.status).toBe("exact");
  expect(source.slice(link.locations[0].start.offset, link.locations[0].end.offset)).toBe(
    'href="jav&#x61;script:alert(1)"',
  );
  const imageSource = '<img src="data:image/png;base64,YQ==">';
  const image = (await locate(imageSource, "<p></p>")).find((x) => x.category === "images")!;
  expect(image.status).toBe("exact");
  expect(imageSource.slice(image.locations[0].start.offset, image.locations[0].end.offset)).toBe(
    'src="data:image/png;base64,YQ=="',
  );
});

test("locates punctuation-only differences", async () => {
  const source = "~one~ and ~~two~~";
  const f = (await locate(source)).find((x) => x.category === "text")!;
  expect(f.status).toBe("exact");
  expect(source.slice(f.locations[0].start.offset, f.locations[0].end.offset)).toBe("~one~");
});

test("maps decoded Markdown escapes back to the source range", async () => {
  const root = `${import.meta.dir}/fixtures/e2e/cases/style-footnote-bracket-multiline`;
  const source = await Bun.file(`${root}.md`).text();
  const github = await Bun.file(`${root}.html`).text();
  const f = (await locate(source, github)).find((x) => x.category === "text")!;
  expect(f.status).toBe("exact");
  expect(source.slice(f.locations[0].start.offset, f.locations[0].end.offset)).toBe("^x\\]y");
});

test("locates nested list differences when an item contains a soft line break", async () => {
  const root = `${import.meta.dir}/fixtures/e2e/cases/quirks-probe-list-after`;
  const source = await Bun.file(`${root}.md`).text();
  const github = await Bun.file(`${root}.html`).text();
  const f = (await locate(source, github)).find((x) => x.category === "structure")!;
  expect(f.status).toBe("block");
  expect(f.locations[0].start.line).toBe(11);
});

test("CLI JSON reports the input source range", () => {
  const root = new URL("..", import.meta.url).pathname;
  const child = Bun.spawnSync([process.execPath, "scripts/compare-renderers.ts", "--json"], {
    cwd: root,
    stdin: Buffer.from("prefix $a <b > c$"),
    stdout: "pipe",
    stderr: "pipe",
  });
  expect(child.exitCode).toBe(0);
  const result = JSON.parse(child.stdout.toString());
  const f = result.findings.find((x: any) => x.category === "math");
  expect(f.locations[0].start.column).toBe(8);
  expect(f.status).toBe("exact");
  expect(result.file).toBeNull();
});

test("the compare entry point exposes location-aware findings", () => {
  const result = compareMarkdown("hello missing", "<p>hello</p>", "<p>hello missing</p>");
  expect(result.findings.find((x) => x.category === "text")?.locations[0].start.column).toBe(7);
});

test("compares typed nodes without reporting math as missing prose", async () => {
  const source = "before $x$ and `code` after";
  const { compareMarkdown } = await import("../src/compare");
  const { render } = await import("../src/vscode");
  const result = compareMarkdown(
    source,
    "<p>before $x$ and <code>code</code> after</p>",
    await render(source),
  );
  expect(result.categories).toContain("math");
  expect(result.categories).not.toContain("text");
  expect(result.vscode.nodes.map((node) => node.kind)).toEqual([
    "text",
    "math",
    "text",
    "code",
    "text",
  ]);
  expect(result.vscode.nodes.find((node) => node.kind === "math")?.source?.status).toBe("block");
  expect(result.vscode.text).not.toContain("[math]");
});

test("marks repeated node source blocks as ambiguous even without differences", async () => {
  const source = "$x$\n\n$x$";
  const { compareMarkdown } = await import("../src/compare");
  const { render } = await import("../src/vscode");
  const html = await render(source);
  const result = compareMarkdown(source, html, html);
  expect(result.categories).toEqual([]);
  const math = result.vscode.nodes.filter((node) => node.kind === "math");
  expect(math).toHaveLength(2);
  expect(math[0].source?.status).toBe("ambiguous");
  expect(math[0].source?.locations).toHaveLength(2);
});

test("matches prose across adjacent HTML blocks without flattening math into text", async () => {
  const { detectPotentialBreakage } = await import("../src/compare");
  expect(
    detectPotentialBreakage("<p>hello</p><p>world</p>", "<p>hello world</p>").categories,
  ).not.toContain("text");
  expect(
    detectPotentialBreakage("<p>hello</p><p>world</p>", "<p>hello missing world</p>").categories,
  ).toContain("text");
});
