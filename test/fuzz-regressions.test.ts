import { expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { readdirSync } from "node:fs";
import { join } from "node:path";
import { fuzzCase } from "../scripts/fuzz-cases";
import { renderIsolated } from "../scripts/fuzz-isolated";
import { htmlAst } from "./support/html-ast";
import { normalizeGithubImages } from "./support/github-images";

const directory = join(import.meta.dir, "fixtures/fuzz");
const sha = (value: string) => createHash("sha256").update(value).digest("hex");
const canonical = (html: string) => JSON.stringify(htmlAst(html, { ignoreHighlighting: true }));
for (const name of readdirSync(directory)
  .filter((name) => /^fuzz-[a-f0-9]{24}\.json$/.test(name))
  .sort()) {
  const base = join(directory, name.slice(0, -5));
  test(`saved fuzz regression: ${name.slice(0, -5)}`, async () => {
    const metadata = await Bun.file(base + ".json").json();
    const source = await Bun.file(base + ".md").text();
    expect(sha(source)).toBe(metadata.sourceSha256);
    if (metadata.origin !== "manual") expect(source).toBe(fuzzCase(metadata.seed, metadata.index));
    const result = await renderIsolated(source);
    if (metadata.failure === "crash") {
      expect(result.timedOut).toBe(false);
      expect(result.exitCode).toBe(0);
      if (!metadata.htmlSha256) return;
    }
    const html = await Bun.file(base + ".html").text();
    expect(sha(html)).toBe(metadata.htmlSha256);
    if (metadata.failure === "mismatch") {
      expect(result.exitCode).toBe(0);
      expect(canonical(result.html)).not.toBe(canonical(normalizeGithubImages(html)));
      // A fix changes this hash and requires review and promotion to a passing case.
      expect(sha(canonical(result.html))).toBe(metadata.actualAstSha256);
    } else expect(canonical(result.html)).toBe(canonical(normalizeGithubImages(html)));
  }, 60_000);
}
