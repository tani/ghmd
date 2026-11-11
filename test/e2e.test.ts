import { describe, expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { join } from "node:path";
import { parseFragment } from "parse5";
import { render } from "../src/quirks/index";
import manifest from "./fixtures/e2e/manifest.json";
import { htmlAst } from "./support/html-ast";

const directory = join(import.meta.dir, "fixtures/e2e/cases");
const hash = (value: string) => createHash("sha256").update(value).digest("hex");
function imageUrls(html: string): Map<string, string> {
  const result = new Map<string, string>();
  const visit = (node: any) => {
    if (node.tagName === "img") {
      const attrs = new Map<string, string>(node.attrs.map((attr: any) => [attr.name, attr.value]));
      const original = attrs.get("data-canonical-src");
      const proxied = attrs.get("src");
      if (original && proxied?.startsWith("https://camo.githubusercontent.com/"))
        result.set(original, proxied);
    }
    for (const child of node.childNodes ?? []) visit(child);
  };
  visit(parseFragment(html));
  return result;
}

describe("independent GitHub Markdown API references", () => {
  for (const fixture of manifest.fixtures) {
    const aliases = fixture.cases.map((c) => c.id).join(", ");
    test(aliases, async () => {
      const source = await Bun.file(join(directory, `${fixture.id}.md`)).text();
      const expected = await Bun.file(join(directory, `${fixture.id}.html`)).text();
      expect(hash(source)).toBe(fixture.sourceSha256);
      expect(hash(expected)).toBe(fixture.htmlSha256);
      const runId = expected.match(
        /(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/,
      )?.[1];
      const urls = imageUrls(expected);
      const actual = await render(source, {
        runId,
        imageProxy: (url) => {
          const proxied = urls.get(url);
          if (!proxied) throw new Error(`No saved GitHub image URL for ${url}`);
          return proxied;
        },
      });
      expect(htmlAst(actual, { ignoreHighlighting: true })).toEqual(
        htmlAst(expected, { ignoreHighlighting: true }),
      );
    });
  }
});
