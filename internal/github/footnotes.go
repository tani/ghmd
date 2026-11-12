package github

import (
	"bytes"
	"sort"

	"ghmd/internal/cmark-gfm"
)

type recoverySpan struct {
	start, end int
}

type recoveryEdit struct {
	start, end int
	text       string
}

type labelRecovery struct {
	source []byte
	lines  []int
	edits  []recoveryEdit
}

func (r *labelRecovery) location(node cmarkghmd.Node) (recoverySpan, bool) {
	a, b := node.StartLine(), node.StartColumn()
	c, d := node.EndLine(), node.EndColumn()
	if a < 1 || b < 1 || c < 1 || d < 1 || a > len(r.lines) || c > len(r.lines) {
		return recoverySpan{}, false
	}
	start, end := r.lines[a-1], r.lines[c-1]
	aEnd, cEnd := len(r.source), len(r.source)
	if a < len(r.lines) {
		aEnd = r.lines[a] - 1
	}
	if c < len(r.lines) {
		cEnd = r.lines[c] - 1
	}
	if b-1 < aEnd-start {
		start += b - 1
	} else {
		start = aEnd
	}
	if d < cEnd-end {
		end += d
	} else {
		end = cEnd
	}
	return recoverySpan{start: start, end: end}, start <= end
}

func recoveryOverlap(spans []recoverySpan, start, end int) bool {
	for _, span := range spans {
		if start < span.end && end > span.start {
			return true
		}
	}
	return false
}

func recoveryEscaped(source []byte, start, at int) bool {
	count := 0
	for at > start {
		at--
		if source[at] != '\\' {
			break
		}
		count++
	}
	return count%2 != 0
}

func (r *labelRecovery) addEdit(start, end, labelStart, labelEnd int, colon bool) {
	text := "[^" + string(r.source[labelStart:labelEnd]) + "]"
	if colon {
		text += "\\:"
	}
	r.edits = append(r.edits, recoveryEdit{start: start, end: end, text: text})
}

func (r *labelRecovery) paragraphEdits(paragraph cmarkghmd.Node) {
	scope, ok := r.location(paragraph)
	if !ok {
		return
	}
	var text []byte
	var protectedSpans []recoverySpan
	stack := []cmarkghmd.Node{paragraph}
	for len(stack) != 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]
		kind := node.Type()
		if kind == "text" {
			text = append(text, node.Literal()...)
		}
		if kind == "code" || kind == "code_block" || kind == "html_inline" ||
			kind == "html_block" || kind == "link" || kind == "image" {
			if span, ok := r.location(node); ok {
				protectedSpans = append(protectedSpans, span)
			}
		}
		for child := node.FirstChild(); child.Valid(); child = child.Next() {
			stack = append(stack, child)
		}
	}

	source := r.source
	if bytes.Contains(text, []byte("[^]")) {
		for i := scope.start; i+2 < scope.end; i++ {
			if source[i] != '[' || source[i+1] != '^' {
				continue
			}
			label := i + 2
			p := label
			for p < scope.end && source[p] != ']' && source[p] != '\n' {
				p++
			}
			if p == label || p == scope.end || source[p] != '\n' {
				continue
			}
			labelEnd := p
			p++
			for p < scope.end && source[p] != ']' && source[p] != '\n' {
				p++
			}
			if p == scope.end || source[p] != ']' {
				continue
			}
			end := p + 1
			if !recoveryOverlap(protectedSpans, i, end) {
				colon := end < len(source) && source[end] == ':'
				editEnd := end
				if colon {
					editEnd++
				}
				r.addEdit(i, editEnd, label, labelEnd, colon)
			}
			i = p
		}
	}

	// Find unresolved references containing an inline image. Nested brackets,
	// parentheses, quoted destinations, and escaped delimiters are supported.
	for i := scope.start; i+3 < scope.end; i++ {
		if source[i] != '[' || source[i+1] != '^' || recoveryEscaped(source, scope.start, i) {
			continue
		}
		image := i + 2
		for image < scope.end && source[image] != '\n' &&
			(source[image] != ']' || recoveryEscaped(source, scope.start, image)) &&
			!(source[image] == '!' && image+1 < scope.end && source[image+1] == '[') {
			if source[image] == '[' && image+1 < scope.end && source[image+1] == '^' &&
				!recoveryEscaped(source, scope.start, image) {
				break
			}
			image++
		}
		if image+1 >= scope.end || source[image] != '!' || source[image+1] != '[' ||
			recoveryEscaped(source, scope.start, image) {
			continue
		}
		p := image + 2
		depth := 1
		for p < scope.end && source[p] != '\n' && depth != 0 {
			if source[p] == '\\' && p+1 < scope.end && source[p+1] != '\n' {
				p += 2
				continue
			}
			if source[p] == '[' {
				depth++
			} else if source[p] == ']' {
				depth--
			}
			p++
		}
		if depth != 0 || p >= scope.end || source[p] != '(' {
			continue
		}
		p++
		depth = 1
		angle := false
		var quote byte
		for p < scope.end && source[p] != '\n' && depth != 0 {
			ch := source[p]
			if ch == '\\' && p+1 < scope.end && source[p+1] != '\n' {
				p += 2
				continue
			}
			if quote != 0 {
				if ch == quote {
					quote = 0
				}
			} else if angle {
				if ch == '>' {
					angle = false
				}
			} else if ch == '<' {
				angle = true
			} else if (ch == '"' || ch == '\'') && p > scope.start &&
				(source[p-1] == ' ' || source[p-1] == '\t' || source[p-1] == '\r' ||
					source[p-1] == '\f' || source[p-1] == '\v') {
				quote = ch
			} else if ch == '(' {
				depth++
			} else if ch == ')' {
				depth--
			}
			p++
		}
		if depth != 0 || angle || quote != 0 || p >= scope.end || source[p] != ']' {
			continue
		}
		end := p + 1
		if bytes.Contains(text, source[i:end]) && !recoveryOverlap(protectedSpans, i, end) {
			r.addEdit(i, end, i+2, image, false)
		}
		i = p
	}
}

// recoverLabels repairs malformed footnote labels in source order while
// leaving protected code, HTML, links, and images untouched.
func recoverLabels(root cmarkghmd.Node, source []byte) (data []byte, changed bool) {
	r := labelRecovery{source: source, lines: []int{0}}
	for i, ch := range source {
		if ch == '\n' {
			r.lines = append(r.lines, i+1)
		}
	}

	node := root
	for node.Valid() {
		if node.Type() == "paragraph" {
			r.paragraphEdits(node)
		}
		if child := node.FirstChild(); child.Valid() {
			node = child
			continue
		}
		for node.Valid() {
			if next := node.Next(); next.Valid() {
				node = next
				break
			}
			node = node.Parent()
		}
	}

	if len(r.edits) > 1 {
		sort.Slice(r.edits, func(i, j int) bool {
			return r.edits[i].start > r.edits[j].start
		})
	}
	output := append([]byte(nil), source...)
	for _, edit := range r.edits {
		current := len(output)
		start, end := edit.start, edit.end
		if start > current {
			start = current
		}
		if end > current {
			end = current
		}
		if end < start {
			end = start
		}
		total := current - (end - start) + len(edit.text)
		if len(edit.text) > end-start {
			output = append(output, make([]byte, len(edit.text)-(end-start))...)
		}
		copy(output[start+len(edit.text):total], output[end:current])
		copy(output[start:start+len(edit.text)], edit.text)
		output = output[:total]
	}
	return output, !bytes.Equal(output, source)
}
