package lint

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type mathPolicySpan struct {
	start     int
	end       int // exclusive, including the delimiters
	bodyStart int
	bodyEnd   int
	display   bool
	code      bool
}

func mathPolicyMaskComment(line string) string {
	if at := strings.Index(line, "<!--"); at >= 0 {
		return line[:at] + strings.Repeat(" ", len(line)-at)
	}
	return line
}

// mathPolicyMaskHTMLTags prevents dollars in attributes from becoming math;
// the text between HTML tags remains available to the policy checks.
func mathPolicyMaskHTMLTags(line string) string {
	data := []byte(line)
	for i := 0; i < len(data); {
		if data[i] != '<' {
			i++
			continue
		}
		end := mathPolicyHTMLTagEnd(string(data), i)
		if end < 0 {
			i++
			continue
		}
		for j := i; j < end; j++ {
			data[j] = ' '
		}
		i = end
	}
	return string(data)
}

func mathPolicyHTMLTagEnd(line string, at int) int {
	if at >= len(line) || line[at] != '<' {
		return -1
	}
	i := at + 1
	if i < len(line) && line[i] == '/' {
		i++
	}
	if i >= len(line) || !asciiLetter(line[i]) {
		return -1
	}
	for i < len(line) && (asciiLetter(line[i]) || line[i] >= '0' && line[i] <= '9' || line[i] == '-') {
		i++
	}
	if i < len(line) && line[i] != '>' && line[i] != '/' && line[i] != ' ' && line[i] != '\t' {
		return -1
	}
	quote := byte(0)
	for ; i < len(line); i++ {
		if quote != 0 {
			if line[i] == quote {
				quote = 0
			}
			continue
		}
		if line[i] == '\'' || line[i] == '"' {
			quote = line[i]
		} else if line[i] == '>' {
			return i + 1
		}
	}
	return -1
}

func mathPolicyBackslashRun(line string, at int) int {
	start := at
	for start > 0 && line[start-1] == '\\' {
		start--
	}
	return at - start
}

func mathPolicyDollarRun(line string, at int) int {
	end := at
	for end < len(line) && line[end] == '$' {
		end++
	}
	return end - at
}

func mathPolicyDollarSpans(line string) []mathPolicySpan {
	var spans []mathPolicySpan
	for i := 0; i < len(line); {
		if line[i] != '$' || mathPolicyBackslashRun(line, i)%2 != 0 {
			i++
			continue
		}
		size := mathPolicyDollarRun(line, i)
		if size > 2 || i+size >= len(line) {
			i += size
			continue
		}
		close := -1
		for j := i + size; j < len(line); {
			if line[j] != '$' || mathPolicyBackslashRun(line, j)%2 != 0 {
				j++
				continue
			}
			closeSize := mathPolicyDollarRun(line, j)
			if size == 1 {
				// In $x$$y$, the first dollar closes x and the second
				// opens y; it is not a display delimiter.
				close = j
			} else if closeSize == 2 {
				close = j
			}
			if close >= 0 {
				break
			}
			j += closeSize
		}
		if close < 0 {
			i += size
			continue
		}
		span := mathPolicySpan{start: i, end: close + size, bodyStart: i + size, bodyEnd: close, display: size == 2}
		body := line[span.bodyStart:span.bodyEnd]
		first, _ := utf8.DecodeRuneInString(body)
		last, _ := utf8.DecodeLastRuneInString(body)
		after := rune(0)
		if span.end < len(line) {
			after, _ = utf8.DecodeRuneInString(line[span.end:])
		}
		valid := body != "" && !unicode.IsSpace(first) && !unicode.IsSpace(last) &&
			!(unicode.IsLetter(after) || unicode.IsDigit(after) || after == '_') &&
			mathPolicyOpeningBoundary(line, i)
		if unicode.IsDigit(first) && unicode.IsDigit(after) {
			valid = false // adjacent monetary amounts, not a rich-math assertion
		}
		if valid {
			spans = append(spans, span)
			i = span.end
		} else {
			i += size
		}
	}
	return spans
}

