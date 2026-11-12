# GitHub Markdown renderer in Go

This project renders GitHub-flavored Markdown to GitHub-compatible HTML and compares that output with a second GitHub-dialect parser. It is a Go rewrite of the former Bun/TypeScript implementation.

- **cmark-gfm** (vendored) is transpiled by `modernc.org/ccgo/v4` and supplies the GitHub AST without cgo.
- **golang.org/x/net/html** owns HTML5 parsing, tree repair, sanitization and serialization.
- **md4c** (vendored) is similarly transpiled by ccgo and executed by `modernc.org/libc`, with LaTeX math spans enabled.

## Build and run

Go 1.26 or later is required. There is **no C compiler or cgo requirement** at build or test time: the cmark-gfm and md4c sources are transpiled by `modernc.org/ccgo/v4` and executed by `modernc.org/libc`. Use the same guard in CI:

```sh
go vet -unsafeptr=false -unreachable=false ./...
CGO_ENABLED=0 go test -count=1 ./...
```

`unsafeptr` and `unreachable` are intentionally disabled for the generated sources: they coerce C heap addresses, not movable Go allocations, and ccgo's platform-cleanups contain unreachable branches.

The checked-in generated Go files target `linux/arm64`; run `scripts/gen-native.sh` on the desired platform to regenerate them.

```sh
go build ./cmd/ghmd
./ghmd render input.md > output.html
./ghmd render --run-id 0123456789abcdef0123456789abcdef input.md
./ghmd diff input.md        # cmark-gfm vs md4c structural differences
./ghmd lint input.md        # suspicious raw HTML and standard-renderer regressions
```

`render` accepts standard input when the path is `-` or omitted. Input is limited to 1 MB and output to 32 MB.

The public Go API is:

```go
import "ghmd"

html, err := ghmd.Render(source, ghmd.RenderOptions{RunID: ""})
```

An empty run ID generates a random 32-character lowercase hexadecimal identifier. Reusing an ID in the same page can cause anchor collisions.

## Tests

```sh
go test ./...
```

The GitHub reference suite renders 617 captured inputs and compares canonical HTML trees against responses captured from GitHub's Markdown API in GFM mode. Fixtures live in `test/fixtures/e2e/`; source and response hashes are checked before rendering. The comparison ignores syntax-highlighting wrappers, normalizes generated footnote and math IDs, and collapses ordinary HTML whitespace; it still checks tags, nesting, attributes and content.

`test/fixtures/lint/styleguide-bad.md` is a long document that deliberately violates every rule in `GFM_STYLEGUIDE.md`; `styleguide-bad.expected.json` records the violations the linter is expected to catch. `TestStyleGuideBadMarkdownRecall` lints each rule section independently (avoiding cross-section cancellation and ambiguity) and measures **100% recall for all 53 automatically detectable guide rules** (76/76 expected findings) using the generated 344-line violation document. The other 11 rules (line-break intent, interactive-checkbox intent, compatibility-policy intent, and seven implementer requirements) are registered as Manual review guidance rather than incorrectly auto-detected.

`test/fixtures/lint/cases.json` is a labeled corpus of suspicious raw HTML and renderer-regression cases. `TestLintCorpusRecall` requires every expected rule at the expected line (and column where repeated findings on one line need disambiguation), and separately reports additional related authoring-policy findings produced by the generalized rules. `TestFixtureDifferentialNoiseBudget` guards the 617 GitHub fixtures against differential-noise regressions.

`test/fixtures/native-pipeline/` contains exact-byte regression cases for the postprocessing pipeline. The vendored cmark-gfm sources and generated headers are verified against `test/fixtures/cmark-gfm-upstream.json`.

`ghmd diff` is intentionally lossy: it projects both rendered fragments onto a normalized tree and reports structural differences. It does not claim that either parser is correct.

`ghmd lint` reports source locations that are likely to break on GitHub:

- Raw HTML that GitHub removes, unwraps or relocates: executable/filtered containers (`script`, `style`, `iframe`, `svg`, ...), event and namespace attributes, unsafe URLs, source-authored `input`, `template`, nested anchors, block elements inside `p`, unclosed and stray tags.
- Math that GitHub leaves literal or rejects: boundary violations (`a$x$`, `$ x$`, `$x $`, `&` in the expression), unclosed dollars (currency like `$5` is ignored), math inside footnote definitions, and non-isolated `$$...$$` display delimiters. When a boundary rule and `standard/lost-math` both fire on one line, the renderer comparison is kept and the boundary finding is dropped; boundary findings remain when md4c also rejects the syntax.
- Regressions relative to the md4c "standard" renderer: visible text, code, links, images, math, task state, inline formatting (`strong`, `em`, `del`) and table shape produced by the standard parser but missing or changed in the GitHub output. md4c runs with LaTeX math spans enabled, so `standard/lost-math` reports contexts where the standard parse renders math but GitHub leaves it literal (for example emphasis, link labels and raw inline tags). Known dialect noise is suppressed: math delimiters, footnote fallback links, alert markers, autolinks, table-pipe code, currency-like `$...$`, and presentation-only attributes. Differential checks are skipped for documents using footnote syntax or SVG, where md4c's parser semantics are not a usable baseline; strikethrough is compared through `del` formatting while its delimiter-level text noise is suppressed. Raw-HTML rules still apply in all cases.

