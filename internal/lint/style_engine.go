package lint

import (
	"strings"
	"unicode/utf8"
)

// styleViolation is a concrete authoring violation recognized from source, as
// distinct from a rendering or parser differential. ID matches RULES.md.
type styleViolation struct {
	ID     string
	Line   int
	Column int
	Detail string
}

func (v styleViolation) finding() Finding {
	return Finding{Line: v.Line, Column: v.Column, Severity: Warning,
		Rule: "style/" + v.ID, Message: v.ID + ": " + v.Detail}
}

// styleCheck is one independent authoring rule group. Each group has only a
// source contract and cannot affect renderer output.
type styleCheck func(string) []styleViolation

var styleChecks = []struct {
	name  string
	check styleCheck
}{
	{"blocks", coreStyleViolations},
	{"tables", tableStrikeStyleViolations},
	{"footnotes", footnoteViolations},
	{"alerts", alertViolations},
	{"math-policy", mathPolicyStyleViolations},
	{"widgets", widgetViolations},
}

// StyleRules evaluates authoring rules with a stable, extensible rule registry.
func StyleRules(source string) []Finding {
	var violations []styleViolation
	for _, group := range styleChecks {
		violations = append(violations, group.check(source)...)
	}
	findings := make([]Finding, 0, len(violations))
	for _, violation := range violations {
		findings = append(findings, violation.finding())
	}
	return findings
}

func sourceLines(source string) []string {
	return strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n")
}

func styleColumn(line, prefix string) int {
	if len(prefix) > len(line) {
		prefix = line
	}
	return utf8.RuneCountInString(prefix) + 1
}

func indentOf(line string) int { return len(line) - len(strings.TrimLeft(line, " \t")) }

func fenceAt(line string) (marker rune, length int, rest string, ok bool) {
	trimmed := strings.TrimLeft(line, " ")
	if len(line)-len(trimmed) > 3 || len(trimmed) < 3 {
		return 0, 0, "", false
	}
	marker = rune(trimmed[0])
	if marker != '`' && marker != '~' {
		return 0, 0, "", false
	}
	for length < len(trimmed) && trimmed[length] == trimmed[0] {
		length++
	}
	if length < 3 {
		return 0, 0, "", false
	}
	return marker, length, strings.TrimSpace(trimmed[length:]), true
}

// fencedCodeRegions returns opening and interior markers for all code fences.
// Opening rows are in both slices; closing rows are only in inside.
func fencedCodeRegions(lines []string) (inside, opens []bool) {
	inside = make([]bool, len(lines))
	opens = make([]bool, len(lines))
	var marker rune
	minimum := 0
	for i, line := range lines {
		m, n, rest, ok := fenceAt(line)
		if marker == 0 {
			if ok {
				inside[i], opens[i] = true, true
				marker, minimum = m, n
			}
			continue
		}
		inside[i] = true
		if ok && m == marker && n >= minimum && rest == "" {
			marker = 0
		}
	}
	return inside, opens
}

// maskLineCode blanks backtick spans and HTML comments, preserving byte offsets.
func maskLineCode(line string) string {
	data := []byte(line)
	if start := strings.Index(string(data), "<!--"); start >= 0 {
		for i := start; i < len(data); i++ {
			data[i] = ' '
		}
	}
	for i := 0; i < len(data); {
		if data[i] != '`' {
			i++
			continue
		}
		run := 1
		for i+run < len(data) && data[i+run] == '`' {
			run++
		}
		close := -1
		for j := i + run; j+run <= len(data); j++ {
			if data[j] != '`' {
				continue
			}
			count := 1
			for j+count < len(data) && data[j+count] == '`' {
				count++
			}
			if count == run {
				close = j
				break
			}
			j += count - 1
		}
		if close >= 0 {
			for j := i; j < close+run; j++ {
				data[j] = ' '
			}
			i = close + run
		} else {
			for j := i; j < i+run; j++ {
				data[j] = ' '
			}
			i += run
		}
	}
	return string(data)
}
