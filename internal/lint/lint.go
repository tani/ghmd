// Package lint reports source constructs that GitHub's renderer removes, rewrites,
// or allows to interfere with surrounding Markdown.
package lint

import (
	"fmt"
	"html"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	cmarkgfm "gfm/internal/cmark-gfm"
)

const (
	ruleRemovedContainer = "raw-html/removed-container"
	ruleUnsafeAttribute  = "raw-html/unsafe-attribute"
	ruleUnsafeURL        = "raw-html/unsafe-url"
	ruleInput            = "raw-html/input"
	ruleTemplate         = "raw-html/template"
	ruleNestedAnchor     = "html/nested-anchor"
	ruleBlockInP         = "html/block-in-p"
	ruleUnclosedTag      = "html/unclosed-tag"
	ruleStrayClose       = "html/stray-close"
)

// Severity is the seriousness of a finding.
type Severity int

const (
	Error Severity = iota
	Warning
)

func (s Severity) String() string {
	switch s {
	case Error:
		return "error"
	case Warning:
		return "warning"
	default:
		return "unknown"
	}
}

// Finding describes one source location that needs attention.
type Finding struct {
	Line, Column  int
	Severity      Severity
	Rule, Message string
}

func (f Finding) String() string {
	return fmt.Sprintf("%d:%d: %s %s: %s", f.Line, f.Column, f.Severity, f.Rule, f.Message)
}

// Source lints raw HTML and reports semantic regressions from the standard
// Markdown rendering.
func Source(source string) ([]Finding, error) {
	findings, err := rawHTML(source)
	if err != nil {
		return nil, err
	}
	differences, err := differential(source)
	if err != nil {
		return nil, err
	}
	findings = append(findings, differences...)
	findings = append(findings, mathFindings(source)...)
	findings = dedupeMathCauses(findings)
	sortFindings(findings)
	return findings, nil
}

func rawHTML(source string) ([]Finding, error) {
	document, err := cmarkgfm.Parse(source)
	if err != nil {
		return nil, err
	}
	defer document.Close()

	s := htmlScanner{source: source, lineStarts: lineStartOffsets(source)}
	s.walk(document.Root())
	for _, open := range s.open {
		s.add(open.at, Warning, ruleUnclosedTag,
			fmt.Sprintf("unclosed <%s> can swallow following Markdown", open.name))
	}
	return s.findings, nil
}

func sortFindings(findings []Finding) {
	sort.SliceStable(findings, func(i, j int) bool {
		a, b := findings[i], findings[j]
		if a.Line != b.Line {
			return a.Line < b.Line
		}
		if a.Column != b.Column {
			return a.Column < b.Column
		}
		return a.Rule < b.Rule
	})
}

type sourcePosition struct {
	line, column int
}

type openTag struct {
	name string
	at   sourcePosition
}

type htmlScanner struct {
	source     string
	lineStarts []int
	findings   []Finding
	open       []openTag
	rawTag     string

	literal    string
	offset     int
	line       int
	byteColumn int
}

func lineStartOffsets(source string) []int {
	starts := []int{0}
	for i := range len(source) {
		if source[i] == '\n' {
			starts = append(starts, i+1)
		}
	}
	return starts
}

func (s *htmlScanner) walk(node cmarkgfm.Node) {
	for ; node.Valid(); node = node.Next() {
		if node.Type() == "html_inline" || node.Type() == "html_block" {
			s.scanHTML(node)
		}
		if child := node.FirstChild(); child.Valid() {
			s.walk(child)
		}
	}
}

func (s *htmlScanner) scanHTML(node cmarkgfm.Node) {
	s.literal = node.Literal()
	s.offset = 0
	s.line = node.StartLine()
	s.byteColumn = node.StartColumn()
	if s.line < 1 {
		s.line = 1
	}
	if s.byteColumn < 1 {
		s.byteColumn = 1
	}

	for s.offset < len(s.literal) {
		if s.rawTag != "" {
			if s.rawTag == "plaintext" {
				s.advance(len(s.literal))
				continue
			}
			end := rawEndTag(s.literal, s.offset, s.rawTag)
			if end < 0 {
				s.advance(len(s.literal))
				continue
			}
			s.advance(end)
			s.rawTag = ""
			continue
		}

		if s.literal[s.offset] != '<' {
			s.advance(s.offset + 1)
			continue
		}
		rest := s.literal[s.offset:]
		switch {
		case strings.HasPrefix(rest, "<!--"):
			end := strings.Index(rest[4:], "-->")
			if end < 0 {
				s.advance(len(s.literal))
			} else {
				s.advance(s.offset + 4 + end + 3)
			}
		case strings.HasPrefix(rest, "<!"), strings.HasPrefix(rest, "<?"):
			end := strings.IndexByte(rest[2:], '>')
			if end < 0 {
				s.advance(len(s.literal))
			} else {
				s.advance(s.offset + 2 + end + 1)
			}
		default:
			s.scanTag()
		}
	}
}

