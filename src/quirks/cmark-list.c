/* A small tree of already-serialized, sanitized DOM nodes. It preserves every
 * byte except for the list moves and classList edits performed below. */
typedef struct {
  size_t start, open_end, close_start, end, tag_start, tag_end;
  size_t parent, first, last, prev, next;
  unsigned char element, inert, removed, normalize_class, before_newline, append_newline, wrap_table, suppress_table_wrapper, checkbox, checked, synthetic, alert, footnote, arrow, relative_image, image_link, image_rejected, image_paragraph;
} gfm_html_node;
typedef struct {
  const unsigned char *source;
  size_t length, count, capacity;
  gfm_html_node *nodes;
  const char *token;
  const char *table_marker;
  const char *text_boundary;
  int had_text_boundary;
} gfm_html_tree;
#define GFM_NONE SIZE_MAX

static size_t gfm_tree_append(gfm_html_tree *tree, size_t parent, gfm_html_node node) {
  if (!gfm_grow((void **)&tree->nodes, &tree->capacity, tree->count + 1, sizeof(*tree->nodes))) return GFM_NONE;
  size_t id = tree->count++;
  node.parent = parent; node.first = node.last = node.next = GFM_NONE;
  node.prev = tree->nodes[parent].last;
  tree->nodes[id] = node;
  if (node.prev != GFM_NONE) tree->nodes[node.prev].next = id;
  else tree->nodes[parent].first = id;
  tree->nodes[parent].last = id;
  return id;
}
static int gfm_node_is(gfm_html_tree *tree, size_t id, const char *name) {
  gfm_html_node *node = &tree->nodes[id];
  return node->element && gfm_ascii_equal(tree->source + node->tag_start, node->tag_end - node->tag_start, name);
}
static int gfm_void_element(const unsigned char *name, size_t length) {
  static const char *const tags[] = {"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"};
  for (size_t i = 0; i < sizeof(tags) / sizeof(tags[0]); i++)
    if (gfm_ascii_equal(name, length, tags[i])) return 1;
  return 0;
}
static int gfm_raw_element(const unsigned char *name, size_t length) {
  static const char *const tags[] = {"script", "style"};
  for (size_t i = 0; i < sizeof(tags) / sizeof(tags[0]); i++)
    if (gfm_ascii_equal(name, length, tags[i])) return 1;
  return 0;
}
static int gfm_name_char(unsigned char c) {
  return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == ':';
}
static int gfm_parse_serialized_html(gfm_html_tree *tree) {
  tree->count = tree->capacity = 0; tree->nodes = NULL;
  if (!gfm_grow((void **)&tree->nodes, &tree->capacity, 1, sizeof(*tree->nodes))) return 0;
  tree->nodes[tree->count++] = (gfm_html_node){ .parent=GFM_NONE, .first=GFM_NONE, .last=GFM_NONE, .prev=GFM_NONE, .next=GFM_NONE };
  size_t *stack = NULL, depth = 0, capacity = 0;
  if (!gfm_grow((void **)&stack, &capacity, 1, sizeof(*stack))) { free(tree->nodes); tree->nodes = NULL; return 0; }
  stack[depth++] = 0;
  const unsigned char *s = tree->source;
  size_t i = 0, text_start = 0;
  while (i < tree->length) {
    if (s[i] != '<') { i++; continue; }
    if (i + 4 <= tree->length && !memcmp(s + i, "<!--", 4)) {
      size_t marker_start = i;
      i += 4;
      while (i + 3 <= tree->length && memcmp(s + i, "-->", 3)) i++;
      size_t end = i + 3 <= tree->length ? i + 3 : tree->length;
      if (tree->text_boundary && end >= text_start + 7 &&
          end - marker_start == strlen(tree->text_boundary) + 7 &&
          !memcmp(s + marker_start + 4, tree->text_boundary, strlen(tree->text_boundary))) {
        size_t marker = marker_start;
        if (marker > text_start) {
          gfm_html_node text = { .start=text_start, .end=marker };
          if (gfm_tree_append(tree, stack[depth - 1], text) == GFM_NONE) goto fail;
        }
        tree->had_text_boundary = 1;
        text_start = end;
      }
      i = end;
      continue;
    }
    size_t tag_start = i + 1;
    int closing = tag_start < tree->length && s[tag_start] == '/';
    if (closing) tag_start++;
    size_t tag_end = tag_start;
    while (tag_end < tree->length && gfm_name_char(s[tag_end])) tag_end++;
    if (tag_end == tag_start) { i++; continue; }
    size_t current = stack[depth - 1];
    if (current && gfm_raw_element(s + tree->nodes[current].tag_start, tree->nodes[current].tag_end - tree->nodes[current].tag_start) &&
        !(closing &&
          tag_end - tag_start == tree->nodes[current].tag_end - tree->nodes[current].tag_start &&
          !memcmp(s + tag_start, s + tree->nodes[current].tag_start, tag_end - tag_start))) { i++; continue; }
    size_t end = gfm_tag_end(s, tree->length, tag_end);
    if (i > text_start) {
      gfm_html_node text = { .start=text_start, .end=i };
      if (gfm_tree_append(tree, current, text) == GFM_NONE) goto fail;
    }
    if (closing) {
      if (depth > 1 && tag_end - tag_start == tree->nodes[current].tag_end - tree->nodes[current].tag_start &&
          !memcmp(s + tag_start, s + tree->nodes[current].tag_start, tag_end - tag_start)) {
        tree->nodes[current].close_start = i; tree->nodes[current].end = end; depth--;
      } else {
        gfm_html_node text = { .start=i, .end=end };
        if (gfm_tree_append(tree, current, text) == GFM_NONE) goto fail;
      }
    } else {
      int inert = tree->nodes[current].inert || gfm_node_is(tree, current, "template");
      gfm_html_node element = { .start=i, .open_end=end, .close_start=end, .end=end,
        .tag_start=tag_start, .tag_end=tag_end, .element=1, .inert=inert };
      size_t id = gfm_tree_append(tree, current, element);
      if (id == GFM_NONE) goto fail;
      if (!gfm_void_element(s + tag_start, tag_end - tag_start)) {
        if (!gfm_grow((void **)&stack, &capacity, depth + 1, sizeof(*stack))) goto fail;
        stack[depth++] = id;
      }
    }
    i = text_start = end;
  }
  if (text_start < tree->length) {
    gfm_html_node text = { .start=text_start, .end=tree->length };
    if (gfm_tree_append(tree, stack[depth - 1], text) == GFM_NONE) goto fail;
  }
  free(stack); return 1;
fail:
  free(stack); free(tree->nodes); tree->nodes = NULL; return 0;
}

