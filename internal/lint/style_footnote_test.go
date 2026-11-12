package lint

import (
	"strings"
	"testing"
)

func requireFootnoteRule(t *testing.T, source, id string) []styleViolation {
	t.Helper()
	all := append(footnoteViolations(source), alertViolations(source)...)
	var matches []styleViolation
	for _, violation := range all {
		if violation.ID == id {
			matches = append(matches, violation)
		}
	}
	if len(matches) == 0 {
		t.Fatalf("%s not reported for:\n%s\nall findings: %#v", id, source, all)
	}
	return matches
}

func rejectFootnoteRule(t *testing.T, source, id string) {
	t.Helper()
	all := append(footnoteViolations(source), alertViolations(source)...)
	for _, violation := range all {
		if violation.ID == id {
			t.Fatalf("unexpected %s finding %#v for:\n%s", id, violation, source)
		}
	}
}

func TestFootnoteIdentifiersAndCase(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{name: "uppercase reference", source: "Text[^A]"},
		{name: "space in definition", source: "[^bad id]: note"},
		{name: "colon", source: "Text[^a:b]"},
		{name: "uppercase and colon", source: "Text[^A:B]"},
		{name: "escaped closing bracket", source: `Text[^a\]]`},
		{name: "Unicode", source: "Text[^日本語]"},
		{name: "empty", source: "Text[^]"},
		{name: "leading hyphen", source: "Text[^-bad]"},
		{name: "underscore", source: "[^bad_id]: note"},
		{name: "multiline reference", source: "Text[^a\nb]"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			violations := requireFootnoteRule(t, test.source, "FOOT-001")
			if violations[0].finding().Severity != Warning {
				t.Fatalf("severity = %v, want warning", violations[0].finding().Severity)
			}
		})
	}

	violations := requireFootnoteRule(t, "Text[^A]\n\n[^a]: note", "FOOT-001")
	foundCaseWarning := false
	for _, violation := range violations {
		if strings.Contains(violation.Detail, "differ in case") && strings.Contains(violation.Detail, "[^A]") && strings.Contains(violation.Detail, "[^a]") {
			foundCaseWarning = true
		}
	}
	if !foundCaseWarning {
		t.Fatalf("no actionable exact-case mismatch warning: %#v", violations)
	}
}

func TestFootnoteSafeIdentifiersAndSurfaces(t *testing.T) {
	source := "Prefix 日本 Text[^a][^a-2]\n\n[^a]: note\n    continued prose\n[^a-2]: another note"
	for _, id := range []string{"FOOT-001", "FOOT-002", "FOOT-003", "FOOT-004"} {
		rejectFootnoteRule(t, source, id)
	}
	rejectFootnoteRule(t, "Text[^a]\n\n[^a]: note", "FOOT-004")
	rejectFootnoteRule(t, "A wiki can be discussed without using footnote syntax.", "FOOT-004")
	rejectFootnoteRule(t, "wiki documentation\n\n```md\n[^bad id]: note\n```", "FOOT-004")
}

func TestFootnoteDuplicateDefinitionsAndEndPlacement(t *testing.T) {
	violations := requireFootnoteRule(t, "[^a]: one\n[^a]: duplicate", "FOOT-001")
	foundDuplicate := false
	for _, violation := range violations {
		foundDuplicate = foundDuplicate || strings.Contains(violation.Detail, "duplicate")
	}
	if !foundDuplicate {
		t.Fatalf("no duplicate-definition warning: %#v", violations)
	}

	source := "Text[^a]\n\n[^a]: note\n\nLater non-indented prose."
	violations = requireFootnoteRule(t, source, "FOOT-001")
	foundPlacement := false
	for _, violation := range violations {
		if violation.Line == 3 && strings.Contains(violation.Detail, "end of the document") {
			foundPlacement = true
		}
	}
	if !foundPlacement {
		t.Fatalf("definition placement was not diagnosed at its definition: %#v", violations)
	}
	requireFootnoteRule(t, "[^a]: note\n\n## Later appendix", "FOOT-001")

	safe := "[^a]: one\n    continuation\n\n[^b]: two\n    final continuation"
	for _, violation := range footnoteViolations(safe) {
		if violation.ID == "FOOT-001" && strings.Contains(violation.Detail, "end of the document") {
			t.Fatalf("indented continuation or later definition caused a placement warning: %#v", violation)
		}
	}
}

func TestFootnoteStructuralLabelRules(t *testing.T) {
	source := "[text[^a]](https://example.com)"
	violations := requireFootnoteRule(t, source, "FOOT-002")
	if violations[0].Line != 1 || violations[0].Column != 6 {
		t.Fatalf("link-label finding location = %d:%d, want 1:6", violations[0].Line, violations[0].Column)
	}
	rejectFootnoteRule(t, source, "FOOT-001")

	source = `[^![image](x)]`
	requireFootnoteRule(t, source, "FOOT-002")
	rejectFootnoteRule(t, source, "FOOT-001")

	source = "[^a\nb]: note"
	requireFootnoteRule(t, source, "FOOT-001")
	violations = requireFootnoteRule(t, source, "FOOT-002")
	if violations[0].Line != 1 || !strings.Contains(violations[0].Detail, "physical newline") {
		t.Fatalf("multiline definition-label finding = %#v", violations)
	}

	for _, safe := range []string{
		"Text[^a]\n\n[^a]: note",
		"[ordinary label](https://example.com)",
		"![ordinary image](image.png)",
		"Escaped opening bracket \\[^a] is prose punctuation",
	} {
		rejectFootnoteRule(t, safe, "FOOT-002")
	}
}

