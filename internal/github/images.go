package github

import "golang.org/x/net/html"

func externalImage(value string) bool {
	return asciiPrefixCase(value, "http://") ||
		asciiPrefixCase(value, "https://") ||
		asciiPrefixCase(value, "//")
}

// adjustImages marks images that require paragraph or relative-link wrapping.
func adjustImages(t *Tree) bool {
	changed := false
	for id := 1; id < len(t.Nodes); id++ {
		if t.Nodes[id].Inert || t.Nodes[id].Removed || !t.NodeIs(id, "img") {
			continue
		}
		start, end, ok := t.GetAttribute(id, "src")
		if !ok {
			continue
		}
		parent := t.Nodes[id].Parent
		if parent == 0 || (parent != None && t.NodeIs(parent, "b")) {
			t.Nodes[id].ImageParagraph = true
			changed = true
		}
		src := html.UnescapeString(t.Source[start:end])
		if !externalImage(src) {
			t.Nodes[id].RelativeImage = true
			t.Nodes[id].ImageRejected = asciiPrefixCase(src, "data:")
			t.Nodes[id].ImageLink = t.Nearest(id, "a", "") == None
			changed = true
		}
	}
	return changed
}

// writeImage serializes a relative image and its optional preview link.
func writeImage(out *Buffer, t *Tree, node *Node) {
	attrs := attributesRead(t, node)
	srcAttr := attrs.find("src")
	if srcAttr == nil {
		return
	}
	src := attributeValue(t, srcAttr)
	imageHref := ""
	if !node.ImageRejected {
		imageHref = src
	}
	if node.ImageRejected {
		attrs.remove("src")
	}
	attrs.set("style", "max-width: 100%;")
	if node.ImageLink {
		out.Put("<a target=\"_blank\" rel=\"")
		if asciiPrefixCase(imageHref, "mailto:") {
			out.Put("nofollow")
		} else {
			out.Put("noopener noreferrer")
		}
		out.Put("\" href=\"")
		writeAttributeEscape(out, imageHref)
		out.Put("\">")
	}
	out.Put(t.Source[node.Start:node.TagEnd])
	for i := range attrs.Items {
		attr := &attrs.Items[i]
		if attr.Removed {
			continue
		}
		if !attr.Changed {
			out.Put(t.Source[attr.RawStart:attr.RawEnd])
		} else {
			out.Putc(' ')
			out.Put(attr.Name)
			out.Put("=\"")
			writeAttributeEscape(out, attr.Value)
			out.Putc('"')
		}
	}
	out.Putc('>')
	if node.ImageLink {
		out.Put("</a>")
	}
}