static void gfm_detach_node(gfm_html_tree *tree, size_t id) {
  gfm_html_node *n = &tree->nodes[id], *parent = &tree->nodes[n->parent];
  if (n->prev == GFM_NONE) parent->first = n->next; else tree->nodes[n->prev].next = n->next;
  if (n->next == GFM_NONE) parent->last = n->prev; else tree->nodes[n->next].prev = n->prev;
  n->parent = n->prev = n->next = GFM_NONE;
}
static void gfm_insert_before(gfm_html_tree *tree, size_t parent_id, size_t id, size_t anchor) {
  gfm_html_node *parent = &tree->nodes[parent_id], *n = &tree->nodes[id];
  n->parent = parent_id; n->next = anchor;
  n->prev = anchor == GFM_NONE ? parent->last : tree->nodes[anchor].prev;
  if (n->prev == GFM_NONE) parent->first = id; else tree->nodes[n->prev].next = id;
  if (anchor == GFM_NONE) parent->last = id; else tree->nodes[anchor].prev = id;
}
/* Read an attribute value from a normalized DOM start tag. */
static int gfm_get_attribute(gfm_html_tree *tree, size_t id, const char *name, size_t *start, size_t *end) {
  const unsigned char *s = tree->source;
  gfm_html_node *node = &tree->nodes[id];
  size_t p = node->tag_end;
  while (p < node->open_end && s[p] != '>') {
    while (p < node->open_end && gfm_ascii_space(s[p])) p++;
    size_t attr = p;
    while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '=' && s[p] != '>' && s[p] != '/') p++;
    if (p == attr) { p++; continue; }
    size_t attr_end = p;
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
    if (gfm_ascii_equal(s + attr, attr_end - attr, name)) {
      if (start) *start = a;
      if (end) *end = b;
      return 1;
    }
  }
  return 0;
}
static int gfm_attribute_equals(gfm_html_tree *tree, size_t id, const char *name, const char *wanted) {
  size_t start, end;
  if (!gfm_get_attribute(tree, id, name, &start, &end)) return 0;
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + start, end - start);
  int matches = value.size == strlen(wanted) && !memcmp(value.ptr, wanted, value.size);
  cmark_strbuf_free(&value);
  return matches;
}
static int gfm_class_contains(gfm_html_tree *tree, size_t id, const char *wanted) {
  size_t start, end;
  if (!gfm_get_attribute(tree, id, "class", &start, &end)) return 0;
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + start, end - start);
  int matches = 0;
  for (size_t i = 0; i < (size_t)value.size;) {
    while (i < (size_t)value.size && gfm_ascii_space(value.ptr[i])) i++;
    size_t begin = i;
    while (i < (size_t)value.size && !gfm_ascii_space(value.ptr[i])) i++;
    if (i - begin == strlen(wanted) && !memcmp(value.ptr + begin, wanted, i - begin)) { matches = 1; break; }
  }
  cmark_strbuf_free(&value);
  return matches;
}
static size_t gfm_nearest(gfm_html_tree *tree, size_t id, const char *a, const char *b) {
  for (size_t n = id; n != GFM_NONE; n = tree->nodes[n].parent)
    if (gfm_node_is(tree, n, a) || (b && gfm_node_is(tree, n, b))) return n;
  return GFM_NONE;
}
static int gfm_checkbox(gfm_html_tree *tree, size_t id) {
  return gfm_node_is(tree, id, "input") && gfm_attribute_equals(tree, id, "type", "checkbox");
}

