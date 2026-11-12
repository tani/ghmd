/* Runtime-independent Lexbor parsing, DOM metadata and sanitizer. */
#include "../../vendor/lexbor/lexbor.h"

#define H_ALLOC_FAILED 2
#define H_DEPTH_EXCEEDED 3
#define H_LIMIT ((size_t)64 * 1024 * 1024)
typedef struct { unsigned char *data; size_t length, capacity; int failed; } h_buffer;
static void h_put(h_buffer *b, const void *data, size_t length) {
  if (b->failed || length > H_LIMIT - b->length) { b->failed = 1; return; }
  size_t need = b->length + length + 1;
  if (need > b->capacity) {
    size_t capacity = b->capacity ? b->capacity : 256;
    while (capacity < need) capacity *= 2;
    unsigned char *next = realloc(b->data, capacity);
    if (!next) { b->failed = H_ALLOC_FAILED; return; }
    b->data = next; b->capacity = capacity;
  }
  if (length) memcpy(b->data + b->length, data, length);
  b->length += length; b->data[b->length] = 0;
}
static void h_s(h_buffer *b, const char *s) { h_put(b, s, strlen(s)); }
static void h_json(h_buffer *b, const unsigned char *s, size_t length) {
  h_s(b, "\"");
  for (size_t i = 0; i < length; i++) {
    unsigned char c = s[i];
    if (c == '"') h_s(b, "\\\"");
    else if (c == '\\') h_s(b, "\\\\");
    else if (c < 32) { char escaped[7]; snprintf(escaped, sizeof(escaped), "\\u%04x", c); h_s(b, escaped); }
    else h_put(b, &c, 1);
  }
  h_s(b, "\"");
}
static void h_number(h_buffer *b, size_t n) { char s[32]; snprintf(s, sizeof(s), "%zu", n); h_s(b, s); }

typedef struct { const unsigned char *name; size_t length, start, end; } h_attr_location;
typedef struct { size_t start, open_end, close_start, end, count; h_attr_location *attrs; } h_location;
typedef struct {
  lxb_html_parser_t *parser; lxb_html_document_t *document;
  lxb_dom_node_t *root; const unsigned char *source; size_t length, token_end;
  size_t *utf16; int track, failed;
  lxb_html_tokenizer_token_f token_callback; void *token_context;
  int (*token_filter)(lxb_html_tokenizer_t *, lxb_html_token_t *, void *); void *filter_context;
} h_document;

static size_t h_offset(h_document *d, const unsigned char *pointer) {
  uintptr_t p = (uintptr_t)pointer, start = (uintptr_t)d->source;
  return p >= start && p <= start + d->length ? p - start : SIZE_MAX;
}
static h_location *h_location_new(lxb_html_tree_t *tree) {
  h_location *loc = lexbor_mraw_calloc(tree->document->dom_document.mraw, sizeof(*loc));
  if (loc) loc->close_start = SIZE_MAX;
  return loc;
}

/* The generated upstream definition is renamed, so all original call sites
 * pass here. Parsing and tree-repair algorithms remain untouched. */
lxb_html_element_t *lxb_html_tree_create_element_for_token(lxb_html_tree_t *tree, lxb_html_token_t *token, lxb_ns_id_t ns) {
  lxb_html_element_t *element = gfm_lexbor_create_element_original(tree, token, ns);
  if (!element) return NULL;
  h_document *d = tree->document->dom_document.user;
  if (!d || !d->track) return element;
  lxb_dom_node_t *node = lxb_dom_interface_node(element);
  if (token->base_element) {
    h_location *old = lxb_dom_interface_node(token->base_element)->user;
    if (old) {
      h_location *loc = h_location_new(tree);
      if (!loc) { d->failed = 1; return element; }
      *loc = *old; loc->close_start = SIZE_MAX; loc->end = 0; node->user = loc;
    }
    return element;
  }
  size_t begin = h_offset(d, token->begin);
  if (begin == SIZE_MAX || begin == 0 || token->type & LXB_HTML_TOKEN_TYPE_CLOSE) return element;
  h_location *loc = h_location_new(tree);
  if (!loc) { d->failed = 1; return element; }
  loc->start = begin - 1; loc->open_end = d->token_end;
  size_t count = 0;
  for (lxb_html_token_attr_t *a = token->attr_first; a; a = a->next) count++;
  if (count) {
    loc->attrs = lexbor_mraw_calloc(tree->document->dom_document.mraw, count * sizeof(*loc->attrs));
    if (!loc->attrs) { d->failed = 1; return element; }
  }
  for (lxb_html_token_attr_t *a = token->attr_first; a; a = a->next) {
    size_t start = h_offset(d, a->name_begin), end = h_offset(d, a->name_end);
    if (start == SIZE_MAX || end == SIZE_MAX) continue;
    size_t attr_end = end;
    if (a->value_begin) {
      attr_end = h_offset(d, a->value_end);
      if (attr_end == SIZE_MAX) continue;
      size_t value_start = h_offset(d, a->value_begin);
      if (value_start && value_start != SIZE_MAX && (d->source[value_start - 1] == '"' || d->source[value_start - 1] == '\'')) attr_end++;
    }
    int duplicate = 0;
    for (size_t i = 0; i < loc->count; i++) {
      h_attr_location *old = &loc->attrs[i];
      if (old->length != end - start) continue;
      size_t n = 0;
      for (; n < old->length; n++) {
        unsigned char x = old->name[n], y = a->name_begin[n];
        if (x >= 'A' && x <= 'Z') x += 32;
        if (y >= 'A' && y <= 'Z') y += 32;
        if (x != y) break;
      }
      if (n == old->length) { duplicate = 1; break; }
    }
    if (!duplicate) loc->attrs[loc->count++] = (h_attr_location){a->name_begin, end - start, start, attr_end};
  }
  node->user = loc;
  return element;
}

