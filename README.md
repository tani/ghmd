# ghmd

GitHub-flavored Markdown renderer in Go. It uses vendored cmark-gfm for parsing, `golang.org/x/net/html` for HTML5 handling, and md4c for parser differential checks.

## Build

Requires Go 1.26+, a C compiler, and cgo.

```sh
go build ./cmd/ghmd
./ghmd render input.md > output.html
./ghmd diff input.md
./ghmd lint input.md
```

`render` reads stdin when the path is `-` or omitted. Input is limited to 1 MB; output is limited to 32 MB.

```go
html, err := ghmd.Render(source, ghmd.RenderOptions{})
```

### Make targets

| Target | Output |
| --- | --- |
| `make go-native` | `build/go/<arch>/ghmd` |
| `make tinygo-native` | `build/tinygo/<arch>/ghmd` |
| `make tinygo-wasi` | `build/tinygo/wasi/ghmd.wasm` |
| `make web` | `web/dist/` |
| `make all` | All non-web builds |
| `make clean` | Remove generated files |

TinyGo 0.44.0 is tested. Override tool paths with `GO=` or `TINYGO=`.

### Web UI

```sh
make web
python3 -m http.server 8765 --directory web/dist
```

Open <http://localhost:8765/>. Render, lint, and diff run locally in the browser. Pushes to `main` deploy `web/dist/` through GitHub Pages; enable “GitHub Actions” as the Pages source first.

## Lint

`ghmd lint` reports source locations that may render incorrectly on GitHub. `ghmd diff` reports structural differences between cmark-gfm and md4c. Findings use:

```text
path:line:column: severity rule: message
```

The command exits 1 when findings exist. See [RULES.md](RULES.md) for the complete rule list.

## Test

```sh
go test ./...
```

The suite includes 617 GitHub Markdown API fixtures, native-pipeline regression cases, and 29 automatic authoring rules (47 expected findings).

## Packages

- `internal/cmark-gfm`: cmark-gfm binding
- `internal/standard`: md4c binding
- `internal/github`: GitHub compatibility pipeline
- `internal/lint`: source and differential linting
- `internal/diff`: normalized parser comparison
- `cmd/ghmd`: CLI
- `cmd/ghmd-wasm`: browser WebAssembly bridge
- `web`: browser UI

## Limits

- 1 MB Markdown, 32 MB HTML output, tree depth 512.
- Raw HTML is allowlist-sanitized; unsafe tags, attributes, and URLs are removed.
- No syntax highlighting or remote image proxying.

## License

Original code is MIT licensed. `third_party/cmark-gfm` is cmark-gfm 0.29.0.gfm.13; `third_party/md4c` is md4c 0.5.2. See their provenance and license files.
