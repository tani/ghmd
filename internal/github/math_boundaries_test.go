package github

import (
	"strings"
	"testing"
)

const mathBoundaryRunID = "0123456789abcdef0123456789abcdef"

func mathBoundaryElement(body string, display bool, runID string) string {
	kind, displayStyle := "inline", "inline-block"
	if display {
		kind, displayStyle = "display", "block"
	}
	return `<math-renderer class="js-` + kind + `-math" style="display: ` + displayStyle + `" data-run-id="` + runID + `">` + body + `</math-renderer>`
}

func renderMathBoundaryDocument(t *testing.T, source, runID string) string {
	t.Helper()
	// RenderDocument is the internal entry point used by ghmd.Render. Using it
	// here avoids an import cycle while retaining the fixed run ID assertions.
	html, err := RenderDocument(source, runID, "01234567-89ab-cdef-0123-456789abcdef")
	if err != nil {
		t.Fatal(err)
	}
	return html
}

func TestMathBoundaryBehavior(t *testing.T) {
	math := func(body string) string { return mathBoundaryElement(body, false, mathBoundaryRunID) }
	tests := []struct {
		name, source, want string
	}{
		{
			"second adjacent bare formula stays literal",
			"$a$$b$",
			`<p>` + math("$a$") + `$b$</p>`,
		},
		{
			"bare and protected formulas adjacent to letters",
			"a$x$ ($x$) $x$a $`x`$a a$`x`$",
			`<p>a$x$ (` + math("$x$") + `) $x$a ` + math("$x$") + `a a$<code class="notranslate">x</code>$</p>`,
		},
		{
			"escaped HTML in protected math",
			"$`a <b > c`$",
			`<p>` + math("$a &amp;lt;b &amp;gt; c$") + `</p>`,
		},
		{
			"protected math in nested emphasis",
			"***value $`x`$ here***",
			`<p><em><strong>value ` + math("$x$") + ` here</strong></em></p>`,
		},
		{
			"display delimiters in details stay literal",
			"<details>\n\n$$\nx\n$$\n\n</details>",
			"<details><summary>Details</summary>\n<p>$$<br>\nx<br>\n$$</p>\n</details>",
		},
		{
			"later expression recovered after rejected candidate",
			`$\sqrt{\$4}$, $x = \$$`,
			`<p>$\sqrt{$4}$, ` + math("$x = $") + `$</p>`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := renderMathBoundaryDocument(t, test.source, mathBoundaryRunID); got != test.want {
				t.Fatalf("got:\n%s\nwant:\n%s", got, test.want)
			}
		})
	}
}

func TestMathDisplayBlocksAndFencedMath(t *testing.T) {
	html := renderMathBoundaryDocument(t, "$$\na\n$$\n$$\nb\n$$\n\n```math\nx$y\n```", "00000000000000000000000000000000")
	if got := strings.Count(html, `class="js-display-math"`); got != 3 {
		t.Fatalf("got %d display renderers, want 3:\n%s", got, html)
	}
	if !strings.Contains(html, `$$x\$y$$`) {
		t.Fatalf("fenced math was not escaped:\n%s", html)
	}
}

func TestMathStageRecognizesEmptyDisplayBlock(t *testing.T) {
	ids := RenderIDs{RunID: mathBoundaryRunID, Boundary: "boundary"}
	html, _, err := renderMathStage("<p>$$\n\n$$</p>", ids)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(html, `class="js-display-math"`) {
		t.Fatalf("empty block was not recognized:\n%s", html)
	}
}

func TestMathDetailsSummaryPrecedesFencedText(t *testing.T) {
	html := renderMathBoundaryDocument(t, `<div>
<pre lang="math"><details></details></pre>
</div>`, mathBoundaryRunID)
	if !strings.Contains(html, "$$Details$$") {
		t.Fatalf("details summary missing from fenced math:\n%s", html)
	}
}

func TestMathPreservesAdjacentTextAcrossRemovedComments(t *testing.T) {
	source := "<div>\n$$\nx\n$$<!-- <table> --><input type=\"checkbox\" checked>$`x`$<div>😀</div>[!NOTE]<br><span class=\"task-list-item x x\">$`x`$<!-- <table> -->$`x`$$$\nx\n$$x</span>\n</div>"
	html := renderMathBoundaryDocument(t, source, mathBoundaryRunID)
	if got := strings.Count(html, `class="js-inline-math"`); got != 2 {
		t.Fatalf("got %d inline renderers, want 2:\n%s", got, html)
	}
	if strings.Contains(html, "<!--") {
		t.Fatalf("comment survived rendering:\n%s", html)
	}
}
