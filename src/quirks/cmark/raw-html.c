/* Raw cmark HTML literals are filtered before rendering: remove only actual
 * source-authored input tags, preserving the other source bytes. */
static int gfm_dom_has_input(const unsigned char *s, size_t length) {
  for (size_t i = 0; i + 6 <= length; i++)
    if (s[i] == '<' && gfm_ascii_prefix_case(s + i + 1, length - i - 1, "input") &&
        (i + 6 == length || gfm_ascii_space(s[i + 6]) || s[i + 6] == '/' || s[i + 6] == '>'))
      return 1;
  return 0;
}
typedef struct {
  size_t start, end;
} gfm_raw_span;
static int gfm_raw_span_compare(const void *a, const void *b) {
  const gfm_raw_span *x = a, *y = b;
  return x->start < y->start ? -1 : x->start > y->start ? 1 : 0;
}
static gfm_status gfm_dom_input_spans(lxb_dom_node_t *node, gfm_raw_span **spans, size_t *count,
                                      size_t *capacity, unsigned depth) {
  if (depth > 512)
    return GFM_LIMIT_EXCEEDED;
  if (node->local_name == LXB_TAG_INPUT && node->user) {
    h_location *loc = node->user;
    if (!gfm_grow((void **)spans, capacity, *count + 1, sizeof(**spans)))
      return GFM_NO_MEMORY;
    (*spans)[(*count)++] = (gfm_raw_span){loc->start, loc->open_end};
  }
  for (lxb_dom_node_t *child = h_children(node); child; child = child->next) {
    gfm_status status = gfm_dom_input_spans(child, spans, count, capacity, depth + 1);
    if (status != GFM_OK)
      return status;
  }
  return GFM_OK;
}
static gfm_status gfm_filter_raw_inputs(gfm_view source, gfm_string *result) {
  gfm_status status = GFM_NO_MEMORY;
  h_buffer out = {0};
  h_document d = {0};
  gfm_raw_span *spans = NULL;
  size_t count = 0, capacity = 0;
  if (gfm_dom_has_input(source.data, source.length)) {
    if (!h_parse(&d, source.data, source.length, H_TRACK_BYTES))
      goto fail;
    status = gfm_dom_input_spans(d.root, &spans, &count, &capacity, 0);
    if (status != GFM_OK)
      goto fail;
  }
  if (count > 1)
    qsort(spans, count, sizeof(*spans), gfm_raw_span_compare);
  size_t start = 0;
  for (size_t i = 0; i < count; i++) {
    if (spans[i].start < start || spans[i].end > source.length)
      continue;
    h_put(&out, source.data + start, spans[i].start - start);
    start = spans[i].end;
  }
  h_put(&out, source.data + start, source.length - start);
  h_destroy(&d);
  free(spans);
  if (out.failed) {
    free(out.data);
    return GFM_NO_MEMORY;
  }
  *result = (gfm_string){out.data, out.length};
  return GFM_OK;
fail:
  h_destroy(&d);
  free(spans);
  free(out.data);
  return status;
}
