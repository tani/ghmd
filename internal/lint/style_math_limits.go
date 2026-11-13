package lint

import (
	"fmt"
	"strings"
)

// maxMathBlockBytes bounds a single display-math block. Inline expressions
// have a tighter limit in checkSpan. The upstream cases 23 and 24 from
// https://github.com/nschloe/markdown-math-acid-test use blocks well beyond
// this size to exercise resource behavior.
const maxMathBlockBytes = 512

func mathPolicyBlockLimitViolations(lines []string) []styleViolation {
	var out []styleViolation
	for i := 0; i < len(lines); i++ {
		marker, size, info, fenced := fenceAt(lines[i])
		if fenced {
			closeLine := -1
			for j := i + 1; j < len(lines); j++ {
				m, n, rest, ok := fenceAt(lines[j])
				if ok && m == marker && n >= size && rest == "" {
					closeLine = j
					break
				}
			}
			if closeLine < 0 {
				return out
			}
			fields := strings.Fields(info)
			if len(fields) > 0 && strings.EqualFold(fields[0], "math") {
				out = appendMathBlockLimit(out, lines[i+1:closeLine], i, lines[i])
			}
			i = closeLine
			continue
		}

		if !mathBlockDelimiter(lines[i]) {
			continue
		}
		closeLine := -1
		for j := i + 1; j < len(lines); j++ {
			if _, _, _, ok := fenceAt(lines[j]); ok {
				break
			}
			if mathBlockDelimiter(lines[j]) {
				closeLine = j
				break
			}
		}
		if closeLine < 0 {
			continue
		}
		out = appendMathBlockLimit(out, lines[i+1:closeLine], i, lines[i])
		i = closeLine
	}
	return out
}

func mathPolicyBlockContextViolations(lines []string, htmlStacks [][]mathPolicyHTMLTag) []styleViolation {
	var out []styleViolation
	for i := 0; i < len(lines); i++ {
		marker, size, info, fenced := fenceAt(lines[i])
		if fenced {
			closeLine := -1
			for j := i + 1; j < len(lines); j++ {
				m, n, rest, ok := fenceAt(lines[j])
				if ok && m == marker && n >= size && rest == "" {
					closeLine = j
					break
				}
			}
			if closeLine < 0 {
				return out
			}
			fields := strings.Fields(info)
			if len(fields) > 0 && strings.EqualFold(fields[0], "math") {
				out = appendMathBlockContext(out, i, lines[i], htmlStacks[i])
				out = appendMathFencedBodyViolations(out, lines[i+1:closeLine], i+1)
			}
			i = closeLine
			continue
		}

		if !mathBlockDelimiter(lines[i]) {
			continue
		}
		closeLine := -1
		for j := i + 1; j < len(lines); j++ {
			if _, _, _, ok := fenceAt(lines[j]); ok {
				break
			}
			if mathBlockDelimiter(lines[j]) {
				closeLine = j
				break
			}
		}
		if closeLine < 0 {
			continue
		}
		out = appendMathBlockContext(out, i, lines[i], htmlStacks[i])
		i = closeLine
	}
	return out
}

func appendMathFencedBodyViolations(out []styleViolation, body []string, firstLine int) []styleViolation {
	for offset, line := range body {
		for _, at := range mathPolicyNestedTextDollars(line) {
			out = append(out, styleViolation{ID: "MATH-006", Line: firstLine + offset + 1,
				Column: styleColumn(line, line[:at]),
				Detail: "avoid nested dollar math inside \\text{...}; use supported TeX constructs without changing meaning"})
		}
	}
	return out
}

func appendMathBlockContext(out []styleViolation, line int, opening string, stack []mathPolicyHTMLTag) []styleViolation {
	if len(stack) == 0 {
		return out
	}
	column := indentOf(opening) + 1
	out = append(out, styleViolation{ID: "MATH-005", Line: line + 1, Column: column,
		Detail: "move block math outside HTML containers, or verify the exact nested target structure"})
	if mathPolicyDetailsOpen(stack) {
		out = append(out, styleViolation{ID: "DETAILS-002", Line: line + 1, Column: column,
			Detail: "move equation widgets outside collapsible sections or test the actual initial expansion"})
	}
	return out
}

func mathBlockDelimiter(line string) bool {
	indent := indentOf(line)
	return indent <= 3 && !strings.Contains(line[:indent], "\t") && strings.TrimSpace(line) == "$$"
}

func appendMathBlockLimit(out []styleViolation, body []string, line int, opening string) []styleViolation {
	if mathBlockBytes(body) <= maxMathBlockBytes {
		return out
	}
	return append(out, styleViolation{ID: "MATH-009", Line: line + 1,
		Column: indentOf(opening) + 1,
		Detail: fmt.Sprintf("display math exceeds %d bytes; split it at valid TeX boundaries or attach readable source", maxMathBlockBytes)})
}

func mathBlockBytes(lines []string) int {
	total := 0
	for i, line := range lines {
		if i > 0 {
			total++
		}
		total += len(line)
	}
	return total
}
