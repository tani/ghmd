package test

import (
	"reflect"
	"strings"
	"testing"

	"ghmd"
	cmarkghmd "ghmd/internal/cmark-gfm"
	"ghmd/internal/verify"
)

func renderForBehaviorTest(t *testing.T, source, runID string) string {
	t.Helper()
	html, err := ghmd.Render(source, ghmd.RenderOptions{RunID: runID})
	if err != nil {
		t.Fatal(err)
	}
	return html
}

func requireHTMLContains(t *testing.T, html string, wanted ...string) {
	t.Helper()
	for _, value := range wanted {
		if !strings.Contains(html, value) {
			t.Fatalf("missing %q in:\n%s", value, html)
		}
	}
}

func requireHTMLNotContains(t *testing.T, html string, rejected ...string) {
	t.Helper()
	for _, value := range rejected {
		if strings.Contains(html, value) {
			t.Fatalf("unexpected %q in:\n%s", value, html)
		}
	}
}

func requireSameHTMLAST(t *testing.T, got, want string) {
	t.Helper()
	gotAST := verify.HTMLAST(got, false)
	wantAST := verify.HTMLAST(want, false)
	if !reflect.DeepEqual(gotAST, wantAST) {
		t.Fatalf("HTML AST mismatch\ngot:  %#v\nwant: %#v", gotAST, wantAST)
	}
}

// The bun:test null-argument and shared-global concurrency checks are omitted:
// Go rejects non-string input at compile time and has no global document object.
// Typed invalid identifier/source-size checks live in render_test.go.

func TestRenderGFMKeepsTaskStatesSeparate(t *testing.T) {
	if got := cmarkghmd.Version(); got != "0.29.0.gfm.13" {
		t.Fatalf("cmark-gfm version = %q", got)
	}
	html := renderForBehaviorTest(t,
		"- [ ] `[x]`\n- [x] done\n\n| A | B |\n|---|---|\n| 1 | 2 |", "")
	requireHTMLContains(t, html,
		`aria-label="Incomplete task"`,
		`aria-label="Completed task"`,
		`<markdown-accessiblity-table><table role="table">`)
}

func TestRenderGitHubAlertsAndExternalLinks(t *testing.T) {
	html := renderForBehaviorTest(t, "> [!NOTE]\n> Message.\n\n[link](https://example.com)", "")
	requireHTMLContains(t, html,
		`class="markdown-alert markdown-alert-note"`,
		"<p>Message.</p>",
		`href="https://example.com" rel="nofollow"`)
}

func TestNativeParserRendersUTF8(t *testing.T) {
	const source = "あ **太字**\n\n[^a]\n\n[^a]: note"
	document, err := cmarkghmd.Parse(source)
	if err != nil {
		t.Fatal(err)
	}
	defer document.Close()
	requireHTMLContains(t, document.HTML(), "<strong>太字</strong>")
	// renderAstXml and source-location AST assertions are intentionally omitted:
	// those native/XML inspection APIs were intentionally removed from the Go rewrite.
}

func TestRenderPreservesSafeHTMLAndRemovesExecutableContent(t *testing.T) {
	html := renderForBehaviorTest(t,
		"<details>\n<summary>More</summary>\n\nContent\n\n</details>\n\n<img src=x onerror=alert(1)>\n\n<a href=\"javascript:alert(1)\">bad</a>", "")
	requireHTMLContains(t, html, "<details>\n<summary>More</summary>\n<p>Content</p>\n</details>")
	requireHTMLNotContains(t, html, "onerror", "javascript:")
}

func TestRenderMathPlaceholdersAndFootnoteText(t *testing.T) {
	html := renderForBehaviorTest(t, "$`x^2`$\n\n```math\nx^2\n```\n\nText[^a]\n\n[^a]: $a$", "")
	requireHTMLContains(t, html,
		`class="js-inline-math"`,
		`class="js-display-math"`,
		">$$x^2$$</math-renderer>",
		"<p>$a$ <a")
}

func TestRenderMultilineFootnoteLabelsAndImageLabels(t *testing.T) {
	html := renderForBehaviorTest(t, "[^a\nb]: # c\n\nxxx[^a\nb].\n\n[^a]: note\n\n[^![image](#)]", "")
	if got := strings.Count(html, "data-footnote-ref="); got != 2 {
		t.Fatalf("data-footnote-ref count = %d in:\n%s", got, html)
	}
	requireHTMLContains(t, html, "<p>[^]</p>")
}

