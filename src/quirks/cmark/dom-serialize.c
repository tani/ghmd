/* Serialize DOM nodes with the existing renderer's entity spelling. Attribute
 * angle brackets stay literal inside quotes, as in the former DOM serializer. */
static void gfm_dom_escape(h_buffer *out, const unsigned char *s, size_t length, int attribute) {
  for (size_t i = 0; i < length; i++) {
    if (s[i] == '&')
      h_s(out, "&amp;");
    else if (s[i] == '"' && attribute)
      h_s(out, "&quot;");
    else if (s[i] == '<' && !attribute)
      h_s(out, "&lt;");
    else if (s[i] == '>' && !attribute)
      h_s(out, "&gt;");
    else if (s[i] == 0xc2 && i + 1 < length && s[i + 1] == 0xa0) {
      h_s(out, "&nbsp;");
      i++;
    } else
      h_put(out, s + i, 1);
  }
}
static void gfm_dom_serialize_node(h_buffer *out, lxb_dom_node_t *node, unsigned depth, int raw,
                                   int aliases) {
  if (out->failed)
    return;
  if (depth > 512) {
    out->failed = H_DEPTH_EXCEEDED;
    return;
  }
  if (node->type == LXB_DOM_NODE_TYPE_TEXT || node->type == LXB_DOM_NODE_TYPE_COMMENT) {
    lexbor_str_t *text = &((lxb_dom_character_data_t *)node)->data;
    if (node->type == LXB_DOM_NODE_TYPE_COMMENT) {
      h_s(out, "<!--");
      h_put(out, text->data, text->length);
      h_s(out, "-->");
    } else if (raw)
      h_put(out, text->data, text->length);
    else
      gfm_dom_escape(out, text->data, text->length, 0);
    return;
  }
  if (node->type != LXB_DOM_NODE_TYPE_ELEMENT)
    return;
  size_t size;
  const unsigned char *name =
      aliases ? gfm_dom_name(node, &size)
              : lxb_dom_element_qualified_name(lxb_dom_interface_element(node), &size);
  h_s(out, "<");
  h_put(out, name, size);
  for (lxb_dom_attr_t *attr = lxb_dom_interface_element(node)->first_attr; attr;
       attr = attr->next) {
    size_t n;
    const unsigned char *key = lxb_dom_attr_qualified_name(attr, &n);
    h_s(out, " ");
    h_put(out, key, n);
    h_s(out, "=\"");
    const unsigned char *value = lxb_dom_attr_value(attr, &n);
    gfm_dom_escape(out, value, n, 1);
    h_s(out, "\"");
  }
  h_s(out, ">");
  if (node->ns == LXB_NS_HTML &&
      h_list("area base br col embed hr img input keygen link meta param source track wbr", name,
             size))
    return;
  int literal = node->ns == LXB_NS_HTML &&
                h_list("style script xmp iframe noembed noframes plaintext noscript", name, size);
  for (lxb_dom_node_t *child = h_children(node); child; child = child->next)
    gfm_dom_serialize_node(out, child, depth + 1, literal, aliases);
  h_s(out, "</");
  h_put(out, name, size);
  h_s(out, ">");
}
static void gfm_dom_serialize(h_buffer *out, lxb_dom_node_t *root, unsigned depth, int aliases) {
  for (lxb_dom_node_t *child = h_children(root); child; child = child->next)
    gfm_dom_serialize_node(out, child, depth + 1, 0, aliases);
}
