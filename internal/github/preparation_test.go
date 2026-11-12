package github

import (
	"strings"
	"testing"
)

func TestPreparationPreservesNestedTemplateAndCommentBoundaries(t *testing.T) {
	prepared, needsReparse := prepareNativeStage(t,
		"<p>$<template>x<template>y</template></template><!-- one --><!-- two -->$</p>")
	if needsReparse {
		t.Fatal("template and comment removal requested reparsing")
	}
	if strings.Contains(prepared, "<template>") {
		t.Fatalf("template survived preparation:\n%s", prepared)
	}
	if !strings.Contains(prepared, "$<!--test-boundary-->x<!--test-boundary-->y") {
		t.Fatalf("text boundaries were not retained:\n%s", prepared)
	}

	// The expression spans separate text nodes, so it must stay literal.
	html, needsReparse := renderNativeMathStage(t, prepared)
	if html != "<p>$xy$</p>" || needsReparse {
		t.Fatalf("math across text boundaries changed: html=%q needsReparse=%v", html, needsReparse)
	}
}

func TestPreparationRetainsCommentShapedRawText(t *testing.T) {
	prepared, _ := prepareNativeStage(t,
		"<textarea><!--literal--><b>text</b></textarea><p>a<!--removed-->b</p>")
	html, needsReparse := renderNativeMathStage(t, prepared)
	want := "<textarea><!--literal--><b>text</b></textarea><p>ab</p>"
	if html != want || needsReparse {
		t.Fatalf("raw-text preservation mismatch: html=%q needsReparse=%v", html, needsReparse)
	}
}

func TestMathStageWrapsTablesAccordingToMathContext(t *testing.T) {
	math := `<math-renderer class="js-inline-math" style="display: inline-block" data-run-id="` + nativeTestToken + `">$x$</math-renderer>`
	table := "<table><tbody><tr><td>v</td></tr></tbody></table>"
	markedTable := strings.Replace(table, "<table>", `<table data-test-table="">`, 1)
	wrap := func(value string) string {
		return "<markdown-accessiblity-table>" + value + "</markdown-accessiblity-table>"
	}

	for _, test := range []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "math in paragraph",
			source: "<p>$x$" + table + "</p>",
			want:   "<p>" + math + markedTable + "</p>",
		},
		{
			name:   "plain paragraph",
			source: "<p>plain" + table + "</p>",
			want:   "<p>plain" + wrap(markedTable) + "</p>",
		},
		{
			name:   "formatting outside paragraph",
			source: "<b><p>$x$" + table + "</p></b>",
			want:   "<b><p>$x$" + wrap(markedTable) + "</p></b>",
		},
		{
			name:   "details outside paragraph",
			source: "<details><p>$x$" + table + "</p></details>",
			want:   "<details><p>" + math + markedTable + "</p></details>",
		},
		{
			name:   "formatting and details outside paragraph",
			source: "<b><details><p>$x$" + table + "</p></details></b>",
			want:   "<b><details><p>" + math + markedTable + "</p></details></b>",
		},
		{
			name:   "adjacent paragraphs",
			source: "<p>$x$" + table + "</p><p>plain" + table + "</p>",
			want:   "<p>" + math + markedTable + "</p><p>plain" + wrap(markedTable) + "</p>",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			html, _ := renderNativeMathStage(t, test.source)
			if html != test.want {
				t.Fatalf("table wrapping mismatch\ngot:  %s\nwant: %s", html, test.want)
			}
		})
	}
}
