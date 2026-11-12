package github

import "golang.org/x/net/html"

type domAlias struct {
	original string
	neutral  string
}

var domAliases = [...]domAlias{
	{original: "template", neutral: "gfm-native-template"},
	{original: "table", neutral: "gfm-native-table"},
	{original: "thead", neutral: "gfm-native-thead"},
	{original: "tbody", neutral: "gfm-native-tbody"},
	{original: "tfoot", neutral: "gfm-native-tfoot"},
	{original: "tr", neutral: "gfm-native-tr"},
	{original: "td", neutral: "gfm-native-td"},
	{original: "th", neutral: "gfm-native-th"},
	{original: "caption", neutral: "gfm-native-caption"},
	{original: "colgroup", neutral: "gfm-native-colgroup"},
	{original: "col", neutral: "gfm-native-col"},
	{original: "p", neutral: "gfm-native-p"},
}

func neutralDOMName(name string) string {
	for _, alias := range domAliases {
		if name == alias.original {
			return alias.neutral
		}
	}
	return name
}

func originalDOMName(name string) string {
	for _, alias := range domAliases {
		if name == alias.neutral {
			return alias.original
		}
	}
	return name
}

func domIs(node *html.Node, name string) bool {
	return node != nil && node.Type == html.ElementNode && originalDOMName(node.Data) == name
}

func domIn(node *html.Node, names string) bool {
	return node != nil && node.Type == html.ElementNode && wordList(names, originalDOMName(node.Data))
}

func domPermitted(parent *html.Node) string {
	if domIs(parent, "table") {
		return "caption colgroup thead tfoot tbody"
	}
	if domIn(parent, "thead tbody tfoot") {
		return "tr"
	}
	if domIs(parent, "tr") {
		return "td th"
	}
	if domIs(parent, "colgroup") {
		return "col"
	}
	return ""
}

func domRequired(node *html.Node) string {
	if domIn(node, "thead tbody tfoot caption") {
		return "table"
	}
	if domIs(node, "tr") {
		return "thead tbody tfoot"
	}
	if domIn(node, "td th") {
		return "tr"
	}
	if domIs(node, "col") {
		return "colgroup"
	}
	return ""
}

func domCloseParent(parent, node *html.Node) bool {
	if domIs(parent, "p") {
		return domIn(node, "address article aside blockquote details dialog div dl fieldset figcaption figure footer form h1 h2 h3 h4 h5 h6 header hgroup hr main menu nav ol p pre section table ul")
	}
	if domIn(parent, "td th") {
		return domIn(node, "td th tr tbody tfoot thead")
	}
	if domIs(parent, "caption") {
		return domIn(node, "table tbody thead tfoot tr td th col colgroup")
	}
	if domIs(parent, "option") {
		return domIn(node, "option optgroup")
	}
	return false
}

func domInsertAfter(after, node *html.Node) {
	after.Parent.InsertBefore(node, after.NextSibling)
}

func domChildrenBefore(node *html.Node) {
	for node.FirstChild != nil {
		child := node.FirstChild
		node.RemoveChild(child)
		node.Parent.InsertBefore(child, node)
	}
}

// domMoveAfter moves a complete sibling tail in order and returns its new end.
func domMoveAfter(after, first *html.Node) *html.Node {
	for first != nil {
		next := first.NextSibling
		first.Parent.RemoveChild(first)
		domInsertAfter(after, first)
		after = first
		first = next
	}
	return after
}

func domUnwrap(node *html.Node) {
	domChildrenBefore(node)
	node.Parent.RemoveChild(node)
}

func domSameAncestor(root, node *html.Node) *html.Node {
	name := originalDOMName(node.Data)
	for parent := root; parent != nil; parent = parent.Parent {
		if parent.Type == html.ElementNode && originalDOMName(parent.Data) == name {
			return parent
		}
	}
	return nil
}

func domCloseRelocated(root, tail, ancestor *html.Node) {
	after := ancestor
	for root != nil {
		nextRoot := root.Parent
		nextTail := root.NextSibling
		after = domMoveAfter(after, tail)
		if root == ancestor {
			break
		}
		root = nextRoot
		tail = nextTail
	}
}

func domCompat(root *html.Node, depth int) bool {
	if depth > 512 {
		return false
	}
	node := root.FirstChild
	for node != nil {
		if node.Type != html.ElementNode {
			node = node.NextSibling
			continue
		}
		// A neutralized col still has HTML void-element semantics.
		if domIs(node, "col") {
			domMoveAfter(node, node.FirstChild)
		}
		if domCloseParent(root, node) && root.Parent != nil {
			domMoveAfter(root, node)
			return true
		}
		tbody := domIs(root, "table") && domIs(node, "tr")
		permitted := domPermitted(root)
		if permitted != "" && !domIn(node, permitted) && !tbody {
			close := (domIn(root, "thead tbody tfoot") && domIn(node, "caption colgroup thead tfoot tbody")) ||
				(domIs(root, "tr") && domIn(node, "caption colgroup thead tfoot tbody tr"))
			if close && root.Parent != nil {
				domMoveAfter(root, node)
				return true
			}
			// Forbidden elements go before the table; their contents remain at
			// the current insertion point. Text is not foster-parented.
			hadChildren := node.FirstChild != nil
			tail := node.NextSibling
			closes := domSameAncestor(root, node)
			next := node.NextSibling
			if hadChildren {
				next = node.FirstChild
			}
			domChildrenBefore(node)
			if closes != nil {
				domCloseRelocated(root, tail, closes)
			}
			before := root
			for before.Parent != nil {
				allowed := domPermitted(before.Parent)
				if allowed == "" || domIn(node, allowed) {
					break
				}
				before = before.Parent
			}
			node.Parent.RemoveChild(node)
			if before.Parent != nil {
				before.Parent.InsertBefore(node, before)
			} else {
				before.AppendChild(node)
			}
			if closes != nil && !hadChildren {
				return true
			}
			node = next
			continue
		}
		if tbody {
			body := &html.Node{Type: html.ElementNode, Data: "gfm-native-tbody"}
			node.Parent.InsertBefore(body, node)
			child := node
			for child != nil && domIs(child, "tr") {
				next := child.NextSibling
				child.Parent.RemoveChild(child)
				body.AppendChild(child)
				child = next
			}
			node = body
		}
		parents := domRequired(node)
		if parents != "" && !domIn(root, parents) {
			next := node.NextSibling
			if node.FirstChild != nil {
				next = node.FirstChild
			}
			domChildrenBefore(node)
			node.Parent.RemoveChild(node)
			node = next
			continue
		}
		if !domCompat(node, depth+1) {
			return false
		}
		node = node.NextSibling
	}
	return true
}
