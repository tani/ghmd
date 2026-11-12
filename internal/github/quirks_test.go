package github

import (
	"strings"
	"testing"

	"yagfm/internal/cmark-gfm"
)

func renderAdapted(t *testing.T, source string) string {
	t.Helper()
	document, err := cmarkgfm.Parse(source)
	if err != nil {
		t.Fatal(err)
	}
	defer document.Close()
	if !adaptNodes(document.Root(), []byte(source)) {
		t.Fatal("adapt failed")
	}
	return document.HTML()
}

func TestTaskAdaptation(t *testing.T) {
	html := renderAdapted(t, "- [ ] `[x]`\n- [x] yes\n\n`[^a\\]b]`")
	for _, wanted := range []string{
		`<input type="checkbox" disabled="" /> <code>[x]</code>`,
		`<input type="checkbox" checked="" disabled="" /> yes`,
		`<code>[^a\]b]</code>`,
	} {
		if !strings.Contains(html, wanted) {
			t.Fatalf("missing %q in\n%s", wanted, html)
		}
	}
}

func TestNestedTaskMarkers(t *testing.T) {
	html := renderAdapted(t, "日本語\x00\n\n- [x] 一\n  - [ ] 二\n  - [X] 三")
	if strings.Count(html, `<input type="checkbox" checked="" disabled="" />`) != 2 {
		t.Fatalf("checked count wrong:\n%s", html)
	}
	if strings.Count(html, `<input type="checkbox" disabled="" />`) != 1 {
		t.Fatalf("unchecked count wrong:\n%s", html)
	}
}

func TestRecoveredEscapedLabel(t *testing.T) {
	html := renderAdapted(t, "[^a\\]b] and `[^a\\]b]`")
	if !strings.Contains(html, "[^a]b] and <code>[^a\\]b]</code>") {
		t.Fatalf("recovery mismatch:\n%s", html)
	}
}