/* DOM textContent ignores template contents and comments. */
static int gfm_text_has_visible_content(gfm_html_tree *tree, size_t id) {
  gfm_html_node *node = &tree->nodes[id];
  if (node->element) return 1;
  if (node->synthetic) return !gfm_js_whitespace(node->synthetic);
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + node->start, node->end - node->start);
  int visible = 0;
  for (size_t i = 0; i < (size_t)value.size;) {
    uint32_t cp = gfm_utf8_at(value.ptr, value.size, i);
    if (!gfm_js_whitespace(cp)) { visible = 1; break; }
    i++;
    while (i < (size_t)value.size && (value.ptr[i] & 0xc0) == 0x80) i++;
  }
  cmark_strbuf_free(&value);
  return visible;
}
static int gfm_contains_math(gfm_html_tree *tree, size_t id) {
  for (size_t n = id; n != GFM_NONE;) {
    if (gfm_node_is(tree, n, "math-renderer")) return 1;
    if (!gfm_node_is(tree, n, "template") && tree->nodes[n].first != GFM_NONE) { n = tree->nodes[n].first; continue; }
    while (n != id && tree->nodes[n].next == GFM_NONE) n = tree->nodes[n].parent;
    if (n == id) break;
    n = tree->nodes[n].next;
  }
  return 0;
}
static int gfm_text_ends_in_newline(gfm_html_tree *tree, size_t id) {
  size_t n = tree->nodes[id].last;
  while (n != GFM_NONE && n != id) {
    gfm_html_node *node = &tree->nodes[n];
    if (node->element && !gfm_node_is(tree, n, "template") && node->last != GFM_NONE) { n = node->last; continue; }
    if (!node->element && node->synthetic) return node->synthetic == '\n';
    if (!node->element && node->end > node->start) {
      cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      houdini_unescape_html_f(&value, tree->source + node->start, node->end - node->start);
      if (value.size) { int result = value.ptr[value.size - 1] == '\n'; cmark_strbuf_free(&value); return result; }
      cmark_strbuf_free(&value);
    }
    while (n != id && tree->nodes[n].prev == GFM_NONE) n = tree->nodes[n].parent;
    if (n == id) break;
    n = tree->nodes[n].prev;
  }
  return 0;
}
static int gfm_fix_list_math_spacing(gfm_html_tree *tree) {
  int changed = 0;
  for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || !gfm_node_is(tree, id, "li")) continue;
    size_t first = GFM_NONE, second = GFM_NONE, last = GFM_NONE, count = 0;
    for (size_t child = tree->nodes[id].first; child != GFM_NONE; child = tree->nodes[child].next) {
      if (!gfm_text_has_visible_content(tree, child)) continue;
      if (!count) first = child; else if (count == 1) second = child;
      last = child; count++;
    }
    if (count > 1 && tree->nodes[last].element && gfm_contains_math(tree, last) && !gfm_text_ends_in_newline(tree, id)) {
      tree->nodes[id].append_newline = 1; changed = 1;
    }
    if (count > 1 && first == tree->nodes[id].first && gfm_node_is(tree, first, "math-renderer") &&
        !tree->nodes[second].element && gfm_text_has_visible_content(tree, second)) {
      tree->nodes[first].before_newline = 1; changed = 1;
    }
  }
  return changed;
}

