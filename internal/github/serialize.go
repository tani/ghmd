package github

func writeTableTag(out *Buffer, t *Tree, id int, node *Node) bool {
	attrs := attributesRead(t, node)
	if t.TableMarker != "" {
		attrs.remove(t.TableMarker)
	}
	attrs.set("role", "table")
	if !node.SuppressTableWrapper {
		out.Put("<markdown-accessiblity-table>")
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
	return true
}

func serializeTree(t *Tree, out *Buffer) bool {
	id := t.Nodes[0].First
	for id != None {
		node := &t.Nodes[id]
		if node.BeforeNewline {
			out.Putc('\n')
		}
		if node.ImageParagraph {
			out.Put("<p>")
		}
		if node.Synthetic {
			out.Putc(node.SyntheticByte)
		} else if !node.Element {
			out.Put(t.Source[node.Start:node.End])
		} else if node.Checkbox {
			writeCheckbox(out, t, node)
		} else if node.RelativeImage {
			if !writeImage(out, t, node) {
				return false
			}
		} else if node.Footnote != 0 {
			arrow := false
			if !writeFootnoteTag(out, t, node, t.Token, &arrow) {
				return false
			}
			node.Arrow = arrow
		} else if node.Alert != 0 {
			kind := node.Alert - 1
			out.Put("<div class=\"markdown-alert markdown-alert-")
			out.Put(alertNames[kind])
			out.Put("\"><p class=\"markdown-alert-title\">")
			out.Put(alertTitles[kind])
			out.Put("</p>")
		} else if node.WrapTable {
			if !writeTableTag(out, t, id, node) {
				return false
			}
		} else if node.NormalizeClass != 0 {
			writeNormalizedTag(out, t, id, node)
		} else {
			out.Put(t.Source[node.Start:node.OpenEnd])
		}
		if !node.Arrow && node.First != None {
			id = node.First
			continue
		}
		for id != 0 {
			node = &t.Nodes[id]
			if node.AppendNewline {
				out.Putc('\n')
			}
			if !node.Arrow {
				if node.Alert != 0 {
					out.Put("</div>")
				} else if node.Element && node.End > node.CloseStart {
					out.Put(t.Source[node.CloseStart:node.End])
				}
			}
			if node.WrapTable && !node.SuppressTableWrapper {
				out.Put("</markdown-accessiblity-table>")
			}
			if node.ImageParagraph {
				out.Put("</p>")
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
	return true
}

// FinalizeSerializedHTML applies the final feature adjustments and serializes.
func FinalizeSerializedHTML(source, token, tableMarker string) (string, bool) {
	tree := NewTree(source)
	tree.Token = token
	tree.TableMarker = tableMarker
	if !tree.ParseSerializedHTML() {
		return "", false
	}
	changed := adjustAlerts(tree)
	if !adjustCheckboxes(tree, &changed) {
		return "", false
	}
	if fixListMathSpacing(tree) {
		changed = true
	}
	if adjustImages(tree) {
		changed = true
	}
	if adjustFootnoteNodes(tree) {
		changed = true
	}
	if flattenLists(tree) {
		changed = true
	}
	for i := 1; i < len(tree.Nodes); i++ {
		if !tree.Nodes[i].Inert && tree.NodeIs(i, "table") {
			tree.Nodes[i].WrapTable = true
			tree.Nodes[i].SuppressTableWrapper = tableMarker != "" && hasAttribute(tree, i, tableMarker)
			changed = true
		}
		if !tree.Nodes[i].Inert && (tree.NodeIs(i, "code") || tree.NodeIs(i, "pre")) {
			tree.Nodes[i].NormalizeClass = 2
			changed = true
		}
	}
	if !changed {
		return source, true
	}
	var out Buffer
	if !serializeTree(tree, &out) {
		return "", false
	}
	return out.String(), true
}

func hasAttribute(t *Tree, id int, name string) bool {
	_, _, ok := t.GetAttribute(id, name)
	return ok
}
