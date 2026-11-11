# gfm-bun-render

An in-process Bun renderer for GitHub-style Markdown HTML fragments. Its reference tests compare individual Markdown inputs with saved GitHub API responses in `test/fixtures/e2e/`.

```ts
import { render } from "./src/index.ts";
const html = await render("# Hello");
```

## Setup and use

Bun 1.4.4 is required. On first import, `bun:ffi` compiles the unmodified cmark-gfm sources in-process via `cc()`. A small C entry point uses preprocessor includes and isolates colliding internal names; it is not a pre-generated copy of the source files. No separate compiler, CMake, make, or setup step is required. Standard C headers must be available to Bun’s C compiler.

```sh
bun install --frozen-lockfile
bun test
bun run render input.md > output.html
```

The CLI also accepts standard input. The library returns a Promise containing an HTML fragment. The native compilation at import and native parsing are synchronous in the calling process; the Promise also covers the asynchronous HTML transformations. It uses the vendored cmark-gfm parser and happy-dom, and limits input to 1 MB and output to 32 MB. The `cc()` integration has been tested here with Bun 1.4.4 on Linux ARM64; other platforms have not been verified in this project.

## Compatibility status

The offline E2E tests cover 580 distinct inputs, representing 624 named cases, with responses individually captured from GitHub's Markdown API in GFM mode. The cases cover positive and negative syntax, math in different contexts, task-list and document boundaries, HTML sanitization, code highlighting, images, widgets, and larger documents. Each input and its exact response are stored as an `.md`/`.html` pair; the manifest records aliases, source, rule IDs, hashes, and capture time. Run `bun run test:e2e` to run only these tests; no API request is made by the tests.

The comparison parses HTML into a tree, sorts attributes, normalizes generated footnote and math IDs, and collapses ordinary HTML whitespace. It ignores formatting whitespace at block boundaries but keeps meaningful inline spacing and preserves whitespace inside code, preformatted text, and math. It checks tags, nesting, attributes, and content; it does not establish visual equivalence or validate the browser's rendering of math or diagrams. Saved signed image URLs are supplied via `imageProxy`.

Known differences are listed with their reasons in `test/fixtures/e2e/known-differences.json`. Those tests still verify the GitHub fixture and the current output tree, so an unrelated regression or an improvement requires review rather than silently passing. The current list includes syntax highlighting for several languages, raw checkboxes in lists, SVG text, and deep list nesting.

To capture missing cases from `scripts/new-e2e-cases.ts`, set `GITHUB_PAT` or `GITHUB_TOKEN` and run `bun scripts/capture-e2e.ts`. The script reuses existing references, waits at least 650 ms between requests by default, saves each successful response before continuing, and stops on an HTTP error so it can be resumed. `GFM_CAPTURE_DELAY_MS` may increase the interval (minimum 250 ms). GitHub's authenticated rate limit is typically 5,000 requests per hour, and secondary limits can apply.

For example, `matrix-bare-prefix-0` captures recognized `$x$`, while `matrix-bare-prefix-1` captures `a$x$` remaining literal. Both expected responses come from GitHub, rather than being inferred from the renderer.

The original inputs include cases split from the Markdown batches previously provided in this workspace, including examples attributed there to the markdown-math-acid-test project. The independently captured outputs reflect each case on its own; document-wide interactions can produce different results.

Applications can supply a synchronous or asynchronous image proxy resolver:

```ts
const html = await render(markdown, {
  imageProxy: async (canonicalUrl) => getSignedProxyUrl(canonicalUrl),
});
```

For reproducible output, pass a 32-character lowercase hexadecimal `runId` to `render`, or use `--run-id ID` on the CLI. Normally the renderer generates a fresh ID for each fragment, as GitHub does; reusing an ID in the same page can cause anchor collisions.

The resolver must return an HTTP(S) URL. Without one, the renderer leaves remote image URLs unproxied. Matching these fixtures does not prove compatibility with every possible Markdown input; math, sanitization, and GitHub's other server-side transformations have broader behavior than the cases covered here. Raw HTML passes through an HTML5 parser and DOMPurify before renderer transformations; this prototype is not a security-audited service.

The native source is based on cmark-gfm 0.29.0.gfm.13. The vendored upstream C files are unmodified and are checked against the release hashes in `test/fixtures/cmark-gfm-upstream.json`. Compatibility adjustments live in `src/native.ts`. Multiline footnote compatibility can require reparsing transformed source, so AST XML source positions in those cases refer to the transformed input. See `vendor/cmark-gfm/PROVENANCE.json` and `COPYING` for provenance and licensing. Original prototype code is MIT licensed; third-party code retains its own license.
