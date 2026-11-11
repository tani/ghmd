/* Attribute edits on sanitized, DOM-serialized HTML. Unchanged attributes are
 * copied byte-for-byte; modified values are escaped exactly once. */
typedef struct {
  const unsigned char *name;
  size_t name_length, raw_start, raw_end, value_start, value_end;
  cmark_strbuf value;
  unsigned char removed, changed, loaded;
} gfm_attribute;
typedef struct { gfm_attribute *items; size_t count, capacity; } gfm_attributes;

static void gfm_attributes_free(gfm_attributes *attrs) {
  for (size_t i = 0; i < attrs->count; i++) cmark_strbuf_free(&attrs->items[i].value);
  free(attrs->items);
}
static int gfm_attributes_read(gfm_html_tree *tree, gfm_html_node *node, gfm_attributes *attrs) {
  const unsigned char *s = tree->source;
  size_t p = node->tag_end;
  while (p < node->open_end && s[p] != '>') {
    size_t space = p;
    while (p < node->open_end && gfm_ascii_space(s[p])) p++;
    size_t name = p;
    while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '=' && s[p] != '>' && s[p] != '/') p++;
    if (p == name) { p++; continue; }
    size_t name_end = p;
    while (p < node->open_end && gfm_ascii_space(s[p])) p++;
    size_t a = p, b = p;
    if (p < node->open_end && s[p] == '=') {
      p++; while (p < node->open_end && gfm_ascii_space(s[p])) p++;
      if (p < node->open_end && (s[p] == '"' || s[p] == '\'')) {
        unsigned char quote = s[p++]; a = p;
        while (p < node->open_end && s[p] != quote) p++;
        b = p; if (p < node->open_end) p++;
      } else { a = p; while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '>') p++; b = p; }
    }
    if (!gfm_grow((void **)&attrs->items, &attrs->capacity, attrs->count + 1, sizeof(*attrs->items))) return 0;
    gfm_attribute *attr = &attrs->items[attrs->count++];
    *attr = (gfm_attribute){ .name=s + name, .name_length=name_end-name, .raw_start=space, .raw_end=p, .value_start=a, .value_end=b };
    cmark_strbuf_init(cmark_get_default_mem_allocator(), &attr->value, 0);
  }
  return 1;
}
static gfm_attribute *gfm_attribute_find(gfm_attributes *attrs, const char *name) {
  for (size_t i = 0; i < attrs->count; i++)
    if (!attrs->items[i].removed && gfm_ascii_equal(attrs->items[i].name, attrs->items[i].name_length, name)) return &attrs->items[i];
  return NULL;
}
static cmark_strbuf *gfm_attribute_value(gfm_html_tree *tree, gfm_attribute *attr) {
  if (!attr->loaded) {
    houdini_unescape_html_f(&attr->value, tree->source + attr->value_start, attr->value_end - attr->value_start);
    attr->loaded = 1;
  }
  return &attr->value;
}
static void gfm_attribute_remove(gfm_attributes *attrs, const char *name) {
  gfm_attribute *attr = gfm_attribute_find(attrs, name);
  if (attr) attr->removed = 1;
}
static int gfm_attribute_set(gfm_attributes *attrs, const char *name, const unsigned char *value, size_t length) {
  gfm_attribute *attr = gfm_attribute_find(attrs, name);
  if (!attr) {
    if (!gfm_grow((void **)&attrs->items, &attrs->capacity, attrs->count + 1, sizeof(*attrs->items))) return 0;
    attr = &attrs->items[attrs->count++];
    *attr = (gfm_attribute){ .name=(const unsigned char *)name, .name_length=strlen(name) };
    cmark_strbuf_init(cmark_get_default_mem_allocator(), &attr->value, 0);
  }
  cmark_strbuf_set(&attr->value, value, length);
  attr->loaded = attr->changed = 1;
  return 1;
}
static void gfm_attribute_escape(cmark_strbuf *out, const unsigned char *value, size_t length) {
  for (size_t i = 0; i < length; i++) {
    if (value[i] == '&') cmark_strbuf_puts(out, "&amp;");
    else if (value[i] == '"') cmark_strbuf_puts(out, "&quot;");
    else cmark_strbuf_putc(out, value[i]);
  }
}
/* JavaScript String.slice counts UTF-16 code units, even within non-ASCII IDs. */
static void gfm_slice_utf16(cmark_strbuf *out, const unsigned char *text, size_t length, size_t units) {
  size_t i = 0;
  while (i < length && units) {
    uint32_t cp = gfm_utf8_at(text, length, i);
    size_t width = 1;
    while (i + width < length && (text[i + width] & 0xc0) == 0x80) width++;
    if (cp > 0xffff && units == 1) { cmark_strbuf_puts(out, "\xef\xbf\xbd"); i += width; units = 0; break; }
    units -= cp > 0xffff ? 2 : 1;
    i += width;
  }
  cmark_strbuf_put(out, text + i, length - i);
}
static int gfm_make_reference_value(gfm_html_tree *tree, gfm_attributes *attrs, const char *name,
                                    const char *prefix, size_t slice, const char *token,
                                    int colon_removes) {
  gfm_attribute *attr = gfm_attribute_find(attrs, name);
  if (!attr) return 1;
  cmark_strbuf *value = gfm_attribute_value(tree, attr);
  if (!value->size) return 1;
  cmark_strbuf sliced = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  gfm_slice_utf16(&sliced, value->ptr, value->size, slice);
  if (colon_removes && sliced.size == 1 && sliced.ptr[0] == ':') {
    gfm_attribute_remove(attrs, name); cmark_strbuf_free(&sliced); return 1;
  }
  cmark_strbuf output = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_strbuf_puts(&output, prefix);
  cmark_strbuf_put(&output, sliced.ptr, sliced.size);
  cmark_strbuf_putc(&output, '-'); cmark_strbuf_puts(&output, token);
  int ok = gfm_attribute_set(attrs, name, output.ptr, output.size);
  cmark_strbuf_free(&sliced); cmark_strbuf_free(&output);
  return ok;
}

