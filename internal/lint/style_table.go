package lint

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// tableStrikeStyleViolations implements the table and strikethrough authoring
// rules. Table rows are scanned lexically: escaped pipes are content, while
// code and math spans retain their offsets so a pipe can be diagnosed without
// being mistaken for a cell boundary.
func tableStrikeStyleViolations(source string) []styleViolation {
	lines := sourceLines(source)
	inside, fenceOpens := fencedCodeRegions(tableFenceScanLines(lines))
	violations := tableStyleLines(lines, inside, fenceOpens)
	return append(violations, strikeStyleSource(source, lines, inside, fenceOpens)...)
}

// At the start of a line, a long strikeout run can resemble a tilde fence to a
// block scanner. Mask that opening run only for fence detection when a matching
// closing run occurs on the same line; the original source remains unchanged.
func tableFenceScanLines(lines []string) []string {
	var adjusted []string
	for i, line := range lines {
		if !tableLongStrikeLine(line) {
			continue
		}
		if adjusted == nil {
			adjusted = append([]string(nil), lines...)
		}
		start := len(line) - len(strings.TrimLeft(line, " "))
		size := tableByteRun(line, start, '~')
		adjusted[i] = line[:start] + strings.Repeat(" ", size) + line[start+size:]
	}
	if adjusted != nil {
		return adjusted
	}
	return lines
}

func tableLongStrikeLine(line string) bool {
	start := len(line) - len(strings.TrimLeft(line, " "))
	if start > 3 || start >= len(line) || line[start] != '~' {
		return false
	}
	size := tableByteRun(line, start, '~')
	if size < 3 || start+size >= len(line) {
		return false
	}
	next, _ := utf8.DecodeRuneInString(line[start+size:])
	if unicode.IsSpace(next) {
		return false
	}
	for i := start + size; i < len(line); {
		if line[i] != '~' {
			i++
			continue
		}
		run := tableByteRun(line, i, '~')
		if run == size {
			previous, _ := utf8.DecodeLastRuneInString(line[:i])
			return !unicode.IsSpace(previous)
		}
		i += run
	}
	return false
}

type tableSpan struct {
	start, end int
	display    bool
}

type tableCell struct {
	start, end int
}

type tablePipeRow struct {
	line       int
	text       string
	cells      []tableCell
	boundaries []int
	code       []tableSpan
	math       []tableSpan
	leading    bool
	trailing   bool
}

func tableStyleLines(lines []string, inside, fenceOpens []bool) []styleViolation {
	var violations []styleViolation
	for i := 0; i+1 < len(lines); {
		if inside[i] || fenceOpens[i] || inside[i+1] || fenceOpens[i+1] {
			i++
			continue
		}
		header := parseTablePipeRow(lines[i], i)
		delimiter := parseTablePipeRow(lines[i+1], i+1)
		delimiterOK, delimiterValid := tableDelimiterRow(delimiter)
		if !delimiterOK || len(delimiter.boundaries) == 0 || strings.TrimSpace(header.text) == "" {
			i++
			continue
		}

		rows := []tablePipeRow{header, delimiter}
		j := i + 2
		for ; j < len(lines); j++ {
			if inside[j] || fenceOpens[j] || strings.TrimSpace(lines[j]) == "" {
				break
			}
			row := parseTablePipeRow(lines[j], j)
			// A pipe in code, math, or an HTML comment is not enough to establish
			// a table body row. This is what separates subsequent prose from a
			// malformed row with missing or extra cells.
			if len(row.boundaries) == 0 {
				break
			}
			rows = append(rows, row)
		}

		validTable := delimiterValid && len(header.cells) == len(delimiter.cells)
		for rowIndex, row := range rows {
			label := "body row"
			if rowIndex == 0 {
				label = "header"
			} else if rowIndex == 1 {
				label = "delimiter row"
			}
			violations = append(violations, tableOuterPipeViolations(row, label)...)

			if rowIndex == 1 && len(row.cells) != len(header.cells) {
				violations = append(violations, tableViolation("TABLE-001", row, 0,
					fmt.Sprintf("delimiter has %d cells but the header has %d; make every table row have the header's cell count", len(row.cells), len(header.cells))))
			}
			if rowIndex >= 2 && len(row.cells) != len(header.cells) {
				violations = append(violations, tableViolation("TABLE-001", row, 0,
					fmt.Sprintf("body row has %d cells but the header has %d; remove extras or supply the missing cells", len(row.cells), len(header.cells))))
			}

			if rowIndex == 1 {
				for _, cell := range row.cells {
					content := tableDelimiterContent(row.text[cell.start:cell.end])
					if content != "" && tableHyphenCount(content) < 3 {
						violations = append(violations, tableViolation("TABLE-001", row, cell.start,
							fmt.Sprintf("delimiter cell %q needs at least three hyphens, for example ---", content)))
					}
				}
			} else {
				for _, cell := range row.cells {
					if reason := tableComplexCell(row, cell); reason != "" {
						violations = append(violations, tableViolation("TABLE-002", row, cell.start,
							fmt.Sprintf("table cell contains %s; move complex content to a section instead of a pipe table", reason)))
					}
				}
			}

			violations = append(violations, tableLiteralPipeViolations(row)...)
			if validTable {
				violations = append(violations, tableMathPipeViolations(row)...)
			}
		}
		i = j
	}
	return violations
}

