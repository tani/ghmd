package test

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"yagfm"
)

func TestSanitizerNeutralizesAttacks(t *testing.T) {
	attacks := []string{
		`<svg><foreignObject><iframe srcdoc="<script>alert(1)</script>"></iframe></foreignObject></svg>`,
		`<math><annotation-xml encoding="text/html"><img src=x onerror=alert(1)></annotation-xml></math>`,
		`<math><mtext><table><mglyph><style><!--</style><img title="--><img src=x onerror=alert(1)>">`,
		`<svg><style><a title="</style><img src=x onerror=alert(1)>">x</a></style></svg>`,
		`<template><template><svg><g onload=alert(1)></g></svg><img src=x onerror=alert(1)></template></template>`,
		`<table><caption><template><textarea></textarea><img src=x onerror=alert(1)></template></caption></table>`,
		`<noscript><p title="</noscript><img src=x onerror=alert(1)>">x</p></noscript>`,
		"<textarea>&lt;/textarea>&lt;img src=x onerror=alert(1)></textarea>",
		`<div title="&quot; onmouseover=alert(1) &lt;script&gt;">text</div>`,
		`<object><form action="javascript:alert(1)"><input name=action></form></object>`,
		`<iframe><img src=x onerror=alert(1)></iframe><p>safe</p>`,
		`<!--removed--><div><img src=x onerror=alert(1)><a href="javascript:alert(2)" onclick=alert(3)>bad</a><template><img src=x onload=alert(4)></template></div>`,
	}
	for _, attack := range attacks {
		output, err := yagfm.Render(attack, yagfm.RenderOptions{})
		if err != nil {
			t.Fatalf("%v: %s", err, attack)
		}
		assertSafeTree(t, output)
	}
}

func assertSafeTree(t *testing.T, fragment string) {
	t.Helper()
	context := &html.Node{Type: html.ElementNode, Data: "div", DataAtom: atom.Div}
	nodes, err := html.ParseFragment(strings.NewReader(fragment), context)
	if err != nil {
		t.Fatal(err)
	}
	var visit func(*html.Node)
	visit = func(node *html.Node) {
		if node.Type == html.ElementNode {
			switch strings.ToLower(node.Data) {
			case "script", "style", "iframe", "object", "embed", "base", "link", "meta":
				t.Fatalf("unsafe element <%s> in\n%s", node.Data, fragment)
			}
			for _, attr := range node.Attr {
				name := strings.ToLower(attr.Key)
				if strings.HasPrefix(name, "on") || name == "srcdoc" || name == "formaction" || name == "is" ||
					name == "xmlns" || strings.HasPrefix(name, "xmlns:") {
					t.Fatalf("unsafe attribute %s in\n%s", attr.Key, fragment)
				}
				if name == "href" || name == "src" || name == "action" || name == "poster" || name == "background" || name == "cite" {
					clean := strings.ToLower(strings.Map(func(r rune) rune {
						if r <= 0x20 || r == 0x7f {
							return -1
						}
						return r
					}, attr.Val))
					if strings.HasPrefix(clean, "javascript:") || strings.HasPrefix(clean, "vbscript:") ||
						strings.HasPrefix(clean, "file:") || strings.HasPrefix(clean, "blob:") {
						t.Fatalf("unsafe URL %q in\n%s", attr.Val, fragment)
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