static int gfm_text_starts_newline(gfm_html_tree *tree, size_t id) {
  if (id == GFM_NONE || tree->nodes[id].element) return 0;
  if (tree->nodes[id].synthetic) return tree->nodes[id].synthetic == '\n';
  gfm_html_node *node = &tree->nodes[id];
  cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&value, tree->source + node->start, node->end - node->start);
  int yes = value.size && value.ptr[0] == '\n';
  cmark_strbuf_free(&value);
  return yes;
}
static int gfm_adjust_checkboxes(gfm_html_tree *tree, int *changed) {
  size_t original_count = tree->count;
  for (size_t id = 1; id < original_count; id++) {
    if (tree->nodes[id].inert || !gfm_checkbox(tree, id)) continue;
    size_t item = gfm_nearest(tree, id, "li", NULL);
    if (item == GFM_NONE) { gfm_detach_node(tree, id); tree->nodes[id].removed = 1; *changed = 1; continue; }
    size_t list = gfm_nearest(tree, id, "ul", "ol");
    tree->nodes[item].normalize_class = 3;
    if (list != GFM_NONE) tree->nodes[list].normalize_class = 4;
    tree->nodes[id].checkbox = 1;
    tree->nodes[id].checked = gfm_get_attribute(tree, id, "checked", NULL, NULL);
    *changed = 1;
  }
  for (size_t id = 1; id < original_count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed || !gfm_node_is(tree, id, "li") ||
        (tree->nodes[id].normalize_class != 3 && !gfm_class_contains(tree, id, "task-list-item"))) continue;
    size_t input = tree->nodes[id].first;
    while (input != GFM_NONE && !tree->nodes[input].element) input = tree->nodes[input].next;
    if (input == GFM_NONE || !gfm_checkbox(tree, input) || tree->nodes[input].removed) continue;
    size_t paragraph = tree->nodes[input].next;
    while (paragraph != GFM_NONE && !tree->nodes[paragraph].element) paragraph = tree->nodes[paragraph].next;
    if (paragraph == GFM_NONE || !gfm_node_is(tree, paragraph, "p")) continue;
    size_t whitespace = tree->nodes[input].next;
    if (whitespace != GFM_NONE && !tree->nodes[whitespace].element && !gfm_text_has_visible_content(tree, whitespace)) {
      gfm_detach_node(tree, whitespace); tree->nodes[whitespace].removed = 1;
    }
    gfm_detach_node(tree, input);
    gfm_html_node space = { .synthetic = ' ' };
    size_t space_id = gfm_tree_append(tree, 0, space);
    if (space_id == GFM_NONE) return 0;
    gfm_detach_node(tree, space_id);
    gfm_insert_before(tree, paragraph, space_id, tree->nodes[paragraph].first);
    gfm_insert_before(tree, paragraph, input, space_id);
    if (!gfm_text_starts_newline(tree, tree->nodes[id].first)) {
      gfm_html_node newline = { .synthetic = '\n' };
      size_t newline_id = gfm_tree_append(tree, 0, newline);
      if (newline_id == GFM_NONE) return 0;
      gfm_detach_node(tree, newline_id);
      gfm_insert_before(tree, id, newline_id, tree->nodes[id].first);
    }
    *changed = 1;
  }
  return 1;
}

