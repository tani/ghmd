package verify

import (
	"reflect"
	"testing"
)

func same(a, b string, ignoreHighlighting bool) bool {
	return reflect.DeepEqual(HTMLAST(a, ignoreHighlighting), HTMLAST(b, ignoreHighlighting))
}

func TestHTMLASTIgnoresFormattingAndAttributeOrder(t *testing.T) {
	if !same("<ul>\n<li class=\"a\" id=\"b\">x</li>\n</ul>", `<ul><li id="b" class="a">x</li></ul>`, false) {
		t.Fatal("attribute order or block formatting changed the tree")
	}
	if !same("<li>hello <code>x</code>\n</li>", "<li>hello <code>x</code></li>", false) {
		t.Fatal("trailing layout whitespace changed the tree")
	}
	if !same("<p>a\n  b</p>", "<p>a b</p>", false) {
		t.Fatal("collapsed text whitespace changed the tree")
	}
}

func TestHTMLASTRetainsMeaningfulWhitespaceAndMarkup(t *testing.T) {
	cases := [][2]string{
		{"<p>a <em>b</em></p>", "<p>a<em>b</em></p>"},
		{"<p>a b</p>", "<p>ab</p>"},
		{"<pre>a\n b</pre>", "<pre>a b</pre>"},
		{"<code>a  b</code>", "<code>a b</code>"},
		{"<math-renderer>$a  b$</math-renderer>", "<math-renderer>$a b$</math-renderer>"},
		{"<p>x</p>", "<div>x</div>"},
		{`<a href="/a">x</a>`, `<a href="/b">x</a>`},
	}
	for _, test := range cases {
		if same(test[0], test[1], false) {
			t.Fatalf("meaningful difference collapsed: %q vs %q", test[0], test[1])
		}
	}
}

func TestHTMLASTNormalizesGeneratedIDsOnlyInAttributes(t *testing.T) {
	a := "0123456789abcdef0123456789abcdef"
	b := "abcdef0123456789abcdef0123456789"
	if !same(`<math-renderer data-run-id="`+a+`">x</math-renderer>`,
		`<math-renderer data-run-id="`+b+`">x</math-renderer>`, false) {
		t.Fatal("generated run IDs were not normalized")
	}
	if same("<p>"+a+"</p>", "<p>"+b+"</p>", false) {
		t.Fatal("generated-looking text was normalized")
	}
}

func TestHTMLASTIgnoresSyntaxHighlighting(t *testing.T) {
	highlighted := `<div class="highlight highlight-source-js"><pre class="notranslate"><span class="pl-k">const</span> x</pre></div>`
	plain := `<pre lang="js" class="notranslate"><code class="notranslate">const x
</code></pre>`
	if !same(highlighted, plain, true) {
		t.Fatal("highlighting wrapper was not ignored")
	}
	if same(highlighted, plain, false) {
		t.Fatal("highlighting wrapper should differ without the option")
	}
	if same(highlighted, replaceOnce(plain, "const x", "const y"), true) {
		t.Fatal("code content difference was ignored")
	}
	if same("<pre><code><em>x</em></code></pre>", "<pre><code>x</code></pre>", true) {
		t.Fatal("markup inside code was ignored")
	}
	if same(`<div class="highlight"><pre><a href="/x">x</a></pre></div>`, "<pre><code>x</code></pre>", true) {
		t.Fatal("non-highlighting markup inside a highlight wrapper was ignored")
	}
}

func replaceOnce(value, old, replacement string) string {
	for i := 0; i+len(old) <= len(value); i++ {
		if value[i:i+len(old)] == old {
			return value[:i] + replacement + value[i+len(old):]
		}
	}
	return value
}
