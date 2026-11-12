package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"gfm/internal/lint"
)

type lintCase struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Want   []struct {
		Rule   string `json:"rule"`
		Line   int    `json:"line"`
		Column int    `json:"column"`
	} `json:"want"`
}

func TestLintCorpusRecall(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("fixtures", "lint", "cases.json"))
	if err != nil {
		t.Fatal(err)
	}
	var cases []lintCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatal(err)
	}

	wantByRule := map[string]int{}
	hitByRule := map[string]int{}
	totalWant, totalHit, totalFindings := 0, 0, 0
	var misses, extras []string
	for _, test := range cases {
		findings, err := lint.Source(test.Source)
		if err != nil {
			t.Fatalf("%s: %v", test.ID, err)
		}
		totalFindings += len(findings)
		remaining := append([]lint.Finding(nil), findings...)
		for _, want := range test.Want {
			wantByRule[want.Rule]++
			totalWant++
			matched := -1
			for i, finding := range remaining {
				if finding.Rule != want.Rule || finding.Line != want.Line {
					continue
				}
				if want.Column != 0 && finding.Column != want.Column {
					continue
				}
				matched = i
				break
			}
			if matched < 0 {
				misses = append(misses, test.ID+" "+want.Rule)
				continue
			}
			remaining = append(remaining[:matched], remaining[matched+1:]...)
			hitByRule[want.Rule]++
			totalHit++
		}
		for _, finding := range remaining {
			extras = append(extras, test.ID+" "+finding.Rule)
		}
	}

	rules := make([]string, 0, len(wantByRule))
	for rule := range wantByRule {
		rules = append(rules, rule)
	}
	sort.Strings(rules)
	for _, rule := range rules {
		t.Logf("%-40s recall %d/%d", rule, hitByRule[rule], wantByRule[rule])
	}
	recall := float64(totalHit) / float64(totalWant)
	precision := float64(totalFindings-len(extras)) / float64(totalFindings)
	t.Logf("overall recall %.3f (%d/%d), precision %.3f (%d/%d findings), extras %d",
		recall, totalHit, totalWant, precision, totalHit, totalFindings, len(extras))
	if totalHit != totalWant {
		t.Fatalf("recall gaps: %v", misses)
	}
	if len(extras) > 0 {
		sort.Strings(extras)
		t.Fatalf("unexpected findings: %v", extras)
	}
}
