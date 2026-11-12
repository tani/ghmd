# VS Code-style Markdown renderer

This renderer approximates VS Code's built-in Markdown preview with
`markdown-it` and `@vscode/markdown-it-katex`. It provides the comparison target
for the [GitHub-style renderer](../quirks/README.md). It uses a fresh Markdown
engine per render and does not call cmark-gfm.

## API

```ts
import { render } from "gfm-bun-render/vscode";

const html = await render("# Example\n\n$x^2$\n", {
  linkify: true,
  breaks: false,
  math: true,
});
```

`render(source: string, options?): Promise<string>` returns an HTML fragment.
Conversion runs synchronously; the function wraps its result in a Promise.
Input is limited to 1,000,000 UTF-8 bytes and output to 32 MiB.

| Option | Default | Preview behavior |
| --- | --- | --- |
| `linkify` | `true` | Recognize plain-text URLs, corresponding to `markdown.preview.linkify`. |
| `breaks` | `false` | Turn ordinary newlines into line breaks, corresponding to `markdown.preview.breaks`. |
| `math` | `true` | Enable the KaTeX plugin, corresponding to `markdown.math.enabled`. |

[index.ts](index.ts) also exports `createMarkdownEngine(options?)`, which returns
a configured `MarkdownIt` instance for rendering and source-map analysis. It
enables raw HTML, disables typographic substitutions and fuzzy links, and uses
the math plugin with fenced blocks enabled, global grouping, and fresh macros.
The renderer performs no code syntax highlighting.

## Comparison and source mapping

The [comparison module](../compare/README.md) calls `createMarkdownEngine()` with
default options and adds `markdown-it-source-map` to a separate analysis engine.
Only that analysis rendering receives temporary mapping attributes. Normal
`render()` output has no added source-map attributes. Markdown locations come
from the plugin, tokens, and source matching; the GitHub-style renderer does not
embed a corresponding Markdown source map.

## Scope and checks

This module does not reproduce the entire VS Code preview: webview styles,
security settings, additional extensions, and interactive widgets are outside
its scope. GitHub-specific features such as footnotes, task checkboxes, alerts,
and Mermaid widgets do not have dedicated implementations here. KaTeX error
markup is reported separately by comparison.

Returned HTML is unsanitized. Consumers displaying it must apply a sanitizer
appropriate to their environment.

Run from the repository root:

```sh
bun test test/vscode.test.ts test/source-map.test.ts
```

The [root README](../../README.md) describes CLI comparisons, saved references,
and the limitations of interpreting renderer differences.
