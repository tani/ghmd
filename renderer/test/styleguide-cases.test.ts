import { describe, expect, test } from "bun:test";
import { cases } from "./styleguide/cases";
import { diagnoseCases } from "./styleguide/diagnostics";
import snapshots from "./fixtures/styleguide/github-batches.json";
import { render } from "../src/index";

const separator = (id: string) => `\n\n## AUDITBOUNDARY-${id}\n\n`;
const cache = new Map<string, Promise<ReturnType<typeof diagnoseCases>>>();
function diagnostics(group: string, id: string) {
  const members = cases.filter((c) => c.group === group);
  const set = group === "footnote" ? members.filter((c) => c.id === id) : members;
  const key = set.length === 1 ? set[0].id : group;
  let promise = cache.get(key);
  if (!promise) {
    promise = (async () => {
      const source = set.map((c) => separator(c.id) + c.source).join("");
      const expected = (snapshots as Record<string, { html: string }>)[key].html;
      const runId = expected.match(
        /(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/,
      )?.[1];
      const proxy = new Map<string, string>();
      for (const match of expected.matchAll(
        /src="(https:\/\/camo\.githubusercontent\.com\/[^"\s]+)"[^>]*data-canonical-src="([^"]+)"/g,
      ))
        proxy.set(match[2], match[1]);
      const actual = await render(source, {
        runId,
        imageProxy: (url) => {
          const value = proxy.get(url);
          if (!value) throw Error(`No saved proxy URL for ${url}`);
          return value;
        },
      });
      return diagnoseCases(
        set.map((c) => c.id),
        expected,
        actual,
        true,
      );
    })();
    cache.set(key, promise);
  }
  return promise;
}
describe("style guide cases (saved GitHub batch context)", () => {
  for (const item of cases)
    test(`${item.id} [${item.rules.join(", ")}]`, async () => {
      const result = (await diagnostics(item.group, item.id)).find((x) => x.id === item.id);
      expect(result).toBeDefined();
      if (result?.status !== "pass") throw Error(JSON.stringify(result, null, 2));
    });
});
