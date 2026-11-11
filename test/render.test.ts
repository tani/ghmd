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

describe("GitHub code blocks and image proxy", () => {
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
  test("uses a caller-supplied signed image URL and rejects unsafe schemes", async () => {
    const input = "![alt](https://example.com/a.png)";
    const html = await render(input, {
      imageProxy: (url) => `https://proxy.example/${encodeURIComponent(url)}`,
    });
    expect(html).toContain(
      '<a target="_blank" rel="noopener noreferrer nofollow" href="https://proxy.example/',
    );
    expect(html).toContain(
      'data-canonical-src="https://example.com/a.png" style="max-width: 100%;"',
    );
    await expect(render(input, { imageProxy: () => "javascript:alert(1)" })).rejects.toBeInstanceOf(
      TypeError,
    );
  });
});

test("linked images retain their original link when proxied", async () => {
  const html = await render("[![alt](https://example.com/a.png)](https://example.com)", {
    imageProxy: () => "https://proxy.example/a",
  });
  expect(html).toBe(
    '<p><a href="https://example.com" rel="nofollow"><img src="https://proxy.example/a" alt="alt" data-canonical-src="https://example.com/a.png" style="max-width: 100%;"></a></p>',
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
