// Package verify canonicalizes HTML for fixture comparison.
package verify

import (
	"sort"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Canon is a canonical HTML node for structural comparison.
type Canon struct {
	Type      string      `json:"type"`
	Namespace string      `json:"namespace,omitempty"`
	Attrs     [][2]string `json:"attrs,omitempty"`
	Children  []Canon     `json:"children,omitempty"`
	Value     string      `json:"value,omitempty"`
}

var blockTags = toSet("address article aside blockquote body caption dd details dialog div dl dt fieldset figcaption figure footer form h1 h2 h3 h4 h5 h6 header hgroup hr li main menu nav ol p pre section summary table tbody td tfoot th thead tr ul")
var preserveTags = toSet("pre code textarea script style math-renderer")

func toSet(value string) map[string]bool {
	set := map[string]bool{}
	for _, item := range strings.Fields(value) {
		set[item] = true
	}
	return set
}

func parseFragment(fragment string) []*html.Node {
	context := &html.Node{Type: html.ElementNode, Data: "template", DataAtom: atom.Template}
	nodes, err := html.ParseFragment(strings.NewReader(fragment), context)
	if err != nil {
		return nil
	}
	return nodes
}

func attrValue(node *html.Node, name string) string {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

func generated(name, value string) string {
	if name == "data-run-id" && isRunID(value) {
		return "GENERATED"
	}
	if name == "id" || name == "href" {
		// user-content-fn...-<32 hex> suffix
		if index := strings.LastIndex(value, "-"); index >= 0 && isRunID(value[index+1:]) &&
			strings.Contains(value[:index], "user-content-fn") {
			return value[:index+1] + "GENERATED"
		}
	}
	return value
}

func isRunID(value string) bool {
	if len(value) != 32 {
		return false
	}
	for i := 0; i < len(value); i++ {
		c := value[i]
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'f') {
			return false
		}
	}
	return true
}

func textContents(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	var out strings.Builder
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		out.WriteString(textContents(child))
	}
	return out.String()
}

func codeBlock(node *html.Node) Canon {
	var attrs [][2]string
	for _, attr := range node.Attr {
		if attr.Key == "lang" || attr.Key == "class" && attr.Val == "notranslate" {
			continue
		}
		attrs = append(attrs, [2]string{attr.Key, generated(attr.Key, attr.Val)})
	}
	sortAttrs(attrs)
	return Canon{Type: "pre", Attrs: attrs, Children: []Canon{{
		Type:  "#text",
		Value: strings.TrimSuffix(textContents(node), "\n"),
	}}}
}

func highlightedChildren(node *html.Node) bool {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.TextNode {
			continue
		}
		if child.Data != "span" {
			return false
		}
		for _, attr := range child.Attr {
			if attr.Key != "class" {
				return false
			}
			for _, name := range strings.Fields(attr.Val) {
				if !strings.HasPrefix(name, "pl-") {
					return false
				}
			}
		}
		if !highlightedChildren(child) {
			return false
		}
	}
	return true
}

func sortAttrs(attrs [][2]string) {
	sort.Slice(attrs, func(i, j int) bool { return attrs[i][0] < attrs[j][0] })
}

func collapse(value string) string {
	var out strings.Builder
	for i := 0; i < len(value); {
		c := value[i]
		if c == '\t' || c == '\n' || c == '\f' || c == '\r' || c == ' ' {
			out.WriteByte(' ')
			for i < len(value) {
				c = value[i]
				if c != '\t' && c != '\n' && c != '\f' && c != '\r' && c != ' ' {
					break
				}
				i++
			}
			continue
		}
		out.WriteByte(c)
		i++
	}
	return out.String()
}

// HTMLAST returns the canonical tree used by the GitHub reference tests.
func HTMLAST(fragment string, ignoreHighlighting bool) Canon {
	root := Canon{Type: "#document"}
	for _, node := range parseFragment(fragment) {
		root.Children = append(root.Children, visit(node, false, ignoreHighlighting))
	}
	return root
}

