// Package cmarkgfm exposes the vendored cmark-gfm parser through cgo.
package cmarkgfm

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/../../third_party/cmark-gfm/src -I${SRCDIR}/../../third_party/cmark-gfm/extensions -I${SRCDIR}/../../third_party/cmark-gfm/generated
#include <stdlib.h>
#include "bridge.h"
*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

// Node is a borrowed cmark AST node. It is valid while its Document is open.
type Node struct {
	p unsafe.Pointer
}

// Document owns one parsed cmark document and its parser.
type Document struct {
	p *C.bridge_document
}

// Version returns the vendored cmark-gfm release string.
func Version() string { return C.GoString(C.bridge_version()) }

// Parse parses Markdown with the GitHub extensions enabled.
func Parse(source string) (*Document, error) {
	C.bridge_register_extensions()
	data := C.CBytes([]byte(source))
	defer C.free(data)
	p := C.bridge_parse((*C.char)(data), C.size_t(len(source)))
	if p == nil {
		return nil, errors.New("cannot parse Markdown")
	}
	d := &Document{p: p}
	runtime.SetFinalizer(d, (*Document).Close)
	return d, nil
}

// Close releases the parser and AST.
func (d *Document) Close() {
	if d.p != nil {
		C.bridge_free(d.p)
		d.p = nil
	}
}

// Root returns the document node.
func (d *Document) Root() Node { return Node{p: C.bridge_root(d.p)} }

// HTML renders the AST to HTML using cmark-gfm's renderer.
func (d *Document) HTML() string {
	p := C.bridge_render_html(d.p)
	if p == nil {
		return ""
	}
	defer C.bridge_free_string(p)
	return C.GoString(p)
}

// Type returns the cmark node type name, such as "paragraph" or "text".
func (n Node) Type() string { return C.GoString(C.bridge_node_type(n.p)) }

// FirstChild returns the first child, or a zero node.
func (n Node) FirstChild() Node { return Node{p: C.bridge_first_child(n.p)} }

// Next returns the next sibling, or a zero node.
func (n Node) Next() Node { return Node{p: C.bridge_next(n.p)} }

// Parent returns the parent, or a zero node.
func (n Node) Parent() Node { return Node{p: C.bridge_parent(n.p)} }

// Valid reports whether the node refers to an AST node.
func (n Node) Valid() bool { return n.p != nil }

// Literal returns the node literal, or "".
func (n Node) Literal() string {
	p := C.bridge_node_literal(n.p)
	if p == nil {
		return ""
	}
	return C.GoString(p)
}

// SetLiteral replaces a node literal.
func (n Node) SetLiteral(value string) bool {
	c := C.CString(value)
	defer C.free(unsafe.Pointer(c))
	return C.bridge_set_literal(n.p, c, C.size_t(len(value))) != 0
}

// StartLine returns the one-based source line, or 0.
func (n Node) StartLine() int { return int(C.bridge_start_line(n.p)) }

// StartColumn returns the one-based source byte column, or 0.
func (n Node) StartColumn() int { return int(C.bridge_start_column(n.p)) }

// EndLine returns the one-based ending source line, or 0.
func (n Node) EndLine() int { return int(C.bridge_end_line(n.p)) }

// EndColumn returns the one-based ending source byte column, or 0.
func (n Node) EndColumn() int { return int(C.bridge_end_column(n.p)) }

// SetExtension attaches a named syntax extension to the node.
func (n Node) SetExtension(name string) bool {
	c := C.CString(name)
	defer C.free(unsafe.Pointer(c))
	return C.bridge_set_extension(n.p, c) != 0
}

// SetTaskChecked marks a task-list item.
func (n Node) SetTaskChecked(checked bool) bool {
	v := C.int(0)
	if checked {
		v = 1
	}
	return C.bridge_set_task_checked(n.p, v) != 0
}

// TaskChecked reports a task-list item's checked state.
func (n Node) TaskChecked() bool { return C.bridge_get_task_checked(n.p) != 0 }
