package lint

import "testing"

func TestMathPolicySourceChecks(t *testing.T) {
	violations := mathPolicyStyleViolations("$x$")
	requireStyleViolation(t, violations, "MATH-002")
	violations = mathPolicyStyleViolations("\\(x\\) and \\[y\\]")
	requireStyleViolation(t, violations, "MATH-003")
	violations = tableStrikeStyleViolations("a~~***nested***~~a")
	requireStyleViolation(t, violations, "STRIKE-002")
	violations = mathPolicyStyleViolations("cost $5 and $6 total")
	if len(violations) != 0 {
		t.Fatalf("currency should remain ordinary prose: %v", violations)
	}
	violations = mathPolicyStyleViolations("[$x$](https://example.com)")
	requireStyleViolation(t, violations, "MATH-007")
}

func TestDetailsMathContext(t *testing.T) {
	violations := mathPolicyStyleViolations("<details><summary>x</summary>\n\n$y$\n")
	requireStyleViolation(t, violations, "DETAILS-002")
	violations = mathPolicyStyleViolations("$$x$$")
	requireStyleViolation(t, violations, "MATH-002")
	violations = mathPolicyStyleViolations(`$\text{$x$}$`)
	requireStyleViolation(t, violations, "MATH-008")
	violations = mathPolicyStyleViolations(`$\\$x$`)
	requireStyleViolation(t, violations, "MATH-009")
}
