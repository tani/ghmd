package github

// adjustCheckboxes marks task inputs and moves leading checkboxes into paragraphs.
func adjustCheckboxes(t *Tree, changed *bool) {
	originalCount := len(t.Nodes)
	for id := 1; id < originalCount; id++ {
		if t.Nodes[id].Inert || !t.Checkbox(id) {
			continue
		}
		item := t.Nearest(id, "li", "")
		if item == None {
			t.Detach(id)
			t.Nodes[id].Removed = true
			*changed = true
			continue
		}
		list := t.Nearest(id, "ul", "ol")
		t.Nodes[item].NormalizeClass = 3
		if list != None {
			t.Nodes[list].NormalizeClass = 4
		}
		t.Nodes[id].Checkbox = true
		_, _, t.Nodes[id].Checked = t.GetAttribute(id, "checked")
		*changed = true
	}
	for id := 1; id < originalCount; id++ {
		if t.Nodes[id].Inert || t.Nodes[id].Removed || !t.NodeIs(id, "li") ||
			(t.Nodes[id].NormalizeClass != 3 && !t.ClassContains(id, "task-list-item")) {
			continue
		}
		input := t.Nodes[id].First
		for input != None && !t.Nodes[input].Element {
			input = t.Nodes[input].Next
		}
		if input == None || !t.Checkbox(input) || t.Nodes[input].Removed {
			continue
		}
		paragraph := t.Nodes[input].Next
		for paragraph != None && !t.Nodes[paragraph].Element {
			paragraph = t.Nodes[paragraph].Next
		}
		if paragraph == None || !t.NodeIs(paragraph, "p") {
			continue
		}
		whitespace := t.Nodes[input].Next
		if whitespace != None && !t.Nodes[whitespace].Element && !t.TextHasVisibleContent(whitespace) {
			t.Detach(whitespace)
			t.Nodes[whitespace].Removed = true
		}
		t.Detach(input)
		space := t.appendNode(0, Node{Synthetic: true, SyntheticByte: ' '})
		t.Detach(space)
		t.InsertBefore(paragraph, space, t.Nodes[paragraph].First)
		t.InsertBefore(paragraph, input, space)
		if !t.TextStartsNewline(t.Nodes[id].First) {
			newline := t.appendNode(0, Node{Synthetic: true, SyntheticByte: '\n'})
			t.Detach(newline)
			t.InsertBefore(id, newline, t.Nodes[id].First)
		}
		*changed = true
	}
}

// writeCheckbox serializes a sanitized task checkbox with its generated state attributes.
func writeCheckbox(out *Buffer, t *Tree, node *Node) {
	s := t.Source
	p, written := node.TagEnd, node.Start
	hasID, hasClass, hasLabel := false, false, false
	label := "Incomplete task"
	if node.Checked {
		label = "Completed task"
	}
	for p < node.OpenEnd && s[p] != '>' {
		space := p
		for p < node.OpenEnd && asciiSpace(s[p]) {
			p++
		}
		attr := p
		for p < node.OpenEnd && !asciiSpace(s[p]) && s[p] != '=' && s[p] != '>' && s[p] != '/' {
			p++
		}
		if p == attr {
			p++
			continue
		}
		attrEnd := p
		for p < node.OpenEnd && asciiSpace(s[p]) {
			p++
		}
		a, b := p, p
		if p < node.OpenEnd && s[p] == '=' {
			p++
			for p < node.OpenEnd && asciiSpace(s[p]) {
				p++
			}
			if p < node.OpenEnd && (s[p] == '"' || s[p] == '\'') {
				quote := s[p]
				p++
				a = p
				for p < node.OpenEnd && s[p] != quote {
					p++
				}
				b = p
				if p < node.OpenEnd {
					p++
				}
			} else {
				a = p
				for p < node.OpenEnd && !asciiSpace(s[p]) && s[p] != '>' {
					p++
				}
				b = p
			}
		}
		checked := asciiEqual(s[attr:attrEnd], "checked")
		disabled := asciiEqual(s[attr:attrEnd], "disabled")
		id := asciiEqual(s[attr:attrEnd], "id")
		className := asciiEqual(s[attr:attrEnd], "class")
		aria := asciiEqual(s[attr:attrEnd], "aria-label")
		if !checked && !disabled && !id && !className && !aria {
			continue
		}
		out.Put(s[written:space])
		if id {
			out.Put(" id=\"\"")
			hasID = true
		}
		if className {
			writeClass(out, s[a:b], 5)
			hasClass = true
		}
		if aria {
			out.Put(" aria-label=\"")
			out.Put(label)
			out.Putc('"')
			hasLabel = true
		}
		written = p
	}
	out.Put(s[written : node.OpenEnd-1])
	if !hasID {
		out.Put(" id=\"\"")
	}
	out.Put(" disabled=\"\"")
	if !hasClass {
		out.Put(" class=\"task-list-item-checkbox\"")
	}
	if !hasLabel {
		out.Put(" aria-label=\"")
		out.Put(label)
		out.Putc('"')
	}
	if node.Checked {
		out.Put(" checked=\"\"")
	}
	out.Putc('>')
}
