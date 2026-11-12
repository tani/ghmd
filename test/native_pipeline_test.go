package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gfm"
)

func TestNativePipeline(t *testing.T) {
	directory := filepath.Join("fixtures", "native-pipeline")
	entries, err := os.ReadDir(directory)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") || entry.Name() == "README.md" {
			continue
		}
		t.Run(entry.Name(), func(t *testing.T) {
			source, err := os.ReadFile(filepath.Join(directory, entry.Name()))
			if err != nil {
				t.Fatal(err)
			}
			expected, err := os.ReadFile(filepath.Join(directory, strings.TrimSuffix(entry.Name(), ".md")+".html"))
			if err != nil {
				t.Fatal(err)
			}
			actual, err := gfm.Render(string(source), gfm.RenderOptions{RunID: strings.Repeat("a", 32)})
			if err != nil {
				t.Fatal(err)
			}
			if actual != string(expected) {
				t.Fatalf("mismatch\ngot:\n%s\nwant:\n%s", actual, expected)
			}
		})
	}
}