func TestRenderCodeBlocksWithoutSyntaxHighlighting(t *testing.T) {
	html := renderForBehaviorTest(t,
		"```mermaid\nflowchart TD\n A --> B\n```\n\n```json\n{\"type\":\"FeatureCollection\"}\n```", "")
	requireHTMLContains(t, html,
		"<pre lang=\"mermaid\" class=\"notranslate\"><code class=\"notranslate\">flowchart TD\n A --&gt; B\n</code></pre>",
		"<pre lang=\"json\" class=\"notranslate\"><code class=\"notranslate\">{\"type\":\"FeatureCollection\"}\n</code></pre>")
	requireHTMLNotContains(t, html, `class="pl-`)
}

func TestRenderRemoteImagesWithoutProxy(t *testing.T) {
	html := renderForBehaviorTest(t, "![alt](https://example.com/a.png)", "")
	requireHTMLContains(t, html, `<img src="https://example.com/a.png" alt="alt">`)
	requireHTMLNotContains(t, html, "data-canonical-src")
}

func TestRenderLinkedImagesRetainOriginalURLs(t *testing.T) {
	html := renderForBehaviorTest(t, "[![alt](https://example.com/a.png)](https://example.com)", "")
	want := `<p><a href="https://example.com" rel="nofollow"><img src="https://example.com/a.png" alt="alt"></a></p>`
	if html != want {
		t.Fatalf("got %q, want %q", html, want)
	}
}

func TestRenderDeterministicIDAndMalformedID(t *testing.T) {
	const runID = "1234567890abcdef1234567890abcdef"
	html := renderForBehaviorTest(t, "$x$ and [^a]\n\n[^a]: note", runID)
	requireHTMLContains(t, html,
		`data-run-id="`+runID+`"`,
		`user-content-fn-a-`+runID)
	if _, err := ghmd.Render("text", ghmd.RenderOptions{RunID: "bad"}); err == nil || !strings.Contains(err.Error(), "runId must be 32") {
		t.Fatalf("malformed run ID error = %v", err)
	}
}

func TestRenderRawAnchorAndRejectedDestination(t *testing.T) {
	html := renderForBehaviorTest(t, "<a>plain anchor</a> and [unsafe](javascript:alert(1))", "")
	requireHTMLContains(t, html, "<a>plain anchor</a>", " and unsafe")
	requireHTMLNotContains(t, html, "javascript:", "data-gfm-had-href")
}

func TestRenderUnwrapsTemplatesAndKeepsContent(t *testing.T) {
	for _, test := range []struct {
		source string
		want   string
	}{
		{source: "<template><code>x</code></template>", want: `<p><code class="notranslate">x</code></p>`},
		{source: "<template><b>x</b><i>y</i></template>", want: "<p><b>x</b><i>y</i></p>"},
		{source: "<div><template>before <strong>x</strong> after</template></div>", want: "<div>before <strong>x</strong> after</div>"},
	} {
		if got := renderForBehaviorTest(t, test.source, ""); got != test.want {
			t.Fatalf("render(%q) = %q, want %q", test.source, got, test.want)
		}
	}
}

func TestRenderTrimsFinalNewlineForOpenRawHTML(t *testing.T) {
	for _, test := range []struct {
		source string
		want   string
	}{
		{source: "<div>\n<b>hello\n", want: "<div>\n<b>hello</b></div>"},
		{source: "<div>\n<b>hello\n</div>", want: "<div>\n<b>hello\n</b></div>"},
		{source: "<pre>\nfoo\n", want: `<pre class="notranslate">foo</pre>`},
		{source: "<details><summary>x</summary>\nhello\n", want: "<details><summary>x</summary>\nhello</details>"},
	} {
		if got := renderForBehaviorTest(t, test.source, ""); got != test.want {
			t.Fatalf("render(%q) = %q, want %q", test.source, got, test.want)
		}
	}
}

func TestRenderSanitizesTemplatesBeforeUnwrapping(t *testing.T) {
	html := renderForBehaviorTest(t,
		`<template><img src=x onerror=alert(1)><a href="javascript:alert(1)">x</a></template>`, "")
	requireHTMLNotContains(t, html, "onerror", "javascript:")
	requireHTMLContains(t, html, "</a>x</p>")
}

func TestRenderNestedTemplatesAndParsedTableStructure(t *testing.T) {
	if got := renderForBehaviorTest(t, "<template><template>hi <b>x</b></template></template>", ""); got != "<p>hi <b>x</b></p>" {
		t.Fatalf("nested template output = %q", got)
	}
	html := renderForBehaviorTest(t, "<template><table><tr><td>x</td></tr></table></template>", "")
	requireSameHTMLAST(t, html,
		`<p><markdown-accessiblity-table></markdown-accessiblity-table></p><table role="table"><tbody><tr><td>x</td></tr></tbody></table><p></p>`)
}

