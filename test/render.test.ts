import { describe, test, expect } from "bun:test";
import { render } from "../src/quirks/index";
import { nativeRender, nativeVersion } from "../src/quirks/native";

describe("render", () => {
  test("renders GFM and keeps task states separate from later content", async () => {
    expect(nativeVersion).toBe("0.29.0.gfm.13");
    const html = await render("- [ ] `[x]`\n- [x] done\n\n| A | B |\n|---|---|\n| 1 | 2 |");
    expect(html).toContain('aria-label="Incomplete task"');
    expect(html).toContain('aria-label="Completed task"');
    expect(html).toContain('<markdown-accessiblity-table><table role="table">');
  });
  test("renders GitHub alerts and external link attributes", async () => {
    const html = await render("> [!NOTE]\n> Message.\n\n[link](https://example.com)");
    expect(html).toContain('class="markdown-alert markdown-alert-note"');
    expect(html).toContain("<p>Message.</p>");
    expect(html).toContain('href="https://example.com" rel="nofollow"');
  });
  test("native parser retains UTF-8 source positions", () => {
    const { html, astXML } = nativeRender("あ **太字**\n\n[^a]\n\n[^a]: note");
    expect(html).toContain("<strong>太字</strong>");
    expect(Bun.XML.parse(astXML, { compact: false }).name).toBe("document");
  });
  test("rejects invalid inputs and does not leak globals across concurrent calls", async () => {
    const before = Object.getOwnPropertyDescriptor(globalThis, "document");
    const outputs = await Promise.all(["Alpha", "Beta"].map(render));
    expect(outputs).toEqual(["<p>Alpha</p>", "<p>Beta</p>"]);
    expect(Object.getOwnPropertyDescriptor(globalThis, "document")).toEqual(before);
    await expect(render(null as any)).rejects.toBeInstanceOf(TypeError);
    await expect(render("x".repeat(1000001))).rejects.toBeInstanceOf(RangeError);
  });
});

describe("HTML and math compatibility", () => {
  test("preserves safe raw HTML and removes executable attributes and URLs", async () => {
    const html = await render(
      '<details>\n<summary>More</summary>\n\nContent\n\n</details>\n\n<img src=x onerror=alert(1)>\n\n<a href="javascript:alert(1)">bad</a>',
    );
    expect(html).toContain("<details>\n<summary>More</summary>\n<p>Content</p>\n</details>");
    expect(html).not.toContain("onerror");
    expect(html).not.toContain("javascript:");
  });
  test("produces math placeholders and preserves math inside footnotes as text", async () => {
    const html = await render("$`x^2`$\n\n```math\nx^2\n```\n\nText[^a]\n\n[^a]: $a$");
    expect(html).toContain('class="js-inline-math"');
    expect(html).toContain('class="js-display-math"');
    expect(html).toContain(">$$x^2$$</math-renderer>");
    expect(html).toContain("<p>$a$ <a");
  });
  test("handles multiline footnote labels and image nodes in labels", async () => {
    const html = await render("[^a\nb]: # c\n\nxxx[^a\nb].\n\n[^a]: note\n\n[^![image](#)]");
    expect(html.match(/data-footnote-ref=/g)).toHaveLength(2);
    expect(html).toContain("<p>[^]</p>");
  });
});

describe("GitHub code blocks and images", () => {
  test("preserves code and language without syntax highlighting", async () => {
    const html = await render(
      '```mermaid\nflowchart TD\n A --> B\n```\n\n```json\n{"type":"FeatureCollection"}\n```',
    );
    expect(html).toContain(
      '<pre lang="mermaid" class="notranslate"><code class="notranslate">flowchart TD\n A --&gt; B\n</code></pre>',
    );
    expect(html).toContain(
      '<pre lang="json" class="notranslate"><code class="notranslate">{"type":"FeatureCollection"}\n</code></pre>',
    );
    expect(html).not.toContain('class="pl-');
  });
  test("preserves remote image URLs without a proxy", async () => {
    const html = await render("![alt](https://example.com/a.png)");
    expect(html).toContain('<img src="https://example.com/a.png" alt="alt">');
    expect(html).not.toContain("data-canonical-src");
  });
});

test("linked images retain their original link and image URL", async () => {
  const html = await render("[![alt](https://example.com/a.png)](https://example.com)");
  expect(html).toBe(
    '<p><a href="https://example.com" rel="nofollow"><img src="https://example.com/a.png" alt="alt"></a></p>',
  );
});

test("accepts a deterministic render ID and rejects malformed IDs", async () => {
  const runId = "1234567890abcdef1234567890abcdef";
  const html = await render("$x$ and [^a]\n\n[^a]: note", { runId });
  expect(html).toContain(`data-run-id="${runId}"`);
  expect(html).toContain(`user-content-fn-a-${runId}`);
  await expect(render("text", { runId: "bad" })).rejects.toThrow("runId must be 32");
});

test("keeps a raw anchor without href while stripping a rejected destination", async () => {
  const html = await render("<a>plain anchor</a> and [unsafe](javascript:alert(1))");
  expect(html).toContain("<a>plain anchor</a>");
  expect(html).toContain(" and unsafe");
  expect(html).not.toContain("javascript:");
  expect(html).not.toContain("data-gfm-had-href");
});

test("unwraps template elements while keeping their parsed content", async () => {
  expect(await render("<template><code>x</code></template>")).toBe(
    '<p><code class="notranslate">x</code></p>',
  );
  expect(await render("<template><b>x</b><i>y</i></template>")).toBe("<p><b>x</b><i>y</i></p>");
  expect(await render("<div><template>before <strong>x</strong> after</template></div>")).toBe(
    "<div>before <strong>x</strong> after</div>",
  );
});

