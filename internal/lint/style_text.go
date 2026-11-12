package lint

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type coreEmphasisRun struct {
	start, end        int
	marker            byte
	canOpen, canClose bool
}

func textStyleViolations(lines, masked []string, insideFence []bool) []styleViolation {
	var violations []styleViolation
	for i, line := range lines {
		if insideFence[i] {
			continue
		}
		analysis := masked[i]
		if coreThematicDelimiterLine(analysis) {
			continue
		}
		runs := coreEmphasisDelimiters(analysis)
		if at, delimiter, opening, ok := coreEmphasisBoundarySpace(analysis, runs); ok {
			side := "opening"
			if !opening {
				side = "closing"
			}
			violations = append(violations, styleViolation{
				ID: "TEXT-001", Line: i + 1, Column: coreColumn(line, at),
				Detail: "remove whitespace inside the " + side + " " + delimiter + " delimiter in " + strings.TrimSpace(line),
			})
		} else if at, delimiter, nested, ok := coreOverlappingEmphasis(analysis, runs); ok {
			kind := "overlapping"
			if nested {
				kind = "nested"
			}
			violations = append(violations, styleViolation{
				ID: "TEXT-001", Line: i + 1, Column: coreColumn(line, at),
				Detail: "rewrite " + kind + " " + delimiter + " emphasis as separate, non-overlapping delimiters",
			})
		}
		if at, punctuation, ok := coreCJKClosingPunctuation(analysis, runs); ok {
			violations = append(violations, styleViolation{
				ID: "TEXT-002", Line: i + 1, Column: coreColumn(line, at),
				Detail: "move CJK punctuation " + punctuation + " outside the closing emphasis delimiter",
			})
		}
	}
	return violations
}

func coreEmphasisDelimiters(line string) []coreEmphasisRun {
	var runs []coreEmphasisRun
	for at := 0; at < len(line); {
		if line[at] != '*' && line[at] != '_' || coreEscaped(line, at) {
			at++
			continue
		}
		marker := line[at]
		end := at + 1
		for end < len(line) && line[end] == marker {
			end++
		}
		prev, prevOK := coreRuneBefore(line, at)
		next, nextSize := utf8.DecodeRuneInString(line[end:])
		nextOK := nextSize > 0
		if !nextOK {
			next = 0
		}
		prevSpace := !prevOK || unicode.IsSpace(prev)
		nextSpace := !nextOK || unicode.IsSpace(next)
		prevPunct := prevOK && (unicode.IsPunct(prev) || unicode.IsSymbol(prev))
		nextPunct := nextOK && (unicode.IsPunct(next) || unicode.IsSymbol(next))
		canOpen := !nextSpace && (!nextPunct || prevSpace || prevPunct)
		canClose := !prevSpace && (!prevPunct || nextSpace || nextPunct)
		if marker == '_' && prevOK && nextOK && coreAlphaNumeric(prev) && coreAlphaNumeric(next) {
			canOpen = false
			canClose = false
		}
		runs = append(runs, coreEmphasisRun{start: at, end: end, marker: marker, canOpen: canOpen, canClose: canClose})
		at = end
	}
	return runs
}

func coreEmphasisBoundarySpace(line string, runs []coreEmphasisRun) (at int, delimiter string, opening, ok bool) {
	for _, run := range runs {
		length := run.end - run.start
		if length > 2 || length == 1 && run.marker == '_' && run.start < len(line) && run.start > 0 && coreEscaped(line, run.start) {
			continue
		}
		prev, prevOK := coreRuneBefore(line, run.start)
		next, nextSize := utf8.DecodeRuneInString(line[run.end:])
		nextOK := nextSize > 0
		prevSpace := !prevOK || unicode.IsSpace(prev)
		nextSpace := !nextOK || unicode.IsSpace(next)
		prevPunct := prevOK && (unicode.IsPunct(prev) || unicode.IsSymbol(prev))
		nextPunct := nextOK && (unicode.IsPunct(next) || unicode.IsSymbol(next))
		isBullet := length == 1 && run.marker == '*' && strings.TrimSpace(line[:run.start]) == "" && nextSpace
		isolatedAsterisk := length == 1 && prevOK && nextOK && prevSpace && nextSpace
		if isBullet || isolatedAsterisk {
			continue
		}
		if nextOK && nextSpace && (prevSpace || prevPunct) {
			return run.start, line[run.start:run.end], true, true
		}
		if prevOK && prevSpace && (nextSpace || nextPunct) {
			return run.start, line[run.start:run.end], false, true
		}
	}
	return 0, "", false, false
}

