package test

import (
	"strings"
	"testing"

	"gfm/internal/github"
)

func TestValidateIdentifiers(t *testing.T) {
	valid := "01234567-89ab-cdef-0123-456789abcdef"
	if err := github.Validate("x", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", valid); err != nil {
		t.Fatal(err)
	}
	for _, runID := range []string{
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\x00",
	} {
		if err := github.Validate("x", runID, valid); err == nil {
			t.Fatalf("accepted run ID %q", runID)
		}
	}
	for _, boundary := range []string{
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		valid[:35] + ">",
		valid + "\x00",
	} {
		if err := github.Validate("x", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", boundary); err == nil {
			t.Fatalf("accepted boundary %q", boundary)
		}
	}
}

func TestSourceLimit(t *testing.T) {
	valid := "01234567-89ab-cdef-0123-456789abcdef"
	runID := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	if err := github.Validate(repeat("x", 1000001), runID, valid); err == nil {
		t.Fatal("accepted oversized source")
	}
}

func TestRenderDocumentEmptyAndUnicode(t *testing.T) {
	const boundary = "01234567-89ab-cdef-0123-456789abcdef"
	const runID = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	for _, test := range []struct {
		name   string
		source string
		want   string
	}{
		{name: "empty input", source: "", want: ""},
		{name: "Unicode and NUL", source: "日本語\x00😀", want: "<p>日本語�😀</p>"},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, err := github.RenderDocument(test.source, runID, boundary)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Fatalf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestRenderDocumentRejectsInvalidIdentifiers(t *testing.T) {
	const validBoundary = "01234567-89ab-cdef-0123-456789abcdef"
	const validRunID = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	// Go's string type replaces the TypeScript null/undefined/non-string cases;
	// empty, malformed, and NUL-containing strings exercise the equivalent validation.
	for _, test := range []struct {
		name      string
		runID     string
		boundary  string
		wantError string
	}{
		{name: "empty run ID", runID: "", boundary: validBoundary, wantError: "runId must be 32"},
		{name: "short run ID", runID: "bad", boundary: validBoundary, wantError: "runId must be 32"},
		{name: "NUL in run ID", runID: validRunID[:31] + "\x00", boundary: validBoundary, wantError: "runId must be 32"},
		{name: "empty boundary", runID: validRunID, boundary: "", wantError: "Invalid render boundary"},
		{name: "short boundary", runID: validRunID, boundary: "bad", wantError: "Invalid render boundary"},
		{name: "NUL in boundary", runID: validRunID, boundary: validBoundary[:35] + "\x00", wantError: "Invalid render boundary"},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := github.RenderDocument("x", test.runID, test.boundary)
			if err == nil || !strings.Contains(err.Error(), test.wantError) {
				t.Fatalf("error = %v, want one containing %q", err, test.wantError)
			}
		})
	}
}

// TypeScript production-module export-surface assertions are intentionally not
// ported: Go package exports are checked statically and have no Object.keys API.

func repeat(value string, count int) string {
	out := make([]byte, 0, len(value)*count)
	for i := 0; i < count; i++ {
		out = append(out, value...)
	}
	return string(out)
}