/* Apply the alert conversion in document order, as the DOM loop does. */
static int gfm_adjust_alerts(gfm_html_tree *tree) {
  int changed = 0;
  for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed || !gfm_node_is(tree, id, "blockquote")) continue;
    int nested = 0;
    for (size_t parent = tree->nodes[id].parent; parent != GFM_NONE; parent = tree->nodes[parent].parent)
      if (gfm_node_is(tree, parent, "blockquote") && !tree->nodes[parent].alert) { nested = 1; break; }
    if (nested) continue;
    size_t paragraph = tree->nodes[id].first;
    while (paragraph != GFM_NONE && !tree->nodes[paragraph].element) paragraph = tree->nodes[paragraph].next;
    if (paragraph == GFM_NONE || !gfm_node_is(tree, paragraph, "p")) continue;
    size_t marker = tree->nodes[paragraph].first;
    if (marker == GFM_NONE || tree->nodes[marker].element || tree->nodes[marker].synthetic) continue;
    size_t br = tree->nodes[marker].next;
    if (br == GFM_NONE || !gfm_node_is(tree, br, "br")) continue;
    cmark_strbuf text = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    houdini_unescape_html_f(&text, tree->source + tree->nodes[marker].start,
                            tree->nodes[marker].end - tree->nodes[marker].start);
    size_t kind = 0;
    static const char *const markers[] = {"[!NOTE]", "[!TIP]", "[!IMPORTANT]", "[!WARNING]", "[!CAUTION]"};
    while (kind < sizeof(markers) / sizeof(markers[0]) &&
           (text.size != strlen(markers[kind]) || memcmp(text.ptr, markers[kind], text.size))) kind++;
    cmark_strbuf_free(&text);
    if (kind == sizeof(markers) / sizeof(markers[0])) continue;
    gfm_detach_node(tree, br); tree->nodes[br].removed = 1;
    gfm_detach_node(tree, marker); tree->nodes[marker].removed = 1;
    size_t first = tree->nodes[paragraph].first;
    if (first != GFM_NONE && !tree->nodes[first].element && !tree->nodes[first].synthetic &&
        tree->nodes[first].start < tree->nodes[first].end && tree->source[tree->nodes[first].start] == '\n')
      tree->nodes[first].start++;
    first = tree->nodes[id].first;
    if (first != GFM_NONE && !tree->nodes[first].element && !gfm_text_has_visible_content(tree, first)) {
      gfm_detach_node(tree, first); tree->nodes[first].removed = 1;
    }
    tree->nodes[id].alert = kind + 1; changed = 1;
  }
  return changed;
}

