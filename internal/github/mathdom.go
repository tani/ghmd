package github

import (
	"html"
	"strings"
)

type mathEdit struct {
	outer, inner        string
	trimFirst, trimLast bool
}

func mathText(out *Buffer, tree *Tree, edits []mathEdit, id int) {
	current := id
	for current != None {
		node := &tree.Nodes[current]
		if node.Element && !tree.NodeIs(current, "template") && node.First != None {
			current = node.First
			continue
		}
		if !node.Element {
			decoded := html.UnescapeString(tree.Source[node.Start:node.End])
			start := 0
			if edits[current].trimFirst && len(decoded) != 0 {
				start = 1
			}
			end := len(decoded)
			if edits[current].trimLast && end > start {
				end--
			}
			out.Put(decoded[start:end])
		}
		for current != id && tree.Nodes[current].Next == None {
			current = tree.Nodes[current].Parent
		}
		if current == id {
			break
		}
		current = tree.Nodes[current].Next
	}
}

func mathEscape(out *Buffer, text string, protectAngles bool) {
	for i := 0; i < len(text); i++ {
		switch {
		case text[i] == '&':
			out.Put("&amp;")
		case text[i] == '<':
			if protectAngles {
				out.Put("&amp;lt;")
			} else {
				out.Put("&lt;")
			}
		case text[i] == '>':
			if protectAngles {
				out.Put("&amp;gt;")
			} else {
				out.Put("&gt;")
			}
		case text[i] == 0xc2 && i+1 < len(text) && text[i+1] == 0xa0:
			out.Put("&nbsp;")
			i++
		default:
			out.Putc(text[i])
		}
	}
}

func mathElement(out *Buffer, text string, display, protectAngles bool, token string) {
	if display {
		out.Put("<math-renderer class=\"js-display-math\" style=\"display: block\" data-run-id=\"")
	} else {
		out.Put("<math-renderer class=\"js-inline-math\" style=\"display: inline-block\" data-run-id=\"")
	}
	out.Put(mathCString(token))
	out.Put("\">")
	mathEscape(out, text, protectAngles)
	out.Put("</math-renderer>")
}

func mathClosest(tree *Tree, id int, names ...string) bool {
	for p := id; p != None; p = tree.Nodes[p].Parent {
		for _, name := range names {
			if tree.NodeIs(p, name) {
				return true
			}
		}
	}
	return false
}

func mathBlocked(tree *Tree, id int) bool {
	for p := tree.Nodes[id].Parent; p != None; p = tree.Nodes[p].Parent {
		if tree.NodeIs(p, "details") {
			return false
		}
		if tree.NodeIs(p, "b") || tree.NodeIs(p, "strong") || tree.NodeIs(p, "em") ||
			tree.NodeIs(p, "a") || tree.NodeIs(p, "code") || tree.NodeIs(p, "pre") ||
			tree.NodeIs(p, "section") {
			return true
		}
	}
	return false
}

var mathTargetNames = [...]string{
	"p", "li", "h1", "h2", "h3", "h4", "h5", "h6",
	"td", "th", "strong", "em", "i", "del", "span", "summary",
}

func mathTarget(tree *Tree, id int) bool {
	for _, name := range mathTargetNames {
		if tree.NodeIs(id, name) {
			return true
		}
	}
	return false
}

func mathReplacedAncestor(tree *Tree, edits []mathEdit, id int) bool {
	for p := tree.Nodes[id].Parent; p != None; p = tree.Nodes[p].Parent {
		if edits[p].inner != "" || edits[p].outer != "" {
			return true
		}
	}
	return false
}

// mathCString reproduces the C call sites that consume a detached string with
// strlen or cmark_strbuf_puts rather than an explicit byte length.
func mathCString(s string) string {
	if i := strings.IndexByte(s, 0); i >= 0 {
		return s[:i]
	}
	return s
}

