/* Structural rules are adapted from happy-dom 21.0.4's HTML element
 * constraints (MIT; Copyright (c) 2019 David Ortner). See DOM-COMPAT-LICENSE.
 * Lexbor parsing, native DOM edits and renderer coordination are implemented
 * here; no happy-dom parser or runtime is included.
 *
 * HTML5 tree adjustments that previously ran in render.ts. Lexbor owns parsing
 * and tree repair; these edits preserve the renderer's compatibility policy. */
/* Lexbor supplies tokenization and DOM construction. At internal reparse
 * boundaries, neutral tag names retain table/template content until these
 * structural compatibility rules run. This is never used to sanitize raw HTML. */
typedef struct {
  lxb_tag_id_t original;
  const char *neutral;
} gfm_dom_alias;
static const gfm_dom_alias gfm_dom_aliases[] = {
    {LXB_TAG_TEMPLATE, "gfm-native-template"},
    {LXB_TAG_TABLE, "gfm-native-table"},
    {LXB_TAG_THEAD, "gfm-native-thead"},
    {LXB_TAG_TBODY, "gfm-native-tbody"},
    {LXB_TAG_TFOOT, "gfm-native-tfoot"},
    {LXB_TAG_TR, "gfm-native-tr"},
    {LXB_TAG_TD, "gfm-native-td"},
    {LXB_TAG_TH, "gfm-native-th"},
    {LXB_TAG_CAPTION, "gfm-native-caption"},
    {LXB_TAG_COLGROUP, "gfm-native-colgroup"},
    {LXB_TAG_COL, "gfm-native-col"},
    {LXB_TAG_P, "gfm-native-p"},
};
enum { GFM_DOM_ALIAS_COUNT = sizeof(gfm_dom_aliases) / sizeof(gfm_dom_aliases[0]) };
typedef struct {
  lxb_tag_id_t ids[GFM_DOM_ALIAS_COUNT];
} gfm_dom_tokens;
static int gfm_dom_filter_token(lxb_html_tokenizer_t *tkz, lxb_html_token_t *token, void *context) {
  gfm_dom_tokens *masks = context;
  for (size_t i = 0; i < GFM_DOM_ALIAS_COUNT; i++)
    if (token->tag_id == gfm_dom_aliases[i].original) {
      if (!masks->ids[i]) {
        size_t length = strlen(gfm_dom_aliases[i].neutral);
        const lxb_tag_data_t *name = lxb_tag_append_lower(
            tkz->tags, (const unsigned char *)gfm_dom_aliases[i].neutral, length);
        if (!name)
          return 0;
        masks->ids[i] = name->tag_id;
      }
      token->tag_id = masks->ids[i];
      break;
    }
  return 1;
}
static const unsigned char *gfm_dom_name(lxb_dom_node_t *node, size_t *length) {
  const unsigned char *name =
      lxb_dom_element_qualified_name(lxb_dom_interface_element(node), length);
  const size_t prefix = sizeof("gfm-native-") - 1;
  if (*length <= prefix || memcmp(name, "gfm-native-", prefix))
    return name;
  for (size_t i = 0; i < GFM_DOM_ALIAS_COUNT; i++) {
    if (h_equal(name, *length, gfm_dom_aliases[i].neutral)) {
      name += prefix;
      *length -= prefix;
      break;
    }
  }
  return name;
}
static int gfm_dom_is(lxb_dom_node_t *node, const char *name) {
  if (node->type != LXB_DOM_NODE_TYPE_ELEMENT)
    return 0;
  size_t length;
  const unsigned char *tag = gfm_dom_name(node, &length);
  return h_equal(tag, length, name);
}
static int gfm_dom_in(lxb_dom_node_t *node, const char *names) {
  if (node->type != LXB_DOM_NODE_TYPE_ELEMENT)
    return 0;
  size_t length;
  const unsigned char *tag = gfm_dom_name(node, &length);
  return h_list(names, tag, length);
}
static const char *gfm_dom_permitted(lxb_dom_node_t *parent) {
  if (gfm_dom_is(parent, "table"))
    return "caption colgroup thead tfoot tbody";
  if (gfm_dom_in(parent, "thead tbody tfoot"))
    return "tr";
  if (gfm_dom_is(parent, "tr"))
    return "td th";
  if (gfm_dom_is(parent, "colgroup"))
    return "col";
  return NULL;
}
static const char *gfm_dom_required(lxb_dom_node_t *node) {
  if (gfm_dom_in(node, "thead tbody tfoot caption"))
    return "table";
  if (gfm_dom_is(node, "tr"))
    return "thead tbody tfoot";
  if (gfm_dom_in(node, "td th"))
    return "tr";
  if (gfm_dom_is(node, "col"))
    return "colgroup";
  return NULL;
}
static int gfm_dom_close_parent(lxb_dom_node_t *parent, lxb_dom_node_t *node) {
  if (gfm_dom_is(parent, "p"))
    return gfm_dom_in(
        node,
        "address article aside blockquote details dialog div dl fieldset figcaption figure footer form h1 h2 h3 h4 h5 h6 header hgroup hr main menu nav ol p pre section table ul");
  if (gfm_dom_in(parent, "td th"))
    return gfm_dom_in(node, "td th tr tbody tfoot thead");
  if (gfm_dom_is(parent, "caption"))
    return gfm_dom_in(node, "table tbody thead tfoot tr td th col colgroup");
  if (gfm_dom_is(parent, "option"))
    return gfm_dom_in(node, "option optgroup");
  return 0;
}
static void gfm_dom_children_before(lxb_dom_node_t *node) {
  while (node->first_child) {
    lxb_dom_node_t *child = node->first_child;
    lxb_dom_node_remove(child);
    lxb_dom_node_insert_before(node, child);
  }
}
/* Move a complete sibling tail, preserving order and returning its new end. */
static lxb_dom_node_t *gfm_dom_move_after(lxb_dom_node_t *after, lxb_dom_node_t *first) {
  while (first) {
    lxb_dom_node_t *next = first->next;
    lxb_dom_node_remove(first);
    lxb_dom_node_insert_after(after, first);
    after = first;
    first = next;
  }
  return after;
}
/* A relocated start tag never entered the old open-element stack. Its end
 * tag could therefore close an existing ancestor with the same name. */
