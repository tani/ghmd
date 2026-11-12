# GitHub Markdown renderer in Go

This project renders GitHub-flavored Markdown to GitHub-compatible HTML and compares that output with a second GitHub-dialect parser. It is a Go rewrite of the former Bun/TypeScript implementation.

- **cmark-gfm** (vendored) parses Markdown through cgo and supplies the GitHub AST.
- **golang.org/x/net/html** owns HTML5 parsing, tree repair, sanitization and serialization.
- **md4c** (vendored) parses the same Markdown through cgo as an independent GitHub dialect for differential checks.

## Build and run

Go 1.26 or later, a C compiler and cgo are required.

```sh
go build ./cmd/gfm
./gfm render input.md > output.html
./gfm render --run-id 0123456789abcdef0123456789abcdef input.md
./gfm diff input.md        # cmark-gfm vs md4c structural differences
```

`render` accepts standard input when the path is `-` or omitted. Input is limited to 1 MB and output to 32 MB.

The public Go API is:

```go
import "gfm"

html, err := gfm.Render(source, gfm.RenderOptions{RunID: ""})
```

An empty run ID generates a random 32-character lowercase hexadecimal identifier. Reusing an ID in the same page can cause anchor collisions.

## Tests

```sh
go test ./...
```

The GitHub reference suite renders 617 captured inputs and compares canonical HTML trees against responses captured from GitHub's Markdown API in GFM mode. Fixtures live in `test/fixtures/e2e/`; source and response hashes are checked before rendering. The comparison ignores syntax-highlighting wrappers, normalizes generated footnote and math IDs, and collapses ordinary HTML whitespace; it still checks tags, nesting, attributes and content.

`test/fixtures/native-pipeline/` contains exact-byte regression cases for the postprocessing pipeline. The vendored cmark-gfm sources and generated headers are verified against `test/fixtures/cmark-gfm-upstream.json`.

`gfm diff` is intentionally lossy: it projects both rendered fragments onto a normalized tree and reports structural differences. It does not claim that either parser is correct.

## Architecture

| Package | Responsibility |
| --- | --- |
| `internal/cmark-gfm` | cgo binding to the unmodified vendored cmark-gfm parser and AST. |
| `internal/standard` | cgo binding to the vendored md4c parser (CommonMark plus GitHub extensions). |
| `internal/github` | Full GitHub compatibility pipeline: AST adaptation, sanitization, DOM repair, math, tables, alerts, footnotes, task lists, images, links and final serialization. |
| `internal/diff` | cmark-gfm vs md4c normalized tree comparison. |
| `internal/verify` | Fixture canonicalization and GitHub image-proxy normalization. |
| `cmd/gfm` | `render` and `diff` command-line interface. |

The pipeline keeps the former renderer's byte-range tree for already-sanitized HTML. HTML5 boundaries use `x/net/html`: raw-input filtering, sanitizer traversal, DOM compatibility repair and DOM serialization. The renderer makes no JavaScript or Node calls.

## Rendering policy

- Raw HTML passes through an allowlist sanitizer. Executable containers, event handlers, inline styles, unsafe URLs, namespace attributes and dangerous `id`/`name` values are removed. Template contents are checked recursively; safe HTML and basic MathML are retained. Unknown tags are unwrapped.
- The allowlist data is adapted from DOMPurify 3.4.1 under Apache-2.0; see `NOTICE` and `POLICY-LICENSE` files retained with the pipeline sources.
- Neither renderer performs syntax highlighting.
- Input limits: 1 MB of Markdown, 32 MB of HTML output, tree depth 512.
- Remote image URLs are left unproxied.

## Vendored sources and licensing

`third_party/cmark-gfm` is based on cmark-gfm 0.29.0.gfm.13 and is unmodified; see `PROVENANCE.json` and `COPYING`. `third_party/md4c` is md4c 0.5.2; see `PROVENANCE.json` and `LICENSE.md`. Original project code is MIT licensed; third-party code keeps its own license.
