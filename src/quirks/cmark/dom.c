/* Structural rules are adapted from happy-dom 21.0.4's HTML element
 * constraints (MIT; Copyright (c) 2019 David Ortner). See DOM-COMPAT-LICENSE.
 * Lexbor parsing, native DOM edits and renderer coordination are implemented
 * here; no happy-dom parser or runtime is included.
 *
 * HTML5 tree adjustments that previously ran in render.ts. Lexbor owns parsing
 * and tree repair; these edits preserve the renderer's compatibility policy. */
static void gfm_dom_unwrap(lxb_dom_node_t *node) {
  while (node->first_child) {
    lxb_dom_node_t *child = node->first_child;
    lxb_dom_node_remove(child); lxb_dom_node_insert_before(node, child);
  }
  lxb_dom_node_remove(node);
}
static int gfm_dom_anchors(lxb_dom_node_t *root, unsigned char *seen, unsigned depth) {
  if (depth > 512) return 0;
  lxb_dom_node_t *node = root->first_child;
  while (node) {
    h_location *loc = node->user;
    if (node->type == LXB_DOM_NODE_TYPE_ELEMENT && node->local_name == LXB_TAG_A && loc) {
      if (seen[loc->start]) {
        lxb_dom_node_t *next = node->first_child ? node->first_child : node->next;
        gfm_dom_unwrap(node); node = next; continue;
      }
      seen[loc->start] = 1;
    }
    if (!gfm_dom_anchors(node, seen, depth + 1)) return 0;
    node = node->next;
  }
  return 1;
}
static int gfm_dom_has_earlier(lxb_dom_node_t *root, lxb_dom_node_t *tail, h_location *loc, unsigned depth) {
  if (depth > 512 || root == tail) return 0;
  h_location *other = root->user;
  if (root->local_name == tail->local_name && other && other->start == loc->start) return 1;
  for (lxb_dom_node_t *n = root->first_child; n; n = n->next)
    if (gfm_dom_has_earlier(n, tail, loc, depth + 1)) return 1;
  return 0;
}
static int gfm_dom_whitespace(lxb_dom_node_t *root) {
  for (lxb_dom_node_t *n = root->first_child; n; n = n->next) {
    if (n->type != LXB_DOM_NODE_TYPE_TEXT) return 0;
    lexbor_str_t *text = &((lxb_dom_character_data_t *)n)->data;
    for (size_t i = 0; i < text->length;) {
      uint32_t cp = gfm_utf8_at(text->data, text->length, i);
      if (!gfm_js_whitespace(cp)) return 0;
      i++; while (i < text->length && (text->data[i] & 0xc0) == 0x80) i++;
    }
  }
  return 1;
}
static int gfm_dom_svg(lxb_dom_node_t *root, int *changed, unsigned depth) {
  if (depth > 512) return 0;
  lxb_dom_node_t *n = h_children(root);
  while (n) {
    if (n->ns == LXB_NS_SVG && n->local_name != LXB_TAG_A) {
      lxb_dom_node_t *next = n->first_child ? n->first_child : n->next;
      gfm_dom_unwrap(n); n = next; *changed = 1; continue;
    }
    if (!gfm_dom_svg(n, changed, depth + 1)) return 0;
    n = n->next;
  }
  return 1;
}
static int gfm_dom_mark(lxb_dom_node_t *root, const char *marker, unsigned depth) {
  if (depth > 512) return 0;
  if (root->type == LXB_DOM_NODE_TYPE_ELEMENT && root->local_name == LXB_TAG_A) {
    lxb_dom_element_t *e = lxb_dom_interface_element(root);
    size_t marker_length = strlen(marker);
    if (lxb_dom_element_has_attribute(e, (const unsigned char *)"href", 4) &&
        !lxb_dom_element_set_attribute(e, (const unsigned char *)"data-gfm-had-href", 17, (const unsigned char *)marker, marker_length)) return 0;
  }
  for (lxb_dom_node_t *n = h_children(root); n; n = n->next)
    if (!gfm_dom_mark(n, marker, depth + 1)) return 0;
  return 1;
}
static void gfm_dom_trim(h_document *d) {
  lxb_dom_node_t *last = d->root->last_child;
  h_location *loc = last ? last->user : NULL;
  if (last && last->type == LXB_DOM_NODE_TYPE_ELEMENT && loc && loc->close_start == SIZE_MAX) {
    lxb_dom_node_t *leaf = last;
    while (h_children(leaf)) {
      leaf = h_children(leaf); while (leaf->next) leaf = leaf->next;
    }
    if (leaf->type == LXB_DOM_NODE_TYPE_TEXT) {
      lexbor_str_t *text = &((lxb_dom_character_data_t *)leaf)->data;
      if (text->length && text->data[text->length - 1] == '\n') text->data[--text->length] = 0;
    }
  }
  last = d->root->last_child; loc = last ? last->user : NULL;
  if (last && last->type == LXB_DOM_NODE_TYPE_ELEMENT && loc && gfm_dom_whitespace(last) && gfm_dom_has_earlier(d->root, last, loc, 0))
    lxb_dom_node_remove(last);
}
/* Lexbor supplies tokenization and DOM construction. At internal reparse
 * boundaries, neutral tag names retain table/template content until these
 * structural compatibility rules run. This is never used to sanitize raw HTML. */