func (s *htmlScanner) scanTag() {
	start := s.offset
	closing := start+1 < len(s.literal) && s.literal[start+1] == '/'
	p := start + 1
	if closing {
		p++
	}
	if p >= len(s.literal) || !asciiLetter(s.literal[p]) {
		s.advance(start + 1)
		return
	}
	nameStart := p
	for p < len(s.literal) && tagNameByte(s.literal[p]) {
		p++
	}
	if p < len(s.literal) && !asciiSpace(s.literal[p]) && s.literal[p] != '/' && s.literal[p] != '>' {
		s.advance(start + 1)
		return
	}

	name := strings.ToLower(s.literal[nameStart:p])
	at := s.position()
	if closing {
		end := strings.IndexByte(s.literal[p:], '>')
		if end < 0 {
			s.advance(len(s.literal))
			return
		}
		p += end + 1
		s.advance(p)
		s.closeTag(name, at)
		return
	}

	findingCount := len(s.findings)
	selfClosing := false
	complete := false
	for p < len(s.literal) {
		for p < len(s.literal) && asciiSpace(s.literal[p]) {
			p++
		}
		if p >= len(s.literal) {
			break
		}
		if s.literal[p] == '>' {
			p++
			complete = true
			break
		}
		if s.literal[p] == '/' {
			if p+1 < len(s.literal) && s.literal[p+1] == '>' {
				selfClosing = true
				complete = true
				p += 2
				break
			}
			p++
			continue
		}

		attributeStart := p
		for p < len(s.literal) && !asciiSpace(s.literal[p]) &&
			s.literal[p] != '=' && s.literal[p] != '/' && s.literal[p] != '>' {
			p++
		}
		if p == attributeStart {
			p++
			continue
		}
		attribute := strings.ToLower(s.literal[attributeStart:p])
		for p < len(s.literal) && asciiSpace(s.literal[p]) {
			p++
		}
		value := ""
		hasValue := false
		if p < len(s.literal) && s.literal[p] == '=' {
			hasValue = true
			p++
			for p < len(s.literal) && asciiSpace(s.literal[p]) {
				p++
			}
			if p < len(s.literal) && (s.literal[p] == '\'' || s.literal[p] == '"') {
				quote := s.literal[p]
				p++
				valueStart := p
				for p < len(s.literal) && s.literal[p] != quote {
					p++
				}
				value = s.literal[valueStart:p]
				if p < len(s.literal) {
					p++
				}
			} else {
				valueStart := p
				for p < len(s.literal) && !asciiSpace(s.literal[p]) && s.literal[p] != '>' {
					p++
				}
				value = s.literal[valueStart:p]
			}
		}
		s.checkAttribute(name, attribute, value, hasValue, at)
	}
	if !complete {
		s.findings = s.findings[:findingCount]
		s.advance(len(s.literal))
		return
	}

	s.openTag(name, selfClosing, at)
	s.advance(p)
}

func (s *htmlScanner) checkAttribute(tag, attribute, value string, hasValue bool, at sourcePosition) {
	if strings.HasPrefix(attribute, "on") || word(attribute, "srcdoc formaction is xmlns") || strings.HasPrefix(attribute, "xmlns:") {
		s.add(at, Error, ruleUnsafeAttribute, fmt.Sprintf("unsafe attribute %q", attribute))
	}
	if !hasValue || !word(attribute, "href src action poster background cite") {
		return
	}
	if scheme, unsafe := unsafeURL(tag, attribute, value); unsafe {
		s.add(at, Error, ruleUnsafeURL, fmt.Sprintf("unsafe %s URL in attribute %q", scheme, attribute))
	}
}