static int gfm_is_list(gfm_html_tree *tree, size_t id) {
  return id != GFM_NONE && (gfm_node_is(tree, id, "ul") || gfm_node_is(tree, id, "ol"));
}
static int gfm_flatten_lists(gfm_html_tree *tree) {
  int changed = 0;
  for (size_t id = 1; id < tree->count; id++) {
    gfm_html_node *list = &tree->nodes[id];
    if (list->inert || list->removed || !gfm_is_list(tree, id)) continue;
    size_t same = 1;
    for (size_t parent = list->parent; parent != GFM_NONE; parent = tree->nodes[parent].parent)
      if (gfm_is_list(tree, parent) && gfm_node_is(tree, parent, gfm_node_is(tree, id, "ul") ? "ul" : "ol")) same++;
    if (same <= 10) continue;
    size_t item = list->parent;
    if (item == GFM_NONE || !gfm_node_is(tree, item, "li")) continue;
    size_t parent = tree->nodes[item].parent;
    if (!gfm_is_list(tree, parent)) continue;
    size_t anchor = tree->nodes[item].next;
    for (size_t child = list->first; child != GFM_NONE;) {
      size_t next = tree->nodes[child].next;
      if (gfm_node_is(tree, child, "li")) {
        gfm_detach_node(tree, child);
        gfm_insert_before(tree, parent, child, anchor);
        tree->nodes[child].normalize_class = 1;
      }
      child = next;
    }
    gfm_detach_node(tree, id); list->removed = 1; changed = 1;
  }
  return changed;
}

/* Apply DOMTokenList normalization while preserving the attribute order. */
static void gfm_write_class(cmark_strbuf *out, const unsigned char *value, size_t length, int action) {
  const char *added = action == 2 ? "notranslate" : action == 3 ? "task-list-item" : action == 4 ? "contains-task-list" : action == 5 ? "task-list-item-checkbox" : NULL;
  cmark_strbuf decoded = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_strbuf tokens = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  houdini_unescape_html_f(&decoded, value, length);
  size_t n = decoded.size;
  for (size_t i = 0; i < n;) {
    while (i < n && gfm_ascii_space(decoded.ptr[i])) i++;
    size_t start = i;
    while (i < n && !gfm_ascii_space(decoded.ptr[i])) i++;
    size_t count = i - start;
    if (!count || (action == 1 && count == 14 && !memcmp(decoded.ptr + start, "task-list-item", 14))) continue;
    int duplicate = 0;
    for (size_t j = 0; j < (size_t)tokens.size;) {
      size_t previous = j;
      while (j < (size_t)tokens.size && tokens.ptr[j] != ' ') j++;
      if (j - previous == count && !memcmp(tokens.ptr + previous, decoded.ptr + start, count)) { duplicate = 1; break; }
      j++;
    }
    if (!duplicate) {
      if (tokens.size) cmark_strbuf_putc(&tokens, ' ');
      cmark_strbuf_put(&tokens, decoded.ptr + start, count);
    }
  }
  if (added) {
    size_t added_length = strlen(added);
    int found = 0;
    for (size_t j = 0; j < (size_t)tokens.size;) {
      size_t previous = j;
      while (j < (size_t)tokens.size && tokens.ptr[j] != ' ') j++;
      if (j - previous == added_length && !memcmp(tokens.ptr + previous, added, added_length)) found = 1;
      j++;
    }
    if (!found) { if (tokens.size) cmark_strbuf_putc(&tokens, ' '); cmark_strbuf_puts(&tokens, added); }
  }
  if (tokens.size) {
    cmark_strbuf_puts(out, " class=\"");
    for (size_t i = 0; i < (size_t)tokens.size; i++) {
      unsigned char ch = tokens.ptr[i];
      if (ch == '&') cmark_strbuf_puts(out, "&amp;");
      else if (ch == '"') cmark_strbuf_puts(out, "&quot;");
      else cmark_strbuf_putc(out, ch);
    }
    cmark_strbuf_putc(out, '"');
  }
  cmark_strbuf_free(&decoded); cmark_strbuf_free(&tokens);
}

