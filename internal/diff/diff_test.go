package diff

import (
	"strings"
	"testing"
)

func TestGitHubParsersIdenticalInput(t *testing.T) {
	findings, err := GitHubParsers("# hi\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) != 0 {
		t.Fatalf("unexpected findings for identical output: %v", findings)
	}
}

func TestGitHubParsersReportsDialectDifferences(t *testing.T) {
	findings, err := GitHubParsers("- [x] task\n\n| a | b |\n|---|---|\n| 1 | 2 |\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) == 0 {
		t.Fatal("expected task-list and table differences")
	}
	var sawTable, sawTask bool
	for _, finding := range findings {
		if strings.Contains(finding.Want, "markdown-accessiblity-table") || strings.Contains(finding.Got, "<table") {
			sawTable = true
		}
		if strings.Contains(finding.Want, "task-list") || strings.Contains(finding.Got, "task-list") {
			sawTask = true
		}
	}
	if !sawTable || !sawTask {
		t.Fatalf("findings did not cover table and task differences: %v", findings)
	}
}

func TestCompareFindsTextDifference(t *testing.T) {
	if findings := compare("<p>a</p>", "<p>a</p>"); len(findings) != 0 {
		t.Fatalf("identical fragments differ: %v", findings)
	}
	findings := compare("<p>a</p>", "<p>b</p>")
	if len(findings) == 0 || findings[0].Want != "a" || findings[0].Got != "b" {
		t.Fatalf("text difference not reported: %v", findings)
	}
}
