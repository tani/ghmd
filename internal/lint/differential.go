package lint

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"gfm/internal/github"
	"gfm/internal/standard"
)

const (
	differentialRunID    = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	differentialBoundary = "aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaaa"

	ruleLostText       = "standard/lost-text"
	ruleLostFormatting = "standard/lost-formatting"
	ruleLostStructure  = "standard/lost-structure"
	ruleLostCode       = "standard/lost-code"
	ruleLostLink       = "standard/lost-link"
	ruleLostImage      = "standard/lost-image"
	ruleLostMath       = "standard/lost-math"
	ruleTaskState      = "standard/task-state"
)

type featureID struct {
	rule string
	key  string
}

type featureDescription struct {
	search  string
	message string
}

type semanticFeatures struct {
	counts       map[featureID]int
	descriptions map[featureID]featureDescription
	source       string
}

func newSemanticFeatures() semanticFeatures {
	return semanticFeatures{
		counts:       make(map[featureID]int),
		descriptions: make(map[featureID]featureDescription),
	}
}

func (f *semanticFeatures) add(rule, key, search, message string) {
	id := featureID{rule: rule, key: key}
	f.counts[id]++
	if _, exists := f.descriptions[id]; !exists {
		f.descriptions[id] = featureDescription{search: search, message: message}
	}
}

func differential(source string) ([]Finding, error) {
	lowerSource := strings.ToLower(source)
	if strings.Contains(source, "[^") || strings.Contains(lowerSource, "<svg") {
		// md4c has no footnote support and SVG is foreign content; these
		// outputs are not a usable expected baseline. Raw-HTML rules still
		// apply to the same source.
		return nil, nil
	}
	githubHTML, err := github.RenderDocument(source, differentialRunID, differentialBoundary)
	if err != nil {
		return nil, err
	}
	standardHTML, err := standard.Render(source, true)
	if err != nil {
		return nil, err
	}
	got, err := extractFeatures(githubHTML, source)
	if err != nil {
		return nil, err
	}
	want, err := extractFeatures(standardHTML, source)
	if err != nil {
		return nil, err
	}
	return compareFeatures(source, want, got), nil
}

func extractFeatures(fragment, source string) (semanticFeatures, error) {
	features := newSemanticFeatures()
	features.source = source
	context := &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body}
	nodes, err := html.ParseFragment(strings.NewReader(fragment), context)
	if err != nil {
		return features, err
	}
	for _, node := range nodes {
		features.walk(node)
	}
	features.addVisibleText(nodes)
	return features, nil
}