static int gfm_ascii_prefix_case(const unsigned char *value, size_t size, const char *prefix) {
  size_t length = strlen(prefix);
  if (size < length) return 0;
  for (size_t i = 0; i < length; i++) {
    unsigned char c = value[i];
    if (c >= 'A' && c <= 'Z') c += 'a' - 'A';
    if (c != (unsigned char)prefix[i]) return 0;
  }
  return 1;
}
static int gfm_external_image(const unsigned char *value, size_t length) {
  return gfm_ascii_prefix_case(value, length, "http://") ||
         gfm_ascii_prefix_case(value, length, "https://") ||
         gfm_ascii_prefix_case(value, length, "//");
}
static int gfm_adjust_images(gfm_html_tree *tree) {
  int changed = 0;
  for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed || !gfm_node_is(tree, id, "img")) continue;
    size_t start, end;
    if (!gfm_get_attribute(tree, id, "src", &start, &end)) continue;
    size_t parent = tree->nodes[id].parent;
    if (parent == 0 || (parent != GFM_NONE && gfm_node_is(tree, parent, "b"))) {
      tree->nodes[id].image_paragraph = 1; changed = 1;
    }
    cmark_strbuf src = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    houdini_unescape_html_f(&src, tree->source + start, end - start);
    if (!gfm_external_image(src.ptr, src.size)) {
      tree->nodes[id].relative_image = 1;
      tree->nodes[id].image_rejected = gfm_ascii_prefix_case(src.ptr, src.size, "data:");
      tree->nodes[id].image_link = gfm_nearest(tree, id, "a", NULL) == GFM_NONE;
      changed = 1;
    }
    cmark_strbuf_free(&src);
  }
  return changed;
}
static int gfm_write_image(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node) {
  gfm_attributes attrs = {0};
  if (!gfm_attributes_read(tree, node, &attrs)) { gfm_attributes_free(&attrs); return 0; }
  gfm_attribute *src_attr = gfm_attribute_find(&attrs, "src");
  cmark_strbuf *src = gfm_attribute_value(tree, src_attr);
  cmark_strbuf image_href = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  if (!node->image_rejected) cmark_strbuf_put(&image_href, src->ptr, src->size);
  if (node->image_rejected) gfm_attribute_remove(&attrs, "src");
  int ok = gfm_attribute_set(&attrs, "style", (const unsigned char *)"max-width: 100%;", 16);
  if (ok && node->image_link) {
    cmark_strbuf_puts(out, "<a target=\"_blank\" rel=\"");
    if (gfm_ascii_prefix_case(image_href.ptr, image_href.size, "mailto:")) cmark_strbuf_puts(out, "nofollow");
    else cmark_strbuf_puts(out, "noopener noreferrer");
    cmark_strbuf_puts(out, "\" href=\"");
    gfm_attribute_escape(out, image_href.ptr, image_href.size); cmark_strbuf_puts(out, "\">");
  }
  if (ok) {
    cmark_strbuf_put(out, tree->source + node->start, node->tag_end - node->start);
    for (size_t i = 0; i < attrs.count; i++) {
      gfm_attribute *attr = &attrs.items[i];
      if (attr->removed) continue;
      if (!attr->changed) cmark_strbuf_put(out, tree->source + attr->raw_start, attr->raw_end - attr->raw_start);
      else {
        cmark_strbuf_putc(out, ' '); cmark_strbuf_put(out, attr->name, attr->name_length);
        cmark_strbuf_puts(out, "=\""); gfm_attribute_escape(out, attr->value.ptr, attr->value.size); cmark_strbuf_putc(out, '"');
      }
    }
    cmark_strbuf_putc(out, '>');
    if (node->image_link) cmark_strbuf_puts(out, "</a>");
  }
  cmark_strbuf_free(&image_href); gfm_attributes_free(&attrs);
  return ok;
}

