package github

import (
	"errors"
	"strings"

	"ghmd/internal/cmark-gfm"
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
func adaptText(node cmarkghmd.Node) error {
	input := node.Literal()
	if input == "" || !strings.Contains(input, "[^") {
		return nil
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
	if input == output {
		return nil
	}
	if !node.SetLiteral(output) {
		return errors.New("cannot replace Markdown literal")
	}
	return nil
}

// adaptListNode applies source-based task-list adjustments to an item node.
func adaptListNode(node cmarkghmd.Node, nodeType string, source []byte, lines []int) error {
	line, column := node.StartLine(), node.StartColumn()
	if line < 1 || line > len(lines) || column < 1 {
		return nil
	}
	begin := lines[line-1]
	end := len(source)
	if line < len(lines) {
		end = lines[line] - 1
	}
	if column-1 > end-begin {
		return nil
	}
	marker, ok := listMarker(source, begin+column-1, end)
	if !ok {
		return nil
	}
	if nodeType == "tasklist" {
		if end-marker >= 3 && source[marker] == '[' &&
			(source[marker+1] == ' ' || source[marker+1] == 'x' || source[marker+1] == 'X') &&
			source[marker+2] == ']' {
			node.SetTaskChecked(source[marker+1] != ' ')
		}
		return nil
	}

	paragraph := node.FirstChild()
	if !paragraph.Valid() || paragraph.Type() != "paragraph" {
		return nil
	}
	first := paragraph.FirstChild()
	if !first.Valid() || first.Type() != "text" {
		return nil
	}
	text := first.Literal()
	if len(text) < 4 || text[0] != '[' ||
		(text[1] != ' ' && text[1] != 'x' && text[1] != 'X') || text[2] != ']' ||
		(text[3] != ' ' && text[3] != '\t') {
		return nil
	}
	if !node.SetExtension("tasklist") {
		return errors.New("cannot attach tasklist extension")
	}
	node.SetTaskChecked(text[1] != ' ')
	if !first.SetLiteral(text[4:]) {
		return errors.New("cannot replace task literal")
	}
	return nil
}

// adaptNodes applies GitHub compatibility adjustments to the unmodified cmark AST.
func adaptNodes(root cmarkghmd.Node, source []byte) error {
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

	var visit func(cmarkghmd.Node) error
	visit = func(node cmarkghmd.Node) error {
		nodeType := node.Type()
		if nodeType == "text" {
			if err := adaptText(node); err != nil {
				return err
			}
		} else if nodeType == "item" || nodeType == "tasklist" {
			if err := adaptListNode(node, nodeType, source, lines); err != nil {
				return err
			}
		}
		for child := node.FirstChild(); child.Valid(); child = child.Next() {
			if err := visit(child); err != nil {
				return err
			}
		}
		return nil
	}
	return visit(root)
}