typedef struct { lxb_tag_id_t ids[12]; } gfm_dom_tokens;
static const lxb_tag_id_t gfm_dom_mask_tags[] = {LXB_TAG_TEMPLATE, LXB_TAG_TABLE, LXB_TAG_THEAD, LXB_TAG_TBODY, LXB_TAG_TFOOT, LXB_TAG_TR, LXB_TAG_TD, LXB_TAG_TH, LXB_TAG_CAPTION, LXB_TAG_COLGROUP, LXB_TAG_COL, LXB_TAG_P};
static const char *gfm_dom_mask_names[] = {"gfm-native-template", "gfm-native-table", "gfm-native-thead", "gfm-native-tbody", "gfm-native-tfoot", "gfm-native-tr", "gfm-native-td", "gfm-native-th", "gfm-native-caption", "gfm-native-colgroup", "gfm-native-col", "gfm-native-p"};
static int gfm_dom_filter_token(lxb_html_tokenizer_t *tkz, lxb_html_token_t *token, void *context) {
  gfm_dom_tokens *masks = context;
  for (size_t i = 0; i < 12; i++) if (token->tag_id == gfm_dom_mask_tags[i]) {
    if (!masks->ids[i]) {
      size_t length = strlen(gfm_dom_mask_names[i]);
      const lxb_tag_data_t *name = lxb_tag_append_lower(tkz->tags, (const unsigned char *)gfm_dom_mask_names[i], length);
      if (!name) return 0;
      masks->ids[i] = name->tag_id;
    }
    token->tag_id = masks->ids[i]; break;
  }
  return 1;
}
static const unsigned char *gfm_dom_name(lxb_dom_node_t *node, size_t *length) {
  const unsigned char *name = lxb_dom_element_qualified_name(lxb_dom_interface_element(node), length);
  if (*length > 11 && !memcmp(name, "gfm-native-", 11) && h_list("template table thead tbody tfoot tr td th caption colgroup col p", name + 11, *length - 11)) {
    name += 11; *length -= 11;
  }
  return name;
}
static int gfm_dom_is(lxb_dom_node_t *node, const char *name) {
  if (node->type != LXB_DOM_NODE_TYPE_ELEMENT) return 0;
  size_t length; const unsigned char *tag = gfm_dom_name(node, &length);
  return h_equal(tag, length, name);
}
static int gfm_dom_in(lxb_dom_node_t *node, const char *names) {
  if (node->type != LXB_DOM_NODE_TYPE_ELEMENT) return 0;
  size_t length; const unsigned char *tag = gfm_dom_name(node, &length);
  return h_list(names, tag, length);
}
static const char *gfm_dom_permitted(lxb_dom_node_t *parent) {
  if (gfm_dom_is(parent, "table")) return "caption colgroup thead tfoot tbody";
  if (gfm_dom_in(parent, "thead tbody tfoot")) return "tr";
  if (gfm_dom_is(parent, "tr")) return "td th";
  if (gfm_dom_is(parent, "colgroup")) return "col";
  return NULL;
}
static const char *gfm_dom_required(lxb_dom_node_t *node) {
  if (gfm_dom_in(node, "thead tbody tfoot caption")) return "table";
  if (gfm_dom_is(node, "tr")) return "thead tbody tfoot";
  if (gfm_dom_in(node, "td th")) return "tr";
  if (gfm_dom_is(node, "col")) return "colgroup";
  return NULL;
}
static int gfm_dom_close_parent(lxb_dom_node_t *parent, lxb_dom_node_t *node) {
  if (gfm_dom_is(parent, "p")) return gfm_dom_in(node, "address article aside blockquote details dialog div dl fieldset figcaption figure footer form h1 h2 h3 h4 h5 h6 header hgroup hr main menu nav ol p pre section table ul");
  if (gfm_dom_in(parent, "td th")) return gfm_dom_in(node, "td th tr tbody tfoot thead");
  if (gfm_dom_is(parent, "caption")) return gfm_dom_in(node, "table tbody thead tfoot tr td th col colgroup");
  if (gfm_dom_is(parent, "option")) return gfm_dom_in(node, "option optgroup");
  return 0;
}
static void gfm_dom_children_before(lxb_dom_node_t *node) {
  while (node->first_child) {
    lxb_dom_node_t *child = node->first_child;
    lxb_dom_node_remove(child); lxb_dom_node_insert_before(node, child);
  }
}
/* A relocated start tag never entered the old open-element stack. Its end
 * tag could therefore close an existing ancestor with the same name. */
