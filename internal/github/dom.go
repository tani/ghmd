package github

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type domLocation struct {
	start   int
	openEnd int
	closed  bool
}

type domDocument struct {
	root      *html.Node
	locations map[*html.Node]*domLocation
}

type domSourceToken struct {
	kind    html.TokenType
	name    string
	start   int
	end     int
	nameEnd int
	attrs   []html.Attribute
}

func sourceNameEnd(raw string, prefix int) int {
	i := prefix
	for i < len(raw) && !asciiSpace(raw[i]) && raw[i] != '/' && raw[i] != '>' {
		i++
	}
	return i
}

func tokenizeDOM(source string, keepAttrs bool) []domSourceToken {
	z := html.NewTokenizer(strings.NewReader(source))
	var tokens []domSourceToken
	offset := 0
	for {
		kind := z.Next()
		raw := string(z.Raw())
		start := offset
		offset += len(raw)
		if kind == html.ErrorToken {
			break
		}
		if kind != html.StartTagToken && kind != html.SelfClosingTagToken && kind != html.EndTagToken {
			continue
		}
		token := z.Token()
		prefix := 1
		if kind == html.EndTagToken {
			prefix = 2
		}
		var attrs []html.Attribute
		if keepAttrs {
			attrs = append(attrs, token.Attr...)
		}
		tokens = append(tokens, domSourceToken{
			kind:    kind,
			name:    token.Data,
			start:   start,
			end:     offset,
			nameEnd: start + sourceNameEnd(raw, prefix),
			attrs:   attrs,
		})
	}
	return tokens
}

func provenanceAttribute(tokens []domSourceToken) string {
	used := make(map[string]bool)
	for _, token := range tokens {
		if token.kind == html.EndTagToken {
			continue
		}
		for _, attr := range token.attrs {
			used[attr.Key] = true
		}
	}
	const base = "data-gfm-dom-provenance"
	if !used[base] {
		return base
	}
	for i := 1; ; i++ {
		name := base + "-" + strconv.Itoa(i)
		if !used[name] {
			return name
		}
	}
}

func markExplicitEnds(tokens []domSourceToken, locations []*domLocation) {
	type openTag struct {
		name string
		loc  *domLocation
	}
	var open []openTag
	nextLocation := 0
	for _, token := range tokens {
		if token.kind == html.EndTagToken {
			for i := len(open) - 1; i >= 0; i-- {
				if open[i].name == token.name {
					open[i].loc.closed = true
					open = open[:i]
					break
				}
			}
			continue
		}
		loc := locations[nextLocation]
		nextLocation++
		if wordList("area base br col embed hr img input keygen link meta param source track wbr", token.name) {
			continue
		}
		open = append(open, openTag{name: token.name, loc: loc})
	}
}

func rewriteDOMSource(source string, tokens []domSourceToken, aliases, tracked bool) (string, string, []*domLocation) {
	attribute := ""
	if tracked {
		attribute = provenanceAttribute(tokens)
	}
	var locations []*domLocation
	if tracked {
		for _, token := range tokens {
			if token.kind != html.EndTagToken {
				locations = append(locations, &domLocation{start: token.start, openEnd: token.end})
			}
		}
		markExplicitEnds(tokens, locations)
	}
	var out strings.Builder
	last, nextLocation := 0, 0
	for _, token := range tokens {
		id := -1
		if token.kind != html.EndTagToken && tracked {
			id = nextLocation
			nextLocation++
		}
		name := token.name
		if aliases {
			name = neutralDOMName(name)
		}
		if id < 0 && name == token.name {
			continue
		}
		prefix := 1
		if token.kind == html.EndTagToken {
			prefix = 2
		}
		nameStart := token.start + prefix
		if token.nameEnd < nameStart || token.nameEnd > len(source) {
			continue
		}
		out.WriteString(source[last:nameStart])
		if name == token.name {
			out.WriteString(source[nameStart:token.nameEnd])
		} else {
			out.WriteString(name)
		}
		if id >= 0 {
			out.WriteByte(' ')
			out.WriteString(attribute)
			out.WriteString("=\"")
			out.WriteString(strconv.Itoa(id))
			out.WriteByte('"')
		}
		last = token.nameEnd
	}
	out.WriteString(source[last:])
	return out.String(), attribute, locations
}

