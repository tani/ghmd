package lint

import (
	"sort"
	"strings"
	"unicode/utf8"

	"gfm/internal/github"
)

const (
	ruleMathNotRecognized   = "math/not-recognized"
	ruleMathUnclosedDollar  = "math/unclosed-dollar"
	ruleMathFootnoteLiteral = "math/footnote-literal"
	ruleMathDisplayIsolated = "math/display-not-isolated"
)

// mathFindings reports dollar math that GitHub renders literally or rejects.
func mathFindings(source string) []Finding {
	if !strings.ContainsRune(source, '$') {
		return nil
	}

	starts := lineStartOffsets(source)
	s := mathScanner{
		source:    source,
		starts:    starts,
		ignored:   mathIgnoredRegions(source, starts),
		footnotes: mathFootnoteLines(source, starts),
	}
	s.scan()
	return s.findings
}

type mathScanner struct {
	source    string
	starts    []int
	ignored   []bool
	footnotes []bool
	findings  []Finding
}

func (s *mathScanner) scan() {
	for i := 0; i < len(s.source); {
		if !s.dollar(i) {
			i++
			continue
		}
		if s.dollar(i + 1) {
			i = s.display(i)
			continue
		}

		close := s.inlineClose(i + 1)
		if close < 0 {
			if s.unclosedOpening(i) {
				s.add(i, ruleMathUnclosedDollar, "unclosed dollar prevents GitHub from recognizing inline math")
			}
			i++
			continue
		}
		if s.currencyPair(i, close) {
			i = close + 1
			continue
		}
		if !s.inlineBoundaries(i, close) {
			s.add(i, ruleMathNotRecognized, "inline math does not satisfy GitHub's dollar boundaries")
		} else if s.footnotes[s.line(i)] {
			s.add(i, ruleMathFootnoteLiteral, "GitHub renders inline math in footnote definitions as literal text")
		}
		i = close + 1
	}
}

func (s *mathScanner) display(open int) int {
	closeIsolated := s.nextDisplay(open+2, true)
	if mathDelimitersOnOwnLines(s.source, s.starts, open) && closeIsolated >= 0 {
		return closeIsolated + 2
	}
	s.add(open, ruleMathDisplayIsolated, "display math requires isolated opening and closing $$ lines")
	close := s.nextDisplay(open+2, false)
	if close >= 0 {
		return close + 2
	}
	return open + 2
}

func (s *mathScanner) dollar(i int) bool {
	return i >= 0 && i < len(s.source) && s.source[i] == '$' &&
		!s.ignored[i] && !mathEscaped(s.source, i)
}

func (s *mathScanner) inlineClose(from int) int {
	for i := from; i < len(s.source) && s.source[i] != '\n'; i++ {
		if s.dollar(i) {
			return i
		}
	}
	return -1
}

func (s *mathScanner) unclosedOpening(open int) bool {
	if open+1 >= len(s.source) || s.source[open+1] == '\n' {
		return false
	}
	if s.source[open+1] == '$' || s.source[open+1] >= '0' && s.source[open+1] <= '9' {
		return false
	}
	return !github.IsJSWhitespace(runeAt(s.source, open+1))
}

// currencyPair keeps adjacent monetary amounts such as "$5 and $6" from
// being interpreted as one rejected inline-math span.
func (s *mathScanner) currencyPair(open, close int) bool {
	return open+1 < len(s.source) && s.source[open+1] >= '0' && s.source[open+1] <= '9' &&
		close+1 < len(s.source) && s.source[close+1] >= '0' && s.source[close+1] <= '9'
}

func (s *mathScanner) nextDisplay(from int, isolated bool) int {
	for i := from; i+1 < len(s.source); i++ {
		if s.dollar(i) && s.dollar(i+1) &&
			(!isolated || mathDelimitersOnOwnLines(s.source, s.starts, i)) {
			return i
		}
	}
	return -1
}

func (s *mathScanner) line(offset int) int {
	return sort.Search(len(s.starts), func(i int) bool { return s.starts[i] > offset }) - 1
}

func (s *mathScanner) add(offset int, rule, message string) {
	line := s.line(offset)
	s.findings = append(s.findings, Finding{
		Line:     line + 1,
		Column:   utf8.RuneCountInString(s.source[s.starts[line]:offset]) + 1,
		Severity: Warning,
		Rule:     rule,
		Message:  message,
	})
}

func (s *mathScanner) inlineBoundaries(open, close int) bool {
	before := runeBefore(s.source, open)
	if open > 0 && s.ignored[open-1] {
		// An inline HTML tag or code span forms a separate text-node boundary.
		before = 0
	} else if before == '[' && mathLinkLabel(s.source, close) {
		// Link labels are scanned as their own text by the renderer.
		before = 0
	}
	return mathInlineBoundaries(s.source, open, close, before)
}

func mathLinkLabel(source string, close int) bool {
	for i := close + 1; i < len(source) && source[i] != '\n'; i++ {
		if source[i] == ']' {
			return i+1 < len(source) && (source[i+1] == '(' || source[i+1] == '[')
		}
	}
	return false
}

func mathInlineBoundaries(source string, open, close int, before rune) bool {
	if close == open+1 {
		return false
	}
	after := runeAt(source, close+1)
	first := runeAt(source, open+1)
	last := runeBefore(source, close)
	invalidAfter := after >= 'a' && after <= 'z' || after >= 'A' && after <= 'Z' ||
		after >= '0' && after <= '9' || after == '_'
	return (before == 0 || github.IsJSWhitespace(before) || before == '(') && !invalidAfter &&
		!github.IsJSWhitespace(first) &&
		(!github.IsJSWhitespace(last) || after == '$') &&
		!strings.ContainsRune(source[open+1:close], '&')
}

