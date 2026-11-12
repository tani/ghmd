/* Structural rules are adapted from happy-dom 21.0.4's HTML element
 * constraints (MIT; Copyright (c) 2019 David Ortner). See DOM-COMPAT-LICENSE.
 * Lexbor parsing, native DOM edits and renderer coordination are implemented
 * here; no happy-dom parser or runtime is included.
 *
 * HTML5 tree adjustments that previously ran in render.ts. Lexbor owns parsing
 * and tree repair; these edits preserve the renderer's compatibility policy. */
static void gfm_dom_unwrap(lxb_dom_node_t *node) {
  gfm_dom_children_before(node);
  lxb_dom_node_remove(node);
}
static int gfm_dom_anchors(lxb_dom_node_t *root, unsigned char *seen, unsigned depth) {
  if (depth > 512)
    return 0;
  lxb_dom_node_t *node = root->first_child;
  while (node) {
    h_location *loc = node->user;
    if (node->type == LXB_DOM_NODE_TYPE_ELEMENT && node->local_name == LXB_TAG_A && loc) {
      if (seen[loc->start]) {
        lxb_dom_node_t *next = node->first_child ? node->first_child : node->next;
        gfm_dom_unwrap(node);
        node = next;
        continue;
      }
      seen[loc->start] = 1;
    }
    if (!gfm_dom_anchors(node, seen, depth + 1))
      return 0;
    node = node->next;
  }
  return 1;
}
static int gfm_dom_has_earlier(lxb_dom_node_t *root, lxb_dom_node_t *tail, h_location *loc,
                               unsigned depth) {
  if (depth > 512 || root == tail)
    return 0;
  h_location *other = root->user;
  if (root->local_name == tail->local_name && other && other->start == loc->start)
    return 1;
  for (lxb_dom_node_t *n = root->first_child; n; n = n->next)
    if (gfm_dom_has_earlier(n, tail, loc, depth + 1))
      return 1;
  return 0;
}
static int gfm_dom_whitespace(lxb_dom_node_t *root) {
  for (lxb_dom_node_t *n = root->first_child; n; n = n->next) {
    if (n->type != LXB_DOM_NODE_TYPE_TEXT)
      return 0;
    lexbor_str_t *text = &((lxb_dom_character_data_t *)n)->data;
    for (size_t i = 0; i < text->length;) {
      uint32_t cp = gfm_utf8_at(text->data, text->length, i);
      if (!gfm_js_whitespace(cp))
        return 0;
      i++;
      while (i < text->length && (text->data[i] & 0xc0) == 0x80)
        i++;
    }
  }
  return 1;
}
static int gfm_dom_svg(lxb_dom_node_t *root, int *changed, unsigned depth) {
  if (depth > 512)
    return 0;
  lxb_dom_node_t *n = h_children(root);
  while (n) {
    if (n->ns == LXB_NS_SVG && n->local_name != LXB_TAG_A) {
      lxb_dom_node_t *next = n->first_child ? n->first_child : n->next;
      gfm_dom_unwrap(n);
      n = next;
      *changed = 1;
      continue;
    }
    if (!gfm_dom_svg(n, changed, depth + 1))
      return 0;
    n = n->next;
  }
  return 1;
}
static int gfm_dom_mark(lxb_dom_node_t *root, const char *marker, unsigned depth) {
  if (depth > 512)
    return 0;
  if (root->type == LXB_DOM_NODE_TYPE_ELEMENT && root->local_name == LXB_TAG_A) {
    lxb_dom_element_t *e = lxb_dom_interface_element(root);
    size_t marker_length = strlen(marker);
    if (lxb_dom_element_has_attribute(e, (const unsigned char *)"href", 4) &&
        !lxb_dom_element_set_attribute(e, (const unsigned char *)"data-gfm-had-href", 17,
                                       (const unsigned char *)marker, marker_length))
      return 0;
  }
  for (lxb_dom_node_t *n = h_children(root); n; n = n->next)
    if (!gfm_dom_mark(n, marker, depth + 1))
      return 0;
  return 1;
}
static void gfm_dom_trim(h_document *d) {
  lxb_dom_node_t *last = d->root->last_child;
  h_location *loc = last ? last->user : NULL;
  if (last && last->type == LXB_DOM_NODE_TYPE_ELEMENT && loc && loc->close_start == SIZE_MAX) {
    lxb_dom_node_t *leaf = last;
    while (h_children(leaf)) {
      leaf = h_children(leaf);
      while (leaf->next)
        leaf = leaf->next;
    }
    if (leaf->type == LXB_DOM_NODE_TYPE_TEXT) {
      lexbor_str_t *text = &((lxb_dom_character_data_t *)leaf)->data;
      if (text->length && text->data[text->length - 1] == '\n')
        text->data[--text->length] = 0;
    }
  }
  last = d->root->last_child;
  loc = last ? last->user : NULL;
  if (last && last->type == LXB_DOM_NODE_TYPE_ELEMENT && loc && gfm_dom_whitespace(last) &&
      gfm_dom_has_earlier(d->root, last, loc, 0))
    lxb_dom_node_remove(last);
}
static char *gfm_dom_reparse(const unsigned char *source, size_t length, gfm_status *status) {
  if (length > GFM_MAX_OUTPUT_BYTES) {
    *status = GFM_LIMIT_EXCEEDED;
    return NULL;
  }
  gfm_dom_tokens masks = {0};
  h_document d = {.token_filter = gfm_dom_filter_token, .filter_context = &masks};
  h_buffer out = {0};
  if (!h_parse(&d, source, length, H_TRACK_NONE))
    out.failed = H_ALLOC_FAILED;
  else if (!gfm_dom_compat(d.root, 0))
    out.failed = H_DEPTH_EXCEEDED;
  else
    gfm_dom_serialize(&out, d.root, 0, 1);
  h_destroy(&d);
  if (out.failed) {
    *status = out.failed == H_ALLOC_FAILED ? GFM_NO_MEMORY : GFM_LIMIT_EXCEEDED;
    free(out.data);
    return NULL;
  }
  if (!out.data)
    out.data = (unsigned char *)calloc(1, 1);
  return (char *)out.data;
}
static char *gfm_dom_sanitize(const unsigned char *source, size_t length, const char *marker,
                              gfm_status *status) {
  h_document d = {0};
  h_buffer out = {0};
  unsigned char *seen = NULL;
  int svg = 0;
  if (!h_parse(&d, source, length, H_TRACK_BYTES))
    goto fail;
  seen = calloc(length + 1, 1);
  if (!seen)
    goto fail;
  if (!gfm_dom_anchors(d.root, seen, 0)) {
    *status = GFM_LIMIT_EXCEEDED;
    goto fail;
  }
  free(seen);
  seen = NULL;
  gfm_dom_trim(&d);
  if (!gfm_dom_svg(d.root, &svg, 0)) {
    *status = GFM_LIMIT_EXCEEDED;
    goto fail;
  }
  if (svg) {
    gfm_dom_serialize(&out, d.root, 0, 0);
    if (out.failed) {
      *status = out.failed == H_ALLOC_FAILED ? GFM_NO_MEMORY : GFM_LIMIT_EXCEEDED;
      goto fail;
    }
    h_document next = {0};
    if (!h_parse(&next, out.data, out.length, H_TRACK_NONE)) {
      h_destroy(&next);
      goto fail;
    }
    h_destroy(&d);
    d = next;
    free(out.data);
    out = (h_buffer){0};
  }
  if (!gfm_dom_mark(d.root, marker, 0))
    goto fail;
  h_sanitize(&out, &d, d.root, marker);
  h_destroy(&d);
  if (out.failed) {
    *status = out.failed == H_ALLOC_FAILED ? GFM_NO_MEMORY : GFM_LIMIT_EXCEEDED;
    free(out.data);
    return NULL;
  }
  if (!out.data)
    out.data = (unsigned char *)calloc(1, 1);
  return (char *)out.data;
fail:
  free(seen);
  h_destroy(&d);
  free(out.data);
  return NULL;
}
