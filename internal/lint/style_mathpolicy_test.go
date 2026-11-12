package lint

import (
	"strings"
	"testing"
)

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

func TestMathBlockSizeLimit(t *testing.T) {
	limit := strings.Repeat("a", maxMathBlockBytes)
	over := strings.Repeat("a", maxMathBlockBytes+1)
	tests := []struct {
		name   string
		source string
		line   int
		column int
	}{
		{name: "display at limit", source: "$$\n" + limit + "\n$$\n"},
		{name: "fenced at limit", source: "```math\n" + limit + "\n```\n"},
		{name: "unclosed display", source: "$$\n" + over + "\n"},
		{name: "indented code", source: "    $$\n" + over + "\n    $$\n"},
		{name: "tab-indented code", source: "\t$$\n" + over + "\n\t$$\n"},
		{name: "non-math fence", source: "```text\n" + over + "\n```\n"},
		{name: "display over limit", source: "$$\n" + over + "\n$$\n", line: 1, column: 1},
		{name: "indented display over limit", source: "  $$\n" + over + "\n  $$\n", line: 1, column: 3},
		{name: "fenced over limit", source: "```math\n" + over + "\n```\n", line: 1, column: 1},
		{name: "indented fenced over limit", source: "  ```math\n" + over + "\n  ```\n", line: 1, column: 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := 0
			for _, violation := range mathPolicyStyleViolations(test.source) {
				if violation.ID != "MATH-011" {
					continue
				}
				count++
				if violation.Line != test.line || violation.Column != test.column {
					t.Errorf("MATH-011 at %d:%d, want %d:%d", violation.Line, violation.Column, test.line, test.column)
				}
			}
			want := 0
			if test.line != 0 {
				want = 1
			}
			if count != want {
				t.Errorf("MATH-011 findings = %d, want %d", count, want)
			}
		})
	}
}

func TestMathLayoutReviewSignals(t *testing.T) {
	for _, source := range []string{
		"$\\sum_i x_i$",
		"$\\prod_{i=0}^{n} x_i$",
		"$\\sqrt{\\frac{1}{2}}$",
		"$\\begin{pmatrix}a & b\\\\ c & d\\end{pmatrix}$",
	} {
		requireStyleViolation(t, mathPolicyStyleViolations(source), "MATH-010")
	}
	for _, violation := range mathPolicyStyleViolations("$x + y$") {
		if violation.ID == "MATH-010" {
			t.Fatalf("plain inline math triggered a layout review: %v", violation)
		}
	}
}

func TestMathFencedBodyPolicy(t *testing.T) {
	requireStyleViolation(t,
		mathPolicyStyleViolations("```math\n\\text{$\\alpha$}\n```\n"),
		"MATH-008")
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
		requireStyleViolation(t, violations, "MATH-007")
		requireStyleViolation(t, violations, "DETAILS-002")
	}
	if got := mathPolicyStyleViolations("```math\nA = 5\n```\n"); len(got) != 0 {
		t.Fatalf("standalone fenced math was flagged: %v", got)
	}
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