func (f *semanticFeatures) walk(node *html.Node) {
	if node.Type == html.TextNode {
		return
	}
	if node.Type != html.ElementNode {
		return
	}

	switch node.Data {
	case "script", "style", "iframe", "noembed", "noframes", "template", "head", "title":
		return
	case "math-renderer", "x-equation":
		expression := normalizeMath(nodeText(node))
		if expression != "" {
			f.add(ruleLostMath, expression, expression, fmt.Sprintf("standard math expression %q is missing or changed", expression))
		}
		return
	case "pre":
		f.addCode(node, "preformatted")
		return
	case "code":
		f.addCode(node, "inline")
		return
	case "strong", "b":
		f.addFormatting(node, "strong")
	case "em", "i":
		f.addFormatting(node, "em")
	case "del", "s", "strike":
		f.addFormatting(node, "del")
	case "table":
		f.addTable(node)
	}

	switch node.Data {
	case "a":
		sourceLower := strings.ToLower(f.source)
		if strings.Contains(sourceLower, "<a") && !strings.Contains(sourceLower, "</a>") {
			break
		}
		if href, exists := attribute(node, "href"); exists && !strings.HasPrefix(normalizedText(nodeText(node)), "^") {
			text := normalizedText(nodeText(node))
			if _, unsafe := unsafeURL("a", "href", href); !unsafe && href != text && stripScheme(href) != text {
				f.add(ruleLostLink, href, href, fmt.Sprintf("standard link %q is missing or changed", href))
			}
		}
	case "img":
		src, hasSource := attribute(node, "src")
		alt, hasAlt := attribute(node, "alt")
		if _, unsafe := unsafeURL("img", "src", src); unsafe || strings.HasPrefix(strings.ToLower(src), "data:") {
			break
		}
		if hasSource || hasAlt {
			search := src
			if search == "" {
				search = alt
			}
			f.add(ruleLostImage, src+"\x00"+alt, search,
				fmt.Sprintf("standard image src %q alt %q is missing or changed", src, alt))
		}
	case "input":
		if strings.Contains(strings.ToLower(f.source), "<input") {
			break
		}
		if inputType, _ := attribute(node, "type"); strings.EqualFold(inputType, "checkbox") {
			_, checked := attribute(node, "checked")
			search, state := "[ ]", "unchecked"
			if checked {
				search, state = "[x]", "checked"
			}
			f.add(ruleTaskState, state, search, fmt.Sprintf("standard %s task state is missing or changed", state))
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		f.walk(child)
	}
}

func (f *semanticFeatures) addVisibleText(nodes []*html.Node) {
	current := &strings.Builder{}
	flush := func() {
		for _, text := range visibleTextFragments(current.String()) {
			if strings.ContainsRune(text, '$') || strings.Contains(f.source, "~~") {
				// Dollar remnants are math-rule territory; strikethrough
				// delimiter differences are reported by formatting rules.
				continue
			}
			f.add(ruleLostText, text, text, fmt.Sprintf("standard text %q is missing or changed", text))
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
		if node.Data == "a" && strings.HasPrefix(normalizedText(nodeText(node)), "^") {
			// md4c treats an unresolved footnote marker as a link; its visible
			// text is footnote syntax, not expected document text.
			flush()
			return
		}
		switch node.Data {
		case "script", "style", "iframe", "noembed", "noframes", "template", "head", "title":
			return
		case "code":
			// In the standard renderer, math written as $`expr`$ appears as
			// text, inline code, text. Drop the adjacent math delimiters.
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

func (f *semanticFeatures) addCode(node *html.Node, kind string) {
	if mathCode(node) {
		return
	}
	text := nodeText(node)
	if text == "" || strings.ContainsRune(text, '$') {
		return
	}
	if strings.Contains(text, "|") && strings.Contains(f.source, "|") {
		return
	}
	if kind == "inline" && (strings.Contains(f.source, "$`"+text+"`$") || strings.Contains(f.source, "`$"+text+"$`") ||
		containsMathCode(f.source, text)) {
		return
	}
	f.add(ruleLostCode, kind+"\x00"+text, text,
		fmt.Sprintf("standard %s code %q is missing or changed", kind, text))
}

func compareFeatures(source string, want, got semanticFeatures) []Finding {
	ids := make([]featureID, 0, len(want.counts))
	for id := range want.counts {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		if ids[i].rule != ids[j].rule {
			return ids[i].rule < ids[j].rule
		}
		return ids[i].key < ids[j].key
	})

	starts := lineStartOffsets(source)
	var findings []Finding
	for _, id := range ids {
		missing := want.counts[id] - got.counts[id]
		for range missing {
			description := want.descriptions[id]
			at, located := locateFeature(source, starts, id.rule, description.search)
			message := description.message
			if !located {
				at = sourcePosition{}
				message += " (unlocated)"
			}
			findings = append(findings, Finding{
				Line: at.line, Column: at.column, Severity: Warning, Rule: id.rule, Message: message,
			})
		}
	}
	return findings
}

func locateFeature(source string, starts []int, rule, value string) (sourcePosition, bool) {
	if rule == ruleLostMath {
		for _, candidate := range []string{"$" + value + "$", "$`" + value + "`$", "$$\n" + value + "\n$$"} {
			if at, found := exactUniqueLocation(source, starts, candidate); found {
				return at, true
			}
		}
	}
	if at, found := exactUniqueLocation(source, starts, value); found {
		return at, true
	}
	if rule == ruleTaskState && value == "[x]" {
		return exactUniqueLocation(source, starts, "[X]")
	}
	return sourcePosition{}, false
}

func exactUniqueLocation(source string, starts []int, value string) (sourcePosition, bool) {
	if value == "" {
		return sourcePosition{}, false
	}
	offset := strings.Index(source, value)
	if offset < 0 || strings.Contains(source[offset+1:], value) {
		return sourcePosition{}, false
	}
	lineIndex := sort.Search(len(starts), func(i int) bool { return starts[i] > offset }) - 1
	if lineIndex < 0 {
		return sourcePosition{}, false
	}
	return sourcePosition{
		line:   lineIndex + 1,
		column: utf8.RuneCountInString(source[starts[lineIndex]:offset]) + 1,
	}, true
}

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

func mathCode(node *html.Node) bool {
	if classContains(node, "language-math") {
		return true
	}
	if node.Data == "pre" {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.ElementNode && child.Data == "code" && classContains(child, "language-math") {
				return true
			}
		}
	}
	return false
}

func classContains(node *html.Node, class string) bool {
	value, exists := attribute(node, "class")
	if !exists {
		return false
	}
	for field := range strings.FieldsSeq(value) {
		if field == class {
			return true
		}
	}
	return false
}

func attribute(node *html.Node, name string) (string, bool) {
	for _, attribute := range node.Attr {
		if attribute.Namespace == "" && attribute.Key == name {
			return attribute.Val, true
		}
	}
	return "", false
}

func nodeText(node *html.Node) string {
	var text strings.Builder
	var walk func(*html.Node)
	walk = func(node *html.Node) {
		if node.Type == html.TextNode {
			text.WriteString(node.Data)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			walk(child)
		}
	}
	walk(node)
	return text.String()
}

func normalizedText(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

// inlineMathEnd recognizes the same inline math boundaries as the renderer,
// so currency-like text such as "$5 and $6" is not treated as math.
func inlineMathEnd(s string, i int) int {
	if i+1 >= len(s) || s[i+1] == '$' || s[i+1] == '\n' {
		return -1
	}
	end := i + 1
	for end < len(s) && s[end] != '$' && s[end] != '\n' {
		end++
	}
	if end == len(s) || s[end] != '$' || end == i+1 {
		return -1
	}
	before := runeBefore(s, i)
	after := runeAt(s, end+1)
	first := runeAt(s, i+1)
	last := runeBefore(s, end)
	invalidAfter := after >= 'a' && after <= 'z' || after >= 'A' && after <= 'Z' ||
		after >= '0' && after <= '9' || after == '_'
	if i != 0 && !github.IsJSWhitespace(before) && before != '(' || invalidAfter ||
		github.IsJSWhitespace(first) || github.IsJSWhitespace(last) && after != '$' ||
		strings.ContainsRune(s[i+1:end], '&') {
		return -1
	}
	return end + 1
}

func runeAt(s string, i int) rune {
	if i < 0 || i >= len(s) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(s[i:])
	return r
}

func runeBefore(s string, i int) rune {
	if i <= 0 {
		return 0
	}
	r, _ := utf8.DecodeLastRuneInString(s[:i])
	return r
}

// meaningfulText filters punctuation-only leftovers from dialect differences.
func meaningfulText(text string) bool {
	if strings.HasPrefix(text, "[ ]") || strings.HasPrefix(text, "[x]") || strings.HasPrefix(text, "[X]") {
		return false
	}
	if strings.Contains(text, "://") || strings.Contains(text, "]( ") || strings.Contains(text, "](") ||
		strings.Contains(text, "|") {
		return false
	}
	if strings.ContainsRune(text, '$') && strings.ContainsAny(text, `\^_{}`) {
		return false
	}
	letters := 0
	for _, r := range text {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			letters++
		}
	}
	return letters >= 2 && len([]rune(text)) >= 3
}

func stripScheme(href string) string {
	if i := strings.Index(href, "://"); i >= 0 {
		return href[i+3:]
	}
	return href
}

// containsMathCode reports inline code wrapped as $`...`$ in the source,
// allowing arbitrary whitespace between the markers and code text.
func containsMathCode(source, text string) bool {
	want := normalizedText(text)
	for offset := 0; ; {
		open := strings.Index(source[offset:], "$`")
		if open < 0 {
			return false
		}
		open += offset + 2
		close := strings.Index(source[open:], "`$")
		if close < 0 {
			return false
		}
		if normalizedText(source[open:open+close]) == want {
			return true
		}
		offset = open + close + 2
	}
}

// normalizeMath removes renderer-specific delimiters so cmark-gfm placeholders
// and md4c x-equation elements compare by expression.
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

func (f *semanticFeatures) addFormatting(node *html.Node, kind string) {
	text := normalizedText(nodeText(node))
	if text == "" {
		return
	}
	f.add(ruleLostFormatting, kind, text,
		fmt.Sprintf("standard %s formatting is missing or changed", kind))
}

func (f *semanticFeatures) addTable(node *html.Node) {
	rows, columns := 0, 0
	var visit func(*html.Node, int)
	visit = func(current *html.Node, depth int) {
		if current.Type == html.ElementNode {
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
				visit(child, depth+1)
			}
		}
	}
	visit(node, 0)
	if rows == 0 && columns == 0 {
		return
	}
	search := firstCellText(node)
	f.add(ruleLostStructure, fmt.Sprintf("table:%d:%d", rows, columns), search,
		fmt.Sprintf("standard table shape %d rows x %d columns is missing or changed", rows, columns))
}

func firstCellText(node *html.Node) string {
	var visit func(*html.Node) string
	visit = func(current *html.Node) string {
		if current.Type == html.ElementNode && (current.Data == "td" || current.Data == "th") {
			if text := normalizedText(nodeText(current)); meaningfulText(text) {
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
