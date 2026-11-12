# GitHub Markdown renderer in Go

This project renders GitHub-flavored Markdown to GitHub-compatible HTML and compares that output with a second GitHub-dialect parser. It is a Go rewrite of the former Bun/TypeScript implementation.

- **cmark-gfm** (vendored) parses Markdown through cgo and supplies the GitHub AST.
- **golang.org/x/net/html** owns HTML5 parsing, tree repair, sanitization and serialization.
- **md4c** (vendored) parses the same Markdown through cgo as an independent GitHub dialect for differential checks, with LaTeX math spans enabled.

## Build and run

Go 1.26 or later, a C compiler and cgo are required.

```sh
go build ./cmd/ghmd
./ghmd render input.md > output.html
./ghmd render --run-id 0123456789abcdef0123456789abcdef input.md
./ghmd diff input.md        # cmark-gfm vs md4c structural differences
./ghmd lint input.md        # suspicious raw HTML and renderer regressions
```

`render` accepts standard input when the path is `-` or omitted. Input is limited to 1 MB and output to 32 MB. Reusing a run ID in the same page can cause anchor collisions; omitting it generates a random 32-character lowercase hexadecimal identifier.

The public Go API is:

```go
import "ghmd"

html, err := ghmd.Render(source, ghmd.RenderOptions{RunID: ""})
```

### Make targets

`make` builds with the host Go toolchain and with a TinyGo installation on `PATH`. Override either compiler with `GO=` or `TINYGO=`.

| Target | Output |
| --- | --- |
| `make go-native` | `build/go/<arch>/ghmd` using cgo |
| `make tinygo-native` | `build/tinygo/<arch>/ghmd` using TinyGo + cgo |
| `make tinygo-wasi` | `build/tinygo/wasi/ghmd.wasm` |
| `make all` | All three builds |
| `make clean` | Removes the generated binaries |

`<arch>` defaults to `go env GOHOSTARCH`. TinyGo 0.44.0 is the tested toolchain. It is used as a system executable rather than a Go module tool because building TinyGo from source requires matching LLVM and libclang development files. The Makefile also supplies absolute include paths for the vendored cmark-gfm headers, which TinyGo's cgo handling does not infer.

The WASI build includes `third_party/cmark-gfm-bridge/wasi_shim.c`, which defines the ASCII `tolower` function missing from the available `wasi-libc` link set. It is compiled only under `__wasip1__` or `__wasi__`, so native builds are unaffected.

## Lint behavior

`ghmd diff` is intentionally lossy: it projects both rendered fragments onto a normalized tree and reports structural differences. It does not claim that either parser is correct.

`ghmd lint` reports source locations that are likely to break on GitHub:

- Raw HTML that GitHub removes, unwraps or relocates: executable/filtered containers (`script`, `style`, `iframe`, `svg`, ...), event and namespace attributes, unsafe URLs, source-authored `input`, `template`, nested anchors, block elements inside `p`, unclosed and stray tags.
- Math that GitHub leaves literal or rejects: boundary violations (`a$x$`, `$ x$`, `$x $`, `&` in the expression), unclosed dollars (currency like `$5` is ignored), math inside footnote definitions, and non-isolated `$$...$$` display delimiters. When a boundary rule and `standard/lost-math` both fire on one line, the renderer comparison is kept and the boundary finding is dropped; boundary findings remain when md4c also rejects the syntax.
- Regressions relative to the md4c standard renderer: visible text, code, links, images, math, task state, inline formatting (`strong`, `em`, `del`) and table shape produced by the standard parser but missing or changed in the GitHub output. md4c runs with LaTeX math spans enabled, so `standard/lost-math` reports contexts where the standard parse renders math but GitHub leaves it literal (for example emphasis, link labels and raw inline tags). Known dialect noise is suppressed: math delimiters, footnote fallback links, alert markers, autolinks, table-pipe code, currency-like `$...$`, and presentation-only attributes. Differential checks are skipped for documents using footnote syntax or SVG, where md4c's parser semantics are not a usable baseline; strikethrough is compared through `del` formatting while its delimiter-level text noise is suppressed. Raw-HTML rules still apply in all cases.

Findings are printed as `path:line:column: severity rule: message` with one-based line and Unicode code-point columns. Unlocated differential findings use line 0, column 0. The command exits 1 when any finding is reported and prints nothing on clean input.

### Authoring-style checks

