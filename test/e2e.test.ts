import { describe, expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { join } from "node:path";
import { render } from "../src/quirks/index";
import manifest from "./fixtures/e2e/manifest.json";
import { htmlAst } from "./support/html-ast";
import { normalizeGithubImages } from "./support/github-images";

const directory = join(import.meta.dir, "fixtures/e2e/cases");
const hash = (value: string) => createHash("sha256").update(value).digest("hex");

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
      const actual = await render(source, { runId });
      expect(htmlAst(actual, { ignoreHighlighting: true })).toEqual(
        htmlAst(normalizeGithubImages(expected), { ignoreHighlighting: true }),
      );
    });
  }
});
