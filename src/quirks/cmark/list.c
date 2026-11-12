/* List flattening and spacing around rendered math. */
static int gfm_fix_list_math_spacing(gfm_html_tree *tree) {
  int changed = 0;
  for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || !gfm_node_is(tree, id, "li"))
      continue;
    size_t first = GFM_NONE, second = GFM_NONE, last = GFM_NONE, count = 0;
    for (size_t child = tree->nodes[id].first; child != GFM_NONE; child = tree->nodes[child].next) {
      if (!gfm_text_has_visible_content(tree, child))
        continue;
      if (!count)
        first = child;
      else if (count == 1)
        second = child;
      last = child;
      count++;
    }
    if (count > 1 && tree->nodes[last].element && gfm_contains_math(tree, last) &&
        !gfm_text_ends_in_newline(tree, id)) {
      tree->nodes[id].append_newline = 1;
      changed = 1;
    }
    if (count > 1 && first == tree->nodes[id].first && gfm_node_is(tree, first, "math-renderer") &&
        !tree->nodes[second].element && gfm_text_has_visible_content(tree, second)) {
      tree->nodes[first].before_newline = 1;
      changed = 1;
    }
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
    if (list->inert || list->removed || !gfm_is_list(tree, id))
      continue;
    size_t same = 1;
    for (size_t parent = list->parent; parent != GFM_NONE; parent = tree->nodes[parent].parent)
      if (gfm_is_list(tree, parent) &&
          gfm_node_is(tree, parent, gfm_node_is(tree, id, "ul") ? "ul" : "ol"))
        same++;
    if (same <= 10)
      continue;
    size_t item = list->parent;
    if (item == GFM_NONE || !gfm_node_is(tree, item, "li"))
      continue;
    size_t parent = tree->nodes[item].parent;
    if (!gfm_is_list(tree, parent))
      continue;
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
    gfm_detach_node(tree, id);
    list->removed = 1;
    changed = 1;
  }
  return changed;
}