func tableViolation(id string, row tablePipeRow, offset int, detail string) styleViolation {
	if offset < 0 || offset > len(row.text) {
		offset = 0
	}
	return styleViolation{
		ID:     id,
		Line:   row.line + 1,
		Column: styleColumn(row.text, row.text[:offset]),
		Detail: detail,
	}
}

func tableOuterPipeViolations(row tablePipeRow, label string) []styleViolation {
	if row.leading && row.trailing {
		return nil
	}
	missing := make([]string, 0, 2)
	if !row.leading {
		missing = append(missing, "leading")
	}
	if !row.trailing {
		missing = append(missing, "trailing")
	}
	return []styleViolation{tableViolation("TABLE-001", row, 0,
		fmt.Sprintf("table %s needs %s pipe%s; use | ... | on every row", label, strings.Join(missing, " and "), tablePlural(len(missing))))}
}

func tablePlural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func tableDelimiterContent(content string) string {
	return strings.Trim(content, " \t\r")
}

func tableHyphenCount(content string) int {
	return strings.Count(content, "-")
}

// tableDelimiterRow recognizes intended delimiter rows, including rows with
// too-short hyphen runs so TABLE-001 can repair them. Cells containing anything
// other than colons and hyphens remain ordinary prose.
func tableDelimiterRow(row tablePipeRow) (recognized, valid bool) {
	if len(row.cells) == 0 {
		return false, false
	}
	valid = true
	for _, cell := range row.cells {
		content := tableDelimiterContent(row.text[cell.start:cell.end])
		if content == "" {
			return false, false
		}
		if content[0] == ':' {
			content = content[1:]
		}
		if len(content) > 0 && content[len(content)-1] == ':' {
			content = content[:len(content)-1]
		}
		if content == "" || strings.Trim(content, "-") != "" {
			return false, false
		}
		if len(content) < 3 {
			valid = false
		}
	}
	return true, valid
}

func parseTablePipeRow(line string, lineNumber int) tablePipeRow {
	row := tablePipeRow{line: lineNumber, text: line}
	row.code = tableInlineCodeSpans(line)
	row.math = tableInlineMathSpans(line, row.code)
	masked := maskLineCode(line)
	inMath := func(offset int) bool {
		return tableOffsetInSpans(row.math, offset)
	}
	for i := range len(line) {
		if line[i] != '|' || masked[i] != '|' || inMath(i) {
			continue
		}
		// GitHub protects a table pipe whenever a backslash immediately
		// precedes it, including repeated-backslash paths and regular
		// expressions. TABLE-003 diagnoses those ambiguous runs separately.
		if tableBackslashRun(line, i) == 0 {
			row.boundaries = append(row.boundaries, i)
		}
	}
	if len(row.boundaries) == 0 {
		row.cells = []tableCell{{start: 0, end: tableRowEnd(line)}}
		return row
	}

	start := 0
	firstCell := 0
	lastBoundary := len(row.boundaries)
	trimmedStart := 0
	for trimmedStart < len(line) && (line[trimmedStart] == ' ' || line[trimmedStart] == '\t') {
		trimmedStart++
	}
	if row.boundaries[0] == trimmedStart {
		row.leading = true
		start = row.boundaries[0] + 1
		firstCell = 1
	}
	end := tableRowEnd(line)
	if lastBoundary > firstCell && row.boundaries[lastBoundary-1] == end-1 {
		row.trailing = true
		end = row.boundaries[lastBoundary-1]
		lastBoundary--
	}

	for _, boundary := range row.boundaries[firstCell:lastBoundary] {
		row.cells = append(row.cells, tableCell{start: start, end: boundary})
		start = boundary + 1
	}
	row.cells = append(row.cells, tableCell{start: start, end: end})
	return row
}

