package lint

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	footnoteWikiPattern         = regexp.MustCompile(`(?i)\bwiki(?:s)?\b`)
	footnoteWikiMetadataPattern = regexp.MustCompile(`(?i)\bwiki\s*:\s*footnotes?\b`)
)

type footnoteToken struct {
	identifier string
	start      int
	line       int
	column     int
	endLine    int
	bodyColumn int
	definition bool
}

func footnoteViolations(source string) []styleViolation {
	lines := sourceLines(source)
	fenced, _ := fencedCodeRegions(lines)
	commentsMasked := coreMaskComments(lines, fenced)
	codeSpans := coreInlineCodeSpans(commentsMasked, fenced)
	maskedLines := coreMaskInlineCode(commentsMasked, codeSpans)

	rawDocument := strings.Join(lines, "\n")
	maskedDocument := strings.Join(maskedLines, "\n")
	lineStarts := footnoteLineStarts(lines)
	tokens := scanFootnoteTokens(rawDocument, maskedDocument, lines, lineStarts)
	linkLabels := footnoteLinkLabelRanges(maskedDocument)

	var definitions []footnoteToken
	exactDefinitions := make(map[string]footnoteToken)
	foldedDefinitions := make(map[string]footnoteToken)
	definitionOpenLine := make([]bool, len(lines))
	definitionStructureLine := make([]bool, len(lines))
	for _, token := range tokens {
		if !token.definition {
			continue
		}
		definitions = append(definitions, token)
		definitionOpenLine[token.line] = true
		for line := token.line; line <= token.endLine && line < len(lines); line++ {
			definitionStructureLine[line] = true
		}
		if _, exists := exactDefinitions[token.identifier]; !exists {
			exactDefinitions[token.identifier] = token
		}
		folded := strings.ToLower(token.identifier)
		if _, exists := foldedDefinitions[folded]; !exists {
			foldedDefinitions[folded] = token
		}
	}

	var out []styleViolation
	add := func(token footnoteToken, id, detail string) {
		out = append(out, styleViolation{
			ID: id, Line: token.line + 1, Column: coreColumn(lines[token.line], token.column), Detail: detail,
		})
	}
	for _, token := range tokens {
		compoundImage := footnoteCompoundImageLabel(token.identifier)
		if !validFootnoteIdentifier(token.identifier) && !compoundImage {
			kind := "reference"
			if token.definition {
				kind = "definition"
			}
			add(token, "FOOT-001", fmt.Sprintf("footnote %s identifier %q must match [a-z][a-z0-9-]*; use lowercase ASCII letters, digits, and hyphens only", kind, token.identifier))
		}
		if compoundImage {
			add(token, "FOOT-002", fmt.Sprintf("do not put image %q inside footnote-like brackets; write the image and footnote reference separately", token.identifier))
		}
		if token.definition {
			if token.endLine != token.line {
				add(token, "FOOT-002", fmt.Sprintf("footnote definition label %q contains a physical newline; keep the complete [^identifier]: label on one source line", token.identifier))
			}
			if first := exactDefinitions[token.identifier]; first.start != token.start {
				add(token, "FOOT-001", fmt.Sprintf("duplicate footnote definition %q; keep one definition for this identifier", token.identifier))
			}
			continue
		}
		if footnoteInsideLabels(token, linkLabels) {
			add(token, "FOOT-002", fmt.Sprintf("move footnote reference [^%s] outside the link or image label", token.identifier))
		}
		if _, exact := exactDefinitions[token.identifier]; exact {
			continue
		}
		if definition, found := foldedDefinitions[strings.ToLower(token.identifier)]; found && definition.identifier != token.identifier && strings.EqualFold(definition.identifier, token.identifier) {
			add(token, "FOOT-001", fmt.Sprintf("footnote reference [^%s] and definition [^%s] differ in case; use exactly the same spelling", token.identifier, definition.identifier))
		}
	}

	// A definition block may only be followed by blank lines, other definitions,
	// and indented continuation content. This suffix scan avoids confusing a
	// continuation paragraph with prose that moves a definition away from the end.
	topLevelProse := make([]bool, len(lines))
	for i, line := range lines {
		if fenced[i] || definitionStructureLine[i] || indentOf(line) != 0 || strings.TrimSpace(maskedLines[i]) == "" {
			continue
		}
		topLevelProse[i] = true
	}
	proseAfter := make([]bool, len(lines)+1)
	for i := len(lines) - 1; i >= 0; i-- {
		proseAfter[i] = proseAfter[i+1] || topLevelProse[i]
	}
	for _, definition := range definitions {
		if definition.endLine+1 < len(lines) && proseAfter[definition.endLine+1] {
			add(definition, "FOOT-001", fmt.Sprintf("move footnote definition [^%s] and its continuations to the end of the document; non-indented prose follows it", definition.identifier))
		}
	}

	for _, definition := range definitions {
		checkMath := func(line, base int) {
			if line < 0 || line >= len(lines) || base < 0 || base > len(lines[line]) {
				return
			}
			raw := lines[line][base:]
			masked := maskedLines[line][base:]
			if offset := footnoteMathOffset(raw, masked); offset >= 0 {
				out = append(out, styleViolation{
					ID: "FOOT-003", Line: line + 1, Column: coreColumn(lines[line], base+offset),
					Detail: fmt.Sprintf("move mathematical expression %q out of footnote [^%s]; give it ordinary paragraph context instead", raw[offset:], definition.identifier),
				})
			}
		}
		checkMath(definition.endLine, definition.bodyColumn)
		for line := definition.endLine + 1; line < len(lines); line++ {
			if definitionOpenLine[line] {
				break
			}
			if fenced[line] {
				continue
			}
			if strings.TrimSpace(maskedLines[line]) == "" {
				continue
			}
			if indentOf(lines[line]) == 0 {
				break
			}
			checkMath(line, 0)
		}
	}

	wikiLine, declarationLine := -1, -1
	for i, line := range maskedLines {
		if !footnoteWikiPattern.MatchString(line) {
			continue
		}
		if wikiLine < 0 {
			wikiLine = i
		}
		lower := strings.ToLower(line)
		if declarationLine < 0 && (footnoteWikiMetadataPattern.MatchString(line) ||
			strings.Contains(lower, "footnote") && strings.Contains(lower, "assumed to work in a wiki")) {
			declarationLine = i
		}
	}
	if wikiLine >= 0 {
		switch {
		case len(definitions) > 0:
			definition := definitions[0]
			add(definition, "FOOT-004", fmt.Sprintf("footnotes such as [^%s] are not supported in GitHub wikis; use an ordinary numbered references section", definition.identifier))
		case declarationLine >= 0:
			out = append(out, styleViolation{ID: "FOOT-004", Line: declarationLine + 1, Column: 1,
				Detail: "footnotes are not supported in GitHub wikis; use an ordinary numbered references section"})
		default:
			for _, token := range tokens {
				if !token.definition {
					add(token, "FOOT-004", fmt.Sprintf("footnote reference [^%s] is not supported in GitHub wikis; use an ordinary numbered references section", token.identifier))
					break
				}
			}
		}
	}
	return out
}

