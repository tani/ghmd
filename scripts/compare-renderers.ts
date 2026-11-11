import { join } from "node:path";
import { render as renderQuirks } from "../src/quirks";
import { render as renderVsCode } from "../src/vscode";
import { compareHtml, detectPotentialBreakage, type SemanticKind } from "../src/compare";
import manifest from "../test/fixtures/e2e/manifest.json";
import { fetchGitHubHtml } from "../src/compare/github";

const args = process.argv.slice(2);
let fixtures = false,
  json = false,
  all = false,
  live = false,
  only: string | undefined,
  kind: SemanticKind | undefined;
const paths: string[] = [];
const kinds: SemanticKind[] = [
  "text",
  "math",
  "code",
  "links",
  "images",
  "tasks",
  "structure",
  "formatting",
];
for (let i = 0; i < args.length; i++) {
  const arg = args[i];
  if (arg === "--fixtures") fixtures = true;
  else if (arg === "--json") json = true;
  else if (arg === "--all") all = true;
  else if (arg === "--github-live") live = true;
  else if (arg === "--only" && args[i + 1]) {
    only = args[++i];
    fixtures = true;
  } else if (arg === "--kind" && args[i + 1]) {
    const value = args[++i] as SemanticKind;
    if (!kinds.includes(value)) throw Error(`Unknown category: ${value}`);
    kind = value;
  } else if (arg.startsWith("-") && arg !== "-") throw Error(`Unknown option: ${arg}`);
  else paths.push(arg);
}
if (paths.length > 1 || (fixtures && paths.length) || (all && !fixtures) || (live && fixtures))
  throw new Error(
    "Usage: bun run compare [--fixtures | --only ID | input.md|-] [--kind CATEGORY] [--all] [--github-live] [--json]",
  );
