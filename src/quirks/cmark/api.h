#ifndef GFM_API_H
#define GFM_API_H
#include <stddef.h>

typedef struct { const unsigned char *data; size_t length; } gfm_view;
typedef struct { unsigned char *data; size_t length; } gfm_string;
typedef enum { GFM_OK, GFM_BAD_ARGUMENT, GFM_NO_MEMORY, GFM_LIMIT_EXCEEDED } gfm_status;
typedef struct { gfm_string text; const char *error; } gfm_result;
#define GFM_MAX_OUTPUT_BYTES ((size_t)32 * 1024 * 1024)

/* Render Markdown to final sanitized HTML. Inputs are borrowed UTF-8 views;
 * the run ID is 32 lowercase hex characters and the boundary is UUID-shaped.
 * Views need not be NUL-terminated. Data pointers and result must be non-NULL.
 * Successful output is malloc-owned with an explicit byte length. The caller
 * frees result->text.data. Failure returns no owned output and a static error.
 * There are no callbacks, operations, or intermediate HTML entry points.
 */
gfm_status gfm_render(gfm_view source, gfm_view run_id, gfm_view boundary, gfm_result *result);
#endif
