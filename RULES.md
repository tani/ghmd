# Lint rules

`ghmd lint` reports source locations as:

```text
path:line:column: severity rule: message
```

Line and column numbers are one-based. Columns count Unicode code points. A finding without a source position uses line 0, column 0. The command exits 1 when it reports at least one finding.

There are three rule families:

- **Source rules** detect raw HTML, math syntax, and DOM hazards directly in Markdown.
- **Differential rules** compare GitHub rendering against the independent md4c renderer.
- **Authoring rules** enforce source-observable Markdown style and compatibility policy. Their emitted identifiers are `style/<ID>`.

Only `raw-html/removed-container`, `raw-html/unsafe-attribute`, and `raw-html/unsafe-url` are errors. The remaining rules are warnings.

## Raw HTML and DOM

| Rule | Severity | Finds |
| --- | --- | --- |
| `raw-html/removed-container` | error | Executable or filtered containers whose contents GitHub removes. |
| `raw-html/unsafe-attribute` | error | Event handlers, namespace attributes, inline styles, and other unsafe attributes. |
| `raw-html/unsafe-url` | error | Unsafe URL schemes in source HTML attributes. |
| `raw-html/input` | warning | Source-authored `<input>` elements, which GitHub removes. |
| `raw-html/template` | warning | `<template>` contents that are unwrapped or re-parented by GitHub. |
| `html/nested-anchor` | warning | An `<a>` element nested inside another anchor. |
| `html/block-in-p` | warning | A block element inside a `<p>` element. |
| `html/unclosed-tag` | warning | An unclosed tag that can swallow following Markdown. |
| `html/stray-close` | warning | A closing tag without a matching open tag. |

## Standard-renderer differentials

These rules compare the GitHub output with md4c's GitHub dialect. Missing or changed content produces a warning:

| Rule | Finds |
| --- | --- |
| `standard/lost-text` | Visible text lost or changed. |
| `standard/lost-formatting` | Lost or changed inline formatting such as `strong`, `em`, and `del`. |
| `standard/lost-structure` | Lost or changed structural elements. |
| `standard/lost-code` | Lost or changed code content. |
| `standard/lost-link` | Lost or changed links. |
| `standard/lost-image` | Lost or changed images. |
| `standard/lost-math` | Math present in the standard parse but missing or changed in GitHub output. |
| `standard/task-state` | Changed task-list state. |

Differential checks are skipped where md4c semantics are not a usable baseline: footnote syntax and SVG. Known dialect noise is suppressed for math delimiters, footnote fallback links, alert markers, autolinks, table-pipe code, currency-like `$...$`, and presentation-only attributes.

## Math syntax

| Rule | Finds |
| --- | --- |
| `math/not-recognized` | Inline math that violates GitHub's dollar boundaries, including dollar-backtick math immediately after a letter. |
| `math/unclosed-dollar` | An unclosed inline dollar delimiter. |
| `math/footnote-literal` | Inline math inside a footnote definition, which GitHub leaves literal. |
| `math/display-not-isolated` | `$$...$$` display delimiters that are not isolated on their own lines. |
| `math/stripped-escape` | ASCII-punctuation escapes such as `\{`, `\_`, `\#`, and `\$` in bare `$...$` or isolated `$$...$$` math, where GitHub removes the backslash before math rendering. |
| `math/disallowed-macro` | `\operatorname` and `\operatorname*`, which GitHub rejects with "The following macros are not allowed", plus TeX macro-definition commands such as `\def`, `\newcommand`, `\let`, and `\DeclareMathOperator`, which GitHub does not execute as definitions. |
| `math/html-lookalike` | A `<tag`-like sequence inside display or fenced math that GitHub may interpret as HTML and remove from the expression. |
| `math/unknown-macro` | A TeX command absent from the combined KaTeX/MathJax command allowlist. Document-local definitions do not suppress this finding. |

The command allowlist is generated from the KaTeX v0.19.1 support table and the MathJax documentation command index. It is a syntax-recognition aid, not a GitHub capability guarantee: a command present in KaTeX or MathJax can still be rejected or rendered literally by GitHub, as `\operatorname` and macro definitions are. `math/disallowed-macro` therefore remains an explicit GitHub-specific denial list.

## Authoring rules

Authoring checks produce `style/<ID>` findings. IDs are stable. The following 49 automatic rules emit a direct `style/` finding:

