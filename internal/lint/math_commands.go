package lint

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

// GitHub's current math renderer rejects operatorname with "The following
// macros are not allowed". TeX macro definitions are rendered literally rather
// than executed. Both behaviors were observed on rendered GitHub math.
const mathDefinitionGuidance = `GitHub renders TeX macro definitions literally rather than executing them; inline the expanded expression or use a built-in command`

var disallowedMathMacros = []struct {
	name     string
	guidance string
}{
	{`\operatorname`, `GitHub's math renderer rejects \operatorname; use \mathrm{...} or \mathop{\mathrm{...}} and verify limit placement`},
	{`\def`, mathDefinitionGuidance},
	{`\gdef`, mathDefinitionGuidance},
	{`\edef`, mathDefinitionGuidance},
	{`\xdef`, mathDefinitionGuidance},
	{`\let`, mathDefinitionGuidance},
	{`\newcommand`, mathDefinitionGuidance},
	{`\renewcommand`, mathDefinitionGuidance},
	{`\providecommand`, mathDefinitionGuidance},
	{`\newenvironment`, mathDefinitionGuidance},
	{`\renewenvironment`, mathDefinitionGuidance},
	{`\DeclareMathOperator`, `GitHub rejects \DeclareMathOperator with "The following macros are not allowed"; use \mathrm or \mathop{\mathrm{...}}`},
}

type mathRegionForm uint8

const (
	mathFormInlineBare mathRegionForm = iota
	mathFormInlineCode
	mathFormDisplay
	mathFormFence
)

type mathRegion struct {
	form             mathRegionForm
	start, bodyStart int
	bodyEnd          int
}

func mathRegionFindings(source string) []Finding {
	starts := lineStartOffsets(source)
	var out []Finding
	for _, region := range mathRegions(source) {
		body := source[region.bodyStart:region.bodyEnd]
		for _, use := range mathUnknownCommands(body) {
			out = append(out, mathFindingAt(source, starts, region.bodyStart+use.at,
				ruleMathUnknownMacro, fmt.Sprintf(
					"\\%s is not in the combined KaTeX/MathJax command allowlist; verify support in the target renderer",
					use.name)))
		}
		for _, macro := range disallowedMathMacros {
			for from := 0; from < len(body); {
				at := strings.Index(body[from:], macro.name)
				if at < 0 {
					break
				}
				at += from
				from = at + len(macro.name)
				absolute := region.bodyStart + at
				if mathPolicyBackslashRun(source, absolute+1)%2 == 0 {
					continue
				}
				if end := absolute + len(macro.name); end < len(source) && asciiLetter(source[end]) {
					continue
				}
				out = append(out, mathFindingAt(source, starts, absolute,
					ruleMathDisallowedMacro, macro.guidance))
			}
		}

		if region.form == mathFormInlineCode && region.start > 0 &&
			unicode.IsLetter(runeBefore(source, region.start)) {
			out = append(out, mathFindingAt(source, starts, region.start, ruleMathNotRecognized,
				"dollar-backtick inline math immediately after a letter is not recognized; add whitespace before the opening dollar"))
		}

		if region.form == mathFormDisplay || region.form == mathFormFence {
			if at := mathHTMLLookalike(body); at >= 0 {
				out = append(out, mathFindingAt(source, starts, region.bodyStart+at, ruleMathHTMLLookalike,
					"GitHub may treat this <... sequence inside display math as HTML and drop math content; use \\lt or escape the delimiter and verify"))
			}
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Line != out[j].Line {
			return out[i].Line < out[j].Line
		}
		if out[i].Column != out[j].Column {
			return out[i].Column < out[j].Column
		}
		return out[i].Rule < out[j].Rule
	})
	return out
}

func mathRegions(source string) []mathRegion {
	lines := sourceLines(source)
	starts := lineStartOffsets(source)
	ignored := mathIgnoredRegions(source, starts)
	var out []mathRegion
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
				break
			}
			fields := strings.Fields(info)
			if len(fields) > 0 && strings.EqualFold(fields[0], "math") {
				out = append(out, mathRegion{form: mathFormFence, start: starts[i],
					bodyStart: starts[i+1], bodyEnd: starts[closeLine]})
			}
			i = closeLine
			continue
		}

		if mathBlockDelimiter(lines[i]) {
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
			if closeLine >= 0 {
				out = append(out, mathRegion{form: mathFormDisplay, start: starts[i],
					bodyStart: starts[i+1], bodyEnd: starts[closeLine]})
				i = closeLine
				continue
			}
		}

		base := starts[i]
		line := mathPolicyMaskComment(lines[i])
		for _, span := range mathPolicyDollarBacktickSpans(line) {
			start := base + span.start
			if start > 0 && ignored[start-1] {
				continue
			}
			out = append(out, mathRegion{form: mathFormInlineCode, start: start,
				bodyStart: base + span.bodyStart, bodyEnd: base + span.bodyEnd})
		}
		for _, span := range mathPolicyDollarSpans(maskLineCode(line)) {
			if span.display {
				continue
			}
			out = append(out, mathRegion{form: mathFormInlineBare, start: base + span.start,
				bodyStart: base + span.bodyStart, bodyEnd: base + span.bodyEnd})
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].start < out[j].start })
	return out
}

type mathCommandUse struct {
	at   int
	name string
}

func mathUnknownCommands(source string) []mathCommandUse {
	var out []mathCommandUse
	for i := 0; i+1 < len(source); i++ {
		if source[i] == '%' && mathPolicyBackslashRun(source, i)%2 == 0 {
			for i < len(source) && source[i] != '\n' {
				i++
			}
			continue
		}
		if source[i] != '\\' || !asciiLetter(source[i+1]) {
			continue
		}
		if mathPolicyBackslashRun(source, i+1)%2 == 0 {
			continue
		}
		start := i
		end := i + 2
		for end < len(source) && asciiLetter(source[end]) {
			end++
		}
		name := source[i+1 : end]
		i = end - 1
		if _, ok := knownTeXCommands[name]; ok {
			continue
		}
		out = append(out, mathCommandUse{at: start, name: name})
	}
	return out
}

func mathHTMLLookalike(body string) int {
	for i := 0; i+1 < len(body); i++ {
		if body[i] != '<' {
			continue
		}
		next := i + 1
		if body[next] == '/' && next+1 < len(body) {
			next++
		}
		if asciiLetter(body[next]) {
			return i
		}
	}
	return -1
}

func mathFindingAt(source string, starts []int, at int, rule, message string) Finding {
	line := sort.Search(len(starts), func(i int) bool { return starts[i] > at }) - 1
	return Finding{
		Line:     line + 1,
		Column:   utf8.RuneCountInString(source[starts[line]:at]) + 1,
		Severity: Warning,
		Rule:     rule,
		Message:  message,
	}
}
