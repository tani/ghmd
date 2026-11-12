/* Apply DOMTokenList normalization while preserving the attribute order. */
static void gfm_write_class(cmark_strbuf *out, const unsigned char *value, size_t length,
                            int action) {
  const char *added = action == 2   ? "notranslate"
                      : action == 3 ? "task-list-item"
                      : action == 4 ? "contains-task-list"
                      : action == 5 ? "task-list-item-checkbox"
                                    : NULL;
  cmark_strbuf decoded = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_strbuf tokens = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&decoded, value, length);
  size_t n = decoded.size;
  for (size_t i = 0; i < n;) {
    while (i < n && gfm_ascii_space(decoded.ptr[i]))
      i++;
    size_t start = i;
    while (i < n && !gfm_ascii_space(decoded.ptr[i]))
      i++;
    size_t count = i - start;
    if (!count ||
        (action == 1 && count == 14 && !memcmp(decoded.ptr + start, "task-list-item", 14)))
      continue;
    int duplicate = 0;
    for (size_t j = 0; j < (size_t)tokens.size;) {
      size_t previous = j;
      while (j < (size_t)tokens.size && tokens.ptr[j] != ' ')
        j++;
      if (j - previous == count && !memcmp(tokens.ptr + previous, decoded.ptr + start, count)) {
        duplicate = 1;
        break;
      }
      j++;
    }
    if (!duplicate) {
      if (tokens.size)
        cmark_strbuf_putc(&tokens, ' ');
      cmark_strbuf_put(&tokens, decoded.ptr + start, count);
    }
  }
  if (added) {
    size_t added_length = strlen(added);
    int found = 0;
    for (size_t j = 0; j < (size_t)tokens.size;) {
      size_t previous = j;
      while (j < (size_t)tokens.size && tokens.ptr[j] != ' ')
        j++;
      if (j - previous == added_length && !memcmp(tokens.ptr + previous, added, added_length))
        found = 1;
      j++;
    }
    if (!found) {
      if (tokens.size)
        cmark_strbuf_putc(&tokens, ' ');
      cmark_strbuf_puts(&tokens, added);
    }
  }
  if (tokens.size) {
    cmark_strbuf_puts(out, " class=\"");
    for (size_t i = 0; i < (size_t)tokens.size; i++) {
      unsigned char ch = tokens.ptr[i];
      if (ch == '&')
        cmark_strbuf_puts(out, "&amp;");
      else if (ch == '"')
        cmark_strbuf_puts(out, "&quot;");
      else
        cmark_strbuf_putc(out, ch);
    }
    cmark_strbuf_putc(out, '"');
  }
  cmark_strbuf_free(&decoded);
  cmark_strbuf_free(&tokens);
}

static void gfm_write_normalized_tag(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node) {
  const unsigned char *s = tree->source;
  size_t p = node->tag_end, written = node->start;
  int saw_class = 0, action = node->normalize_class;
  while (p < node->open_end && s[p] != '>') {
    size_t space = p;
    while (p < node->open_end && gfm_ascii_space(s[p]))
      p++;
    size_t attr = p;
    while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '=' && s[p] != '>' &&
           s[p] != '/')
      p++;
    if (p == attr) {
      p++;
      continue;
    }
    size_t attr_end = p;
    while (p < node->open_end && gfm_ascii_space(s[p]))
      p++;
    size_t a = p, b = p;
    if (p < node->open_end && s[p] == '=') {
      p++;
      while (p < node->open_end && gfm_ascii_space(s[p]))
        p++;
      if (p < node->open_end && (s[p] == '"' || s[p] == '\'')) {
        unsigned char quote = s[p++];
        a = p;
        while (p < node->open_end && s[p] != quote)
          p++;
        b = p;
        if (p < node->open_end)
          p++;
      } else {
        a = p;
        while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '>')
          p++;
        b = p;
      }
    }
    if (gfm_ascii_equal(s + attr, attr_end - attr, "class")) {
      cmark_strbuf_put(out, s + written, space - written);
      gfm_write_class(out, s + a, b - a, action);
      written = p;
      saw_class = 1;
    } else if (action == 2 && gfm_node_is(tree, (size_t)(node - tree->nodes), "pre") &&
               gfm_ascii_equal(s + attr, attr_end - attr, "lang")) {
      cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      houdini_unescape_html_f(&value, s + a, b - a);
      int math = value.size == 4 && !memcmp(value.ptr, "math", 4);
      cmark_strbuf_free(&value);
      if (math) {
        cmark_strbuf_put(out, s + written, space - written);
        written = p;
      }
    }
  }
  if (action != 1 && !saw_class) {
    cmark_strbuf_put(out, s + written, node->open_end - 1 - written);
    gfm_write_class(out, (const unsigned char *)"", 0, action);
    cmark_strbuf_putc(out, '>');
  } else
    cmark_strbuf_put(out, s + written, node->open_end - written);
}
