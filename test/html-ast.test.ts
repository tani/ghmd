import { expect, test } from "bun:test";
import { htmlAst } from "./support/html-ast";

test("ignores HTML formatting and attribute order", () => {
  expect(htmlAst('<ul>\n<li class="a" id="b">x</li>\n</ul>')).toEqual(
    htmlAst('<ul><li id="b" class="a">x</li></ul>'),
  );
  expect(htmlAst("<li>hello <code>x</code>\n</li>")).toEqual(
    htmlAst("<li>hello <code>x</code></li>"),
  );
  expect(htmlAst("<p>a\n  b</p>")).toEqual(htmlAst("<p>a b</p>"));
});
test("retains meaningful whitespace and markup", () => {
  expect(htmlAst("<p>a <em>b</em></p>")).not.toEqual(htmlAst("<p>a<em>b</em></p>"));
  expect(htmlAst("<p>a b</p>")).not.toEqual(htmlAst("<p>ab</p>"));
  expect(htmlAst("<pre>a\n b</pre>")).not.toEqual(htmlAst("<pre>a b</pre>"));
  expect(htmlAst("<code>a  b</code>")).not.toEqual(htmlAst("<code>a b</code>"));
  expect(htmlAst("<math-renderer>$a  b$</math-renderer>")).not.toEqual(
    htmlAst("<math-renderer>$a b$</math-renderer>"),
  );
  expect(htmlAst("<p>x</p>")).not.toEqual(htmlAst("<div>x</div>"));
  expect(htmlAst('<a href="/a">x</a>')).not.toEqual(htmlAst('<a href="/b">x</a>'));
});
test("normalizes generated IDs only in their attributes", () => {
  const a = "0123456789abcdef0123456789abcdef",
    b = "abcdef0123456789abcdef0123456789";
  expect(htmlAst(`<math-renderer data-run-id="${a}">x</math-renderer>`)).toEqual(
    htmlAst(`<math-renderer data-run-id="${b}">x</math-renderer>`),
  );
  expect(htmlAst(`<p>${a}</p>`)).not.toEqual(htmlAst(`<p>${b}</p>`));
});

test("ignores syntax highlighting while preserving code and unrelated markup", () => {
  const options = { ignoreHighlighting: true };
  const highlighted =
    '<div class="highlight highlight-source-js"><pre class="notranslate"><span class="pl-k">const</span> x</pre></div>';
  const plain =
    '<pre lang="js" class="notranslate"><code class="notranslate">const x\n</code></pre>';
  expect(htmlAst(highlighted, options)).toEqual(htmlAst(plain, options));
  expect(htmlAst(highlighted, options)).not.toEqual(
    htmlAst(plain.replace("const x", "const y"), options),
  );
  expect(htmlAst("<pre><code><em>x</em></code></pre>", options)).not.toEqual(
    htmlAst("<pre><code>x</code></pre>", options),
  );
  expect(
    htmlAst('<div class="highlight"><pre><a href="/x">x</a></pre></div>', options),
  ).not.toEqual(htmlAst("<pre><code>x</code></pre>", options));
});
