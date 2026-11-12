package github

import "errors"

// Errors returned by the rendering pipeline. Callers should use errors.Is
// rather than matching message text.
var (
	ErrDepth      = errors.New("HTML tree exceeds depth limit")
	ErrSize       = errors.New("HTML output exceeds byte limit")
	ErrParse      = errors.New("cannot parse HTML")
	ErrSourceSize = errors.New("Markdown exceeds 1 MB")
	ErrOutputSize = errors.New("Rendered HTML exceeds 32 MB")
	ErrRunID      = errors.New("runId must be 32 lowercase hexadecimal characters")
	ErrBoundary   = errors.New("Invalid render boundary")
	ErrMarkdown   = errors.New("Cannot parse Markdown")
	ErrDocument   = errors.New("Cannot render HTML document")
)