func mathApply(tree *Tree, edits []mathEdit) {
	for id := 1; id < len(tree.Nodes); id++ {
		if tree.Nodes[id].Inert || mathReplacedAncestor(tree, edits, id) || !mathTarget(tree, id) {
			continue
		}
		if mathBlocked(tree, id) && !(tree.NodeIs(id, "strong") &&
			!mathClosest(tree, id, "b", "section", "a", "pre", "code")) {
			continue
		}

		onlyTextBR, onlyTextBREm, hasEm := true, true, false
		var reconstructed Buffer
		for child := tree.Nodes[id].First; child != None; child = tree.Nodes[child].Next {
			if !tree.Nodes[child].Element {
				mathText(&reconstructed, tree, edits, child)
			} else if tree.NodeIs(child, "em") {
				onlyTextBR = false
				hasEm = true
				reconstructed.Putc('_')
				mathText(&reconstructed, tree, edits, child)
				reconstructed.Putc('_')
			} else if !tree.NodeIs(child, "br") {
				onlyTextBR = false
				onlyTextBREm = false
			}
		}
		if hasEm && onlyTextBREm && !mathClosest(tree, id, "b") {
			display := emphasisDisplay(reconstructed.String())
			if display != "" {
				var out Buffer
				mathElement(&out, mathCString(display), true, false, tree.Token)
				edits[id].inner = out.String()
				continue
			}
		}

		if onlyTextBR && !tree.NodeIs(id, "li") && !mathClosest(tree, id, "b", "details") {
			var text Buffer
			mathText(&text, tree, edits, id)
			blocks := displayMath(text.String())
			if blocks == nil {
				return
			}
			if len(blocks) != 0 {
				var out Buffer
				valueText := text.String()
				for i := range len(blocks) / 2 {
					start, end := blocks[i*2], blocks[i*2+1]
					var value Buffer
					mathPutLines(&value, valueText[start:end])
					if i != 0 {
						out.Putc('\n')
					}
					mathElement(&out, value.String(), true, false, tree.Token)
				}
				edits[id].inner = out.String()
			}
			if edits[id].inner != "" {
				continue
			}
		}

		for code := tree.Nodes[id].First; code != None; code = tree.Nodes[code].Next {
			if !tree.NodeIs(code, "code") {
				continue
			}
			before, after := tree.Nodes[code].Prev, tree.Nodes[code].Next
			if before == None || after == None || tree.Nodes[before].Element || tree.Nodes[after].Element {
				continue
			}
			var a, b, c Buffer
			mathText(&a, tree, edits, before)
			mathText(&b, tree, edits, after)
			mathText(&c, tree, edits, code)
			value, ok := protectedMath(a.String(), b.String(), c.String(), tree.NodeIs(id, "em"))
			if ok {
				var out Buffer
				mathElement(&out, mathCString(value), false, false, tree.Token)
				edits[code].outer = out.String()
				edits[before].trimLast = true
				edits[after].trimFirst = true
			}
		}

		if tree.NodeIs(id, "em") || tree.NodeIs(id, "strong") && mathClosest(tree, id, "em") {
			continue
		}
		for child := tree.Nodes[id].First; child != None; child = tree.Nodes[child].Next {
			if tree.Nodes[child].Element {
				continue
			}
			var text Buffer
			mathText(&text, tree, edits, child)
			decoded := text.String()
			matches := inlineMath(decoded)
			if matches == nil {
				return
			}
			if len(matches) != 0 {
				var out Buffer
				previous := uint32(0)
				for i := range len(matches) / 2 {
					start, end := matches[i*2], matches[i*2+1]
					mathEscape(&out, decoded[previous:start], false)
					mathElement(&out, decoded[start:end], false, true, tree.Token)
					previous = end
				}
				mathEscape(&out, decoded[previous:], false)
				edits[child].outer = out.String()
			}
		}
	}

	for id := 1; id < len(tree.Nodes); id++ {
		if tree.Nodes[id].Inert || mathReplacedAncestor(tree, edits, id) ||
			!tree.NodeIs(id, "pre") || !tree.AttributeEquals(id, "lang", "math") || mathBlocked(tree, id) {
			continue
		}
		var text, out Buffer
		mathText(&text, tree, edits, id)
		value := mathCString(fencedMath(text.String()))
		mathElement(&out, value, true, false, tree.Token)
		edits[id].outer = out.String()
	}
}

func mathDOM(source, token, boundary string) string {
	tree := NewTree(source)
	tree.Token = token
	tree.TextBoundary = boundary
	tree.ParseSerializedHTML()
	edits := make([]mathEdit, len(tree.Nodes))
	mathApply(tree, edits)
	changed := tree.HadTextBoundary
	var out Buffer
	id := tree.Nodes[0].First
	for id != None {
		node := &tree.Nodes[id]
		if edits[id].outer != "" {
			out.Put(mathCString(edits[id].outer))
			changed = true
		} else if !node.Element {
			if edits[id].trimFirst || edits[id].trimLast {
				var text Buffer
				mathText(&text, tree, edits, id)
				mathEscape(&out, text.String(), false)
				changed = true
			} else {
				out.Put(source[node.Start:node.End])
			}
		} else {
			out.Put(source[node.Start:node.OpenEnd])
		}
		if edits[id].outer == "" && edits[id].inner != "" {
			out.Put(mathCString(edits[id].inner))
			changed = true
		}
		if edits[id].outer == "" && edits[id].inner == "" && node.First != None {
			id = node.First
			continue
		}
		for id != 0 {
			node = &tree.Nodes[id]
			if node.Element && edits[id].outer == "" && node.End > node.CloseStart {
				out.Put(source[node.CloseStart:node.End])
			}
			if node.Next != None {
				id = node.Next
				break
			}
			id = node.Parent
		}
		if id == 0 {
			break
		}
	}
	if !changed {
		return source
	}
	return out.String()
}
