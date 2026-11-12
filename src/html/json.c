/* JSON tree adapter for comparison and source mapping. */
static void h_json(h_buffer *b, const unsigned char *s, size_t length) {
  h_s(b, "\"");
  for (size_t i = 0; i < length; i++) {
    unsigned char c = s[i];
    if (c == '"')
      h_s(b, "\\\"");
    else if (c == '\\')
      h_s(b, "\\\\");
    else if (c < 32) {
      char escaped[7];
      snprintf(escaped, sizeof(escaped), "\\u%04x", c);
      h_s(b, escaped);
    } else
      h_put(b, &c, 1);
  }
  h_s(b, "\"");
}
static void h_number(h_buffer *b, size_t n) {
  char s[32];
  snprintf(s, sizeof(s), "%zu", n);
  h_s(b, s);
}

static const char *h_namespace(uintptr_t ns) {
  switch (ns) {
  case LXB_NS_HTML:
    return "http://www.w3.org/1999/xhtml";
  case LXB_NS_SVG:
    return "http://www.w3.org/2000/svg";
  case LXB_NS_MATH:
    return "http://www.w3.org/1998/Math/MathML";
  case LXB_NS_XLINK:
    return "http://www.w3.org/1999/xlink";
  case LXB_NS_XML:
    return "http://www.w3.org/XML/1998/namespace";
  case LXB_NS_XMLNS:
    return "http://www.w3.org/2000/xmlns/";
  default:
    return "";
  }
}
static void h_range(h_buffer *b, h_document *d, size_t start, size_t end) {
  if (start > d->length || end > d->length) {
    b->failed = 1;
    return;
  }
  h_s(b, "{\"startOffset\":");
  h_number(b, d->utf16[start]);
  h_s(b, ",\"endOffset\":");
  h_number(b, d->utf16[end]);
  h_s(b, "}");
}
static void h_json_node(h_buffer *b, h_document *d, lxb_dom_node_t *node, int root,
                        unsigned depth) {
  if (depth > 512) {
    b->failed = H_DEPTH_EXCEEDED;
    return;
  }
  h_s(b, "{\"nodeName\":");
  if (root)
    h_json(b, (const unsigned char *)"#document-fragment", 18);
  else if (node->type == LXB_DOM_NODE_TYPE_TEXT || node->type == LXB_DOM_NODE_TYPE_COMMENT) {
    int text = node->type == LXB_DOM_NODE_TYPE_TEXT;
    h_json(b, (const unsigned char *)(text ? "#text" : "#comment"), text ? 5 : 8);
    h_s(b, text ? ",\"value\":" : ",\"data\":");
    lexbor_str_t *data = &((lxb_dom_character_data_t *)node)->data;
    h_json(b, data->data, data->length);
    h_s(b, "}");
    return;
  } else if (node->type == LXB_DOM_NODE_TYPE_ELEMENT) {
    size_t length;
    const unsigned char *name =
        lxb_dom_element_qualified_name(lxb_dom_interface_element(node), &length);
    h_json(b, name, length);
    h_s(b, ",\"tagName\":");
    h_json(b, name, length);
    h_s(b, ",\"namespaceURI\":");
    h_json(b, (const unsigned char *)h_namespace(node->ns), strlen(h_namespace(node->ns)));
    h_s(b, ",\"attrs\":[");
    int first = 1;
    for (lxb_dom_attr_t *attr = lxb_dom_interface_element(node)->first_attr; attr;
         attr = attr->next) {
      if (!first)
        h_s(b, ",");
      first = 0;
      const unsigned char *attr_name =
          attr->node.ns >= LXB_NS_XLINK && attr->node.ns <= LXB_NS_XMLNS
              ? lxb_dom_attr_local_name(attr, &length)
              : lxb_dom_attr_qualified_name(attr, &length);
      h_s(b, "{\"name\":");
      h_json(b, attr_name, length);
      if (attr->node.ns >= LXB_NS_XLINK && attr->node.ns <= LXB_NS_XMLNS) {
        h_s(b, ",\"namespace\":");
        h_json(b, (const unsigned char *)h_namespace(attr->node.ns),
               strlen(h_namespace(attr->node.ns)));
        const char *prefix = attr->node.ns == LXB_NS_XLINK                     ? "xlink"
                             : attr->node.ns == LXB_NS_XML                     ? "xml"
                             : (length == 5 && !memcmp(attr_name, "xmlns", 5)) ? ""
                                                                               : "xmlns";
        h_s(b, ",\"prefix\":");
        h_json(b, (const unsigned char *)prefix, strlen(prefix));
      }
      const unsigned char *value = lxb_dom_attr_value(attr, &length);
      h_s(b, ",\"value\":");
      h_json(b, value, length);
      h_s(b, "}");
    }
    h_s(b, "]");
    h_location *loc = node->user;
    if (d->track && loc) {
      h_s(b, ",\"sourceCodeLocation\":{\"startOffset\":");
      h_number(b, d->utf16[loc->start]);
      h_s(b, ",\"startTag\":");
      h_range(b, d, loc->start, loc->open_end);
      if (loc->close_start != SIZE_MAX) {
        h_s(b, ",\"endTag\":");
        h_range(b, d, loc->close_start, loc->end);
      }
      h_s(b, ",\"attrs\":{");
      for (size_t i = 0; i < loc->count; i++) {
        if (i)
          h_s(b, ",");
        h_attr_location *a = &loc->attrs[i];
        h_buffer name = {0};
        for (size_t n = 0; n < a->length; n++) {
          unsigned char c = a->name[n];
          if (c >= 'A' && c <= 'Z')
            c += 32;
          h_put(&name, &c, 1);
        }
        if (name.failed)
          b->failed = name.failed;
        h_json(b, name.data, name.length);
        free(name.data);
        h_s(b, ":");
        h_range(b, d, a->start, a->end);
      }
      h_s(b, "}}");
    }
  } else {
    h_json(b, (const unsigned char *)"#unknown", 8);
  }
  int is_template = !root && node->type == LXB_DOM_NODE_TYPE_ELEMENT &&
                    node->local_name == LXB_TAG_TEMPLATE && node->ns == LXB_NS_HTML;
  if (is_template)
    h_s(b, ",\"content\":{\"nodeName\":\"#document-fragment\"");
  h_s(b, ",\"childNodes\":[");
  int first = 1;
  for (lxb_dom_node_t *child = h_children(node); child; child = child->next) {
    if (!first)
      h_s(b, ",");
    first = 0;
    h_json_node(b, d, child, 0, depth + 1);
  }
  h_s(b, "]}");
  if (is_template)
    h_s(b, "}");
}
