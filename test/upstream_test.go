package test

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"ghmd/internal/github"
)

type upstreamManifest struct {
	Commit string            `json:"commit"`
	Files  map[string]string `json:"files"`
}

func TestUpstreamVendorHashes(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("fixtures", "cmark-gfm-upstream.json"))
	if err != nil {
		t.Fatal(err)
	}
	var manifest upstreamManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		t.Fatal(err)
	}
	if manifest.Commit != "587a12bb54d95ac37241377e6ddc93ea0e45439b" {
		t.Fatalf("unexpected commit %s", manifest.Commit)
	}
	vendor := filepath.Join("..", "third_party", "cmark-gfm")
	var actual []string
	for _, dir := range []string{"src", "extensions", "generated"} {
		entries, err := os.ReadDir(filepath.Join(vendor, dir))
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				actual = append(actual, dir+"/"+entry.Name())
			}
		}
	}
	actual = append(actual, "COPYING")
	var expected []string
	for path := range manifest.Files {
		expected = append(expected, path)
	}
	sort.Strings(actual)
	sort.Strings(expected)
	if strings.Join(actual, ",") != strings.Join(expected, ",") {
		t.Fatalf("vendor file list changed:\n%v\n%v", actual, expected)
	}
	for path, want := range manifest.Files {
		bytes, err := os.ReadFile(filepath.Join(vendor, path))
		if err != nil {
			t.Fatal(err)
		}
		sum := sha256.Sum256(bytes)
		if hex.EncodeToString(sum[:]) != want {
			t.Fatalf("hash mismatch for %s", path)
		}
	}
}

func finish(html string) string {
	return github.FinalizeSerializedHTML(html, "00000000000000000000000000000000", "")
}

func TestFinalizeRules(t *testing.T) {
	cases := []struct{ name, input, want string }{
		{"table wrap", `<style>a{content:"<table>"}</style><table class="x" role="grid"><tr><td>1</td></tr></table>`,
			`<style>a{content:"<table>"}</style><markdown-accessiblity-table><table class="x" role="table"><tr><td>1</td></tr></table></markdown-accessiblity-table>`},
		{"template table", `<template><table><tbody><tr><td>x</td></tr></tbody></table></template><table></table>`,
			`<template><table><tbody><tr><td>x</td></tr></tbody></table></template><markdown-accessiblity-table><table role="table"></table></markdown-accessiblity-table>`},
		{"code classes", `<pre lang="math" class="a a"><code class="x x">z</code></pre><template><code>z</code></template>`,
			`<pre class="a notranslate"><code class="x notranslate">z</code></pre><template><code>z</code></template>`},
		{"list math spacing", `<ul><li><math-renderer>$x$</math-renderer> y</li><li>x <em><math-renderer>$y$</math-renderer></em></li></ul>`,
			"<ul><li>\n<math-renderer>$x$</math-renderer> y</li><li>x <em><math-renderer>$y$</math-renderer></em>\n</li></ul>"},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			if got := finish(test.input); got != test.want {
				t.Fatalf("got:\n%s\nwant:\n%s", got, test.want)
			}
		})
	}
}

func TestFinalizeFootnotes(t *testing.T) {
	token := "0123456789abcdef0123456789abcdef"
	input := `<p><sup class="footnote-ref"><a href="#fn-a&amp;b" id="fnref-a" data-footnote-ref="">1</a></sup></p><section class="old" data-footnotes=""><ol><li id="fn-a"><p>note <a href="#fnref-a" data-footnote-backref="" data-footnote-backref-idx="1" class="old">↩</a></p></li></ol></section>`
	output := github.FinalizeSerializedHTML(input, token, "")
	for _, wanted := range []string{
		`href="#user-content-fn-a&amp;b-` + token + `"`,
		`id="user-content-fnref-a-` + token + `" data-footnote-ref="" aria-describedby="footnote-label"`,
		`<section data-footnotes="" class="footnotes"><h2 id="footnote-label" class="sr-only">Footnotes</h2>`,
	} {
		if !strings.Contains(output, wanted) {
			t.Fatalf("missing %q in\n%s", wanted, output)
		}
	}
}

func TestDeepListFlattening(t *testing.T) {
	deep := strings.Repeat("<ul><li>parent", 10) + `<ul><li class="task-list-item other other">child</li></ul>` + strings.Repeat("</li></ul>", 10)
	output := finish(deep)
	if strings.Count(output, "<ul>") != 10 {
		t.Fatalf("expected 10 lists in\n%s", output)
	}
	if !strings.Contains(output, `<li class="other">child</li>`) || strings.Contains(output, "task-list-item") {
		t.Fatalf("class normalization failed:\n%s", output)
	}
}
