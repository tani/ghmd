# gfm-bun-render

An in-process Bun renderer for GitHub-style Markdown HTML fragments. Its current goal is to reproduce the saved GitHub API results in `../github-html/` for the inputs in `../github-markdown/`.

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
# Resolve signed image URLs through GitHub (requires network access):
bun run render --github-image-proxy input.md > output.html
```

The CLI also accepts standard input. The library returns a Promise containing an HTML fragment. The native compilation at import and native parsing are synchronous in the calling process; the Promise also covers the asynchronous HTML transformations. It uses the vendored cmark-gfm parser and happy-dom, and limits input to 1 MB and output to 32 MB. The `cc()` integration has been tested here with Bun 1.4.4 on Linux ARM64; other platforms have not been verified in this project.

## Compatibility status

The reference tests compare all four saved fragments byte-for-byte by supplying their saved per-render footnote and math IDs via `runId`. GitHub signs image proxy URLs with its own server-side key. For the image in batch 1, the test supplies the already-saved signed URL via `imageProxy`; without a resolver, the renderer leaves the original image URL in place and batch 1 does not match exactly.

The CLI flag sends each unique remote image URL to GitHub's Markdown API to obtain its signed Camo URL. It does not send the rest of the Markdown. GitHub API rate limits apply; optionally set `GITHUB_PAT` or `GITHUB_TOKEN` for authenticated requests. To use a workspace-root `.env` from this directory, run `bun --env-file=../.env src/cli.ts --github-image-proxy input.md`. Failed lookups cause rendering to fail. Without the flag, rendering remains offline and leaves remote image URLs unproxied.

The library exposes the same opt-in resolver:

```ts
import { render, createGitHubImageProxy } from "./src/index.ts";
const imageProxy = createGitHubImageProxy(); // optionally pass {token}
const html = await render(markdown, { imageProxy });
```

Applications with a different trusted image proxy can supply a synchronous or asynchronous resolver:

```ts
const html = await render(markdown, {
  imageProxy: async (canonicalUrl) => getSignedProxyUrl(canonicalUrl),
});
```

For reproducible output, pass a 32-character lowercase hexadecimal `runId` to `render`, or use `--run-id ID` on the CLI. Normally the renderer generates a fresh ID for each fragment, as GitHub does; reusing an ID in the same page can cause anchor collisions.

The resolver must return an HTTP(S) URL. GitHub's signing key is not part of this project; the opt-in resolver obtains signatures from GitHub rather than generating them locally. Matching the four fixtures does not prove compatibility with every possible Markdown input; math, sanitization, and GitHub's other server-side transformations have broader behavior than the cases covered here. Raw HTML passes through an HTML5 parser and DOMPurify before renderer transformations; this prototype is not a security-audited service.

The native source is based on cmark-gfm 0.29.0.gfm.13. The vendored upstream C files are unmodified and are checked against the release hashes in `test/fixtures/cmark-gfm-upstream.json`. Compatibility adjustments live in `src/native.ts`. Multiline footnote compatibility can require reparsing transformed source, so AST XML source positions in those cases refer to the transformed input. See `vendor/cmark-gfm/PROVENANCE.json` and `COPYING` for provenance and licensing. Original prototype code is MIT licensed; third-party code retains its own license.