func tableRowEnd(line string) int {
	end := len(line)
	for end > 0 && (line[end-1] == ' ' || line[end-1] == '\t' || line[end-1] == '\r') {
		end--
	}
	return end
}

func tableOffsetInSpans(spans []tableSpan, offset int) bool {
	index := sort.Search(len(spans), func(i int) bool { return spans[i].end > offset })
	return index < len(spans) && spans[index].start <= offset
}

func tableSpanAt(spans []tableSpan, offset int) (tableSpan, bool) {
	index := sort.Search(len(spans), func(i int) bool { return spans[i].end > offset })
	if index < len(spans) && spans[index].start <= offset {
		return spans[index], true
	}
	return tableSpan{}, false
}

// tableInlineCodeSpans returns complete CommonMark code spans, including their
// backtick delimiters. Backslash escaping applies to opening backticks, but
// backslashes inside a span are ordinary code content.
func tableInlineCodeSpans(line string) []tableSpan {
	var spans []tableSpan
	for i := 0; i < len(line); {
		if strings.HasPrefix(line[i:], "<!--") {
			if end := strings.Index(line[i+4:], "-->"); end >= 0 {
				i += 4 + end + 3
			} else {
				return spans
			}
			continue
		}
		if line[i] != '`' || tableBackslashRun(line, i)%2 != 0 {
			i++
			continue
		}
		size := tableByteRun(line, i, '`')
		close := tableClosingBackticks(line, i+size, size)
		if close < 0 {
			i += size
			continue
		}
		spans = append(spans, tableSpan{start: i, end: close + size})
		i = close + size
	}
	return spans
}

func tableClosingBackticks(text string, from, size int) int {
	for i := from; i < len(text); {
		if text[i] != '`' {
			i++
			continue
		}
		run := tableByteRun(text, i, '`')
		if run == size {
			return i
		}
		i += run
	}
	return -1
}

func tableInlineMathSpans(line string, code []tableSpan) []tableSpan {
	var spans []tableSpan
	masked := maskLineCode(line)
	for i := 0; i < len(line); {
		if line[i] != '$' || masked[i] != '$' || tableBackslashRun(line, i)%2 != 0 || tableOffsetInSpans(code, i) {
			i++
			continue
		}
		size := tableByteRun(line, i, '$')
		if size > 2 {
			i += size
			continue
		}
		if i+size >= len(line) {
			break
		}
		first, _ := utf8.DecodeRuneInString(line[i+size:])
		if size == 1 && (unicode.IsSpace(first) || unicode.IsDigit(first)) {
			i += size
			continue
		}

		close := -1
		for j := i + size; j < len(line); {
			if masked[j] != '$' || line[j] != '$' || tableBackslashRun(line, j)%2 != 0 {
				j++
				continue
			}
			closeSize := tableByteRun(line, j, '$')
			if closeSize == size {
				close = j
				break
			}
			j += closeSize
		}
		if close < 0 {
			i += size
			continue
		}
		last, _ := utf8.DecodeLastRuneInString(line[i+size : close])
		if last == utf8.RuneError && close == i+size || size == 1 && unicode.IsSpace(last) {
			i += size
			continue
		}
		spans = append(spans, tableSpan{start: i, end: close + size, display: size == 2})
		i = close + size
	}
	return spans
}

func tableByteRun(text string, at int, value byte) int {
	end := at
	for end < len(text) && text[end] == value {
		end++
	}
	return end - at
}

