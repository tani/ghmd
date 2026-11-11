/** Offline experiment: independent GitHub captures remain the reference. */
import { createHash } from "node:crypto";
import { mkdir } from "node:fs/promises";
import { join } from "node:path";
import { render as experimental } from "../src/experimental/render";
import { render as baseline } from "../src/quirks";
import { compareHtml } from "../src/compare";
import { htmlAst } from "../test/support/html-ast";
import { normalizeGithubImages } from "../test/support/github-images";
import { fuzzCase } from "./fuzz-cases";
import { htmlScopeCases } from "./html-scope-cases";
import { htmlRewriterProbes } from "./html-rewriter-probes";
import { unsafeHtmlFindings } from "../test/support/unsafe-html";
import manifest from "../test/fixtures/e2e/manifest.json";

const directory = join(import.meta.dir, "../test/fixtures/e2e/cases");
const output = join(import.meta.dir, "../experiments/html-rewriter");
await mkdir(join(output, "cases"), { recursive: true });
const canonical = (html: string) => JSON.stringify(htmlAst(html, { ignoreHighlighting: true }));
const sha = (source: string) => createHash("sha256").update(source).digest("hex");
const runId = "0123456789abcdef0123456789abcdef";
const cases: any[] = [];
let samples = 0;
for (const fixture of manifest.fixtures) {
  const source = await Bun.file(join(directory, fixture.id + ".md")).text();
  const reference = await Bun.file(join(directory, fixture.id + ".html")).text();
  if (sha(source) !== fixture.sourceSha256 || sha(reference) !== fixture.htmlSha256)
    throw Error(`Fixture hash mismatch: ${fixture.id}`);
  const expected = normalizeGithubImages(reference);
  const old = await baseline(source, { runId });
  try {
    const actual = await experimental(source, { runId });
    const streamed = await experimental(source, { runId, strategy: "stream" });
    const streamEqual = canonical(streamed) === canonical(expected);
    const categories = compareHtml(expected, actual);
    const equal = canonical(actual) === canonical(expected);
    const item: any = {
      id: fixture.id,
      astEquivalent: equal,
      streamAstEquivalent: streamEqual,
      streamSemanticDifferences: compareHtml(expected, streamed),
      semanticDifferences: categories,
      baselineAstEquivalent: canonical(old) === canonical(expected),
      baselineSemanticDifferences: compareHtml(expected, old),
    };
    if ((!equal || !streamEqual) && samples++ < 20) {
      const prefix = join(output, "cases", fixture.id);
      await Bun.write(prefix + ".md", source);
      await Bun.write(prefix + ".reference.html", expected);
      await Bun.write(prefix + ".baseline.html", old);
      await Bun.write(prefix + ".experimental.html", actual);
      await Bun.write(prefix + ".stream.html", streamed);
      item.savedExample = "cases/" + fixture.id;
    }
    cases.push(item);
  } catch (error) {
    cases.push({ id: fixture.id, error: String(error) });
  }
}
const byCategory: Record<string, number> = {};
for (const item of cases)
  for (const category of item.semanticDifferences ?? [])
    byCategory[category] = (byCategory[category] ?? 0) + 1;
