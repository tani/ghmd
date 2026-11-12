package github

import (
	"strings"
	"testing"
)

func mathStage(t *testing.T, html string) string {
	t.Helper()
	ids := NewRenderIDs("00000000000000000000000000000000", "boundary")
	marked, ok := prepareMathTables(html, ids.RunID, ids.Boundary, "")
	if !ok {
		t.Fatal("prepare tables failed")
	}
	out, ok := mathDOM(marked, ids.RunID, ids.Boundary)
	if !ok {
		t.Fatal("math DOM failed")
	}
	return out
}

func TestInlineMathBoundaries(t *testing.T) {
	html := mathStage(t, "<p>日本語 $x$&nbsp;$y$ z$x$</p>")
	if strings.Count(html, `class="js-inline-math"`) != 2 {
		t.Fatalf("expected two inline matches, got:\n%s", html)
	}
	if !strings.Contains(html, " z$x$") {
		t.Fatalf("literal math missing:\n%s", html)
	}
	if got := mathStage(t, "<p>a$x$</p>"); got != "<p>a$x$</p>" {
		t.Fatalf("adjacent math should stay literal: %s", got)
	}
}
