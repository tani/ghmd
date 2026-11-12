package github

// prepareDetails inserts GitHub's default summary into details elements without one.
func prepareDetails(source string) string {
	tree := NewTree(source)
	tree.ParseSerializedHTML()
	var out Buffer
	id := tree.Nodes[0].First
	for id != None {
		node := &tree.Nodes[id]
		if !node.Element {
			out.Put(source[node.Start:node.End])
		} else {
			out.Put(source[node.Start:node.OpenEnd])
			if !node.Inert && tree.NodeIs(id, "details") {
				hasSummary := false
				for child := node.First; child != None; child = tree.Nodes[child].Next {
					if tree.NodeIs(child, "summary") {
						hasSummary = true
						break
					}
				}
				if !hasSummary {
					out.Put("<summary>Details</summary>")
				}
			}
		}
		if node.First != None {
			id = node.First
			continue
		}
		for id != 0 {
			node = &tree.Nodes[id]
			if node.Element && node.End > node.CloseStart {
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

// unwrapTemplates moves template children into their parents without merging text nodes.
func unwrapTemplates(t *Tree) {
	for id := 1; id < len(t.Nodes); id++ {
		if !t.NodeIs(id, "template") || t.Nodes[id].Parent == None {
			continue
		}
		parent := t.Nodes[id].Parent
		for t.Nodes[id].First != None {
			child := t.Nodes[id].First
			t.Detach(child)
			t.InsertBefore(parent, child, id)
		}
		t.Detach(id)
	}
	for id := 1; id < len(t.Nodes); id++ {
		t.Nodes[id].Inert = false
	}
}

// prepareWrite serializes an existing tree without invoking another HTML parser.
func prepareWrite(t *Tree, out *Buffer, root int, boundary, tableMarker string, stripComments, preserveBoundaries bool) {
	id := root
	if root == 0 {
		id = t.Nodes[0].First
	}
	if id == None {
		return
	}
	for id != None {
		node := &t.Nodes[id]
		if node.Comment {
			if !stripComments {
				out.Put(t.Source[node.Start:node.End])
			}
		} else if !node.Element {
			if preserveBoundaries && node.Prev != None {
				prev := node.Prev
				for stripComments && prev != None && t.Nodes[prev].Comment {
					prev = t.Nodes[prev].Prev
				}
				if prev != None && !t.Nodes[prev].Element && !t.Nodes[prev].Comment {
					out.Put("<!--")
					out.Put(boundary)
					out.Put("-->")
				}
			}
			out.Put(t.Source[node.Start:node.End])
		} else {
			for range node.TableWrappers {
				out.Put("<markdown-accessiblity-table>")
			}
			if node.TableMarks {
				out.Put(t.Source[node.Start : node.OpenEnd-1])
				out.Putc(' ')
				out.Put(tableMarker)
				out.Put("=\"\">")
			} else {
				out.Put(t.Source[node.Start:node.OpenEnd])
			}
		}
		if node.First != None {
			id = node.First
			continue
		}
		for id != 0 {
			node = &t.Nodes[id]
			if node.Element && node.End > node.CloseStart {
				out.Put(t.Source[node.CloseStart:node.End])
			}
			for range node.TableWrappers {
				out.Put("</markdown-accessiblity-table>")
			}
			if id == root {
				return
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
}

// flattenTemplates exposes template contents while retaining text boundaries.
func flattenTemplates(source, boundary string) string {
	tree := NewTree(source)
	tree.SplitComments = true
	tree.ParseSerializedHTML()
	unwrapTemplates(tree)
	var out Buffer
	prepareWrite(tree, &out, 0, boundary, "", false, true)
	return out.String()
}

// removeBoundaries removes only generated text-boundary comments.
func removeBoundaries(source, boundary string) string {
	var out Buffer
	n, start := len(boundary), 0
	for i := 0; i+n+7 <= len(source); i++ {
		if source[i:i+4] != "<!--" || source[i+4:i+4+n] != boundary ||
			source[i+4+n:i+n+7] != "-->" {
			continue
		}
		out.Put(source[start:i])
		i += n + 6
		start = i + 1
	}
	out.Put(source[start:])
	return out.String()
}

// prepareMathTables decides paragraph-local table wrapping before math rendering.
func prepareMathTables(source, token, boundary, tableMarker string) (string, error) {
	tree := NewTree(source)
	tree.SplitComments = true
	tree.ParseSerializedHTML()
	subtreeEnd := make([]int, len(tree.Nodes))
	for id := range subtreeEnd {
		subtreeEnd[id] = id + 1
	}
	for id := len(tree.Nodes) - 1; id >= 1; id-- {
		parent := tree.Nodes[id].Parent
		if parent != None && subtreeEnd[parent] < subtreeEnd[id] {
			subtreeEnd[parent] = subtreeEnd[id]
		}
	}
	for id := 1; id < len(tree.Nodes); id++ {
		if tree.Nodes[id].Inert || !tree.NodeIs(id, "p") {
			continue
		}
		blocked, hasTable := false, false
		for parent := tree.Nodes[id].Parent; parent != None; parent = tree.Nodes[parent].Parent {
			if tree.NodeIs(parent, "details") {
				break
			}
			if tree.NodeIs(parent, "b") || tree.NodeIs(parent, "strong") ||
				tree.NodeIs(parent, "em") || tree.NodeIs(parent, "a") ||
				tree.NodeIs(parent, "code") || tree.NodeIs(parent, "pre") ||
				tree.NodeIs(parent, "section") {
				blocked = true
				break
			}
		}
		for n := id + 1; n < subtreeEnd[id]; n++ {
			if tree.NodeIs(n, "table") {
				hasTable = true
				break
			}
		}
		if !hasTable {
			continue
		}
		hasMath := false
		if !blocked {
			var paragraph Buffer
			prepareWrite(tree, &paragraph, id, boundary, tableMarker, true, false)
			paragraphSource := paragraph.String()
			rendered := mathDOM(paragraphSource, token, boundary)
			hasMath = rendered != paragraphSource
		}
		for n := id + 1; n < subtreeEnd[id]; n++ {
			if !tree.NodeIs(n, "table") {
				continue
			}
			tree.Nodes[n].TableMarks = true
			if !hasMath {
				tree.Nodes[n].TableWrappers++
			}
		}
	}
	var out Buffer
	prepareWrite(tree, &out, 0, boundary, tableMarker, true, true)
	return out.String(), nil
}
