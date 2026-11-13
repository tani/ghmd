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

func TestFootnoteIdentifiers(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
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

}

func TestFootnoteSafeIdentifiersAndSurfaces(t *testing.T) {
	source := "Prefix 日本 Text[^a][^a-2]\n\n[^a]: note\n    continued prose\n[^a-2]: another note"
	for _, id := range []string{"FOOT-001", "FOOT-002", "FOOT-003"} {
		rejectFootnoteRule(t, source, id)
	}
	rejectFootnoteRule(t, "Text[^A]\n\n[^A]: note", "FOOT-001")
}

func TestFootnoteDuplicateDefinitions(t *testing.T) {
	violations := requireFootnoteRule(t, "[^a]: one\n[^a]: duplicate", "FOOT-001")
	foundDuplicate := false
	for _, violation := range violations {
		foundDuplicate = foundDuplicate || strings.Contains(violation.Detail, "duplicate")
	}
	if !foundDuplicate {
		t.Fatalf("no duplicate-definition warning: %#v", violations)
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
	rejectFootnoteRule(t, "Ordinary prose with [!NOTE] but no alert claim.", "ALERT-001")
}
