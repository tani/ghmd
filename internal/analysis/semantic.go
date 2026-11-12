package analysis

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/net/html"
)

// FeatureKind identifies a category of semantic content.
type FeatureKind string

const (
	featureText       FeatureKind = "standard/lost-text"
	featureFormatting FeatureKind = "standard/lost-formatting"
	featureStructure  FeatureKind = "standard/lost-structure"
	featureCode       FeatureKind = "standard/lost-code"
	featureLink       FeatureKind = "standard/lost-link"
	featureImage      FeatureKind = "standard/lost-image"
	featureMath       FeatureKind = "standard/lost-math"
	featureTaskState  FeatureKind = "standard/task-state"
)

// Feature describes semantic content independently of its source markup.
// Value is the identity within Kind; Search and Message describe the content.
type Feature struct {
	Kind    FeatureKind
	Value   string
	Search  string
	Message string
	Count   int
}

// FeatureSet contains features keyed by both kind and value.
type FeatureSet map[string]Feature

// Add inserts a feature or merges its count with an existing feature having the
// same kind and value. The first Search and Message values are retained. A zero
// Count denotes one occurrence.
func (features FeatureSet) Add(feature Feature) {
	if feature.Count == 0 {
		feature.Count = 1
	}
	key := string(feature.Kind) + "\x00" + feature.Value
	if current, exists := features[key]; exists {
		feature.Count += current.Count
		feature.Search = current.Search
		feature.Message = current.Message
	}
	features[key] = feature
}

// SemanticFeatures extracts normalized semantic content from an HTML fragment.
// It does not inspect Markdown source or suppress renderer-specific source noise.
func SemanticFeatures(source string) (FeatureSet, error) {
	features := make(FeatureSet)
	nodes, err := Paragraphs(source)
	if err != nil {
		return features, err
	}
	for _, node := range nodes {
		features.walk(node)
	}
	features.addVisibleText(nodes)
	return features, nil
}

func (features FeatureSet) walk(node *html.Node) {
	if node.Type != html.ElementNode {
		return
	}

	switch node.Data {
	case "script", "style", "iframe", "noembed", "noframes", "template", "head", "title", "svg":
		return
	case "math-renderer", "x-equation":
		expression := normalizeMath(textContents(node))
		if expression != "" {
			features.Add(Feature{
				Kind: featureMath, Value: expression, Search: expression,
				Message: fmt.Sprintf("standard math expression %q is missing or changed", expression),
			})
		}
		return
	case "pre":
		features.addCode(node, "preformatted")
		return
	case "code":
		features.addCode(node, "inline")
		return
	case "strong", "b":
		features.addFormatting(node, "strong")
	case "em", "i":
		features.addFormatting(node, "em")
	case "del", "s", "strike":
		features.addFormatting(node, "del")
	case "table":
		features.addTable(node)
	}

	switch node.Data {
	case "a":
		href, exists := attribute(node, "href")
		if !exists {
			break
		}
		text := normalizedText(textContents(node))
		if !strings.HasPrefix(text, "^") && href != text && stripScheme(href) != text {
			features.Add(Feature{
				Kind: featureLink, Value: href, Search: href,
				Message: fmt.Sprintf("standard link %q is missing or changed", href),
			})
		}
	case "img":
		src, hasSource := attribute(node, "src")
		alt, hasAlt := attribute(node, "alt")
		if hasSource || hasAlt {
			search := src
			if search == "" {
				search = alt
			}
			features.Add(Feature{
				Kind: featureImage, Value: src + "\x00" + alt, Search: search,
				Message: fmt.Sprintf("standard image src %q alt %q is missing or changed", src, alt),
			})
		}
	case "input":
		inputType, _ := attribute(node, "type")
		if strings.EqualFold(inputType, "checkbox") {
			_, checked := attribute(node, "checked")
			search, state := "[ ]", "unchecked"
			if checked {
				search, state = "[x]", "checked"
			}
			features.Add(Feature{
				Kind: featureTaskState, Value: state, Search: search,
				Message: fmt.Sprintf("standard %s task state is missing or changed", state),
			})
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		features.walk(child)
	}
}

func (features FeatureSet) addVisibleText(nodes []*html.Node) {
	current := &strings.Builder{}
	flush := func() {
		for _, text := range visibleTextFragments(current.String()) {
			features.Add(Feature{
				Kind: featureText, Value: text, Search: text,
				Message: fmt.Sprintf("standard text %q is missing or changed", text),
			})
		}
		current.Reset()
	}
	afterCode := false
	var walk func(*html.Node)
	walk = func(node *html.Node) {
		if node.Type == html.TextNode {
			data := node.Data
			if afterCode {
				data = strings.TrimPrefix(data, "$")
				afterCode = false
			}
			current.WriteString(data)
			return
		}
		if node.Type != html.ElementNode {
			return
		}
		if _, generated := attribute(node, "data-footnote-ref"); generated {
			flush()
			return
		}
		if _, generated := attribute(node, "data-footnote-backref"); generated {
			flush()
			return
		}
		if node.Data == "a" && strings.HasPrefix(normalizedText(textContents(node)), "^") {
			flush()
			return
		}
		switch node.Data {
		case "script", "style", "iframe", "noembed", "noframes", "template", "head", "title", "svg":
			return
		case "code":
			trimmed := strings.TrimSuffix(current.String(), "$")
			current.Reset()
			current.WriteString(trimmed)
			flush()
			afterCode = true
			return
		case "br":
			afterCode = false
			current.WriteByte(' ')
			return
		case "pre", "math-renderer", "x-equation", "hr":
			afterCode = false
			flush()
			return
		}
		if textBlock(node.Data) {
			flush()
			parent := current
			current = &strings.Builder{}
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				walk(child)
			}
			flush()
			current = parent
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			walk(child)
		}
	}
	for _, node := range nodes {
		walk(node)
	}
	flush()
}