func TestRenderTableHandlingIsLocalToMathParagraph(t *testing.T) {
	const runID = "0123456789abcdef0123456789abcdef"
	math := `<math-renderer class="js-inline-math" style="display: inline-block" data-run-id="` + runID + `">$x$</math-renderer>`
	table := `<table role="table"><tbody><tr><td>x</td></tr></tbody></table>`

	same := renderForBehaviorTest(t,
		"<template><table><tr><td>x</td></tr></table></template> $x$", runID)
	requireSameHTMLAST(t, same, "<p></p>"+table+" "+math)

	separate := renderForBehaviorTest(t,
		"<template><table><tr><td>x</td></tr></table></template>\n\n$x$", runID)
	requireSameHTMLAST(t, separate,
		"<p><markdown-accessiblity-table></markdown-accessiblity-table></p>"+table+"<p></p>\n<p>"+math+"</p>")

	nested := renderForBehaviorTest(t,
		"<div><template><table><tr><td>$x$</td></tr></table></template></div>", runID)
	requireSameHTMLAST(t, nested,
		"<div><markdown-accessiblity-table><table role=\"table\"><tbody><tr><td>"+math+"</td></tr></tbody></table></markdown-accessiblity-table></div>")

	for _, html := range []string{same, separate, nested} {
		requireHTMLNotContains(t, html, "data-gfm-table-")
	}
}

func TestRenderEmptySafeLinksAndEncodedDangerousURLs(t *testing.T) {
	html := renderForBehaviorTest(t, `<a href="">empty</a>`, "")
	requireHTMLContains(t, html, `<a href="">empty</a>`)
	for _, url := range []string{
		"javascript:alert(1)",
		"jav&#x61;script:alert(1)",
		"javascript&#58;alert(1)",
		"vbscript:alert(1)",
	} {
		t.Run(url, func(t *testing.T) {
			html := renderForBehaviorTest(t, `<a href="`+url+`">bad</a>`, "")
			requireHTMLNotContains(t, html, "href=")
			requireHTMLContains(t, html, "bad")
		})
	}
}

func TestRenderProcessesLinkMarkersAndDestinationsAfterSanitization(t *testing.T) {
	html := renderForBehaviorTest(t,
		`<a href="https://example.com/?a=1&amp;b=2" rel="ugc">web</a> <a href="//example.com">protocol relative</a> <a href="MAILTO:a@example.com">mail</a> <a href="/local" rel="ugc">local</a> <a href="javascript:alert(1)" id="anchor">kept</a> <a href="javascript:alert(1)"><strong>unwrapped</strong></a>`, "")
	requireHTMLContains(t, html,
		`href="https://example.com/?a=1&amp;b=2" rel="nofollow"`,
		`href="//example.com" rel="nofollow"`,
		`href="MAILTO:a@example.com" rel="nofollow"`,
		`href="/local" rel="ugc"`,
		`id="anchor">kept</a>`,
		"<strong>unwrapped</strong>")
	requireHTMLNotContains(t, html, "data-gfm-had-href")
}

func TestRenderInsertsDefaultDetailsSummaryOnlyWhenNeeded(t *testing.T) {
	html := renderForBehaviorTest(t,
		"<details><div><summary>Nested</summary></div>body</details>\n\n<details><summary>Own</summary>body</details>", "")
	requireHTMLContains(t, html,
		"<details><summary>Details</summary><div><summary>Nested</summary></div>body</details>",
		"<details><summary>Own</summary>body</details>")
}

func TestRenderPreservesMathTextBoundariesAcrossTemplates(t *testing.T) {
	if got := renderForBehaviorTest(t, "$<template>x</template>$", ""); got != "<p>$x$</p>" {
		t.Fatalf("template math boundary output = %q", got)
	}
	if got := renderForBehaviorTest(t, "a<template>b<template>c</template></template>d", ""); got != "<p>abcd</p>" {
		t.Fatalf("nested template text output = %q", got)
	}
}

func TestRenderDetailsTemplateTableAndMathTogether(t *testing.T) {
	got := renderForBehaviorTest(t,
		"<details><template><table><tr><td>x</td></tr></table></template> $x$</details>",
		strings.Repeat("a", 32))
	want := "<details><summary>Details</summary><markdown-accessiblity-table><table role=\"table\"><tbody><tr><td>x</td></tr></tbody></table></markdown-accessiblity-table> $x$</details>"
	if got != want {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}
