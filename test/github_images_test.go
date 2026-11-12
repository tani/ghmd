package test

import (
	"reflect"
	"strings"
	"testing"

	"ghmd/internal/verify"
)

const (
	githubImageProxy    = "https://camo.githubusercontent.com/signature/encoded"
	githubImageOriginal = "https://example.com/image.png?a=1&b=2"
	githubImageHTML     = `<img src="https://camo.githubusercontent.com/signature/encoded" alt="x" data-canonical-src="https://example.com/image.png?a=1&amp;b=2" style="max-width: 100%;">`
)

func TestNormalizeGitHubImagesRestoresOriginalAndGeneratedWrapper(t *testing.T) {
	html := `<p><a target="_blank" rel="noopener noreferrer nofollow" href="` + githubImageProxy + `">` + githubImageHTML + `</a></p>`
	want := `<p><img src="https://example.com/image.png?a=1&amp;b=2" alt="x"></p>`
	if got := verify.NormalizeGitHubImages(html); !sameHTML(got, want) {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestNormalizeGitHubImagesPreservesUserLink(t *testing.T) {
	html := `<p><a href="https://example.com/page" rel="nofollow">` + githubImageHTML + `</a></p>`
	want := `<p><a href="https://example.com/page" rel="nofollow"><img src="https://example.com/image.png?a=1&amp;b=2" alt="x"></a></p>`
	if got := verify.NormalizeGitHubImages(html); !sameHTML(got, want) {
		t.Fatalf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestNormalizeGitHubImagesLeavesUnrelatedMarkupAndStyles(t *testing.T) {
	noCanonical := `<a href="` + githubImageProxy + `"><img src="` + githubImageProxy + `"></a>`
	if got := verify.NormalizeGitHubImages(noCanonical); !sameHTML(got, noCanonical) {
		t.Fatalf("unrelated image changed:\ngot:  %s\nwant: %s", got, noCanonical)
	}

	extraAttr := `<a target="_blank" rel="noopener noreferrer nofollow" href="` + githubImageProxy + `" title="kept">` + githubImageHTML + `</a>`
	if got := verify.NormalizeGitHubImages(extraAttr); !strings.Contains(got, `title="kept"`) {
		t.Fatalf("extra wrapper attribute removed:\n%s", got)
	}

	style := `<img src="` + githubImageProxy + `" data-canonical-src="` + strings.ReplaceAll(githubImageOriginal, "&", "&amp;") + `" style="width: 20px;">`
	if got := verify.NormalizeGitHubImages(style); !strings.Contains(got, `style="width: 20px;"`) {
		t.Fatalf("additional image style removed:\n%s", got)
	}
}

func sameHTML(a, b string) bool {
	return reflect.DeepEqual(verify.HTMLAST(a, false), verify.HTMLAST(b, false))
}