const generated: any[] = [];
const generatedSamples = new Set<string>();
let streamGeneratedEqual = 0;
for (let index = 0; index < 1000; index++) {
  const source = fuzzCase(123, index);
  const old = await baseline(source, { runId });
  try {
    const actual = await experimental(source, { runId });
    const astEquivalent = canonical(actual) === canonical(old);
    const streamed = await experimental(source, { runId, strategy: "stream" });
    const streamAstEquivalent = canonical(streamed) === canonical(old);
    if (streamAstEquivalent) streamGeneratedEqual++;
    // Refresh previously saved examples even when this strategy now fixes them.
    const existingPrefix = join(output, "cases", `generated-123-${index}`);
    if (await Bun.file(existingPrefix + ".md").exists()) {
      await Bun.write(existingPrefix + ".baseline.html", old);
      await Bun.write(existingPrefix + ".stream.html", streamed);
      await Bun.write(existingPrefix + ".experimental.html", actual);
    }
    if (!astEquivalent) {
      const semanticDifferences = compareHtml(old, actual);
      const item: any = { index, semanticDifferences, streamAstEquivalent };
      const signature = semanticDifferences.join(",") || "ast-only";
      if (!generatedSamples.has(signature)) {
        generatedSamples.add(signature);
        const id = `generated-123-${index}`;
        const prefix = join(output, "cases", id);
        await Bun.write(prefix + ".md", source);
        await Bun.write(prefix + ".baseline.html", old);
        await Bun.write(prefix + ".experimental.html", actual);
        await Bun.write(prefix + ".stream.html", streamed);
        item.savedExample = "cases/" + id;
      }
      generated.push(item);
    }
  } catch (error) {
    generated.push({ index, error: String(error) });
  }
}
const holdout: any[] = [];
for (let index = 0; index < 1000; index++) {
  const source = fuzzCase(456, index);
  const old = await baseline(source, { runId });
  const streamed = await experimental(source, { runId, strategy: "stream" });
  const actual = await experimental(source, { runId });
  holdout.push({
    index,
    astEquivalent: canonical(actual) === canonical(old),
    streamAstEquivalent: canonical(streamed) === canonical(old),
  });
}
const scopeCases: any[] = [];
let scopeSamples = 0;
for (const { id, source } of htmlScopeCases()) {
  const old = await baseline(source, { runId });
  const streamed = await experimental(source, { runId, strategy: "stream" });
  const actual = await experimental(source, { runId });
  const aggressive = await experimental(source, { runId, strategy: "closure-events-aggressive" });
  const item: any = {
    id,
    source,
    astEquivalent: canonical(actual) === canonical(old),
    streamAstEquivalent: canonical(streamed) === canonical(old),
    aggressiveAstEquivalent: canonical(aggressive) === canonical(old),
    semanticDifferences: compareHtml(old, actual),
  };
  if (
    (!item.astEquivalent && scopeSamples++ < 5) ||
    (!item.aggressiveAstEquivalent && item.streamAstEquivalent)
  ) {
    const prefix = join(output, "cases", "scope-" + id);
    await Bun.write(prefix + ".md", source);
    await Bun.write(prefix + ".baseline.html", old);
    await Bun.write(prefix + ".stream.html", streamed);
    await Bun.write(prefix + ".experimental.html", actual);
    await Bun.write(prefix + ".aggressive.html", aggressive);
    item.savedExample = "cases/scope-" + id;
  }
  scopeCases.push(item);
}
const probes: any[] = [];
for (const source of htmlRewriterProbes) {
  try {
    const actual = await experimental(source, { runId });
    const findings = unsafeHtmlFindings(actual);
    probes.push({ source, findings, actual });
  } catch (error) {
    probes.push({ source, error: String(error) });
  }
}
const bundle = await Bun.build({
  entrypoints: [join(import.meta.dir, "../src/experimental/render.ts")],
  target: "bun",
  metafile: true,
});
if (!bundle.success) throw Error("Cannot inspect experimental dependency graph");
const inputs = Object.keys(bundle.metafile!.inputs);
if (inputs.some((path) => /node_modules\/(?:parse5|dompurify|happy-dom)\//.test(path)))
  throw Error("HTML parser dependency leaked into experiment");
const report = {
  description:
    "HTMLRewriter + existing cmark C quirks, without a DOM parser in the experimental rendering path",
  strategy: "closure-events",
  strategyNotes:
    "Bun onEndTag events bound native analysis scopes. Temporary closing tags are removed before output; anchors retain block-boundary closure. Math-blocking formatting is excluded because formatting reconstruction is unavailable. The broader aggressive variant is evaluated only on the scope corpus to expose regressions.",
  bun: Bun.version,
  platform: process.platform,
  arch: process.arch,
  manifestSha256: sha(await Bun.file(join(directory, "../manifest.json")).text()),
  assessment:
    "Offline HTML AST and lossy semantic comparison; browser layout and XSS safety are not established by these metrics.",
  summary: {
    total: cases.length,
    astEquivalent: cases.filter((c) => c.astEquivalent).length,
    semanticEquivalent: cases.filter((c) => c.semanticDifferences?.length === 0).length,
    semanticDifferences: cases.filter((c) => c.semanticDifferences?.length > 0).length,
    errors: cases.filter((c) => c.error).length,
    streamAstEquivalent: cases.filter((c) => c.streamAstEquivalent).length,
    baselineAstEquivalent: cases.filter((c) => c.baselineAstEquivalent).length,
    byCategory,
  },
  runtimeInputs: inputs,
  implementationHashes: Object.fromEntries(
    await Promise.all(
      [
        ...inputs,
        "scripts/experiment-html-rewriter.ts",
        "scripts/html-rewriter-probes.ts",
        "scripts/html-scope-cases.ts",
        "test/support/unsafe-html.ts",
        ...Array.from(
          new Bun.Glob("src/quirks/cmark-*.c").scanSync({ cwd: join(import.meta.dir, "..") }),
        ),
      ].map(async (path) => [path, sha(await Bun.file(join(import.meta.dir, "..", path)).text())]),
    ),
  ),
  generatedComparison: {
    reference: "current renderer, not GitHub API",
    seed: 123,
    total: 1000,
    astEquivalent: 1000 - generated.length,
    streamAstEquivalent: streamGeneratedEqual,
    regressions: generated.filter((item) => item.streamAstEquivalent).length,
    semanticEquivalent:
      1000 - generated.filter((item) => item.error || item.semanticDifferences.length).length,
    differences: generated,
  },
  heldOutGenerated: {
    reference: "current renderer, not GitHub API",
    seed: 456,
    total: holdout.length,
    astEquivalent: holdout.filter((c) => c.astEquivalent).length,
    streamAstEquivalent: holdout.filter((c) => c.streamAstEquivalent).length,
    regressions: holdout.filter((c) => !c.astEquivalent && c.streamAstEquivalent),
    differences: holdout.filter((c) => !c.astEquivalent),
  },
  malformedScopeCases: {
    reference: "current renderer, not GitHub API",
    total: scopeCases.length,
    astEquivalent: scopeCases.filter((c) => c.astEquivalent).length,
    streamAstEquivalent: scopeCases.filter((c) => c.streamAstEquivalent).length,
    regressions: scopeCases.filter((c) => !c.astEquivalent && c.streamAstEquivalent).length,
    aggressiveAstEquivalent: scopeCases.filter((c) => c.aggressiveAstEquivalent).length,
    aggressiveRegressions: scopeCases.filter(
      (c) => !c.aggressiveAstEquivalent && c.streamAstEquivalent,
    ).length,
    cases: scopeCases,
  },
  securityProbes: {
    assessment: "Limited parse5 inspection; not a browser execution or XSS safety guarantee",
    total: probes.length,
    passed: probes.filter((p) => !p.error && p.findings.length === 0).length,
    cases: probes,
  },
  cases,
};
await Bun.write(join(output, "report.json"), JSON.stringify(report, null, 2) + "\n");
console.log(JSON.stringify(report.summary, null, 2));
console.log("Generated", JSON.stringify({ total: 1000, astEquivalent: 1000 - generated.length }));
console.log(
  "Probes",
  JSON.stringify({ total: probes.length, passed: report.securityProbes.passed }),
);
console.log(
  "Holdout",
  JSON.stringify({
    total: report.heldOutGenerated.total,
    astEquivalent: report.heldOutGenerated.astEquivalent,
    streamAstEquivalent: report.heldOutGenerated.streamAstEquivalent,
    regressions: report.heldOutGenerated.regressions.length,
  }),
);
console.log(
  "Scopes",
  JSON.stringify({
    total: report.malformedScopeCases.total,
    astEquivalent: report.malformedScopeCases.astEquivalent,
    streamAstEquivalent: report.malformedScopeCases.streamAstEquivalent,
    regressions: report.malformedScopeCases.regressions,
  }),
);
console.log(`Report: ${join(output, "report.json")}`);

if (
  report.summary.errors ||
  report.summary.baselineAstEquivalent !== report.summary.total ||
  report.securityProbes.passed !== report.securityProbes.total ||
  generated.some((item) => item.error)
)
  process.exitCode = 1;