func recordDOMLocations(node *html.Node, document *domDocument, attribute string, locations []*domLocation) {
	if node.Type == html.ElementNode {
		id := -1
		attrs := node.Attr[:0]
		for _, attr := range node.Attr {
			if attr.Namespace == "" && attr.Key == attribute {
				parsed, err := strconv.Atoi(attr.Val)
				if err == nil && parsed >= 0 && parsed < len(locations) {
					id = parsed
				}
				continue
			}
			attrs = append(attrs, attr)
		}
		node.Attr = attrs
		if id >= 0 {
			document.locations[node] = locations[id]
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		recordDOMLocations(child, document, attribute, locations)
	}
}

func parseDOM(source string, aliases, tracked bool) (*domDocument, error) {
	tokens := tokenizeDOM(source, tracked)
	rewritten, attribute, locations := rewriteDOMSource(source, tokens, aliases, tracked)
	// A template fragment context is the direct x/net counterpart of the
	// former parser's template context. Using html.Parse and taking its body would drop
	// standalone rows and cells before the compatibility pass can repair them;
	// this context retains those nodes and template children.
	context := &html.Node{Type: html.ElementNode, Data: "template", DataAtom: atom.Template}
	children, err := html.ParseFragment(strings.NewReader(rewritten), context)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParse, err)
	}
	document := &domDocument{
		root:      &html.Node{Type: html.DocumentNode},
		locations: make(map[*html.Node]*domLocation),
	}
	for _, child := range children {
		document.root.AppendChild(child)
	}
	if tracked {
		recordDOMLocations(document.root, document, attribute, locations)
	}
	return document, nil
}

func domChildren(node *html.Node) *html.Node {
	return node.FirstChild
}

func domLastChild(node *html.Node) *html.Node {
	return node.LastChild
}

// domReparse neutralizes table/template names before parsing so HTML's native
// insertion modes do not erase structure that the compatibility pass owns.
func domReparse(source string) (string, error) {
	if len(source) > maxOutputBytes {
		return "", ErrSize
	}
	document, err := parseDOM(source, true, false)
	if err != nil {
		return "", err
	}
	if !domCompat(document.root, 0) {
		return "", ErrDepth
	}
	return domSerialize(document.root, true)
}

func domAnchors(root *html.Node, document *domDocument, seen []bool, depth int) error {
	if depth > 512 {
		return ErrDepth
	}
	node := domChildren(root)
	for node != nil {
		if node.Type == html.ElementNode && node.Data == "a" {
			if loc := document.locations[node]; loc != nil && loc.start >= 0 && loc.start < len(seen) {
				if seen[loc.start] {
					next := node.NextSibling
					if node.FirstChild != nil {
						next = node.FirstChild
					}
					domUnwrap(node)
					node = next
					continue
				}
				seen[loc.start] = true
			}
		}
		if err := domAnchors(node, document, seen, depth+1); err != nil {
			return err
		}
		node = node.NextSibling
	}
	return nil
}

func domHasEarlier(root, tail *html.Node, document *domDocument, loc *domLocation, depth int) bool {
	if depth > 512 || root == tail {
		return false
	}
	other := document.locations[root]
	if root.Data == tail.Data && other != nil && other.start == loc.start {
		return true
	}
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		if domHasEarlier(child, tail, document, loc, depth+1) {
			return true
		}
	}
	return false
}