func visit(node *html.Node, literal, ignoreHighlighting bool) Canon {
	if node.Type == html.TextNode {
		value := node.Data
		if !literal {
			value = collapse(value)
		}
		return Canon{Type: "#text", Value: value}
	}
	if node.Type == html.CommentNode {
		return Canon{Type: "#comment", Value: node.Data}
	}
	tag := node.Data
	if ignoreHighlighting && tag == "div" && hasClass(node, "highlight") {
		var children []*html.Node
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type != html.TextNode || strings.TrimSpace(child.Data) != "" {
				children = append(children, child)
			}
		}
		if len(children) == 1 && children[0].Data == "pre" && highlightedChildren(children[0]) {
			return codeBlock(children[0])
		}
	}
	if ignoreHighlighting && tag == "pre" {
		if node.FirstChild != nil && node.FirstChild == node.LastChild && node.FirstChild.Data == "code" &&
			onlyText(node.FirstChild) {
			return codeBlock(node)
		}
	}
	keep := literal || preserveTags[tag]
	var children []Canon
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, visit(child, keep, ignoreHighlighting))
	}
	if !keep {
		for i := len(children) - 1; i >= 0; i-- {
			child := &children[i]
			if child.Type != "#text" || child.Value == "" {
				continue
			}
			previous, next := "", ""
			if i > 0 {
				previous = children[i-1].Type
			}
			if i+1 < len(children) {
				next = children[i+1].Type
			}
			leftBlock := previous == ""
			if leftBlock {
				leftBlock = tag == "" || blockTags[tag]
			} else {
				leftBlock = blockTags[previous]
			}
			rightBlock := next == ""
			if rightBlock {
				rightBlock = tag == "" || blockTags[tag]
			} else {
				rightBlock = blockTags[next]
			}
			if child.Value == " " && (leftBlock || rightBlock) {
				children = append(children[:i], children[i+1:]...)
			} else {
				if leftBlock {
					child.Value = strings.TrimPrefix(child.Value, " ")
				}
				if rightBlock {
					child.Value = strings.TrimSuffix(child.Value, " ")
				}
				if child.Value == "" {
					children = append(children[:i], children[i+1:]...)
				}
			}
		}
	}
	result := Canon{Type: tag}
	if node.Namespace != "" && node.Namespace != "http://www.w3.org/1999/xhtml" {
		result.Namespace = node.Namespace
	}
	var attrs [][2]string
	for _, attr := range node.Attr {
		name := attr.Key
		if attr.Namespace != "" {
			name = attr.Namespace + ":" + name
		}
		attrs = append(attrs, [2]string{name, generated(name, attr.Val)})
	}
	sortAttrs(attrs)
	result.Attrs = attrs
	result.Children = children
	return result
}

func hasClass(node *html.Node, name string) bool {
	for _, value := range strings.Fields(attrValue(node, "class")) {
		if value == name {
			return true
		}
	}
	return false
}

func onlyText(node *html.Node) bool {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type != html.TextNode {
			return false
		}
	}
	return true
}

// NormalizeGitHubImages undoes GitHub's Camo proxy transformation.
func NormalizeGitHubImages(fragment string) string {
	nodes := parseFragment(fragment)
	for _, node := range nodes {
		normalizeImages(node)
	}
	var out strings.Builder
	for _, node := range nodes {
		html.Render(&out, node)
	}
	return out.String()
}

func normalizeImages(node *html.Node) {
	for child := node.FirstChild; child != nil; {
		next := child.NextSibling
		if isProxyWrapper(child) {
			image := child.FirstChild
			child.RemoveChild(image)
			node.RemoveChild(child)
			node.InsertBefore(image, next)
			normalizeImages(image)
		} else {
			normalizeImages(child)
		}
		child = next
	}
	if node.Type == html.ElementNode && node.Data == "img" && isCamo(node) {
		canonical := attrValue(node, "data-canonical-src")
		var attrs []html.Attribute
		for _, attr := range node.Attr {
			if attr.Key == "data-canonical-src" || attr.Key == "style" && attr.Val == "max-width: 100%;" {
				continue
			}
			if attr.Key == "src" {
				attr.Val = canonical
			}
			attrs = append(attrs, attr)
		}
		node.Attr = attrs
	}
}

func isCamo(node *html.Node) bool {
	src := attrValue(node, "src")
	canonical := attrValue(node, "data-canonical-src")
	return strings.HasPrefix(src, "https://camo.githubusercontent.com/") &&
		(strings.HasPrefix(strings.ToLower(canonical), "http://") || strings.HasPrefix(strings.ToLower(canonical), "https://"))
}

func isProxyWrapper(node *html.Node) bool {
	if node.Data != "a" || node.FirstChild == nil || node.FirstChild != node.LastChild ||
		node.FirstChild.Data != "img" {
		return false
	}
	image := node.FirstChild
	return isCamo(image) && attrValue(node, "href") == attrValue(image, "src") &&
		attrValue(node, "target") == "_blank" &&
		attrValue(node, "rel") == "noopener noreferrer nofollow" && len(node.Attr) == 3
}
