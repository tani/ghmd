package lint

import "testing"

type expectedMathFinding struct {
	rule         string
	line, column int
}

func requireMathFindings(t *testing.T, source string, want ...expectedMathFinding) {
	t.Helper()
	findings := mathFindings(source)
	if len(findings) != len(want) {
		t.Fatalf("mathFindings(%q) = %v, want %v", source, findings, want)
	}
	for i, finding := range findings {
		expected := want[i]
		if finding.Rule != expected.rule || finding.Line != expected.line || finding.Column != expected.column {
			t.Errorf("finding %d = %v, want %s at %d:%d", i, finding, expected.rule, expected.line, expected.column)
		}
		if finding.Severity != Warning {
			t.Errorf("finding %d severity = %v, want warning", i, finding.Severity)
		}
		if finding.Message == "" {
			t.Errorf("finding %d has an empty message", i)
		}
	}
}

func TestMathNotRecognizedBoundaries(t *testing.T) {
	tests := []struct {
		name   string
		source string
		column int
	}{
		{name: "letter before", source: "a$x$", column: 2},
		{name: "space after opening", source: "$ x$", column: 1},
		{name: "space before closing", source: "$x $", column: 1},
		{name: "ampersand", source: "$x&y$", column: 1},
		{name: "letter after", source: "$x$y", column: 1},
		{name: "digit after", source: "$x$1", column: 1},
		{name: "underscore after", source: "$x$_", column: 1},
		{name: "multibyte column", source: "é🙂 $ x$", column: 4},
		{name: "second span", source: "$x$ and $ y$", column: 9},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requireMathFindings(t, test.source, expectedMathFinding{ruleMathNotRecognized, 1, test.column})
		})
	}
}

func TestMathRecognizedInlineControls(t *testing.T) {
	for _, source := range []string{
		"$x$",
		"before ($x$).",
		"before\u00a0$x+y$!",
		"$5$",
	} {
		requireMathFindings(t, source)
	}
}

func TestMathUnclosedDollar(t *testing.T) {
	requireMathFindings(t, "before $x", expectedMathFinding{ruleMathUnclosedDollar, 1, 8})
	requireMathFindings(t, "first $x\nnext $y$", expectedMathFinding{ruleMathUnclosedDollar, 1, 7})
	requireMathFindings(t, "é $open", expectedMathFinding{ruleMathUnclosedDollar, 1, 3})
}

func TestMathCurrencyControls(t *testing.T) {
	for _, source := range []string{
		"$5",
		"$5.99",
		"cost $5 and $6",
		"cost $5 and $6 total",
		"prices $5.99, $6, and $7.25",
	} {
		requireMathFindings(t, source)
	}
}

func TestMathFootnoteDefinitions(t *testing.T) {
	requireMathFindings(t, "[^a]: note $x$\n",
		expectedMathFinding{ruleMathFootnoteLiteral, 1, 12})

	source := "reference[^a]\n\n[^a]: note\n    $x + y$\n\n\t$z$\n\nafter $ok$\n"
	requireMathFindings(t, source,
		expectedMathFinding{ruleMathFootnoteLiteral, 4, 5},
		expectedMathFinding{ruleMathFootnoteLiteral, 6, 2})

	requireMathFindings(t, "[^a]: $ x$",
		expectedMathFinding{ruleMathNotRecognized, 1, 7})
	requireMathFindings(t, "[^a]: $x",
		expectedMathFinding{ruleMathUnclosedDollar, 1, 7})
}

func TestMathDisplayIsolation(t *testing.T) {
	for _, source := range []string{
		"$$\nx\n$$\n",
		"  $$  \nx\n\t$$ \n",
		"$$\nx $y$ z\n$$\nafter $x$\n",
	} {
		requireMathFindings(t, source)
	}

	tests := []struct {
		name   string
		source string
		line   int
		column int
	}{
		{name: "one line", source: "$$x$$", line: 1, column: 1},
		{name: "opening has text", source: "prefix $$\nx\n$$", line: 1, column: 8},
		{name: "opening trailing text", source: "$$ x\ny\n$$", line: 1, column: 1},
		{name: "closing has text", source: "$$\nx\ntext $$", line: 1, column: 1},
		{name: "closing trailing text", source: "$$\nx\n$$ trailing", line: 1, column: 1},
		{name: "unclosed display", source: "before $$", line: 1, column: 8},
		{name: "multibyte column", source: "é $$x$$", line: 1, column: 3},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requireMathFindings(t, test.source,
				expectedMathFinding{ruleMathDisplayIsolated, test.line, test.column})
		})
	}
}

func TestMathExcludedRegions(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{name: "backtick fence", source: "   ```tex\n$ x$\n$x\n   ```\n"},
		{name: "tilde fence", source: "~~~\n$ x$\n$x\n~~~~\n"},
		{name: "unclosed fence", source: "```\n$ x$"},
		{name: "inline code", source: "`$ x$` and ``$x $``"},
		{name: "math code", source: "$`x`$ a$` y `$"},
		{name: "HTML comment", source: "<!-- $ x$ -->\n<!--\n$x\n-->"},
		{name: "HTML attributes", source: "<span title=\"$ x$\" data-value='$x'>safe</span>"},
		{name: "escaped dollars", source: `\$ x\$ and \$open`},
		{name: "currency in excluded regions", source: "`<span title=\"$ x$\">$x$</span>`"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requireMathFindings(t, test.source)
		})
	}

	requireMathFindings(t, "`<code>$x$</code>`\n$ x$",
		expectedMathFinding{ruleMathNotRecognized, 2, 1})
	requireMathFindings(t, "<span title=\"<!-- $x$ -->\">safe</span> $ x$",
		expectedMathFinding{ruleMathNotRecognized, 1, 40})
}

func TestMathEscapedBackslashes(t *testing.T) {
	requireMathFindings(t, `\$x$`)
	requireMathFindings(t, `\\$ x$`,
		expectedMathFinding{ruleMathNotRecognized, 1, 3})
}

func TestSourceIncludesMathFindings(t *testing.T) {
	findings, err := Source("é $ x$")
	if err != nil {
		t.Fatal(err)
	}
	var boundary, lost int
	for _, finding := range findings {
		switch finding.Rule {
		case ruleMathNotRecognized:
			boundary++
		case ruleLostMath:
			lost++
		}
	}
	if boundary != 0 || lost != 1 {
		t.Fatalf("Source findings = %v, want only the lost-math comparison", findings)
	}
}

func TestSourceKeepsMathCauseWhenStandardAlsoRejects(t *testing.T) {
	findings, err := Source("a$x$")
	if err != nil {
		t.Fatal(err)
	}
	seen := false
	for _, finding := range findings {
		if finding.Rule == ruleMathNotRecognized {
			seen = true
			break
		}
	}
	if !seen {
		t.Fatalf("Source findings = %v, want one boundary finding", findings)
	}
}
