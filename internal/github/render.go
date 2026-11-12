package github

import (
	"errors"
	"strings"

	"yagfm/internal/cmark-gfm"
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
		return errors.New("Markdown exceeds 1 MB")
	}
	if len(runID) != runIDLength {
		return errors.New("runId must be 32 lowercase hexadecimal characters")
	}
	for i := 0; i < len(runID); i++ {
		if !isHex(runID[i]) {
			return errors.New("runId must be 32 lowercase hexadecimal characters")
		}
	}
	if len(boundary) != boundaryLength {
		return errors.New("Invalid render boundary")
	}
	for i := 0; i < len(boundary); i++ {
		dash := i == 8 || i == 13 || i == 18 || i == 23
		if dash {
			if boundary[i] != '-' {
				return errors.New("Invalid render boundary")
			}
		} else if !isHex(boundary[i]) {
			return errors.New("Invalid render boundary")
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
		return "", errors.New("Rendered HTML exceeds 32 MB")
	}
	html, ok := domSanitize(html, ids.Boundary)
	if !ok {
		return "", htmlStageError(html, "Cannot render HTML document")
	}
	html, ok = domReparse(html)
	if !ok {
		return "", htmlStageError(html, "Cannot render HTML document")
	}
	html, reparse, ok := prepareHTMLStage(html, ids.Boundary)
	if !ok {
		return "", htmlStageError(html, "Cannot render HTML document")
	}
	if reparse {
		if html, ok = domReparse(html); !ok {
			return "", htmlStageError(html, "Cannot render HTML document")
		}
	}
	html, reparse, ok = renderMathStage(html, ids)
	if !ok {
		return "", htmlStageError(html, "Cannot render HTML document")
	}
	if reparse {
		if html, ok = domReparse(html); !ok {
			return "", htmlStageError(html, "Cannot render HTML document")
		}
	}
	html, ok = finishHTML(html, ids)
	if !ok {
		return "", errors.New("Cannot render HTML document")
	}
	if len(html) > 0 && html[len(html)-1] == '\n' {
		html = html[:len(html)-1]
	}
	if len(html) > maxOutputBytes {
		return "", errors.New("Rendered HTML exceeds 32 MB")
	}
	return html, nil
}

func htmlStageError(_ string, message string) error { return errors.New(message) }

func renderMarkdown(source []byte) (string, error) {
	document, err := cmarkgfm.Parse(string(source))
	if err != nil {
		return "", errors.New("Cannot parse Markdown")
	}
	defer document.Close()
	root := document.Root()
	recovered, changed, ok := recoverLabels(root, source)
	if !ok {
		return "", errors.New("Cannot parse Markdown")
	}
	if changed {
		document.Close()
		document, err = cmarkgfm.Parse(string(recovered))
		if err != nil {
			return "", errors.New("Cannot parse Markdown")
		}
		defer document.Close()
		root = document.Root()
		source = recovered
	}
	if !adaptNodes(root, source) {
		return "", errors.New("Cannot parse Markdown")
	}
	if !walkNodes(root.FirstChild(), filterNode) {
		return "", errors.New("Cannot filter raw HTML")
	}
	return document.HTML(), nil
}

func walkNodes(node cmarkgfm.Node, visit func(cmarkgfm.Node) bool) bool {
	for ; node.Valid(); node = node.Next() {
		if !visit(node) {
			return false
		}
		if !walkNodes(node.FirstChild(), visit) {
			return false
		}
	}
	return true
}

func filterNode(node cmarkgfm.Node) bool {
	kind := node.Type()
	if kind != "html_inline" && kind != "html_block" {
		return true
	}
	original := node.Literal()
	filtered, ok := filterRawInputs(original)
	if !ok {
		return false
	}
	if filtered != original && !node.SetLiteral(filtered) {
		return false
	}
	return true
}

func prepareHTMLStage(source, boundary string) (string, bool, bool) {
	flattened, ok := flattenTemplates(source, boundary)
	if !ok {
		return "", false, false
	}
	details, ok := prepareDetails(flattened)
	if !ok {
		return "", false, false
	}
	reparse := flattened != details
	if reparse {
		details = removeBoundaries(details, boundary)
	}
	return details, reparse, true
}

func renderMathStage(source string, ids RenderIDs) (string, bool, bool) {
	marked, ok := prepareMathTables(source, ids.RunID, ids.Boundary, ids.TableMarker)
	if !ok {
		return "", false, false
	}
	plain := removeBoundaries(marked, ids.Boundary)
	math, ok := mathDOM(marked, ids.RunID, ids.Boundary)
	if !ok {
		return "", false, false
	}
	return math, plain != math, true
}

func finishHTML(source string, ids RenderIDs) (string, bool) {
	linked, ok := processLinks(source, ids.Boundary)
	if !ok {
		return "", false
	}
	return FinalizeSerializedHTML(linked, ids.RunID, ids.TableMarker)
}