func textBlock(tag string) bool {
	switch tag {
	case "address", "article", "aside", "blockquote", "body", "caption", "dd", "details", "dialog", "div", "dl", "dt",
		"fieldset", "figcaption", "figure", "footer", "form", "h1", "h2", "h3", "h4", "h5", "h6", "header", "hgroup",
		"li", "main", "menu", "nav", "ol", "p", "section", "summary", "table", "tbody", "td", "tfoot", "th", "thead", "tr", "ul":
		return true
	}
	return false
}

func (features FeatureSet) addCode(node *html.Node, kind string) {
	if mathCode(node) {
		return
	}
	text := textContents(node)
	if text == "" {
		return
	}
	features.Add(Feature{
		Kind: featureCode, Value: kind + "\x00" + text, Search: text,
		Message: fmt.Sprintf("standard %s code %q is missing or changed", kind, text),
	})
}

func (features FeatureSet) addFormatting(node *html.Node, kind string) {
	text := normalizedText(textContents(node))
	if text == "" {
		return
	}
	features.Add(Feature{
		Kind: featureFormatting, Value: kind, Search: text,
		Message: fmt.Sprintf("standard %s formatting is missing or changed", kind),
	})
}

func (features FeatureSet) addTable(node *html.Node) {
	rows, columns := 0, 0
	var visit func(*html.Node)
	visit = func(current *html.Node) {
		if current.Type != html.ElementNode {
			return
		}
		if current.Data == "tr" {
			rows++
			count := 0
			for child := current.FirstChild; child != nil; child = child.NextSibling {
				if child.Type == html.ElementNode && (child.Data == "td" || child.Data == "th") {
					count++
				}
			}
			if count > columns {
				columns = count
			}
		}
		for child := current.FirstChild; child != nil; child = child.NextSibling {
			visit(child)
		}
	}
	visit(node)
	if rows == 0 && columns == 0 {
		return
	}
	shape := fmt.Sprintf("table:%d:%d", rows, columns)
	features.Add(Feature{
		Kind: featureStructure, Value: shape, Search: firstCellText(node),
		Message: fmt.Sprintf("standard table shape %d rows x %d columns is missing or changed", rows, columns),
	})
}

func firstCellText(node *html.Node) string {
	var visit func(*html.Node) string
	visit = func(current *html.Node) string {
		if current.Type == html.ElementNode && (current.Data == "td" || current.Data == "th") {
			if text := normalizedText(textContents(current)); meaningfulText(text) {
				return text
			}
		}
		for child := current.FirstChild; child != nil; child = child.NextSibling {
			if text := visit(child); text != "" {
				return text
			}
		}
		return ""
	}
	return visit(node)
}

// normalizeMath removes delimiters used by equivalent math element encodings.
func normalizeMath(value string) string {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "$$") && strings.HasSuffix(value, "$$") && len(value) >= 4 {
		value = value[2 : len(value)-2]
	} else if strings.HasPrefix(value, "$") && strings.HasSuffix(value, "$") && len(value) >= 2 {
		value = value[1 : len(value)-1]
	}
	value = strings.ReplaceAll(value, "`", "")
	return strings.TrimSpace(value)
}