static lxb_dom_node_t *gfm_dom_same_ancestor(lxb_dom_node_t *root, lxb_dom_node_t *node) {
  size_t size;
  const unsigned char *name = gfm_dom_name(node, &size);
  for (lxb_dom_node_t *parent = root; parent; parent = parent->parent) {
    if (parent->type != LXB_DOM_NODE_TYPE_ELEMENT)
      continue;
    size_t length;
    const unsigned char *tag = gfm_dom_name(parent, &length);
    if (size == length && !memcmp(name, tag, size))
      return parent;
  }
  return NULL;
}
static void gfm_dom_close_relocated(lxb_dom_node_t *root, lxb_dom_node_t *tail,
                                    lxb_dom_node_t *ancestor) {
  lxb_dom_node_t *after = ancestor;
  while (root) {
    lxb_dom_node_t *next_root = root->parent, *next_tail = root->next;
    after = gfm_dom_move_after(after, tail);
    if (root == ancestor)
      break;
    root = next_root;
    tail = next_tail;
  }
}
static int gfm_dom_compat(lxb_dom_node_t *root, unsigned depth) {
  if (depth > 512)
    return 0;
  lxb_dom_node_t *node = root->first_child;
  while (node) {
    if (node->type != LXB_DOM_NODE_TYPE_ELEMENT) {
      node = node->next;
      continue;
    }
    /* A neutralized col still has HTML void-element semantics. */
    if (gfm_dom_is(node, "col")) {
      gfm_dom_move_after(node, node->first_child);
    }
    if (gfm_dom_close_parent(root, node) && root->parent) {
      /* A start tag closes only a direct forbidden parent. Everything after
       * it follows that parent, rather than adding a synthetic empty p. */
      gfm_dom_move_after(root, node);
      return 1;
    }
    int tbody = gfm_dom_is(root, "table") && gfm_dom_is(node, "tr");
    const char *permitted = gfm_dom_permitted(root);
    if (permitted && !gfm_dom_in(node, permitted) && !tbody) {
      int close =
          (gfm_dom_in(root, "thead tbody tfoot") &&
           gfm_dom_in(node, "caption colgroup thead tfoot tbody")) ||
          (gfm_dom_is(root, "tr") && gfm_dom_in(node, "caption colgroup thead tfoot tbody tr"));
      if (close && root->parent) {
        gfm_dom_move_after(root, node);
        return 1;
      }
      /* Forbidden elements are placed before the table; their contents stay
       * at the current insertion point. Text itself is never foster-parented. */
      int had_children = node->first_child != NULL;
      lxb_dom_node_t *tail = node->next, *closes = gfm_dom_same_ancestor(root, node);
      lxb_dom_node_t *next = had_children ? node->first_child : node->next;
      gfm_dom_children_before(node);
      if (closes)
        gfm_dom_close_relocated(root, tail, closes);
      lxb_dom_node_t *before = root;
      while (before->parent) {
        const char *allowed = gfm_dom_permitted(before->parent);
        if (!allowed || gfm_dom_in(node, allowed))
          break;
        before = before->parent;
      }
      lxb_dom_node_remove(node);
      if (before->parent)
        lxb_dom_node_insert_before(before, node);
      else
        lxb_dom_node_insert_child(before, node);
      if (closes && !had_children)
        return 1;
      node = next;
      continue;
    }
    if (tbody) {
      const char *name = "gfm-native-tbody";
      lxb_dom_element_t *body = lxb_dom_document_create_element(
          node->owner_document, (const unsigned char *)name, 16, NULL);
      if (!body)
        return 0;
      lxb_dom_node_t *body_node = lxb_dom_interface_node(body);
      lxb_dom_node_insert_before(node, body_node);
      lxb_dom_node_t *child = node;
      while (child && gfm_dom_is(child, "tr")) {
        lxb_dom_node_t *next = child->next;
        lxb_dom_node_remove(child);
        lxb_dom_node_insert_child(body_node, child);
        child = next;
      }
      node = body_node;
    }
    const char *parents = gfm_dom_required(node);
    if (parents && !gfm_dom_in(root, parents)) {
      lxb_dom_node_t *next = node->first_child ? node->first_child : node->next;
      gfm_dom_children_before(node);
      lxb_dom_node_remove(node);
      node = next;
      continue;
    }
    if (!gfm_dom_compat(node, depth + 1))
      return 0;
    node = node->next;
  }
  return 1;
}