/* Flags: 1 reference, 2 back-reference, 4 superscript, 8 section, 16 list item. */
static int gfm_write_footnote_tag(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node, const char *token, int *arrow) {
  gfm_attributes attrs = {0};
  if (!gfm_attributes_read(tree, node, &attrs)) { gfm_attributes_free(&attrs); return 0; }
  int ok = 1;
  if (node->footnote & 4) gfm_attribute_remove(&attrs, "class");
  if (node->footnote & 1) {
    ok = gfm_make_reference_value(tree, &attrs, "href", "#user-content-fn-", 4, token, 1) &&
         gfm_make_reference_value(tree, &attrs, "id", "user-content-", 0, token, 0) &&
         gfm_attribute_set(&attrs, "aria-describedby", (const unsigned char *)"footnote-label", 14);
  }
  if (ok && (node->footnote & 2)) {
    gfm_attribute *href = gfm_attribute_find(&attrs, "href");
    if (href) {
      cmark_strbuf *value = gfm_attribute_value(tree, href);
      cmark_strbuf sliced = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      gfm_slice_utf16(&sliced, value->ptr, value->size, 7);
      *arrow = sliced.size == 1 && sliced.ptr[0] == ':';
      cmark_strbuf_free(&sliced);
      if (!*arrow) ok = gfm_make_reference_value(tree, &attrs, "href", "#user-content-", 1, token, 0);
    }
    if (!*arrow) {
      gfm_attribute_remove(&attrs, "class");
      gfm_attribute_remove(&attrs, "data-footnote-backref-idx");
      ok = ok && gfm_attribute_set(&attrs, "class", (const unsigned char *)"data-footnote-backref", 21);
    }
  }
  if (ok && (node->footnote & 8)) {
    gfm_attribute_remove(&attrs, "class");
    gfm_attribute_remove(&attrs, "data-footnotes");
    ok = gfm_attribute_set(&attrs, "data-footnotes", (const unsigned char *)"", 0) &&
         gfm_attribute_set(&attrs, "class", (const unsigned char *)"footnotes", 9);
  }
  if (ok && (node->footnote & 16)) {
    /* li.id is rewritten even when the original id is empty. */
    gfm_attribute *attr = gfm_attribute_find(&attrs, "id");
    cmark_strbuf *value = gfm_attribute_value(tree, attr);
    cmark_strbuf updated = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    cmark_strbuf_puts(&updated, "user-content-"); cmark_strbuf_put(&updated, value->ptr, value->size);
    cmark_strbuf_putc(&updated, '-'); cmark_strbuf_puts(&updated, token);
    ok = gfm_attribute_set(&attrs, "id", updated.ptr, updated.size);
    cmark_strbuf_free(&updated);
  }
  if (ok && *arrow) cmark_strbuf_puts(out, "↩");
  else if (ok) {
    cmark_strbuf_put(out, tree->source + node->start, node->tag_end - node->start);
    int class_seen = 0;
    for (size_t i = 0; i < attrs.count; i++) {
      gfm_attribute *attr = &attrs.items[i];
      if (attr->removed) continue;
      if (node->normalize_class && gfm_ascii_equal(attr->name, attr->name_length, "class")) {
        if (attr->changed) {
          /* Class on a footnote section/backref is not normalized elsewhere. */
          cmark_strbuf_putc(out, ' '); cmark_strbuf_put(out, attr->name, attr->name_length);
          cmark_strbuf_puts(out, "=\""); gfm_attribute_escape(out, attr->value.ptr, attr->value.size); cmark_strbuf_putc(out, '"');
        } else gfm_write_class(out, tree->source + attr->value_start, attr->value_end - attr->value_start, node->normalize_class);
        class_seen = 1; continue;
      }
      if (!attr->changed) cmark_strbuf_put(out, tree->source + attr->raw_start, attr->raw_end - attr->raw_start);
      else {
        cmark_strbuf_putc(out, ' '); cmark_strbuf_put(out, attr->name, attr->name_length);
        cmark_strbuf_puts(out, "=\""); gfm_attribute_escape(out, attr->value.ptr, attr->value.size); cmark_strbuf_putc(out, '"');
      }
    }
    if (node->normalize_class > 1 && !class_seen) gfm_write_class(out, (const unsigned char *)"", 0, node->normalize_class);
    cmark_strbuf_putc(out, '>');
    if (node->footnote & 8) cmark_strbuf_puts(out, "<h2 id=\"footnote-label\" class=\"sr-only\">Footnotes</h2>");
  }
  gfm_attributes_free(&attrs);
  return ok;
}

