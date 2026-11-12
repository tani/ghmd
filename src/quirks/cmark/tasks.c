/* Task-list state, placement, and checkbox serialization. */
static int gfm_adjust_checkboxes(gfm_html_tree *tree, int *changed) {
  size_t original_count = tree->count;
  for (size_t id = 1; id < original_count; id++) {
    if (tree->nodes[id].inert || !gfm_checkbox(tree, id))
      continue;
    size_t item = gfm_nearest(tree, id, "li", NULL);
    if (item == GFM_NONE) {
      gfm_detach_node(tree, id);
      tree->nodes[id].removed = 1;
      *changed = 1;
      continue;
    }
    size_t list = gfm_nearest(tree, id, "ul", "ol");
    tree->nodes[item].normalize_class = 3;
    if (list != GFM_NONE)
      tree->nodes[list].normalize_class = 4;
    tree->nodes[id].checkbox = 1;
    tree->nodes[id].checked = gfm_get_attribute(tree, id, "checked", NULL, NULL);
    *changed = 1;
  }
  for (size_t id = 1; id < original_count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed || !gfm_node_is(tree, id, "li") ||
        (tree->nodes[id].normalize_class != 3 && !gfm_class_contains(tree, id, "task-list-item")))
      continue;
    size_t input = tree->nodes[id].first;
    while (input != GFM_NONE && !tree->nodes[input].element)
      input = tree->nodes[input].next;
    if (input == GFM_NONE || !gfm_checkbox(tree, input) || tree->nodes[input].removed)
      continue;
    size_t paragraph = tree->nodes[input].next;
    while (paragraph != GFM_NONE && !tree->nodes[paragraph].element)
      paragraph = tree->nodes[paragraph].next;
    if (paragraph == GFM_NONE || !gfm_node_is(tree, paragraph, "p"))
      continue;
    size_t whitespace = tree->nodes[input].next;
    if (whitespace != GFM_NONE && !tree->nodes[whitespace].element &&
        !gfm_text_has_visible_content(tree, whitespace)) {
      gfm_detach_node(tree, whitespace);
      tree->nodes[whitespace].removed = 1;
    }
    gfm_detach_node(tree, input);
    gfm_html_node space = {.synthetic = ' '};
    size_t space_id = gfm_tree_append(tree, 0, space);
    if (space_id == GFM_NONE)
      return 0;
    gfm_detach_node(tree, space_id);
    gfm_insert_before(tree, paragraph, space_id, tree->nodes[paragraph].first);
    gfm_insert_before(tree, paragraph, input, space_id);
    if (!gfm_text_starts_newline(tree, tree->nodes[id].first)) {
      gfm_html_node newline = {.synthetic = '\n'};
      size_t newline_id = gfm_tree_append(tree, 0, newline);
      if (newline_id == GFM_NONE)
        return 0;
      gfm_detach_node(tree, newline_id);
      gfm_insert_before(tree, id, newline_id, tree->nodes[id].first);
    }
    *changed = 1;
  }
  return 1;
}

static void gfm_write_checkbox(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node) {
  const unsigned char *s = tree->source;
  size_t p = node->tag_end, written = node->start;
  int has_id = 0, has_class = 0, has_label = 0;
  const char *label = node->checked ? "Completed task" : "Incomplete task";
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
    int checked = gfm_ascii_equal(s + attr, attr_end - attr, "checked");
    int disabled = gfm_ascii_equal(s + attr, attr_end - attr, "disabled");
    int id = gfm_ascii_equal(s + attr, attr_end - attr, "id");
    int class_name = gfm_ascii_equal(s + attr, attr_end - attr, "class");
    int aria = gfm_ascii_equal(s + attr, attr_end - attr, "aria-label");
    if (!(checked || disabled || id || class_name || aria))
      continue;
    cmark_strbuf_put(out, s + written, space - written);
    if (id) {
      cmark_strbuf_puts(out, " id=\"\"");
      has_id = 1;
    }
    if (class_name) {
      gfm_write_class(out, s + a, b - a, 5);
      has_class = 1;
    }
    if (aria) {
      cmark_strbuf_puts(out, " aria-label=\"");
      cmark_strbuf_puts(out, label);
      cmark_strbuf_putc(out, '"');
      has_label = 1;
    }
    written = p;
  }
  cmark_strbuf_put(out, s + written, node->open_end - 1 - written);
  if (!has_id)
    cmark_strbuf_puts(out, " id=\"\"");
  cmark_strbuf_puts(out, " disabled=\"\"");
  if (!has_class)
    cmark_strbuf_puts(out, " class=\"task-list-item-checkbox\"");
  if (!has_label) {
    cmark_strbuf_puts(out, " aria-label=\"");
    cmark_strbuf_puts(out, label);
    cmark_strbuf_putc(out, '"');
  }
  if (node->checked)
    cmark_strbuf_puts(out, " checked=\"\"");
  cmark_strbuf_putc(out, '>');
}
