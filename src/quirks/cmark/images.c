/* Relative-image wrapping and image attribute serialization. */
static int gfm_external_image(const unsigned char *value, size_t length) {
  return gfm_ascii_prefix_case(value, length, "http://") ||
         gfm_ascii_prefix_case(value, length, "https://") ||
         gfm_ascii_prefix_case(value, length, "//");
}
static int gfm_adjust_images(gfm_html_tree *tree) {
  int changed = 0;
  for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed || !gfm_node_is(tree, id, "img"))
      continue;
    size_t start, end;
    if (!gfm_get_attribute(tree, id, "src", &start, &end))
      continue;
    size_t parent = tree->nodes[id].parent;
    if (parent == 0 || (parent != GFM_NONE && gfm_node_is(tree, parent, "b"))) {
      tree->nodes[id].image_paragraph = 1;
      changed = 1;
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
  if (!gfm_attributes_read(tree, node, &attrs)) {
    gfm_attributes_free(&attrs);
    return 0;
  }
  gfm_attribute *src_attr = gfm_attribute_find(&attrs, "src");
  cmark_strbuf *src = gfm_attribute_value(tree, src_attr);
  cmark_strbuf image_href = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  if (!node->image_rejected)
    cmark_strbuf_put(&image_href, src->ptr, src->size);
  if (node->image_rejected)
    gfm_attribute_remove(&attrs, "src");
  int ok = gfm_attribute_set(&attrs, "style", (const unsigned char *)"max-width: 100%;", 16);
  if (ok && node->image_link) {
    cmark_strbuf_puts(out, "<a target=\"_blank\" rel=\"");
    if (gfm_ascii_prefix_case(image_href.ptr, image_href.size, "mailto:"))
      cmark_strbuf_puts(out, "nofollow");
    else
      cmark_strbuf_puts(out, "noopener noreferrer");
    cmark_strbuf_puts(out, "\" href=\"");
    gfm_attribute_escape(out, image_href.ptr, image_href.size);
    cmark_strbuf_puts(out, "\">");
  }
  if (ok) {
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
    if (node->image_link)
      cmark_strbuf_puts(out, "</a>");
  }
  cmark_strbuf_free(&image_href);
  gfm_attributes_free(&attrs);
  return ok;
}
