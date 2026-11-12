// Package ghmd renders GitHub-flavored Markdown to GitHub-compatible HTML.
package ghmd

import (
	"ghmd/internal/github"
)

// RenderOptions controls one render.
type RenderOptions struct {
	// RunID is a 32-character lowercase hexadecimal render ID. Empty generates one.
	RunID string
}

// Render converts Markdown to a sanitized GitHub-style HTML fragment.
func Render(source string, options RenderOptions) (string, error) {
	runID := options.RunID
	if runID == "" {
		var err error
		if runID, err = github.NewRunID(); err != nil {
			return "", err
		}
	}
	boundary, err := github.NewBoundary()
	if err != nil {
		return "", err
	}
	return github.RenderDocument(source, runID, boundary)
}
