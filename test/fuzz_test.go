package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"

	"gfm"
	"gfm/internal/verify"
)

var fuzzName = regexp.MustCompile(`^fuzz-[a-f0-9]{24}\.json$`)

func TestSavedFuzzRegressions(t *testing.T) {
	directory := filepath.Join("fixtures", "fuzz")
	entries, err := os.ReadDir(directory)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if !fuzzName.MatchString(entry.Name()) {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			base := filepath.Join(directory, entry.Name()[:len(entry.Name())-5])
			metadata, err := os.ReadFile(base + ".json")
			if err != nil {
				t.Fatal(err)
			}
			var meta struct {
				Failure   string `json:"failure"`
				SourceSHA string `json:"sourceSha256"`
				HTMLSHA   string `json:"htmlSha256"`
			}
			if err := json.Unmarshal(metadata, &meta); err != nil {
				t.Fatal(err)
			}
			source, err := os.ReadFile(base + ".md")
			if err != nil {
				t.Fatal(err)
			}
			if hash(source) != meta.SourceSHA {
				t.Fatalf("source hash mismatch")
			}
			html, err := os.ReadFile(base + ".html")
			if err != nil {
				t.Fatal(err)
			}
			if hash(html) != meta.HTMLSHA {
				t.Fatalf("html hash mismatch")
			}
			actual, err := gfm.Render(string(source), gfm.RenderOptions{})
			if err != nil {
				t.Fatal(err)
			}
			got := verify.HTMLAST(actual, true)
			want := verify.HTMLAST(verify.NormalizeGitHubImages(string(html)), true)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("saved regression no longer matches")
			}
		})
	}
}
