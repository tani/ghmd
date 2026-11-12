package lint

import (
	"sort"
	"strings"
)

type mathPolicyLineRules struct {
	profile     mathPolicyProfile
	htmlBefore  []mathPolicyHTMLTag
	original    string
	codeMasked  string
	masked      string
	commentMask string
	line        int
	add         func(string, int, string)
}

func mathPolicyStyleViolations(source string) []styleViolation {
	lines := sourceLines(source)
	fenced, _ := fencedCodeRegions(lines)
	profile := mathPolicyDocumentProfile(lines, fenced)
	htmlStacks := mathPolicyHTMLStacks(lines, fenced)
	var out []styleViolation

	for lineIndex, original := range lines {
		if fenced[lineIndex] {
			continue
		}
		commentMasked := mathPolicyMaskComment(original)
		codeMasked := maskLineCode(commentMasked)
		masked := mathPolicyMaskHTMLTags(codeMasked)
		lineNumber := lineIndex + 1
		rules := mathPolicyLineRules{
			profile: profile, htmlBefore: htmlStacks[lineIndex], original: original,
			codeMasked: codeMasked, masked: masked, commentMask: commentMasked, line: lineNumber,
		}
		rules.add = func(id string, offset int, detail string) {
			if offset < 0 {
				offset = 0
			}
			if offset > len(original) {
				offset = len(original)
			}
			out = append(out, styleViolation{ID: id, Line: lineNumber,
				Column: styleColumn(original, original[:offset]), Detail: detail})
		}
		if profile == mathPolicyProfileUnspecified || profile == mathPolicyProfileRich {
			for _, at := range mathPolicyAlternateDelimiters(masked) {
				rules.add("MATH-003", at, "verify parser-level support for alternate LaTeX delimiters; prefer $`...`$ inline and fenced math display in github-rich")
			}
		}

		spans := mathPolicyDollarSpans(masked)
		spans = append(spans, mathPolicyDollarBacktickSpans(commentMasked)...)
		sort.SliceStable(spans, func(i, j int) bool { return spans[i].start < spans[j].start })
		for _, span := range spans {
			bodySource := codeMasked
			if span.code {
				bodySource = commentMasked
			}
			rules.checkSpan(span, bodySource[span.bodyStart:span.bodyEnd])
		}
		rules.adjacentDollars(spans)
		for _, at := range mathPolicyNestedTextDollars(codeMasked) {
			rules.add("MATH-008", at, "avoid nested dollar math inside \\text{...}; use supported TeX constructs without changing meaning")
		}
		for _, at := range mathPolicyDoubledDollarEscapes(codeMasked) {
			rules.add("MATH-009", at, "use ordinary TeX escapes such as \\$; do not compensate with doubled backslashes")
		}
	}
	return out
}

func (r *mathPolicyLineRules) checkSpan(span mathPolicySpan, body string) {
	if !span.code && (r.profile == mathPolicyProfileUnspecified || r.profile == mathPolicyProfileRich) {
		if span.display {
			r.add("MATH-002", span.start, "same-line $$...$$ display math is not isolated; put delimiters on separate lines or use a fenced math block")
		} else {
			r.add("MATH-002", span.start, "avoid bare $...$ in newly authored github-rich content; prefer dollar-backtick $`...`$ inline math")
		}
	}
	if span.display {
		return
	}
	checks := []struct {
		id    string
		check func() (detail string, ok bool)
	}{
		{"MATH-005", func() (string, bool) {
			return "inline math contains a TeX % comment; move comments and multiline expressions to a fenced math block", mathPolicyTeXComment(body) >= 0
		}},
		{"MATH-010", func() (string, bool) {
			return "verify whether \\displaystyle is semantically needed; do not force it merely to enlarge inline output", strings.Contains(body, `\displaystyle`)
		}},
		{"MATH-011", func() (string, bool) {
			return "inline math exceeds 80 bytes; split only at valid boundaries or preserve a readable code fallback", len(body) > 80
		}},
		{"DETAILS-002", func() (string, bool) {
			return "move equation widgets outside collapsible sections or test the actual initial expansion", mathPolicyDetailsOpen(r.htmlBefore)
		}},
		{"MATH-007", func() (string, bool) {
			nested := mathPolicyNestedContext(r.masked, span) ||
				mathPolicyEmphasisContext(r.masked, span) ||
				mathPolicyHTMLContext(r.original, span.start, r.htmlBefore) ||
				mathPolicyTableContext(r.masked, span)
			return "move math outside link labels, emphasis, table cells, and HTML containers, or verify the exact nested target structure", nested
		}},
	}
	for _, check := range checks {
		if detail, ok := check.check(); ok {
			r.add(check.id, span.start, detail)
		}
	}
}

func (r *mathPolicyLineRules) adjacentDollars(spans []mathPolicySpan) {
	for i := 1; i < len(spans); i++ {
		previous, current := spans[i-1], spans[i]
		if previous.display || current.display || previous.end != current.start {
			continue
		}
		r.add("MATH-006", previous.end-1, "adjacent closing/opening math dollars merge expressions; separate them with prose or whitespace")
	}
}