// mathIgnoredRegions marks fenced code, code spans, math-code, comments, and
// HTML tags without changing any source offsets or newlines.
func mathIgnoredRegions(source string, starts []int) []bool {
	ignored := make([]bool, len(source))
	mark := func(from, to int) {
		for i := from; i < to; i++ {
			ignored[i] = true
		}
	}

	for line := 0; line < len(starts); line++ {
		start := starts[line]
		end := len(source)
		if line+1 < len(starts) {
			end = starts[line+1]
		}
		marker, size, ok := mathOpeningFence(source[start:end])
		if !ok {
			continue
		}
		closeLine := line + 1
		for closeLine < len(starts) {
			closeEnd := len(source)
			if closeLine+1 < len(starts) {
				closeEnd = starts[closeLine+1]
			}
			if mathClosingFence(source[starts[closeLine]:closeEnd], marker, size) {
				break
			}
			closeLine++
		}
		if closeLine == len(starts) {
			mark(start, len(source))
			break
		}
		end = len(source)
		if closeLine+1 < len(starts) {
			end = starts[closeLine+1]
		}
		mark(start, end)
		line = closeLine
	}

	for i := 0; i < len(source); {
		if ignored[i] {
			i++
			continue
		}
		switch {
		case strings.HasPrefix(source[i:], "<!--"):
			end := strings.Index(source[i+4:], "-->")
			if end < 0 {
				mark(i, len(source))
				i = len(source)
			} else {
				end += i + 7
				mark(i, end)
				i = end
			}
		case source[i] == '<':
			end := mathHTMLTagEnd(source, i)
			if end < 0 {
				i++
			} else {
				mark(i, end)
				i = end
			}
		case source[i] == '`' && !mathEscaped(source, i):
			size := mathByteRun(source, i, '`')
			end := mathCodeSpanEnd(source, i+size, size, ignored)
			if end < 0 {
				i += size
				continue
			}
			start := i
			if i > 0 && source[i-1] == '$' && !ignored[i-1] && !mathEscaped(source, i-1) &&
				end < len(source) && source[end] == '$' && !ignored[end] {
				start = i - 1
				end++
			}
			mark(start, end)
			i = end
		default:
			i++
		}
	}
	return ignored
}

func mathOpeningFence(line string) (byte, int, bool) {
	i := mathFenceIndent(line)
	if i < 0 || i >= len(line) || line[i] != '`' && line[i] != '~' {
		return 0, 0, false
	}
	marker := line[i]
	size := mathByteRun(line, i, marker)
	if size < 3 || marker == '`' && strings.ContainsRune(line[i+size:], '`') {
		return 0, 0, false
	}
	return marker, size, true
}

func mathClosingFence(line string, marker byte, minimum int) bool {
	i := mathFenceIndent(line)
	if i < 0 || i >= len(line) || line[i] != marker || mathByteRun(line, i, marker) < minimum {
		return false
	}
	i += mathByteRun(line, i, marker)
	return strings.Trim(line[i:], " \t\r\n") == ""
}

func mathFenceIndent(line string) int {
	i := 0
	for i < len(line) && i < 4 && line[i] == ' ' {
		i++
	}
	if i > 3 {
		return -1
	}
	return i
}

func mathByteRun(source string, at int, marker byte) int {
	end := at
	for end < len(source) && source[end] == marker {
		end++
	}
	return end - at
}

func mathCodeSpanEnd(source string, from, size int, ignored []bool) int {
	for i := from; i < len(source); {
		if source[i] != '`' {
			i++
			continue
		}
		end := i + mathByteRun(source, i, '`')
		if end-i == size && !ignored[i] {
			return end
		}
		i = end
	}
	return -1
}

func mathHTMLTagEnd(source string, open int) int {
	i := open + 1
	if i < len(source) && source[i] == '/' {
		i++
	}
	if i >= len(source) || !asciiLetter(source[i]) {
		return -1
	}
	var quote byte
	for ; i < len(source); i++ {
		c := source[i]
		if quote != 0 {
			if c == quote {
				quote = 0
			}
			continue
		}
		if c == '\'' || c == '"' {
			quote = c
		} else if c == '>' {
			return i + 1
		}
	}
	return -1
}

func mathEscaped(source string, at int) bool {
	backslashes := 0
	for i := at - 1; i >= 0 && source[i] == '\\'; i-- {
		backslashes++
	}
	return backslashes%2 == 1
}

func mathDelimitersOnOwnLines(source string, starts []int, at int) bool {
	line := sort.Search(len(starts), func(i int) bool { return starts[i] > at }) - 1
	end := len(source)
	if line+1 < len(starts) {
		end = starts[line+1] - 1
	}
	return strings.Trim(source[starts[line]:at], " \t\r") == "" &&
		strings.Trim(source[at+2:end], " \t\r") == ""
}

func mathFootnoteLines(source string, starts []int) []bool {
	lines := make([]bool, len(starts))
	continuation := false
	for i, start := range starts {
		end := len(source)
		if i+1 < len(starts) {
			end = starts[i+1]
		}
		line := source[start:end]
		if mathFootnoteDefinition(line) {
			lines[i] = true
			continuation = true
			continue
		}
		if continuation && (strings.TrimSpace(line) == "" || line[0] == ' ' || line[0] == '\t') {
			lines[i] = true
			continue
		}
		continuation = false
	}
	return lines
}

func mathFootnoteDefinition(line string) bool {
	i := mathFenceIndent(line)
	if i < 0 || !strings.HasPrefix(line[i:], "[^") {
		return false
	}
	close := strings.Index(line[i+2:], "]:")
	return close >= 1
}
