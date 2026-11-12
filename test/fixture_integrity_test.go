package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// The old github-reference suite exercised the removed compare/live-fetch
// client, so its HTTP request and token-redaction cases have no Go equivalent.

func TestFixtureManifestIntegrityAndCaseParity(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("fixtures", "e2e", "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var manifest struct {
		Fixtures []struct {
			ID        string `json:"id"`
			SourceSHA string `json:"sourceSha256"`
			Cases     []struct {
				ID string `json:"id"`
			} `json:"cases"`
		} `json:"fixtures"`
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		t.Fatal(err)
	}

	fixtureIDs := make(map[string]bool, len(manifest.Fixtures))
	caseIDs := make(map[string]bool)
	sourceHashes := make(map[string]bool, len(manifest.Fixtures))
	for _, fixture := range manifest.Fixtures {
		if fixtureIDs[fixture.ID] {
			t.Fatalf("duplicate fixture identity %q", fixture.ID)
		}
		fixtureIDs[fixture.ID] = true
		if sourceHashes[fixture.SourceSHA] {
			t.Fatalf("duplicate source hash %q", fixture.SourceSHA)
		}
		sourceHashes[fixture.SourceSHA] = true

		fixtureIsCaptured := false
		for _, captured := range fixture.Cases {
			if caseIDs[captured.ID] {
				t.Fatalf("duplicate capture case identity %q", captured.ID)
			}
			caseIDs[captured.ID] = true
			if captured.ID == fixture.ID {
				fixtureIsCaptured = true
			}
		}
		if !fixtureIsCaptured {
			t.Errorf("fixture %q is absent from the manifest's capture cases", fixture.ID)
		}
	}

	directory := filepath.Join("fixtures", "e2e", "cases")
	entries, err := os.ReadDir(directory)
	if err != nil {
		t.Fatal(err)
	}
	files := make(map[string]map[string]bool)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		extension := filepath.Ext(entry.Name())
		if extension != ".md" && extension != ".html" {
			continue
		}
		id := strings.TrimSuffix(entry.Name(), extension)
		if files[id] == nil {
			files[id] = make(map[string]bool)
		}
		files[id][extension] = true
		if !fixtureIDs[id] {
			t.Errorf("case file %q has no manifest fixture", entry.Name())
		}
	}
	for id := range fixtureIDs {
		if !files[id][".md"] || !files[id][".html"] {
			t.Errorf("manifest fixture %q lacks a Markdown/HTML case pair", id)
		}
	}
}
