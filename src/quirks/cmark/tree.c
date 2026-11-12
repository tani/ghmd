/* Byte-range tree and queries for already serialized, sanitized HTML.
 * Nodes retain source spans while feature modules edit relationships and flags.
 * Arbitrary raw HTML is parsed by Lexbor before reaching this representation. */
typedef struct {
  size_t start, open_end, close_start, end, tag_start, tag_end;
  size_t parent, first, last, prev, next, table_wrappers;
  unsigned char element, inert, removed, normalize_class, before_newline, append_newline,
      wrap_table, suppress_table_wrapper, checkbox, checked, synthetic, alert, footnote, arrow,
      relative_image, image_link, image_rejected, image_paragraph, comment, table_marks;
} gfm_html_node;
typedef struct {
  const unsigned char *source;
  size_t length, count, capacity;
  gfm_html_node *nodes;
  const char *token;
  const char *table_marker;
  const char *text_boundary;
  int had_text_boundary, split_comments;
} gfm_html_tree;
#define GFM_NONE SIZE_MAX

static size_t gfm_tree_append(gfm_html_tree *tree, size_t parent, gfm_html_node node) {
  if (!gfm_grow((void **)&tree->nodes, &tree->capacity, tree->count + 1, sizeof(*tree->nodes)))
    return GFM_NONE;
  size_t id = tree->count++;
  node.parent = parent;
  node.first = node.last = node.next = GFM_NONE;
  node.prev = tree->nodes[parent].last;
  tree->nodes[id] = node;
  if (node.prev != GFM_NONE)
    tree->nodes[node.prev].next = id;
  else
    tree->nodes[parent].first = id;
  tree->nodes[parent].last = id;
  return id;
}
static int gfm_node_is(gfm_html_tree *tree, size_t id, const char *name) {
  gfm_html_node *node = &tree->nodes[id];
  return node->element &&
         gfm_ascii_equal(tree->source + node->tag_start, node->tag_end - node->tag_start, name);
}
static int gfm_void_element(const unsigned char *name, size_t length) {
  static const char *const tags[] = {"area",  "base", "br",   "col",   "embed",  "hr",    "img",
                                     "input", "link", "meta", "param", "source", "track", "wbr"};
  for (size_t i = 0; i < sizeof(tags) / sizeof(tags[0]); i++)
    if (gfm_ascii_equal(name, length, tags[i]))
      return 1;
  return 0;
}
static int gfm_raw_element(const unsigned char *name, size_t length) {
  static const char *const tags[] = {"script", "style",   "textarea", "title",    "xmp",
                                     "iframe", "noembed", "noframes", "plaintext"};
  for (size_t i = 0; i < sizeof(tags) / sizeof(tags[0]); i++)
    if (gfm_ascii_equal(name, length, tags[i]))
      return 1;
  return 0;
}
static int gfm_name_char(unsigned char c) {
  return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' ||
         c == ':';
}
static int gfm_parse_serialized_html(gfm_html_tree *tree) {
  tree->count = tree->capacity = 0;
  tree->nodes = NULL;
  if (!gfm_grow((void **)&tree->nodes, &tree->capacity, 1, sizeof(*tree->nodes)))
    return 0;
  tree->nodes[tree->count++] = (gfm_html_node){
      .parent = GFM_NONE, .first = GFM_NONE, .last = GFM_NONE, .prev = GFM_NONE, .next = GFM_NONE};
  size_t *stack = NULL, depth = 0, capacity = 0;
  if (!gfm_grow((void **)&stack, &capacity, 1, sizeof(*stack))) {
    free(tree->nodes);
    tree->nodes = NULL;
    return 0;
  }
  stack[depth++] = 0;
  const unsigned char *s = tree->source;
  size_t i = 0, text_start = 0;
  while (i < tree->length) {
    if (s[i] != '<') {
      i++;
      continue;
    }
    size_t active = stack[depth - 1];
    if (active && gfm_raw_element(s + tree->nodes[active].tag_start,
                                  tree->nodes[active].tag_end - tree->nodes[active].tag_start)) {
      size_t name_length = tree->nodes[active].tag_end - tree->nodes[active].tag_start;
      size_t after = i + 2 + name_length;
      int closes = i + 2 <= tree->length && s[i + 1] == '/' && after <= tree->length &&
                   (after == tree->length || gfm_ascii_space(s[after]) || s[after] == '>' ||
                    s[after] == '/');
      for (size_t k = 0; closes && k < name_length; k++) {
        unsigned char a = s[i + 2 + k], b = s[tree->nodes[active].tag_start + k];
        if (a >= 'A' && a <= 'Z')
          a += 'a' - 'A';
        if (b >= 'A' && b <= 'Z')
          b += 'a' - 'A';
        if (a != b)
          closes = 0;
      }
      if (!closes) {
        i++;
        continue;
      }
    }
    if (i + 4 <= tree->length && !memcmp(s + i, "<!--", 4)) {
      size_t marker_start = i;
      i += 4;
      while (i + 3 <= tree->length && memcmp(s + i, "-->", 3))
        i++;
      size_t end = i + 3 <= tree->length ? i + 3 : tree->length;
      if (tree->split_comments) {
        size_t current = stack[depth - 1];
        if (marker_start > text_start) {
          gfm_html_node text = {.start = text_start, .end = marker_start};
          if (gfm_tree_append(tree, current, text) == GFM_NONE)
            goto fail;
        }
        gfm_html_node comment = {.start = marker_start, .end = end, .comment = 1};
        if (gfm_tree_append(tree, current, comment) == GFM_NONE)
          goto fail;
        i = text_start = end;
        continue;
      }
      if (tree->text_boundary && end >= text_start + 7 &&
          end - marker_start == strlen(tree->text_boundary) + 7 &&
          !memcmp(s + marker_start + 4, tree->text_boundary, strlen(tree->text_boundary))) {
        size_t marker = marker_start;
        if (marker > text_start) {
          gfm_html_node text = {.start = text_start, .end = marker};
          if (gfm_tree_append(tree, stack[depth - 1], text) == GFM_NONE)
            goto fail;
        }
        tree->had_text_boundary = 1;
        text_start = end;
      }
      i = end;
      continue;
    }
    size_t tag_start = i + 1;
    int closing = tag_start < tree->length && s[tag_start] == '/';
    if (closing)
      tag_start++;
    size_t tag_end = tag_start;
    while (tag_end < tree->length && gfm_name_char(s[tag_end]))
      tag_end++;
    if (tag_end == tag_start) {
      i++;
      continue;
    }
    size_t current = stack[depth - 1];
    if (current &&
        gfm_raw_element(s + tree->nodes[current].tag_start,
                        tree->nodes[current].tag_end - tree->nodes[current].tag_start) &&
        !(closing &&
          tag_end - tag_start == tree->nodes[current].tag_end - tree->nodes[current].tag_start &&
          !memcmp(s + tag_start, s + tree->nodes[current].tag_start, tag_end - tag_start))) {
      i++;
      continue;
    }
    size_t end = gfm_tag_end(s, tree->length, tag_end);
    if (i > text_start) {
      gfm_html_node text = {.start = text_start, .end = i};
      if (gfm_tree_append(tree, current, text) == GFM_NONE)
        goto fail;
    }
    if (closing) {
      if (depth > 1 &&
          tag_end - tag_start == tree->nodes[current].tag_end - tree->nodes[current].tag_start &&
          !memcmp(s + tag_start, s + tree->nodes[current].tag_start, tag_end - tag_start)) {
        tree->nodes[current].close_start = i;
        tree->nodes[current].end = end;
        depth--;
      } else {
        gfm_html_node text = {.start = i, .end = end};
        if (gfm_tree_append(tree, current, text) == GFM_NONE)
          goto fail;
      }
    } else {
      int inert = tree->nodes[current].inert || gfm_node_is(tree, current, "template");
      gfm_html_node element = {.start = i,
                               .open_end = end,
                               .close_start = end,
                               .end = end,
                               .tag_start = tag_start,
                               .tag_end = tag_end,
                               .element = 1,
                               .inert = inert};
      size_t id = gfm_tree_append(tree, current, element);
      if (id == GFM_NONE)
        goto fail;
      if (!gfm_void_element(s + tag_start, tag_end - tag_start)) {
        if (!gfm_grow((void **)&stack, &capacity, depth + 1, sizeof(*stack)))
          goto fail;
        stack[depth++] = id;
      }
    }
    i = text_start = end;
  }
  if (text_start < tree->length) {
    gfm_html_node text = {.start = text_start, .end = tree->length};
    if (gfm_tree_append(tree, stack[depth - 1], text) == GFM_NONE)
      goto fail;
  }
  free(stack);
  return 1;
fail:
  free(stack);
  free(tree->nodes);
  tree->nodes = NULL;
  return 0;
}

