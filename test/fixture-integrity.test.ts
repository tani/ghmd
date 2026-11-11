import { expect, test } from "bun:test";
import manifest from "./fixtures/e2e/manifest.json";
import capturedCases from "../scripts/new-e2e-cases";

test("reference manifest covers every capture case without duplicate identities", () => {
  const ids = manifest.fixtures.flatMap((f) => f.cases.map((c) => c.id));
  const sources = manifest.fixtures.map((f) => f.sourceSha256);
  expect(new Set(ids).size).toBe(ids.length);
  expect(new Set(sources).size).toBe(sources.length);
  for (const item of capturedCases) expect(ids).toContain(item.id);
});
