package lint

import (
	"fmt"
	"testing"
)

var tableStrikeRuleIDs = []string{
	"TABLE-001", "TABLE-002", "TABLE-003", "TABLE-004", "TABLE-005",
	"STRIKE-001", "STRIKE-002",
}

func tableStrikeTestCounts(t *testing.T, source string, want map[string]int) {
	t.Helper()
	violations := tableStrikeStyleViolations(source)
	got := make(map[string]int)
	for _, violation := range violations {
		got[violation.ID]++
		if violation.Detail == "" || violation.Line < 1 || violation.Column < 1 {
			t.Errorf("invalid violation location or detail: %#v", violation)
		}
	}
	for _, id := range tableStrikeRuleIDs {
		if got[id] != want[id] {
			t.Errorf("%s count = %d, want %d\nsource:\n%s\nviolations: %#v", id, got[id], want[id], source, violations)
		}
	}
}

func TestTableCellCountsAndDelimiters(t *testing.T) {
	tableStrikeTestCounts(t, "| A | B |\n| --- | --- |\n| 1 |\n| 1 | 2 | 3 |\n", map[string]int{
		"TABLE-001": 2, // short body row, then excess body row
	})
	tableStrikeTestCounts(t, "| A | B |\n| -- |\n", map[string]int{
		"TABLE-001": 2, // wrong cell count and a two-hyphen delimiter cell
	})
	tableStrikeTestCounts(t, "A | B\n--- | ---\n1 | 2\n", map[string]int{
		"TABLE-001": 3, // leading/trailing pipes on header, delimiter, and body
	})
}

func TestTableEscapedPipesPreserveCellCounts(t *testing.T) {
	tableStrikeTestCounts(t, "| A | B |\n| --- | --- |\n| left \\| right | `a\\|b` |\n", map[string]int{})
	tableStrikeTestCounts(t, "Use A | B in prose, not in a table.\nwith --- | prose in a paragraph.\n", map[string]int{})
}

func TestTableLiteralPipesInCodeAndBackslashRuns(t *testing.T) {
	tableStrikeTestCounts(t, "| A |\n| --- |\n| `a|b` |\n| `a\\|b` |\n", map[string]int{
		"TABLE-003": 1,
	})
	tableStrikeTestCounts(t, "| A |\n| --- |\n| `a\\\\|b` |\n| C \\\\| D |\n| E \\\\\\\\| F |\n", map[string]int{
		"TABLE-003": 3,
	})

	violations := tableStrikeStyleViolations("| A |\n| --- |\n| `界|x` |\n")
	found := false
	for _, violation := range violations {
		if violation.ID == "TABLE-003" {
			found = true
			if violation.Column != 5 {
				t.Errorf("TABLE-003 Unicode column = %d, want 5: %#v", violation.Column, violation)
			}
		}
	}
	if !found {
		t.Error("TABLE-003 was not reported for a literal pipe after a CJK rune")
	}
}

func TestTableComplexCells(t *testing.T) {
	tableStrikeTestCounts(t, "| A |\n| --- |\n| - list item |\n| 1. numbered item |\n| ```go |\n| $$x$$ |\n", map[string]int{
		"TABLE-002": 4,
	})
	tableStrikeTestCounts(t, "| A |\n| --- |\n| a - b |\n| `code` |\n| ```inline code``` |\n| $x$ |\n", map[string]int{})
}

func TestTableCosmeticPadding(t *testing.T) {
	tableStrikeTestCounts(t, "|   a   |  b |\n| --- | --- |\n|   1   |  2 |\n", map[string]int{
		"TABLE-004": 2,
	})
	tableStrikeTestCounts(t, "| 日本語 | 😀 |\n| --- | --- |\n| 値 | 絵文字 |\n", map[string]int{})
	tableStrikeTestCounts(t, "| A | B |\n| --- | --- |\n|    | 内容 |\n", map[string]int{})
}

func TestTableMathLiteralPipe(t *testing.T) {
	tableStrikeTestCounts(t, "| A |\n| --- |\n| $a | b$ |\n| $a \\vert b$ |\n| $a \\mid b$ |\n| $a \\| b$ |\n", map[string]int{
		"TABLE-005": 1,
	})
	tableStrikeTestCounts(t, "| A |\n| --- |\n| $$a|b$$ |\n| `$a|b$` |\n", map[string]int{
		"TABLE-002": 1,
		"TABLE-003": 1,
		"TABLE-005": 0,
	})
	tableStrikeTestCounts(t, "$a | b$\n\n$$a | b$$\n", map[string]int{})
	tableStrikeTestCounts(t, "| A |\n| -- |\n| $a|b$ |\n", map[string]int{
		"TABLE-001": 1,
		"TABLE-005": 0,
	})
}

func TestStrikeTildeRunLengths(t *testing.T) {
	tableStrikeTestCounts(t, "~single~ and ~~~~long~~~~ and ~~~three~~~\n~~double~~\n", map[string]int{
		"STRIKE-001": 3,
	})
	tableStrikeTestCounts(t, "2~3, ~home, bitwise ~value, and an ordinary ~ are not paired strikeout.\n~ not strike ~\n`~code~` and `~~~~code~~~~`\n<!-- ~comment~ -->\n", map[string]int{})
	tableStrikeTestCounts(t, "~~~\n~fenced code~\n~~~\n", map[string]int{})
	tableStrikeTestCounts(t, "~~~~long~~~~\n~~after~~\n", map[string]int{
		"STRIKE-001": 1,
	})
}

func TestStrikeLineAndLetterBoundaries(t *testing.T) {
	tableStrikeTestCounts(t, "~~across\nlines~~\n", map[string]int{
		"STRIKE-002": 1,
	})
	tableStrikeTestCounts(t, "a~~joined~~b\n", map[string]int{
		"STRIKE-002": 1,
	})
	tableStrikeTestCounts(t, "(~~punctuated~~). A ~~separate~~ correction.\n", map[string]int{})
}

func TestStrikeOnlyCrossedEmphasisOverlaps(t *testing.T) {
	for _, source := range []string{
		"~~$x$~~\n",
		"~~$a*b*c$~~\n",
		"~~*x*~~\n",
		"~~**strong**~~\n",
		"**~~wrapped~~**\n",
	} {
		t.Run(fmt.Sprintf("valid_%q", source), func(t *testing.T) {
			tableStrikeTestCounts(t, source, map[string]int{})
		})
	}
	for _, source := range []string{
		"**~~crossed**~~\n",
		"~~**crossed~~**\n",
	} {
		t.Run(fmt.Sprintf("overlap_%q", source), func(t *testing.T) {
			tableStrikeTestCounts(t, source, map[string]int{"STRIKE-002": 1})
		})
	}
}
