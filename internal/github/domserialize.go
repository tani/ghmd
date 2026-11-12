package github

import "golang.org/x/net/html"

func domSerializeEscape(out *Buffer, value string, attribute bool) {
	for i := 0; i < len(value); i++ {
		if value[i] == '&' {
			out.Put("&amp;")
		} else if value[i] == '"' && attribute {
			out.Put("&quot;")
		} else if value[i] == '<' && !attribute {
			out.Put("&lt;")
		} else if value[i] == '>' && !attribute {
			out.Put("&gt;")
		} else if value[i] == 0xc2 && i+1 < len(value) && value[i+1] == 0xa0 {
			out.Put("&nbsp;")
			i++
		} else {
			out.Putc(value[i])
		}
	}
}

func domAttributeName(attr html.Attribute) string {
	if attr.Namespace != "" {
		return attr.Namespace + ":" + attr.Key
	}
	return attr.Key
}

func domSerializeNode(out *Buffer, node *html.Node, depth int, raw, aliases bool) bool {
	if len(out.Data) > maxDOMBufferBytes {
		return false
	}
	if depth > 512 {
		return false
	}
	if node.Type == html.TextNode || node.Type == html.CommentNode {
		if node.Type == html.CommentNode {
			out.Put("<!--")
			out.Put(node.Data)
			out.Put("-->")
		} else if raw {
			out.Put(node.Data)
		} else {
			domSerializeEscape(out, node.Data, false)
		}
		return len(out.Data) <= maxDOMBufferBytes
	}
	if node.Type != html.ElementNode {
		return true
	}
	name := node.Data
	if aliases {
		name = originalDOMName(name)
	}
	out.Putc('<')
	out.Put(name)
	for _, attr := range node.Attr {
		out.Putc(' ')
		out.Put(domAttributeName(attr))
		out.Put("=\"")
		domSerializeEscape(out, attr.Val, true)
		out.Putc('"')
	}
	out.Putc('>')
	if node.Namespace == "" && wordList("area base br col embed hr img input keygen link meta param source track wbr", name) {
		return len(out.Data) <= maxDOMBufferBytes
	}
	literal := node.Namespace == "" && wordList("style script xmp iframe noembed noframes plaintext noscript", name)
	for child := domChildren(node); child != nil; child = child.NextSibling {
		if !domSerializeNode(out, child, depth+1, literal, aliases) {
			return false
		}
	}
	out.Put("</")
	out.Put(name)
	out.Putc('>')
	return len(out.Data) <= maxDOMBufferBytes
}

func domSerialize(root *html.Node, aliases bool) (string, bool) {
	var out Buffer
	for child := domChildren(root); child != nil; child = child.NextSibling {
		if !domSerializeNode(&out, child, 1, false, aliases) {
			return "", false
		}
	}
	if len(out.Data) > maxDOMBufferBytes {
		return "", false
	}
	return out.String(), true
}
