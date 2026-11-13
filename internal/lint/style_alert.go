package lint

import "strings"

type alertSourceMarker struct {
	name         string
	markerOffset int
	afterMarker  int
	depth        int
}

func alertViolations(source string) []styleViolation {
	lines := sourceLines(source)
	fenced, _ := fencedCodeRegions(lines)
	commentsMasked := coreMaskComments(lines, fenced)
	codeSpans := coreInlineCodeSpans(commentsMasked, fenced)
	maskedLines := coreMaskInlineCode(commentsMasked, codeSpans)

	var out []styleViolation
	quoteOpen, quoteHasContent, quoteHasAlert := false, false, false
	for i, line := range lines {
		masked := maskedLines[i]
		depth, contentOffset := alertQuoteContent(masked)
		if depth == 0 {
			quoteOpen, quoteHasContent, quoteHasAlert = false, false, false
			continue
		}
		if !quoteOpen {
			quoteOpen, quoteHasContent, quoteHasAlert = true, false, false
		}
		marker, ok := parseAlertSourceMarker(line)
		if ok {
			add := func(detail string) {
				out = append(out, styleViolation{
					ID: "ALERT-001", Line: i + 1, Column: coreColumn(line, marker.markerOffset), Detail: detail,
				})
			}
			upper := strings.ToUpper(marker.name)
			switch {
			case alertSupportedName(marker.name):
			case alertSupportedName(upper):
				add("alert marker [!" + marker.name + "] has invalid capitalization; use [!" + upper + "]")
			default:
				add("unsupported alert marker [!" + marker.name + "]; use [!NOTE], [!TIP], [!IMPORTANT], [!WARNING], or [!CAUTION]")
			}
			if marker.depth > 1 {
				add("do not nest alerts inside another blockquote or alert; use one top-level blockquote with a single alert marker")
			}
			if quoteHasContent || quoteHasAlert {
				add("place an alert marker on the blockquote's first content line; do not insert a second or nested alert in the same quote")
			}
			if strings.TrimSpace(masked[marker.afterMarker:]) != "" {
				add("move the alert message to subsequent quoted lines; leave the alert marker alone on its line")
			}
			quoteHasAlert = true
		}
		if strings.TrimSpace(masked[contentOffset:]) != "" {
			quoteHasContent = true
		}
	}

	// These collapsible-section source checks belonged to the former combined
	// footnote/alert rule file. Keep their behavior while ownership is split.
	for i, line := range lines {
		if fenced[i] {
			continue
		}
		if start := strings.Index(line, "<summary>"); start >= 0 {
			if end := strings.Index(line[start+len("<summary>"):], "</summary>"); end >= 0 {
				if after := strings.TrimSpace(line[start+len("<summary>")+end+len("</summary>"):]); after != "" {
					out = append(out, styleViolation{ID: "DETAILS-001", Line: i + 1, Column: start + 1, Detail: "separate summary and body with a blank line"})
				}
			}
			if i+1 < len(lines) && strings.TrimSpace(lines[i+1]) != "" && !strings.HasPrefix(strings.TrimSpace(lines[i+1]), "<") {
				out = append(out, styleViolation{ID: "DETAILS-001", Line: i + 1, Column: start + 1, Detail: "leave a blank line before the collapsed section's Markdown body"})
			}
		}
	}
	return out
}

func alertSupportedName(name string) bool {
	switch name {
	case "NOTE", "TIP", "IMPORTANT", "WARNING", "CAUTION":
		return true
	default:
		return false
	}
}

func alertQuoteContent(line string) (depth, contentOffset int) {
	at := 0
	for at < len(line) && at < 3 && line[at] == ' ' {
		at++
	}
	for at < len(line) && line[at] == '>' {
		depth++
		at++
		if at < len(line) && (line[at] == ' ' || line[at] == '\t') {
			at++
		}
	}
	return depth, at
}

func parseAlertSourceMarker(line string) (alertSourceMarker, bool) {
	depth, content := alertQuoteContent(line)
	if depth == 0 || content+2 > len(line) || line[content:content+2] != "[!" {
		return alertSourceMarker{}, false
	}
	relativeClose := strings.IndexByte(line[content+2:], ']')
	if relativeClose < 0 {
		return alertSourceMarker{}, false
	}
	close := content + 2 + relativeClose
	return alertSourceMarker{
		name: line[content+2 : close], markerOffset: content, afterMarker: close + 1,
		depth: depth,
	}, true
}
