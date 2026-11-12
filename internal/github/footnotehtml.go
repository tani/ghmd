package github

// sliceFootnoteUTF16 implements JavaScript String.slice semantics for the
// nonnegative, code-unit offsets used by the footnote renderer.
func sliceFootnoteUTF16(text string, units int) string {
	i := 0
	for i < len(text) && units != 0 {
		cp := utf8At(text, i)
		width := 1
		for i+width < len(text) && text[i+width]&0xc0 == 0x80 {
			width++
		}
		if cp > 0xffff && units == 1 {
			return "\xef\xbf\xbd" + text[i+width:]
		}
		if cp > 0xffff {
			units -= 2
		} else {
			units--
		}
		i += width
	}
	return text[i:]
}

func makeFootnoteReferenceValue(t *Tree, attrs *Attributes, name, prefix string, slice int,
	token string, colonRemoves, lowercase bool) bool {
	attr := attrs.find(name)
	if attr == nil {
		return true
	}
	value := attributeValue(t, attr)
	if value == "" {
		return true
	}
	sliced := sliceFootnoteUTF16(value, slice)
	if colonRemoves && sliced == ":" {
		attrs.remove(name)
		return true
	}
	if lowercase {
		sliced = unicodeLower(sliced)
	}
	attrs.set(name, prefix+sliced+"-"+token)
	return true
}

// Footnote flags: 1 reference, 2 back-reference, 4 superscript, 8 section,
// 16 list item.
func writeFootnoteTag(out *Buffer, t *Tree, node *Node, token string, arrow *bool) bool {
	attrs := attributesRead(t, node)
	ok := true
	if node.Footnote&4 != 0 {
		attrs.remove("class")
	}
	if node.Footnote&1 != 0 {
		ok = makeFootnoteReferenceValue(t, &attrs, "href", "#user-content-fn-", 4, token, true, false) &&
			makeFootnoteReferenceValue(t, &attrs, "id", "user-content-", 0, token, false, true)
		if ok {
			attrs.set("aria-describedby", "footnote-label")
		}
	}
	if ok && node.Footnote&2 != 0 {
		if href := attrs.find("href"); href != nil {
			value := attributeValue(t, href)
			sliced := sliceFootnoteUTF16(value, 7)
			*arrow = sliced == ":"
			if !*arrow {
				ok = makeFootnoteReferenceValue(t, &attrs, "href", "#user-content-", 1, token, false, false)
			}
		}
		if !*arrow {
			attrs.remove("class")
			attrs.remove("data-footnote-backref-idx")
			if ok {
				attrs.set("class", "data-footnote-backref")
			}
		}
	}
	if ok && node.Footnote&8 != 0 {
		attrs.remove("class")
		attrs.remove("data-footnotes")
		attrs.set("data-footnotes", "")
		attrs.set("class", "footnotes")
	}
	if ok && node.Footnote&16 != 0 {
		// A list item id is rewritten even when the original id is absent or empty.
		attr := attrs.find("id")
		value := ""
		if attr != nil {
			value = attributeValue(t, attr)
		}
		attrs.set("id", "user-content-"+value+"-"+token)
	}
	if ok && *arrow {
		out.Put("↩")
	} else if ok {
		out.Put(t.Source[node.Start:node.TagEnd])
		classSeen := false
		for i := range attrs.Items {
			attr := &attrs.Items[i]
			if attr.Removed {
				continue
			}
			if node.NormalizeClass != 0 && asciiEqual(attr.Name, "class") {
				if attr.Changed {
					// Class on a footnote section/backref is not normalized elsewhere.
					out.Putc(' ')
					out.Put(attr.Name)
					out.Put("=\"")
					writeAttributeEscape(out, attr.Value)
					out.Putc('"')
				} else {
					writeClass(out, t.Source[attr.ValueStart:attr.ValueEnd], node.NormalizeClass)
				}
				classSeen = true
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
		if node.NormalizeClass > 1 && !classSeen {
			writeClass(out, "", node.NormalizeClass)
		}
		out.Putc('>')
		if node.Footnote&8 != 0 {
			out.Put("<h2 id=\"footnote-label\" class=\"sr-only\">Footnotes</h2>")
		}
	}
	return ok
}

func adjustFootnoteNodes(t *Tree) bool {
	changed := false
	section := None
	for id := 1; id < len(t.Nodes); id++ {
		if t.Nodes[id].Inert || t.Nodes[id].Removed {
			continue
		}
		if t.NodeIs(id, "sup") && t.ClassContains(id, "footnote-ref") {
			t.Nodes[id].Footnote |= 4
		}
		if t.NodeIs(id, "a") {
			if _, _, ok := t.GetAttribute(id, "data-footnote-ref"); ok {
				t.Nodes[id].Footnote |= 1
			}
			if _, _, ok := t.GetAttribute(id, "data-footnote-backref"); ok {
				t.Nodes[id].Footnote |= 2
			}
		}
		if section == None && t.NodeIs(id, "section") {
			if _, _, ok := t.GetAttribute(id, "data-footnotes"); ok {
				section = id
				t.Nodes[id].Footnote |= 8
				next := t.Nodes[id].Next
				if t.Nodes[id].Parent != 0 && next != None && !t.Nodes[next].Element &&
					t.Nodes[next].Next == None && !t.TextHasVisibleContent(next) {
					t.Detach(next)
					t.Nodes[next].Removed = true
				}
			}
		}
		if t.Nodes[id].Footnote != 0 {
			changed = true
		}
	}
	if section != None {
		for id := 1; id < len(t.Nodes); id++ {
			if t.Nodes[id].Inert || t.Nodes[id].Removed || !t.NodeIs(id, "li") {
				continue
			}
			if _, _, ok := t.GetAttribute(id, "id"); !ok {
				continue
			}
			for parent := t.Nodes[id].Parent; parent != None; parent = t.Nodes[parent].Parent {
				if parent == section {
					t.Nodes[id].Footnote |= 16
					break
				}
			}
		}
	}
	return changed
}
