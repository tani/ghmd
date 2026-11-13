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
	for i, line := range lines {
		if insideFence[i] {
			continue
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