func tableBackslashRun(text string, at int) int {
	start := at
	for start > 0 && text[start-1] == '\\' {
		start--
	}
	return at - start
}

func tableComplexCell(row tablePipeRow, cell tableCell) string {
	content := strings.TrimSpace(row.text[cell.start:cell.end])
	if content == "" {
		return ""
	}
	if tableListItem(content) {
		return "an inline list item"
	}
	if strings.HasPrefix(content, "~~~") || tableUnclosedCodeFence(row, cell) {
		return "a fenced code structure"
	}
	lower := strings.ToLower(content)
	for _, tag := range []string{"<ul", "<ol", "<li", "<pre", "<table", "<blockquote", "<p>", "<div"} {
		if strings.HasPrefix(lower, tag) {
			return "a complex HTML block"
		}
	}
	if strings.Contains(lower, "<") && (strings.Contains(lower, "colspan") || strings.Contains(lower, "rowspan")) {
		return "merged cells"
	}
	for _, span := range row.math {
		if span.display && span.start >= cell.start && span.end <= cell.end {
			return "display mathematics"
		}
	}
	return ""
}

func tableUnclosedCodeFence(row tablePipeRow, cell tableCell) bool {
	for i := cell.start; i < cell.end; {
		if row.text[i] != '`' {
			i++
			continue
		}
		size := tableByteRun(row.text, i, '`')
		if size >= 3 && tableBackslashRun(row.text, i)%2 == 0 && !tableOffsetInSpans(row.code, i) {
			return true
		}
		i += size
	}
	return false
}

func tableListItem(content string) bool {
	if len(content) >= 2 && (content[0] == '-' || content[0] == '+' || content[0] == '*') &&
		(content[1] == ' ' || content[1] == '\t') {
		return true
	}
	i := 0
	for i < len(content) && content[i] >= '0' && content[i] <= '9' {
		i++
	}
	return i > 0 && i+1 < len(content) && (content[i] == '.' || content[i] == ')') &&
		(content[i+1] == ' ' || content[i+1] == '\t')
}

func tableLiteralPipeViolations(row tablePipeRow) []styleViolation {
	var violations []styleViolation
	for i := range len(row.text) {
		if row.text[i] != '|' {
			continue
		}
		slashes := tableBackslashRun(row.text, i)
		_, inCode := tableSpanAt(row.code, i)
		if inCode && slashes != 1 {
			detail := "inline code contains an unescaped literal pipe; write it as \\| or move the code out of the table"
			if slashes >= 2 {
				detail = fmt.Sprintf("%d backslashes precede this pipe inside inline code; move the exact code outside the table instead of collapsing the backslash run", slashes)
			}
			violations = append(violations, tableViolation("TABLE-003", row, i, detail))
			continue
		}
		if slashes >= 2 {
			// Repeated backslashes have renderer-specific escaping behavior. Quote
			// the count rather than silently reducing them to one escape.
			violations = append(violations, tableViolation("TABLE-003", row, i,
				fmt.Sprintf("%d consecutive backslashes precede this pipe; move the exact content outside the table instead of silently discarding backslashes", slashes)))
		}
	}
	return violations
}

func tableMathPipeViolations(row tablePipeRow) []styleViolation {
	var violations []styleViolation
	for _, span := range row.math {
		if span.display {
			continue
		}
		for i := span.start + 1; i < span.end-1; i++ {
			if row.text[i] != '|' || tableBackslashRun(row.text, i)%2 != 0 {
				continue
			}
			violations = append(violations, tableViolation("TABLE-004", row, i,
				"inline math contains a literal pipe; use \\vert or \\mid where semantically appropriate, or move the math out of the table"))
		}
	}
	return violations
}

type tableTildeRun struct {
	start, end int
}

type tableTildePair struct {
	open, close tableTildeRun
}

