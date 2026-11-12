package github

import (
	"strings"

	"golang.org/x/net/html"
)

// None is the absent node index.
const None = -1

// Node is a byte-range HTML node used by the renderer's postprocessing.
type Node struct {
	Start, OpenEnd, CloseStart, End                         int
	TagStart, TagEnd                                        int
	Parent, First, Last, Prev, Next                         int
	TableWrappers                                           int
	Element, Inert, Removed                                 bool
	NormalizeClass                                          int
	BeforeNewline, AppendNewline                            bool
	WrapTable, SuppressTableWrapper                         bool
	Checkbox, Checked, Synthetic                            bool
	Alert                                                   int
	Footnote                                                int
	Arrow                                                   bool
	RelativeImage, ImageLink, ImageRejected, ImageParagraph bool
	Comment, TableMarks                                     bool
	SyntheticByte                                           byte
}

// Tree is a parse of already serialized, sanitized HTML.
type Tree struct {
	Source          string
	Nodes           []Node
	Token           string
	TableMarker     string
	TextBoundary    string
	HadTextBoundary bool
	SplitComments   bool
}

// NewTree creates a tree rooted at node 0.
func NewTree(source string) *Tree {
	return &Tree{Source: source, Nodes: []Node{{Parent: None, First: None, Last: None, Prev: None, Next: None}}}
}

func (t *Tree) appendNode(parent int, node Node) int {
	node.Parent = parent
	node.First, node.Last, node.Next = None, None, None
	if parent == None {
		return None
	}
	node.Prev = t.Nodes[parent].Last
	id := len(t.Nodes)
	t.Nodes = append(t.Nodes, node)
	if node.Prev != None {
		t.Nodes[node.Prev].Next = id
	} else {
		t.Nodes[parent].First = id
	}
	t.Nodes[parent].Last = id
	return id
}

// NodeIs reports whether the element's tag name equals name (ASCII case-insensitive).
func (t *Tree) NodeIs(id int, name string) bool {
	if id < 0 || id >= len(t.Nodes) {
		return false
	}
	n := &t.Nodes[id]
	return n.Element && asciiEqual(t.Source[n.TagStart:n.TagEnd], name)
}

func asciiSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f'
}

func asciiEqual(s string, expected string) bool {
	if len(s) != len(expected) {
		return false
	}
	for i := range s {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != expected[i] {
			return false
		}
	}
	return true
}

func asciiPrefixCase(value string, prefix string) bool {
	if len(value) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		c := value[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != prefix[i] {
			return false
		}
	}
	return true
}

func tagEnd(s string, length, start int) int {
	var quote byte
	for i := start; i < length; i++ {
		if quote != 0 {
			if s[i] == quote {
				quote = 0
			}
		} else if s[i] == '\'' || s[i] == '"' {
			quote = s[i]
		} else if s[i] == '>' {
			return i + 1
		}
	}
	return length
}

func nameChar(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-' || c == ':'
}

var voidElements = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}
var rawElements = []string{"script", "style", "textarea", "title", "xmp", "iframe", "noembed", "noframes", "plaintext"}

func inList(names []string, value string) bool {
	for _, name := range names {
		if asciiEqual(value, name) {
			return true
		}
	}
	return false
}

func voidElement(name string) bool { return inList(voidElements, name) }
func rawElement(name string) bool  { return inList(rawElements, name) }

