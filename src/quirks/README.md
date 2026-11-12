# GitHub-style Markdown renderer

This is the package's default renderer. It uses unmodified cmark-gfm sources,
project-owned C compatibility transforms, and the [native HTML core](../html/README.md)
to approximate GitHub's Markdown API output, including known rendering quirks.
Compatibility is checked against saved API responses; those cases do not prove
equivalence for every Markdown document.

## API and CLI

```ts
import { render } from "gfm-bun-render";

const html = await render("# Example\n\n$x^2$\n", {
  runId: "a".repeat(32),
});
```

`render(source: string, options?): Promise<string>` returns a sanitized HTML
fragment. Its only public option is `runId`, a 32-character lowercase hexadecimal
identifier used for generated math and footnote IDs. Omitting it generates a
fresh ID. Each render also creates a separate internal random boundary marker.
Reusing a `runId` in one page can produce anchor collisions.

Both compilation at module import and rendering run synchronously in the calling
process. The Promise preserves the public API. Input is limited to 1,000,000
UTF-8 bytes and final output to 32 MiB. The final operation removes one trailing
LF, when present.

From the repository root:

```sh
bun run render input.md
bun run render --run-id aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa input.md
bun run render < input.md
```

## Processing and ownership

TypeScript compiles [cmark/bun.c](cmark/bun.c) through the
[runtime `cc()` adapter](../ffi/README.md) once per module instance. Bun uses its
built-in compiler; Node.js uses the optional TinyCC/Koffi backend. A render makes
one native call, `gfm_bun_render`, for the complete operation:

1. Parse Markdown with cmark-gfm, apply AST compatibility transforms, filter raw
   source-authored inputs, and render HTML.
2. Parse HTML with Lexbor, apply provenance-dependent compatibility adjustments,
   and sanitize it with the C allowlist policy.
3. Reparse sanitized HTML and process templates, details, tables, and math,
   performing further reparses when a transformation requires them.
4. Finish links and other HTML adjustments, serialize the result, enforce output
   limits, and release intermediate strings and native trees.

All of these stages stay in C. JavaScript generates identifiers and handles the
public API and CLI; intermediate HTML, reparse flags, and native pointers do not
cross the binding.

[cmark/api.h](cmark/api.h) exposes the runtime-independent `gfm_render` function.
Its three inputs are borrowed UTF-8 views for Markdown, the run ID, and the
boundary; they need not be NUL-terminated. Successful output is a malloc-owned
string with an explicit byte length, freed by the caller. Failures return a
status and a static error message. `bun.c` owns string conversion, frees native
output after copying it into JavaScript, and translates errors into exceptions.

## Code layout

| File or group | Responsibility |
| --- | --- |
| [index.ts](index.ts) | Public `render` API and random identifiers. |
| [native-core.ts](native-core.ts) | Compiler configuration and the internal `renderDocument` binding. |
| [cli.ts](cli.ts) | File/stdin rendering and `--run-id`. |
| [cmark/unity.c](cmark/unity.c) | Includes upstream cmark-gfm, the HTML core, and owned modules; isolates colliding internal C names. |
| [cmark/document.c](cmark/document.c) | The complete operation, validation, intermediate-string ownership, and cleanup. |
| [cmark/markdown.c](cmark/markdown.c), [cmark/quirks.c](cmark/quirks.c), [cmark/math.c](cmark/math.c), [cmark/footnotes.c](cmark/footnotes.c) | Markdown parsing, AST changes, and initial rendering. |
| [cmark/dom.c](cmark/dom.c), [cmark/dom-compat.c](cmark/dom-compat.c), [cmark/dom-serialize.c](cmark/dom-serialize.c), [cmark/raw-html.c](cmark/raw-html.c) | Native HTML repair policy, serialization, provenance-based edits, and raw input filtering. |
| [cmark/tree.c](cmark/tree.c), [cmark/attributes.c](cmark/attributes.c), [cmark/classes.c](cmark/classes.c), [cmark/serialize.c](cmark/serialize.c) | Queries and edits over serialized HTML spans. |
| [cmark/prepare.c](cmark/prepare.c), [cmark/stages.c](cmark/stages.c), [cmark/math-dom.c](cmark/math-dom.c), and feature modules | Preparation, math placeholders, tables, lists, tasks, images, alerts, links, and footnotes. |
| [cmark/unicode.c](cmark/unicode.c), [cmark/unicode-lower-data.h](cmark/unicode-lower-data.h) | Unicode lowercase behavior, including contextual Greek sigma and generated Unicode tables. |

Lexbor performs HTML5 tokenization and tree construction. The serialized-HTML
span helpers operate on pipeline output; they do not replace parsing of raw
HTML. Source provenance supports the compatibility adjustments described in the
[HTML core README](../html/README.md#why-html-source-locations-are-retained).

Neither syntax highlighting nor remote image proxying is performed. Generated
math markup reproduces GitHub-style placeholders; comparison normalizes it with
the [VS Code-style renderer's](../vscode/README.md) KaTeX output.

## Maintenance and checks

Keep upstream sources unchanged and compatibility work in owned modules. The
preprocessor includes the sources at compile time; no pre-generated merged C
copy or external compiler process is required at runtime. Use Bun 1.4.4 or later,
or Node.js 24 or later with the optional peers installed and built. Both need the
standard C headers used by their compiler. Format handwritten C
with the repository's `.clang-format`; exclude vendored sources and generated
Unicode tables.

Run from the repository root:

```sh
# Complete suite
bun run test
# Saved references only, a subset of the complete suite
bun run test:e2e
```

The full suite covers bindings, native ownership, upstream integrity, individual
features, and saved references. `test:e2e` runs the saved API cases and fuzz
regressions offline. Migration-specific Markdown/HTML pairs live separately in
[test/fixtures/native-pipeline](../../test/fixtures/native-pipeline/README.md).
Private diagnostic helpers are confined to `test/support/native.c`.
`npm run test:node` separately checks the Node adapter and saved API references.