static lxb_dom_node_t *gfm_dom_same_ancestor(lxb_dom_node_t *root, lxb_dom_node_t *node) {
  size_t size; const unsigned char *name = gfm_dom_name(node, &size);
  for (lxb_dom_node_t *parent = root; parent; parent = parent->parent) {
    if (parent->type != LXB_DOM_NODE_TYPE_ELEMENT) continue;
    size_t length; const unsigned char *tag = gfm_dom_name(parent, &length);
    if (size == length && !memcmp(name, tag, size)) return parent;
  }
  return NULL;
}
static void gfm_dom_close_relocated(lxb_dom_node_t *root, lxb_dom_node_t *tail, lxb_dom_node_t *ancestor) {
  lxb_dom_node_t *after = ancestor;
  while (root) {
    lxb_dom_node_t *next_root = root->parent, *next_tail = root->next;
    while (tail) {
      lxb_dom_node_t *next = tail->next;
      lxb_dom_node_remove(tail); lxb_dom_node_insert_after(after, tail); after = tail; tail = next;
    }
    if (root == ancestor) break;
    root = next_root; tail = next_tail;
  }
}
static int gfm_dom_compat(lxb_dom_node_t *root, unsigned depth) {
  if (depth > 512) return 0;
  lxb_dom_node_t *node = root->first_child;
  while (node) {
    if (node->type != LXB_DOM_NODE_TYPE_ELEMENT) { node = node->next; continue; }
    /* A neutralized col still has HTML void-element semantics. */
    if (gfm_dom_is(node, "col")) {
      lxb_dom_node_t *after = node;
      while (node->first_child) {
        lxb_dom_node_t *child = node->first_child;
        lxb_dom_node_remove(child); lxb_dom_node_insert_after(after, child); after = child;
      }
    }
    if (gfm_dom_close_parent(root, node) && root->parent) {
      /* A start tag closes only a direct forbidden parent. Everything after
       * it follows that parent, rather than adding a synthetic empty p. */
      lxb_dom_node_t *after = root, *moved = node;
      while (moved) {
        lxb_dom_node_t *next = moved->next;
        lxb_dom_node_remove(moved); lxb_dom_node_insert_after(after, moved); after = moved; moved = next;
      }
      return 1;
    }
    int tbody = gfm_dom_is(root, "table") && gfm_dom_is(node, "tr");
    const char *permitted = gfm_dom_permitted(root);
    if (permitted && !gfm_dom_in(node, permitted) && !tbody) {
      int close = (gfm_dom_in(root, "thead tbody tfoot") && gfm_dom_in(node, "caption colgroup thead tfoot tbody")) ||
                  (gfm_dom_is(root, "tr") && gfm_dom_in(node, "caption colgroup thead tfoot tbody tr"));
      if (close && root->parent) {
        lxb_dom_node_t *after = root, *moved = node;
        while (moved) { lxb_dom_node_t *next = moved->next; lxb_dom_node_remove(moved); lxb_dom_node_insert_after(after, moved); after = moved; moved = next; }
        return 1;
      }
      /* Forbidden elements are placed before the table; their contents stay
       * at the current insertion point. Text itself is never foster-parented. */
      int had_children = node->first_child != NULL;
      lxb_dom_node_t *tail = node->next, *closes = gfm_dom_same_ancestor(root, node);
      lxb_dom_node_t *next = had_children ? node->first_child : node->next;
      gfm_dom_children_before(node);
      if (closes) gfm_dom_close_relocated(root, tail, closes);
      lxb_dom_node_t *before = root;
      while (before->parent) {
        const char *allowed = gfm_dom_permitted(before->parent);
        if (!allowed || gfm_dom_in(node, allowed)) break;
        before = before->parent;
      }
      lxb_dom_node_remove(node);
      if (before->parent) lxb_dom_node_insert_before(before, node);
      else lxb_dom_node_insert_child(before, node);
      if (closes && !had_children) return 1;
      node = next; continue;
    }
    if (tbody) {
      const char *name = "gfm-native-tbody";
      lxb_dom_element_t *body = lxb_dom_document_create_element(node->owner_document, (const unsigned char *)name, 16, NULL);
      if (!body) return 0;
      lxb_dom_node_t *body_node = lxb_dom_interface_node(body);
      lxb_dom_node_insert_before(node, body_node);
      lxb_dom_node_t *child = node;
      while (child && gfm_dom_is(child, "tr")) { lxb_dom_node_t *next = child->next; lxb_dom_node_remove(child); lxb_dom_node_insert_child(body_node, child); child = next; }
      node = body_node;
    }
    const char *parents = gfm_dom_required(node);
    if (parents && !gfm_dom_in(root, parents)) {
      lxb_dom_node_t *next = node->first_child ? node->first_child : node->next;
      gfm_dom_children_before(node); lxb_dom_node_remove(node); node = next; continue;
    }
    if (!gfm_dom_compat(node, depth + 1)) return 0;
    node = node->next;
  }
  return 1;
}
/* Serialize DOM nodes with the existing renderer's entity spelling. Attribute
 * angle brackets stay literal inside quotes, as in the former DOM serializer. */