static void gfm_detach_node(gfm_html_tree *tree, size_t id) {
  gfm_html_node *n = &tree->nodes[id], *parent = &tree->nodes[n->parent];
  if (n->prev == GFM_NONE)
    parent->first = n->next;
  else
    tree->nodes[n->prev].next = n->next;
  if (n->next == GFM_NONE)
    parent->last = n->prev;
  else
    tree->nodes[n->next].prev = n->prev;
  n->parent = n->prev = n->next = GFM_NONE;
}
static void gfm_insert_before(gfm_html_tree *tree, size_t parent_id, size_t id, size_t anchor) {
  gfm_html_node *parent = &tree->nodes[parent_id], *n = &tree->nodes[id];
  n->parent = parent_id;
  n->next = anchor;
  n->prev = anchor == GFM_NONE ? parent->last : tree->nodes[anchor].prev;
  if (n->prev == GFM_NONE)
    parent->first = id;
  else
    tree->nodes[n->prev].next = id;
  if (anchor == GFM_NONE)
    parent->last = id;
  else
    tree->nodes[anchor].prev = id;
}
/* Read an attribute value from a normalized DOM start tag. */
static int gfm_get_attribute(gfm_html_tree *tree, size_t id, const char *name, size_t *start,
                             size_t *end) {
  const unsigned char *s = tree->source;
  gfm_html_node *node = &tree->nodes[id];
  size_t p = node->tag_end;
  while (p < node->open_end && s[p] != '>') {
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
    if (gfm_ascii_equal(s + attr, attr_end - attr, name)) {
      if (start)
        *start = a;
      if (end)
        *end = b;
      return 1;
    }
  }
  return 0;
}
static int gfm_attribute_equals(gfm_html_tree *tree, size_t id, const char *name,
                                const char *wanted) {
  size_t start, end;
  if (!gfm_get_attribute(tree, id, name, &start, &end))
    return 0;
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + start, end - start);
  int matches = value.size == strlen(wanted) && !memcmp(value.ptr, wanted, value.size);
  cmark_strbuf_free(&value);
  return matches;
}
static int gfm_class_contains(gfm_html_tree *tree, size_t id, const char *wanted) {
  size_t start, end;
  if (!gfm_get_attribute(tree, id, "class", &start, &end))
    return 0;
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + start, end - start);
  int matches = 0;
  for (size_t i = 0; i < (size_t)value.size;) {
    while (i < (size_t)value.size && gfm_ascii_space(value.ptr[i]))
      i++;
    size_t begin = i;
    while (i < (size_t)value.size && !gfm_ascii_space(value.ptr[i]))
      i++;
    if (i - begin == strlen(wanted) && !memcmp(value.ptr + begin, wanted, i - begin)) {
      matches = 1;
      break;
    }
  }
  cmark_strbuf_free(&value);
  return matches;
}
static size_t gfm_nearest(gfm_html_tree *tree, size_t id, const char *a, const char *b) {
  for (size_t n = id; n != GFM_NONE; n = tree->nodes[n].parent)
    if (gfm_node_is(tree, n, a) || (b && gfm_node_is(tree, n, b)))
      return n;
  return GFM_NONE;
}
static int gfm_checkbox(gfm_html_tree *tree, size_t id) {
  return gfm_node_is(tree, id, "input") && gfm_attribute_equals(tree, id, "type", "checkbox");
}

