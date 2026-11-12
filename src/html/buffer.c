#define H_ALLOC_FAILED 2
#define H_DEPTH_EXCEEDED 3
#define H_LIMIT ((size_t)64 * 1024 * 1024)
typedef struct {
  unsigned char *data;
  size_t length, capacity;
  int failed;
} h_buffer;
static void h_put(h_buffer *b, const void *data, size_t length) {
  if (b->failed || length > H_LIMIT - b->length) {
    b->failed = 1;
    return;
  }
  size_t need = b->length + length + 1;
  if (need > b->capacity) {
    size_t capacity = b->capacity ? b->capacity : 256;
    while (capacity < need)
      capacity *= 2;
    unsigned char *next = realloc(b->data, capacity);
    if (!next) {
      b->failed = H_ALLOC_FAILED;
      return;
    }
    b->data = next;
    b->capacity = capacity;
  }
  if (length)
    memcpy(b->data + b->length, data, length);
  b->length += length;
  b->data[b->length] = 0;
}
static void h_s(h_buffer *b, const char *s) {
  h_put(b, s, strlen(s));
}