func unsafeURL(tag, attribute, value string) (string, bool) {
	normalized := strings.Map(func(r rune) rune {
		if r <= 0x20 || r == 0x7f || unicode.IsSpace(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, html.UnescapeString(value))
	for _, scheme := range [...]string{"javascript", "vbscript", "file", "blob"} {
		if strings.HasPrefix(normalized, scheme+":") {
			return scheme + ":", true
		}
	}
	if !strings.HasPrefix(normalized, "data:") {
		return "", false
	}
	if (tag == "img" || tag == "source") && (attribute == "src" || attribute == "srcset") {
		mime := normalized[len("data:"):]
		if end := strings.IndexAny(mime, ";,"); end >= 0 {
			mime = mime[:end]
		}
		if word(mime, "image/png image/jpeg image/gif image/webp image/avif image/bmp image/x-icon image/svg+xml") {
			return "", false
		}
	}
	return "data:", true
}

func (s *htmlScanner) openTag(name string, selfClosing bool, at sourcePosition) {
	if word(name, "script style iframe object embed base link meta svg title xmp noembed noframes plaintext foreignobject annotation annotation-xml") {
		s.add(at, Error, ruleRemovedContainer, fmt.Sprintf("GitHub removes <%s> and its contents", name))
	}
	if name == "input" {
		s.add(at, Warning, ruleInput, "GitHub removes source-authored <input> tags")
	}
	if name == "template" {
		s.add(at, Warning, ruleTemplate, "<template> contents are unwrapped and re-parented")
	}
	if name == "a" && s.tagOpen("a") {
		s.add(at, Warning, ruleNestedAnchor, "nested <a> element")
	}
	if s.tagOpen("p") && word(name, "address article aside blockquote details dialog div dl fieldset figcaption figure footer form h1 h2 h3 h4 h5 h6 header hgroup hr main menu nav ol p pre section table ul") {
		s.add(at, Warning, ruleBlockInP, fmt.Sprintf("block element <%s> inside <p>", name))
	}
	if selfClosing || word(name, "area base br col embed hr img input link meta param source track wbr") {
		return
	}
	s.open = append(s.open, openTag{name: name, at: at})
	if word(name, "script style textarea title xmp iframe noembed noframes plaintext") {
		s.rawTag = name
	}
}

func (s *htmlScanner) closeTag(name string, at sourcePosition) {
	for i := len(s.open) - 1; i >= 0; i-- {
		if s.open[i].name == name {
			s.open = s.open[:i]
			return
		}
	}
	s.add(at, Warning, ruleStrayClose, fmt.Sprintf("stray closing tag </%s>", name))
}

func (s *htmlScanner) tagOpen(name string) bool {
	for i := len(s.open) - 1; i >= 0; i-- {
		if s.open[i].name == name {
			return true
		}
	}
	return false
}

func (s *htmlScanner) add(at sourcePosition, severity Severity, rule, message string) {
	s.findings = append(s.findings, Finding{
		Line: at.line, Column: at.column, Severity: severity, Rule: rule, Message: message,
	})
}

func (s *htmlScanner) position() sourcePosition {
	line := s.line
	if line > len(s.lineStarts) {
		return sourcePosition{line: line, column: s.byteColumn}
	}
	start := s.lineStarts[line-1]
	absolute := start + s.byteColumn - 1
	if absolute < start {
		absolute = start
	}
	if absolute > len(s.source) {
		absolute = len(s.source)
	}
	return sourcePosition{line: line, column: utf8.RuneCountInString(s.source[start:absolute]) + 1}
}

func (s *htmlScanner) advance(to int) {
	for s.offset < to {
		c := s.literal[s.offset]
		s.offset++
		if c == '\n' {
			s.line++
			s.byteColumn = 1
		} else {
			s.byteColumn++
		}
	}
}

func rawEndTag(literal string, offset int, name string) int {
	for offset+1 < len(literal) {
		relative := strings.Index(literal[offset:], "</")
		if relative < 0 {
			return -1
		}
		candidate := offset + relative
		nameStart := candidate + 2
		nameEnd := nameStart + len(name)
		if nameEnd <= len(literal) && strings.EqualFold(literal[nameStart:nameEnd], name) &&
			nameEnd < len(literal) && (asciiSpace(literal[nameEnd]) || literal[nameEnd] == '/' || literal[nameEnd] == '>') {
			return candidate
		}
		offset = nameStart
	}
	return -1
}

func asciiLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func tagNameByte(c byte) bool {
	return asciiLetter(c) || c >= '0' && c <= '9' || c == '-' || c == ':' || c == '_'
}

func asciiSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f'
}

func word(value, words string) bool {
	for len(words) > 0 {
		end := strings.IndexByte(words, ' ')
		if end < 0 {
			return value == words
		}
		if value == words[:end] {
			return true
		}
		words = words[end+1:]
	}
	return false
}

// dedupeMathCauses drops a boundary-level math finding when the standard
// renderer comparison already reports that GitHub lost the same expression on
// the same line. Syntax problems that md4c also rejects (unclosed dollars,
// non-isolated display delimiters) remain visible.
func dedupeMathCauses(findings []Finding) []Finding {
	lostLines := make(map[int]bool)
	for _, finding := range findings {
		if finding.Rule == ruleLostMath {
			lostLines[finding.Line] = true
		}
	}
	out := findings[:0]
	for _, finding := range findings {
		if finding.Rule == ruleMathNotRecognized && lostLines[finding.Line] {
			continue
		}
		out = append(out, finding)
	}
	return out
}
