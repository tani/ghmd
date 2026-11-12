/* Insert GitHub's default summary before math reads details text. */
char *gfm_prepare_details(const unsigned char *source, size_t length) {
  gfm_html_tree tree = {.source = source, .length = length};
  if (!gfm_parse_serialized_html(&tree))
    return NULL;
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  size_t id = tree.nodes[0].first;
  while (id != GFM_NONE) {
    gfm_html_node *node = &tree.nodes[id];
    if (!node->element)
      cmark_strbuf_put(&out, source + node->start, node->end - node->start);
    else {
      cmark_strbuf_put(&out, source + node->start, node->open_end - node->start);
      if (!node->inert && gfm_node_is(&tree, id, "details")) {
        int has_summary = 0;
        for (size_t child = node->first; child != GFM_NONE; child = tree.nodes[child].next)
          if (gfm_node_is(&tree, child, "summary")) {
            has_summary = 1;
            break;
          }
        if (!has_summary)
          cmark_strbuf_puts(&out, "<summary>Details</summary>");
      }
    }
    if (node->first != GFM_NONE) {
      id = node->first;
      continue;
    }
    while (id != 0) {
      node = &tree.nodes[id];
      if (node->element && node->end > node->close_start)
        cmark_strbuf_put(&out, source + node->close_start, node->end - node->close_start);
      if (node->next != GFM_NONE) {
        id = node->next;
        break;
      }
      id = node->parent;
    }
    if (id == 0)
      break;
  }
  free(tree.nodes);
  return (char *)cmark_strbuf_detach(&out);
}

/* Template contents are moved into the parent without merging adjacent text
 * nodes. Comments representing those boundaries survive serialization. */
static int gfm_unwrap_templates(gfm_html_tree *tree) {
  for (size_t id = 1; id < tree->count; id++) {
    if (!gfm_node_is(tree, id, "template") || tree->nodes[id].parent == GFM_NONE)
      continue;
    size_t parent = tree->nodes[id].parent;
    while (tree->nodes[id].first != GFM_NONE) {
      size_t child = tree->nodes[id].first;
      gfm_detach_node(tree, child);
      gfm_insert_before(tree, parent, child, id);
    }
    gfm_detach_node(tree, id);
  }
  for (size_t id = 1; id < tree->count; id++)
    tree->nodes[id].inert = 0;
  return 1;
}

/* Serialize an existing tree without invoking a second HTML parser. */
static int gfm_prepare_write(gfm_html_tree *tree, cmark_strbuf *out, size_t root,
                             const char *boundary, const char *table_marker, int strip_comments,
                             int preserve_boundaries) {
  size_t id = root == 0 ? tree->nodes[0].first : root;
  if (id == GFM_NONE)
    return 1;
  while (id != GFM_NONE) {
    gfm_html_node *node = &tree->nodes[id];
    if (node->comment) {
      if (!strip_comments)
        cmark_strbuf_put(out, tree->source + node->start, node->end - node->start);
    } else if (!node->element) {
      if (preserve_boundaries && node->prev != GFM_NONE) {
        size_t prev = node->prev;
        while (strip_comments && prev != GFM_NONE && tree->nodes[prev].comment)
          prev = tree->nodes[prev].prev;
        if (prev != GFM_NONE && !tree->nodes[prev].element && !tree->nodes[prev].comment) {
          cmark_strbuf_puts(out, "<!--");
          cmark_strbuf_puts(out, boundary);
          cmark_strbuf_puts(out, "-->");
        }
      }
      cmark_strbuf_put(out, tree->source + node->start, node->end - node->start);
    } else {
      for (size_t n = 0; n < node->table_wrappers; n++)
        cmark_strbuf_puts(out, "<markdown-accessiblity-table>");
      if (node->table_marks) {
        cmark_strbuf_put(out, tree->source + node->start, node->open_end - node->start - 1);
        cmark_strbuf_putc(out, ' ');
        cmark_strbuf_puts(out, table_marker);
        cmark_strbuf_puts(out, "=\"\">");
      } else
        cmark_strbuf_put(out, tree->source + node->start, node->open_end - node->start);
    }
    if (node->first != GFM_NONE) {
      id = node->first;
      continue;
    }
    while (id != 0) {
      node = &tree->nodes[id];
      if (node->element && node->end > node->close_start)
        cmark_strbuf_put(out, tree->source + node->close_start, node->end - node->close_start);
      for (size_t n = 0; n < node->table_wrappers; n++)
        cmark_strbuf_puts(out, "</markdown-accessiblity-table>");
      if (id == root)
        return 1;
      if (node->next != GFM_NONE) {
        id = node->next;
        break;
      }
      id = node->parent;
    }
    if (id == 0)
      break;
  }
  return 1;
}

