package github

import "golang.org/x/net/html"

// Attribute is one parsed start-tag attribute.
type Attribute struct {
	Name                     string
	NameLength               int
	RawStart, RawEnd         int
	ValueStart, ValueEnd     int
	Value                    string
	Loaded, Removed, Changed bool
}

// Attributes is the ordered attribute list of a start tag.
type Attributes struct {
	Items []Attribute
}

func attributesRead(t *Tree, node *Node) Attributes {
	s := t.Source
	attrs := Attributes{}
	p := node.TagEnd
	for p < node.OpenEnd && s[p] != '>' {
		space := p
		for p < node.OpenEnd && asciiSpace(s[p]) {
			p++
		}
		name := p
		for p < node.OpenEnd && !asciiSpace(s[p]) && s[p] != '=' && s[p] != '>' && s[p] != '/' {
			p++
		}
		if p == name {
			p++
			continue
		}
		nameEnd := p
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
		attrs.Items = append(attrs.Items, Attribute{
			Name: s[name:nameEnd], NameLength: nameEnd - name,
			RawStart: space, RawEnd: p, ValueStart: a, ValueEnd: b,
		})
	}
	return attrs
}

func (a *Attributes) find(name string) *Attribute {
	for i := range a.Items {
		if !a.Items[i].Removed && asciiEqual(a.Items[i].Name, name) {
			return &a.Items[i]
		}
	}
	return nil
}

func attributeValue(t *Tree, attr *Attribute) string {
	if !attr.Loaded {
		attr.Value = html.UnescapeString(t.Source[attr.ValueStart:attr.ValueEnd])
		attr.Loaded = true
	}
	return attr.Value
}

func (a *Attributes) remove(name string) {
	if attr := a.find(name); attr != nil {
		attr.Removed = true
	}
}

func (a *Attributes) set(name, value string) {
	attr := a.find(name)
	if attr == nil {
		a.Items = append(a.Items, Attribute{Name: name, NameLength: len(name)})
		attr = &a.Items[len(a.Items)-1]
	}
	attr.Value = value
	attr.Loaded = true
	attr.Changed = true
}