func TestFootnoteMathInDefinitionAndContinuation(t *testing.T) {
	for _, source := range []string{
		"[^a]: note $x$",
		"[^a]: note\n    continued with $y$",
		"[^a]: $`x`$",
		"[^a]: note\n    $$x^2$$",
	} {
		violations := requireFootnoteRule(t, source, "FOOT-003")
		if violations[0].finding().Rule != "style/FOOT-003" || !strings.Contains(violations[0].Detail, "move mathematical expression") {
			t.Fatalf("unexpected math finding: %#v", violations[0])
		}
	}
	for _, safe := range []string{
		"[^a]: a simple definition",
		"[^a]: literal math `$x$`",
		"[^a]: note\n\n$outside$",
	} {
		rejectFootnoteRule(t, safe, "FOOT-003")
	}
}

func TestFootnoteCodeRegionsAreIgnored(t *testing.T) {
	source := "Inline `[^A]` and `[text[^bad]](url)`.\n\n```md\n[^bad id]: note $x$\n> [!DANGER]\n```\n\nText[^a]\n\n[^a]: note"
	for _, id := range []string{"FOOT-001", "FOOT-002", "FOOT-003", "ALERT-001"} {
		rejectFootnoteRule(t, source, id)
	}
}

func TestFootnoteWikiSurface(t *testing.T) {
	for _, source := range []string{
		"Footnote assumed to work in a wiki.",
		"wiki: footnote",
		"wiki: documentation\n\nText[^a]\n\n[^a]: note",
		"This wiki uses references.\n\nText[^a]",
	} {
		violations := requireFootnoteRule(t, source, "FOOT-004")
		if !strings.Contains(violations[0].Detail, "numbered references section") {
			t.Fatalf("wiki finding lacks an actionable fallback: %#v", violations[0])
		}
	}
	rejectFootnoteRule(t, "Text[^a]\n\n[^a]: note", "FOOT-004")
	rejectFootnoteRule(t, "The word wiki in literal code is ignored: `wiki: footnote`", "FOOT-004")
}

func TestAlertMarkers(t *testing.T) {
	for _, source := range []string{
		"> [!DANGER]\n> message",
		"> [!note]\n> message",
		"> [!NoTe]\n> message",
		"> [!])\n> message",
	} {
		violations := requireFootnoteRule(t, source, "ALERT-001")
		if !strings.Contains(violations[0].Detail, "use [!NOTE]") && !strings.Contains(violations[0].Detail, "use [!NOTE],") {
			t.Fatalf("alert finding lacks an allowed replacement: %#v", violations[0])
		}
	}

	violations := requireFootnoteRule(t, "> > [!NOTE]\n> > Nested.", "ALERT-001")
	if violations[0].Column != 5 {
		t.Fatalf("nested alert column = %d, want 5", violations[0].Column)
	}
	violations = requireFootnoteRule(t, "> [!NOTE]\n> body\n> [!IMPORTANT]", "ALERT-001")
	if violations[0].Line != 3 {
		t.Fatalf("second alert finding line = %d, want 3", violations[0].Line)
	}
	requireFootnoteRule(t, "> introductory text\n> [!TIP]", "ALERT-001")
	requireFootnoteRule(t, "> [!NOTE] message belongs below", "ALERT-001")
}

func TestSafeStandaloneAlerts(t *testing.T) {
	source := "> [!NOTE]\n> Message.\n\n> [!TIP]\n> Message.\n\n> [!IMPORTANT]\n> Message.\n\n> [!WARNING]\n> Message.\n\n> [!CAUTION]\n> Message."
	rejectFootnoteRule(t, source, "ALERT-001")
	rejectFootnoteRule(t, source, "ALERT-002")
	rejectFootnoteRule(t, "Ordinary prose with [!NOTE] but no alert claim.", "ALERT-001")
}

func TestPortableAlertFallback(t *testing.T) {
	for _, source := range []string{
		"profile: portable-gfm\n\n> [!NOTE]\n> message",
		"portable profile assumes the alert works here",
		"portable [!NOTE]",
		"> [!WARNING] portable",
	} {
		violations := requireFootnoteRule(t, source, "ALERT-002")
		if !strings.Contains(violations[0].Detail, "bold textual label") || !strings.Contains(violations[0].Detail, "**Warning:**") && !strings.Contains(violations[0].Detail, "**Note:**") {
			t.Fatalf("portable alert lacks a bold-label fallback: %#v", violations[0])
		}
	}
	for _, safe := range []string{
		"> [!NOTE]\n> ordinary alert on a rich surface",
		"profile: portable-gfm\n\nNo alert is used.",
		"profile: portable-gfm\n\n```md\n> [!NOTE]\n```",
		"portable prose with a literal `[!NOTE]` code span",
	} {
		rejectFootnoteRule(t, safe, "ALERT-002")
	}
}
