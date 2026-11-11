import { createHash } from "node:crypto";
import { readdir, readFile, writeFile, mkdir } from "node:fs/promises";
import { resolve, join } from "node:path";
import { cases } from "../test/styleguide/cases";
import { render } from "../src/index";
import { diagnoseCases, normalizeGeneratedIds } from "../test/styleguide/diagnostics";

const root = resolve(import.meta.dir, "..");
const args = process.argv.slice(2),
  live = args.includes("--live"),
  capture = args.includes("--capture");
const acidOnly = args.includes("--acid-only");
const isolateIndex = args.indexOf("--isolate");
const isolate = isolateIndex < 0 ? undefined : args[isolateIndex + 1];
if (isolateIndex >= 0 && !isolate) throw Error("--isolate requires a case ID or upstream filename");
const upstreamIndex = args.indexOf("--acid-dir");
const referenceIndex = args.indexOf("--reference-report");
const referencePath = referenceIndex < 0 ? undefined : args[referenceIndex + 1];
if (referenceIndex >= 0 && !referencePath) throw Error("--reference-report requires a path");
const referenceReport = referencePath
  ? JSON.parse(await readFile(resolve(referencePath), "utf8"))
  : undefined;
const onlyIndex = args.indexOf("--only");
const only = onlyIndex < 0 ? undefined : args[onlyIndex + 1];
if (onlyIndex >= 0 && !only) throw Error("--only requires a bundle ID");
const upstream = upstreamIndex < 0 ? undefined : resolve(args[upstreamIndex + 1] ?? "");
if (capture && !live) throw Error("--capture requires --live");
if (acidOnly && !upstream) throw Error("--acid-only requires --acid-dir");
if (isolate && !live && !referenceReport)
  throw Error(
    "--isolate requires --live or --reference-report with an independently captured result",
  );
if (isolate && only) throw Error("--isolate and --only cannot be used together");
if (isolate && acidOnly && cases.some((c) => c.id === isolate))
  throw Error("--acid-only cannot select a synthetic case");
