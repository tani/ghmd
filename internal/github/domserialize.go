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

func domSerializeNode(out *Buffer, node *html.Node, depth int, raw, aliases bool) error {
	if len(out.Data) > maxDOMBufferBytes {
		return ErrSize
	}
	if depth > 512 {
		return ErrDepth
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
		if len(out.Data) > maxDOMBufferBytes {
			return ErrSize
		}
		return nil
	}
	if node.Type != html.ElementNode {
		return nil
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
		if len(out.Data) > maxDOMBufferBytes {
			return ErrSize
		}
		return nil
	}
	literal := node.Namespace == "" && wordList("style script xmp iframe noembed noframes plaintext noscript", name)
	for child := domChildren(node); child != nil; child = child.NextSibling {
		if err := domSerializeNode(out, child, depth+1, literal, aliases); err != nil {
			return err
		}
	}
	out.Put("</")
	out.Put(name)
	out.Putc('>')
	if len(out.Data) > maxDOMBufferBytes {
		return ErrSize
	}
	return nil
}

func domSerialize(root *html.Node, aliases bool) (string, error) {
	var out Buffer
	for child := domChildren(root); child != nil; child = child.NextSibling {
		if err := domSerializeNode(&out, child, 1, false, aliases); err != nil {
			return "", err
		}
	}
	if len(out.Data) > maxDOMBufferBytes {
		return "", ErrSize
	}
	return out.String(), nil
}