static void gfm_write_normalized_tag(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node) {
  const unsigned char *s = tree->source;
  size_t p = node->tag_end, written = node->start;
  int saw_class = 0, action = node->normalize_class;
  while (p < node->open_end && s[p] != '>') {
    size_t space = p;
    while (p < node->open_end && gfm_ascii_space(s[p])) p++;
    size_t attr = p;
    while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '=' && s[p] != '>' && s[p] != '/') p++;
    if (p == attr) { p++; continue; }
    size_t attr_end = p;
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
    if (gfm_ascii_equal(s + attr, attr_end - attr, "class")) {
      cmark_strbuf_put(out, s + written, space - written);
      gfm_write_class(out, s + a, b - a, action);
      written = p; saw_class = 1;
    } else if (action == 2 && gfm_node_is(tree, (size_t)(node - tree->nodes), "pre") &&
               gfm_ascii_equal(s + attr, attr_end - attr, "lang")) {
      cmark_strbuf value = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      houdini_unescape_html_f(&value, s + a, b - a);
      int math = value.size == 4 && !memcmp(value.ptr, "math", 4);
      cmark_strbuf_free(&value);
      if (math) { cmark_strbuf_put(out, s + written, space - written); written = p; }
    }
  }
  if (action != 1 && !saw_class) {
    cmark_strbuf_put(out, s + written, node->open_end - 1 - written);
    gfm_write_class(out, (const unsigned char *)"", 0, action);
    cmark_strbuf_putc(out, '>');
  } else cmark_strbuf_put(out, s + written, node->open_end - written);
}

static void gfm_write_checkbox(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node) {
  const unsigned char *s = tree->source;
  size_t p = node->tag_end, written = node->start;
  int has_id = 0, has_class = 0, has_label = 0;
  const char *label = node->checked ? "Completed task" : "Incomplete task";
  while (p < node->open_end && s[p] != '>') {
    size_t space = p;
    while (p < node->open_end && gfm_ascii_space(s[p])) p++;
    size_t attr = p;
    while (p < node->open_end && !gfm_ascii_space(s[p]) && s[p] != '=' && s[p] != '>' && s[p] != '/') p++;
    if (p == attr) { p++; continue; }
    size_t attr_end = p;
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
    int checked = gfm_ascii_equal(s + attr, attr_end - attr, "checked");
    int disabled = gfm_ascii_equal(s + attr, attr_end - attr, "disabled");
    int id = gfm_ascii_equal(s + attr, attr_end - attr, "id");
    int class_name = gfm_ascii_equal(s + attr, attr_end - attr, "class");
    int aria = gfm_ascii_equal(s + attr, attr_end - attr, "aria-label");
    if (!(checked || disabled || id || class_name || aria)) continue;
    cmark_strbuf_put(out, s + written, space - written);
    if (id) { cmark_strbuf_puts(out, " id=\"\""); has_id = 1; }
    if (class_name) { gfm_write_class(out, s + a, b - a, 5); has_class = 1; }
    if (aria) { cmark_strbuf_puts(out, " aria-label=\""); cmark_strbuf_puts(out, label); cmark_strbuf_putc(out, '"'); has_label = 1; }
    written = p;
  }
  cmark_strbuf_put(out, s + written, node->open_end - 1 - written);
  if (!has_id) cmark_strbuf_puts(out, " id=\"\"");
  cmark_strbuf_puts(out, " disabled=\"\"");
  if (!has_class) cmark_strbuf_puts(out, " class=\"task-list-item-checkbox\"");
  if (!has_label) { cmark_strbuf_puts(out, " aria-label=\""); cmark_strbuf_puts(out, label); cmark_strbuf_putc(out, '"'); }
  if (node->checked) cmark_strbuf_puts(out, " checked=\"\"");
  cmark_strbuf_putc(out, '>');
}

#include "cmark-footnote-html.c"

static int gfm_write_table_tag(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node) {
  gfm_attributes attrs = {0};
  if (!gfm_attributes_read(tree, node, &attrs)) { gfm_attributes_free(&attrs); return 0; }
  if (tree->table_marker && *tree->table_marker) gfm_attribute_remove(&attrs, tree->table_marker);
  if (!gfm_attribute_set(&attrs, "role", (const unsigned char *)"table", 5)) { gfm_attributes_free(&attrs); return 0; }
  if (!node->suppress_table_wrapper) cmark_strbuf_puts(out, "<markdown-accessiblity-table>");
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
  gfm_attributes_free(&attrs);
  return 1;
}

