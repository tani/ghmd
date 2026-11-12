#include "../../third_party/md4c/md4c-html.c"
#include "../../third_party/md4c/entity.c"
#include "bridge.h"
#include <stdlib.h>
#include <string.h>

typedef struct {
  char *data;
  size_t length, capacity;
  int failed;
} md_buffer;

static void md_collect(const MD_CHAR *text, MD_SIZE size, void *userdata) {
  md_buffer *buffer = userdata;
  if (buffer->failed)
    return;
  if (buffer->length + size + 1 > buffer->capacity) {
    size_t capacity = buffer->capacity ? buffer->capacity : 1024;
    while (capacity < buffer->length + size + 1)
      capacity *= 2;
    char *data = realloc(buffer->data, capacity);
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
  md_buffer buffer = {0};
  unsigned flags = github_dialect ? MD_DIALECT_GITHUB : 0;
  int result = md_html(source, length, md_collect, &buffer, flags, 0);
  if (result != 0 || buffer.failed) {
    free(buffer.data);
    return NULL;
  }
  if (!buffer.data) {
    buffer.data = calloc(1, 1);
    if (!buffer.data)
      return NULL;
  }
  return buffer.data;
}
