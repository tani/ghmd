# Lint rules

`ghmd lint` reports source locations as:

```text
path:line:column: severity rule: message
```

Line and column numbers are one-based. Columns count Unicode code points. A finding without a source position uses line 0, column 0. The command exits 1 when it reports at least one finding.

There are three rule families:

- **Source rules** detect raw HTML, math syntax, and DOM hazards directly in Markdown.
- **Differential rules** compare GitHub rendering against the independent md4c renderer.
- **Authoring rules** detect source patterns that can change GitHub output. Their emitted identifiers are `style/<ID>`.

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

Authoring checks produce `style/<ID>` findings. IDs are stable. The following 36 automatic rules emit a direct `style/` finding:

| Rule | Checks |
| --- | --- |
| `ALERT-001` | Alert type is one of `NOTE`, `TIP`, `IMPORTANT`, `WARNING`, or `CAUTION`; markers are not nested. |
| `BLOCK-001` | One title, ATX heading hierarchy, and conservative blank-line separation. |
| `BLOCK-002` | List continuation blocks are indented relative to the item content column. |
| `CODE-001` | Multiline code and literal Markdown examples are fenced with a safe fence length and language identifier. |
| `CODE-002` | Inline-code delimiters do not collide with embedded backticks; exact line breaks use fenced code. |
| `DETAILS-001` | Plain-text summaries, required blank lines, and closed `<details>` sections. |
| `DETAILS-002` | Math and diagram widgets are not placed inside collapsible sections without explicit verification. |
| `EXT-001` | Unsupported extensions such as highlight syntax, wiki links, colon directives, and PlantUML. |
| `FOOT-001` | Footnote identifiers are unique, lowercase, and consistently referenced; definitions are at the end. |
| `FOOT-002` | Footnote references are outside link/image labels, and labels stay on one line. |
| `FOOT-003` | Math is kept out of footnotes. |
| `LINK-001` | Important references use explicit descriptive links rather than bare-URL detection. |
| `MATH-001` | Alternate TeX delimiters require explicit parser support; GitHub's dollar-backtick syntax is not generic. |
| `MATH-003` | Inline math stays on one line and avoids TeX `%` comments. |
| `MATH-004` | Separate inline expressions are separated by prose or whitespace; adjacent dollars are rejected. |
| `MATH-005` | Complex math is moved outside links, footnotes, tables, emphasis, alerts, and HTML containers. Also checks block math in HTML containers. |
| `MATH-006` | Nested dollar math inside `\text{...}` is avoided, including in fenced math. |
| `MATH-007` | Normal TeX escapes are retained and not doubled to compensate for parser behavior. |
| `MATH-008` | Layout-sensitive TeX is reviewed: `\displaystyle`, large operators, roots around fractions, and matrices. |
| `MATH-009` | Unusually large inline or display expressions are split at valid TeX boundaries. |
| `MERMAID-001` | A lowercase `mermaid` fence starts with a valid diagram type. |
| `MERMAID-002` | Established diagram types are preferred; newer types are gated on capability tests. |
| `MERMAID-003` | Identifiers, labels, reserved words, Unicode quoting, and edge syntax are safe. |
| `MERMAID-004` | HTML labels, callbacks, click directives, external assets, and per-diagram configuration are not required for meaning. |
| `MERMAID-005` | Comments and configuration directives are kept explicit and portable. |
| `STRIKE-001` | Strikeout uses exactly two tildes on each side. |
| `STRIKE-002` | Strikeout stays on one line without overlapping emphasis. |
| `TABLE-001` | Tables have consistent cell counts, one header row, and a valid delimiter row. |
| `TABLE-002` | Table cells stay inline on one source line. |
| `TABLE-003` | Literal pipes in tables are escaped as `\|`. |
| `TABLE-004` | Math is moved out of tables and avoids literal relation pipes. |
| `TASK-001` | Task markers use `[ ]` or `[x]` after a list item and a space. |
| `TASK-002` | Extra literal checkbox markers are not placed on the first task line. |
| `TEXT-001` | Bold and emphasis delimiters are not nested, overlapping, or padded with whitespace. |
| `TEXT-002` | Emphasis adjacent to CJK text and punctuation is reviewed. |
| `WIDGET-001` | ASCII STL data is structurally valid; large models are linked. |

### Automatic rules covered by lower-level checks

Three automatic authoring rules intentionally do not emit `style/<ID>` themselves:

| Authoring rule | Emitted rule |
| --- | --- |
| `AUTO-001` | `style/LINK-001` |
| `HTML-001` | `raw-html/removed-container`, `raw-html/unsafe-attribute`, `raw-html/unsafe-url` |
| `MATH-002` | `standard/lost-math` |

## Coverage summary

- 25 source and differential finding rules.
- 36 direct automatic `style/` rules.
- 3 automatic authoring rules reported through lower-level rules.

The fixture suite in `test/fixtures/lint/` exercises automatic detection and negative controls. The command is not a proof that a target renderer produced correct visual output.