static lxb_html_token_t *h_token(lxb_html_tokenizer_t *tkz, lxb_html_token_t *token, void *context) {
  h_document *d = context;
  if (d->token_filter && !d->token_filter(tkz, token, d->filter_context)) { d->failed = 1; return NULL; }
  if (d->track && token->type & LXB_HTML_TOKEN_TYPE_CLOSE) {
    size_t begin = h_offset(d, token->begin);
    if (begin != SIZE_MAX && begin >= 2) {
      lexbor_array_t *open = d->parser->tree->open_elements;
      for (size_t i = open->length; i-- > 0;) {
        lxb_dom_node_t *node = open->list[i];
        if (node->local_name == token->tag_id) {
          h_location *loc = node->user;
          if (loc) { loc->close_start = begin - 2; loc->end = d->token_end; }
          break;
        }
      }
    }
  }
  return d->token_callback(tkz, token, d->token_context);
}

static void h_destroy(h_document *d) {
  if (d->parser) lxb_html_parser_destroy(d->parser);
  if (d->document) lxb_html_document_destroy(d->document);
  free(d->utf16);
}
static int h_parse(h_document *d, const unsigned char *source, size_t length, int track) {
  d->source = source; d->length = length; d->track = track;
  d->document = lxb_html_document_create();
  d->parser = lxb_html_parser_create();
  if (!d->document || !d->parser || lxb_html_parser_init(d->parser) != LXB_STATUS_OK) return 0;
  if (track == 1) {
    if (length > SIZE_MAX / sizeof(size_t) - 1) return 0;
    d->utf16 = malloc((length + 1) * sizeof(size_t));
    if (!d->utf16) return 0;
    size_t units = 0;
    for (size_t i = 0; i < length; i++) {
      d->utf16[i] = units;
      if ((source[i] & 0xc0) != 0x80) units += source[i] >= 0xf0 ? 2 : 1;
    }
    d->utf16[length] = units;
  }
  if (lxb_html_parse_fragment_chunk_begin(d->parser, d->document, LXB_TAG_TEMPLATE, LXB_NS_HTML) != LXB_STATUS_OK) return 0;
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
    for (size_t i = 0; i < length; i++) if (source[i] == '>') {
      d->token_end = i + 1;
      if (lxb_html_parse_fragment_chunk_process(d->parser, source + start, i + 1 - start) != LXB_STATUS_OK) return 0;
      start = i + 1;
    }
    d->token_end = length;
    if (start < length && lxb_html_parse_fragment_chunk_process(d->parser, source + start, length - start) != LXB_STATUS_OK) return 0;
  } else if (lxb_html_parse_fragment_chunk_process(d->parser, source, length) != LXB_STATUS_OK) return 0;
  d->root = lxb_html_parse_fragment_chunk_end(d->parser);
  return d->root != NULL && !d->failed;
}
static const char *h_namespace(uintptr_t ns) {
  switch (ns) {
    case LXB_NS_HTML: return "http://www.w3.org/1999/xhtml";
    case LXB_NS_SVG: return "http://www.w3.org/2000/svg";
    case LXB_NS_MATH: return "http://www.w3.org/1998/Math/MathML";
    case LXB_NS_XLINK: return "http://www.w3.org/1999/xlink";
    case LXB_NS_XML: return "http://www.w3.org/XML/1998/namespace";
    case LXB_NS_XMLNS: return "http://www.w3.org/2000/xmlns/";
    default: return "";
  }
}
static void h_range(h_buffer *b, h_document *d, size_t start, size_t end) {
  if (start > d->length || end > d->length) { b->failed = 1; return; }
  h_s(b, "{\"startOffset\":"); h_number(b, d->utf16[start]);
  h_s(b, ",\"endOffset\":"); h_number(b, d->utf16[end]); h_s(b, "}");
}
static lxb_dom_node_t *h_children(lxb_dom_node_t *node) {
  if (node->type == LXB_DOM_NODE_TYPE_ELEMENT && node->ns == LXB_NS_HTML && node->local_name == LXB_TAG_TEMPLATE)
    return lxb_dom_interface_node(lxb_html_interface_template(node)->content)->first_child;
  return node->first_child;
}
static void h_json_node(h_buffer *b, h_document *d, lxb_dom_node_t *node, int root, unsigned depth) {
  if (depth > 512) { b->failed = H_DEPTH_EXCEEDED; return; }
  h_s(b, "{\"nodeName\":");
  if (root) h_json(b, (const unsigned char *)"#document-fragment", 18);
  else if (node->type == LXB_DOM_NODE_TYPE_TEXT || node->type == LXB_DOM_NODE_TYPE_COMMENT) {
    int text = node->type == LXB_DOM_NODE_TYPE_TEXT;
    h_json(b, (const unsigned char *)(text ? "#text" : "#comment"), text ? 5 : 8);
    h_s(b, text ? ",\"value\":" : ",\"data\":");
    lexbor_str_t *data = &((lxb_dom_character_data_t *)node)->data;
    h_json(b, data->data, data->length); h_s(b, "}"); return;
  } else if (node->type == LXB_DOM_NODE_TYPE_ELEMENT) {
    size_t length; const unsigned char *name = lxb_dom_element_qualified_name(lxb_dom_interface_element(node), &length);
    h_json(b, name, length); h_s(b, ",\"tagName\":"); h_json(b, name, length);
    h_s(b, ",\"namespaceURI\":"); h_json(b, (const unsigned char *)h_namespace(node->ns), strlen(h_namespace(node->ns)));
    h_s(b, ",\"attrs\":["); int first = 1;
    for (lxb_dom_attr_t *attr = lxb_dom_interface_element(node)->first_attr; attr; attr = attr->next) {
      if (!first) h_s(b, ","); first = 0;
      const unsigned char *attr_name = attr->node.ns >= LXB_NS_XLINK && attr->node.ns <= LXB_NS_XMLNS ? lxb_dom_attr_local_name(attr, &length) : lxb_dom_attr_qualified_name(attr, &length);
      h_s(b, "{\"name\":"); h_json(b, attr_name, length);
      if (attr->node.ns >= LXB_NS_XLINK && attr->node.ns <= LXB_NS_XMLNS) {
        h_s(b, ",\"namespace\":"); h_json(b, (const unsigned char *)h_namespace(attr->node.ns), strlen(h_namespace(attr->node.ns)));
        const char *prefix = attr->node.ns == LXB_NS_XLINK ? "xlink" : attr->node.ns == LXB_NS_XML ? "xml" : (length == 5 && !memcmp(attr_name, "xmlns", 5)) ? "" : "xmlns";
        h_s(b, ",\"prefix\":"); h_json(b, (const unsigned char *)prefix, strlen(prefix));
      }
      const unsigned char *value = lxb_dom_attr_value(attr, &length);
      h_s(b, ",\"value\":"); h_json(b, value, length); h_s(b, "}");
    }
    h_s(b, "]");
    h_location *loc = node->user;
    if (d->track && loc) {
      h_s(b, ",\"sourceCodeLocation\":{\"startOffset\":"); h_number(b, d->utf16[loc->start]);
      h_s(b, ",\"startTag\":"); h_range(b, d, loc->start, loc->open_end);
      if (loc->close_start != SIZE_MAX) { h_s(b, ",\"endTag\":"); h_range(b, d, loc->close_start, loc->end); }
      h_s(b, ",\"attrs\":{");
      for (size_t i = 0; i < loc->count; i++) {
        if (i) h_s(b, ",");
        h_attr_location *a = &loc->attrs[i];
        h_buffer name = {0};
        for (size_t n = 0; n < a->length; n++) { unsigned char c = a->name[n]; if (c >= 'A' && c <= 'Z') c += 32; h_put(&name, &c, 1); }
        if (name.failed) b->failed = name.failed;
        h_json(b, name.data, name.length); free(name.data); h_s(b, ":");
        h_range(b, d, a->start, a->end);
      }
      h_s(b, "}}");
    }
  } else { h_json(b, (const unsigned char *)"#unknown", 8); }
  int is_template = !root && node->type == LXB_DOM_NODE_TYPE_ELEMENT && node->local_name == LXB_TAG_TEMPLATE && node->ns == LXB_NS_HTML;
  if (is_template) h_s(b, ",\"content\":{\"nodeName\":\"#document-fragment\"");
  h_s(b, ",\"childNodes\":["); int first = 1;
  for (lxb_dom_node_t *child = h_children(node); child; child = child->next) {
    if (!first) h_s(b, ","); first = 0;
    h_json_node(b, d, child, 0, depth + 1);
  }
  h_s(b, "]}"); if (is_template) h_s(b, "}");
}

#include "sanitize.c"