static void gfm_dom_escape(h_buffer *out, const unsigned char *s, size_t length, int attribute) {
  for (size_t i = 0; i < length; i++) {
    if (s[i] == '&') h_s(out, "&amp;");
    else if (s[i] == '"' && attribute) h_s(out, "&quot;");
    else if (s[i] == '<' && !attribute) h_s(out, "&lt;");
    else if (s[i] == '>' && !attribute) h_s(out, "&gt;");
    else if (s[i] == 0xc2 && i + 1 < length && s[i + 1] == 0xa0) { h_s(out, "&nbsp;"); i++; }
    else h_put(out, s + i, 1);
  }
}
static void gfm_dom_serialize_node(h_buffer *out, lxb_dom_node_t *node, unsigned depth, int raw, int aliases) {
  if (out->failed) return;
  if (depth > 512) { out->failed = H_DEPTH_EXCEEDED; return; }
  if (node->type == LXB_DOM_NODE_TYPE_TEXT || node->type == LXB_DOM_NODE_TYPE_COMMENT) {
    lexbor_str_t *text = &((lxb_dom_character_data_t *)node)->data;
    if (node->type == LXB_DOM_NODE_TYPE_COMMENT) { h_s(out, "<!--"); h_put(out, text->data, text->length); h_s(out, "-->"); }
    else if (raw) h_put(out, text->data, text->length);
    else gfm_dom_escape(out, text->data, text->length, 0);
    return;
  }
  if (node->type != LXB_DOM_NODE_TYPE_ELEMENT) return;
  size_t size; const unsigned char *name = aliases ? gfm_dom_name(node, &size) : lxb_dom_element_qualified_name(lxb_dom_interface_element(node), &size);
  h_s(out, "<"); h_put(out, name, size);
  for (lxb_dom_attr_t *attr = lxb_dom_interface_element(node)->first_attr; attr; attr = attr->next) {
    size_t n; const unsigned char *key = lxb_dom_attr_qualified_name(attr, &n);
    h_s(out, " "); h_put(out, key, n); h_s(out, "=\"");
    const unsigned char *value = lxb_dom_attr_value(attr, &n);
    gfm_dom_escape(out, value, n, 1); h_s(out, "\"");
  }
  h_s(out, ">");
  if (node->ns == LXB_NS_HTML && h_list("area base br col embed hr img input keygen link meta param source track wbr", name, size)) return;
  int literal = node->ns == LXB_NS_HTML && h_list("style script xmp iframe noembed noframes plaintext noscript", name, size);
  for (lxb_dom_node_t *child = h_children(node); child; child = child->next) gfm_dom_serialize_node(out, child, depth + 1, literal, aliases);
  h_s(out, "</"); h_put(out, name, size); h_s(out, ">");
}
static void gfm_dom_serialize(h_buffer *out, lxb_dom_node_t *root, unsigned depth, int aliases) {
  for (lxb_dom_node_t *child = h_children(root); child; child = child->next) gfm_dom_serialize_node(out, child, depth + 1, 0, aliases);
}
static char *gfm_dom_reparse(const unsigned char *source, size_t length, gfm_status *status) {
  if (length > GFM_MAX_OUTPUT_BYTES) { *status = GFM_LIMIT_EXCEEDED; return NULL; }
  gfm_dom_tokens masks = {0};
  h_document d = {.token_filter=gfm_dom_filter_token, .filter_context=&masks}; h_buffer out = {0};
  if (!h_parse(&d, source, length, 0)) out.failed = H_ALLOC_FAILED;
  else if (!gfm_dom_compat(d.root, 0)) out.failed = H_DEPTH_EXCEEDED;
  else gfm_dom_serialize(&out, d.root, 0, 1);
  h_destroy(&d);
  if (out.failed) { *status = out.failed == H_ALLOC_FAILED ? GFM_NO_MEMORY : GFM_LIMIT_EXCEEDED; free(out.data); return NULL; }
  if (!out.data) out.data = (unsigned char *)calloc(1, 1);
  return (char *)out.data;
}
static char *gfm_dom_sanitize(const unsigned char *source, size_t length, const char *marker, gfm_status *status) {
  h_document d = {0}; h_buffer out = {0}; unsigned char *seen = NULL; int svg = 0;
  if (!h_parse(&d, source, length, 2)) goto fail;
  seen = calloc(length + 1, 1);
  if (!seen) goto fail;
  if (!gfm_dom_anchors(d.root, seen, 0)) { *status = GFM_LIMIT_EXCEEDED; goto fail; }
  free(seen); seen = NULL;
  gfm_dom_trim(&d);
  if (!gfm_dom_svg(d.root, &svg, 0)) { *status = GFM_LIMIT_EXCEEDED; goto fail; }
  if (svg) {
    gfm_dom_serialize(&out, d.root, 0, 0);
    if (out.failed) { *status = out.failed == H_ALLOC_FAILED ? GFM_NO_MEMORY : GFM_LIMIT_EXCEEDED; goto fail; }
    h_document next = {0};
    if (!h_parse(&next, out.data, out.length, 0)) { h_destroy(&next); goto fail; }
    h_destroy(&d); d = next;
    free(out.data); out = (h_buffer){0};
  }
  if (!gfm_dom_mark(d.root, marker, 0)) goto fail;
  h_sanitize(&out, &d, d.root, marker);
  h_destroy(&d);
  if (out.failed) { *status = out.failed == H_ALLOC_FAILED ? GFM_NO_MEMORY : GFM_LIMIT_EXCEEDED; free(out.data); return NULL; }
  if (!out.data) out.data = (unsigned char *)calloc(1, 1);
  return (char *)out.data;
fail:
  free(seen); h_destroy(&d); free(out.data); return NULL;
}
/* Raw cmark HTML literals are filtered before rendering: remove only actual
 * source-authored input tags, preserving the other source bytes. */
