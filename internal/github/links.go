package github

func writeLink(out *Buffer, t *Tree, id int, marker string) (unwrap bool) {
	node := &t.Nodes[id]
	attrs := attributesRead(t, node)
	marked := attrs.find("data-gfm-had-href")
	hadHref := false
	if marked != nil {
		hadHref = attributeValue(t, marked) == marker
		marked.Removed = true
	}
	href := attrs.find("href")
	unwrap = hadHref && href == nil && attrs.find("id") == nil && attrs.find("name") == nil
	if !unwrap && href != nil {
		value := attributeValue(t, href)
		if externalImage(value) || asciiPrefixCase(value, "mailto:") {
			attrs.set("rel", "nofollow")
		}
	}
	if !unwrap {
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
	}
	return unwrap
}

// processLinks restores sanitized links and applies GitHub's external-link rel policy.
func processLinks(source, boundary string) string {
	tree := NewTree(source)
	tree.ParseSerializedHTML()
	var out Buffer
	id := tree.Nodes[0].First
	for id != None {
		node := &tree.Nodes[id]
		if !node.Element {
			out.Put(source[node.Start:node.End])
		} else if !node.Inert && tree.NodeIs(id, "a") {
			node.Removed = writeLink(&out, tree, id, boundary)
		} else {
			out.Put(source[node.Start:node.OpenEnd])
		}
		if node.First != None {
			id = node.First
			continue
		}
		for id != 0 {
			node = &tree.Nodes[id]
			if node.Element && !node.Removed && node.End > node.CloseStart {
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
	return out.String()
}
