import { describe, expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { join } from "node:path";
import { render } from "../src/index";
import manifest from "./fixtures/known-differences/manifest.json";

const dir = join(import.meta.dir, "fixtures/known-differences");
const sha = (s: string) => createHash("sha256").update(s).digest("hex");

describe("GitHub reference pairs: resolved renderer differences", () => {
  for (const item of manifest.pairs) {
    test(`${item.id}: input and GitHub HTML are intact`, async () => {
      expect(sha(await Bun.file(join(dir, `${item.id}.md`)).text())).toBe(item.sourceSha256);
      expect(sha(await Bun.file(join(dir, `${item.id}.html`)).text())).toBe(item.expectedSha256);
    });
    test(`${item.id}: matches the saved GitHub HTML`, async () => {
      const source = await Bun.file(join(dir, `${item.id}.md`)).text();
      const expected = await Bun.file(join(dir, `${item.id}.html`)).text();
      const runId = expected.match(
        /(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/,
      )?.[1];
      const actual = await render(source, { runId });
      expect(actual).toBe(expected);
    });
  }
});
