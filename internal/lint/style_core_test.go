package lint

import (
	"fmt"
	"strings"
	"testing"
)

type coreStyleCase struct {
	source string
	line   int
	column int
}

func TestCoreStyleRules(t *testing.T) {
	tests := []struct {
		id   string
		bad  []coreStyleCase
		safe []string
	}{
		{
			id: "BLOCK-001",
			bad: []coreStyleCase{
				{source: "- item\ncontinuation at zero", line: 2, column: 1},
				{source: " - item\n  continuation", line: 2, column: 3},
				{source: "- one\n two", line: 2, column: 2},
			},
			safe: []string{
				"- one\n  two",
				"- one\n\nseparate paragraph",
				"- outer\n  - nested\n    continuation",
			},
		},
		{
			id: "TEXT-001",
			bad: []coreStyleCase{
				{source: "**outer **inner** outer**", line: 1, column: 9},
				{source: "*outer **inner* outer**", line: 1},
				{source: "** bold **", line: 1, column: 1},
				{source: "*bold *", line: 1, column: 7},
			},
			safe: []string{
				"**bold** and *emphasis*",
				"`**literal ** delimiters`",
				"**one** and **two**",
				"2 * 3 * 4",
			},
		},
		{
			id:  "TEXT-002",
			bad: []coreStyleCase{{source: "これは**強調です。**です", line: 1, column: 11}},
			safe: []string{
				"これは**強調**です。",
				"These are **strong。** sentences",
				"文。**強調**です",
			},
		},
		{
			id: "CODE-001",
			bad: []coreStyleCase{
				{source: "```markdown\n```html\nliteral\n```", line: 2, column: 1},
				{source: "````markdown\n`````HTML\n<script>alert(1)</script>\n```\n````", line: 2, column: 1},
			},
			safe: []string{
				"```html\n<p>x</p>\n```",
				"````markdown\n```html\nx\n```\n````",
				"~~~js\nconst answer = 42\n~~~",
			},
		},
		{
			id: "CODE-002",
			bad: []coreStyleCase{
				{source: "`code with ` embedded backtick`", line: 1},
				{source: "`code with ` embedded backtick", line: 1},
				{source: "`exact\nline`", line: 1},
			},
			safe: []string{
				"`code` and `more`",
				"``code with ` backtick``",
				"```text\n`literal\n```",
			},
		},
		{
			id: "TASK-001",
			bad: []coreStyleCase{
				{source: "-[ ] missing space", line: 1, column: 1},
				{source: "- [x]no space", line: 1, column: 3},
				{source: "-[x] text", line: 1, column: 1},
			},
			safe: []string{
				"- [ ] todo\n- [x] done\n- [X] uppercase is accepted by GitHub",
				"-[todo] is not a checkbox",
				"`-[ ] literal task syntax`",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			for n, bad := range test.bad {
				t.Run(fmt.Sprintf("violation_%d", n+1), func(t *testing.T) {
					violations := coreStyleViolations(bad.source)
					matched := false
					for _, violation := range violations {
						if violation.ID != test.id || bad.line != 0 && violation.Line != bad.line || bad.column != 0 && violation.Column != bad.column {
							continue
						}
						matched = true
						if strings.TrimSpace(violation.Detail) == "" {
							t.Errorf("empty detail: %#v", violation)
						}
						if got := violation.finding().Message; !strings.HasPrefix(got, test.id+": ") {
							t.Errorf("message lacks actionable ID and detail: %q", got)
						}
					}
					if !matched {
						t.Fatalf("missing %s at %d:%d in %#v", test.id, bad.line, bad.column, violations)
					}
				})
			}
			for n, source := range test.safe {
				t.Run(fmt.Sprintf("safe_%d", n+1), func(t *testing.T) {
					for _, violation := range coreStyleViolations(source) {
						if violation.ID == test.id {
							t.Fatalf("unexpected safe-source finding: %#v", violation)
						}
					}
				})
			}
		})
	}
}