func mathPolicyOpeningBoundary(line string, at int) bool {
	if at == 0 {
		return true
	}
	before, _ := utf8.DecodeLastRuneInString(line[:at])
	return unicode.IsSpace(before) || strings.ContainsRune("([{<>:*|_~$", before)
}

func mathPolicyDollarBacktickSpans(line string) []mathPolicySpan {
	var spans []mathPolicySpan
	for i := 0; i+2 < len(line); {
		if line[i] != '$' || line[i+1] != '`' || mathPolicyBackslashRun(line, i)%2 != 0 {
			i++
			continue
		}
		run := 1
		for i+1+run < len(line) && line[i+1+run] == '`' {
			run++
		}
		bodyStart := i + 1 + run
		close := -1
		for j := bodyStart; j+run < len(line); j++ {
			if line[j] != '`' || line[j+run] != '$' {
				continue
			}
			matched := true
			for k := 1; k < run; k++ {
				if line[j+k] != '`' {
					matched = false
					break
				}
			}
			if matched {
				close = j
				break
			}
		}
		if close < 0 {
			i++
			continue
		}
		spans = append(spans, mathPolicySpan{start: i, end: close + run + 1,
			bodyStart: bodyStart, bodyEnd: close, code: true})
		i = close + run + 1
	}
	return spans
}

func mathPolicyAlternateDelimiters(line string) []int {
	var out []int
	for i := 0; i+3 < len(line); i++ {
		if line[i] != '\\' || line[i+1] != '(' && line[i+1] != '[' || mathPolicyBackslashRun(line, i)%2 != 0 {
			continue
		}
		closing := "\\" + string(map[byte]byte{'(': ')', '[': ']'}[line[i+1]])
		if end := strings.Index(line[i+2:], closing); end > 0 {
			out = append(out, i)
			i += 1 + end + len(closing)
		}
	}
	return out
}

func mathPolicyTeXComment(body string) int {
	for i := 0; i < len(body); i++ {
		if body[i] == '%' && mathPolicyBackslashRun(body, i)%2 == 0 {
			return i
		}
	}
	return -1
}

func mathPolicyNestedTextDollars(line string) []int {
	var out []int
	for from := 0; from < len(line); {
		at := strings.Index(line[from:], `\text`)
		if at < 0 {
			break
		}
		at += from
		open := at + len(`\text`)
		for open < len(line) && unicode.IsSpace(rune(line[open])) {
			open++
		}
		if open >= len(line) || line[open] != '{' {
			from = open
			continue
		}
		depth, end := 1, open+1
		for ; end < len(line) && depth > 0; end++ {
			if mathPolicyBackslashRun(line, end)%2 != 0 {
				continue
			}
			if line[end] == '{' {
				depth++
			} else if line[end] == '}' {
				depth--
			}
		}
		if depth != 0 {
			break
		}
		body := line[open+1 : end-1]
		for _, span := range mathPolicyDollarSpans(body) {
			if !span.display {
				out = append(out, at)
				break
			}
		}
		from = end
	}
	return out
}

func mathPolicyDoubledDollarEscapes(line string) []int {
	var out []int
	for i := 1; i < len(line); i++ {
		if line[i] != '$' {
			continue
		}
		run := mathPolicyBackslashRun(line, i)
		if run < 2 || run%2 != 0 {
			continue
		}
		// The characteristic source is $\\$x$: an opening dollar occurs
		// before the doubled escape. Requiring that opener avoids treating a
		// literal path ending in a double backslash as a TeX assertion.
		if strings.Contains(line[:i-run], "$") {
			out = append(out, i-run)
		}
	}
	return out
}

