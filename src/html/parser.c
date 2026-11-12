typedef enum { H_TRACK_NONE = 0, H_TRACK_UTF16 = 1, H_TRACK_BYTES = 2 } h_tracking;

typedef struct {
  const unsigned char *name;
  size_t length, start, end;
} h_attr_location;
typedef struct {
  size_t start, open_end, close_start, end, count;
  h_attr_location *attrs;
} h_location;
typedef struct {
  lxb_html_parser_t *parser;
  lxb_html_document_t *document;
  lxb_dom_node_t *root;
  const unsigned char *source;
  size_t length, token_end;
  size_t *utf16;
  h_tracking track;
  int failed;
  lxb_html_tokenizer_token_f token_callback;
  void *token_context;
  int (*token_filter)(lxb_html_tokenizer_t *, lxb_html_token_t *, void *);
  void *filter_context;
} h_document;

static size_t h_offset(h_document *d, const unsigned char *pointer) {
  uintptr_t p = (uintptr_t)pointer, start = (uintptr_t)d->source;
  return p >= start && p <= start + d->length ? p - start : SIZE_MAX;
}
static h_location *h_location_new(lxb_html_tree_t *tree) {
  h_location *loc = lexbor_mraw_calloc(tree->document->dom_document.mraw, sizeof(*loc));
  if (loc)
    loc->close_start = SIZE_MAX;
  return loc;
}

/* lexbor.c routes upstream calls here using scoped preprocessor macros.
 * The original implementation creates the node; this wrapper adds locations. */
lxb_html_element_t *lxb_html_tree_create_element_for_token(lxb_html_tree_t *tree,
                                                           lxb_html_token_t *token,
                                                           lxb_ns_id_t ns) {
  lxb_html_element_t *element = gfm_lexbor_create_element_original(tree, token, ns);
  if (!element)
    return NULL;
  h_document *d = tree->document->dom_document.user;
  if (!d || !d->track)
    return element;
  lxb_dom_node_t *node = lxb_dom_interface_node(element);
  if (token->base_element) {
    h_location *old = lxb_dom_interface_node(token->base_element)->user;
    if (old) {
      h_location *loc = h_location_new(tree);
      if (!loc) {
        d->failed = 1;
        return element;
      }
      *loc = *old;
      loc->close_start = SIZE_MAX;
      loc->end = 0;
      node->user = loc;
    }
    return element;
  }
  size_t begin = h_offset(d, token->begin);
  if (begin == SIZE_MAX || begin == 0 || token->type & LXB_HTML_TOKEN_TYPE_CLOSE)
    return element;
  h_location *loc = h_location_new(tree);
  if (!loc) {
    d->failed = 1;
    return element;
  }
  loc->start = begin - 1;
  loc->open_end = d->token_end;
  size_t count = 0;
  for (lxb_html_token_attr_t *a = token->attr_first; a; a = a->next)
    count++;
  if (count) {
    loc->attrs = lexbor_mraw_calloc(tree->document->dom_document.mraw, count * sizeof(*loc->attrs));
    if (!loc->attrs) {
      d->failed = 1;
      return element;
    }
  }
  for (lxb_html_token_attr_t *a = token->attr_first; a; a = a->next) {
    size_t start = h_offset(d, a->name_begin), end = h_offset(d, a->name_end);
    if (start == SIZE_MAX || end == SIZE_MAX)
      continue;
    size_t attr_end = end;
    if (a->value_begin) {
      attr_end = h_offset(d, a->value_end);
      if (attr_end == SIZE_MAX)
        continue;
      size_t value_start = h_offset(d, a->value_begin);
      if (value_start && value_start != SIZE_MAX &&
          (d->source[value_start - 1] == '"' || d->source[value_start - 1] == '\''))
        attr_end++;
    }
    int duplicate = 0;
    for (size_t i = 0; i < loc->count; i++) {
      h_attr_location *old = &loc->attrs[i];
      if (old->length != end - start)
        continue;
      size_t n = 0;
      for (; n < old->length; n++) {
        unsigned char x = old->name[n], y = a->name_begin[n];
        if (x >= 'A' && x <= 'Z')
          x += 32;
        if (y >= 'A' && y <= 'Z')
          y += 32;
        if (x != y)
          break;
      }
      if (n == old->length) {
        duplicate = 1;
        break;
      }
    }
    if (!duplicate)
      loc->attrs[loc->count++] = (h_attr_location){a->name_begin, end - start, start, attr_end};
  }
  node->user = loc;
  return element;
}

