// Package analysis holds the common parsing and comparison plumbing shared by
// the diff command and the linter. It does not know authoring policies.
package analysis

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Text returns all text content below a node, excluding non-text nodes.
func Text(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	var out strings.Builder
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		out.WriteString(Text(child))
	}
	return out.String()
}

// Paragraphs returns top-level fragment nodes using body fragment semantics.
func Paragraphs(source string) ([]*html.Node, error) {
	context := &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body}
	return html.ParseFragment(strings.NewReader(source), context)
}

// Walk visits nodes in document order, skipping template/script/style/head content.
func Walk(node *html.Node, visit func(*html.Node)) {
	if skipTag(node.Data) {
		return
	}
	if node.Type != html.ElementNode && node.Type != html.TextNode {
		return
	}
	visit(node)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		Walk(child, visit)
	}
}

// Attr reports the first attribute value and whether it exists.
func Attr(node *html.Node, name string) (string, bool) {
	for _, attr := range node.Attr {
		if attr.Namespace == "" && attr.Key == name {
			return attr.Val, true
		}
	}
	return "", false
}

func HasAttr(node *html.Node, name string) bool { _, ok := Attr(node, name); return ok }

// Class reports a token membership on a class attribute.
func Class(node *html.Node, wanted string) bool {
	value, ok := Attr(node, "class")
	return ok && wordList(value, wanted)
}

// WordList is the same ASCII-whitespace token list helper used by the sanitizer.
func wordList(value, wanted string) bool {
	for _, candidate := range strings.Fields(value) {
		if candidate == wanted {
			return true
		}
	}
	return false
}

func skipTag(data string) bool {
	switch data {
	case "script", "style", "head", "template":
		return true
	}
	return false
}

// CollapseWhitespace replaces runs of ASCII layout whitespace with one space.
func CollapseWhitespace(value string) string {
	var out strings.Builder
	pending := false
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case ' ', '\t', '\n', '\r', '\f':
			pending = true
		default:
			if pending {
				out.WriteByte(' ')
			}
			pending = false
			out.WriteByte(value[i])
		}
	}
	if pending {
		out.WriteByte(' ')
	}
	return out.String()
}
