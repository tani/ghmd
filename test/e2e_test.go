package test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"

	"gfm"
	"gfm/internal/verify"
)

type manifest struct {
	Fixtures []struct {
		ID        string `json:"id"`
		SourceSHA string `json:"sourceSha256"`
		HTMLSAH   string `json:"htmlSha256"`
	} `json:"fixtures"`
}

var runIDPattern = regexp.MustCompile(`(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(["#\s])`)

func TestGitHubReferences(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("fixtures", "e2e", "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m manifest
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	for _, fixture := range m.Fixtures {
		t.Run(fixture.ID, func(t *testing.T) {
			directory := filepath.Join("fixtures", "e2e", "cases")
			source, err := os.ReadFile(filepath.Join(directory, fixture.ID+".md"))
			if err != nil {
				t.Fatal(err)
			}
			expected, err := os.ReadFile(filepath.Join(directory, fixture.ID+".html"))
			if err != nil {
				t.Fatal(err)
			}
			if hash(source) != fixture.SourceSHA || hash(expected) != fixture.HTMLSAH {
				t.Fatalf("fixture hash mismatch for %s", fixture.ID)
			}
			runID := runIDPattern.FindStringSubmatch(string(expected))
			options := gfm.RenderOptions{}
			if len(runID) > 1 {
				options.RunID = runID[1]
			}
			actual, err := gfm.Render(string(source), options)
			if err != nil {
				t.Fatal(err)
			}
			want := verify.HTMLAST(verify.NormalizeGitHubImages(string(expected)), true)
			got := verify.HTMLAST(actual, true)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("HTML AST mismatch\ngot:  %#v\nwant: %#v", got, want)
			}
		})
	}
}

func hash(value []byte) string {
	sum := sha256.Sum256(value)
	return hex.EncodeToString(sum[:])
}
