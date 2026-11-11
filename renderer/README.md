# gfm-bun-render

A runnable prototype with one public API:

```ts
import {render} from './src/index.ts';

const html: string = await render('# Hello\n\nMath $`\\frac{1}{2}`$.');
```

`render(source: string): Promise<string>` resolves to an HTML **fragment**. It is a renderer, not a linter, API client, or comparison report. It uses the requested stack: **bun:ffi + cmark-gfm + Bun.XML + happy-dom + MathJax + Mermaid**. Rendering runs in the calling Bun process. No child process, Worker, Python, jsdom, Chromium, or remote service is used.

## Run

Bun 1.4.4 and a POSIX C compiler are required for setup. The native library has been built and tested on Linux x86-64; the build script also has an untested macOS path. Windows is not supported by the build script.

```sh
bun install --frozen-lockfile
bun run setup
bun test
bun run render input.md > output.html
```

For stdin: `bun run render < input.md > output.html`.

## What renders

- GFM headings, paragraphs, lists, tables, autolinks, strikethrough, task lists, and footnotes through cmark-gfm 0.29.0.gfm.13.
- Inline `$...$`, protected dollar-backtick math, display `$$...$$`, and fenced `math` blocks into **real MathJax SVG**, with local glyph paths rather than external font resources.
- Fenced `mermaid` diagrams into **real Mermaid SVG**. Strict Mermaid sanitization remains enabled.
- Invalid TeX or Mermaid becomes an escaped, visible `math-error` or `mermaid-error` element containing the original source. Operational rendering failures reject the Promise. A missing native library fails module loading; run setup first.
- Ordinary code blocks, inline code, and link destinations are opaque to the math recognizer. LaTeX `\(...\)` and `\[...\]` delimiters are not enabled.

The renderer returns only an HTML fragment, without a doctype or stylesheet for the surrounding page. The adjacent `../github-html/` directory contains the four previously saved GitHub Markdown API response fragments, unchanged.

## In-process asynchronous API

`await render(source)` directly calls cmark-gfm with `bun:ffi`, transforms the HTML with happy-dom, and awaits Mermaid's own asynchronous rendering API. There are no subprocess calls in `src/`. The setup script invokes the C compiler only when building the native library.

Mermaid is a singleton and DOMPurify retains its first DOM, so the module reuses one private DOM. Complete renders are queued: `Promise.all(inputs.map(render))` is supported, with each request executing in order rather than simultaneously. A rejected request does not poison the queue. Each request resets Mermaid configuration, uses fresh MathJax input state, and clears the document afterward.

Browser globals are installed temporarily while a render executes. Their original property descriptors are restored in `finally`, including on error. They are visible to other application code while the Promise is pending; this is not an isolated JavaScript realm. Other code must not change or use these DOM globals concurrently, and should not initialize another Mermaid/DOMPurify environment in the same module graph. There is no per-call process startup or module reimport.

Maximum Markdown size is 1 MB, formula size 20,000 characters, diagram size 50,000 characters, and returned HTML limit 32 MB. The old child-process timeout has been removed: in-process synchronous work cannot be forcibly cancelled by a Promise timeout.

## Native and XML boundaries

The vendored C source and its provenance are under `vendor/cmark-gfm/`. `scripts/build-native.ts` builds one shared library containing cmark-gfm, its GFM extensions, and a small same-library memory-free function. Every returned C string is copied before release, ASTs and parsers are released in `finally`, and Markdown is passed with its UTF-8 byte length.

`Bun.XML.parse(..., {compact:false})` reads cmark's source-position AST. The math recognizer uses that AST to open math only inside text, before a second cmark pass can destroy the TeX. It is a bounded delimiter recognizer, not a complete parser for every Markdown/math ambiguity. Fenced math is recovered from the cmark code element. Generated SVG is parsed and serialized with **Bun.XML's ordered shape**, preserving child and text order.

One documented native patch adds XML-safe type names for footnote definition/reference nodes. The upstream version otherwise emits malformed `<<unknown>` tags in its XML output. The patch changes those names, not Markdown parsing or HTML rendering. See `patches/footnote-xml.patch`.

## happy-dom compatibility adaptations

The renderer includes two narrow adapters for happy-dom 21.0.4:

1. MathJax gets the exported happy-dom base constructors, because its `instanceof` document checks do not accept happy-dom's per-window constructor chain.
2. DOMPurify gets a standard-result `Node.nodeName` getter. Standalone SVG containing `<style>` is parsed as XML and imported into HTML, working around happy-dom's HTML parser dropping such SVG subtrees. DOMPurify's strict sanitization is **not disabled**. Sanitized diagram CSS is placed alongside the SVG in the returned figure, so reading the fragment back through happy-dom also retains the diagram nodes.

These adaptations belong to the private happy-dom window; browser globals are restored after each render. They do not replace geometry measurements or relax Mermaid to `securityLevel: 'loose'`.

## Prototype limits

**happy-dom does not perform browser layout. Its SVG `getBBox()` and text measurements return empty/zero values.** Mermaid SVG generation and preservation of nodes, edges, labels, and XML are tested, but diagram sizing, spacing, clipping, and font accuracy are **not** guaranteed. No fake bounding-box implementation is installed. A mathematically nonempty SVG is not a claim of correct visual layout. For example, the simple flowchart can retain both labels while receiving an undersized viewBox.

This is not a byte-compatible GitHub renderer. In particular, it generates MathJax/Mermaid SVG rather than GitHub's client widgets, retains cmark's own parser quirks, and does not reproduce GitHub's syntax highlighting, full sanitizer policy, alerts, footnote ID tokens, map widgets, or STL viewer. Raw Markdown HTML stays disabled by cmark. It is a local prototype, not a security-audited service for arbitrary hostile input.

## Dependencies and licenses

Direct npm dependencies are pinned to `happy-dom@21.0.4`, `@mathjax/src@4.1.3`, and `mermaid@12.3.0`, with `bun.lock`. The native source is extracted from the cmarkgfm 2025.10.22 source distribution and identifies itself as 0.29.0.gfm.13; the source archive SHA-256 is recorded. cmark-gfm's license is retained in `vendor/cmark-gfm/COPYING`. npm dependency license notices remain in the installed packages. Original prototype code is MIT licensed; third-party code retains its own license.
