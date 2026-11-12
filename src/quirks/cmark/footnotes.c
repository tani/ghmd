/* Recover malformed footnote references using cmark source positions.  This
 * runs before the final parse and never modifies the vendored parser. */
typedef struct {
  size_t start, end;
} gfm_span;
typedef struct {
  size_t start, end;
  char *text;
  size_t length;
} gfm_edit;
typedef struct {
  size_t length;
  int changed;
  unsigned char data[];
} gfm_recovered;
typedef struct {
  const unsigned char *source;
  size_t length, *lines, line_count;
  gfm_edit *edits;
  size_t edit_count, edit_capacity;
  int ok;
} gfm_recovery;

static int gfm_grow(void **data, size_t *capacity, size_t count, size_t width) {
  if (count <= *capacity)
    return 1;
  size_t next = *capacity ? *capacity * 2 : 16;
  if (next < count)
    next = count;
  if (next > SIZE_MAX / width)
    return 0;
  void *p = realloc(*data, next * width);
  if (!p)
    return 0;
  *data = p;
  *capacity = next;
  return 1;
}

static int gfm_location(gfm_recovery *r, cmark_node *node, gfm_span *span) {
  int a = cmark_node_get_start_line(node), b = cmark_node_get_start_column(node);
  int c = cmark_node_get_end_line(node), d = cmark_node_get_end_column(node);
  if (a < 1 || b < 1 || c < 1 || d < 1 || (size_t)a > r->line_count || (size_t)c > r->line_count)
    return 0;
  size_t start = r->lines[a - 1], end = r->lines[c - 1];
  size_t a_end = (size_t)a < r->line_count ? r->lines[a] - 1 : r->length;
  size_t c_end = (size_t)c < r->line_count ? r->lines[c] - 1 : r->length;
  start += (size_t)(b - 1) < a_end - start ? (size_t)(b - 1) : a_end - start;
  end += (size_t)d < c_end - end ? (size_t)d : c_end - end;
  span->start = start;
  span->end = end;
  return start <= end;
}

static int gfm_overlap(gfm_span *spans, size_t count, size_t start, size_t end) {
  for (size_t i = 0; i < count; i++)
    if (start < spans[i].end && end > spans[i].start)
      return 1;
  return 0;
}

static int gfm_escaped(const unsigned char *s, size_t start, size_t at) {
  size_t count = 0;
  while (at > start && s[--at] == '\\')
    count++;
  return count % 2 != 0;
}

static int gfm_add_edit(gfm_recovery *r, size_t start, size_t end, size_t label_start,
                        size_t label_end, int colon) {
  if (!gfm_grow((void **)&r->edits, &r->edit_capacity, r->edit_count + 1, sizeof(gfm_edit)))
    return 0;
  size_t n = label_end - label_start;
  if (n > SIZE_MAX - 5)
    return 0;
  char *text = malloc(n + 5);
  if (!text)
    return 0;
  memcpy(text, "[^", 2);
  memcpy(text + 2, r->source + label_start, n);
  text[n + 2] = ']';
  if (colon) {
    text[n + 3] = '\\';
    text[n + 4] = ':';
  }
  r->edits[r->edit_count++] = (gfm_edit){start, end, text, n + 3 + (colon ? 2 : 0)};
  return 1;
}

static int gfm_in_text(const cmark_strbuf *text, const unsigned char *raw, size_t length) {
  if (length > (size_t)text->size)
    return 0;
  for (size_t i = 0; i + length <= (size_t)text->size; i++)
    if (memcmp(text->ptr + i, raw, length) == 0)
      return 1;
  return 0;
}

