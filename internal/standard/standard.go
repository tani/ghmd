// Package standard exposes the vendored md4c CommonMark parser through cgo.
package standard

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/../../third_party/md4c
#include <stdlib.h>
#include "bridge.h"
*/
import "C"

import (
	"errors"
)

// Render converts Markdown to HTML with md4c. github enables the GitHub dialect.
func Render(source string, github bool) (string, error) {
	data := C.CBytes([]byte(source))
	defer C.free(data)
	flag := C.int(0)
	if github {
		flag = 1
	}
	out := C.bridge_md4c_render((*C.char)(data), C.size_t(len(source)), flag)
	if out == nil {
		return "", errors.New("cannot render Markdown with md4c")
	}
	defer C.bridge_md4c_free(out)
	return C.GoString(out), nil
}
