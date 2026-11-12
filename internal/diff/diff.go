// Package diff compares the GitHub renderer with the md4c parser.
package diff

import (
	"fmt"

	"ghmd/internal/analysis"
	"ghmd/internal/github"
	"ghmd/internal/standard"
)

// Finding is one normalized HTML difference.
type Finding struct {
	Path string
	Want string
	Got  string
}

func (f Finding) String() string {
	return fmt.Sprintf("%s: cmark-gfm %q, md4c %q", f.Path, f.Want, f.Got)
}

// GitHubParsers renders with cmark-gfm and md4c and reports structural differences.
func GitHubParsers(source string) ([]Finding, error) {
	runID, err := github.NewRunID()
	if err != nil {
		return nil, err
	}
	boundary, err := github.NewBoundary()
	if err != nil {
		return nil, err
	}
	got, err := github.RenderDocument(source, runID, boundary)
	if err != nil {
		return nil, err
	}
	want, err := standard.Render(source, true)
	if err != nil {
		return nil, err
	}
	differences := analysis.CompareFragments(got, want)
	if len(differences) == 0 {
		return nil, nil
	}
	findings := make([]Finding, len(differences))
	for i, difference := range differences {
		findings[i] = Finding{
			Path: difference.Path,
			Want: difference.Want,
			Got:  difference.Got,
		}
	}
	return findings, nil
}

