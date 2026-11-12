package test

import (
	"strings"
	"testing"

	"ghmd"
)

const upstreamExtraRunID = "00000000000000000000000000000000"

func renderUpstreamExtra(t *testing.T, source string) string {
	t.Helper()
	html, err := ghmd.Render(source, ghmd.RenderOptions{RunID: upstreamExtraRunID})
	if err != nil {
		t.Fatal(err)
	}
	return html
}

func TestUpstreamTaskStatesPreserveLiteralCode(t *testing.T) {
	html := renderUpstreamExtra(t, "- [ ] `[x]`\n- [x] yes\n\n`[^a\\]b]`")
	if got := strings.Count(html, `type="checkbox"`); got != 2 {
		t.Fatalf("got %d checkboxes, want 2:\n%s", got, html)
	}
	if got := strings.Count(html, `checked=""`); got != 1 {
		t.Fatalf("got %d checked tasks, want 1:\n%s", got, html)
	}
	for _, wanted := range []string{`<code class="notranslate">[x]</code>`, `<code class="notranslate">[^a\]b]</code>`} {
		if !strings.Contains(html, wanted) {
			t.Fatalf("missing %q in\n%s", wanted, html)
		}
	}
}

func TestUpstreamNestedTasksAfterMultibyteAndNUL(t *testing.T) {
	html := renderUpstreamExtra(t, "日本語\x00\n\n- [x] 一\n  - [ ] 二\n  - [X] 三")
	if got := strings.Count(html, `checked=""`); got != 2 {
		t.Fatalf("got %d checked tasks, want 2:\n%s", got, html)
	}
	checkboxes := strings.Count(html, `type="checkbox"`)
	if checkboxes != 3 {
		t.Fatalf("got %d checkboxes, want 3:\n%s", checkboxes, html)
	}
	if unchecked := checkboxes - strings.Count(html, `checked=""`); unchecked != 1 {
		t.Fatalf("got %d unchecked tasks, want 1:\n%s", unchecked, html)
	}
}

func TestUpstreamRecoversEscapedUnresolvedLabel(t *testing.T) {
	html := renderUpstreamExtra(t, "[^a\\]b] and `[^a\\]b]`")
	if !strings.Contains(html, `[^a]b] and <code class="notranslate">[^a\]b]</code>`) {
		t.Fatalf("label recovery mismatch:\n%s", html)
	}
}

func TestUpstreamRetainsFootnoteLabelBeforeImage(t *testing.T) {
	html := renderUpstreamExtra(t, "[^a![x](#)]\n\n[^a]: note")
	for _, wanted := range []string{"data-footnote-ref", "<p>note "} {
		if !strings.Contains(html, wanted) {
			t.Fatalf("missing %q in\n%s", wanted, html)
		}
	}
}

func TestUpstreamRecoversMultilineFootnoteLabel(t *testing.T) {
	html := renderUpstreamExtra(t, "日本語\x00 [^] [^a\nb]\n\n[^a]: note")
	for _, wanted := range []string{
		"日本語� [^] <sup",
		`href="#user-content-fn-a-` + upstreamExtraRunID + `"`,
		"<p>note ",
	} {
		if !strings.Contains(html, wanted) {
			t.Fatalf("missing %q in\n%s", wanted, html)
		}
	}
}

func TestUpstreamInlineMathAtUnicodeWhitespace(t *testing.T) {
	html := renderUpstreamExtra(t, "日本語 $x$\u00a0$y$ z$x$")
	if got := strings.Count(html, `class="js-inline-math"`); got != 2 {
		t.Fatalf("got %d inline renderers, want 2:\n%s", got, html)
	}
	if !strings.Contains(html, " z$x$") {
		t.Fatalf("letter-adjacent formula did not stay literal:\n%s", html)
	}
	if got := renderUpstreamExtra(t, "a$x$"); got != "<p>a$x$</p>" {
		t.Fatalf("letter-adjacent formula changed:\n%s", got)
	}
}

func TestUpstreamDeepListInTemplateRemainsInert(t *testing.T) {
	deep := strings.Repeat("<ul><li>", 12) + "child" + strings.Repeat("</li></ul>", 12)
	input := `<template>` + deep + `</template>`
	if got := finish(input); got != input {
		t.Fatalf("inert list changed:\ngot:\n%s\nwant:\n%s", got, input)
	}
}

func TestUpstreamImageAndCheckboxCorrectionsSkipTemplates(t *testing.T) {
	input := `<ul><li><input type="checkbox" checked="" disabled=""> <p><img src="a?x=1&amp;y=2"></p></li></ul><template><img src="b"><input type="checkbox"></template>`
	output := finish(input)
	for _, wanted := range []string{
		`<ul class="contains-task-list"><li class="task-list-item">
<p><input type="checkbox" id="" disabled="" class="task-list-item-checkbox" aria-label="Completed task" checked=""> `,
		`<a target="_blank" rel="noopener noreferrer" href="a?x=1&amp;y=2"><img src="a?x=1&amp;y=2" style="max-width: 100%;"></a>`,
		`<template><img src="b"><input type="checkbox"></template>`,
	} {
		if !strings.Contains(output, wanted) {
			t.Fatalf("missing %q in\n%s", wanted, output)
		}
	}
}

func TestUpstreamMultilineFootnoteSyntaxInCodeStaysLiteral(t *testing.T) {
	html := renderUpstreamExtra(t, "[^] and `[^a\nb]`\n\n[^a]: note")
	if !strings.Contains(html, `<code class="notranslate">[^a b]</code>`) {
		t.Fatalf("code was treated as footnote syntax:\n%s", html)
	}
	if strings.Contains(html, "data-footnote-ref") {
		t.Fatalf("code produced a footnote reference:\n%s", html)
	}
}
