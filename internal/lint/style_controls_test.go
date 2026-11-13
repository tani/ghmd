package lint

import (
	"strings"
	"testing"
)

type styleControl struct {
	id, source string
}

func TestRuleNegativeControls(t *testing.T) {
	mathFence := "```math\nx = 1\n```"
	mermaid := "Diagram description.\n\n```mermaid\nflowchart TD\n    A-->B\n```"
	controls := []styleControl{
		{"BLOCK-001", "- item\n  continued text"},
		{"TEXT-001", "**bold** and *emphasis*"},
		{"TEXT-002", "これは**強調**です。"},
		{"CODE-001", "```html\nliteral <text>\n```"},
		{"CODE-002", "`simple inline code`"},
		{"TABLE-001", "| a | b |\n| --- | --- |\n| 1 | 2 |"},
		{"TABLE-002", "| a | b |\n| --- | --- |\n| plain | text |"},
		{"TABLE-003", "| A |\n| --- |\n| `a\\|b` |"},
		{"TABLE-004", "| a |\n| --- |\n| $a \\vert b$ |"},
		{"STRIKE-001", "~~obsolete~~ and ~~also obsolete~~"},
		{"STRIKE-002", "~~removed~~ text"},
		{"TASK-001", "- [ ] first task\n- [x] second task"},
		{"HTML-001", "<strong>safe text</strong>"},
		{"FOOT-001", "Text[^a]\n\n[^a]: note"},
		{"FOOT-002", "Text[^a]\n\n[^a]: note"},
		{"FOOT-003", "Text[^a]\n\n[^a]: note"},
		{"ALERT-001", "> [!NOTE]\n> information"},
		{"DETAILS-001", "<details>\n<summary>More</summary>\n\nSafe body.\n\n</details>"},
		{"DETAILS-002", "<details>\n<summary>More</summary>\n\nSafe body.\n\n</details>"},
		{"WIDGET-001", "```stl\nsolid model\nfacet normal 0 0 1\nendloop\n```"},
		{"EXT-001", "Ordinary **Markdown**, without custom directives."},
		{"MATH-001", mathFence},
		{"MATH-002", "$`\\{x\\}$`"},
		{"MATH-003", mathFence},
		{"MATH-004", "$`x`$ and $`y`$"},
		{"MATH-005", "$`x`$"},
		{"MATH-006", "$`\\text{plain text}`$"},
		{"MATH-007", "$`\\$x`$"},
		{"MERMAID-001", mermaid},
	}
	for _, control := range controls {
		t.Run(control.id, func(t *testing.T) {
			findings, err := Source(control.source)
			if err != nil {
				t.Fatal(err)
			}
			for _, finding := range findings {
				if finding.Rule == "style/"+control.id {
					t.Fatalf("%s fired on compliant input %q: %v", control.id, control.source, findings)
				}
			}
			if strings.Contains(control.source, "```") {
				return
			}
		})
	}
}
