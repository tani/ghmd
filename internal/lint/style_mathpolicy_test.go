package lint

import "testing"

func TestMathPolicySourceChecks(t *testing.T) {
	violations := mathPolicyStyleViolations("\\(x\\) and \\[y\\]")
	requireStyleViolation(t, violations, "MATH-001")
	violations = tableStrikeStyleViolations("a~~***nested***~~a")
	requireStyleViolation(t, violations, "STRIKE-002")
	violations = mathPolicyStyleViolations("cost $5 and $6 total")
	if len(violations) != 0 {
		t.Fatalf("currency should remain ordinary prose: %v", violations)
	}
	violations = mathPolicyStyleViolations("[$x$](https://example.com)")
	requireStyleViolation(t, violations, "MATH-005")
}

func TestMathFencedBodyPolicy(t *testing.T) {
	requireStyleViolation(t,
		mathPolicyStyleViolations("```math\n\\text{$\\alpha$}\n```\n"),
		"MATH-006")
	if got := mathPolicyStyleViolations("```math\n\\text{plain text}\n```\n"); len(got) != 0 {
		t.Fatalf("plain text fenced math was flagged: %v", got)
	}
}

func TestMathBlockHTMLContext(t *testing.T) {
	for _, source := range []string{
		"<details>\n<summary>More</summary>\n\n```math\nA = 5\n```\n\n</details>\n",
		"<details>\n<summary>More</summary>\n\n$$\nA = 5\n$$\n\n</details>\n",
	} {
		violations := mathPolicyStyleViolations(source)
		requireStyleViolation(t, violations, "MATH-005")
		requireStyleViolation(t, violations, "DETAILS-002")
	}
	if got := mathPolicyStyleViolations("```math\nA = 5\n```\n"); len(got) != 0 {
		t.Fatalf("standalone fenced math was flagged: %v", got)
	}
}

func TestDetailsMathContext(t *testing.T) {
	violations := mathPolicyStyleViolations("<details><summary>x</summary>\n\n$y$\n")
	requireStyleViolation(t, violations, "DETAILS-002")
	violations = mathPolicyStyleViolations(`$\text{$x$}$`)
	requireStyleViolation(t, violations, "MATH-006")
	violations = mathPolicyStyleViolations(`$\\$x$`)
	requireStyleViolation(t, violations, "MATH-007")
}
