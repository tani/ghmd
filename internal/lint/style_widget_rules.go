package lint

import (
	"regexp"
	"strings"
)

func widgetViolations(source string) []styleViolation {
	lines := sourceLines(source)
	fenced, opens := fencedCodeRegions(lines)
	var out []styleViolation
	for i, line := range lines {
		if fenced[i] && !opens[i] {
			continue
		}
		if opens[i] {
			info := strings.Fields(strings.TrimSpace(line))
			language := ""
			if len(info) > 0 {
				language = strings.TrimPrefix(strings.TrimLeft(info[0], "`~"), "```")
			}
			switch {
			case language == "Mermaid":
				out = append(out, styleViolation{ID: "MERMAID-001", Line: i + 1, Column: 1, Detail: "use a lowercase mermaid fence"})
			case language == "mermaid":
				if i+1 >= len(lines) || !validMermaidType(strings.TrimSpace(lines[i+1])) {
					out = append(out, styleViolation{ID: "MERMAID-001", Line: i + 1, Column: 1, Detail: "start a mermaid diagram with a valid diagram type"})
				}
				for j := i + 1; j < len(lines) && fenced[j] && !opens[j]; j++ {
					body := lines[j]
					if strings.TrimSpace(body) == "```" {
						break
					}
					if strings.Contains(body, "sankey-beta") {
						out = append(out, styleViolation{ID: "MERMAID-002", Line: j + 1, Column: 1, Detail: "gate optional/newer Mermaid diagram types on a successful capability test"})
					}
					if match := regexp.MustCompile(`(?:^|\s)end\s*(?:-->|$)`).FindStringIndex(body); match != nil {
						out = append(out, styleViolation{ID: "MERMAID-003", Line: j + 1, Column: match[0] + 1, Detail: "do not use the reserved lowercase identifier end; rename it (for example End)"})
					}
					if strings.Contains(body, "[ラベル]") || strings.Contains(body, "[a b]") {
						out = append(out, styleViolation{ID: "MERMAID-003", Line: j + 1, Column: 1, Detail: "quote Mermaid labels that contain Unicode or punctuation"})
					}
					if strings.Contains(body, "<b>") || strings.Contains(body, "click ") || strings.Contains(body, "callback") {
						out = append(out, styleViolation{ID: "MERMAID-004", Line: j + 1, Column: 1, Detail: "do not require HTML labels, click directives, or callbacks for diagram meaning"})
					}
					if strings.Contains(body, "%%{init:") || strings.Contains(body, "theme") {
						out = append(out, styleViolation{ID: "MERMAID-005", Line: j + 1, Column: 1, Detail: "avoid initialization/theme directives and undocumented configuration tolerance"})
					}
				}
			case language == "plantuml":
				out = append(out, styleViolation{ID: "EXT-001", Line: i + 1, Column: 1, Detail: "PlantUML is not supported on GitHub; use a descriptive link or literal code unless verified by a custom profile"})
			case language == "stl":
				body := strings.Join(lines[i+1:], "\n")
				if strings.Index(body, "```") >= 0 {
					body = body[:strings.Index(body, "```")]
				}
				trimmed := strings.TrimSpace(body)
				if !strings.HasPrefix(trimmed, "solid ") && !strings.HasPrefix(trimmed, "facet normal") {
					out = append(out, styleViolation{ID: "WIDGET-001", Line: i + 1, Column: 1, Detail: "validate the ASCII STL structure and link large models rather than embedding invalid data"})
				}
			}
			continue
		}
		if fenced[i] {
			continue
		}
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "==") {
			out = append(out, styleViolation{ID: "EXT-001", Line: i + 1, Column: 1, Detail: "GitHub does not support ==highlight== without a custom profile; use ordinary Markdown"})
		}
		if strings.Contains(trimmed, "[[") && strings.Contains(trimmed, "]]") {
			out = append(out, styleViolation{ID: "EXT-001", Line: i + 1, Column: 1, Detail: "do not assume wiki-link or other host-extension support"})
		}
		if strings.Contains(trimmed, ":::") {
			out = append(out, styleViolation{ID: "EXT-001", Line: i + 1, Column: 1, Detail: "colon-fenced directives are not GitHub Markdown; use normal sections or links"})
		}
	}
	return out
}

func validMermaidType(line string) bool {
	first := strings.Fields(line)
	if len(first) == 0 {
		return true // blank line is permitted before a type; evaluate the first nonblank instead
	}
	return regexp.MustCompile(`^(?:flowchart|graph|sequenceDiagram|classDiagram|stateDiagram|stateDiagram-v2|erDiagram|pie|info)$`).MatchString(first[0])
}
