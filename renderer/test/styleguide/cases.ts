/** Original synthetic inputs, grouped to keep independent footnotes and raw
 * HTML from changing the interpretation of unrelated cases. */
export type Case = { id: string; rules: string[]; group: string; source: string };
export const cases: Case[] = [
  {
    id: "cjk-emphasis",
    rules: ["TEXT-002"],
    group: "syntax",
    source: "これは**重要。**です\n\nこれは**重要**です。\n\n日本語 $`x`$。",
  },
  {
    id: "table-cell-count",
    rules: ["TABLE-001"],
    group: "syntax",
    source: "| A | B |\n| --- | --- |\n| 1 | 2 | 3 |\n| short |",
  },
  {
    id: "table-pipes",
    rules: ["TABLE-003"],
    group: "syntax",
    source: "| A |\n| --- |\n| `a\\|b` |\n| C \\\\| D |\n| E \\\\\\\\| F |",
  },
  {
    id: "table-unicode-empty",
    rules: ["TABLE-004"],
    group: "syntax",
    source: "| 日本語 | 😀 |\n| --- | --- |",
  },
  {
    id: "autolink-brackets",
    rules: ["LINK-001", "AUTO-001"],
    group: "syntax",
    source: "[ www.example.com\n\n[ https://example.com\n\na www. b\n\nWWW.COMMONMARK.ORG",
  },
  {
    id: "links-relative-unicode",
    rules: ["LINK-001", "LINK-002"],
    group: "syntax",
    source:
      "日本語https://example.com/a。\n\n[relative](../docs/a.md)\n\nhttps://example.com/\\(foo\\)",
  },
  {
    id: "tasks-markers-nesting",
    rules: ["TASK-001", "TASK-003"],
    group: "syntax",
    source: "- [ ] `[x]`\n- [x] done\n  - [ ] nested\n- [ ] literal [x] appears later",
  },
  {
    id: "strike-delimiters",
    rules: ["STRIKE-001", "STRIKE-002"],
    group: "syntax",
    source: "~single~ ~~double~~\n\na~~***nested***~~a\n\n~~across\nlines~~",
  },
  {
    id: "footnote-case-repeat",
    rules: ["FOOT-001"],
    group: "footnote",
    source: "Text[^A][^a]\n\n[^A]: Upper\n[^a]: Lower",
  },
  {
    id: "footnote-whitespace-colon",
    rules: ["FOOT-001"],
    group: "footnote",
    source: "[^:][^ spaced ]\n\n[^:]: colon\n[^ spaced ]: spaces",
  },
  {
    id: "footnote-bracket-multiline",
    rules: ["FOOT-001", "FOOT-002"],
    group: "footnote",
    source: "[^a\nb]: line\n\nText[^a\nb] and [^x\\]y].\n\n[^a]: note\n[^x\\]y]: other",
  },
  {
    id: "footnote-link-image",
    rules: ["FOOT-002"],
    group: "footnote",
    source: "[text [^b] end](https://example.com)\n\n[^![image](#)]\n\n[^b]: note",
  },
  {
    id: "footnote-math",
    rules: ["FOOT-003"],
    group: "footnote",
    source: "Text[^a].\n\n[^a]: $x$ and $`x`$",
  },
  {
    id: "alerts-five",
    rules: ["ALERT-001"],
    group: "rich",
    source: ["NOTE", "TIP", "IMPORTANT", "WARNING", "CAUTION"]
      .map((x) => `> [!${x}]\n> Message.`)
      .join("\n\n"),
  },
  {
    id: "alerts-invalid-nested",
    rules: ["ALERT-001"],
    group: "rich",
    source: "> [!UNKNOWN]\n> Message.\n\n> > [!NOTE]\n> > Nested.",
  },
  {
    id: "mermaid-source",
    rules: ["MERMAID-003", "MERMAID-004"],
    group: "rich",
    source:
      '```mermaid\nflowchart TD\n A["日本語"] --> B["end"]\n A ---o C\n```\n\n```mermaid\nnot a diagram\n```',
  },
  {
    id: "data-widgets",
    rules: ["WIDGET-001", "WIDGET-002"],
    group: "rich",
    source:
      '```geojson\n{"type":"FeatureCollection","features":[]}\n```\n\n```topojson\ninvalid\n```\n\n```stl\nsolid a\nendsolid a\n```',
  },
  {
    id: "literal-source",
    rules: ["CODE-001", "MATH-003"],
    group: "rich",
    source: "`$x$` and `<script>`\n\n```text\n$`x`$\n<details>\nflowchart TD\n```",
  },
  {
    id: "html-dangerous",
    rules: ["HTML-001", "RENDER-003"],
    group: "security",
    source:
      "<script>alert(1)</script>\n\n<img src=x onerror=alert(1)>\n\n[bad](javascript:alert(1))",
  },
  {
    id: "html-misnested",
    rules: ["RENDER-001"],
    group: "html",
    source: "$a <b > c$\n\nMore $x$\n\n</b>\n\nLast $x$",
  },
];
