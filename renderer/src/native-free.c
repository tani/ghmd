#include <stdlib.h>
/* Renderer strings use cmark's default malloc allocator. Free them in the
 * same shared library, avoiding assumptions about the host C runtime. */
void gfm_free(void *p) { free(p); }