`StyleRules` composes source-observable checks by `GFM_STYLEGUIDE.md` ID. It covers blocks and list continuation, emphasis and CJK punctuation, code fences and inline code, links and heading aliases, meaningful image alt text, table cells and pipe escaping, tildes, task markers, footnote definitions and labels, alerts, details, math delimiters and nesting context, widget data, Mermaid policy, mentions and host extensions, and compatibility-profile advisories. Source positions use one-based lines and Unicode code-point columns; examples in inline code, fences and HTML comments are ignored.

Rules without a reliable source signal are registered as Manual review guidance instead of false-positive proxy checks: `BLOCK-003`, `TASK-002`, `COMPAT-002`, `COMPAT-003`, and `RENDER-001` through `RENDER-007`. Compatibility-profile selection (`COMPAT-001`) is detected when rich features have no declaration; an explicit declaration avoids the advisory.

## Tests

```sh
go test ./...
```

The GitHub reference suite renders 617 captured inputs and compares canonical HTML trees against responses captured from GitHub's Markdown API in GFM mode. Fixtures live in `test/fixtures/e2e/`; source and response hashes are checked before rendering. The comparison ignores syntax-highlighting wrappers, normalizes generated footnote and math IDs, and collapses ordinary HTML whitespace, but still checks tags, nesting, attributes and content.

`test/fixtures/lint/styleguide-bad.md` deliberately violates every rule in `GFM_STYLEGUIDE.md`; `styleguide-bad.expected.json` records the expected violations. `TestStyleGuideBadMarkdownRecall` lints each rule section independently to avoid cross-section cancellation and measures 100% recall for all 53 automatically detectable guide rules (76/76 findings) in the 344-line document. The other 11 rules are Manual review guidance.

`test/fixtures/lint/cases.json` is a labeled corpus of suspicious raw HTML and renderer-regression cases. `TestLintCorpusRecall` requires every expected rule at the expected line and, where repeated findings need disambiguation, column. `TestFixtureDifferentialNoiseBudget` guards the 617 GitHub fixtures against differential-noise regressions.

`test/fixtures/native-pipeline/` contains exact-byte regression cases for the postprocessing pipeline. Vendored cmark-gfm sources and generated headers are verified against `test/fixtures/cmark-gfm-upstream.json`.

## Architecture

| Package | Responsibility |
| --- | --- |
| `internal/cmark-gfm` | cgo binding to the unmodified vendored cmark-gfm parser and AST. |
| `internal/standard` | cgo binding to the vendored md4c parser (CommonMark, GitHub extensions and LaTeX math spans). |
| `internal/github` | Full GitHub compatibility pipeline: AST adaptation, sanitization, DOM repair, math, tables, alerts, footnotes, task lists, images, links and final serialization. |
| `internal/analysis` | Shared HTML fragment parsing, structural tree comparison and semantic feature extraction. |
| `internal/diff` | Thin wrapper over `analysis.CompareFragments` for cmark-gfm vs md4c structural differences. |
| `internal/lint` | Raw-HTML source and policy rules plus the standard-renderer differential. Uses `analysis.SemanticFeatures`, not `internal/diff`. |
| `internal/verify` | Fixture canonicalization and GitHub image-proxy normalization. |
| `cmd/ghmd` | `render`, `diff` and `lint` command-line interface. |

The pipeline keeps the former renderer's byte-range tree for already-sanitized HTML. HTML5 boundaries use `x/net/html`: raw-input filtering, sanitizer traversal, DOM compatibility repair and DOM serialization. The renderer makes no JavaScript or Node calls.

## Rendering policy

- Raw HTML passes through an allowlist sanitizer. Executable containers, event handlers, inline styles, unsafe URLs, namespace attributes and dangerous `id`/`name` values are removed. Template contents are checked recursively; safe HTML and basic MathML are retained. Unknown tags are unwrapped.
- The allowlist data is adapted from DOMPurify 3.4.1 under Apache-2.0; see `NOTICE` and `POLICY-LICENSE` with the pipeline sources.
- Neither renderer performs syntax highlighting.
- Input limits: 1 MB of Markdown, 32 MB of HTML output, tree depth 512.
- Remote image URLs are left unproxied.

## Vendored sources and licensing

`third_party/cmark-gfm` is based on cmark-gfm 0.29.0.gfm.13 and is unmodified; see `PROVENANCE.json` and `COPYING`. `third_party/md4c` is md4c 0.5.2; see `PROVENANCE.json` and `LICENSE.md`. Original project code is MIT licensed; third-party code keeps its own license.
