package lint

import (
	"sort"
	"strings"
	"unicode/utf8"
)

// coreStyleViolations is the entry point for the source-observable rules owned
// by the core style checker. The other style engines append their own rules in
// StyleRules.
func coreStyleViolations(source string) []styleViolation {
	lines := sourceLines(source)
	insideFence, _ := fencedCodeRegions(lines)
	commentsMasked := coreMaskComments(lines, insideFence)
	codeSpans := coreInlineCodeSpans(commentsMasked, insideFence)
	proseMasked := coreMaskInlineCode(commentsMasked, codeSpans)

	var violations []styleViolation
	violations = append(violations, blockStyleViolations(lines, proseMasked, insideFence)...)
	violations = append(violations, textStyleViolations(lines, proseMasked, insideFence)...)
	violations = append(violations, codeStyleViolations(lines, codeSpans)...)
	violations = append(violations, taskStyleViolations(lines, proseMasked, insideFence)...)
	sort.SliceStable(violations, func(i, j int) bool {
		a, b := violations[i], violations[j]
		if a.Line != b.Line {
			return a.Line < b.Line
		}
		if a.Column != b.Column {
			return a.Column < b.Column
		}
		return a.ID < b.ID
	})
	return violations
}

type coreHeading struct {
	level  int
	marker int
}

func coreATXHeading(line string) (coreHeading, bool) {
	at := 0
	for at < len(line) && at < 3 && line[at] == ' ' {
		at++
	}
	marker := at
	for at < len(line) && line[at] == '#' && at-marker < 6 {
		at++
	}
	level := at - marker
	if level == 0 || at < len(line) && line[at] == '#' {
		return coreHeading{}, false
	}
	if at < len(line) && line[at] != ' ' && line[at] != '\t' {
		return coreHeading{}, false
	}
	for at < len(line) && (line[at] == ' ' || line[at] == '\t') {
		at++
	}
	return coreHeading{level: level, marker: marker}, true
}

func blockStyleViolations(lines, masked []string, insideFence []bool) []styleViolation {
	var violations []styleViolation
	violations = append(violations, listContinuationViolations(lines, masked, insideFence)...)
	return violations
}

type coreListMarker struct {
	indent        int
	contentIndent int
}

func coreListItem(line string) (coreListMarker, bool) {
	at := 0
	for at < len(line) && (line[at] == ' ' || line[at] == '\t') {
		at++
	}
	markerStart := at
	if at >= len(line) {
		return coreListMarker{}, false
	}
	if line[at] == '-' || line[at] == '+' || line[at] == '*' {
		at++
	} else {
		digits := at
		for at < len(line) && line[at] >= '0' && line[at] <= '9' && at-digits < 9 {
			at++
		}
		if at == digits || at >= len(line) || line[at] != '.' && line[at] != ')' {
			return coreListMarker{}, false
		}
		at++
	}
	if at == markerStart || at >= len(line) || line[at] != ' ' && line[at] != '\t' {
		return coreListMarker{}, false
	}
	for at < len(line) && (line[at] == ' ' || line[at] == '\t') {
		at++
	}
	indent := indentOf(line[:markerStart])
	contentIndent := indent + 2
	if at < len(line) {
		contentIndent = at
	}
	return coreListMarker{indent: indent, contentIndent: contentIndent}, true
}

func listContinuationViolations(lines, masked []string, insideFence []bool) []styleViolation {
	var violations []styleViolation
	var marker coreListMarker
	active := false
	reported := false
	for i, line := range lines {
		if insideFence[i] {
			active = false
			continue
		}
		if strings.TrimSpace(masked[i]) == "" {
			active = false
			continue
		}
		if item, ok := coreListItem(line); ok {
			marker = item
			active = true
			reported = false
			continue
		}
		if !active {
			continue
		}
		trimmed := strings.TrimLeft(line, " \t")
		if _, heading := coreATXHeading(line); heading || strings.HasPrefix(trimmed, ">") {
			active = false
			continue
		}
		indent := indentOf(line)
		zeroColumn := indent == 0 && marker.indent == 0
		oneOrMoreShort := indent < marker.contentIndent
		if reported || !zeroColumn && !oneOrMoreShort {
			continue
		}
		at := 0
		for at < len(line) && (line[at] == ' ' || line[at] == '\t') {
			at++
		}
		violations = append(violations, styleViolation{
			ID: "BLOCK-001", Line: i + 1, Column: coreColumn(line, at),
			Detail: "indent this list continuation to the item content column; do not use an ambiguous lazy continuation",
		})
		reported = true
	}
	return violations
}

func coreColumn(line string, byteOffset int) int {
	if byteOffset < 0 {
		byteOffset = 0
	}
	if byteOffset > len(line) {
		byteOffset = len(line)
	}
	return utf8.RuneCountInString(line[:byteOffset]) + 1
}

func taskStyleViolations(lines, masked []string, insideFence []bool) []styleViolation {
	var violations []styleViolation
	for i, line := range lines {
		if insideFence[i] {
			continue
		}
		analysis := masked[i]
		bullet, bracket, hasBulletSpace, ok := coreTaskLikeItem(analysis)
		if !ok {
			continue
		}
		var problems []string
		if !hasBulletSpace {
			problems = append(problems, "add a space after the list bullet")
		}
		if bracket+3 < len(analysis) && analysis[bracket+3] != ' ' && analysis[bracket+3] != '\t' {
			problems = append(problems, "add a space after the [x] or [ ] status marker")
		}
		if len(problems) > 0 {
			columnAt := bracket
			if !hasBulletSpace {
				columnAt = bullet
			}
			violations = append(violations, styleViolation{
				ID: "TASK-001", Line: i + 1, Column: coreColumn(line, columnAt),
				Detail: strings.Join(problems, "; "),
			})
		}
	}
	return violations
}

func coreTaskLikeItem(line string) (bullet, bracket int, hasBulletSpace, ok bool) {
	at := 0
	for at < len(line) && (line[at] == ' ' || line[at] == '\t') {
		at++
	}
	bullet = at
	if at < len(line) && (line[at] == '-' || line[at] == '+' || line[at] == '*') {
		at++
	} else {
		digits := at
		for at < len(line) && line[at] >= '0' && line[at] <= '9' && at-digits < 9 {
			at++
		}
		if at == digits || at >= len(line) || line[at] != '.' && line[at] != ')' {
			return 0, 0, false, false
		}
		at++
	}
	if at >= len(line) {
		return 0, 0, false, false
	}
	hasBulletSpace = line[at] == ' ' || line[at] == '\t'
	if hasBulletSpace {
		for at < len(line) && (line[at] == ' ' || line[at] == '\t') {
			at++
		}
	}
	bracket = at
	if bracket+3 > len(line) || line[bracket] != '[' || line[bracket+2] != ']' {
		return 0, 0, false, false
	}
	if line[bracket+1] != ' ' && line[bracket+1] != 'x' && line[bracket+1] != 'X' {
		return 0, 0, false, false
	}
	return bullet, bracket, hasBulletSpace, true
}