func validFootnoteIdentifier(identifier string) bool {
	if identifier == "" || identifier[0] < 'a' || identifier[0] > 'z' {
		return false
	}
	for i := 1; i < len(identifier); i++ {
		c := identifier[i]
		if c < 'a' || c > 'z' {
			if (c < '0' || c > '9') && c != '-' {
				return false
			}
		}
	}
	return true
}

func footnoteLineStarts(lines []string) []int {
	starts := make([]int, len(lines))
	start := 0
	for i, line := range lines {
		starts[i] = start
		start += len(line) + 1
	}
	return starts
}

func footnoteLocation(starts []int, offset int) (line, column int) {
	line = sort.Search(len(starts), func(i int) bool { return starts[i] > offset }) - 1
	if line < 0 {
		return 0, offset
	}
	return line, offset - starts[line]
}

func scanFootnoteTokens(raw, masked string, lines []string, starts []int) []footnoteToken {
	var tokens []footnoteToken
	for open := 0; open+1 < len(masked); {
		relative := strings.Index(masked[open:], "[^")
		if relative < 0 {
			break
		}
		open += relative
		if footnoteEscaped(masked, open) {
			open += 2
			continue
		}
		close := footnoteMatchingBracket(masked, open)
		if close < 0 {
			open += 2
			continue
		}
		line, column := footnoteLocation(starts, open)
		endLine, closeColumn := footnoteLocation(starts, close)
		definition := column <= 3 && strings.Trim(lines[line][:column], " ") == "" && close+1 < len(masked) && masked[close+1] == ':'
		tokens = append(tokens, footnoteToken{
			identifier: raw[open+2 : close], start: open,
			line: line, column: column, endLine: endLine, bodyColumn: closeColumn + 2, definition: definition,
		})
		open = close + 1
	}
	return tokens
}

