# GitHub and VS Code Markdown comparison

This project finds Markdown that may render correctly in VS Code's built-in preview but differently on GitHub. It uses `markdown-it` and `@vscode/markdown-it-katex` to approximate VS Code's default preview, then compares its meaningful output with GitHub HTML.

```sh
bun install --frozen-lockfile
bun run compare input.md                    # Offline: approximate GitHub with the local renderer
bun run compare --github-live input.md      # One real GitHub API request; sends the file contents
bun run compare --fixtures                  # Use saved GitHub API responses, offline
bun run compare --fixtures --kind math      # Filter by category
bun run compare --fixtures --all            # List all candidates
bun run compare --only original-0024        # Inspect a saved case
bun run measure:recall                     # Evaluate the small labeled benchmark
```

The live option uses `GITHUB_PAT` or `GITHUB_TOKEN` if available; GitHub's rate limits apply. The offline local renderer is an approximation, and can miss GitHub behavior; fixture comparisons use the exact saved GitHub response, even where the local renderer disagrees. The live option sends the document to GitHub, while ordinary comparisons and tests make no requests. `--json` emits machine-readable results.

The comparison projects HTML onto text, math, code, links, images, task states, formatting, and list/table structure. It extracts TeX from KaTeX's MathML and GitHub's math placeholders, and ignores highlighter markup and nonessential serialization. It looks for content or formatting present in the VS Code output that may be missing or different in GitHub. KaTeX errors are reported separately rather than called GitHub defects. Findings are candidates, not a guarantee of correct visual rendering, security, or author intent; the projection is lossy, and it cannot detect shared renderer errors.

Individual results include source-location candidates. Exact ranges identify matching Markdown expressions, destinations, or words; they do not prove that the span caused the difference. Structural changes and fenced blocks may point to an enclosing block instead. Repeated content is reported as ambiguous, and differences that cannot be aligned are marked unavailable rather than assigned a guessed location. Large candidate lists may be capped. Line and column numbers are one-based; columns count Unicode code points, end positions are exclusive, and JSON offsets use UTF-16 code units. Grouped fixture summaries omit individual ranges; use `--all`, `--only ID`, or `--json` to inspect them.

The semantic summary keeps an ordered `nodes` array with separate text, math, code, image, and math-error nodes. Text comparisons follow text nodes in order without substituting markers for other content. `compareMarkdown` adds a `source` field to VS Code nodes when the source map can associate them with a Markdown block; repeated matches are marked ambiguous, and unmapped nodes are unavailable. These are block-level candidates, while individual findings can have more precise ranges. `text` now contains individual text fragments instead of a single document-wide string.

Location analysis uses `markdown-it-source-map` to associate rendered top-level elements with source lines. The plugin does not provide inline columns, end lines, or mappings for every renderer such as fenced blocks, so token ranges and conservative source matching are still used for those cases. Source-provided attributes are not trusted as plugin-generated mappings, and source-map instrumentation is confined to analysis rather than normal preview output.

For programmatic comparisons of already-rendered fragments:

```ts
import { compareMarkdown } from "gfm-bun-render/compare";
const { findings } = compareMarkdown(markdown, githubHtml, vscodeHtml);
```

The VS Code-style renderer follows default preview settings (`html: true`, `linkify: true`, `breaks: false`, fuzzy links off) and the built-in math plugin's fenced math and macro settings. VS Code itself has additional preview plugins, webview security settings, styles, and configurable behavior, which this approximation does not fully reproduce. In particular, do not assume that all GitHub features, such as footnotes, task lists, alerts, and Mermaid widgets, are reproduced by these libraries. The returned VS Code-style HTML is **not sanitized** and must not be inserted into a trusted page without an appropriate sanitizer. [VS Code Markdown engine](https://github.com/microsoft/vscode/blob/main/extensions/markdown-language-features/src/markdownEngine.ts), [VS Code math extension](https://github.com/microsoft/vscode/blob/main/extensions/markdown-math/src/extension.ts).

The default recall benchmark (`test/benchmarks/vscode-github-risk.json`) contains seven positive examples and four controls, selected manually from existing development fixtures. The cases have assumed intent, some share root causes, and the sample is neither independent nor representative. A document is counted as detected when any candidate category appears; that candidate may not identify the annotated issue. `--style` and `--rendering` select older, broader label sets that include cases VS Code itself does not necessarily render as intended. Scores do not estimate general recall.

## Setup and use

