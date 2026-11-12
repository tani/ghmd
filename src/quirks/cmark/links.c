/* Link edits after sanitization and math, over DOM-serialized HTML. */
static int gfm_write_link(cmark_strbuf *out, gfm_html_tree *tree, size_t id, const char *marker,
                          int *unwrap) {
  gfm_html_node *node = &tree->nodes[id];
  gfm_attributes attrs = {0};
  if (!gfm_attributes_read(tree, node, &attrs)) {
    gfm_attributes_free(&attrs);
    return 0;
  }
  gfm_attribute *marked = gfm_attribute_find(&attrs, "data-gfm-had-href");
  int had_href = 0;
  if (marked) {
    cmark_strbuf *value = gfm_attribute_value(tree, marked);
    had_href = value->size == strlen(marker) && !memcmp(value->ptr, marker, value->size);
    marked->removed = 1;
  }
  gfm_attribute *href = gfm_attribute_find(&attrs, "href");
  *unwrap =
      had_href && !href && !gfm_attribute_find(&attrs, "id") && !gfm_attribute_find(&attrs, "name");
  if (!*unwrap && href) {
    cmark_strbuf *value = gfm_attribute_value(tree, href);
    if (gfm_external_image(value->ptr, value->size) ||
        gfm_ascii_prefix_case(value->ptr, value->size, "mailto:")) {
      if (!gfm_attribute_set(&attrs, "rel", (const unsigned char *)"nofollow", 8)) {
        gfm_attributes_free(&attrs);
        return 0;
      }
    }
  }
  if (!*unwrap) {
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
  }
  gfm_attributes_free(&attrs);
  return 1;
}
char *gfm_process_links(const unsigned char *source, size_t length, const char *marker) {
  gfm_html_tree tree = {.source = source, .length = length};
  if (!gfm_parse_serialized_html(&tree))
    return NULL;
  cmark_strbuf out = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  size_t id = tree.nodes[0].first;
  while (id != GFM_NONE) {
    gfm_html_node *node = &tree.nodes[id];
    if (!node->element)
      cmark_strbuf_put(&out, source + node->start, node->end - node->start);
    else if (!node->inert && gfm_node_is(&tree, id, "a")) {
      int unwrap = 0;
      if (!gfm_write_link(&out, &tree, id, marker, &unwrap)) {
        cmark_strbuf_free(&out);
        free(tree.nodes);
        return NULL;
      }
      node->removed = unwrap;
    } else
      cmark_strbuf_put(&out, source + node->start, node->open_end - node->start);
    if (node->first != GFM_NONE) {
      id = node->first;
      continue;
    }
    while (id != 0) {
      node = &tree.nodes[id];
      if (node->element && !node->removed && node->end > node->close_start)
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
