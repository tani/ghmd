/* Operations on HTML emitted by the already-sanitized DOM serializer. */
static int gfm_ascii_space(unsigned char c) {
  return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f';
}
static int gfm_ascii_equal(const unsigned char *s, size_t n, const char *expected) {
  size_t length = strlen(expected);
  if (n != length)
    return 0;
  for (size_t i = 0; i < n; i++) {
    unsigned char c = s[i];
    if (c >= 'A' && c <= 'Z')
      c += 'a' - 'A';
    if (c != (unsigned char)expected[i])
      return 0;
  }
  return 1;
}
static size_t gfm_tag_end(const unsigned char *s, size_t length, size_t start) {
  unsigned char quote = 0;
  for (size_t i = start; i < length; i++) {
    if (quote) {
      if (s[i] == quote)
        quote = 0;
    } else if (s[i] == '\'' || s[i] == '"')
      quote = s[i];
    else if (s[i] == '>')
      return i + 1;
  }
  return length;
}

static int gfm_ascii_prefix_case(const unsigned char *value, size_t size, const char *prefix) {
  size_t length = strlen(prefix);
  if (size < length)
    return 0;
  for (size_t i = 0; i < length; i++) {
    unsigned char c = value[i];
    if (c >= 'A' && c <= 'Z')
      c += 'a' - 'A';
    if (c != (unsigned char)prefix[i])
      return 0;
  }
  return 1;
}
