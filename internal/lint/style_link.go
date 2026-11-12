package lint

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	coreBareURLPattern  = regexp.MustCompile(`(?i)(https?://|www\.)[^\s<>\[\]]+`)
	coreBareMailPattern = regexp.MustCompile(`[A-Za-z0-9][A-Za-z0-9._%+\-]*@[A-Za-z0-9][A-Za-z0-9.\-]*\.[A-Za-z]{2,}`)
)

func linkStyleViolations(lines, masked []string, insideFence []bool) []styleViolation {
	var violations []styleViolation
	headings := make(map[string]struct{})
	for i, line := range lines {
		if insideFence[i] {
			continue
		}
		if heading, ok := coreATXHeading(line); ok && heading.level > 0 && heading.normalized != "" {
			if _, duplicate := headings[heading.normalized]; duplicate {
				violations = append(violations, styleViolation{
					ID: "LINK-003", Line: i + 1, Column: coreColumn(line, heading.marker),
					Detail: "duplicate heading " + heading.normalized + "; give each section a unique title for stable anchors",
				})
			} else {
				headings[heading.normalized] = struct{}{}
			}
		}

		for _, destination := range coreMarkdownDestinations(masked[i]) {
			if strings.HasPrefix(destination.text, "../") {
				violations = append(violations, styleViolation{
					ID: "LINK-002", Line: i + 1, Column: coreColumn(line, destination.start),
					Detail: "verify relative path " + destination.text + " against the destination document, or use an explicit absolute URL",
				})
			}
		}

		bareAnalysis := coreMaskExplicitLinks(masked[i])
		urlMatches := coreBareURLPattern.FindAllStringIndex(bareAnalysis, -1)
		for _, match := range urlMatches {
			if !coreBareTokenBoundary(bareAnalysis, match[0]) {
				continue
			}
			token := coreTrimBareToken(bareAnalysis[match[0]:match[1]])
			violations = append(violations, styleViolation{
				ID: "LINK-001", Line: i + 1, Column: coreColumn(line, match[0]),
				Detail: "use an explicit descriptive Markdown link instead of bare " + token,
			})
		}
		for _, match := range coreBareMailPattern.FindAllStringIndex(bareAnalysis, -1) {
			insideURL := false
			for _, url := range urlMatches {
				if match[0] >= url[0] && match[0] < url[1] {
					insideURL = true
					break
				}
			}
			if insideURL || !coreBareTokenBoundary(bareAnalysis, match[0]) {
				continue
			}
			token := coreTrimBareToken(bareAnalysis[match[0]:match[1]])
			violations = append(violations, styleViolation{
				ID: "LINK-001", Line: i + 1, Column: coreColumn(line, match[0]),
				Detail: "use an explicit descriptive email link instead of bare " + token,
			})
		}
	}
	return violations
}

func coreBareTokenBoundary(line string, start int) bool {
	if start == 0 {
		return true
	}
	previous, _ := utf8.DecodeLastRuneInString(line[:start])
	if previous >= 'a' && previous <= 'z' || previous >= 'A' && previous <= 'Z' || previous >= '0' && previous <= '9' {
		return false
	}
	return previous != '@' && previous != '_' && previous != '-'
}

func coreTrimBareToken(token string) string {
	return strings.TrimRight(token, ".,;:!?。．，、！？；：")
}

type coreDestination struct {
	start int
	text  string
}

func coreMarkdownDestinations(line string) []coreDestination {
	var destinations []coreDestination
	for at := 0; at+1 < len(line); at++ {
		if line[at] != ']' || line[at+1] != '(' || coreEscaped(line, at) {
			continue
		}
		start := at + 2
		for start < len(line) && (line[start] == ' ' || line[start] == '\t') {
			start++
		}
		if start < len(line) && line[start] == '<' {
			start++
		}
		end := start
		for end < len(line) && line[end] != ')' && line[end] != '>' && line[end] != ' ' && line[end] != '\t' {
			if line[end] == '\\' && end+1 < len(line) {
				end += 2
				continue
			}
			end++
		}
		if end > start {
			destinations = append(destinations, coreDestination{start: start, text: line[start:end]})
		}
	}
	// Reference definitions are explicit links too. Their destination still
	// needs the same relative-path validation as an inline destination.
	at := 0
	for at < len(line) && at < 4 && line[at] == ' ' {
		at++
	}
	if at < len(line) && line[at] == '[' {
		if close := strings.Index(line[at:], "]:"); close >= 0 {
			start := at + close + 2
			for start < len(line) && (line[start] == ' ' || line[start] == '\t') {
				start++
			}
			if start < len(line) && line[start] == '<' {
				start++
			}
			end := start
			for end < len(line) && line[end] != ' ' && line[end] != '\t' && line[end] != '>' {
				end++
			}
			if end > start {
				destinations = append(destinations, coreDestination{start: start, text: line[start:end]})
			}
		}
	}
	return destinations
}

func coreMaskExplicitLinks(line string) string {
	out := []byte(line)
	// Angle-bracket autolinks and raw HTML attributes are explicit syntax, not
	// extended bare-autolink detection.
	for at := 0; at < len(line); at++ {
		if line[at] != '<' {
			continue
		}
		if close := strings.IndexByte(line[at+1:], '>'); close >= 0 {
			close += at + 2
			coreMaskBytes(out, at, close)
			at = close - 1
		}
	}
	for at := 0; at < len(line); at++ {
		open := at
		if line[at] == '!' && at+1 < len(line) && line[at+1] == '[' {
			open = at + 1
		} else if line[at] != '[' {
			continue
		}
		close := coreUnescapedByte(line, ']', open+1)
		if close < 0 || close+1 >= len(line) {
			continue
		}
		switch line[close+1] {
		case '(':
			if end := coreClosingParen(line, close+2); end >= 0 {
				coreMaskBytes(out, at, end+1)
				at = end
			}
		case '[':
			if end := coreUnescapedByte(line, ']', close+2); end >= 0 {
				coreMaskBytes(out, at, end+1)
				at = end
			}
		}
	}
	at := 0
	for at < len(line) && at < 4 && line[at] == ' ' {
		at++
	}
	if at < len(line) && line[at] == '[' {
		if colon := strings.Index(line[at:], "]:"); colon >= 0 {
			coreMaskBytes(out, at, len(line))
		}
	}
	return string(out)
}

func coreUnescapedByte(line string, wanted byte, start int) int {
	for start < len(line) {
		at := strings.IndexByte(line[start:], wanted)
		if at < 0 {
			return -1
		}
		at += start
		if !coreEscaped(line, at) {
			return at
		}
		start = at + 1
	}
	return -1
}

func coreClosingParen(line string, start int) int {
	depth := 0
	for at := start; at < len(line); at++ {
		if coreEscaped(line, at) {
			continue
		}
		switch line[at] {
		case '(':
			depth++
		case ')':
			if depth == 0 {
				return at
			}
			depth--
		}
	}
	return -1
}
