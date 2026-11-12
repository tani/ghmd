package analysis

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

// Difference describes one structural difference between two HTML fragments.
// Want and Got contain the rendered left- and right-hand nodes, respectively.
type Difference struct {
	Path string
	Want string
	Got  string
}

type fragmentNode struct {
	tag      string
	attrs    string
	text     string
	children []fragmentNode
}

// CompareFragments compares the normalized element trees of two HTML fragments.
// Attribute order and layout whitespace do not affect the comparison.
func CompareFragments(left, right string) []Difference {
	leftNodes, rightNodes := parseFragment(left), parseFragment(right)
	var differences []Difference
	var walk func(path string, a, b []fragmentNode)
	walk = func(path string, a, b []fragmentNode) {
		for i := 0; i < len(a) || i < len(b); i++ {
			childPath := fmt.Sprintf("%s/%d", path, i)
			if i >= len(a) {
				differences = append(differences, Difference{Path: childPath, Want: "<missing>", Got: renderFragmentNode(b[i])})
				continue
			}
			if i >= len(b) {
				differences = append(differences, Difference{Path: childPath, Want: renderFragmentNode(a[i]), Got: "<missing>"})
				continue
			}
			if a[i].tag != b[i].tag || a[i].attrs != b[i].attrs || a[i].text != b[i].text {
				differences = append(differences, Difference{Path: childPath, Want: renderFragmentNode(a[i]), Got: renderFragmentNode(b[i])})
				continue
			}
			walk(childPath, a[i].children, b[i].children)
		}
	}
	walk("", leftNodes, rightNodes)
	return differences
}

func parseFragment(fragment string) []fragmentNode {
	nodes, err := Paragraphs(fragment)
	if err != nil {
		return nil
	}
	var convert func(*html.Node) fragmentNode
	convert = func(node *html.Node) fragmentNode {
		if node.Type == html.TextNode {
			return fragmentNode{text: strings.Join(strings.Fields(node.Data), " ")}
		}
		if node.Type != html.ElementNode {
			return fragmentNode{}
		}
		attrs := make([]string, 0, len(node.Attr))
		for _, attr := range node.Attr {
			attrs = append(attrs, attr.Key+"="+attr.Val)
		}
		sort.Strings(attrs)
		result := fragmentNode{tag: node.Data, attrs: strings.Join(attrs, " ")}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			converted := convert(child)
			if converted.tag != "" || converted.text != "" {
				result.children = append(result.children, converted)
			}
		}
		return result
	}
	result := make([]fragmentNode, 0, len(nodes))
	for _, node := range nodes {
		converted := convert(node)
		if converted.tag != "" || converted.text != "" {
			result = append(result, converted)
		}
	}
	return result
}

func renderFragmentNode(node fragmentNode) string {
	if node.text != "" {
		return node.text
	}
	if node.attrs != "" {
		return "<" + node.tag + " " + node.attrs + ">"
	}
	return "<" + node.tag + ">"
}
