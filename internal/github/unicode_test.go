package github

import (
	"strings"
	"testing"
)

const (
	unicodeTestRunID    = "0123456789abcdef0123456789abcdef"
	unicodeTestBoundary = "01234567-89ab-cdef-0123-456789abcdef"
)

func TestUnicodeLowerContextAndSpecialMappings(t *testing.T) {
	tests := []struct{ source, want string }{
		{"ΟΣ", "ος"},
		{"ΟΣΑ", "οσα"},
		{"ΣΣ", "σς"},
		{"AΣ\u0301", "aς\u0301"},
		{"AΣ\u0301B", "aσ\u0301b"},
		{"\u0345Σ", "\u0345σ"},
		{"AΣ\u0345", "aς\u0345"},
		{"İẞ𐐀𐐧", "i\u0307ß\U00010428\U0001044f"},
		// Pinned Unicode 17 additions; Go's unicode tables may lag behind.
		{"\u1C89", "\u1C8A"},
		{"\uA7CB", "\u0264"},
		{"\U00010D50", "\U00010D70"},
		{"\x00ABC\x00", "\x00abc\x00"},
		{"", ""},
	}
	for _, test := range tests {
		if got := unicodeLower(test.source); got != test.want {
			t.Errorf("unicodeLower(%q) = %q, want %q", test.source, got, test.want)
		}
	}
}

func TestUnicodeLowerFootnoteIDsInNativeHTML(t *testing.T) {
	tests := []struct{ id, lower string }{
		{"ΟΣ", "ος"},
		{"ΟΣΑ", "οσα"},
		{"AΣ\u0301", "aς\u0301"},
		{"AΣ\u0301B", "aσ\u0301b"},
		{"İẞ𐐀", "i\u0307ß\U00010428"},
	}
	for _, test := range tests {
		html, err := RenderDocument(`<a id="`+test.id+`" data-footnote-ref>x</a>`, unicodeTestRunID, unicodeTestBoundary)
		if err != nil {
			t.Fatalf("render %q: %v", test.id, err)
		}
		want := `id="user-content-` + test.lower + `-` + unicodeTestRunID + `"`
		if !strings.Contains(html, want) {
			t.Errorf("rendered ID %q does not contain %q:\n%s", test.id, want, html)
		}
	}
}