func attribute(node *html.Node, name string) (string, bool) {
	return Attr(node, name)
}

func textContents(node *html.Node) string {
	var text strings.Builder
	Walk(node, func(current *html.Node) {
		if current.Type == html.TextNode {
			text.WriteString(current.Data)
		}
	})
	return text.String()
}

func normalizedText(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// visibleTextFragments separates embedded syntactic tokens from visible text.
// The returned fragments retain their content; source-specific suppression is
// left to callers comparing feature sets.
func visibleTextFragments(value string) []string {
	var fragments []string
	start := 0
	add := func(end int) {
		text := normalizedText(value[start:end])
		if text != "" && meaningfulText(text) {
			fragments = append(fragments, text)
		}
	}

	for i := 0; i < len(value); {
		end := -1
		switch {
		case value[i] == '$':
			if strings.HasPrefix(value[i:], "$$\n") {
				if close := strings.Index(value[i+3:], "\n$$"); close >= 0 {
					end = i + 3 + close + 3
				}
			} else if close := inlineMathEnd(value, i); close > 0 {
				end = close
			}
		case strings.HasPrefix(value[i:], "[^"):
			if close := strings.IndexByte(value[i+2:], ']'); close >= 0 {
				end = i + 2 + close + 1
				if end < len(value) && value[end] == ':' {
					end++
				}
			}
		case strings.HasPrefix(value[i:], "[!"):
			if close := strings.IndexByte(value[i+2:], ']'); close >= 0 {
				candidateEnd := i + 2 + close + 1
				if alertMarker(value[i:candidateEnd]) {
					end = candidateEnd
				}
			}
		}
		if end < 0 {
			i++
			continue
		}
		add(i)
		i = end
		start = i
	}
	add(len(value))
	return fragments
}

func alertMarker(value string) bool {
	if len(value) < 4 || !strings.HasPrefix(value, "[!") || value[len(value)-1] != ']' {
		return false
	}
	for i := 2; i < len(value)-1; i++ {
		if !asciiLetter(value[i]) {
			return false
		}
	}
	return true
}

func asciiLetter(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z'
}

// inlineMathEnd recognizes inline math boundaries without treating ordinary
// currency text as math.
func inlineMathEnd(value string, start int) int {
	if start+1 >= len(value) || value[start+1] == '$' || value[start+1] == '\n' {
		return -1
	}
	end := start + 1
	for end < len(value) && value[end] != '$' && value[end] != '\n' {
		end++
	}
	if end == len(value) || value[end] != '$' || end == start+1 {
		return -1
	}
	before := runeBefore(value, start)
	after := runeAt(value, end+1)
	first := runeAt(value, start+1)
	last := runeBefore(value, end)
	invalidAfter := after >= 'a' && after <= 'z' || after >= 'A' && after <= 'Z' ||
		after >= '0' && after <= '9' || after == '_'
	if start != 0 && !isJSWhitespace(before) && before != '(' || invalidAfter ||
		isJSWhitespace(first) || isJSWhitespace(last) && after != '$' ||
		strings.ContainsRune(value[start+1:end], '&') {
		return -1
	}
	return end + 1
}

func runeAt(value string, offset int) rune {
	if offset < 0 || offset >= len(value) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(value[offset:])
	return r
}

func runeBefore(value string, offset int) rune {
	if offset <= 0 {
		return 0
	}
	r, _ := utf8.DecodeLastRuneInString(value[:offset])
	return r
}

func isJSWhitespace(r rune) bool {
	return r >= 9 && r <= 13 || r == 0x20 || r == 0xa0 || r == 0x1680 ||
		r >= 0x2000 && r <= 0x200a || r == 0x2028 || r == 0x2029 || r == 0x202f ||
		r == 0x205f || r == 0x3000 || r == 0xfeff
}

// meaningfulText rejects empty and punctuation-only text without interpreting
// Markdown syntax.
func meaningfulText(text string) bool {
	letters, runes := 0, 0
	for _, r := range text {
		runes++
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			letters++
		}
	}
	return letters >= 2 && runes >= 3
}

func stripScheme(href string) string {
	if i := strings.Index(href, "://"); i >= 0 {
		return href[i+3:]
	}
	return href
}

func mathCode(node *html.Node) bool {
	if Class(node, "language-math") {
		return true
	}
	if node.Data == "pre" {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.ElementNode && child.Data == "code" && Class(child, "language-math") {
				return true
			}
		}
	}
	return false
}