static int gfm_dom_has_input(const unsigned char *s, size_t length) {
  for (size_t i = 0; i + 6 <= length; i++)
    if (s[i] == '<' && gfm_ascii_prefix_case(s + i + 1, length - i - 1, "input") &&
        (i + 6 == length || gfm_ascii_space(s[i + 6]) || s[i + 6] == '/' || s[i + 6] == '>')) return 1;
  return 0;
}
typedef struct { size_t start, end; } gfm_raw_span;
static int gfm_raw_span_compare(const void *a, const void *b) {
  const gfm_raw_span *x = a, *y = b;
  return x->start < y->start ? -1 : x->start > y->start ? 1 : 0;
}
static gfm_status gfm_dom_input_spans(lxb_dom_node_t *node, gfm_raw_span **spans, size_t *count, size_t *capacity, unsigned depth) {
  if (depth > 512) return GFM_LIMIT_EXCEEDED;
  if (node->local_name == LXB_TAG_INPUT && node->user) {
    h_location *loc = node->user;
    if (!gfm_grow((void **)spans, capacity, *count + 1, sizeof(**spans))) return GFM_NO_MEMORY;
    (*spans)[(*count)++] = (gfm_raw_span){loc->start, loc->open_end};
  }
  for (lxb_dom_node_t *child = h_children(node); child; child = child->next) {
    gfm_status status = gfm_dom_input_spans(child, spans, count, capacity, depth + 1);
    if (status != GFM_OK) return status;
  }
  return GFM_OK;
}
static gfm_status gfm_filter_raw_inputs(gfm_view source, gfm_string *result, void *context) {
  (void)context;
  gfm_status status = GFM_NO_MEMORY;
  h_buffer out = {0}; h_document d = {0}; gfm_raw_span *spans = NULL; size_t count = 0, capacity = 0;
  if (gfm_dom_has_input(source.data, source.length)) {
    if (!h_parse(&d, source.data, source.length, 2)) goto fail;
    status = gfm_dom_input_spans(d.root, &spans, &count, &capacity, 0);
    if (status != GFM_OK) goto fail;
  }
  if (count > 1) qsort(spans, count, sizeof(*spans), gfm_raw_span_compare);
  size_t start = 0;
  for (size_t i = 0; i < count; i++) {
    if (spans[i].start < start || spans[i].end > source.length) continue;
    h_put(&out, source.data + start, spans[i].start - start); start = spans[i].end;
  }
  h_put(&out, source.data + start, source.length - start);
  h_destroy(&d); free(spans);
  if (out.failed) { free(out.data); return GFM_NO_MEMORY; }
  *result = (gfm_string){out.data, out.length}; return GFM_OK;
fail:
  h_destroy(&d); free(spans); free(out.data); return status;
}
