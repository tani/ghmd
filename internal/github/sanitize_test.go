package github

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func sanitizeStandalone(t *testing.T, source string) string {
	t.Helper()
	document, err := parseDOM(source, false, false)
	if err != nil {
		t.Fatalf("parse failed for %q: %v", source, err)
	}
	output, err := sanitizeDOM(document.root, "")
	if err != nil {
		t.Fatalf("sanitize failed for %q: %v", source, err)
	}
	assertStandaloneSanitizedTree(t, output)
	repeatedDocument, err := parseDOM(output, false, false)
	if err != nil {
		t.Fatalf("second parse failed for %q: %v", output, err)
	}
	repeated, err := sanitizeDOM(repeatedDocument.root, "")
	if err != nil {
		t.Fatalf("second sanitize failed for %q: %v", output, err)
	}
	if repeated != output {
		t.Fatalf("sanitize is not idempotent:\nfirst:\n%s\nsecond:\n%s", output, repeated)
	}
	return output
}

// assertStandaloneSanitizedTree reparses with an independent DOM. Template
// children are ordinary children in x/net/html and are visited recursively.
func assertStandaloneSanitizedTree(t *testing.T, fragment string) {
	t.Helper()
	context := &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body}
	nodes, err := html.ParseFragment(strings.NewReader(fragment), context)
	if err != nil {
		t.Fatal(err)
	}
	var visit func(*html.Node)
	visit = func(node *html.Node) {
		if node.Type == html.ElementNode {
			tag := strings.ToLower(node.Data)
			switch tag {
			case "script", "style", "iframe", "object", "embed", "base", "link", "meta":
				t.Fatalf("unsafe element <%s> in\n%s", tag, fragment)
			}
			for _, attr := range node.Attr {
				name := strings.ToLower(attr.Key)
				localName := name
				if attr.Namespace != "" {
					name = strings.ToLower(attr.Namespace) + ":" + name
				}
				if strings.HasPrefix(name, "on") || name == "srcdoc" || name == "formaction" ||
					name == "is" || name == "xmlns" || strings.HasPrefix(name, "xmlns:") {
					t.Fatalf("unsafe attribute %s in\n%s", name, fragment)
				}
				if name == "style" {
					t.Fatalf("non-generated style attribute in\n%s", fragment)
				}
				switch localName {
				case "href", "src", "action", "poster", "background", "cite":
					clean := strings.Map(func(r rune) rune {
						if r <= 0x20 || r == 0x7f || r == 0xa0 || r == 0x1680 || r == 0x180e ||
							r >= 0x2000 && r <= 0x2029 || r == 0x205f || r == 0x3000 || r == 0xfeff {
							return -1
						}
						return r
					}, attr.Val)
					clean = strings.ToLower(clean)
					for _, scheme := range []string{"javascript:", "vbscript:", "file:", "blob:"} {
						if strings.HasPrefix(clean, scheme) {
							t.Fatalf("unsafe URL %q in\n%s", attr.Val, fragment)
						}
					}
					if (tag != "img" || localName != "src") && strings.HasPrefix(clean, "data:") {
						t.Fatalf("unsafe data URL %q in\n%s", attr.Val, fragment)
					}
				}
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			visit(child)
		}
	}
	for _, node := range nodes {
		visit(node)
	}
}

func TestStandaloneSanitizerRetainsSafeMarkup(t *testing.T) {
	source := `<details open><summary>More</summary><input type="checkbox" checked disabled aria-label="done"><a id="note-1" href="#note-1" data-footnote-ref>note</a><math><mfrac><mi>x</mi><mn>2</mn></mfrac></math></details>`
	output := sanitizeStandalone(t, source)
	for _, wanted := range []string{
		`<details open="">`,
		`checked="" disabled="" aria-label="done"`,
		`data-footnote-ref=""`,
		`<math><mfrac><mi>x</mi><mn>2</mn></mfrac></math>`,
	} {
		if !strings.Contains(output, wanted) {
			t.Fatalf("missing %q in\n%s", wanted, output)
		}
	}
}

func TestStandaloneSanitizerDoesNotSkipAttributesAfterRemovedNodes(t *testing.T) {
	for _, prefix := range []string{"<!--removed-->", "<unknown>keep</unknown>", "<script>removed</script>"} {
		t.Run(prefix, func(t *testing.T) {
			source := prefix + `<div><img src=x onerror=alert(1)><a href="javascript:alert(2)" onclick=alert(3)>bad</a><template><img src=x onload=alert(4)></template></div>`
			output := sanitizeStandalone(t, source)
			lower := strings.ToLower(output)
			for _, unsafe := range []string{"onerror", "onclick", "onload", "javascript:"} {
				if strings.Contains(lower, unsafe) {
					t.Fatalf("unsafe text %q remains in\n%s", unsafe, output)
				}
			}
		})
	}
	// High-level rendering of these attacks is covered in test/sanitizer_test.go.
}

func TestStandaloneSanitizerValidatesURLs(t *testing.T) {
	unsafe := []string{
		"javascript:alert(1)",
		"jav&#x61;script&#58;alert(1)",
		"java&#9;script:alert(1)",
		"\r\nJaVaScRiPt:alert(1)",
		"java\u2003script:alert(1)",
		"\ufeffjavascript:alert(1)",
		"vbscript:msgbox(1)",
		"file:///etc/passwd",
		"data:text/html,<script>alert(1)</script>",
		"blob:https://example.com/id",
	}
	for _, url := range unsafe {
		t.Run(fmt.Sprintf("unsafe_%q", url), func(t *testing.T) {
			output := sanitizeStandalone(t, `<a href="`+url+`">x</a><img src="`+url+`"><form action="`+url+`"></form>`)
			for _, attribute := range []string{"href=", "src=", "action="} {
				if strings.Contains(output, attribute) {
					t.Fatalf("unsafe %s remains for %q:\n%s", attribute, url, output)
				}
			}
		})
	}
	for _, url := range []string{
		"",
		"#part",
		"/path",
		"../image.png",
		"//example.com/a",
		"https://example.com/x?a=1&amp;b=2",
		"MAILTO:a@example.com",
		"tel:+123",
	} {
		t.Run(fmt.Sprintf("safe_%q", url), func(t *testing.T) {
			if output := sanitizeStandalone(t, `<a href="`+url+`">x</a>`); !strings.Contains(output, "href=") {
				t.Fatalf("safe URL %q was removed:\n%s", url, output)
			}
		})
	}
}