func coreOverlappingEmphasis(line string, runs []coreEmphasisRun) (at int, delimiter string, nested, ok bool) {
	type opener struct {
		index  int
		marker byte
		length int
	}
	var stack []opener
	for i, run := range runs {
		length := run.end - run.start
		if length > 2 {
			if run.canOpen || run.canClose {
				return run.start, strings.Repeat(string(run.marker), length), false, true
			}
			continue
		}
		closed := false
		if run.canClose && len(stack) > 0 {
			top := stack[len(stack)-1]
			if top.marker == run.marker && top.length == length {
				stack = stack[:len(stack)-1]
				closed = true
				if len(stack) > 0 {
					return runs[top.index].start, line[runs[top.index].start:runs[top.index].end], true, true
				}
			} else {
				compatible := -1
				for j := len(stack) - 2; j >= 0; j-- {
					if stack[j].marker == run.marker && stack[j].length == length {
						compatible = j
						break
					}
				}
				if compatible >= 0 {
					return run.start, line[run.start:run.end], false, true
				}
			}
		}
		if run.canOpen && !closed {
			stack = append(stack, opener{index: i, marker: run.marker, length: length})
		}
	}
	_ = line
	return 0, "", false, false
}

func coreCJKClosingPunctuation(line string, runs []coreEmphasisRun) (int, string, bool) {
	for i, run := range runs {
		if run.marker != '*' || run.end-run.start > 2 {
			continue
		}
		before := strings.TrimRight(line[:run.start], " \t")
		punctuation, ok := coreRuneBefore(before, len(before))
		if !ok || !strings.ContainsRune("。．，、！？；：｡､", punctuation) {
			continue
		}
		punctuationOffset := len(before) - utf8.RuneLen(punctuation)
		contentRune, contentOK := coreRuneBefore(before, punctuationOffset)
		if !contentOK || !coreCJK(contentRune) {
			continue
		}
		hadEarlierDelimiter := false
		hasLaterDelimiter := false
		for j := range runs {
			if runs[j].marker != run.marker || runs[j].end-runs[j].start != run.end-run.start {
				continue
			}
			hadEarlierDelimiter = hadEarlierDelimiter || j < i
			hasLaterDelimiter = hasLaterDelimiter || j > i
		}
		if hadEarlierDelimiter && !hasLaterDelimiter {
			return run.start, string(punctuation), true
		}
	}
	return 0, "", false
}

func coreThematicDelimiterLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return false
	}
	marker := trimmed[0]
	if marker != '*' && marker != '_' {
		return false
	}
	count := 0
	for _, r := range trimmed {
		switch r {
		case ' ', '\t':
		case rune(marker):
			count++
		default:
			return false
		}
	}
	return count >= 3
}

func coreRuneBefore(line string, offset int) (rune, bool) {
	if offset <= 0 {
		return 0, false
	}
	r, _ := utf8.DecodeLastRuneInString(line[:offset])
	return r, true
}

func coreAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func coreCJK(r rune) bool {
	return unicode.In(r, unicode.Han, unicode.Hiragana, unicode.Katakana, unicode.Hangul)
}

func coreEscaped(line string, at int) bool {
	backslashes := 0
	for at > 0 && line[at-1] == '\\' {
		backslashes++
		at--
	}
	return backslashes%2 == 1
}
