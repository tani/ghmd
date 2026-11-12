package github

// fixListMathSpacing adds the line breaks expected around math in list items.
func fixListMathSpacing(t *Tree) bool {
	changed := false
	for id := 1; id < len(t.Nodes); id++ {
		if t.Nodes[id].Inert || !t.NodeIs(id, "li") {
			continue
		}
		first, second, last := None, None, None
		count := 0
		for child := t.Nodes[id].First; child != None; child = t.Nodes[child].Next {
			if !t.TextHasVisibleContent(child) {
				continue
			}
			if count == 0 {
				first = child
			} else if count == 1 {
				second = child
			}
			last = child
			count++
		}
		if count > 1 && t.Nodes[last].Element && t.ContainsMath(last) && !t.TextEndsInNewline(id) {
			t.Nodes[id].AppendNewline = true
			changed = true
		}
		if count > 1 && first == t.Nodes[id].First && t.NodeIs(first, "math-renderer") &&
			!t.Nodes[second].Element && t.TextHasVisibleContent(second) {
			t.Nodes[first].BeforeNewline = true
			changed = true
		}
	}
	return changed
}

func isList(t *Tree, id int) bool {
	return id != None && (t.NodeIs(id, "ul") || t.NodeIs(id, "ol"))
}

// flattenLists flattens lists nested more than ten levels in the same list type.
func flattenLists(t *Tree) bool {
	changed := false
	for id := 1; id < len(t.Nodes); id++ {
		if t.Nodes[id].Inert || t.Nodes[id].Removed || !isList(t, id) {
			continue
		}
		same := 1
		listName := "ol"
		if t.NodeIs(id, "ul") {
			listName = "ul"
		}
		for parent := t.Nodes[id].Parent; parent != None; parent = t.Nodes[parent].Parent {
			if isList(t, parent) && t.NodeIs(parent, listName) {
				same++
			}
		}
		if same <= 10 {
			continue
		}
		item := t.Nodes[id].Parent
		if item == None || !t.NodeIs(item, "li") {
			continue
		}
		parent := t.Nodes[item].Parent
		if !isList(t, parent) {
			continue
		}
		anchor := t.Nodes[item].Next
		for child := t.Nodes[id].First; child != None; {
			next := t.Nodes[child].Next
			if t.NodeIs(child, "li") {
				t.Detach(child)
				t.InsertBefore(parent, child, anchor)
				t.Nodes[child].NormalizeClass = 1
			}
			child = next
		}
		t.Detach(id)
		t.Nodes[id].Removed = true
		changed = true
	}
	return changed
}