Findings are printed as `path:line:column: severity rule: message` with one-based line and Unicode code-point columns. The command exits 1 when any finding is reported and prints nothing on clean input. Unlocated differential findings use line 0, column 0.

## Authoring-style coverage

`StyleRules` composes source-observable authoring checks by `GFM_STYLEGUIDE.md` ID. It covers blocks and list continuation, emphasis and CJK punctuation, code fences/inline code, links and heading aliases, meaningful image alt text, table cells and pipe escaping, tildes, task markers, footnote definitions and labels, alerts, details, math delimiters and nesting context, widget data, Mermaid authoring policy, mentions/host extensions, and compatibility-profile advisories. Source positions use one-based lines and Unicode code-point columns; examples in inline code, fences, and HTML comments are ignored.

The 53 automatically checkable guide rules are those with a reliable source signal. `BLOCK-003`, `TASK-002`, `COMPAT-002`, `COMPAT-003`, and `RENDER-001` through `RENDER-007` have honest manual-review guidance instead of false-positive proxy checks. Compatibility-profile selection (`COMPAT-001`) is detected when rich features have no declaration; an explicit declaration avoids the advisory.

## Architecture

| Package | Responsibility |
| --- | --- |
| `internal/cmark-gfm` | Go AST wrapper around ccgo-transpiled cmark-gfm and `modernc.org/libc`. |
| `internal/standard` | Go wrapper over ccgo-transpiled md4c and `modernc.org/libc` (CommonMark plus GitHub extensions and LaTeX math spans). |
| `internal/github` | Full GitHub compatibility pipeline: AST adaptation, sanitization, DOM repair, math, tables, alerts, footnotes, task lists, images, links and final serialization. |
| `internal/analysis` | Shared HTML fragment parsing, structural tree comparison, and semantic feature extraction. |
| `internal/diff` | Thin wrapper over `analysis.CompareFragments` for cmark-gfm vs md4c structural differences. |
| `internal/lint` | Raw-HTML source/policy rules plus the standard-renderer differential. Uses `analysis.SemanticFeatures`; does not call `internal/diff`. |
| `internal/verify` | Fixture canonicalization and GitHub image-proxy normalization. |
| `cmd/ghmd` | `render`, `diff` and `lint` command-line interface. |

The pipeline keeps the former renderer's byte-range tree for already-sanitized HTML. HTML5 boundaries use `x/net/html`: raw-input filtering, sanitizer traversal, DOM compatibility repair and DOM serialization. The renderer makes no JavaScript or Node calls.

## Rendering policy

- Raw HTML passes through an allowlist sanitizer. Executable containers, event handlers, inline styles, unsafe URLs, namespace attributes and dangerous `id`/`name` values are removed. Template contents are checked recursively; safe HTML and basic MathML are retained. Unknown tags are unwrapped.
- The allowlist data is adapted from DOMPurify 3.4.1 under Apache-2.0; see `NOTICE` and `POLICY-LICENSE` files retained with the pipeline sources.
- Neither renderer performs syntax highlighting.
- Input limits: 1 MB of Markdown, 32 MB of HTML output, tree depth 512.
- Remote image URLs are left unproxied.

## Vendored sources and licensing

`third_party/cmark-gfm` is based on cmark-gfm 0.29.0.gfm.13 and is unmodified; see `PROVENANCE.json` and `COPYING`. `third_party/md4c` is an unmodified copy of md4c 0.5.2; see `PROVENANCE.json` and `LICENSE.md`. `third_party/md4c-bridge` contains only our output-callback bridge compiled against it. `third_party/md4c-bridge/bridge.c` and `third_party/cmark-gfm-bridge/sources.c` are callback/enumeration glue compiled by ccgo to preserve the public Go render/parse API without a Go-to-C callback. `scripts/gen-native.sh` regenerates `internal/standard/md4c_generated.go` and `internal/cmark-gfm/cmark_generated.go` from those sources. Original project code is MIT licensed; third-party code keeps its own license. ccgo is BSD-3-Clause, and `modernc.org/libc` retains its own license.
