// Package diff compares the GitHub renderer with the md4c parser.
package diff

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"ghmd/internal/github"
	"ghmd/internal/standard"
)

// Finding is one normalized HTML difference.
type Finding struct {
	Path string
	Want string
	Got  string
}

func (f Finding) String() string {
	return fmt.Sprintf("%s: cmark-gfm %q, md4c %q", f.Path, f.Want, f.Got)
}

// GitHubParsers renders with cmark-gfm and md4c and reports structural differences.
func GitHubParsers(source string) ([]Finding, error) {
	runID, err := github.NewRunID()
	if err != nil {
		return nil, err
	}
	boundary, err := github.NewBoundary()
	if err != nil {
		return nil, err
	}
	got, err := github.RenderDocument(source, runID, boundary)
	if err != nil {
		return nil, err
	}
	want, err := standard.Render(source, true)
	if err != nil {
		return nil, err
	}
	return compare(got, want), nil
}

type node struct {
	tag      string
	attrs    string
	text     string
	children []node
}

func parse(fragment string) []node {
	context := &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body}
	nodes, err := html.ParseFragment(strings.NewReader(fragment), context)
	if err != nil {
		return nil
	}
	var convert func(*html.Node) node
	convert = func(n *html.Node) node {
		if n.Type == html.TextNode {
			return node{text: strings.Join(strings.Fields(n.Data), " ")}
		}
		if n.Type != html.ElementNode {
			return node{}
		}
		var attrs []string
		for _, a := range n.Attr {
			attrs = append(attrs, a.Key+"="+a.Val)
		}
		sort.Strings(attrs)
		result := node{tag: n.Data, attrs: strings.Join(attrs, " ")}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			child := convert(c)
			if child.tag != "" || child.text != "" {
				result.children = append(result.children, child)
			}
		}
		return result
	}
	var result []node
	for _, n := range nodes {
		converted := convert(n)
		if converted.tag != "" || converted.text != "" {
			result = append(result, converted)
		}
	}
	return result
}

func compare(got, want string) []Finding {
	left, right := parse(got), parse(want)
	var findings []Finding
	var walk func(path string, a, b []node)
	walk = func(path string, a, b []node) {
		for i := 0; i < len(a) || i < len(b); i++ {
			childPath := fmt.Sprintf("%s/%d", path, i)
			if i >= len(a) {
				findings = append(findings, Finding{childPath, "<missing>", renderNode(b[i])})
				continue
			}
			if i >= len(b) {
				findings = append(findings, Finding{childPath, renderNode(a[i]), "<missing>"})
				continue
			}
			if a[i].tag != b[i].tag || a[i].attrs != b[i].attrs || a[i].text != b[i].text {
				findings = append(findings, Finding{childPath, renderNode(a[i]), renderNode(b[i])})
				continue
			}
			walk(childPath, a[i].children, b[i].children)
		}
	}
	walk("", left, right)
	return findings
}

func renderNode(n node) string {
	if n.text != "" {
		return n.text
	}
	if n.attrs != "" {
		return "<" + n.tag + " " + n.attrs + ">"
	}
	return "<" + n.tag + ">"
}
