package github

import (
	"errors"
	"strings"
	"testing"
)

const (
	nativeTestBoundary   = "test-boundary"
	nativeTestToken      = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	nativeTestTableMark  = "data-test-table"
	nativeTestOutputSize = 32 * 1024 * 1024
)

func nativeStageIDs() RenderIDs {
	ids := NewRenderIDs(nativeTestToken, nativeTestBoundary)
	ids.TableMarker = nativeTestTableMark
	return ids
}

func prepareNativeStage(t *testing.T, source string) (string, bool) {
	t.Helper()
	html, needsReparse := prepareHTMLStage(source, nativeTestBoundary)
	return html, needsReparse
}

func renderNativeMathStage(t *testing.T, source string) (string, bool) {
	t.Helper()
	html, needsReparse, err := renderMathStage(source, nativeStageIDs())
	if err != nil {
		t.Fatal(err)
	}
	return html, needsReparse
}

func finishCompleteResult(source string, ids RenderIDs) (string, error) {
	html := finishHTML(source, ids)
	html = strings.TrimSuffix(html, "\n")
	if len(html) > nativeTestOutputSize {
		return "", errors.New("Rendered HTML exceeds 32 MB")
	}
	return html, nil
}

func finishComplete(t *testing.T, source string, ids RenderIDs) string {
	t.Helper()
	html, err := finishCompleteResult(source, ids)
	if err != nil {
		t.Fatal(err)
	}
	return html
}

func TestHTMLPreparationPreservesTemplateTextBoundaries(t *testing.T) {
	html, needsReparse := prepareNativeStage(t, "<p>日本語</p>")
	if html != "<p>日本語</p>" || needsReparse {
		t.Fatalf("plain HTML changed: html=%q needsReparse=%v", html, needsReparse)
	}

	html, needsReparse = prepareNativeStage(t, "<p>a<template>b</template>c</p>")
	if html != "<p>a<!--test-boundary-->b<!--test-boundary-->c</p>" || needsReparse {
		t.Fatalf("template boundaries mismatch: html=%q needsReparse=%v", html, needsReparse)
	}

	html, needsReparse = prepareNativeStage(t, "<p>a<template>b</template></p><details>x</details>")
	want := "<p>ab</p><details><summary>Details</summary>x</details>"
	if html != want || !needsReparse {
		t.Fatalf("details preparation mismatch: html=%q needsReparse=%v", html, needsReparse)
	}
}

func TestTableWrappingAloneDoesNotRequestDOMReparsing(t *testing.T) {
	table := "<table><tbody><tr><td>v</td></tr></tbody></table>"
	html, needsReparse := renderNativeMathStage(t, "<p>plain"+table+"</p>")
	if needsReparse {
		t.Fatal("table wrapping alone requested reparsing")
	}
	if !strings.Contains(html, `<markdown-accessiblity-table><table data-test-table="">`) {
		t.Fatalf("wrapped table missing:\n%s", html)
	}

	html, needsReparse = renderNativeMathStage(t, "<p>$x$"+table+"</p>")
	if !needsReparse {
		t.Fatal("math rendering did not request reparsing")
	}
	if !strings.Contains(html, "<math-renderer") {
		t.Fatalf("math renderer missing:\n%s", html)
	}
	if strings.Contains(html, "<markdown-accessiblity-table>") {
		t.Fatalf("table was unexpectedly wrapped:\n%s", html)
	}
}

func TestRemovedTextBoundaryMarkersAloneDoNotRequestDOMReparsing(t *testing.T) {
	html, needsReparse := renderNativeMathStage(t, "<p>$<!--test-boundary-->x$</p>")
	if html != "<p>$x$</p>" || needsReparse {
		t.Fatalf("boundary removal mismatch: html=%q needsReparse=%v", html, needsReparse)
	}
}

func TestFinalStageRemovesRejectedAnchorsBeforeImageWrapping(t *testing.T) {
	html := finishComplete(t,
		`<a data-gfm-had-href="test-boundary"><img src="local.png"></a>`,
		nativeStageIDs())
	want := `<p><a target="_blank" rel="noopener noreferrer" href="local.png"><img src="local.png" style="max-width: 100%;"></a></p>`
	if html != want {
		t.Fatalf("final HTML mismatch\ngot:  %s\nwant: %s", html, want)
	}
}

// The Bun FFI "one native call" test is intentionally not ported: Go invokes
// the in-process stages directly and has no bun:ffi compile/invoke surface.

func TestFinalStageRemovesExactlyOneTrailingLF(t *testing.T) {
	for _, source := range []string{
		"",
		"text",
		"text\n",
		"text\n\n",
		"text\r\n",
		"\n\r\n",
		"text\n\r",
		"text\u2028",
	} {
		t.Run(strings.ReplaceAll(source, "\n", `\n`), func(t *testing.T) {
			want := strings.TrimSuffix(source, "\n")
			if got := finishComplete(t, source, nativeStageIDs()); got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

func TestFinalStageEnforcesByteLimitAfterTrimmingAndExpansion(t *testing.T) {
	// The character count is far below the byte limit; UTF-8 bytes fill it exactly.
	exact := strings.Repeat("あ", nativeTestOutputSize/3) + strings.Repeat("x", nativeTestOutputSize%3)
	if got := finishComplete(t, exact+"\n", nativeStageIDs()); got != exact {
		t.Fatal("HTML at the exact output byte limit changed")
	}

	if _, err := finishCompleteResult(exact+"x", nativeStageIDs()); err == nil || err.Error() != "Rendered HTML exceeds 32 MB" {
		t.Fatalf("oversized HTML error = %v", err)
	}

	// The input fits, but relative-image wrapping pushes the final HTML over.
	expanded := strings.Repeat("x", nativeTestOutputSize-40) + `<img src="x">`
	if len(expanded) >= nativeTestOutputSize {
		t.Fatalf("test input is already oversized: %d bytes", len(expanded))
	}
	if _, err := finishCompleteResult(expanded, nativeStageIDs()); err == nil || err.Error() != "Rendered HTML exceeds 32 MB" {
		t.Fatalf("expanded HTML error = %v", err)
	}
	if got := finishComplete(t, "ok\n", nativeStageIDs()); got != "ok" {
		t.Fatalf("small HTML after limit checks = %q", got)
	}
}