| Rule | Checks |
| --- | --- |
| `ALERT-001` | Alert type is one of `NOTE`, `TIP`, `IMPORTANT`, `WARNING`, or `CAUTION`; markers are not nested. |
| `ALERT-002` | Alerts are used sparingly; portable output uses a plain warning paragraph. |
| `BLOCK-001` | One title, ATX heading hierarchy, and conservative blank-line separation. |
| `BLOCK-002` | List continuation blocks are indented relative to the item content column. |
| `CODE-001` | Multiline code and literal Markdown examples are fenced with a safe fence length and language identifier. |
| `CODE-002` | Inline-code delimiters do not collide with embedded backticks; exact line breaks use fenced code. |
| `COMPAT-001` | Rich GitHub features without a declared compatibility profile. |
| `DETAILS-001` | Plain-text summaries, required blank lines, and closed `<details>` sections. |
| `DETAILS-002` | Math and diagram widgets are not placed inside collapsible sections without explicit verification. |
| `EXT-001` | Unsupported extensions such as highlight syntax, wiki links, colon directives, and PlantUML. |
| `FOOT-001` | Footnote identifiers are unique, lowercase, and consistently referenced; definitions are at the end. |
| `FOOT-002` | Footnote references are outside link/image labels, and labels stay on one line. |
| `FOOT-003` | Math is kept out of footnotes. |
| `FOOT-004` | Footnotes are not used where GitHub does not support them, such as wikis. |
| `LINK-001` | Important references use explicit descriptive links rather than bare-URL detection. |
| `LINK-002` | Relative paths match the document location and are not silently rewritten. |
| `LINK-003` | Headings are simple and unique enough for stable section links. |
| `MATH-002` | New `github-rich` content avoids bare `$...$` and `$$...$$` math. |
| `MATH-003` | Alternate TeX delimiters require explicit parser support; GitHub's dollar-backtick syntax is not generic. |
| `MATH-005` | Inline math stays on one line and avoids TeX `%` comments. |
| `MATH-006` | Separate inline expressions are separated by prose or whitespace; adjacent dollars are rejected. |
| `MATH-007` | Complex math is moved outside links, footnotes, tables, emphasis, alerts, and HTML containers. Also checks block math in HTML containers. |
| `MATH-008` | Nested dollar math inside `\text{...}` is avoided, including in fenced math. |
| `MATH-009` | Normal TeX escapes are retained and not doubled to compensate for parser behavior. |
| `MATH-010` | Layout-sensitive TeX is reviewed: `\displaystyle`, large operators, roots around fractions, and matrices. |
| `MATH-011` | Unusually large inline or display expressions are split at valid TeX boundaries. |
| `MEDIA-001` | Images have meaningful alt text and essential visual information has a textual account. |
| `MERMAID-001` | A lowercase `mermaid` fence starts with a valid diagram type. |
| `MERMAID-002` | The target Mermaid engine and diagram type are verified and recorded. |
| `MERMAID-003` | Established diagram types are preferred; newer types are gated on capability tests. |
| `MERMAID-004` | Identifiers, labels, reserved words, Unicode quoting, and edge syntax are safe. |
| `MERMAID-005` | HTML labels, callbacks, click directives, external assets, and per-diagram configuration are not required for meaning. |
| `MERMAID-006` | Comments and configuration directives are kept explicit and portable. |
| `MERMAID-007` | Diagrams are split when they exceed five nodes or relationships. |
| `MERMAID-008` | A nearby prose description and accessible diagram metadata are supplied. |
| `SOCIAL-001` | Mentions, issue references, emoji shortcodes, and color swatches are not the only expression of meaning. |
| `STRIKE-001` | Strikeout uses exactly two tildes on each side. |
| `STRIKE-002` | Strikeout stays on one line without overlapping emphasis. |
| `TABLE-001` | Tables have consistent cell counts, one header row, and a valid delimiter row. |
| `TABLE-002` | Table cells stay inline on one source line. |
| `TABLE-003` | Literal pipes in tables are escaped as `\|`. |
| `TABLE-004` | Table width and padding are conservative; source alignment is cosmetic. |
| `TABLE-005` | Math is moved out of tables and avoids literal relation pipes. |
| `TASK-001` | Task markers use `[ ]` or `[x]` after a list item and a space. |
| `TASK-003` | Extra literal checkbox markers are not placed on the first task line. |
| `TEXT-001` | Bold and emphasis delimiters are not nested, overlapping, or padded with whitespace. |
| `TEXT-002` | Emphasis adjacent to CJK text and punctuation is reviewed. |
| `WIDGET-001` | Every nonessential widget has a textual summary and essential data remains accessible. |
| `WIDGET-002` | GeoJSON/TopoJSON and ASCII STL data are structurally valid; large data is linked. |

### Automatic rules covered by lower-level checks

Four automatic authoring rules intentionally do not emit `style/<ID>` themselves:

| Authoring rule | Emitted rule |
| --- | --- |
| `AUTO-001` | `style/LINK-001` |
| `HTML-001` | `raw-html/removed-container`, `raw-html/unsafe-attribute`, `raw-html/unsafe-url` |
| `MATH-001` | `style/COMPAT-001` |
| `MATH-004` | `standard/lost-math` |

### Manual review rules

These rules require intent, external evidence, or renderer behavior that cannot be inferred reliably from source, so they never produce automatic findings:

| Rule | Review guidance |
| --- | --- |
| `BLOCK-003` | A visible line break is author intent; ordinary soft-wrapped paragraphs are not violations. |
| `TASK-002` | A task checkbox is document data, not a portable interactive control. |
| `COMPAT-002` | Content preservation and compatibility-profile choice require author intent. |
| `COMPAT-003` | Specification behavior, current defects, historical bugs, and untested risks require external evidence. |
| `RENDER-001` | Renderer parsing and math-node construction are implementation responsibilities. |
| `RENDER-002` | Reproducible parser and renderer configuration is a build responsibility. |
| `RENDER-003` | HTML and identifier security policy is an application responsibility. |
| `RENDER-004` | Execution limits and failure isolation are renderer operations. |
| `RENDER-005` | Source preservation and revision-safe edits are renderer operations. |
| `RENDER-006` | Full-document and resource-failure coverage is a test-plan responsibility. |
| `RENDER-007` | Streaming-completion handling is a renderer state-machine responsibility. |

## Coverage summary

- 25 source and differential finding rules.
- 49 direct automatic `style/` rules.
- 4 automatic authoring rules reported through lower-level rules.
- 11 manual review rules.

The fixture suite in `test/fixtures/lint/` exercises automatic detection and negative controls. The command is not a proof that a target renderer produced correct visual output.
