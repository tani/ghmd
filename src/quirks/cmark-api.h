#ifndef GFM_API_H
#define GFM_API_H
#include <stddef.h>

typedef struct { const unsigned char *data; size_t length; } gfm_view;
typedef struct { unsigned char *data; size_t length; } gfm_string;
typedef enum { GFM_OK, GFM_BAD_ARGUMENT, GFM_NO_MEMORY, GFM_FILTER_FAILED } gfm_status;
typedef struct { gfm_string text; const char *error; } gfm_result;

/* Input views are borrowed. The filter returns a malloc-owned, NUL-terminated
 * string (or NULL for an empty string), which the core releases on every path.
 * Calls are synchronous and may be reentered; context belongs to this call. */
typedef gfm_status (*gfm_html_filter)(gfm_view input, gfm_string *output, void *context);

/* Successful output is malloc-owned and has an explicit UTF-8 byte length.
 * Errors contain a static diagnostic, never runtime-specific exception state.
 * The prepareStage and mathStage operations return an ASCII '0'/'1' prefix
 * followed by HTML. '1' means the caller must reparse through its DOM before
 * the next stage; '0' preserves the existing serialized structure as-is.
 * Views need not be NUL-terminated. A {NULL, 0} optional modifier uses the
 * operation's default; required arguments must have non-NULL data. */
gfm_status gfm_call(const char *operation, const gfm_view *arguments, size_t count,
                    gfm_html_filter filter, void *context, gfm_result *result);
#endif
