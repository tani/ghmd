#include "../md4c/md4c-html.h"
#include <stdlib.h>
#include <string.h>

typedef struct {
  char *data;
  size_t length, capacity;
  int failed;
} ghmd_md_buffer;

static void ghmd_md_collect(const MD_CHAR *text, MD_SIZE size, void *userdata) {
  ghmd_md_buffer *buffer = (ghmd_md_buffer *)userdata;
  if (buffer->failed)
    return;
  if (buffer->length + (size_t)size + 1 > buffer->capacity) {
    size_t capacity = buffer->capacity ? buffer->capacity : 1024;
    while (capacity < buffer->length + (size_t)size + 1)
      capacity *= 2;
    char *data = (char *)realloc(buffer->data, capacity);
    if (!data) {
      buffer->failed = 1;
      return;
    }
    buffer->data = data;
    buffer->capacity = capacity;
  }
  memcpy(buffer->data + buffer->length, text, size);
  buffer->length += size;
  buffer->data[buffer->length] = 0;
}

char *bridge_md4c_render(const char *source, size_t length, int github_dialect) {
  ghmd_md_buffer buffer = {0};
  unsigned flags = github_dialect ? MD_DIALECT_GITHUB | MD_FLAG_LATEXMATHSPANS : 0;
  int result = md_html(source, (MD_SIZE)length, ghmd_md_collect, &buffer, flags, 0);
  if (result != 0 || buffer.failed) {
    free(buffer.data);
    return NULL;
  }
  if (!buffer.data) {
    buffer.data = (char *)calloc(1, 1);
  }
  return buffer.data;
}

void bridge_md4c_free(char *value) { free(value); }