func strikeStyleSource(source string, lines []string, inside, fenceOpens []bool) []styleViolation {
	if !strings.Contains(source, "~") {
		return nil
	}
	masked, starts := tableStrikeMaskedSource(source, lines, inside, fenceOpens)
	var runs []tableTildeRun
	for i := 0; i < len(source); {
		if masked[i] != '~' || tableBackslashRun(source, i)%2 != 0 {
			i++
			continue
		}
		end := i + tableByteRun(source, i, '~')
		// Masked bytes have spaces substituted by the code/comment scanner. Only
		// a wholly visible run can be a delimiter.
		visible := true
		for j := i; j < end; j++ {
			if masked[j] != '~' {
				visible = false
				break
			}
		}
		if visible {
			runs = append(runs, tableTildeRun{start: i, end: end})
		}
		i = end
	}

	pending := make(map[int][]int)
	var pairs []tableTildePair
	previousEnd := 0
	for i, run := range runs {
		if strikeBlankLine(source[previousEnd:run.start]) {
			clear(pending)
		}
		previousEnd = run.end
		size := run.end - run.start
		canClose := strikeCanClose(source, run)
		if canClose && len(pending[size]) > 0 {
			openIndex := pending[size][len(pending[size])-1]
			pending[size] = pending[size][:len(pending[size])-1]
			pairs = append(pairs, tableTildePair{open: runs[openIndex], close: run})
			continue
		}
		if strikeCanOpen(source, run) {
			pending[size] = append(pending[size], i)
		}
	}

	var violations []styleViolation
	for _, pair := range pairs {
		size := pair.open.end - pair.open.start
		lineIndex := sort.Search(len(starts), func(i int) bool { return starts[i] > pair.open.start }) - 1
		if lineIndex < 0 {
			lineIndex = 0
		}
		prefix := source[starts[lineIndex]:pair.open.start]
		line := ""
		if lineIndex < len(lines) {
			line = lines[lineIndex]
		}
		if size != 2 {
			violations = append(violations, styleViolation{
				ID:     "STRIKE-001",
				Line:   lineIndex + 1,
				Column: styleColumn(line, prefix),
				Detail: fmt.Sprintf("strikethrough uses %d tildes on each side; use exactly ~~...~~, not single or longer tilde runs", size),
			})
			continue
		}

		var reasons []string
		if strings.Contains(source[pair.open.end:pair.close.start], "\n") {
			reasons = append(reasons, "it crosses source lines")
		}
		if strikeAdjacentLetter(source, pair.open.start, pair.close.end) {
			reasons = append(reasons, "it touches adjacent letters without prose punctuation or spacing")
		}
		if strikeEmphasisOverlap(source, masked, pair) {
			reasons = append(reasons, "it overlaps emphasis delimiters")
		}
		if len(reasons) != 0 {
			violations = append(violations, styleViolation{
				ID:     "STRIKE-002",
				Line:   lineIndex + 1,
				Column: styleColumn(line, prefix),
				Detail: fmt.Sprintf("strikethrough is hard to read because %s; keep one-line strikeout separate from adjacent words and emphasis, or use ordinary explanatory prose", strings.Join(reasons, " and ")),
			})
		}
	}
	return violations
}

func tableStrikeMaskedSource(source string, lines []string, inside, fenceOpens []bool) ([]byte, []int) {
	masked := []byte(source)
	starts := lineStartOffsets(source)
	fenced := make([]bool, len(source))
	for i, line := range lines {
		start := starts[i]
		end := start + len(line)
		if end > len(source) {
			end = len(source)
		}
		if inside[i] || fenceOpens[i] {
			for j := start; j < end; j++ {
				fenced[j] = true
				masked[j] = ' '
			}
			continue
		}
		lineMask := maskLineCode(line)
		copy(masked[start:end], lineMask)
	}

	// Complete the per-line helper's work for comments and code spans whose
	// closing delimiter occurs on another source line.
	for i := 0; i < len(source); {
		if fenced[i] {
			i++
			continue
		}
		if strings.HasPrefix(source[i:], "<!--") {
			end := strings.Index(source[i+4:], "-->")
			to := len(source)
			if end >= 0 {
				to = i + 4 + end + 3
			}
			tableMaskBytes(masked, i, to)
			i = to
			continue
		}
		if source[i] == '`' && tableBackslashRun(source, i)%2 == 0 {
			size := tableByteRun(source, i, '`')
			close := tableClosingBackticks(source, i+size, size)
			if close >= 0 && !tableAnyOffset(fenced, close, close+size) {
				tableMaskBytes(masked, i, close+size)
				i = close + size
				continue
			}
			i += size
			continue
		}
		i++
	}

	// TeX delimiters inside strikeout remain visible, but tildes and emphasis
	// markers in the expression itself are mathematics, not prose formatting.
	for i, line := range lines {
		if inside[i] || fenceOpens[i] {
			continue
		}
		code := tableInlineCodeSpans(line)
		for _, span := range tableInlineMathSpans(line, code) {
			tableMaskBytes(masked, starts[i]+span.start, starts[i]+span.end)
		}
	}
	return masked, starts
}

