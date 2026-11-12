/* Locate GitHub-style inline math in a UTF-8 text node.  DOM construction is
 * performed by the caller after sanitization; returned ranges are byte offsets. */
typedef struct {
  size_t count, capacity;
  uint32_t *ranges;
} gfm_math_matches;

static uint32_t gfm_utf8_at(const unsigned char *s, size_t len, size_t i) {
  if (i >= len)
    return 0;
  unsigned char c = s[i];
  if (c < 0x80)
    return c;
  unsigned n = c >= 0xf0 && c <= 0xf4   ? 4
               : c >= 0xe0 && c <= 0xef ? 3
               : c >= 0xc2 && c <= 0xdf ? 2
                                        : 1;
  if (n == 1 || i + n > len)
    return 0xfffd;
  uint32_t cp = c & (0x7f >> n);
  for (unsigned j = 1; j < n; j++) {
    if ((s[i + j] & 0xc0) != 0x80)
      return 0xfffd;
    cp = (cp << 6) | (s[i + j] & 0x3f);
  }
  return cp;
}
static uint32_t gfm_utf8_before(const unsigned char *s, size_t i) {
  if (!i)
    return 0;
  size_t start = i - 1;
  while (start > 0 && i - start < 4 && (s[start] & 0xc0) == 0x80)
    start--;
  return gfm_utf8_at(s, i, start);
}
static int gfm_js_whitespace(uint32_t cp) {
  return (cp >= 9 && cp <= 13) || cp == 0x20 || cp == 0xa0 || cp == 0x1680 ||
         (cp >= 0x2000 && cp <= 0x200a) || cp == 0x2028 || cp == 0x2029 || cp == 0x202f ||
         cp == 0x205f || cp == 0x3000 || cp == 0xfeff;
}

gfm_math_matches *gfm_inline_math(const unsigned char *s, size_t length) {
  if (length > UINT32_MAX)
    return NULL;
  gfm_math_matches *result = calloc(1, sizeof(*result));
  if (!result)
    return NULL;
  for (size_t i = 0; i < length;) {
    if (s[i] != '$' || i + 1 >= length || s[i + 1] == '$' || s[i + 1] == '\n') {
      i++;
      continue;
    }
    size_t end = i + 1;
    while (end < length && s[end] != '$' && s[end] != '\n')
      end++;
    if (end == length || s[end] != '$' || end == i + 1) {
      i++;
      continue;
    }
    uint32_t before = gfm_utf8_before(s, i);
    unsigned char after = end + 1 < length ? s[end + 1] : 0;
    uint32_t first = gfm_utf8_at(s, length, i + 1), last = gfm_utf8_before(s, end);
    int invalid_after = (after >= 'a' && after <= 'z') || (after >= 'A' && after <= 'Z') ||
                        (after >= '0' && after <= '9') || after == '_';
    if ((i && !gfm_js_whitespace(before) && before != '(') || invalid_after ||
        gfm_js_whitespace(first) || (gfm_js_whitespace(last) && after != '$') ||
        memchr(s + i + 1, '&', end - i - 1)) {
      i++;
      continue;
    }
    if (result->count == result->capacity) {
      size_t capacity = result->capacity ? result->capacity * 2 : 8;
      if (capacity > SIZE_MAX / (2 * sizeof(uint32_t))) {
        free(result->ranges);
        free(result);
        return NULL;
      }
      uint32_t *p = realloc(result->ranges, capacity * 2 * sizeof(uint32_t));
      if (!p) {
        free(result->ranges);
        free(result);
        return NULL;
      }
      result->ranges = p;
      result->capacity = capacity;
    }
    result->ranges[result->count * 2] = (uint32_t)i;
    result->ranges[result->count * 2 + 1] = (uint32_t)(end + 1);
    result->count++;
    i = end + 1;
  }
  return result;
}
void gfm_inline_math_free(gfm_math_matches *m) {
  if (m) {
    free(m->ranges);
    free(m);
  }
}

