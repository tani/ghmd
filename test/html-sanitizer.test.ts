import { expect, test } from "bun:test";
import { createDOM } from "./support/dom";
import { render } from "../src/quirks";
import { sanitizeHtml } from "../src/html";

// An independent DOM implementation reparses the serialized result. Inspect
// templates too: inert content can become active when later unwrapped.
function inspect(html: string, allowGenerated = false) {
  const window = createDOM();
  try {
    window.document.body.innerHTML = html;
    const visit = (node: any) => {
      if (node.nodeType === 1) {
        const tag = node.localName.toLowerCase();
        expect([
          "script",
          "style",
          "iframe",
          "object",
          "embed",
          "base",
          "link",
          "meta",
        ]).not.toContain(tag);
        for (const attr of node.attributes) {
          expect(attr.name).not.toMatch(/^on|^(?:srcdoc|formaction|is|xmlns(?::.*)?)$/i);
          if (attr.name === "style") {
            expect(allowGenerated).toBe(true);
            expect(attr.value).toMatch(/^(?:display: (?:inline-block|block)|max-width: 100%;)$/);
          }
          if (["href", "src", "action", "poster", "background", "cite"].includes(attr.name)) {
            const clean = attr.value.replace(
              /[\u0000-\u0020\u007f\u00a0\u1680\u180e\u2000-\u2029\u205f\u3000\ufeff]/g,
              "",
            );
            expect(clean).not.toMatch(/^(?:javascript|vbscript|file|blob):/i);
            if (tag !== "img" || attr.name !== "src") expect(clean).not.toMatch(/^data:/i);
          }
        }
      }
      for (const child of node.childNodes ?? []) visit(child);
      if (node.content) visit(node.content);
    };
    visit(window.document.body);
  } finally {
    window.happyDOM.abort();
  }
}

test("sanitizer retains safe GFM markup, MathML, task states and footnote data", () => {
  const source =
    '<details open><summary>More</summary><input type="checkbox" checked disabled aria-label="done"><a id="note-1" href="#note-1" data-footnote-ref>note</a><math><mfrac><mi>x</mi><mn>2</mn></mfrac></math></details>';
  const html = sanitizeHtml(source);
  expect(html).toContain('<details open="">');
  expect(html).toContain('checked="" disabled="" aria-label="done"');
  expect(html).toContain('data-footnote-ref=""');
  expect(html).toContain("<math><mfrac><mi>x</mi><mn>2</mn></mfrac></math>");
  inspect(html);
});