func tableAnyOffset(values []bool, from, to int) bool {
	for i := from; i < to && i < len(values); i++ {
		if values[i] {
			return true
		}
	}
	return false
}

func tableMaskBytes(masked []byte, from, to int) {
	if from < 0 {
		from = 0
	}
	if to > len(masked) {
		to = len(masked)
	}
	for i := from; i < to; i++ {
		if masked[i] != '\n' && masked[i] != '\r' {
			masked[i] = ' '
		}
	}
}

func strikeCanOpen(source string, run tableTildeRun) bool {
	if run.end >= len(source) {
		return false
	}
	next, _ := utf8.DecodeRuneInString(source[run.end:])
	return !unicode.IsSpace(next)
}

func strikeCanClose(source string, run tableTildeRun) bool {
	if run.start == 0 {
		return false
	}
	previous, _ := utf8.DecodeLastRuneInString(source[:run.start])
	return !unicode.IsSpace(previous)
}

func strikeBlankLine(text string) bool {
	lineStart := -1
	for i := range len(text) + 1 {
		if i < len(text) && text[i] != '\n' {
			continue
		}
		if lineStart >= 0 && strings.Trim(text[lineStart:i], " \t\r") == "" {
			return true
		}
		if i == len(text) {
			break
		}
		lineStart = i + 1
	}
	return false
}

func strikeAdjacentLetter(source string, open, close int) bool {
	if open > 0 {
		before, _ := utf8.DecodeLastRuneInString(source[:open])
		if unicode.IsLetter(before) || unicode.IsDigit(before) {
			return true
		}
	}
	if close < len(source) {
		after, _ := utf8.DecodeRuneInString(source[close:])
		return unicode.IsLetter(after) || unicode.IsDigit(after)
	}
	return false
}

func strikeEmphasisOverlap(source string, masked []byte, pair tableTildePair) bool {
	// Properly nested emphasis is valid: ~~*text*~~ and **~~text~~** are not
	// overlaps. Diagnose only crossed delimiter runs, such as **~~text**~~.
	before := strikeRunEndingAt(source, masked, pair.open.start)
	suffix := strikeRunEndingAt(source, masked, pair.close.start)
	if before.end > before.start && suffix.end > suffix.start && source[before.start] == source[suffix.start] {
		return true
	}
	prefix := strikeRunStartingAt(source, masked, pair.open.end)
	after := strikeRunStartingAt(source, masked, pair.close.end)
	return after.end > after.start && prefix.end > prefix.start && source[after.start] == source[prefix.start]
}

func strikeRunEndingAt(source string, masked []byte, at int) tableTildeRun {
	if at <= 0 || masked[at-1] != '*' && masked[at-1] != '_' {
		return tableTildeRun{}
	}
	marker := masked[at-1]
	start := at - 1
	for start > 0 && masked[start-1] == marker {
		start--
	}
	if tableBackslashRun(source, start)%2 != 0 {
		return tableTildeRun{}
	}
	return tableTildeRun{start: start, end: at}
}

func strikeRunStartingAt(source string, masked []byte, at int) tableTildeRun {
	if at >= len(source) || masked[at] != '*' && masked[at] != '_' {
		return tableTildeRun{}
	}
	marker := masked[at]
	end := at + tableByteRun(source, at, marker)
	for i := at; i < end; i++ {
		if masked[i] != marker {
			return tableTildeRun{}
		}
	}
	if tableBackslashRun(source, at)%2 != 0 {
		return tableTildeRun{}
	}
	return tableTildeRun{start: at, end: end}
}
