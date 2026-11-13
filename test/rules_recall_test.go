package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"ghmd/internal/lint"
)

type styleGuideExpected struct {
	Source string `json:"source"`
	Rules  []struct {
		ID       string `json:"id"`
		Expected []struct {
			Rule string `json:"rule"`
			Line int    `json:"line"`
		} `json:"expected"`
	} `json:"rules"`
}

// TestRulesBadMarkdownPerBlockRecall isolates each rule section so
// repeated expressions, cross-block formatting multisets, and unrelated
// definitions cannot hide or invent findings.
func TestRulesBadMarkdownRecall(t *testing.T) {
	directory := filepath.Join("fixtures", "lint")
	data, err := os.ReadFile(filepath.Join(directory, "rules-bad.expected.json"))
	if err != nil {
		t.Fatal(err)
	}
	var expected styleGuideExpected
	if err := json.Unmarshal(data, &expected); err != nil {
		t.Fatal(err)
	}
	sourceBytes, err := os.ReadFile(filepath.Join(directory, expected.Source))
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(sourceBytes), "\n")

	content := make(map[string]string)
	known := make(map[string]bool, len(expected.Rules))
	for _, rule := range expected.Rules {
		known[rule.ID] = true
	}
	for i, line := range lines {
		if !strings.HasPrefix(line, "## ") {
			continue
		}
		id := strings.TrimSpace(strings.TrimPrefix(line, "## "))
		if !known[id] {
			continue
		}
		start := i + 2
		if start > len(lines) {
			start = len(lines)
		}
		end := len(lines)
		for j := i + 1; j < len(lines); j++ {
			if strings.HasPrefix(lines[j], "## ") && known[strings.TrimSpace(strings.TrimPrefix(lines[j], "## "))] {
				end = j
				break
			}
		}
		content[id] = strings.TrimRight(strings.Join(lines[start:end], "\n"), "\n")
	}

	inScope, detected, wantTotal, matched, gotTotal := 0, 0, 0, 0, 0
	var misses, extras []string
	var undetected []string
	for _, rule := range expected.Rules {
		block, ok := content[rule.ID]
		if !ok {
			t.Fatalf("missing section %s", rule.ID)
		}
		findings, err := lint.Source(block)
		if err != nil {
			t.Fatalf("%s: %v", rule.ID, err)
		}
		gotTotal += len(findings)
		remaining := make([]lint.Finding, len(findings))
		copy(remaining, findings)
		if len(rule.Expected) > 0 {
			inScope++
		}
		ruleMatched := len(rule.Expected) > 0
		for _, want := range rule.Expected {
			wantTotal++
			line := want.Line
			index := -1
			for i, finding := range remaining {
				if finding.Rule == want.Rule && finding.Line == line {
					index = i
					break
				}
			}
			if index < 0 {
				ruleMatched = false
				misses = append(misses, rule.ID+" "+want.Rule)
				continue
			}
			remaining = append(remaining[:index], remaining[index+1:]...)
			matched++
		}
		if len(rule.Expected) > 0 && ruleMatched {
			detected++
		} else if len(rule.Expected) > 0 {
			undetected = append(undetected, rule.ID)
		}
		for _, extra := range remaining {
			extras = append(extras, rule.ID+" "+extra.Rule)
		}
	}

	recall := float64(matched) / float64(wantTotal)
	if inScope != len(expected.Rules) {
		t.Fatalf("coverage model mismatch: %d automatic rules, want %d", inScope, len(expected.Rules))
	}
	t.Logf("per-block: %d rules, %d detected, undetected %v", len(expected.Rules), detected, undetected)
	t.Logf("per-block finding recall %.3f (%d/%d); additional related findings in this deliberately bad document: %d",
		recall, matched, wantTotal, len(extras))
	t.Logf("rule coverage %.3f (%d/%d)", float64(detected)/float64(inScope), detected, inScope)

	if matched != wantTotal {
		t.Fatalf("per-block recall gaps: %v", misses)
	}
}
