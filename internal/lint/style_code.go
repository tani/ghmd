package lint

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func codeStyleViolations(lines []string, spans []coreInlineSpan) []styleViolation {
	var violations []styleViolation
	violations = append(violations, fenceStyleViolations(lines)...)
	for _, span := range spans {
		open := span.open
		if !span.matched {
			violations = append(violations, styleViolation{
				ID: "CODE-002", Line: open.line + 1, Column: coreColumn(lines[open.line], open.start),
				Detail: "choose matching inline-code backticks; use a longer delimiter when the code contains a backtick",
			})
			continue
		}
		if span.close.line != open.line {
			violations = append(violations, styleViolation{
				ID: "CODE-002", Line: open.line + 1, Column: coreColumn(lines[open.line], open.start),
				Detail: "inline code crosses a physical newline; use a fenced block when exact line breaks matter",
			})
			continue
		}
		for _, embedded := range span.embedded {
			if embedded.length >= open.length {
				violations = append(violations, styleViolation{
					ID: "CODE-002", Line: embedded.line + 1, Column: coreColumn(lines[embedded.line], embedded.start),
					Detail: "inline code contains an embedded backtick run; delimit it with a longer backtick run",
				})
				break
			}
		}
		if open.length == 1 && span.close.start > open.end {
			content := lines[open.line][open.end:span.close.start]
			suffix := strings.TrimSpace(lines[open.line][span.close.end:])
			last, _ := utf8.DecodeLastRuneInString(content)
			if content != "" && unicode.IsSpace(last) && suffix != "" {
				violations = append(violations, styleViolation{
					ID: "CODE-002", Line: span.close.line + 1, Column: coreColumn(lines[open.line], span.close.start),
					Detail: "the single backtick closes before trailing prose; use a longer delimiter for code containing a backtick",
				})
			}
		}
	}
	return violations
}

type coreFence struct {
	char       byte
	length     int
	indent     int
	info       string
	infoOffset int
}

func coreParseFence(line string) (coreFence, bool) {
	at := 0
	for at < len(line) && at < 4 && line[at] == ' ' {
		at++
	}
	if at > 3 || at >= len(line) || line[at] != '`' && line[at] != '~' {
		return coreFence{}, false
	}
	char := line[at]
	end := at
	for end < len(line) && line[end] == char {
		end++
	}
	if end-at < 3 {
		return coreFence{}, false
	}
	info := strings.TrimSpace(line[end:])
	if char == '`' && strings.ContainsRune(info, '`') {
		return coreFence{}, false
	}
	infoOffset := end
	for infoOffset < len(line) && (line[infoOffset] == ' ' || line[infoOffset] == '\t') {
		infoOffset++
	}
	return coreFence{char: char, length: end - at, indent: at, info: info, infoOffset: infoOffset}, true
}

func coreClosesFence(line string, opener coreFence) bool {
	candidate, ok := coreParseFence(line)
	return ok && candidate.char == opener.char && candidate.length >= opener.length && candidate.info == ""
}

func fenceStyleViolations(lines []string) []styleViolation {
	var violations []styleViolation
	var opener coreFence
	open := false
	for i, line := range lines {
		candidate, isFence := coreParseFence(line)
		if open {
			if coreClosesFence(line, opener) {
				open = false
				continue
			}
			// A would-be closing run with an info string remains literal code.
			// Its presence proves that the outer run is not safely longer than
			// the nested example's delimiters.
			if isFence && candidate.char == opener.char && candidate.length >= opener.length && candidate.info != "" {
				detail := "use an outer fence longer than the embedded " + line[candidate.indent:candidate.indent+candidate.length] + " fence"
				if fields := strings.Fields(candidate.info); len(fields) > 0 && strings.ToLower(fields[0]) != fields[0] {
					detail += " and use lowercase language identifier " + strings.ToLower(fields[0])
				}
				violations = append(violations, styleViolation{
					ID: "CODE-001", Line: i + 1, Column: coreColumn(line, candidate.indent),
					Detail: detail,
				})
			}
			continue
		}
		if !isFence {
			continue
		}
		opener = candidate
		open = true
		fields := strings.Fields(candidate.info)
		if len(fields) == 0 {
			continue
		}
		language := fields[0]
		lower := strings.ToLower(language)
		if lower != language {
			violations = append(violations, styleViolation{
				ID: "CODE-001", Line: i + 1, Column: coreColumn(line, candidate.infoOffset),
				Detail: "use lowercase language identifier " + lower + " instead of " + language,
			})
		}
	}
	return violations
}

