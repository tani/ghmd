/** New independently submitted GitHub Markdown API cases. */
export interface CaptureCase {
  id: string;
  source: string;
  origin: string;
  rules: string[];
}
const cases: CaptureCase[] = [];
const add = (origin: string, casesByName: Record<string, string>) => {
  for (const [name, source] of Object.entries(casesByName))
    cases.push({
      id: `expanded-${origin}-${name}`,
      source,
      origin: `expanded:${origin}`,
      rules: [],
    });
};
add("task-context", {
  plain: "- [ ] `$x$`",
  newline: "- [ ] `$x$`\n",
  blank: "- [ ] `$x$`\n\n",
  heading: "## Heading\n\n- [ ] `$x$`",
  followed: "- [ ] `$x$`\n\n## Next",
  surrounded: "## Heading\n\n- [ ] `$x$`\n\n## Next",
  before: "- [ ] before `$x$`",
  empty: "- [ ] `$$`",
  preceding_math: "$x$\n\n- [ ] `$x$`",
  following_math: "- [ ] `$x$`\n\n$x$",
  earlier_task_math: "- [ ] $x$\n\n- [ ] `$x$`",
  checked: "- [x] `$x$`",
  nested: "- parent\n  - [ ] `$x$`",
  loose: "- [ ] `$x$`\n\n  continued paragraph",
  normal_list: "- `$x$`",
  code_plain: "- [ ] `x`",
  code_one_dollar: "- [ ] `$`",
});
add("math-nesting", {
  link_label: "[$x$](https://example.com)",
  link_destination: "[math](https://example.com/$x$)",
  emphasis: "*before $x$ after*",
  strong: "**before $`x`$ after**",
  strike: "~~$x$~~",
  table: "| a | b |\n| --- | --- |\n| $x$ | $`y`$ |",
  table_escaped_pipe: "| a |\n| --- |\n| $a \\| b$ |",
  footnote: "ref[^a]\n\n[^a]: $x$ and $`y`$",
  quote: "> before $x$ after",
  alert: "> [!NOTE]\n> $x$ and $`y`$",
  nested_list: "- parent\n  - $x$\n  - $`y`$",
  details: "<details>\n<summary>Summary $x$</summary>\n\n$x$\n\n</details>",
  raw_span: "<span>$x$</span> and $x$",
  raw_em: "<em>$x$</em> and $x$",
  code: "`$x$` and $x$",
  fenced_math: "```math\na + b\nc + d\n```",
  fenced_math_escape: "```math\na\\\\\nb\n```",
  display_list: "- $$\n  x + y\n  $$",
  display_quote: "> $$\n> x + y\n> $$",
  two_blocks: "$$\na\n$$\n\n$$\nb\n$$",
});
add("html-context", {
  unclosed_b: "$a <b > c$\n\nMore $x$",
  closed_b: "$a <b > c$\n\nMore $x$\n\n</b>\n\nLast $x$",
  unclosed_i: "$a <i > c$\n\nMore $x$",
  unclosed_em: "$a <em > c$\n\nMore $x$",
  nested_bad: "<b><i>one</b> two</i>\n\n$x$",
  link_unclosed: '<a href="https://example.com">x\n\n$x$',
  raw_comment: "<!-- $x$ -->\n\n$x$",
  html_table: "<table><tr><td>$x$</td></tr></table>\n\n$x$",
  html_block: "<div>\n\n$x$\n\n</div>\n\n$x$",
  html_escaped: "$a &lt;b &gt; c$\n\n$x$",
  tag_in_code: "`<b>`\n\n$x$",
});
add("document", {
  footnote_forward: "First[^a] and $x$.\n\n[^a]: A\n\nSecond[^a].",
  footnote_before: "[^a]: Definition $x$\n\nReference[^a] and $x$.",
  footnote_nested: "- task[^a]\n\n[^a]: first\n\n    - nested $x$",
  repeat_labels: "[^A] and [^a]\n\n[^a]: note",
  unclosed_before_list: "<b>\n\n- [ ] $x$\n\n</b>",
  list_after_math: "$$\nx\n$$\n\n- [ ] $x$",
  list_before_math: "- [ ] $x$\n\n$$\ny\n$$",
  list_paragraphs: "- [ ] first $x$\n\n  second $y$\n\n- [x] done $z$",
  references_after: "[reference][x] and $x$\n\n[x]: https://example.com",
});
add("highlight", {
  js: '```js\nconst x = "hello"; // hi\n```',
  javascript: "```javascript\nconst x = /a+/g;\n```",
  typescript: "```ts\ntype X = { x: string };\n```",
  python: '```python\ndef f(x):\n    return f"{x}"\n```',
  rust: '```rust\nfn main() { println!("hi"); }\n```',
  shell: '```sh\necho "$HOME"\n```',
  json: '```json\n{"x": [1, true, null]}\n```',
  yaml: "```yaml\na: [1, 2]\n```",
  html: '```html\n<div class="x">hello</div>\n```',
  css: "```css\n.x { color: red; }\n```",
  sql: "```sql\nSELECT x FROM t WHERE x > 1;\n```",
  unknown: "```unknown-language-xyz\n<a> & $x$\n```",
  text: "```text\n$x$ <b>\n```",
  mermaid: '```mermaid\ngraph TD\n A["hello"] --> B["world"]\n```',
  diff: "```diff\n- old\n+ new\n```",
});
add("sanitize", {
  script: "<script>alert(1)</script>\n\nSafe",
  event: '<img src="https://example.com/a.png" onerror="alert(1)">',
  javascript_link: "[bad](javascript:alert(1))",
  data_image: '<img src="data:image/svg+xml,<svg></svg>">',
  style: '<span style="color:red" onclick="alert(1)">text</span>',
  iframe: '<iframe src="https://example.com"></iframe>',
  svg: '<svg><circle cx="5" cy="5" r="2"/></svg>',
  details: "<details><summary>Title</summary>text</details>",
  input: '<input type="checkbox" checked disabled>',
  anchor: '<a href="javascript:alert(1)" title="ok">label</a>',
  entity_url: '<a href="jav&#x61;script:alert(1)">bad</a>',
  unknown: '<custom-element data-x="1">Text</custom-element>',
});
add("images-widgets", {
  image_https: "![alt](https://example.com/test.png)",
  image_title: '![alt text](https://example.com/a.png "title")',
  image_relative: "![relative](./image.png)",
  image_link: "[![alt](https://example.com/a.png)](https://example.com)",
  image_query: "![alt](https://example.com/a.png?a=1&b=2)",
  image_html: '<img src="https://example.com/a.png" alt="hi" width="20">',
  task_code: "- [ ] `code`\n- [x] done",
  mermaid_sequence: "```mermaid\nsequenceDiagram\n A->>B: hello\n```",
  geojson: '```geojson\n{"type":"Point","coordinates":[1,2]}\n```',
  csv: "```csv\na,b\n1,2\n```",
  alert_invalid: "> [!BOGUS]\n> text",
  alert_math: "> [!WARNING]\n> $x$",
});
add("size", {
  paragraph_16k: "a".repeat(16000) + " $x$",
  repeated_sections: Array.from(
    { length: 80 },
    (_, i) => `## Section ${i}\n\nParagraph $x_${i}$ and **bold**.\n`,
  ).join("\n"),
  nesting: Array.from({ length: 12 }, (_, i) => `${"  ".repeat(i)}- item ${i}`).join("\n"),
});
add("followup", {
  raw_b: "<b>$x$</b> and $x$",
  raw_i: "<i>$x$</i> and $x$",
  raw_em: "<em>$x$</em> and $x$",
  raw_strong: "<strong>$x$</strong> and $x$",
  raw_del: "<del>$x$</del> and $x$",
  raw_span_code: "<span>`$x$`</span> and $x$",
  raw_span_link: "<span>[$x$](https://example.com)</span>",
  nested_del_em: "~~*$x$*~~",
  nested_del_strong: "~~**$x$**~~",
  summary_em: "<details><summary>*$x$*</summary>\n\n$x$\n\n</details>",
  display_list_loose: "- first\n\n  $$\n  x\n  $$",
  display_html: "<div>\n\n$$\nx\n$$\n\n</div>",
  displayed_preceding: "$$\nx\n$$\n\n- [ ] `$y$`",
  inline_prior_code: "`$x$`\n\n- [ ] `$y$`",
  prior_math_after_code: "- [ ] `$y$`\n\n$x$",
  prior_math_same_list: "- $x$\n- [ ] `$y$`",
  prior_math_other_tag: "<span>$x$</span>\n\n- [ ] `$y$`",
  span_style: '<span style="color:red">$x$</span>',
  svg_text: "<svg><text>hello</text></svg>",
  raw_input_in_list: '- <input type="checkbox" disabled> task',
  data_png: '<img src="data:image/png;base64,YQ==">',
  image_root: '<img src="https://example.com/b.png">',
  image_paragraph: 'before <img src="https://example.com/b.png"> after',
  misnested_anchor: '<a href="https://example.com">hello\n\nworld',
});
export default cases;
