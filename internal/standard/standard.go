// Package standard exposes the vendored md4c parser, transpiled to Go by
// modernc.org/ccgo/v4, without using cgo.
package standard

import (
	"errors"
	"unsafe"

	"modernc.org/libc"
)

const githubDialectFlags = MD_FLAG_PERMISSIVEATXHEADERS | MD_FLAG_PERMISSIVEURLAUTOLINKS |
	MD_FLAG_PERMISSIVEWWWAUTOLINKS | MD_FLAG_PERMISSIVEEMAILAUTOLINKS |
	MD_FLAG_TABLES | MD_FLAG_STRIKETHROUGH | MD_FLAG_TASKLISTS | MD_FLAG_LATEXMATHSPANS

// Render converts Markdown to HTML with md4c. github enables the GitHub
// extensions and LaTeX math spans.
func Render(source string, github bool) (string, error) {
	tls := libc.NewTLS()
	defer tls.Close()
	input := libc.Xmalloc(tls, libc.Tsize_t(len(source)+1))
	defer libc.Xfree(tls, input)
	bytes := unsafe.Slice((*byte)(unsafe.Pointer(input)), len(source))
	copy(bytes, source)
	dialect := int32(0)
	if github {
		dialect = 1
	}
	html := bridge_md4c_render(tls, input, libc.Tsize_t(len(source)), dialect)
	if html == 0 {
		return "", errors.New("cannot render Markdown with md4c")
	}
	defer bridge_md4c_free(tls, html)
	return libc.GoString(html), nil
}