// ParseSerializedHTML parses sanitized HTML into the byte-range tree.
// The parser cannot fail: node storage grows by append.
func (t *Tree) ParseSerializedHTML() {
	t.Nodes = []Node{{Parent: None, First: None, Last: None, Prev: None, Next: None}}
	stack := []int{0}
	s := t.Source
	length := len(s)
	i, textStart := 0, 0
	for i < length {
		if s[i] != '<' {
			i++
			continue
		}
		active := stack[len(stack)-1]
		if active != 0 {
			an := &t.Nodes[active]
			if rawElement(s[an.TagStart:an.TagEnd]) {
				nameLength := an.TagEnd - an.TagStart
				after := i + 2 + nameLength
				closes := i+2 <= length && s[i+1] == '/' && after <= length &&
					(after == length || asciiSpace(s[after]) || s[after] == '>' || s[after] == '/')
				for k := 0; closes && k < nameLength; k++ {
					a, b := s[i+2+k], s[an.TagStart+k]
					if a >= 'A' && a <= 'Z' {
						a += 'a' - 'A'
					}
					if b >= 'A' && b <= 'Z' {
						b += 'a' - 'A'
					}
					if a != b {
						closes = false
					}
				}
				if !closes {
					i++
					continue
				}
			}
		}
		if i+4 <= length && s[i:i+4] == "<!--" {
			markerStart := i
			i += 4
			for i+3 <= length && s[i:i+3] != "-->" {
				i++
			}
			end := length
			if i+3 <= length {
				end = i + 3
			}
			if t.SplitComments {
				current := stack[len(stack)-1]
				if markerStart > textStart {
					t.appendNode(current, Node{Start: textStart, End: markerStart, Parent: None, First: None, Last: None, Prev: None, Next: None})
				}
				t.appendNode(current, Node{Start: markerStart, End: end, Comment: true, Parent: None, First: None, Last: None, Prev: None, Next: None})
				i, textStart = end, end
				continue
			}
			if t.TextBoundary != "" && end >= textStart+7 && end-markerStart == len(t.TextBoundary)+7 &&
				s[markerStart+4:markerStart+4+len(t.TextBoundary)] == t.TextBoundary {
				marker := markerStart
				if marker > textStart {
					t.appendNode(stack[len(stack)-1], Node{Start: textStart, End: marker, Parent: None, First: None, Last: None, Prev: None, Next: None})
				}
				t.HadTextBoundary = true
				textStart = end
			}
			i = end
			continue
		}
		tagStart := i + 1
		closing := tagStart < length && s[tagStart] == '/'
		if closing {
			tagStart++
		}
		tagEndPos := tagStart
		for tagEndPos < length && nameChar(s[tagEndPos]) {
			tagEndPos++
		}
		if tagEndPos == tagStart {
			i++
			continue
		}
		current := stack[len(stack)-1]
		if current != 0 {
			cn := &t.Nodes[current]
			if rawElement(s[cn.TagStart:cn.TagEnd]) &&
				!(closing && tagEndPos-tagStart == cn.TagEnd-cn.TagStart &&
					s[tagStart:tagEndPos] == s[cn.TagStart:cn.TagEnd]) {
				i++
				continue
			}
		}
		end := tagEnd(s, length, tagEndPos)
		if i > textStart {
			t.appendNode(current, Node{Start: textStart, End: i, Parent: None, First: None, Last: None, Prev: None, Next: None})
		}
		cn := &t.Nodes[current]
		if closing {
			if len(stack) > 1 && tagEndPos-tagStart == cn.TagEnd-cn.TagStart &&
				s[tagStart:tagEndPos] == s[cn.TagStart:cn.TagEnd] {
				cn.CloseStart = i
				cn.End = end
				stack = stack[:len(stack)-1]
			} else {
				t.appendNode(current, Node{Start: i, End: end, Parent: None, First: None, Last: None, Prev: None, Next: None})
			}
		} else {
			inert := cn.Inert || t.NodeIs(current, "template")
			element := Node{Start: i, OpenEnd: end, CloseStart: end, End: end,
				TagStart: tagStart, TagEnd: tagEndPos, Element: true, Inert: inert,
				Parent: None, First: None, Last: None, Prev: None, Next: None}
			id := t.appendNode(current, element)
			if !voidElement(s[tagStart:tagEndPos]) {
				stack = append(stack, id)
			}
		}
		i, textStart = end, end
	}
	if textStart < length {
		t.appendNode(stack[len(stack)-1], Node{Start: textStart, End: length, Parent: None, First: None, Last: None, Prev: None, Next: None})
	}
}

// Detach removes a node from its parent's child list.
func (t *Tree) Detach(id int) {
	n := &t.Nodes[id]
	parent := &t.Nodes[n.Parent]
	if n.Prev == None {
		parent.First = n.Next
	} else {
		t.Nodes[n.Prev].Next = n.Next
	}
	if n.Next == None {
		parent.Last = n.Prev
	} else {
		t.Nodes[n.Next].Prev = n.Prev
	}
	n.Parent, n.Prev, n.Next = None, None, None
}

// InsertBefore inserts id into parent before anchor (None appends).
func (t *Tree) InsertBefore(parentID, id, anchor int) {
	parent := &t.Nodes[parentID]
	n := &t.Nodes[id]
	n.Parent = parentID
	n.Next = anchor
	if anchor == None {
		n.Prev = parent.Last
	} else {
		n.Prev = t.Nodes[anchor].Prev
	}
	if n.Prev == None {
		parent.First = id
	} else {
		t.Nodes[n.Prev].Next = id
	}
	if anchor == None {
		parent.Last = id
	} else {
		t.Nodes[anchor].Prev = id
	}
}

