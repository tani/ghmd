// Package cmarkghmd exposes the cmark-gfm parser and AST transpiled by
// modernc.org/ccgo/v4, without using cgo.
package cmarkghmd

import (
	"errors"
	"unsafe"

	"modernc.org/libc"
)

// Node is a borrowed cmark AST node. It is valid while its Document is open.
type Node struct{ p uintptr }

// Document owns one parsed cmark document and its parser.
type Document struct{ p uintptr }

// Parse parses Markdown with the GitHub extensions enabled.
func Parse(source string) (*Document, error) {
	tls := libc.NewTLS()
	defer tls.Close()
	input := libc.Xmalloc(tls, libc.Tsize_t(len(source)+1))
	defer libc.Xfree(tls, input)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(input)), len(source)), source)
	doc := bridge_parse(tls, input, libc.Tsize_t(len(source)))
	if doc == 0 {
		return nil, errors.New("cannot parse Markdown")
	}
	return &Document{p: doc}, nil
}

// Close releases the parser and AST.
func (doc *Document) Close() {
	tls := libc.NewTLS()
	defer tls.Close()
	bridge_free(tls, doc.p)
	doc.p = 0
}

// Root returns the document node.
func (doc *Document) Root() Node {
	tls := libc.NewTLS()
	defer tls.Close()
	return Node{p: bridge_root(tls, doc.p)}
}

// HTML renders the AST to HTML using cmark-gfm's renderer.
func (doc *Document) HTML() string {
	tls := libc.NewTLS()
	defer tls.Close()
	html := bridge_render_html(tls, doc.p)
	if html == 0 {
		return ""
	}
	defer bridge_free_string(tls, html)
	return libc.GoString(html)
}

// Version returns the vendored cmark-gfm release string.
func Version() string {
	tls := libc.NewTLS()
	defer tls.Close()
	return libc.GoString(bridge_version(tls))
}

// Type returns the cmark node type name, such as "paragraph" or "text".
func (n Node) Type() string {
	tls := libc.NewTLS()
	defer tls.Close()
	return libc.GoString(bridge_node_type(tls, n.p))
}

// FirstChild returns the first child, or a zero node.
func (n Node) FirstChild() Node {
	tls := libc.NewTLS()
	defer tls.Close()
	return Node{p: bridge_first_child(tls, n.p)}
}

// Next returns the next sibling, or a zero node.
func (n Node) Next() Node {
	tls := libc.NewTLS()
	defer tls.Close()
	return Node{p: bridge_next(tls, n.p)}
}

// Parent returns the parent, or a zero node.
func (n Node) Parent() Node {
	tls := libc.NewTLS()
	defer tls.Close()
	return Node{p: bridge_parent(tls, n.p)}
}

// Valid reports whether the node refers to an AST node.
func (n Node) Valid() bool { return n.p != 0 }

// Literal returns the node literal, or "".
func (n Node) Literal() string {
	tls := libc.NewTLS()
	defer tls.Close()
	literal := bridge_node_literal(tls, n.p)
	if literal == 0 {
		return ""
	}
	return libc.GoString(literal)
}

// SetLiteral replaces a node literal.
func (n Node) SetLiteral(value string) bool {
	tls := libc.NewTLS()
	defer tls.Close()
	data, err := libc.CString(value)
	if err != nil {
		return false
	}
	defer libc.Xfree(tls, data)
	return bridge_set_literal(tls, n.p, data, libc.Tsize_t(len(value))) != 0
}

// StartLine returns the one-based source line, or 0.
func (n Node) StartLine() int {
	tls := libc.NewTLS()
	defer tls.Close()
	return int(bridge_start_line(tls, n.p))
}

// StartColumn returns the one-based source byte column, or 0.
func (n Node) StartColumn() int {
	tls := libc.NewTLS()
	defer tls.Close()
	return int(bridge_start_column(tls, n.p))
}

// EndLine returns the one-based ending source line, or 0.
func (n Node) EndLine() int {
	tls := libc.NewTLS()
	defer tls.Close()
	return int(bridge_end_line(tls, n.p))
}

// EndColumn returns the one-based ending source byte column, or 0.
func (n Node) EndColumn() int {
	tls := libc.NewTLS()
	defer tls.Close()
	return int(bridge_end_column(tls, n.p))
}

// SetExtension attaches a named syntax extension to the node.
func (n Node) SetExtension(name string) bool {
	tls := libc.NewTLS()
	defer tls.Close()
	namePtr, err := libc.CString(name)
	if err != nil {
		return false
	}
	defer libc.Xfree(tls, namePtr)
	return bridge_set_extension(tls, n.p, namePtr) != 0
}

// SetTaskChecked marks a task-list item.
func (n Node) SetTaskChecked(checked bool) bool {
	tls := libc.NewTLS()
	defer tls.Close()
	value := int32(0)
	if checked {
		value = 1
	}
	return bridge_set_task_checked(tls, n.p, value) != 0
}

// TaskChecked reports a task-list item's checked state.
func (n Node) TaskChecked() bool {
	tls := libc.NewTLS()
	defer tls.Close()
	return bridge_get_task_checked(tls, n.p) != 0
}
