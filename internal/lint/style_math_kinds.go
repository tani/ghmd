package lint

import "sort"

type mathPolicyLineRules struct {
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
			htmlBefore: htmlStacks[lineIndex], original: original,
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
		for _, at := range mathPolicyAlternateDelimiters(masked) {
			rules.add("MATH-001", at, "alternate LaTeX delimiters require parser-level support; prefer $`...`$ inline math or a fenced math block")
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
			rules.add("MATH-006", at, "avoid nested dollar math inside \\text{...}; use supported TeX constructs without changing meaning")
		}
		for _, at := range mathPolicyDoubledDollarEscapes(codeMasked) {
			rules.add("MATH-007", at, "use ordinary TeX escapes such as \\$; do not compensate with doubled backslashes")
		}
	}
	return append(out, mathPolicyBlockContextViolations(lines, htmlStacks)...)
}

func (r *mathPolicyLineRules) checkSpan(span mathPolicySpan, body string) {
	if span.display {
		return
	}
	checks := []struct {
		id    string
		check func() (detail string, ok bool)
	}{
		{"MATH-003", func() (string, bool) {
			return "inline math contains a TeX % comment; move comments and multiline expressions to a fenced math block", mathPolicyTeXComment(body) >= 0
		}},
		{"DETAILS-002", func() (string, bool) {
			return "move equation widgets outside collapsible sections or test the actual initial expansion", mathPolicyDetailsOpen(r.htmlBefore)
		}},
		{"MATH-005", func() (string, bool) {
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
		r.add("MATH-004", previous.end-1, "adjacent closing/opening math dollars merge expressions; separate them with prose or whitespace")
	}
}