char *gfm_flatten_templates(const unsigned char *source, size_t length, const char *boundary) {
  gfm_html_tree tree = {.source = source, .length = length, .split_comments = 1};
  if (!gfm_parse_serialized_html(&tree))
    return NULL;
  gfm_unwrap_templates(&tree);
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  gfm_prepare_write(&tree, &out, 0, boundary, "", 0, 1);
  free(tree.nodes);
  return (char *)cmark_strbuf_detach(&out);
}

/* Remove only generated boundary comments before a browser DOM reparse. */
char *gfm_remove_boundaries(const unsigned char *source, size_t length, const char *boundary) {
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  size_t n = strlen(boundary), start = 0;
  for (size_t i = 0; i + n + 7 <= length; i++) {
    if (memcmp(source + i, "<!--", 4) || memcmp(source + i + 4, boundary, n) ||
        memcmp(source + i + 4 + n, "-->", 3))
      continue;
    cmark_strbuf_put(&out, source + start, i - start);
    i += n + 6;
    start = i + 1;
  }
  cmark_strbuf_put(&out, source + start, length - start);
  return (char *)cmark_strbuf_detach(&out);
}

/* Decide paragraph-local table wrapping before math, as the DOM phase did. */
char *gfm_prepare_math_tables(const unsigned char *source, size_t length, const char *token,
                              const char *boundary, const char *table_marker) {
  gfm_html_tree tree = {.source = source, .length = length, .split_comments = 1};
  if (!gfm_parse_serialized_html(&tree))
    return NULL;
  size_t *subtree_end = malloc(tree.count * sizeof(*subtree_end));
  if (!subtree_end) {
    free(tree.nodes);
    return NULL;
  }
  for (size_t id = 0; id < tree.count; id++)
    subtree_end[id] = id + 1;
  for (size_t id = tree.count; id-- > 1;) {
    size_t parent = tree.nodes[id].parent;
    if (parent != GFM_NONE && subtree_end[parent] < subtree_end[id])
      subtree_end[parent] = subtree_end[id];
  }
  for (size_t id = 1; id < tree.count; id++) {
    if (tree.nodes[id].inert || !gfm_node_is(&tree, id, "p"))
      continue;
    int blocked = 0, has_table = 0;
    for (size_t p = tree.nodes[id].parent; p != GFM_NONE; p = tree.nodes[p].parent) {
      if (gfm_node_is(&tree, p, "details"))
        break;
      if (gfm_node_is(&tree, p, "b") || gfm_node_is(&tree, p, "strong") ||
          gfm_node_is(&tree, p, "em") || gfm_node_is(&tree, p, "a") ||
          gfm_node_is(&tree, p, "code") || gfm_node_is(&tree, p, "pre") ||
          gfm_node_is(&tree, p, "section")) {
        blocked = 1;
        break;
      }
    }
    for (size_t n = id + 1; n < subtree_end[id]; n++) {
      if (gfm_node_is(&tree, n, "table")) {
        has_table = 1;
        break;
      }
    }
    if (!has_table)
      continue;
    int has_math = 0;
    if (!blocked) {
      cmark_strbuf paragraph = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      gfm_prepare_write(&tree, &paragraph, id, boundary, table_marker, 1, 0);
      char *rendered = gfm_math_dom(paragraph.ptr, paragraph.size, token, boundary);
      if (!rendered) {
        cmark_strbuf_free(&paragraph);
        free(subtree_end);
        free(tree.nodes);
        return NULL;
      }
      has_math =
          strlen(rendered) != paragraph.size || memcmp(rendered, paragraph.ptr, paragraph.size);
      free(rendered);
      cmark_strbuf_free(&paragraph);
    }
    for (size_t n = id + 1; n < subtree_end[id]; n++) {
      if (!gfm_node_is(&tree, n, "table"))
        continue;
      tree.nodes[n].table_marks = 1;
      if (!has_math)
        tree.nodes[n].table_wrappers++;
    }
  }
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  gfm_prepare_write(&tree, &out, 0, boundary, table_marker, 1, 1);
  free(subtree_end);
  free(tree.nodes);
  return (char *)cmark_strbuf_detach(&out);
}