/* Return consecutive display blocks only when they cover the entire text. */
gfm_math_matches *gfm_display_math(const unsigned char *s, size_t length) {
  if (length > UINT32_MAX)
    return NULL;
  gfm_math_matches *result = calloc(1, sizeof(*result));
  if (!result)
    return NULL;
  size_t start = 0;
  while (start < length) {
    if (length - start < 6 || memcmp(s + start, "$$\n", 3))
      goto mismatch;
    size_t end = start + 3;
    while (end + 2 < length && !(s[end] == '\n' && s[end + 1] == '$' && s[end + 2] == '$'))
      end++;
    if (end + 2 >= length)
      goto mismatch;
    end += 3;
    if (result->count == result->capacity) {
      size_t capacity = result->capacity ? result->capacity * 2 : 4;
      if (capacity > SIZE_MAX / (2 * sizeof(uint32_t))) {
        gfm_inline_math_free(result);
        return NULL;
      }
      uint32_t *p = realloc(result->ranges, capacity * 2 * sizeof(uint32_t));
      if (!p) {
        gfm_inline_math_free(result);
        return NULL;
      }
      result->ranges = p;
      result->capacity = capacity;
    }
    result->ranges[result->count * 2] = (uint32_t)start;
    result->ranges[result->count * 2 + 1] = (uint32_t)end;
    result->count++;
    if (end == length)
      return result;
    if (s[end] != '\n')
      goto mismatch;
    start = end + 1;
    if (start == length)
      goto mismatch;
  }
  return result;
mismatch:
  result->count = 0;
  return result;
}

/* Preserve the renderer's extra escaping of backslash-newline sequences. */
static void gfm_math_put_lines(cmark_strbuf *out, const unsigned char *s, size_t length) {
  for (size_t i = 0; i < length; i++) {
    if (s[i] == '\\' && i + 1 < length && s[i + 1] == '\n')
      cmark_strbuf_putc(out, '\\');
    cmark_strbuf_putc(out, s[i]);
  }
}
/* Display math containing Markdown emphasis is reconstructed with blank lines. */
char *gfm_emphasis_display(const unsigned char *s, size_t length) {
  if (length < 6 || memcmp(s, "$$\n", 3) || memcmp(s + length - 3, "\n$$", 3))
    return NULL;
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_strbuf_puts(&out, "$$\n\n");
  cmark_strbuf_put(&out, s + 3, length - 6);
  cmark_strbuf_puts(&out, "\n\n$$");
  return (char *)cmark_strbuf_detach(&out);
}

char *gfm_fenced_math(const unsigned char *s, size_t length) {
  if (length && s[length - 1] == '\n')
    length--;
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_strbuf_puts(&out, "$$");
  for (size_t i = 0; i < length; i++) {
    if (s[i] == '$' && (!i || s[i - 1] != '\\'))
      cmark_strbuf_putc(&out, '\\');
    if (s[i] == '\\' && i + 1 < length && s[i + 1] == '\n')
      cmark_strbuf_putc(&out, '\\');
    cmark_strbuf_putc(&out, s[i]);
  }
  cmark_strbuf_puts(&out, "$$");
  return (char *)cmark_strbuf_detach(&out);
}

/* Return NULL when the code element is not a protected math expression. */
char *gfm_protected_math(const unsigned char *before, size_t before_length,
                         const unsigned char *after, size_t after_length, const unsigned char *code,
                         size_t code_length, int emphasis) {
  if (!before_length || !after_length || before[before_length - 1] != '$' || after[0] != '$')
    return NULL;
  if (before_length > 1 && (before[before_length - 2] == 'a' || before[before_length - 2] == 'b'))
    return NULL;
  if (emphasis && after_length == 1)
    return NULL;
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_strbuf_putc(&out, '$');
  for (size_t i = 0; i < code_length; i++) {
    if (code[i] == '&')
      cmark_strbuf_puts(&out, "&amp;");
    else if (code[i] == '<')
      cmark_strbuf_puts(&out, "&lt;");
    else if (code[i] == '>')
      cmark_strbuf_puts(&out, "&gt;");
    else
      cmark_strbuf_putc(&out, code[i]);
  }
  cmark_strbuf_putc(&out, '$');
  return (char *)cmark_strbuf_detach(&out);
}