const token = "0123456789abcdef0123456789abcdef";
async function compare(source: string, expected?: string) {
  const [quirks, vscode] = await Promise.all([
    renderQuirks(source, { runId: token }),
    renderVsCode(source),
  ]);
  const analysis = detectPotentialBreakage(expected ?? quirks, vscode);
  const categories = analysis.categories;
  const a = analysis.github,
    b = analysis.vscode;
  return {
    categories,
    vscodeMathErrors: analysis.vscodeMathErrors,
    vscodeUnsupported: analysis.vscodeUnsupported,
    ...(expected === undefined
      ? {}
      : {
          quirksVsGitHub: compareHtml(quirks, expected),
          vscodeVsGitHub: compareHtml(vscode, expected),
        }),
    differences: Object.fromEntries(categories.map((k) => [k, { github: a[k], vscode: b[k] }])),
  };
}
if (fixtures) {
  const cases = [];
  const mathErrors: { id: string; errors: string[] }[] = [];
  const unsupported: { id: string; features: string[] }[] = [];
  const selected = only
    ? manifest.fixtures.filter((f) => f.id === only || f.cases.some((c) => c.id === only))
    : manifest.fixtures;
  if (only && !selected.length) throw new Error(`Unknown fixture: ${only}`);
  for (const fixture of selected) {
    const root = join(import.meta.dir, "../test/fixtures/e2e/cases", fixture.id);
    const result = await compare(
      await Bun.file(root + ".md").text(),
      await Bun.file(root + ".html").text(),
    );
    if (result.vscodeMathErrors.length)
      mathErrors.push({ id: fixture.id, errors: result.vscodeMathErrors });
    if (result.vscodeUnsupported.length)
      unsupported.push({ id: fixture.id, features: result.vscodeUnsupported });
    if (result.categories.length && (!kind || result.categories.includes(kind)))
      cases.push({ id: fixture.id, aliases: fixture.cases.map((c) => c.id), ...result });
  }
  const counts: Partial<Record<SemanticKind, number>> = {};
  for (const entry of cases)
    for (const category of entry.categories) counts[category] = (counts[category] ?? 0) + 1;
  if (json)
    console.log(
      JSON.stringify(
        { total: selected.length, different: cases.length, counts, mathErrors, unsupported, cases },
        null,
        2,
      ),
    );
  else {
    console.log(
      `Compared ${selected.length} inputs; ${cases.length} have${kind ? ` ${kind}` : ""} candidate GitHub issues.`,
    );
    console.log(
      `Categories: ${Object.entries(counts)
        .map(([k, v]) => `${k}=${v}`)
        .join(", ")}`,
    );
    if (mathErrors.length)
      console.log(
        `VS Code math errors in ${mathErrors.length} inputs (those formulas are not evaluated): ${mathErrors
          .slice(0, 5)
          .map((e) => e.id)
          .join(", ")}`,
      );
    if (only && mathErrors.length) console.log(JSON.stringify(mathErrors[0], null, 2));
    if (unsupported.length)
      console.log(
        `VS Code may not render GitHub features in ${unsupported.length} inputs: ${unsupported
          .slice(0, 5)
          .map((e) => e.id)
          .join(", ")}`,
      );
    if (only && unsupported.length) console.log(JSON.stringify(unsupported[0], null, 2));
    if (all || only) {
      for (const c of cases) {
        console.log(
          `${c.id}\t${c.categories.join(",")}\tGitHub: quirks=${c.quirksVsGitHub.join(",") || "match"} vscode=${c.vscodeVsGitHub.join(",") || "match"}${c.vscodeUnsupported.length ? `\tVS Code incomplete: ${c.vscodeUnsupported.join(",")}` : ""}`,
        );
        if (only) console.log(JSON.stringify(c.differences, null, 2));
      }
    } else {
      const groups = new Map<string, string[]>();
      for (const c of cases) {
        const key = `${c.categories.join(",")} | GitHub: quirks=${c.quirksVsGitHub.join(",") || "match"} vscode=${c.vscodeVsGitHub.join(",") || "match"}${c.vscodeUnsupported.length ? ` | VS Code incomplete: ${c.vscodeUnsupported.join(",")}` : ""}`;
        groups.set(key, [...(groups.get(key) ?? []), c.id]);
      }
      for (const [key, ids] of [...groups].sort((a, b) => b[1].length - a[1].length))
        console.log(`${ids.length} cases: ${key} | examples: ${ids.slice(0, 3).join(", ")}`);
      console.log("Use --all for case IDs, --kind CATEGORY to filter, or --only ID for details.");
    }
  }
} else {
  const path = paths[0];
  const source = path && path !== "-" ? await Bun.file(path).text() : await Bun.stdin.text();
  const expected = live
    ? await fetchGitHubHtml(source, { token: process.env.GITHUB_PAT ?? process.env.GITHUB_TOKEN })
    : undefined;
  const result = await compare(source, expected);
  if (kind) result.categories = result.categories.filter((category) => category === kind);
  if (json)
    console.log(
      JSON.stringify({ githubSource: live ? "github-api" : "local-quirks", ...result }, null, 2),
    );
  else if (!result.categories.length) {
    console.log(live ? "GitHub source: live API." : "GitHub source: local approximation.");
    console.log("No candidate GitHub breakage in the selected categories.");
    if (result.vscodeMathErrors.length)
      console.log(`VS Code math errors (not evaluated): ${result.vscodeMathErrors.join(", ")}`);
    if (result.vscodeUnsupported.length)
      console.log(`VS Code may not render GitHub features: ${result.vscodeUnsupported.join(", ")}`);
  } else {
    console.log(live ? "GitHub source: live API." : "GitHub source: local approximation.");
    if (result.vscodeMathErrors.length)
      console.log(`VS Code math errors (not evaluated): ${result.vscodeMathErrors.join(", ")}`);
    if (result.vscodeUnsupported.length)
      console.log(`VS Code may not render GitHub features: ${result.vscodeUnsupported.join(", ")}`);
    for (const category of result.categories) {
      console.log(`\n${category}:`);
      console.log(JSON.stringify(result.differences[category], null, 2));
    }
  }
}
