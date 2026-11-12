package lint

import (
	"strings"
	"testing"
)

type styleControl struct {
	id, source string
}

func TestStyleGuideNegativeControls(t *testing.T) {
	mathFence := "```math\nx = 1\n```"
	mermaid := "Diagram description.\n\n```mermaid\nflowchart TD\n    A-->B\n```"
	controls := []styleControl{
		{"BLOCK-001", "# Title\n\n## Section\n\nText"},
		{"BLOCK-002", "- item\n  continued text"},
		{"TEXT-001", "**bold** and *emphasis*"},
		{"TEXT-002", "これは**強調**です。"},
		{"CODE-001", "```html\nliteral <text>\n```"},
		{"CODE-002", "`simple inline code`"},
		{"LINK-001", "Read [the example](https://example.com)."},
		{"LINK-002", "[relative guide](docs/guide.md)"},
		{"LINK-003", "# First section\n\n## Second section"},
		{"MEDIA-001", "![Diagram of the parser pipeline](diagram.png)"},
		{"TABLE-001", "| a | b |\n| --- | --- |\n| 1 | 2 |"},
		{"TABLE-002", "| a | b |\n| --- | --- |\n| plain | text |"},
		{"TABLE-003", "| A |\n| --- |\n| `a\\|b` |"},
		{"TABLE-004", "| a | b |\n| --- | --- |\n| 1 | 2 |"},
		{"TABLE-005", "| a |\n| --- |\n| $a \\vert b$ |"},
		{"STRIKE-001", "~~obsolete~~ and ~~also obsolete~~"},
		{"STRIKE-002", "~~removed~~ text"},
		{"TASK-001", "- [ ] first task\n- [x] second task"},
		{"TASK-003", "- [x] completed task"},
		{"AUTO-001", "The [project](https://example.com) reference is explicit."},
		{"HTML-001", "<strong>safe text</strong>"},
		{"FOOT-001", "Text[^a]\n\n[^a]: note"},
		{"FOOT-002", "Text[^a]\n\n[^a]: note"},
		{"FOOT-003", "Text[^a]\n\n[^a]: note"},
		{"FOOT-004", "Text[^a]\n\n[^a]: note"},
		{"ALERT-001", "> [!NOTE]\n> information"},
		{"ALERT-002", "> [!TIP]\n> helpful hint"},
		{"DETAILS-001", "<details>\n<summary>More</summary>\n\nSafe body.\n\n</details>"},
		{"DETAILS-002", "<details>\n<summary>More</summary>\n\nSafe body.\n\n</details>"},
		{"WIDGET-001", "```geojson\n{\"type\":\"FeatureCollection\",\"features\":[]}\n```"},
		{"WIDGET-002", "```stl\nsolid model\nfacet normal 0 0 1\nendloop\n```"},
		{"SOCIAL-001", "The renderer maps the parsed tree to HTML."},
		{"EXT-001", "Ordinary **Markdown**, without custom directives."},
		{"MATH-001", "$`x`$"},
		{"MATH-002", "$$\nx\n$$"},
		{"MATH-003", mathFence},
		{"MATH-004", "$`\\{x\\}$`"},
		{"MATH-005", mathFence},
		{"MATH-006", "$`x`$ and $`y`$"},
		{"MATH-007", "$`x`$"},
		{"MATH-008", "$`\\text{plain text}`$"},
		{"MATH-009", "$`\\$x`$"},
		{"MATH-010", "$\\sum_i x_i$"},
		{"MATH-011", "$x + y$"},
		{"MERMAID-001", mermaid},
		{"MERMAID-002", mermaid},
		{"MERMAID-003", mermaid},
		{"MERMAID-004", mermaid},
		{"MERMAID-005", mermaid},
		{"MERMAID-006", mermaid},
		{"MERMAID-007", mermaid},
		{"MERMAID-008", mermaid},
		{"COMPAT-001", "profile: github-rich\n\nText."},
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
