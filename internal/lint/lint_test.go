package lint

import (
	"slices"
	"strings"
	"testing"
)

func lintSource(t *testing.T, source string) []Finding {
	t.Helper()
	findings, err := rawHTML(source)
	if err != nil {
		t.Fatal(err)
	}
	sortFindings(findings)
	return findings
}

func findingRules(findings []Finding) []string {
	rules := make([]string, len(findings))
	for i, finding := range findings {
		rules[i] = finding.Rule
	}
	return rules
}

func requireRules(t *testing.T, source string, want ...string) []Finding {
	t.Helper()
	findings := lintSource(t, source)
	if got := findingRules(findings); !slices.Equal(got, want) {
		t.Fatalf("rules = %v, want %v\nfindings:\n%v", got, want, findings)
	}
	return findings
}

func TestSeverityAndFindingStrings(t *testing.T) {
	if Error.String() != "error" || Warning.String() != "warning" {
		t.Fatalf("unexpected severity strings: %q, %q", Error, Warning)
	}
	finding := Finding{Line: 2, Column: 4, Severity: Warning, Rule: "html/stray-close", Message: "stray closing tag </p>"}
	if got, want := finding.String(), "2:4: warning html/stray-close: stray closing tag </p>"; got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}

func TestSafeHTMLIsClean(t *testing.T) {
	requireRules(t, `<details><summary>Read <kbd>K</kbd></summary>Text<br><a href="https://example.com">safe</a></details>

<img src="data:image/png;base64,AA==">
<source src="DATA:image/webp;base64,AA==">`)
}

func TestNonRawHTMLIsIgnored(t *testing.T) {
	requireRules(t, "`<script onclick=\"alert(1)\">`\n\n```html\n<script>alert(1)</script>\n```\n\n<!-- <iframe src=\"javascript:alert(1)\"> -->\n<!DOCTYPE html>\n<?processing <script>?>")
}

func TestRemovedContainerAndRawText(t *testing.T) {
	findings := requireRules(t, `<SCRIPT><input onclick="not markup"></SCRIPT>`, ruleRemovedContainer)
	if findings[0].Severity != Error || !strings.Contains(findings[0].Message, "script") {
		t.Fatalf("unexpected finding: %v", findings[0])
	}
}

func TestUnsafeAttribute(t *testing.T) {
	findings := requireRules(t, `<a href="/safe" OnClick="alert(1)">link</a>`, ruleUnsafeAttribute)
	if findings[0].Severity != Error || !strings.Contains(findings[0].Message, "onclick") {
		t.Fatalf("unexpected finding: %v", findings[0])
	}
}

func TestUnsafeURLs(t *testing.T) {
	findings := requireRules(t, `<a href="javascript:alert(1)">one</a>
<a href="j&#x61;vascript&colon;alert(1)">two</a>
<a href="java&#x0A;script:alert(1)">three</a>
<img src="data:text/html;base64,PGI+">`,
		ruleUnsafeURL, ruleUnsafeURL, ruleUnsafeURL, ruleUnsafeURL)
	for _, finding := range findings {
		if finding.Severity != Error || !strings.Contains(finding.Message, "URL") {
			t.Fatalf("unexpected finding: %v", finding)
		}
	}
	if !strings.Contains(findings[1].Message, "javascript:") {
		t.Fatalf("entity-obfuscated URL was not decoded: %v", findings[1])
	}
}

func TestSafeDataURLNeedsImageMIMEAndContext(t *testing.T) {
	requireRules(t, `<img src="data:image/svg+xml,<svg/>">`)
	requireRules(t, `<video poster="data:image/png;base64,AA=="></video>`, ruleUnsafeURL)
}

func TestRawInput(t *testing.T) {
	findings := requireRules(t, `<input type="text" name="query">`, ruleInput)
	if findings[0].Severity != Warning {
		t.Fatalf("severity = %v, want warning", findings[0].Severity)
	}
}

func TestTemplate(t *testing.T) {
	requireRules(t, `<template><kbd>K</kbd></template>`, ruleTemplate)
}

func TestNestedAnchor(t *testing.T) {
	requireRules(t, `<a href="/one">one <a href="/two">two</a></a>`, ruleNestedAnchor)
}

func TestBlockInParagraph(t *testing.T) {
	requireRules(t, `<p>before
<div>after</div>
after</p>`, ruleBlockInP)
}

func TestUnclosedTag(t *testing.T) {
	findings := requireRules(t, `<span>hello`, ruleUnclosedTag)
	if findings[0].Line != 1 || findings[0].Column != 1 || !strings.Contains(findings[0].Message, "swallow following Markdown") {
		t.Fatalf("unexpected finding: %v", findings[0])
	}
}

func TestStrayClose(t *testing.T) {
	requireRules(t, `hello </span>`, ruleStrayClose)
}

func TestMultibyteLocation(t *testing.T) {
	findings := requireRules(t, "safe\n\né🙂 <button ONCLICK=\"x\">x</button>", ruleUnsafeAttribute)
	if findings[0].Line != 3 || findings[0].Column != 4 {
		t.Fatalf("location = %d:%d, want 3:4", findings[0].Line, findings[0].Column)
	}
}

func TestFindingsSorted(t *testing.T) {
	findings := requireRules(t, "<span>\n\n<script>\n",
		ruleUnclosedTag, ruleUnclosedTag, ruleRemovedContainer)
	wantLocations := [][2]int{{1, 1}, {3, 1}, {3, 1}}
	for i, finding := range findings {
		if got := [2]int{finding.Line, finding.Column}; got != wantLocations[i] {
			t.Fatalf("finding %d location = %v, want %v", i, got, wantLocations[i])
		}
	}
}

func TestTaskItemsDoNotCreateInputFindings(t *testing.T) {
	requireRules(t, "- [ ] source task\n- [x] done task")
}