test("trims the final newline only when the last raw HTML block remains open", async () => {
  expect(await render("<div>\n<b>hello\n")).toBe("<div>\n<b>hello</b></div>");
  expect(await render("<div>\n<b>hello\n</div>")).toBe("<div>\n<b>hello\n</b></div>");
  expect(await render("<pre>\nfoo\n")).toBe('<pre class="notranslate">foo</pre>');
  expect(await render("<details><summary>x</summary>\nhello\n")).toBe(
    "<details><summary>x</summary>\nhello</details>",
  );
});

test("sanitizes rejected links and attributes inside templates before unwrapping them", async () => {
  const html = await render(
    '<template><img src=x onerror=alert(1)><a href="javascript:alert(1)">x</a></template>',
  );
  expect(html).not.toContain("onerror");
  expect(html).not.toContain("javascript:");
  expect(html).toContain("</a>x</p>");
});

test("retains nested template content and the browser-parsed table structure", async () => {
  const { htmlAst } = await import("./support/html-ast");
  expect(await render("<template><template>hi <b>x</b></template></template>")).toBe(
    "<p>hi <b>x</b></p>",
  );
  const html = await render("<template><table><tr><td>x</td></tr></table></template>");
  expect(htmlAst(html)).toEqual(
    htmlAst(
      '<p><markdown-accessiblity-table></markdown-accessiblity-table></p><table role="table"><tbody><tr><td>x</td></tr></tbody></table><p></p>',
    ),
  );
});

test("keeps table handling local to the paragraph containing math", async () => {
  const { htmlAst } = await import("./support/html-ast");
  const runId = "0123456789abcdef0123456789abcdef";
  const math = `<math-renderer class="js-inline-math" style="display: inline-block" data-run-id="${runId}">$x$</math-renderer>`;
  const table = '<table role="table"><tbody><tr><td>x</td></tr></tbody></table>';
  const same = await render("<template><table><tr><td>x</td></tr></table></template> $x$", {
    runId,
  });
  expect(htmlAst(same)).toEqual(htmlAst(`<p></p>${table} ${math}`));
  const separate = await render("<template><table><tr><td>x</td></tr></table></template>\n\n$x$", {
    runId,
  });
  expect(htmlAst(separate)).toEqual(
    htmlAst(
      `<p><markdown-accessiblity-table></markdown-accessiblity-table></p>${table}<p></p>\n<p>${math}</p>`,
    ),
  );
  const nested = await render(
    "<div><template><table><tr><td>$x$</td></tr></table></template></div>",
    { runId },
  );
  expect(htmlAst(nested)).toEqual(
    htmlAst(
      `<div><markdown-accessiblity-table><table role="table"><tbody><tr><td>${math}</td></tr></tbody></table></markdown-accessiblity-table></div>`,
    ),
  );
  for (const html of [same, separate, nested]) expect(html).not.toContain("data-gfm-table-");
});

test("keeps empty safe links while removing encoded dangerous URLs", async () => {
  expect(await render('<a href="">empty</a>')).toContain('<a href="">empty</a>');
  for (const url of [
    "javascript:alert(1)",
    "jav&#x61;script:alert(1)",
    "javascript&#58;alert(1)",
    "vbscript:alert(1)",
  ]) {
    const html = await render(`<a href="${url}">bad</a>`);
    expect(html).not.toContain("href=");
    expect(html).toContain("bad");
  }
});

test("processes link markers and external destinations after sanitization", async () => {
  const html = await render(
    '<a href="https://example.com/?a=1&amp;b=2" rel="ugc">web</a> <a href="//example.com">protocol relative</a> <a href="MAILTO:a@example.com">mail</a> <a href="/local" rel="ugc">local</a> <a href="javascript:alert(1)" id="anchor">kept</a> <a href="javascript:alert(1)"><strong>unwrapped</strong></a>',
  );
  expect(html).toContain('href="https://example.com/?a=1&amp;b=2" rel="nofollow"');
  expect(html).toContain('href="//example.com" rel="nofollow"');
  expect(html).toContain('href="MAILTO:a@example.com" rel="nofollow"');
  expect(html).toContain('href="/local" rel="ugc"');
  expect(html).toContain('id="anchor">kept</a>');
  expect(html).toContain("<strong>unwrapped</strong>");
  expect(html).not.toContain("data-gfm-had-href");
});

test("inserts a default details summary only when no direct summary exists", async () => {
  const html = await render(
    "<details><div><summary>Nested</summary></div>body</details>\n\n<details><summary>Own</summary>body</details>",
  );
  expect(html).toContain(
    "<details><summary>Details</summary><div><summary>Nested</summary></div>body</details>",
  );
  expect(html).toContain("<details><summary>Own</summary>body</details>");
});

test("preserves math text boundaries across unwrapped templates", async () => {
  expect(await render("$<template>x</template>$")).toBe("<p>$x$</p>");
  expect(await render("a<template>b<template>c</template></template>d")).toBe("<p>abcd</p>");
});

test("handles details, a template table and math together", async () => {
  expect(
    await render("<details><template><table><tr><td>x</td></tr></table></template> $x$</details>", {
      runId: "a".repeat(32),
    }),
  ).toBe(
    '<details><summary>Details</summary><markdown-accessiblity-table><table role="table"><tbody><tr><td>x</td></tr></tbody></table></markdown-accessiblity-table> $x$</details>',
  );
});
