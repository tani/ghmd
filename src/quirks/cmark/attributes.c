/* Attribute edits on sanitized, DOM-serialized HTML. Unchanged attributes are
 * copied byte-for-byte; modified values are escaped exactly once. */
typedef struct {
  const unsigned char *name;
  size_t name_length, raw_start, raw_end, value_start, value_end;
  cmark_strbuf value;
  unsigned char removed, changed, loaded;
} gfm_attribute;
typedef struct {
  gfm_attribute *items;
  size_t count, capacity;
} gfm_attributes;

static void gfm_attributes_free(gfm_attributes *attrs) {
  for (size_t i = 0; i < attrs->count; i++)
    cmark_strbuf_free(&attrs->items[i].value);
  free(attrs->items);
}
static int gfm_attributes_read(gfm_html_tree *tree, gfm_html_node *node, gfm_attributes *attrs) {
  const unsigned char *s = tree->source;
  size_t p = node->tag_end;
  while (p < node->open_end && s[p] != '>') {
    size_t space = p;
    while (p < node->open_end && gfm_ascii_space(s[p]))
      p++;
    size_t name = p;
    while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '=' && s[p] != '>' &&
           s[p] != '/')
      p++;
    if (p == name) {
      p++;
      continue;
    }
    size_t name_end = p;
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
    if (!gfm_grow((void **)&attrs->items, &attrs->capacity, attrs->count + 1,
                  sizeof(*attrs->items)))
      return 0;
    gfm_attribute *attr = &attrs->items[attrs->count++];
    *attr = (gfm_attribute){.name = s + name,
                            .name_length = name_end - name,
                            .raw_start = space,
                            .raw_end = p,
                            .value_start = a,
                            .value_end = b};
    cmark_strbuf_init(cmark_get_default_mem_allocator(), &attr->value, 0);
  }
  return 1;
}
static gfm_attribute *gfm_attribute_find(gfm_attributes *attrs, const char *name) {
  for (size_t i = 0; i < attrs->count; i++)
    if (!attrs->items[i].removed &&
        gfm_ascii_equal(attrs->items[i].name, attrs->items[i].name_length, name))
      return &attrs->items[i];
  return NULL;
}
static cmark_strbuf *gfm_attribute_value(gfm_html_tree *tree, gfm_attribute *attr) {
  if (!attr->loaded) {
    houdini_unescape_html_f(&attr->value, tree->source + attr->value_start,
                            attr->value_end - attr->value_start);
    attr->loaded = 1;
  }
  return &attr->value;
}
static void gfm_attribute_remove(gfm_attributes *attrs, const char *name) {
  gfm_attribute *attr = gfm_attribute_find(attrs, name);
  if (attr)
    attr->removed = 1;
}
static int gfm_attribute_set(gfm_attributes *attrs, const char *name, const unsigned char *value,
                             size_t length) {
  gfm_attribute *attr = gfm_attribute_find(attrs, name);
  if (!attr) {
    if (!gfm_grow((void **)&attrs->items, &attrs->capacity, attrs->count + 1,
                  sizeof(*attrs->items)))
      return 0;
    attr = &attrs->items[attrs->count++];
    *attr = (gfm_attribute){.name = (const unsigned char *)name, .name_length = strlen(name)};
    cmark_strbuf_init(cmark_get_default_mem_allocator(), &attr->value, 0);
  }
  cmark_strbuf_set(&attr->value, value, length);
  attr->loaded = attr->changed = 1;
  return 1;
}
static void gfm_attribute_escape(cmark_strbuf *out, const unsigned char *value, size_t length) {
  for (size_t i = 0; i < length; i++) {
    if (value[i] == '&')
      cmark_strbuf_puts(out, "&amp;");
    else if (value[i] == '"')
      cmark_strbuf_puts(out, "&quot;");
    else
      cmark_strbuf_putc(out, value[i]);
  }
}