func coreMaskComments(lines []string, insideFence []bool) []string {
	masked := make([]string, len(lines))
	inComment := false
	for i, line := range lines {
		if insideFence[i] {
			masked[i] = strings.Repeat(" ", len(line))
			continue
		}
		masked[i] = coreMaskCommentLine(line, &inComment)
	}
	return masked
}

func coreMaskCommentLine(line string, inComment *bool) string {
	out := []byte(line)
	for at := 0; at < len(line); {
		if *inComment {
			relativeEnd := strings.Index(line[at:], "-->")
			if relativeEnd < 0 {
				coreMaskBytes(out, at, len(line))
				return string(out)
			}
			end := at + relativeEnd + 3
			coreMaskBytes(out, at, end)
			*inComment = false
			at = end
			continue
		}
		relativeStart := strings.Index(line[at:], "<!--")
		if relativeStart < 0 {
			break
		}
		start := at + relativeStart
		relativeEnd := strings.Index(line[start+4:], "-->")
		if relativeEnd < 0 {
			coreMaskBytes(out, start, len(line))
			*inComment = true
			break
		}
		end := start + 4 + relativeEnd + 3
		coreMaskBytes(out, start, end)
		at = end
	}
	return string(out)
}

func coreMaskBytes(line []byte, start, end int) {
	for i := start; i < end && i < len(line); i++ {
		line[i] = ' '
	}
}

type coreBacktickRun struct {
	line       int
	start, end int
	length     int
	consumed   bool
}

type coreInlineSpan struct {
	open, close coreBacktickRun
	matched     bool
	embedded    []coreBacktickRun
}

func coreInlineCodeSpans(lines []string, insideFence []bool) []coreInlineSpan {
	var runs []coreBacktickRun
	for i, line := range lines {
		if insideFence[i] {
			continue
		}
		for at := 0; at < len(line); {
			if line[at] != '`' {
				at++
				continue
			}
			end := at + 1
			for end < len(line) && line[end] == '`' {
				end++
			}
			runs = append(runs, coreBacktickRun{line: i, start: at, end: end, length: end - at})
			at = end
		}
	}
	var spans []coreInlineSpan
	for i := range runs {
		if runs[i].consumed {
			continue
		}
		closeAt := -1
		for j := i + 1; j < len(runs); j++ {
			if runs[j].consumed {
				continue
			}
			if runs[j].length != runs[i].length {
				continue
			}
			if coreBlankBetween(lines, runs[i].line, runs[j].line) {
				break
			}
			closeAt = j
			break
		}
		if closeAt < 0 {
			runs[i].consumed = true
			spans = append(spans, coreInlineSpan{open: runs[i]})
			continue
		}
		span := coreInlineSpan{open: runs[i], close: runs[closeAt], matched: true}
		for j := i; j <= closeAt; j++ {
			runs[j].consumed = true
			if j > i && j < closeAt {
				span.embedded = append(span.embedded, runs[j])
			}
		}
		spans = append(spans, span)
	}
	return spans
}

func coreBlankBetween(lines []string, from, to int) bool {
	for i := from + 1; i < to; i++ {
		if strings.TrimSpace(lines[i]) == "" {
			return true
		}
	}
	return false
}

func coreMaskInlineCode(lines []string, spans []coreInlineSpan) []string {
	masked := make([][]byte, len(lines))
	for i, line := range lines {
		masked[i] = []byte(line)
	}
	for _, span := range spans {
		if !span.matched {
			coreMaskBytes(masked[span.open.line], span.open.start, len(masked[span.open.line]))
			continue
		}
		if span.open.line == span.close.line {
			coreMaskBytes(masked[span.open.line], span.open.start, span.close.end)
			continue
		}
		coreMaskBytes(masked[span.open.line], span.open.start, len(masked[span.open.line]))
		for i := span.open.line + 1; i < span.close.line; i++ {
			coreMaskBytes(masked[i], 0, len(masked[i]))
		}
		coreMaskBytes(masked[span.close.line], 0, span.close.end)
	}
	result := make([]string, len(lines))
	for i := range masked {
		result[i] = string(masked[i])
	}
	return result
}
