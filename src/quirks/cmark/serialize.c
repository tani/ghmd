/* Final HTML serialization after all feature edits. */
static int gfm_write_table_tag(cmark_strbuf *out, gfm_html_tree *tree, gfm_html_node *node) {
  gfm_attributes attrs = {0};
  if (!gfm_attributes_read(tree, node, &attrs)) {
    gfm_attributes_free(&attrs);
    return 0;
  }
  if (tree->table_marker && *tree->table_marker)
    gfm_attribute_remove(&attrs, tree->table_marker);
  if (!gfm_attribute_set(&attrs, "role", (const unsigned char *)"table", 5)) {
    gfm_attributes_free(&attrs);
    return 0;
  }
  if (!node->suppress_table_wrapper)
    cmark_strbuf_puts(out, "<markdown-accessiblity-table>");
  cmark_strbuf_put(out, tree->source + node->start, node->tag_end - node->start);
  for (size_t i = 0; i < attrs.count; i++) {
    gfm_attribute *attr = &attrs.items[i];
    if (attr->removed)
      continue;
    if (!attr->changed)
      cmark_strbuf_put(out, tree->source + attr->raw_start, attr->raw_end - attr->raw_start);
    else {
      cmark_strbuf_putc(out, ' ');
      cmark_strbuf_put(out, attr->name, attr->name_length);
      cmark_strbuf_puts(out, "=\"");
      gfm_attribute_escape(out, attr->value.ptr, attr->value.size);
      cmark_strbuf_putc(out, '"');
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
    if (node->before_newline)
      cmark_strbuf_putc(out, '\n');
    if (node->image_paragraph)
      cmark_strbuf_puts(out, "<p>");
    if (node->synthetic)
      cmark_strbuf_putc(out, node->synthetic);
    else if (!node->element)
      cmark_strbuf_put(out, tree->source + node->start, node->end - node->start);
    else if (node->checkbox)
      gfm_write_checkbox(out, tree, node);
    else if (node->relative_image) {
      if (!gfm_write_image(out, tree, node))
        return 0;
    } else if (node->footnote) {
      int arrow = 0;
      if (!gfm_write_footnote_tag(out, tree, node, tree->token, &arrow))
        return 0;
      node->arrow = arrow;
    } else if (node->alert) {
      size_t kind = node->alert - 1;
      cmark_strbuf_puts(out, "<div class=\"markdown-alert markdown-alert-");
      cmark_strbuf_puts(out, gfm_alert_names[kind]);
      cmark_strbuf_puts(out, "\"><p class=\"markdown-alert-title\">");
      cmark_strbuf_puts(out, gfm_alert_titles[kind]);
      cmark_strbuf_puts(out, "</p>");
    } else if (node->wrap_table) {
      if (!gfm_write_table_tag(out, tree, node))
        return 0;
    } else if (node->normalize_class)
      gfm_write_normalized_tag(out, tree, node);
    else
      cmark_strbuf_put(out, tree->source + node->start, node->open_end - node->start);
    if (!node->arrow && node->first != GFM_NONE) {
      id = node->first;
      continue;
    }
    while (id != 0) {
      node = &tree->nodes[id];
      if (node->append_newline)
        cmark_strbuf_putc(out, '\n');
      if (node->arrow) { /* Replaced by the back-reference arrow text. */
      } else if (node->alert)
        cmark_strbuf_puts(out, "</div>");
      else if (node->element && node->end > node->close_start)
        cmark_strbuf_put(out, tree->source + node->close_start, node->end - node->close_start);
      if (node->wrap_table && !node->suppress_table_wrapper)
        cmark_strbuf_puts(out, "</markdown-accessiblity-table>");
      if (node->image_paragraph)
        cmark_strbuf_puts(out, "</p>");
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
char *gfm_finalize_serialized_html(const unsigned char *source, size_t length, const char *token,
                                   const char *table_marker) {
  gfm_html_tree tree = {
      .source = source, .length = length, .token = token, .table_marker = table_marker};
  if (!gfm_parse_serialized_html(&tree))
    return NULL;
  int changed = gfm_adjust_alerts(&tree);
  if (!gfm_adjust_checkboxes(&tree, &changed)) {
    free(tree.nodes);
    return NULL;
  }
  changed |= gfm_fix_list_math_spacing(&tree);
  changed |= gfm_adjust_images(&tree);
  changed |= gfm_adjust_footnote_nodes(&tree);
  changed |= gfm_flatten_lists(&tree);
  for (size_t i = 1; i < tree.count; i++) {
    if (!tree.nodes[i].inert && gfm_node_is(&tree, i, "table")) {
      tree.nodes[i].wrap_table = 1;
      tree.nodes[i].suppress_table_wrapper =
          table_marker && *table_marker && gfm_get_attribute(&tree, i, table_marker, NULL, NULL);
      changed = 1;
    }
    if (!tree.nodes[i].inert && (gfm_node_is(&tree, i, "code") || gfm_node_is(&tree, i, "pre"))) {
      tree.nodes[i].normalize_class = 2;
      changed = 1;
    }
  }
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  if (changed && !gfm_serialize_tree(&tree, &out)) {
    cmark_strbuf_free(&out);
    free(tree.nodes);
    return NULL;
  }
  if (!changed)
    cmark_strbuf_put(&out, source, length);
  free(tree.nodes);
  return (char *)cmark_strbuf_detach(&out);
}