func TestStandaloneSanitizerValidatesEverySrcsetCandidate(t *testing.T) {
	if output := sanitizeStandalone(t, `<img srcset="a.png 1x, https://example.com/b.png 2x">`); !strings.Contains(output, "srcset=") {
		t.Fatalf("safe srcset was removed:\n%s", output)
	}
	for _, value := range []string{
		"a.png 1x, javascript:alert(1) 2x",
		"javascript:alert(1),a.png",
		"a.png 1x, data:text/html,attack 2x",
		"a.png 1x, b.png 2q",
	} {
		if output := sanitizeStandalone(t, `<img srcset="`+value+`">`); strings.Contains(output, "srcset=") {
			t.Fatalf("unsafe srcset %q remains:\n%s", value, output)
		}
	}
}

func TestStandaloneSanitizerDuplicateAndNamespaceAttributes(t *testing.T) {
	if got := sanitizeStandalone(t, `<a href="javascript:alert(1)" href="https://example.com">x</a>`); got != "<a>x</a>" {
		t.Fatalf("unsafe first duplicate was not retained-then-removed: %s", got)
	}
	if got := sanitizeStandalone(t, `<a href="https://example.com" href="javascript:alert(1)">x</a>`); !strings.Contains(got, `href="https://example.com"`) {
		t.Fatalf("safe first duplicate was not kept: %s", got)
	}
	source := `<p style="color:red" onclick="x()" is="evil" xmlns="x" xmlns:xlink="x" xlink:href="javascript:x()">ok</p>`
	if got := sanitizeStandalone(t, source); got != "<p>ok</p>" {
		t.Fatalf("unsafe attributes remain:\ngot:  %s\nwant: <p>ok</p>", got)
	}
	source = `<form id="attributes"><input name="__proto__"><p id="safe-id">ok</p></form>`
	if got := sanitizeStandalone(t, source); got != `<form><input><p id="safe-id">ok</p></form>` {
		t.Fatalf("prototype or id filtering mismatch:\n%s", got)
	}
}

func TestStandaloneSanitizerMalformedForeignMarkupIsIdempotent(t *testing.T) {
	attacks := []string{
		`<svg><foreignObject><iframe srcdoc="<script>alert(1)</script>"></iframe></foreignObject></svg>`,
		`<math><annotation-xml encoding="text/html"><img src=x onerror=alert(1)></annotation-xml></math>`,
		`<math><mtext><table><mglyph><style><!--</style><img title="--><img src=x onerror=alert(1)>">`,
		`<svg><style><a title="</style><img src=x onerror=alert(1)>">x</a></style></svg>`,
		"<template><template><svg><g onload=alert(1)></g></svg><img src=x onerror=alert(1)></template></template>",
		"<table><caption><template><textarea></textarea><img src=x onerror=alert(1)></template></caption></table>",
		`<noscript><p title="</noscript><img src=x onerror=alert(1)>">x</p></noscript>`,
		"<textarea>&lt;/textarea>&lt;img src=x onerror=alert(1)></textarea>",
		`<div title="&quot; onmouseover=alert(1) &lt;script&gt;">text</div>`,
		`<object><form action="javascript:alert(1)"><input name=action></form></object>`,
		"<iframe><img src=x onerror=alert(1)></iframe><p>safe</p>",
	}
	for i, attack := range attacks {
		t.Run(fmt.Sprintf("attack_%d", i+1), func(t *testing.T) {
			sanitizeStandalone(t, attack)
		})
	}
}

func TestStandaloneSanitizerDeepTemplateCombinations(t *testing.T) {
	containers := []string{"div", "span", "template", "math", "mtext", "table", "select", "unknown"}
	for _, outer := range containers {
		for _, inner := range containers {
			t.Run(outer+"/"+inner, func(t *testing.T) {
				source := `<` + outer + `><!--first--><` + inner + `><img src=x onerror=alert(1)><a href="jav&#x61;script:alert(1)">bad</a></` + inner + `></` + outer + `>`
				output := sanitizeStandalone(t, source)
				lower := strings.ToLower(output)
				if strings.Contains(lower, "onerror=") || strings.Contains(lower, `href="javascript:`) {
					t.Fatalf("unsafe combination %s/%s produced:\n%s", outer, inner, output)
				}
			})
		}
	}
}

func TestStandaloneSanitizerRejectsDepthAndMarkerCollisions(t *testing.T) {
	if _, err := domSanitize(strings.Repeat("<template>", 520), ""); err == nil {
		t.Fatal("accepted excessively nested templates")
	}
	if _, err := domSanitize("<p>a<!--x-->b<!--y-->c</p>", "--><img src=x onerror=x()>"); err == nil {
		t.Fatal("accepted a marker that collides with comment text")
	}

	// Null and marker type/length errors belonged to the removed Bun FFI
	// wrapper. Go strings make null/type cases inapplicable, and identifier
	// length validation is covered by TestValidateIdentifiers.
}