func mathPolicyNestedContext(line string, span mathPolicySpan) bool {
	candidate, depth := -1, 0
	for i := 0; i < span.start; i++ {
		if mathPolicyBackslashRun(line, i)%2 != 0 {
			continue
		}
		switch line[i] {
		case '[':
			candidate, depth = i, depth+1
		case ']':
			if depth > 0 {
				depth--
			}
			if depth == 0 {
				candidate = -1
			}
		}
	}
	if candidate < 0 {
		return false
	}
	depth = 1
	for i := span.end; i < len(line); i++ {
		if mathPolicyBackslashRun(line, i)%2 != 0 {
			continue
		}
		if line[i] == '[' {
			depth++
		} else if line[i] == ']' {
			depth--
			if depth == 0 {
				return i+1 < len(line) && (line[i+1] == '(' || line[i+1] == '[')
			}
		}
	}
	return false
}

func mathPolicyEmphasisContext(line string, span mathPolicySpan) bool {
	for _, marker := range []byte{'*', '_'} {
		open := -1
		for i := span.start - 1; i >= 0; {
			if line[i] != marker || mathPolicyBackslashRun(line, i)%2 != 0 {
				i--
				continue
			}
			start := i
			for start > 0 && line[start-1] == marker {
				start--
			}
			if start+1 < len(line) && !unicode.IsSpace(rune(line[start+1])) {
				open = start
			}
			break
		}
		if open < 0 {
			continue
		}
		for i := span.end; i < len(line); i++ {
			if line[i] == marker && mathPolicyBackslashRun(line, i)%2 == 0 && i > 0 && !unicode.IsSpace(rune(line[i-1])) {
				return true
			}
		}
	}
	return false
}

func mathPolicyTableContext(line string, span mathPolicySpan) bool {
	for i := 0; i < len(line); i++ {
		if i > span.start && i < span.end || line[i] != '|' || mathPolicyBackslashRun(line, i)%2 != 0 {
			continue
		}
		return true
	}
	return false
}

type mathPolicyHTMLTag struct {
	name string
}

func mathPolicyDetailsOpen(tags []mathPolicyHTMLTag) bool {
	for _, tag := range tags {
		if tag.name == "details" {
			return true
		}
	}
	return false
}

func mathPolicyHTMLStacks(lines []string, fenced []bool) [][]mathPolicyHTMLTag {
	stacks := make([][]mathPolicyHTMLTag, len(lines))
	var stack []mathPolicyHTMLTag
	for i, line := range lines {
		stacks[i] = append([]mathPolicyHTMLTag(nil), stack...)
		if fenced[i] {
			continue
		}
		stack = mathPolicyScanHTMLTags(maskLineCode(line), len(line), stack)
	}
	return stacks
}

func mathPolicyHTMLContext(line string, at int, before []mathPolicyHTMLTag) bool {
	stack := append([]mathPolicyHTMLTag(nil), before...)
	stack = mathPolicyScanHTMLTags(maskLineCode(line), at, stack)
	return len(stack) > 0
}

func mathPolicyScanHTMLTags(line string, limit int, stack []mathPolicyHTMLTag) []mathPolicyHTMLTag {
	if limit > len(line) {
		limit = len(line)
	}
	for i := 0; i < limit; {
		end := mathPolicyHTMLTagEnd(line, i)
		if i >= limit || line[i] != '<' || end < 0 || end > limit {
			i++
			continue
		}
		literal := line[i:end]
		closing := strings.HasPrefix(literal, "</")
		start := 1
		if closing {
			start = 2
		}
		finish := start
		for finish < len(literal) && (asciiLetter(literal[finish]) || literal[finish] >= '0' && literal[finish] <= '9' || literal[finish] == '-') {
			finish++
		}
		name := strings.ToLower(literal[start:finish])
		if closing {
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j].name == name {
					stack = stack[:j]
					break
				}
			}
		} else if !mathPolicyVoidHTMLTag(name) && !strings.HasSuffix(strings.TrimSpace(literal), "/>") {
			stack = append(stack, mathPolicyHTMLTag{name: name})
		}
		i = end
	}
	return stack
}

func mathPolicyVoidHTMLTag(name string) bool {
	switch name {
	case "area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr":
		return true
	}
	return false
}
