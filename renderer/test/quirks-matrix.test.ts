import { describe, expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { join } from "node:path";
import { render } from "../src/index";
import { diagnoseCases } from "./styleguide/diagnostics";
import manifest from "./fixtures/quirks-matrix/manifest.json";

const directory = join(import.meta.dir, "fixtures/quirks-matrix");
const sha = (value: string) => createHash("sha256").update(value).digest("hex");
const cache = new Map<string, Promise<ReturnType<typeof diagnoseCases>>>();
function resultFor(group: (typeof manifest.groups)[number]) {
  let pending = cache.get(group.id);
  if (!pending) {
    pending = (async () => {
      const source = await Bun.file(join(directory, `${group.id}.md`)).text();
      const expected = await Bun.file(join(directory, `${group.id}.html`)).text();
      const id = expected.match(
        /(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/,
      )?.[1];
      const actual = await render(source, { runId: id });
      return diagnoseCases(
        group.cases.map((c) => c.id),
        expected,
        actual,
        true,
        "QUIRK-",
      );
    })();
    cache.set(group.id, pending);
  }
  return pending;
}
describe("GitHub quirks matrix (batch context)", () => {
  for (const group of manifest.groups) {
    test(`${group.id}: captured inputs and output are intact`, async () => {
      expect(sha(await Bun.file(join(directory, `${group.id}.md`)).text())).toBe(
        group.sourceSha256,
      );
      expect(sha(await Bun.file(join(directory, `${group.id}.html`)).text())).toBe(
        group.expectedSha256,
      );
    });
    for (const entry of group.cases)
      test(`${entry.id}: ${JSON.stringify(entry.source)}`, async () => {
        const result = (await resultFor(group)).find((d) => d.id === entry.id);
        expect(result).toBeDefined();
        if (result?.status !== "pass") throw Error(JSON.stringify(result, null, 2));
      });
  }
});