static int gfm_paragraph_edits(gfm_recovery *r, cmark_node *paragraph) {
  gfm_span scope;
  if (!gfm_location(r, paragraph, &scope))
    return 1;
  cmark_strbuf text = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_node **stack = NULL;
  size_t stack_count = 0, stack_capacity = 0;
  gfm_span *protected = NULL;
  size_t protected_count = 0, protected_capacity = 0;
  int ok = gfm_grow((void **)&stack, &stack_capacity, 1, sizeof(*stack));
  if (ok)
    stack[stack_count++] = paragraph;
  while (ok && stack_count) {
    cmark_node *node = stack[--stack_count];
    const char *type = cmark_node_get_type_string(node);
    if (strcmp(type, "text") == 0)
      cmark_strbuf_puts(&text, cmark_node_get_literal(node));
    if (!strcmp(type, "code") || !strcmp(type, "code_block") || !strcmp(type, "html_inline") ||
        !strcmp(type, "html_block") || !strcmp(type, "link") || !strcmp(type, "image")) {
      gfm_span span;
      if (gfm_location(r, node, &span)) {
        ok = gfm_grow((void **)&protected, &protected_capacity, protected_count + 1,
                      sizeof(*protected));
        if (ok)
        protected[protected_count++] = span;
      }
    }
    for (cmark_node *child = cmark_node_first_child(node); ok && child;
         child = cmark_node_next(child)) {
      ok = gfm_grow((void **)&stack, &stack_capacity, stack_count + 1, sizeof(*stack));
      if (ok)
        stack[stack_count++] = child;
    }
  }
  free(stack);
  const unsigned char *s = r->source;
  if (ok && strstr(cmark_strbuf_cstr(&text), "[^]")) {
    for (size_t i = scope.start; i + 2 < scope.end; i++) {
      if (s[i] != '[' || s[i + 1] != '^')
        continue;
      size_t label = i + 2, p = label;
      while (p < scope.end && s[p] != ']' && s[p] != '\n')
        p++;
      if (p == label || p == scope.end || s[p] != '\n')
        continue;
      size_t label_end = p;
      p++;
      while (p < scope.end && s[p] != ']' && s[p] != '\n')
        p++;
      if (p == scope.end || s[p] != ']')
        continue;
      size_t end = p + 1;
      if (!gfm_overlap(protected, protected_count, i, end)) {
        int colon = end < r->length && s[end] == ':';
        ok = gfm_add_edit(r, i, end + colon, label, label_end, colon);
        if (!ok)
          break;
      }
      i = p;
    }
  }
  /* Find unresolved references containing an inline image; nested brackets,
   * parentheses, quoted destinations and escaped delimiters are supported. */
  for (size_t i = scope.start; ok && i + 3 < scope.end; i++) {
    if (s[i] != '[' || s[i + 1] != '^' || gfm_escaped(s, scope.start, i))
      continue;
    size_t image = i + 2;
    while (image < scope.end && s[image] != '\n' &&
           (s[image] != ']' || gfm_escaped(s, scope.start, image)) &&
           !(s[image] == '!' && image + 1 < scope.end && s[image + 1] == '[')) {
      if (s[image] == '[' && image + 1 < scope.end && s[image + 1] == '^' &&
          !gfm_escaped(s, scope.start, image))
        break;
      image++;
    }
    if (image + 1 >= scope.end || s[image] != '!' || s[image + 1] != '[' ||
        gfm_escaped(s, scope.start, image))
      continue;
    size_t p = image + 2;
    int depth = 1;
    while (p < scope.end && s[p] != '\n' && depth) {
      if (s[p] == '\\' && p + 1 < scope.end && s[p + 1] != '\n') {
        p += 2;
        continue;
      }
      if (s[p] == '[')
        depth++;
      else if (s[p] == ']')
        depth--;
      p++;
    }
    if (depth || p >= scope.end || s[p] != '(')
      continue;
    p++;
    depth = 1;
    int angle = 0;
    unsigned char quote = 0;
    while (p < scope.end && s[p] != '\n' && depth) {
      unsigned char ch = s[p];
      if (ch == '\\' && p + 1 < scope.end && s[p + 1] != '\n') {
        p += 2;
        continue;
      }
      if (quote) {
        if (ch == quote)
          quote = 0;
      } else if (angle) {
        if (ch == '>')
          angle = 0;
      } else if (ch == '<')
        angle = 1;
      else if ((ch == '"' || ch == '\'') && p > scope.start &&
               (s[p - 1] == ' ' || s[p - 1] == '\t' || s[p - 1] == '\r' || s[p - 1] == '\f' ||
                s[p - 1] == '\v'))
        quote = ch;
      else if (ch == '(')
        depth++;
      else if (ch == ')')
        depth--;
      p++;
    }
    if (depth || angle || quote || p >= scope.end || s[p] != ']')
      continue;
    size_t end = p + 1;
    if (gfm_in_text(&text, s + i, end - i) && !gfm_overlap(protected, protected_count, i, end))
      ok = gfm_add_edit(r, i, end, i + 2, image, 0);
    i = p;
  }
  free(protected);
  cmark_strbuf_free(&text);
  return ok;
}

static int gfm_edit_compare(const void *a, const void *b) {
  const gfm_edit *left = a, *right = b;
  return left->start < right->start ? 1 : left->start > right->start ? -1 : 0;
}

gfm_recovered *gfm_recover_labels(cmark_node *root, const unsigned char *source, size_t length) {
  gfm_recovery r = {.source = source, .length = length, .line_count = 1, .ok = 1};
  for (size_t i = 0; i < length; i++)
    if (source[i] == '\n')
      r.line_count++;
  r.lines = malloc(r.line_count * sizeof(*r.lines));
  if (!r.lines)
    return NULL;
  size_t n = 0;
  r.lines[n++] = 0;
  for (size_t i = 0; i < length; i++)
    if (source[i] == '\n')
      r.lines[n++] = i + 1;
  cmark_iter *iter = cmark_iter_new(root);
  if (!iter) {
    free(r.lines);
    return NULL;
  }
  cmark_event_type event;
  while ((event = cmark_iter_next(iter)) != CMARK_EVENT_DONE) {
    if (event != CMARK_EVENT_ENTER)
      continue;
    cmark_node *node = cmark_iter_get_node(iter);
    if (!strcmp(cmark_node_get_type_string(node), "paragraph") && !gfm_paragraph_edits(&r, node)) {
      r.ok = 0;
      break;
    }
  }
  cmark_iter_free(iter);
  free(r.lines);
  gfm_recovered *result = NULL;
  if (r.ok) {
    if (r.edit_count > 1)
      qsort(r.edits, r.edit_count, sizeof(*r.edits), gfm_edit_compare);
    cmark_strbuf output = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    cmark_strbuf_put(&output, source, length);
    for (size_t i = 0; i < r.edit_count; i++) {
      gfm_edit *edit = &r.edits[i];
      size_t current = (size_t)output.size;
      size_t start = edit->start < current ? edit->start : current;
      size_t end = edit->end < current ? edit->end : current;
      if (end < start)
        end = start;
      size_t total = current - (end - start) + edit->length;
      cmark_strbuf_grow(&output, total + 1);
      memmove(output.ptr + start + edit->length, output.ptr + end, current - end);
      memcpy(output.ptr + start, edit->text, edit->length);
      output.size = total;
      output.ptr[total] = 0;
    }
    result = malloc(sizeof(*result) + (size_t)output.size + 1);
    if (result) {
      result->length = output.size;
      result->changed =
          output.size != length || (length && memcmp(output.ptr, source, length) != 0);
      memcpy(result->data, output.ptr, (size_t)output.size + 1);
    }
    cmark_strbuf_free(&output);
  }
  for (size_t i = 0; i < r.edit_count; i++)
    free(r.edits[i].text);
  free(r.edits);
  return result;
}