static lxb_html_token_t *h_token(lxb_html_tokenizer_t *tkz, lxb_html_token_t *token,
                                 void *context) {
  h_document *d = context;
  if (d->token_filter && !d->token_filter(tkz, token, d->filter_context)) {
    d->failed = 1;
    return NULL;
  }
  if (d->track && token->type & LXB_HTML_TOKEN_TYPE_CLOSE) {
    size_t begin = h_offset(d, token->begin);
    if (begin != SIZE_MAX && begin >= 2) {
      lexbor_array_t *open = d->parser->tree->open_elements;
      for (size_t i = open->length; i-- > 0;) {
        lxb_dom_node_t *node = open->list[i];
        if (node->local_name == token->tag_id) {
          h_location *loc = node->user;
          if (loc) {
            loc->close_start = begin - 2;
            loc->end = d->token_end;
          }
          break;
        }
      }
    }
  }
  return d->token_callback(tkz, token, d->token_context);
}

static void h_destroy(h_document *d) {
  if (d->parser)
    lxb_html_parser_destroy(d->parser);
  if (d->document)
    lxb_html_document_destroy(d->document);
  free(d->utf16);
}
static int h_parse(h_document *d, const unsigned char *source, size_t length, h_tracking track) {
  d->source = source;
  d->length = length;
  d->track = track;
  d->document = lxb_html_document_create();
  d->parser = lxb_html_parser_create();
  if (!d->document || !d->parser || lxb_html_parser_init(d->parser) != LXB_STATUS_OK)
    return 0;
  if (track == H_TRACK_UTF16) {
    if (length > SIZE_MAX / sizeof(size_t) - 1)
      return 0;
    d->utf16 = malloc((length + 1) * sizeof(size_t));
    if (!d->utf16)
      return 0;
    size_t units = 0;
    for (size_t i = 0; i < length; i++) {
      d->utf16[i] = units;
      if ((source[i] & 0xc0) != 0x80)
        units += source[i] >= 0xf0 ? 2 : 1;
    }
    d->utf16[length] = units;
  }
  if (lxb_html_parse_fragment_chunk_begin(d->parser, d->document, LXB_TAG_TEMPLATE, LXB_NS_HTML) !=
      LXB_STATUS_OK)
    return 0;
  d->parser->tree->document->dom_document.user = d;
  d->token_callback = d->parser->tkz->callback_token_done;
  d->token_context = d->parser->tkz->callback_token_ctx;
  lxb_html_tokenizer_callback_token_done_set(d->parser->tkz, h_token, d);
  /* Split tracked input after each '>'. The actual HTML tokenizer decides
   * whether this ends a tag, including quoted '>' and duplicate attributes.
   * Its callback then knows the precise consumed tag boundary, without a
   * second tokenizer or any parser-algorithm changes. */
  if (track) {
    size_t start = 0;
    for (size_t i = 0; i < length; i++)
      if (source[i] == '>') {
        d->token_end = i + 1;
        if (lxb_html_parse_fragment_chunk_process(d->parser, source + start, i + 1 - start) !=
            LXB_STATUS_OK)
          return 0;
        start = i + 1;
      }
    d->token_end = length;
    if (start < length && lxb_html_parse_fragment_chunk_process(d->parser, source + start,
                                                                length - start) != LXB_STATUS_OK)
      return 0;
  } else if (lxb_html_parse_fragment_chunk_process(d->parser, source, length) != LXB_STATUS_OK)
    return 0;
  d->root = lxb_html_parse_fragment_chunk_end(d->parser);
  return d->root != NULL && !d->failed;
}
static lxb_dom_node_t *h_children(lxb_dom_node_t *node) {
  if (node->type == LXB_DOM_NODE_TYPE_ELEMENT && node->ns == LXB_NS_HTML &&
      node->local_name == LXB_TAG_TEMPLATE)
    return lxb_dom_interface_node(lxb_html_interface_template(node)->content)->first_child;
  return node->first_child;
}
