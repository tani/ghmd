package lint

import (
	"testing"
)

func requireStyleViolation(t *testing.T, violations []styleViolation, id string) {
	t.Helper()
	for _, violation := range violations {
		if violation.ID == id {
			return
		}
	}
	t.Fatalf("missing %s in %v", id, violations)
}

func TestFootnoteIdentifierPolicy(t *testing.T) {
	violations := footnoteViolations("[^Bad ID]\n\n[^Bad ID]: note\n")
	requireStyleViolation(t, violations, "FOOT-001")
	if got := footnoteViolations("Text[^a]\n\n[^a]: note\n"); len(got) != 0 {
		t.Fatalf("valid footnote was flagged: %v", got)
	}
}

func TestNestedAlertPolicy(t *testing.T) {
	violations := alertViolations("> [!DANGER]\n> message\n")
	requireStyleViolation(t, violations, "ALERT-001")
	violations = alertViolations("> > [!NOTE]\n> > nested\n")
	requireStyleViolation(t, violations, "ALERT-001")
	if got := alertViolations("> [!NOTE]\n> message\n"); len(got) != 0 {
		t.Fatalf("valid alert was flagged: %v", got)
	}
}

func TestDetailsSummaryCorollary(t *testing.T) {
	violations := alertViolations("<details><summary>*em*</summary>\n\ntext\n")
	requireStyleViolation(t, violations, "DETAILS-001")
}

func TestMathPolicyPreservesCurrencyAndEscapes(t *testing.T) {
	violations := mathPolicyStyleViolations("cost $5 and $6 plus \\$x\\$")
	for _, violation := range violations {
		if violation.ID == "MATH-002" || violation.ID == "MATH-009" {
			t.Fatalf("currency or legal escape was flagged: %v", violation)
		}
	}
	requireStyleViolation(t, mathPolicyStyleViolations("$x$"), "MATH-002")
	requireStyleViolation(t, mathPolicyStyleViolations(`$\text{$x$}$`), "MATH-008")
	requireStyleViolation(t, mathPolicyStyleViolations(`$\\$x$`), "MATH-009")
	requireStyleViolation(t, mathPolicyStyleViolations(`$\displaystyle\sum x$`), "MATH-010")
}

func TestWidgetAndSocialPolicies(t *testing.T) {
	widget := widgetViolations("```Mermaid\nflowchart TD\nA-->B\n```\n")
	requireStyleViolation(t, widget, "MERMAID-001")
	widget = widgetViolations("@name #123 :smile:\n")
	requireStyleViolation(t, widget, "SOCIAL-001")
	widget = widgetViolations("==highlight== [[wiki]] ::: warning\n")
	if len(widget) < 3 {
		t.Fatalf("expected all extension violations: %v", widget)
	}
}