func domUTF8At(value string, i int) int32 {
	if i >= len(value) {
		return 0
	}
	c := value[i]
	if c < 0x80 {
		return int32(c)
	}
	n := 1
	if c >= 0xf0 && c <= 0xf4 {
		n = 4
	} else if c >= 0xe0 && c <= 0xef {
		n = 3
	} else if c >= 0xc2 && c <= 0xdf {
		n = 2
	}
	if n == 1 || i+n > len(value) {
		return 0xfffd
	}
	cp := int32(c & byte(0x7f>>uint(n)))
	for j := 1; j < n; j++ {
		if value[i+j]&0xc0 != 0x80 {
			return 0xfffd
		}
		cp = cp<<6 | int32(value[i+j]&0x3f)
	}
	return cp
}

func domWhitespace(root *html.Node) bool {
	for child := root.FirstChild; child != nil; child = child.NextSibling {
		if child.Type != html.TextNode {
			return false
		}
		for i := 0; i < len(child.Data); {
			cp := domUTF8At(child.Data, i)
			if !IsJSWhitespace(cp) {
				return false
			}
			i++
			for i < len(child.Data) && child.Data[i]&0xc0 == 0x80 {
				i++
			}
		}
	}
	return true
}

func domTrim(document *domDocument) {
	last := document.root.LastChild
	loc := document.locations[last]
	if last != nil && last.Type == html.ElementNode && loc != nil && !loc.closed {
		leaf := last
		for child := domLastChild(leaf); child != nil; child = domLastChild(leaf) {
			leaf = child
		}
		if leaf.Type == html.TextNode && len(leaf.Data) != 0 && leaf.Data[len(leaf.Data)-1] == '\n' {
			leaf.Data = leaf.Data[:len(leaf.Data)-1]
		}
	}
	last = document.root.LastChild
	loc = document.locations[last]
	if last != nil && last.Type == html.ElementNode && loc != nil && domWhitespace(last) &&
		domHasEarlier(document.root, last, document, loc, 0) {
		document.root.RemoveChild(last)
	}
}

func domSVG(root *html.Node, changed *bool, depth int) error {
	if depth > 512 {
		return ErrDepth
	}
	node := domChildren(root)
	for node != nil {
		if node.Namespace == "svg" && node.Data != "a" {
			next := node.NextSibling
			if node.FirstChild != nil {
				next = node.FirstChild
			}
			domUnwrap(node)
			node = next
			*changed = true
			continue
		}
		if err := domSVG(node, changed, depth+1); err != nil {
			return err
		}
		node = node.NextSibling
	}
	return nil
}

func domMark(root *html.Node, marker string, depth int) error {
	if depth > 512 {
		return ErrDepth
	}
	if root.Type == html.ElementNode && root.Data == "a" {
		hasHref := false
		for _, attr := range root.Attr {
			if attr.Namespace == "" && attr.Key == "href" {
				hasHref = true
				break
			}
		}
		if hasHref {
			marked := false
			for i := range root.Attr {
				if root.Attr[i].Namespace == "" && root.Attr[i].Key == "data-gfm-had-href" {
					root.Attr[i].Val = marker
					marked = true
					break
				}
			}
			if !marked {
				root.Attr = append(root.Attr, html.Attribute{Key: "data-gfm-had-href", Val: marker})
			}
		}
	}
	for child := domChildren(root); child != nil; child = child.NextSibling {
		if err := domMark(child, marker, depth+1); err != nil {
			return err
		}
	}
	return nil
}

func domSanitize(source, marker string) (string, error) {
	document, err := parseDOM(source, false, true)
	if err != nil {
		return "", err
	}
	if err := domAnchors(document.root, document, make([]bool, len(source)+1), 0); err != nil {
		return "", err
	}
	domTrim(document)
	svg := false
	if err := domSVG(document.root, &svg, 0); err != nil {
		return "", err
	}
	if svg {
		serialized, err := domSerialize(document.root, false)
		if err != nil {
			return "", err
		}
		document, err = parseDOM(serialized, false, false)
		if err != nil {
			return "", err
		}
	}
	if err := domMark(document.root, marker, 0); err != nil {
		return "", err
	}
	return sanitizeDOM(document.root, marker)
}
