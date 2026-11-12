package lint

import (
	"sort"
	"strings"
	"unicode/utf8"

	"ghmd/internal/analysis"
	"ghmd/internal/github"
	"ghmd/internal/standard"
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
	counts       analysis.FeatureSet
	descriptions map[featureID]featureDescription
	source       string
	footnotes    []bool
}

func newSemanticFeatures() semanticFeatures {
	return semanticFeatures{
		counts:       make(analysis.FeatureSet),
		descriptions: make(map[featureID]featureDescription),
	}
}

func (f *semanticFeatures) add(feature analysis.Feature) {
	rule, key, search, message := string(feature.Kind), feature.Value, feature.Search, feature.Message

	switch rule {
	case ruleLostText:
		if textSourceNoise(search) || strings.ContainsRune(search, '$') || strings.Contains(f.source, "~~") ||
			(strings.Contains(f.source, "[^") && strings.HasPrefix(search, "[")) {
			// Dollar remnants are math-rule territory; strikethrough delimiter
			// differences are reported by formatting rules.
			return
		}
	case ruleLostCode:
		text := search
		if text == "" || strings.ContainsRune(text, '$') {
			return
		}
		if strings.Contains(text, "|") && strings.Contains(f.source, "|") {
			return
		}
		kind, _, _ := strings.Cut(key, "\x00")
		if kind == "inline" && (strings.Contains(f.source, "$`"+text+"`$") || strings.Contains(f.source, "`$"+text+"$`") ||
			containsMathCode(f.source, text)) {
			return
		}
	case ruleLostLink:
		sourceLower := strings.ToLower(f.source)
		if strings.Contains(sourceLower, "<a") && !strings.Contains(sourceLower, "</a>") {
			return
		}
		if _, unsafe := unsafeURL("a", "href", key); unsafe {
			return
		}
	case ruleLostImage:
		if strings.Contains(f.source, "[^") {
			return
		}
		src, _, _ := strings.Cut(key, "\x00")
		if _, unsafe := unsafeURL("img", "src", src); unsafe || strings.HasPrefix(strings.ToLower(src), "data:") {
			return
		}
	case ruleTaskState:
		if strings.Contains(strings.ToLower(f.source), "<input") {
			return
		}
	}

	if (rule == ruleLostText || rule == ruleLostCode || rule == ruleLostLink || rule == ruleLostImage || rule == ruleLostMath) &&
		f.onlyInFootnotes(search) {
		return
	}
	id := featureID{rule: rule, key: key}
	f.counts.Add(feature)
	if _, exists := f.descriptions[id]; !exists {
		f.descriptions[id] = featureDescription{search: search, message: message}
	}
}

func differential(source string) ([]Finding, error) {
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
	features.footnotes = mathFootnoteLines(source, lineStartOffsets(source))
	extracted, err := analysis.SemanticFeatures(fragment)
	if err != nil {
		return features, err
	}
	for _, feature := range extracted {
		features.add(feature)
	}
	return features, nil
}

// textSourceNoise filters Markdown remnants from otherwise meaningful text.
// The shared extractor has already rejected empty and punctuation-only text.
func textSourceNoise(text string) bool {
	if strings.HasPrefix(text, "[ ]") || strings.HasPrefix(text, "[x]") || strings.HasPrefix(text, "[X]") {
		return true
	}
	return strings.Contains(text, "://") || strings.Contains(text, "]( ") || strings.Contains(text, "](") ||
		strings.Contains(text, "|")
}

func compareFeatures(source string, want, got semanticFeatures) []Finding {
	gotCounts := make(map[featureID]int, len(got.counts))
	for _, feature := range got.counts {
		gotCounts[featureID{rule: string(feature.Kind), key: feature.Value}] = feature.Count
	}
	type wantedFeature struct {
		id    featureID
		count int
	}
	wanted := make([]wantedFeature, 0, len(want.counts))
	for _, feature := range want.counts {
		wanted = append(wanted, wantedFeature{
			id:    featureID{rule: string(feature.Kind), key: feature.Value},
			count: feature.Count,
		})
	}
	sort.Slice(wanted, func(i, j int) bool {
		if wanted[i].id.rule != wanted[j].id.rule {
			return wanted[i].id.rule < wanted[j].id.rule
		}
		return wanted[i].id.key < wanted[j].id.key
	})

	starts := lineStartOffsets(source)
	var findings []Finding
	for _, feature := range wanted {
		missing := feature.count - gotCounts[feature.id]
		for range missing {
			description := want.descriptions[feature.id]
			at, located := locateFeature(source, starts, feature.id.rule, description.search)
			message := description.message
			if !located {
				at = sourcePosition{}
				message += " (unlocated)"
			}
			findings = append(findings, Finding{
				Line: at.line, Column: at.column, Severity: Warning, Rule: feature.id.rule, Message: message,
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

func normalizedText(value string) string {
	return strings.Join(strings.Fields(value), " ")
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

// onlyInFootnotes reports whether every occurrence of value lies inside a
// footnote definition, where md4c's unsupported footnote syntax makes its
// output unusable as an expected baseline.
func (f *semanticFeatures) onlyInFootnotes(value string) bool {
	if value == "" || len(f.footnotes) == 0 {
		return false
	}
	starts := lineStartOffsets(f.source)
	found := false
	for offset := 0; offset < len(f.source); {
		index := strings.Index(f.source[offset:], value)
		if index < 0 {
			break
		}
		index += offset
		line := sort.Search(len(starts), func(i int) bool { return starts[i] > index }) - 1
		if line < 0 || line >= len(f.footnotes) || !f.footnotes[line] {
			return false
		}
		found = true
		offset = index + len(value)
	}
	return found
}