static int gfm_adjust_footnote_nodes(gfm_html_tree *tree) {
  int changed = 0; size_t section = GFM_NONE;
  for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed) continue;
    if (gfm_node_is(tree, id, "sup") && gfm_class_contains(tree, id, "footnote-ref")) tree->nodes[id].footnote |= 4;
    if (gfm_node_is(tree, id, "a") && gfm_get_attribute(tree, id, "data-footnote-ref", NULL, NULL)) tree->nodes[id].footnote |= 1;
    if (gfm_node_is(tree, id, "a") && gfm_get_attribute(tree, id, "data-footnote-backref", NULL, NULL)) tree->nodes[id].footnote |= 2;
    if (section == GFM_NONE && gfm_node_is(tree, id, "section") && gfm_get_attribute(tree, id, "data-footnotes", NULL, NULL)) {
      section = id; tree->nodes[id].footnote |= 8;
      size_t next = tree->nodes[id].next;
      if (tree->nodes[id].parent != 0 && next != GFM_NONE && !tree->nodes[next].element &&
          tree->nodes[next].next == GFM_NONE && !gfm_text_has_visible_content(tree, next)) {
        gfm_detach_node(tree, next); tree->nodes[next].removed = 1;
      }
    }
    if (tree->nodes[id].footnote) changed = 1;
  }
  if (section != GFM_NONE) for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed || !gfm_node_is(tree, id, "li") || !gfm_get_attribute(tree, id, "id", NULL, NULL)) continue;
    for (size_t parent = tree->nodes[id].parent; parent != GFM_NONE; parent = tree->nodes[parent].parent)
      if (parent == section) { tree->nodes[id].footnote |= 16; break; }
  }
  return changed;
}
