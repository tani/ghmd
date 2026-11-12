# Native HTML core

This directory provides Lexbor HTML5 parsing, HTML source locations, and a C
allowlist sanitizer. The [GitHub-style renderer](../quirks/README.md) includes the
C core directly. Comparison and tests use a separate Bun adapter to obtain a
JavaScript tree.

## JavaScript API

This is an internal module, imported through `src/html/index.ts`; it is not a
package export. Importing it compiles `native.c` once per module instance with
`bun:ffi` `cc()`.

```ts
import { parseFragment, serialize } from "./src/html";

const tree = parseFragment('<p title="a>b">Hello</p>', {
  sourceCodeLocationInfo: true,
});
console.log(tree.childNodes?.[0].sourceCodeLocation?.startTag);
console.log(serialize(tree));
```

`parseFragment(html, options?)` returns an `HtmlNode` tree synchronously. Location
tracking defaults to off. Template contents live in `content`, and ordinary
children have non-enumerable `parentNode` references. Parsing uses a template
fragment context, preserving standalone table rows as well as HTML5 tree repair,
foster parenting, templates, and SVG names and attributes.

`serialize(node)` serializes the node's children, or its template content. It
does not include the supplied node's own opening and closing tags. These APIs
parse and serialize HTML; they do not apply the sanitizer.

The native entry point is `gfm_html_parse`. It returns JSON, which `index.ts`
converts to JavaScript objects. Native trees, input buffers, and output buffers
are freed before the call returns; no native pointer escapes to JavaScript.

## Why HTML source locations are retained

These locations describe the HTML string being parsed. They are separate from
the [Markdown source map used by comparison](../compare/README.md#source-locations).
They retain original start-tag and attribute spans, explicit end tags, and the
shared origin of elements reconstructed during HTML5 repair. Implied elements
such as a generated `tbody` have no original tag location.

The renderer uses this provenance to identify reconstructed anchors, reproduce
trailing-whitespace behavior, and remove source-authored `input` tags by their
exact byte spans while preserving surrounding source. Comparison uses it to
locate raw HTML tags and attributes within Markdown tokens. The metadata stays
in native node data or the optional JSON tree; it is not embedded in rendered
HTML attributes.

| Tracking mode | Use |
| --- | --- |
| `H_TRACK_NONE` | Parsing and internal reparses that need no locations. |
| `H_TRACK_BYTES` | Native renderer edits using UTF-8 byte offsets. |
| `H_TRACK_UTF16` | JavaScript-facing locations using UTF-16 offsets. |

Tracked input is fed in chunks ending at each literal `>`. Lexbor's tokenizer
decides which delimiter actually ends a tag, including quoted `>` characters;
this supplies exact tag ends without a second HTML tokenizer. Duplicate
attributes retain the first attribute's source span.

## Native implementation

| File | Responsibility |
| --- | --- |
| [core.c](core.c) | Assembles the runtime-independent C implementation in an explicit include order. |
| [lexbor.c](lexbor.c) | Scoped `#define` / `#include` adapter that routes element creation through the provenance wrapper. |
| [buffer.c](buffer.c) | Bounded native output buffers and allocation errors. |
| [parser.c](parser.c) | Parser ownership, token callbacks, fragment parsing, and source provenance. |
| [json.c](json.c) | Serializes native trees and optional locations to JSON. |
| [sanitize.c](sanitize.c) | Element/attribute allowlists, URL checks, and recursive sanitization. |
| [native.c](native.c) | Bun/Node-API adapter for parsing only. |
| [index.ts](index.ts) | Native compilation, JavaScript trees, types, and child serialization. |

The sanitizer is private C code used by the renderer. It checks descendants and
template contents, removes executable containers and unsafe attributes or URLs,
and unwraps unknown HTML tags. Its allowlist data is adapted from DOMPurify;
see [NOTICE](NOTICE) and [POLICY-LICENSE](POLICY-LICENSE). Tests access the
standalone sanitizer through [test/support/native.c](../../test/support/native.c).

The [vendored Lexbor distribution](../../vendor/lexbor/README.md) is byte-for-byte
upstream and verified by SHA-256. The preprocessor adapter distinguishes typed
declarations/definitions from calls using the pinned source's argument syntax,
then undefines its macros immediately after inclusion. Review the adapter when
upgrading Lexbor.

The JavaScript parser accepts at most 32 MiB of UTF-8 input. Native buffers are
limited to 64 MiB, and tree serialization and sanitizer traversal enforce a depth
limit of 512. Failures raise exceptions rather than returning partial output.

Run from the repository root:

```sh
bun test test/html-parser.test.ts test/html-sanitizer.test.ts test/html-ast.test.ts
```