static int gfm_serialize_tree(gfm_html_tree *tree, cmark_strbuf *out) {
  size_t id = tree->nodes[0].first;
  while (id != GFM_NONE) {
    gfm_html_node *node = &tree->nodes[id];
    if (node->before_newline) cmark_strbuf_putc(out, '\n');
    if (node->image_paragraph) cmark_strbuf_puts(out, "<p>");
    if (node->synthetic) cmark_strbuf_putc(out, node->synthetic);
    else if (!node->element) cmark_strbuf_put(out, tree->source + node->start, node->end - node->start);
    else if (node->checkbox) gfm_write_checkbox(out, tree, node);
    else if (node->relative_image) { if (!gfm_write_image(out, tree, node)) return 0; }
    else if (node->footnote) {
      int arrow = 0;
      if (!gfm_write_footnote_tag(out, tree, node, tree->token, &arrow)) return 0;
      node->arrow = arrow;
    }
    else if (node->alert) {
      size_t kind = node->alert - 1;
      cmark_strbuf_puts(out, "<div class=\"markdown-alert markdown-alert-");
      cmark_strbuf_puts(out, gfm_alert_names[kind]);
      cmark_strbuf_puts(out, "\"><p class=\"markdown-alert-title\">");
      cmark_strbuf_puts(out, gfm_alert_titles[kind]);
      cmark_strbuf_puts(out, "</p>");
    }
    else if (node->wrap_table) { if (!gfm_write_table_tag(out, tree, node)) return 0; }
    else if (node->normalize_class) gfm_write_normalized_tag(out, tree, node);
    else cmark_strbuf_put(out, tree->source + node->start, node->open_end - node->start);
    if (!node->arrow && node->first != GFM_NONE) { id = node->first; continue; }
    while (id != 0) {
      node = &tree->nodes[id];
      if (node->append_newline) cmark_strbuf_putc(out, '\n');
      if (node->arrow) { /* Replaced by the back-reference arrow text. */ }
      else if (node->alert) cmark_strbuf_puts(out, "</div>");
      else if (node->element && node->end > node->close_start)
        cmark_strbuf_put(out, tree->source + node->close_start, node->end - node->close_start);
      if (node->wrap_table && !node->suppress_table_wrapper) cmark_strbuf_puts(out, "</markdown-accessiblity-table>");
      if (node->image_paragraph) cmark_strbuf_puts(out, "</p>");
      if (node->next != GFM_NONE) { id = node->next; break; }
      id = node->parent;
    }
    if (id == 0) break;
  }
  return 1;
}
char *gfm_finalize_serialized_html(const unsigned char *source, size_t length, const char *token, const char *table_marker) {
  gfm_html_tree tree = { .source = source, .length = length, .token = token, .table_marker = table_marker };
  if (!gfm_parse_serialized_html(&tree)) return NULL;
  int changed = gfm_adjust_alerts(&tree);
  if (!gfm_adjust_checkboxes(&tree, &changed)) { free(tree.nodes); return NULL; }
  changed |= gfm_fix_list_math_spacing(&tree);
  changed |= gfm_adjust_images(&tree);
  changed |= gfm_adjust_footnote_nodes(&tree);
  changed |= gfm_flatten_lists(&tree);
  for (size_t i = 1; i < tree.count; i++) {
    if (!tree.nodes[i].inert && gfm_node_is(&tree, i, "table")) {
      tree.nodes[i].wrap_table = 1;
      tree.nodes[i].suppress_table_wrapper = table_marker && *table_marker && gfm_get_attribute(&tree, i, table_marker, NULL, NULL);
      changed = 1;
    }
    if (!tree.nodes[i].inert && (gfm_node_is(&tree, i, "code") || gfm_node_is(&tree, i, "pre"))) {
      tree.nodes[i].normalize_class = 2; changed = 1;
    }
  }
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  if (changed && !gfm_serialize_tree(&tree, &out)) { cmark_strbuf_free(&out); free(tree.nodes); return NULL; }
  if (!changed) cmark_strbuf_put(&out, source, length);
  free(tree.nodes);
  return (char *)cmark_strbuf_detach(&out);
}