test("removing a comment or unknown tag never skips later unsafe attributes", async () => {
  for (const prefix of ["<!--removed-->", "<unknown>keep</unknown>", "<script>removed</script>"]) {
    const source = `${prefix}<div><img src=x onerror=alert(1)><a href="javascript:alert(2)" onclick=alert(3)>bad</a><template><img src=x onload=alert(4)></template></div>`;
    expect(sanitizeHtml(source)).not.toMatch(/onerror|onclick|onload|javascript:/i);
    inspect(sanitizeHtml(source));
    const result = await render(source);
    expect(result).not.toMatch(/onerror=|onclick=|onload=|href="javascript:/i);
    inspect(result, true);
  }
});

test("URL validation handles entity decoding, mixed case, controls and Unicode spaces", () => {
  const unsafe = [
    "javascript:alert(1)",
    "jav&#x61;script&#58;alert(1)",
    "java&#9;script:alert(1)",
    "\r\nJaVaScRiPt:alert(1)",
    "java\u2003script:alert(1)",
    "\ufeffjavascript:alert(1)",
    "vbscript:msgbox(1)",
    "file:///etc/passwd",
    "data:text/html,<script>alert(1)</script>",
    "blob:https://example.com/id",
  ];
  for (const url of unsafe) {
    const html = sanitizeHtml(
      `<a href="${url}">x</a><img src="${url}"><form action="${url}"></form>`,
    );
    expect(html).not.toContain("href=");
    expect(html).not.toContain("src=");
    expect(html).not.toContain("action=");
    inspect(html);
  }
  for (const url of [
    "",
    "#part",
    "/path",
    "../image.png",
    "//example.com/a",
    "https://example.com/x?a=1&amp;b=2",
    "MAILTO:a@example.com",
    "tel:+123",
  ]) {
    expect(sanitizeHtml(`<a href="${url}">x</a>`)).toContain("href=");
  }
});

test("validates every srcset candidate and preserves safe relative images", () => {
  expect(sanitizeHtml('<img srcset="a.png 1x, https://example.com/b.png 2x">')).toContain(
    "srcset=",
  );
  for (const value of [
    "a.png 1x, javascript:alert(1) 2x",
    "javascript:alert(1),a.png",
    "a.png 1x, data:text/html,attack 2x",
    "a.png 1x, b.png 2q",
  ]) {
    expect(sanitizeHtml(`<img srcset="${value}">`)).not.toContain("srcset=");
  }
});

test("duplicate attributes keep the first value and unsafe namespace attributes are removed", () => {
  expect(sanitizeHtml('<a href="javascript:alert(1)" href="https://example.com">x</a>')).toBe(
    "<a>x</a>",
  );
  expect(sanitizeHtml('<a href="https://example.com" href="javascript:alert(1)">x</a>')).toContain(
    'href="https://example.com"',
  );
  expect(
    sanitizeHtml(
      '<p style="color:red" onclick="x()" is="evil" xmlns="x" xmlns:xlink="x" xlink:href="javascript:x()">ok</p>',
    ),
  ).toBe("<p>ok</p>");
  expect(
    sanitizeHtml('<form id="attributes"><input name="__proto__"><p id="safe-id">ok</p></form>'),
  ).toBe('<form><input><p id="safe-id">ok</p></form>');
});

const attacks = [
  '<svg><foreignObject><iframe srcdoc="<script>alert(1)</script>"></iframe></foreignObject></svg>',
  '<math><annotation-xml encoding="text/html"><img src=x onerror=alert(1)></annotation-xml></math>',
  '<math><mtext><table><mglyph><style><!--</style><img title="--><img src=x onerror=alert(1)>">',
  '<svg><style><a title="</style><img src=x onerror=alert(1)>">x</a></style></svg>',
  "<template><template><svg><g onload=alert(1)></g></svg><img src=x onerror=alert(1)></template></template>",
  "<table><caption><template><textarea></textarea><img src=x onerror=alert(1)></template></caption></table>",
  '<noscript><p title="</noscript><img src=x onerror=alert(1)>">x</p></noscript>',
  "<textarea>&lt;/textarea>&lt;img src=x onerror=alert(1)></textarea>",
  '<div title="&quot; onmouseover=alert(1) &lt;script&gt;">text</div>',
  '<object><form action="javascript:alert(1)"><input name=action></form></object>',
  "<iframe><img src=x onerror=alert(1)></iframe><p>safe</p>",
];
for (const [index, attack] of attacks.entries()) {
  test(`malformed/foreign markup stays inert after reparsing: ${index + 1}`, async () => {
    const sanitized = sanitizeHtml(attack);
    inspect(sanitized);
    expect(sanitizeHtml(sanitized)).toBe(sanitized);
    inspect(await render(attack), true);
  });
}

test("deep templates and removed-node combinations cannot bypass traversal", () => {
  const containers = ["div", "span", "template", "math", "mtext", "table", "select", "unknown"];
  for (const outer of containers)
    for (const inner of containers) {
      const html = sanitizeHtml(
        `<${outer}><!--first--><${inner}><img src=x onerror=alert(1)><a href="jav&#x61;script:alert(1)">bad</a></${inner}></${outer}>`,
      );
      inspect(html);
      expect(html).not.toMatch(/onerror=|href="javascript:/i);
    }
});

test("sanitizer failures throw and do not return partially processed HTML", () => {
  expect(() => sanitizeHtml(null as any)).toThrow(TypeError);
  expect(() => sanitizeHtml("<template>".repeat(520))).toThrow(RangeError);
  expect(() => sanitizeHtml("<p>a<!--x-->b<!--y-->c</p>", "--><img src=x onerror=x()>")).toThrow();
  expect(() => sanitizeHtml("<p>safe</p>", "a".repeat(64))).toThrow(TypeError);
  expect(sanitizeHtml("<p>safe</p>")).toBe("<p>safe</p>");
});