/* DOM textContent ignores template contents and comments. */
static int gfm_text_has_visible_content(gfm_html_tree *tree, size_t id) {
  gfm_html_node *node = &tree->nodes[id];
  if (node->element)
    return 1;
  if (node->synthetic)
    return !gfm_js_whitespace(node->synthetic);
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + node->start, node->end - node->start);
  int visible = 0;
  for (size_t i = 0; i < (size_t)value.size;) {
    uint32_t cp = gfm_utf8_at(value.ptr, value.size, i);
    if (!gfm_js_whitespace(cp)) {
      visible = 1;
      break;
    }
    i++;
    while (i < (size_t)value.size && (value.ptr[i] & 0xc0) == 0x80)
      i++;
  }
  cmark_strbuf_free(&value);
  return visible;
}
static int gfm_contains_math(gfm_html_tree *tree, size_t id) {
  for (size_t n = id; n != GFM_NONE;) {
    if (gfm_node_is(tree, n, "math-renderer"))
      return 1;
    if (!gfm_node_is(tree, n, "template") && tree->nodes[n].first != GFM_NONE) {
      n = tree->nodes[n].first;
      continue;
    }
    while (n != id && tree->nodes[n].next == GFM_NONE)
      n = tree->nodes[n].parent;
    if (n == id)
      break;
    n = tree->nodes[n].next;
  }
  return 0;
}
static int gfm_text_ends_in_newline(gfm_html_tree *tree, size_t id) {
  size_t n = tree->nodes[id].last;
  while (n != GFM_NONE && n != id) {
    gfm_html_node *node = &tree->nodes[n];
    if (node->element && !gfm_node_is(tree, n, "template") && node->last != GFM_NONE) {
      n = node->last;
      continue;
    }
    if (!node->element && node->synthetic)
      return node->synthetic == '\n';
    if (!node->element && node->end > node->start) {
      cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      houdini_unescape_html_f(&value, tree->source + node->start, node->end - node->start);
      if (value.size) {
        int result = value.ptr[value.size - 1] == '\n';
        cmark_strbuf_free(&value);
        return result;
      }
      cmark_strbuf_free(&value);
    }
    while (n != id && tree->nodes[n].prev == GFM_NONE)
      n = tree->nodes[n].parent;
    if (n == id)
      break;
    n = tree->nodes[n].prev;
  }
  return 0;
}
static int gfm_text_starts_newline(gfm_html_tree *tree, size_t id) {
  if (id == GFM_NONE || tree->nodes[id].element)
    return 0;
  if (tree->nodes[id].synthetic)
    return tree->nodes[id].synthetic == '\n';
  gfm_html_node *node = &tree->nodes[id];
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + node->start, node->end - node->start);
  int yes = value.size && value.ptr[0] == '\n';
  cmark_strbuf_free(&value);
  return yes;
}
