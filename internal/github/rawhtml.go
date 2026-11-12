package github

import (
	"sort"

	"golang.org/x/net/html"
)

type rawSpan struct {
	start int
	end   int
}

func domHasInput(source string) bool {
	for i := 0; i+6 <= len(source); i++ {
		if source[i] == '<' && asciiPrefixCase(source[i+1:], "input") &&
			(i+6 == len(source) || asciiSpace(source[i+6]) || source[i+6] == '/' || source[i+6] == '>') {
			return true
		}
	}
	return false
}

func inputSpans(node *html.Node, doc *domDocument, spans *[]rawSpan, depth int) bool {
	if depth > 512 {
		return false
	}
	if node.Type == html.ElementNode && node.Data == "input" {
		if loc := doc.locations[node]; loc != nil {
			*spans = append(*spans, rawSpan{start: loc.start, end: loc.openEnd})
		}
	}
	for child := domChildren(node); child != nil; child = child.NextSibling {
		if !inputSpans(child, doc, spans, depth+1) {
			return false
		}
	}
	return true
}

// filterRawInputs removes only input start tags that became source-authored
// parser elements. Text in attributes, raw-text elements, and ignored tokens is
// left byte-for-byte intact.
func filterRawInputs(source string) (string, bool) {
	if !domHasInput(source) {
		return source, true
	}
	doc, ok := parseDOM(source, false, true)
	if !ok {
		return "", false
	}
	var spans []rawSpan
	if !inputSpans(doc.root, doc, &spans, 0) {
		return "", false
	}
	sort.Slice(spans, func(i, j int) bool { return spans[i].start < spans[j].start })
	var out Buffer
	start := 0
	for _, span := range spans {
		if span.start < start || span.start < 0 || span.end < span.start || span.end > len(source) {
			continue
		}
		out.Put(source[start:span.start])
		start = span.end
	}
	out.Put(source[start:])
	return out.String(), true
}