Bun 1.4.4 is required. On first import of the GitHub-compatible renderer, `bun:ffi` compiles the unmodified cmark-gfm sources in-process via `cc()`. A small C entry point uses preprocessor includes and isolates colliding internal names; it is not a pre-generated copy of the source files. No separate compiler, CMake, make, or setup step is required. Standard C headers must be available to Bun’s C compiler. The VS Code-style renderer does not compile or call cmark-gfm.

```sh
bun install --frozen-lockfile
bun test
bun run render input.md > output.html
```

The CLI uses the GitHub-compatible renderer and also accepts standard input. Both renderers return a Promise containing an HTML fragment and limit input to 1 MB and output to 32 MB. The compatible renderer's native compilation at import and native parsing are synchronous in the calling process; the Promise also covers asynchronous HTML transformations. The `cc()` integration has been tested here with Bun 1.4.4 on Linux ARM64; other platforms have not been verified in this project.

## Compatibility status

The offline E2E tests cover 617 distinct inputs, representing 665 named cases, with responses individually captured from GitHub's Markdown API in GFM mode. The cases cover positive and negative syntax, math in different contexts, task-list and document boundaries, HTML sanitization, code highlighting, images, widgets, and larger documents. Each input and its exact response are stored as an `.md`/`.html` pair; the manifest records aliases, source, rule IDs, hashes, and capture time. Run `bun run test:e2e` to run only these tests; no API request is made by the tests.

The comparison parses HTML into a tree, sorts attributes, normalizes generated footnote and math IDs, and collapses ordinary HTML whitespace. It ignores formatting whitespace at block boundaries but keeps meaningful inline spacing and preserves whitespace inside code, preformatted text, and math. It checks tags, nesting, attributes, and content; it does not establish visual equivalence or validate the browser's rendering of math or diagrams. For captured GitHub responses, proxied image URLs and proxy-only wrappers are normalized back to the original image URLs before comparison.

Neither renderer performs syntax highlighting. E2E comparison ignores GitHub’s highlighting wrappers, token spans, code-block language attributes, and one trailing code-block newline, while still checking code contents; language preservation is covered separately by unit tests.

To capture missing cases from `scripts/new-e2e-cases.ts`, set `GITHUB_PAT` or `GITHUB_TOKEN` and run `bun scripts/capture-e2e.ts`. The script reuses existing references, waits at least 650 ms between requests by default, saves each successful response before continuing, and stops on an HTTP error so it can be resumed. `GFM_CAPTURE_DELAY_MS` may increase the interval (minimum 250 ms). GitHub's authenticated rate limit is typically 5,000 requests per hour, and secondary limits can apply.

For example, `matrix-bare-prefix-0` captures recognized `$x$`, while `matrix-bare-prefix-1` captures `a$x$` remaining literal. Both expected responses come from GitHub, rather than being inferred from the renderer.

The original inputs include cases split from the Markdown batches previously provided in this workspace, including examples attributed there to the markdown-math-acid-test project. The independently captured outputs reflect each case on its own; document-wide interactions can produce different results.

For reproducible output, pass a 32-character lowercase hexadecimal `runId` to `render`, or use `--run-id ID` on the CLI. Normally the renderer generates a fresh ID for each fragment, as GitHub does; reusing an ID in the same page can cause anchor collisions.

The renderer leaves remote image URLs unproxied. Matching these fixtures does not prove compatibility with every possible Markdown input; math, sanitization, and GitHub's other server-side transformations have broader behavior than the cases covered here. Raw HTML passes through an HTML5 parser and DOMPurify before renderer transformations; this prototype is not a security-audited service.

The native compatibility layer owns the cmark parser and AST, adjusts task items and footnote references, and implements math boundary recognition and the post-sanitization handling of alerts, lists, tables, code blocks, footnotes, and images in C. TypeScript still coordinates HTML5 parsing and sanitization, constructs math DOM nodes, performs JavaScript-compatible Unicode lowercase conversion. The post-sanitization C tree consumes HTML serialized by the DOM rather than parsing arbitrary input directly.

The native source is based on cmark-gfm 0.29.0.gfm.13. The vendored upstream C files are unmodified and are checked against the release hashes in `test/fixtures/cmark-gfm-upstream.json`. Compatibility adjustments live in `src/quirks/cmark-*.c` and the TypeScript rendering pipeline; cmark sources remain unmodified. Multiline footnote compatibility can require reparsing transformed source, so AST XML source positions in those cases refer to the transformed input. See `vendor/cmark-gfm/PROVENANCE.json` and `COPYING` for provenance and licensing. Original prototype code is MIT licensed; third-party code retains its own license.
