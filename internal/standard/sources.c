#include "../../third_party/md4c/md4c.c"
#include "bridge.h"
#include <stdlib.h>

void bridge_md4c_free(char *value) { free(value); }
