import { expect, test } from "bun:test";
import { mkdtemp, rm, writeFile } from "node:fs/promises";
import { tmpdir } from "node:os";
import { join } from "node:path";
import { fuzzCase } from "../scripts/fuzz-cases";
import { renderIsolated } from "../scripts/fuzz-isolated";

test("fuzz inputs are deterministic, varied and combine risky features", () => {
  const samples = Array.from({ length: 40 }, (_, index) => fuzzCase(123, index));
  expect(samples).toEqual(Array.from({ length: 40 }, (_, index) => fuzzCase(123, index)));
  expect(new Set(samples).size).toBe(samples.length);
  expect(fuzzCase(124, 0)).not.toBe(samples[0]);
  for (const source of samples) {
    expect(source).toMatch(/\$/);
    expect(source).toMatch(/<(?:b|a|span|details|table|template|input)/);
    expect(
      source.split("\n").filter((line) => /^\s*(?:>\s*)?(?:[-+]|1\.) /.test(line)).length,
    ).toBeGreaterThanOrEqual(8);
  }
});

test("isolates a successful render, an exit and a timeout", async () => {
  expect(await renderIsolated("hello")).toMatchObject({
    html: "<p>hello</p>",
    exitCode: 0,
    timedOut: false,
  });
  const dir = await mkdtemp(join(tmpdir(), "yagfm-fuzz-"));
  try {
    const crash = join(dir, "crash.ts"),
      hang = join(dir, "hang.ts");
    await writeFile(crash, "process.exit(42)");
    await writeFile(hang, "await Bun.sleep(30000)");
    expect(await renderIsolated("", { worker: crash })).toMatchObject({
      exitCode: 42,
      timedOut: false,
    });
    expect(await renderIsolated("", { worker: hang, timeoutMs: 100 })).toMatchObject({
      timedOut: true,
    });
    const envCheck = join(dir, "environment.ts");
    await writeFile(
      envCheck,
      "process.stdout.write(String(Boolean(process.env.GITHUB_PAT || process.env.GITHUB_TOKEN)))",
    );
    const saved = process.env.GITHUB_PAT;
    try {
      process.env.GITHUB_PAT = "test-token";
      expect((await renderIsolated("", { worker: envCheck })).html).toBe("false");
    } finally {
      if (saved === undefined) delete process.env.GITHUB_PAT;
      else process.env.GITHUB_PAT = saved;
    }
  } finally {
    await rm(dir, { recursive: true, force: true });
  }
});
