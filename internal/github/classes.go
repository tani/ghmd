package github

import "golang.org/x/net/html"

func writeClass(out *Buffer, value string, action int) {
	added := ""
	switch action {
	case 2:
		added = "notranslate"
	case 3:
		added = "task-list-item"
	case 4:
		added = "contains-task-list"
	case 5:
		added = "task-list-item-checkbox"
	}
	decoded := html.UnescapeString(value)
	var tokens []string
	seen := map[string]bool{}
	for _, token := range splitSpaces(decoded) {
		if action == 1 && token == "task-list-item" {
			continue
		}
		if seen[token] {
			continue
		}
		seen[token] = true
		tokens = append(tokens, token)
	}
	if added != "" && !seen[added] {
		tokens = append(tokens, added)
	}
	if len(tokens) > 0 {
		out.Put(" class=\"")
		writeAttributeEscape(out, joinSpaces(tokens))
		out.Putc('"')
	}
}

func splitSpaces(value string) []string {
	var result []string
	for i := 0; i < len(value); {
		for i < len(value) && asciiSpace(value[i]) {
			i++
		}
		start := i
		for i < len(value) && !asciiSpace(value[i]) {
			i++
		}
		if i > start {
			result = append(result, value[start:i])
		}
	}
	return result
}

func joinSpaces(tokens []string) string {
	out := ""
	for i, token := range tokens {
		if i > 0 {
			out += " "
		}
		out += token
	}
	return out
}

func writeNormalizedTag(out *Buffer, t *Tree, id int, node *Node) {
	s := t.Source
	p, written := node.TagEnd, node.Start
	sawClass := false
	action := node.NormalizeClass
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
		if asciiEqual(s[attr:attrEnd], "class") {
			out.Put(s[written:space])
			writeClass(out, s[a:b], action)
			written = p
			sawClass = true
		} else if action == 2 && t.NodeIs(id, "pre") && asciiEqual(s[attr:attrEnd], "lang") {
			if html.UnescapeString(s[a:b]) == "math" {
				out.Put(s[written:space])
				written = p
			}
		}
	}
	if action != 1 && !sawClass {
		out.Put(s[written : node.OpenEnd-1])
		writeClass(out, "", action)
		out.Putc('>')
	} else {
		out.Put(s[written:node.OpenEnd])
	}
}