func footnoteEscaped(text string, offset int) bool {
	backslashes := 0
	for offset > 0 && text[offset-1] == '\\' {
		backslashes++
		offset--
	}
	return backslashes%2 != 0
}

func footnoteMatchingBracket(text string, open int) int {
	depth := 0
	for i := open; i < len(text); i++ {
		if text[i] != '[' && text[i] != ']' || footnoteEscaped(text, i) {
			continue
		}
		if text[i] == '[' {
			depth++
			continue
		}
		depth--
		if depth == 0 {
			return i
		}
	}
	return -1
}

func footnoteMatchingParenthesis(text string, open int) int {
	depth := 0
	for i := open; i < len(text); i++ {
		if text[i] != '(' && text[i] != ')' || footnoteEscaped(text, i) {
			continue
		}
		if text[i] == '(' {
			depth++
			continue
		}
		depth--
		if depth == 0 {
			return i
		}
	}
	return -1
}

func footnoteCompoundImageLabel(identifier string) bool {
	if !strings.HasPrefix(identifier, "![") {
		return false
	}
	close := footnoteMatchingBracket(identifier, 1)
	if close < 0 || close+1 >= len(identifier) || identifier[close+1] != '(' {
		return false
	}
	return footnoteMatchingParenthesis(identifier, close+1) == len(identifier)-1
}

func footnoteLinkLabelRanges(masked string) [][2]int {
	var labels [][2]int
	for open := range masked {
		if masked[open] != '[' || footnoteEscaped(masked, open) {
			continue
		}
		close := footnoteMatchingBracket(masked, open)
		if close < 0 || close+1 >= len(masked) {
			continue
		}
		switch masked[close+1] {
		case '(':
			if footnoteMatchingParenthesis(masked, close+1) >= 0 {
				labels = append(labels, [2]int{open, close})
			}
		case '[':
			// A second footnote marker, as in text[^one][^two], is not a
			// reference-link destination and must not nest the first marker.
			if strings.HasPrefix(masked[close+1:], "[^") {
				continue
			}
			if footnoteMatchingBracket(masked, close+1) >= 0 {
				labels = append(labels, [2]int{open, close})
			}
		}
	}
	return labels
}

func footnoteInsideLabels(token footnoteToken, labels [][2]int) bool {
	for _, label := range labels {
		if token.start >= label[0] && token.start < label[1] {
			return true
		}
	}
	return false
}

func footnoteMathOffset(raw, masked string) int {
	best := -1
	if spans := mathPolicyDollarSpans(masked); len(spans) > 0 {
		best = spans[0].start
	}
	for i := range len(raw) - 1 {
		if raw[i] != '$' || masked[i] != '$' || footnoteEscaped(raw, i) {
			continue
		}
		matched := false
		switch raw[i+1] {
		case '`':
			run := 1
			for i+1+run < len(raw) && raw[i+1+run] == '`' {
				run++
			}
			for close := i + 1 + run; close < len(raw); {
				if raw[close] != '`' {
					close++
					continue
				}
				closingRun := 0
				for close+closingRun < len(raw) && raw[close+closingRun] == '`' {
					closingRun++
				}
				if closingRun == run && close+closingRun < len(raw) && raw[close+closingRun] == '$' && masked[close+closingRun] == '$' {
					matched = true
					break
				}
				close += closingRun
			}
		case '$':
			if masked[i+1] == '$' {
				for close := i + 2; close+1 < len(raw); close++ {
					if raw[close] == '$' && raw[close+1] == '$' && masked[close] == '$' && masked[close+1] == '$' && !footnoteEscaped(raw, close) {
						matched = true
						break
					}
				}
			}
		}
		if matched && (best < 0 || i < best) {
			best = i
		}
	}
	return best
}
