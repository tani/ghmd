#ifndef GFM_API_H
#define GFM_API_H
#include <stddef.h>

typedef struct { const unsigned char *data; size_t length; } gfm_view;
typedef struct { unsigned char *data; size_t length; } gfm_string;
typedef enum { GFM_OK, GFM_BAD_ARGUMENT, GFM_NO_MEMORY, GFM_FILTER_FAILED, GFM_LIMIT_EXCEEDED } gfm_status;
typedef struct { gfm_string text; const char *error; } gfm_result;
#define GFM_MAX_OUTPUT_BYTES ((size_t)32 * 1024 * 1024)

/* Input views are borrowed. The filter returns a malloc-owned, NUL-terminated
 * string (or NULL for an empty string), which the core releases on every path.
 * Calls are synchronous and may be reentered; context belongs to this call. */
typedef gfm_status (*gfm_html_filter)(gfm_view input, gfm_string *output, void *context);

/* Successful output is malloc-owned and has an explicit UTF-8 byte length.
 * Errors contain a static diagnostic, never runtime-specific exception state.
 * "document" accepts Markdown, a 32-hex run ID, and a UUID-shaped boundary.
 * It owns raw-input filtering, HTML5 parsing, sanitization, compatibility edits,
 * every reparse, math, links, and final size checks. Only final HTML is returned.
 * "render" (unsanitized Markdown HTML) and "ast" accept one source and an
 * optional diagnostic raw-HTML filter; the default filter is implemented in C.
 * "version", "lowercase", and "rawInputs" are string-only diagnostics.
 * Views need not be NUL-terminated. All arguments must have non-NULL data.
 * Intermediate prepare/math/finish operations are private C helpers.
 */
gfm_status gfm_call(const char *operation, const gfm_view *arguments, size_t count,
                    gfm_html_filter filter, void *context, gfm_result *result);
#endif
