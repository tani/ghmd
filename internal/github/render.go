// Package github implements the GitHub-compatible Markdown pipeline: AST
// adaptation, HTML5 sanitization and repair, math, tables, alerts, footnotes,
// task lists, images, links, and final serialization.
package github

import (
	"errors"
	"fmt"
	"strings"

	"ghmd/internal/cmark-gfm"
)

const (
	maxSourceBytes = 1000000
	maxOutputBytes = 32 * 1024 * 1024
	runIDLength    = 32
	boundaryLength = 36
)

// RenderIDs carries the generated identifiers for one render.
type RenderIDs struct {
	RunID       string
	Boundary    string
	TableMarker string
}

// Validate checks render identifiers the way the native API did.
func Validate(source, runID, boundary string) error {
	if len(source) > maxSourceBytes {
		return ErrSourceSize
	}
	if len(runID) != runIDLength {
		return ErrRunID
	}
	for i := 0; i < len(runID); i++ {
		if !isHex(runID[i]) {
			return ErrRunID
		}
	}
	if len(boundary) != boundaryLength {
		return ErrBoundary
	}
	for i := 0; i < len(boundary); i++ {
		dash := i == 8 || i == 13 || i == 18 || i == 23
		if dash {
			if boundary[i] != '-' {
				return ErrBoundary
			}
		} else if !isHex(boundary[i]) {
			return ErrBoundary
		}
	}
	return nil
}

func isHex(c byte) bool { return c >= 'a' && c <= 'f' || c >= '0' && c <= '9' }

// NewRenderIDs derives the table marker from the boundary.
func NewRenderIDs(runID, boundary string) RenderIDs {
	return RenderIDs{RunID: runID, Boundary: boundary,
		TableMarker: "data-gfm-table-" + strings.ReplaceAll(boundary, "-", "")}
}

// RenderDocument renders Markdown to the final sanitized GitHub HTML fragment.
func RenderDocument(source, runID, boundary string) (string, error) {
	if err := Validate(source, runID, boundary); err != nil {
		return "", err
	}
	ids := NewRenderIDs(runID, boundary)
	html, err := renderMarkdown([]byte(source))
	if err != nil {
		return "", err
	}
	if len(html) > maxOutputBytes {
		return "", ErrOutputSize
	}
	html, err = domSanitize(html, ids.Boundary)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrDocument, err)
	}
	html, err = domReparse(html)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrDocument, err)
	}
	html, reparse := prepareHTMLStage(html, ids.Boundary)
	if reparse {
		if html, err = domReparse(html); err != nil {
			return "", fmt.Errorf("%w: %w", ErrDocument, err)
		}
	}
	html, reparse, err = renderMathStage(html, ids)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrDocument, err)
	}
	if reparse {
		if html, err = domReparse(html); err != nil {
			return "", fmt.Errorf("%w: %w", ErrDocument, err)
		}
	}
	html = finishHTML(html, ids)
	if len(html) > 0 && html[len(html)-1] == '\n' {
		html = html[:len(html)-1]
	}
	if len(html) > maxOutputBytes {
		return "", ErrOutputSize
	}
	return html, nil
}

func renderMarkdown(source []byte) (string, error) {
	document, err := cmarkghmd.Parse(string(source))
	if err != nil {
		return "", ErrMarkdown
	}
	defer document.Close()
	root := document.Root()
	recovered, changed := recoverLabels(root, source)
	if changed {
		document.Close()
		document, err = cmarkghmd.Parse(string(recovered))
		if err != nil {
			return "", ErrMarkdown
		}
		defer document.Close()
		root = document.Root()
		source = recovered
	}
	if err := adaptNodes(root, source); err != nil {
		return "", fmt.Errorf("%w: %w", ErrMarkdown, err)
	}
	if err := walkNodes(root.FirstChild(), filterNode); err != nil {
		return "", fmt.Errorf("Cannot filter raw HTML: %v", err)
	}
	return document.HTML(), nil
}

func walkNodes(node cmarkghmd.Node, visit func(cmarkghmd.Node) error) error {
	for ; node.Valid(); node = node.Next() {
		if err := visit(node); err != nil {
			return err
		}
		if err := walkNodes(node.FirstChild(), visit); err != nil {
			return err
		}
	}
	return nil
}

func filterNode(node cmarkghmd.Node) error {
	kind := node.Type()
	if kind != "html_inline" && kind != "html_block" {
		return nil
	}
	original := node.Literal()
	filtered, err := filterRawInputs(original)
	if err != nil {
		return err
	}
	if filtered != original && !node.SetLiteral(filtered) {
		return errors.New("cannot replace raw HTML")
	}
	return nil
}

func prepareHTMLStage(source, boundary string) (string, bool) {
	flattened := flattenTemplates(source, boundary)
	details := prepareDetails(flattened)
	reparse := flattened != details
	if reparse {
		details = removeBoundaries(details, boundary)
	}
	return details, reparse
}

func renderMathStage(source string, ids RenderIDs) (string, bool, error) {
	marked, err := prepareMathTables(source, ids.RunID, ids.Boundary, ids.TableMarker)
	if err != nil {
		return "", false, err
	}
	plain := removeBoundaries(marked, ids.Boundary)
	math := mathDOM(marked, ids.RunID, ids.Boundary)
	return math, plain != math, nil
}

func finishHTML(source string, ids RenderIDs) string {
	linked := processLinks(source, ids.Boundary)
	return FinalizeSerializedHTML(linked, ids.RunID, ids.TableMarker)
}
