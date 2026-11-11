/** Generate reproducible cases; use --github to capture real API discrepancies. */
import { createHash } from "node:crypto";
import { mkdir } from "node:fs/promises";
import { join } from "node:path";
import { fuzzCase } from "./fuzz-cases";
import { renderIsolated } from "./fuzz-isolated";
import { fetchGitHubHtml } from "../src/compare/github";
import { htmlAst } from "../test/support/html-ast";
import { normalizeGithubImages } from "../test/support/github-images";

const args = process.argv.slice(2);
let seed = 1,
  count = 20,
  live = false,
  delay = 750;
for (let i = 0; i < args.length; i++) {
  const arg = args[i];
  if (arg === "--github") live = true;
  else if (["--seed", "--count", "--delay"].includes(arg) && args[i + 1]) {
    const n = Number(args[++i]);
    if (!Number.isSafeInteger(n)) throw Error(`Invalid ${arg}`);
    if (arg === "--seed") seed = n;
    else if (arg === "--count") count = n;
    else delay = n;
  } else throw Error(`Unknown or incomplete option: ${arg}`);
}
if (seed < 0 || seed > 0xffffffff || count < 1 || count > 1000 || delay < 750)
  throw Error("seed must be 0..4294967295, count 1..1000 and delay at least 750 ms");
const token = process.env.GITHUB_PAT ?? process.env.GITHUB_TOKEN;
if (live && !token) throw Error("Set GITHUB_PAT or GITHUB_TOKEN for --github");
const root = join(import.meta.dir, "../test/fixtures/fuzz");
const sha = (value: string) => createHash("sha256").update(value).digest("hex");
const canonical = (html: string) => JSON.stringify(htmlAst(html, { ignoreHighlighting: true }));
await mkdir(root, { recursive: true });
let captured = 0;
for (let index = 0; index < count; index++) {
  const source = fuzzCase(seed, index);
  const id = `fuzz-${sha(source).slice(0, 24)}`;
  const path = join(root, id);
  if (await Bun.file(path + ".json").exists()) {
    console.log(`Skipping saved ${id}`);
    continue;
  }
  // Obtain the independent reference before rendering, so a native crash can
  // still be saved as a Markdown/HTML pair in live mode.
  let reference: string | undefined;
  if (live) {
    try {
      reference = await fetchGitHubHtml(source, { token });
    } catch (error) {
      console.error(
        `Stopped at ${id}: ${error instanceof Error ? error.message : "API request failed"}`,
      );
      process.exitCode = 1;
      break;
    }
  }
  const { html: actual, stderr, exitCode: code, timedOut } = await renderIsolated(source);
  const failure =
    timedOut || code !== 0
      ? "crash"
      : reference !== undefined && canonical(actual) !== canonical(normalizeGithubImages(reference))
        ? "mismatch"
        : undefined;
  if (failure) {
    await Bun.write(path + ".md", source);
    if (reference !== undefined) await Bun.write(path + ".html", reference);
    await Bun.write(
      path + ".json",
      JSON.stringify(
        {
          id,
          seed,
          index,
          failure,
          ...(reference === undefined
            ? {}
            : { reference: "GitHub REST /markdown (gfm)", capturedAt: new Date().toISOString() }),
          sourceSha256: sha(source),
          ...(reference === undefined ? {} : { htmlSha256: sha(reference) }),
          ...(failure === "mismatch"
            ? { actualAstSha256: sha(canonical(actual)) }
            : { exitCode: code, timedOut, error: stderr.slice(0, 500) }),
        },
        null,
        2,
      ) + "\n",
    );
    console.log(`Saved ${failure}: ${id}`);
    captured++;
    process.exitCode = 1;
  }
  if (live && index + 1 < count) await Bun.sleep(delay);
}
console.log(`Generated up to ${count} cases (seed ${seed}); saved ${captured} failure(s).`);
