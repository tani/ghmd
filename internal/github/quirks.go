package github

import (
	"strings"

	"yagfm/internal/cmark-gfm"
)

// listMarker returns the beginning of list-item content after start.
func listMarker(source []byte, start, end int) (int, bool) {
	p := start
	spaces := 0
	for p < end && source[p] == ' ' && spaces < 3 {
		p++
		spaces++
	}
	if p == end {
		return 0, false
	}
	if source[p] == '-' || source[p] == '+' || source[p] == '*' {
		p++
	} else {
		digits := 0
		for p < end && source[p] >= '0' && source[p] <= '9' && digits < 9 {
			p++
			digits++
		}
		if digits == 0 || p == end || (source[p] != '.' && source[p] != ')') {
			return 0, false
		}
		p++
	}
	if p == end || (source[p] != ' ' && source[p] != '\t') {
		return 0, false
	}
	for {
		p++
		if p >= end || (source[p] != ' ' && source[p] != '\t') {
			break
		}
	}
	return p, true
}

// adaptText applies GitHub's unresolved-footnote recovery to a text node.
func adaptText(node cmarkgfm.Node) bool {
	input := node.Literal()
	if input == "" || !strings.Contains(input, "[^") {
		return true
	}

	var first strings.Builder
	first.Grow(len(input))
	for i := 0; i < len(input); {
		image := i + 2
		if input[i] == '[' && i+1 < len(input) && input[i+1] == '^' {
			for image < len(input) && input[image] != ']' && input[image] != '\n' &&
				!(input[image] == '!' && image+1 < len(input) && input[image+1] == '[') {
				image++
			}
			if image+1 < len(input) && input[image] == '!' && input[image+1] == '[' {
				p := image + 2
				for p < len(input) && input[p] != ']' && input[p] != '\n' {
					p++
				}
				if p+1 < len(input) && input[p] == ']' && input[p+1] == '(' {
					p += 2
					for p < len(input) && input[p] != ')' && input[p] != '\n' {
						if input[p] == '\\' && p+1 < len(input) && input[p+1] != '\n' {
							p += 2
						} else {
							p++
						}
					}
					if p+1 < len(input) && input[p] == ')' && input[p+1] == ']' {
						first.WriteString(input[i:image])
						first.WriteByte(']')
						i = p + 2
						continue
					}
				}
			}
		}
		first.WriteByte(input[i])
		i++
	}

	value := first.String()
	var result strings.Builder
	result.Grow(len(value))
	const escapable = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	for i := 0; i < len(value); {
		if value[i] == '[' && i+1 < len(value) && value[i+1] == '^' {
			p := i + 2
			for p < len(value) && value[p] != ']' && value[p] != '\n' {
				if value[p] == '\\' {
					if p+1 == len(value) || value[p+1] == '\n' {
						break
					}
					p += 2
				} else {
					p++
				}
			}
			if p < len(value) && value[p] == ']' {
				for j := i; j <= p; j++ {
					if value[j] == '\\' && j < p && strings.IndexByte(escapable, value[j+1]) >= 0 {
						j++
					}
					result.WriteByte(value[j])
				}
				i = p + 1
				continue
			}
		}
		result.WriteByte(value[i])
		i++
	}

	output := result.String()
	return input == output || node.SetLiteral(output)
}

// adaptListNode applies source-based task-list adjustments to an item node.
func adaptListNode(node cmarkgfm.Node, nodeType string, source []byte, lines []int) bool {
	line, column := node.StartLine(), node.StartColumn()
	if line < 1 || line > len(lines) || column < 1 {
		return true
	}
	begin := lines[line-1]
	end := len(source)
	if line < len(lines) {
		end = lines[line] - 1
	}
	if column-1 > end-begin {
		return true
	}
	marker, ok := listMarker(source, begin+column-1, end)
	if !ok {
		return true
	}
	if nodeType == "tasklist" {
		if end-marker >= 3 && source[marker] == '[' &&
			(source[marker+1] == ' ' || source[marker+1] == 'x' || source[marker+1] == 'X') &&
			source[marker+2] == ']' {
			node.SetTaskChecked(source[marker+1] != ' ')
		}
		return true
	}

	paragraph := node.FirstChild()
	if !paragraph.Valid() || paragraph.Type() != "paragraph" {
		return true
	}
	first := paragraph.FirstChild()
	if !first.Valid() || first.Type() != "text" {
		return true
	}
	text := first.Literal()
	if len(text) < 4 || text[0] != '[' ||
		(text[1] != ' ' && text[1] != 'x' && text[1] != 'X') || text[2] != ']' ||
		(text[3] != ' ' && text[3] != '\t') {
		return true
	}
	if !node.SetExtension("tasklist") {
		return false
	}
	node.SetTaskChecked(text[1] != ' ')
	return first.SetLiteral(text[4:])
}

// adaptNodes applies GitHub compatibility adjustments to the unmodified cmark AST.
func adaptNodes(root cmarkgfm.Node, source []byte) bool {
	count := 1
	for _, c := range source {
		if c == '\n' {
			count++
		}
	}
	lines := make([]int, count)
	n := 0
	lines[n] = 0
	n++
	for i, c := range source {
		if c == '\n' {
			lines[n] = i + 1
			n++
		}
	}

	var visit func(cmarkgfm.Node) bool
	visit = func(node cmarkgfm.Node) bool {
		nodeType := node.Type()
		if nodeType == "text" {
			if !adaptText(node) {
				return false
			}
		} else if (nodeType == "item" || nodeType == "tasklist") &&
			!adaptListNode(node, nodeType, source, lines) {
			return false
		}
		for child := node.FirstChild(); child.Valid(); child = child.Next() {
			if !visit(child) {
				return false
			}
		}
		return true
	}
	return visit(root)
}
