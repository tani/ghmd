# Markdown rendering comparison

This module finds content or formatting that the VS Code-style renderer preserves
but GitHub may render differently. It accepts HTML from the [local GitHub-style
renderer](../quirks/README.md), saved API responses, or an explicitly requested
live API call. Rendering and fetching are separate from comparison.

## Usage

```ts
import { render as renderGitHub } from "gfm-bun-render";
import { render as renderVSCode } from "gfm-bun-render/vscode";
import { compareMarkdown } from "gfm-bun-render/compare";

const source = "# Example\n\n$x^2$\n";
const githubHtml = await renderGitHub(source);
const vscodeHtml = await renderVSCode(source);
const result = compareMarkdown(source, githubHtml, vscodeHtml);
console.log(result.categories, result.findings);
```

The example runs offline. The comparison returns candidates for investigation;
it does not establish visual equivalence, malicious intent, or which renderer is
correct.

## API

| Function | Result |
| --- | --- |
| `summarizeHtml(html)` | A semantic summary and ordered text, math, code, image, and math-error nodes. |
| `compareHtml(left, right)` | Categories whose summaries differ in either direction. |
| `detectPotentialBreakage(githubHtml, vscodeHtml)` | Directional categories, VS Code math errors, unsupported VS Code features, and both summaries. |
| `compareMarkdown(source, githubHtml, vscodeHtml)` | The directional analysis plus source-location candidates; also attaches source metadata to VS Code summary nodes. |
| `locateDifferences(source, analysis)` | Source-location candidates for an existing directional analysis. |
| `missingTextFragments(expected, actual)` | Expected text fragments that cannot be matched in order. |

The categories are `text`, `math`, `code`, `links`, `images`, `tasks`, `structure`,
and `formatting`. The HTML projection uses the [native HTML parser](../html/README.md),
collapses ordinary text whitespace, extracts TeX from KaTeX annotations and
GitHub math placeholders, and omits highlighter markup. It is deliberately lossy;
differences that both renderers share cannot be detected.

## Source locations

`source-map.ts` creates the default [VS Code-style engine](../vscode/README.md),
parses the Markdown, and renders a separate analysis fragment with
`markdown-it-source-map` and a random marker attribute. It validates those
markers against token indices, tag names, and source lines before using them.
Normal renderer output is unaffected, and source-provided attributes alone do
not establish a mapping.

`locations.ts` associates VS Code semantic nodes with enclosing Markdown blocks,
then combines those blocks, token ranges, and source-text matches to locate
differences. Fences use token ranges because their custom renderers do not carry
the plugin markers. Raw HTML tags and attributes use HTML source offsets from
the native parser. GitHub output carries no Markdown source map.

Findings have an `exact`, `block`, `ambiguous`, or `unavailable` status. An exact
range identifies matching source text, not necessarily the cause of the
difference. Repeated matches remain ambiguous. Lines and columns are one-based;
columns count Unicode code points, offsets count UTF-16 code units, and range
ends are exclusive. Source mapping uses the default preview settings, so callers
using a customized renderer must account for that difference.

## Files and checks

| File | Responsibility |
| --- | --- |
| [index.ts](index.ts) | HTML projection, directional detection, and the public comparison API. |
| [source-map.ts](source-map.ts) | Analysis-only rendering and verified block mappings. |
| [locations.ts](locations.ts) | Mapping nodes and findings back to Markdown. |
| [github.ts](github.ts) | Internal `fetchGitHubHtml` helper: one GFM API request with an optional token and a 30-second timeout. It is not re-exported by the comparison entry point. |

Run from the repository root:

```sh
bun test test/compare.test.ts test/locations.test.ts test/source-map.test.ts test/github-reference.test.ts
```

The tests use local or mocked data. The CLI in
[scripts/compare-renderers.ts](../../scripts/compare-renderers.ts) selects local,
fixture, or explicitly requested live GitHub output.