const snapshots = join(root, "test/fixtures/styleguide/github-batches.json");
const sha = (text: string) => createHash("sha256").update(text).digest("hex");
const generatedId = (text: string) =>
  text.match(/(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/)?.[1];
const canonical = normalizeGeneratedIds;
const separator = (id: string) => `\n\n## AUDITBOUNDARY-${id}\n\n`;
const bundles: {
  id: string;
  source: string;
  caseIds: string[];
  kind: string;
  wrapped?: boolean;
}[] = [];
const isolatedSources = new Map(cases.map((c) => [c.id, c.source]));
for (const group of [...new Set(cases.map((x) => x.group))]) {
  const members = cases.filter((x) => x.group === group);
  // References have document-wide scope: each footnote scenario is isolated.
  const sets = group === "footnote" ? members.map((x) => [x]) : [members];
  for (const set of sets)
    bundles.push({
      id: set.length === 1 ? set[0].id : group,
      source: set.map((x) => separator(x.id) + x.source).join(""),
      caseIds: set.map((x) => x.id),
      kind: "synthetic",
      wrapped: true,
    });
}
if (upstream) {
  const files = (await readdir(upstream)).filter((x) => /^\d\d-.*\.md$/.test(x)).sort();
  if (
    files.length !== 24 ||
    files.some((f, i) => !f.startsWith(String(i + 1).padStart(2, "0") + "-"))
  )
    throw Error("acid-dir must contain all 24 files from the pinned commit");
  // The repository has no license file at the pinned revision. Inputs remain
  // outside this project. Hashes record exactly which bytes were used.
  const manifest = JSON.parse(
    await readFile(join(root, "test/styleguide/acid-manifest.json"), "utf8"),
  );
  const entries = await Promise.all(
    files.map(async (file) => ({ file, source: await readFile(join(upstream, file), "utf8") })),
  );
  for (const entry of entries) {
    if (sha(entry.source) !== manifest.files[entry.file])
      throw Error(`Upstream fixture hash mismatch: ${entry.file}`);
    isolatedSources.set(entry.file, entry.source);
  }
  for (let i = 0; i < entries.length;) {
    const number = i + 1;
    // Case 06 duplicates a footnote label; case 20 contains unclosed HTML.
    const count =
      number === 6 || number === 20
        ? 1
        : Math.min(5, number < 6 ? 6 - number : number < 20 ? 20 - number : 25 - number);
    const set = entries.slice(i, i + count);
    bundles.push({
      id: `acid-${String(number).padStart(2, "0")}-${String(i + count).padStart(2, "0")}`,
      kind: "upstream-batch",
      caseIds: set.map((x) => x.file),
      wrapped: set.length > 1,
      source:
        set.length === 1 ? set[0].source : set.map((x) => separator(x.file) + x.source).join(""),
    });
    i += count;
  }
  const sixth = entries[5];
  const variants = sixth.source
    .split(/\n---\s*(?:\n|$)/)
    .filter((part: string) => part.includes("[^a]:"));
  if (variants.length !== 2) throw Error("Unexpected structure in upstream case 06");
  for (let i = 0; i < variants.length; i++)
    bundles.push({
      id: `acid-06-isolated-${i + 1}`,
      kind: "upstream-variant",
      caseIds: [sixth.file],
      source: variants[i].trim() + "\n",
    });
}
if (isolate) {
  const source = isolatedSources.get(isolate);
  if (source === undefined) throw Error(`Unknown isolated case: ${isolate}`);
  bundles.length = 0;
  bundles.push({
    id: `isolated:${isolate}`,
    source,
    caseIds: [isolate],
    kind: cases.some((c) => c.id === isolate) ? "isolated-synthetic" : "isolated-upstream",
  });
}
let saved: any = {};
try {
  saved = JSON.parse(await readFile(snapshots, "utf8"));
} catch {}
const observed: any[] = [],
  captured: any = {};
let apiUnavailable = "";
for (const bundle of bundles) {
  if (only && bundle.id !== only && !bundle.caseIds.includes(only)) continue;
  if (acidOnly && !bundle.kind.startsWith("upstream-") && bundle.kind !== "isolated-upstream")
    continue;
  let expected: string | undefined,
    reason = "";
  if (live && apiUnavailable) {
    reason = apiUnavailable;
  } else if (live) {
    try {
      const headers: Record<string, string> = {
        Accept: "application/vnd.github+json",
        "Content-Type": "application/json",
        "X-GitHub-Api-Version": "2022-11-28",
      };
      const token = process.env.GITHUB_PAT || process.env.GITHUB_TOKEN;
      if (token) headers.Authorization = `Bearer ${token}`;
      const response = await fetch("https://api.github.com/markdown", {
        method: "POST",
        headers,
        body: JSON.stringify({ text: bundle.source, mode: "gfm" }),
        signal: AbortSignal.timeout(30000),
      });
      if (response.status === 403 || response.status === 429) {
        apiUnavailable = `GitHub returned HTTP ${response.status}; remaining requests skipped`;
        throw Error(apiUnavailable);
      }
      if (!response.ok) throw Error(`HTTP ${response.status}`);
      expected = await response.text();
    } catch (e) {
      reason = String(e);
    }
  } else if (
    referenceReport?.results?.some(
      (x: any) =>
        x.id === bundle.id &&
        x.sourceSha256 === sha(bundle.source) &&
        typeof x.expected === "string",
    )
  ) {
    expected = referenceReport.results.find(
      (x: any) => x.id === bundle.id && x.sourceSha256 === sha(bundle.source),
    ).expected;
  } else if (bundle.kind === "synthetic" && saved[bundle.id]?.sourceSha256 === sha(bundle.source))
    expected = saved[bundle.id].html;
  else
    reason = bundle.kind.startsWith("upstream-")
      ? "Run with --live to obtain a GitHub result for upstream inputs"
      : "No matching saved GitHub response";
  if (capture && expected && bundle.kind === "synthetic")
    captured[bundle.id] = { sourceSha256: sha(bundle.source), html: expected };
  let actual: string | undefined,
    status = "not-run";
  if (expected !== undefined) {
    try {
      const camo = new Map<string, string>();
      for (const match of expected.matchAll(
        /src="(https:\/\/camo\.githubusercontent\.com\/[^"\s]+)"[^>]*data-canonical-src="([^"]+)"/g,
      ))
        camo.set(match[2], match[1]);
      actual = await render(bundle.source, {
        runId: generatedId(expected),
        imageProxy: (url) => {
          const result = camo.get(url);
          if (!result) throw Error(`No saved Camo URL for ${url}`);
          return result;
        },
      });
      status = canonical(actual) === canonical(expected) ? "pass" : "fail";
    } catch (e) {
      status = "fail";
      reason = String(e);
    }
  }
  const diagnostics = diagnoseCases(bundle.caseIds, expected, actual, !!bundle.wrapped).map(
    (d) => ({ ...d, rules: cases.find((c) => c.id === d.id)?.rules ?? [] }),
  );
  if (status === "pass" && diagnostics.some((d) => d.status === "not-run")) {
    status = "fail";
    reason = "Case boundaries could not be verified; inspect the full document.";
  }
  observed.push({
    id: bundle.id,
    caseIds: bundle.caseIds,
    kind: bundle.kind,
    status,
    source:
      bundle.kind === "synthetic" || bundle.kind === "isolated-synthetic"
        ? bundle.source
        : undefined,
    sourceSha256: sha(bundle.source),
    rules: cases.filter((c) => bundle.caseIds.includes(c.id)).flatMap((c) => c.rules),
    expected,
    actual,
    reason,
    diagnostics,
  });
  console.log(`${status.padEnd(8)} ${bundle.id} (${bundle.caseIds.length} cases)`);
  for (const d of diagnostics) {
    if (only && only !== bundle.id && only !== d.id) continue;
    console.log(
      `  ${d.status.padEnd(7)} ${d.id} [${d.scope}]${d.firstDifference ? ` line ${d.firstDifference.line}, column ${d.firstDifference.column}` : ""}${d.reason ? ` — ${d.reason}` : ""}`,
    );
    if (d.firstDifference) {
      const start = Math.max(0, d.firstDifference.column - 61);
      console.log(`    expected: ${d.firstDifference.expected.slice(start, start + 180)}`);
      console.log(`    actual:   ${d.firstDifference.actual.slice(start, start + 180)}`);
    }
  }
}
if (only && !observed.length) throw Error(`Unknown bundle: ${only}`);
if (capture) {
  await mkdir(resolve(snapshots, ".."), { recursive: true });
  await writeFile(snapshots, JSON.stringify({ ...saved, ...captured }, null, 2) + "\n");
}
if (!only && !acidOnly && !isolate) {
  const source = "Text[^a]\n\n[^a]: note";
  try {
    const first = await render(source),
      second = await render(source);
    const a = first.match(/id="(user-content-fnref-[^"]+)"/)?.[1],
      b = second.match(/id="(user-content-fnref-[^"]+)"/)?.[1];
    const status = a && b && a !== b ? "pass" : "fail";
    observed.push({
      id: "footnote-id-collisions",
      caseIds: ["RENDER-003", "RENDER-006"],
      kind: "local",
      status,
      source,
      expected: "Distinct reference IDs for separate renders",
      actual: JSON.stringify([a, b]),
    });
    console.log(`${status.padEnd(8)} footnote-id-collisions`);
  } catch (e) {
    observed.push({
      id: "footnote-id-collisions",
      kind: "local",
      status: "fail",
      reason: String(e),
    });
  }
  try {
    await render("x".repeat(1_000_001));
    observed.push({
      id: "input-size-limit",
      kind: "local",
      status: "fail",
      expected: "RangeError for input above 1 MB",
      actual: "accepted",
    });
    console.log("fail     input-size-limit");
  } catch (e) {
    const status = e instanceof RangeError ? "pass" : "fail";
    observed.push({
      id: "input-size-limit",
      kind: "local",
      status,
      source: "x repeated 1,000,001 times",
      expected: "RangeError for input above 1 MB",
      actual: String(e),
    });
    console.log(`${status.padEnd(8)} input-size-limit`);
  }
}
const report = {
  date: new Date().toISOString(),
  target: "GitHub REST /markdown, mode=gfm, HTML fragment",
  renderer: `Bun ${Bun.version}; gfm-bun-render ${JSON.parse(await readFile(join(root, "package.json"), "utf8")).version}`,
  configuration: {
    live,
    isolate: isolate ?? null,
    referenceReport: referencePath ?? null,
    acidDir: upstream ?? null,
    upstreamCommit: "1a672974525509f26da28abbae230999b49e2f3d",
    normalization: "Only GitHub per-render footnote/math IDs; saved Camo URLs reused",
  },
  results: observed,
  notRun: [
    {
      group: "visual widgets",
      status: "not-run",
      reason:
        "HTML fragments cannot verify rendering, layout, accessibility, resizing, fonts, mobile, export, or Mermaid runtime.",
    },
    {
      group: "streaming",
      status: "unsupported",
      reason: "Renderer accepts complete documents only; there is no streaming interface.",
    },
    {
      group: "historical complexity advisory",
      status: "not-run",
      reason: "No adversarial runtime isolation or resource benchmark configured.",
    },
  ],
};
const out = process.env.STYLEGUIDE_REPORT ?? "/tmp/gfm-styleguide-report.json";
await writeFile(out, JSON.stringify(report, null, 2) + "\n");
console.log(`Report: ${out}`);
if (observed.some((x) => x.status === "fail" || x.status === "not-run")) process.exitCode = 1;
