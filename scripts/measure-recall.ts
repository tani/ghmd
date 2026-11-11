import { join } from "node:path";
import { createHash } from "node:crypto";
import { render as renderQuirks } from "../src/quirks";
import { render as renderVsCode } from "../src/vscode";
import { compareHtml, detectPotentialBreakage } from "../src/compare";
import manifest from "../test/fixtures/e2e/manifest.json";
import styleBenchmark from "../test/benchmarks/style-risk.json";
import renderingBenchmark from "../test/benchmarks/rendering-risk.json";
import vscodeBenchmark from "../test/benchmarks/vscode-github-risk.json";

const benchmark = process.argv.includes("--rendering")
  ? renderingBenchmark
  : process.argv.includes("--style")
    ? styleBenchmark
    : vscodeBenchmark;

const ids = new Set<string>();
const results = [];
for (const item of benchmark.cases) {
  if (ids.has(item.id)) throw Error(`Duplicate benchmark ID: ${item.id}`);
  ids.add(item.id);
  const fixture = manifest.fixtures.find((f) => f.cases.some((c) => c.id === item.id));
  if (!fixture) throw Error(`Missing fixture: ${item.id}`);
  const source = await Bun.file(
    join(import.meta.dir, "../test/fixtures/e2e/cases", `${fixture.id}.md`),
  ).text();
  if (createHash("sha256").update(source).digest("hex") !== item.sourceSha256)
    throw Error(`Benchmark source changed: ${item.id}`);
  const expected = await Bun.file(
    join(import.meta.dir, "../test/fixtures/e2e/cases", `${fixture.id}.html`),
  ).text();
  if (createHash("sha256").update(expected).digest("hex") !== fixture.htmlSha256)
    throw Error(`Reference changed: ${item.id}`);
  const runId = "0123456789abcdef0123456789abcdef";
  const [quirks, vscode] = await Promise.all([
    renderQuirks(source, { runId }),
    renderVsCode(source),
  ]);
  const analysis = detectPotentialBreakage(expected, vscode);
  const differences = analysis.categories;
  results.push({
    ...item,
    detected: differences.length > 0,
    differences,
    vscodeMathErrors: analysis.vscodeMathErrors,
    quirksVsGitHub: compareHtml(quirks, expected),
  });
}
const tp = results.filter((r) => r.positive && r.detected);
const fn = results.filter((r) => r.positive && !r.detected);
const fp = results.filter((r) => !r.positive && r.detected);
const tn = results.filter((r) => !r.positive && !r.detected);
const report = {
  description: benchmark.description,
  profile: benchmark.profile,
  total: results.length,
  positive: tp.length + fn.length,
  negative: fp.length + tn.length,
  truePositive: tp.length,
  falseNegative: fn.length,
  falsePositive: fp.length,
  trueNegative: tn.length,
  recall: tp.length / (tp.length + fn.length),
  precisionOnLabeledCases: tp.length / (tp.length + fp.length),
  results,
};
if (process.argv.includes("--json")) console.log(JSON.stringify(report, null, 2));
else {
  console.log(
    `Document-level recall: ${tp.length}/${tp.length + fn.length} (${(100 * report.recall).toFixed(1)}%).`,
  );
  console.log(`Selected controls flagged: ${fp.length}/${fp.length + tn.length}.`);
  for (const r of results)
    console.log(
      `${r.positive ? (r.detected ? "TP" : "FN") : r.detected ? "FP" : "TN"}\t${r.id}\t${r.rules.join(",")}\t${r.differences.join(",") || "-"}\t${r.reason}`,
    );
  console.log(`Limit: ${benchmark.description}`);
}
