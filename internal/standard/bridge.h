#ifndef YAGFM_MD4C_BRIDGE_H
#define YAGFM_MD4C_BRIDGE_H
#include <stddef.h>
char *bridge_md4c_render(const char *source, size_t length, int github_dialect);
void bridge_md4c_free(char *value);
#endif
