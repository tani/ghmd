/* Runtime-independent HTML implementation shared by the renderer and parser.
 * Include order is explicit; only native.c depends on the Bun/Node-API bridge. */
#include "../../vendor/lexbor/lexbor.h"
#include "buffer.c"
#include "parser.c"
#include "json.c"
#include "sanitize.c"
