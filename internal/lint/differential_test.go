package lint

import (
	"strings"
	"testing"

	"ghmd/internal/analysis"
)

func TestDifferentialFeatureRules(t *testing.T) {
	want, err := extractFeatures(`<p>lost words</p><pre><code>gone()</code></pre><code>inline()</code><a href="/standard">label</a><img src="pic.png" alt="picture"><input type="checkbox" checked><math-renderer>x+y</math-renderer>`, "")
	if err != nil {
		t.Fatal(err)
	}
	got, err := extractFeatures(`<a>label</a><input type="checkbox">`, "")
	if err != nil {
		t.Fatal(err)
	}
	source := "lost words\ngone()\ninline()\n/standard\npic.png\n[x]\nx+y\n"
	findings := compareFeatures(source, want, got)
	counts := make(map[string]int)
	for _, finding := range findings {
		counts[finding.Rule]++
		if finding.Severity != Warning {
			t.Fatalf("severity = %v for %v", finding.Severity, finding)
		}
		if finding.Line == 0 || finding.Column == 0 || strings.Contains(finding.Message, "unlocated") {
			t.Fatalf("feature was not located: %v", finding)
		}
	}
	for rule, count := range map[string]int{
		ruleLostText: 1, ruleLostCode: 2, ruleLostLink: 1,
		ruleLostImage: 1, ruleLostMath: 1, ruleTaskState: 1,
	} {
		if counts[rule] != count {
			t.Fatalf("count for %s = %d, want %d; findings: %v", rule, counts[rule], count, findings)
		}
	}
}

func TestDifferentialAmbiguousLocation(t *testing.T) {
	want, err := extractFeatures(`<p>duplicate</p>`, "")
	if err != nil {
		t.Fatal(err)
	}
	got := newSemanticFeatures()
	findings := compareFeatures("duplicate\nduplicate\n", want, got)
	if len(findings) != 1 || findings[0].Line != 0 || findings[0].Column != 0 || !strings.Contains(findings[0].Message, "unlocated") {
		t.Fatalf("unexpected unlocated finding: %v", findings)
	}
}

func TestDifferentialIgnoresPresentationWrappers(t *testing.T) {
	want, err := extractFeatures(`<p>same <span class="presentation">visible <em>text</em></span></p>`, "")
	if err != nil {
		t.Fatal(err)
	}
	got, err := extractFeatures(`<p>same visible <em>text</em></p>`, "")
	if err != nil {
		t.Fatal(err)
	}
	if findings := compareFeatures("same visible text", want, got); len(findings) != 0 {
		t.Fatalf("presentation-only wrappers differ: %v", findings)
	}
}

func TestSourceAcceptsSemanticallyCleanDocuments(t *testing.T) {
	tests := []string{
		"# Title\n\nText with **strong**, `code`, [a link](https://example.com), and ![an image](image.png).\n\n- [ ] open\n- [x] done\n",
		"<details>\n<summary>Read K</summary>\n\nText<br><a href=\"https://example.com\">safe</a>\n</details>",
		"| a | b |\n| --- | --- |\n| 1 | 2 |\n",
	}
	for _, source := range tests {
		findings, err := Source(source)
		if err != nil {
			t.Fatal(err)
		}
		if len(findings) != 0 {
			t.Fatalf("unexpected findings for %q: %v", source, findings)
		}
	}
}

func TestSourceReportsRendererRegressions(t *testing.T) {
	tests := []struct {
		name   string
		source string
		rule   string
	}{
		{name: "text", source: `<object>fallback text</object>`, rule: ruleLostText},
		{name: "code", source: `<object><code>hidden()</code></object>`, rule: ruleLostCode},
		{name: "link", source: `<object><a href="https://example.com">label</a></object>`, rule: ruleLostLink},
		{name: "image", source: `<object><img src="https://example.com/a.png" alt="x"></object>`, rule: ruleLostImage},
		{name: "math", source: `<math-renderer>x+y</math-renderer>`, rule: ruleLostMath},
		{name: "task", source: `> - [ ]`, rule: ruleTaskState},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			findings, err := Source(test.source)
			if err != nil {
				t.Fatal(err)
			}
			var matching []Finding
			for _, finding := range findings {
				if finding.Rule == test.rule {
					matching = append(matching, finding)
				}
			}
			if len(matching) != 1 {
				t.Fatalf("findings for %s = %v, want one %s", test.source, findings, test.rule)
			}
		})
	}
}

func TestSourceSuppressesKnownDialectNoise(t *testing.T) {
	source := "before $x$ after\n\n$$\ny\n$$\n\nref [^1]\n\n[^1]: footnote text\n\n> [!NOTE]\n> useful alert\n\n```math\nz\n```\n"
	findings, err := Source(source)
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range findings {
		if strings.HasPrefix(finding.Rule, "standard/") {
			t.Fatalf("dialect difference reported as a regression: %v", finding)
		}
	}
}

func TestSourceSortsRawAndDifferentialFindings(t *testing.T) {
	findings, err := Source("é <object>fallback text</object>")
	if err != nil {
		t.Fatal(err)
	}
	var previous Finding
	for i, finding := range findings {
		if i > 0 {
			if finding.Line < previous.Line || finding.Line == previous.Line && finding.Column < previous.Column ||
				finding.Line == previous.Line && finding.Column == previous.Column && finding.Rule < previous.Rule {
				t.Fatalf("findings are not sorted: %v", findings)
			}
		}
		previous = finding
	}
}

func TestVisibleTextFragmentsDoNotHideCurrency(t *testing.T) {
	features, err := analysis.SemanticFeatures(`<p>cost $5 and $6 total</p>`)
	if err != nil {
		t.Fatal(err)
	}
	var fragments []string
	for _, feature := range features {
		if feature.Kind == analysis.FeatureKind(ruleLostText) {
			fragments = append(fragments, feature.Search)
		}
	}
	if len(fragments) != 1 || fragments[0] != "cost $5 and $6 total" {
		t.Fatalf("currency text was rewritten: %q", fragments)
	}
}

func TestCompareFeaturesReportsTableShapeDifference(t *testing.T) {
	want, err := extractFeatures(`<table><tr><td>a</td><td>b</td></tr></table>`, "")
	if err != nil {
		t.Fatal(err)
	}
	got, err := extractFeatures(`<table><tr><td>a</td></tr></table>`, "")
	if err != nil {
		t.Fatal(err)
	}
	var structure []Finding
	for _, finding := range compareFeatures("a b", want, got) {
		if finding.Rule == ruleLostStructure {
			structure = append(structure, finding)
		}
	}
	if len(structure) != 1 || !strings.Contains(structure[0].Message, "1 rows x 2 columns") {
		t.Fatalf("table shape difference not reported: %v", structure)
	}

	sameWant, err := extractFeatures(`<table><tr><td>a</td><td>b</td></tr></table>`, "")
	if err != nil {
		t.Fatal(err)
	}
	sameGot, err := extractFeatures(`<table><tr><th>a</th><th>b</th></tr></table>`, "")
	if err != nil {
		t.Fatal(err)
	}
	for _, finding := range compareFeatures("a b", sameWant, sameGot) {
		if finding.Rule == ruleLostStructure {
			t.Fatalf("identical table shape reported: %v", finding)
		}
	}
}
