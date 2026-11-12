package github

import "strings"

// utf8At decodes the same potentially malformed UTF-8 sequences as the C
// scanner. In particular, continuation bytes are checked but overlong
// sequences and surrogate code points are not rejected.
func utf8At(s string, i int) rune {
	if i < 0 || i >= len(s) {
		return 0
	}
	c := s[i]
	if c < 0x80 {
		return rune(c)
	}
	n := 1
	if c >= 0xf0 && c <= 0xf4 {
		n = 4
	} else if c >= 0xe0 && c <= 0xef {
		n = 3
	} else if c >= 0xc2 && c <= 0xdf {
		n = 2
	}
	if n == 1 || i+n > len(s) {
		return '\ufffd'
	}
	cp := rune(c & (byte(0x7f) >> n))
	for j := 1; j < n; j++ {
		if s[i+j]&0xc0 != 0x80 {
			return '\ufffd'
		}
		cp = cp<<6 | rune(s[i+j]&0x3f)
	}
	return cp
}

func utf8Before(s string, i int) rune {
	if i <= 0 {
		return 0
	}
	if i > len(s) {
		i = len(s)
	}
	start := i - 1
	for start > 0 && i-start < 4 && s[start]&0xc0 == 0x80 {
		start--
	}
	return utf8At(s[:i], start)
}

// inlineMath locates GitHub-style inline math. Each adjacent pair in the
// returned slice is the start and end byte offset of one match, including its
// dollar delimiters.
func inlineMath(s string) []uint32 {
	if uint64(len(s)) > uint64(^uint32(0)) {
		return nil
	}
	matches := make([]uint32, 0)
	for i := 0; i < len(s); {
		if s[i] != '$' || i+1 >= len(s) || s[i+1] == '$' || s[i+1] == '\n' {
			i++
			continue
		}
		end := i + 1
		for end < len(s) && s[end] != '$' && s[end] != '\n' {
			end++
		}
		if end == len(s) || s[end] != '$' || end == i+1 {
			i++
			continue
		}
		before := utf8Before(s, i)
		after := byte(0)
		if end+1 < len(s) {
			after = s[end+1]
		}
		first := utf8At(s, i+1)
		last := utf8Before(s, end)
		invalidAfter := after >= 'a' && after <= 'z' || after >= 'A' && after <= 'Z' ||
			after >= '0' && after <= '9' || after == '_'
		if i != 0 && !IsJSWhitespace(before) && before != '(' || invalidAfter ||
			IsJSWhitespace(first) || IsJSWhitespace(last) && after != '$' ||
			strings.IndexByte(s[i+1:end], '&') >= 0 {
			i++
			continue
		}
		matches = append(matches, uint32(i), uint32(end+1))
		i = end + 1
	}
	return matches
}

// displayMath returns consecutive display blocks only when they cover the
// entire text.
func displayMath(s string) []uint32 {
	if uint64(len(s)) > uint64(^uint32(0)) {
		return nil
	}
	blocks := make([]uint32, 0)
	start := 0
	for start < len(s) {
		if len(s)-start < 6 || s[start:start+3] != "$$\n" {
			return blocks[:0]
		}
		end := start + 3
		for end+2 < len(s) && !(s[end] == '\n' && s[end+1] == '$' && s[end+2] == '$') {
			end++
		}
		if end+2 >= len(s) {
			return blocks[:0]
		}
		end += 3
		blocks = append(blocks, uint32(start), uint32(end))
		if end == len(s) {
			return blocks
		}
		if s[end] != '\n' {
			return blocks[:0]
		}
		start = end + 1
		if start == len(s) {
			return blocks[:0]
		}
	}
	return blocks
}

// mathPutLines preserves the renderer's extra escaping of
// backslash-newline sequences.
func mathPutLines(out *Buffer, s string) {
	for i := range len(s) {
		if s[i] == '\\' && i+1 < len(s) && s[i+1] == '\n' {
			out.Putc('\\')
		}
		out.Putc(s[i])
	}
}

// emphasisDisplay reconstructs display math containing Markdown emphasis with
// blank lines around its body.
func emphasisDisplay(s string) string {
	if len(s) < 6 || s[:3] != "$$\n" || s[len(s)-3:] != "\n$$" {
		return ""
	}
	var out Buffer
	out.Put("$$\n\n")
	out.Put(s[3 : len(s)-3])
	out.Put("\n\n$$")
	return out.String()
}

func fencedMath(s string) string {
	if len(s) != 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	var out Buffer
	out.Put("$$")
	for i := range len(s) {
		if s[i] == '$' && (i == 0 || s[i-1] != '\\') {
			out.Putc('\\')
		}
		if s[i] == '\\' && i+1 < len(s) && s[i+1] == '\n' {
			out.Putc('\\')
		}
		out.Putc(s[i])
	}
	out.Put("$$")
	return out.String()
}

// protectedMath returns false when the code element is not a protected math
// expression.
func protectedMath(before, after, code string, emphasis bool) (string, bool) {
	if len(before) == 0 || len(after) == 0 || before[len(before)-1] != '$' || after[0] != '$' {
		return "", false
	}
	if len(before) > 1 && (before[len(before)-2] == 'a' || before[len(before)-2] == 'b') {
		return "", false
	}
	if emphasis && len(after) == 1 {
		return "", false
	}
	var out Buffer
	out.Putc('$')
	for i := range len(code) {
		switch code[i] {
		case '&':
			out.Put("&amp;")
		case '<':
			out.Put("&lt;")
		case '>':
			out.Put("&gt;")
		default:
			out.Putc(code[i])
		}
	}
	out.Putc('$')
	return out.String(), true
}
