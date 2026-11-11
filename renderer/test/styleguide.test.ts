import { describe, expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { cases } from "./styleguide/cases";
import snapshots from "./fixtures/styleguide/github-batches.json";
import acid from "./styleguide/acid-manifest.json";

const separator = (id: string) => `\n\n## AUDITBOUNDARY-${id}\n\n`;
const sha = (s: string) => createHash("sha256").update(s).digest("hex");
describe("style guide audit fixtures", () => {
  test("every synthetic case has a unique identifier and rule mapping", () => {
    expect(new Set(cases.map((c) => c.id)).size).toBe(cases.length);
    expect(cases.every((c) => c.rules.length > 0 && c.source.length > 0)).toBe(true);
  });
  test("every offline GitHub snapshot corresponds to the exact current input", () => {
    const expectedIds: string[] = [];
    for (const group of [...new Set(cases.map((c) => c.group))]) {
      const members = cases.filter((c) => c.group === group);
      for (const set of group === "footnote" ? members.map((c) => [c]) : [members]) {
        const id = set.length === 1 ? set[0].id : group;
        const source = set.map((c) => separator(c.id) + c.source).join("");
        expectedIds.push(id);
        expect(
          (snapshots as Record<string, { sourceSha256: string; html: string }>)[id]?.sourceSha256,
        ).toBe(sha(source));
        expect(
          (snapshots as Record<string, { sourceSha256: string; html: string }>)[id]?.html.length,
        ).toBeGreaterThan(0);
      }
    }
    expect(Object.keys(snapshots).sort()).toEqual(expectedIds.sort());
  });
  test("pins all 24 upstream inputs to the reviewed revision", () => {
    expect(acid.commit).toBe("1a672974525509f26da28abbae230999b49e2f3d");
    const names = Object.keys(acid.files).sort();
    expect(names).toHaveLength(24);
    expect(names.every((name, i) => name.startsWith(String(i + 1).padStart(2, "0") + "-"))).toBe(
      true,
    );
  });
});