// GetAttribute returns an attribute's raw value span.
func (t *Tree) GetAttribute(id int, name string) (start, end int, ok bool) {
	s := t.Source
	node := &t.Nodes[id]
	p := node.TagEnd
	for p < node.OpenEnd && s[p] != '>' {
		for p < node.OpenEnd && asciiSpace(s[p]) {
			p++
		}
		attr := p
		for p < node.OpenEnd && !asciiSpace(s[p]) && s[p] != '=' && s[p] != '>' && s[p] != '/' {
			p++
		}
		if p == attr {
			p++
			continue
		}
		attrEnd := p
		for p < node.OpenEnd && asciiSpace(s[p]) {
			p++
		}
		a, b := p, p
		if p < node.OpenEnd && s[p] == '=' {
			p++
			for p < node.OpenEnd && asciiSpace(s[p]) {
				p++
			}
			if p < node.OpenEnd && (s[p] == '"' || s[p] == '\'') {
				quote := s[p]
				p++
				a = p
				for p < node.OpenEnd && s[p] != quote {
					p++
				}
				b = p
				if p < node.OpenEnd {
					p++
				}
			} else {
				a = p
				for p < node.OpenEnd && !asciiSpace(s[p]) && s[p] != '>' {
					p++
				}
				b = p
			}
		}
		if asciiEqual(s[attr:attrEnd], name) {
			return a, b, true
		}
	}
	return 0, 0, false
}

// AttributeEquals compares a decoded attribute value.
func (t *Tree) AttributeEquals(id int, name, wanted string) bool {
	start, end, ok := t.GetAttribute(id, name)
	if !ok {
		return false
	}
	return html.UnescapeString(t.Source[start:end]) == wanted
}

// ClassContains checks a decoded class token.
func (t *Tree) ClassContains(id int, wanted string) bool {
	start, end, ok := t.GetAttribute(id, "class")
	if !ok {
		return false
	}
	for _, token := range strings.Fields(html.UnescapeString(t.Source[start:end])) {
		if token == wanted {
			return true
		}
	}
	return false
}

// Nearest walks ancestors for a matching tag.
func (t *Tree) Nearest(id int, a, b string) int {
	for n := id; n != None; n = t.Nodes[n].Parent {
		if t.NodeIs(n, a) || (b != "" && t.NodeIs(n, b)) {
			return n
		}
	}
	return None
}

// Checkbox reports an input[type=checkbox].
func (t *Tree) Checkbox(id int) bool {
	return t.NodeIs(id, "input") && t.AttributeEquals(id, "type", "checkbox")
}

// TextHasVisibleContent implements DOM textContent visibility.
func (t *Tree) TextHasVisibleContent(id int) bool {
	node := &t.Nodes[id]
	if node.Element {
		return true
	}
	if node.Synthetic {
		return !IsJSWhitespace(int32(node.SyntheticByte))
	}
	value := html.UnescapeString(t.Source[node.Start:node.End])
	for _, r := range value {
		if !IsJSWhitespace(r) {
			return true
		}
	}
	return false
}

// ContainsMath reports a math-renderer descendant.
func (t *Tree) ContainsMath(id int) bool {
	for n := id; n != None; {
		if t.NodeIs(n, "math-renderer") {
			return true
		}
		if !t.NodeIs(n, "template") && t.Nodes[n].First != None {
			n = t.Nodes[n].First
			continue
		}
		for n != id && t.Nodes[n].Next == None {
			n = t.Nodes[n].Parent
		}
		if n == id {
			break
		}
		n = t.Nodes[n].Next
	}
	return false
}

// TextEndsInNewline inspects the last visible text descendant.
func (t *Tree) TextEndsInNewline(id int) bool {
	n := t.Nodes[id].Last
	for n != None && n != id {
		node := &t.Nodes[n]
		if node.Element && !t.NodeIs(n, "template") && node.Last != None {
			n = node.Last
			continue
		}
		if !node.Element && node.Synthetic {
			return node.SyntheticByte == '\n'
		}
		if !node.Element && node.End > node.Start {
			value := html.UnescapeString(t.Source[node.Start:node.End])
			if value != "" {
				return value[len(value)-1] == '\n'
			}
		}
		for n != id && t.Nodes[n].Prev == None {
			n = t.Nodes[n].Parent
		}
		if n == id {
			break
		}
		n = t.Nodes[n].Prev
	}
	return false
}

// TextStartsNewline inspects a text node's first byte.
func (t *Tree) TextStartsNewline(id int) bool {
	if id == None || t.Nodes[id].Element {
		return false
	}
	if t.Nodes[id].Synthetic {
		return t.Nodes[id].SyntheticByte == '\n'
	}
	node := &t.Nodes[id]
	value := html.UnescapeString(t.Source[node.Start:node.End])
	return len(value) > 0 && value[0] == '\n'
}
