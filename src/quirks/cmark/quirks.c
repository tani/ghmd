/* GitHub compatibility adjustments applied to the unmodified cmark AST. */
#include <stdlib.h>
#include <string.h>

cmark_parser *gfm_parser_new(int options) {
  static const char *const names[] = {"table", "autolink", "tagfilter", "strikethrough",
                                      "tasklist"};
  cmark_parser *parser = cmark_parser_new(options);
  if (!parser)
    return NULL;
  for (size_t i = 0; i < sizeof(names) / sizeof(names[0]); i++) {
    cmark_syntax_extension *extension = cmark_find_syntax_extension(names[i]);
    if (!extension || !cmark_parser_attach_syntax_extension(parser, extension)) {
      cmark_parser_free(parser);
      return NULL;
    }
  }
  return parser;
}

/* Returns the beginning of list-item content, or NULL if there is no marker. */
static const char *gfm_list_marker(const char *p, const char *end) {
  int spaces = 0;
  while (p < end && *p == ' ' && spaces < 3) {
    p++;
    spaces++;
  }
  if (p == end)
    return NULL;
  if (*p == '-' || *p == '+' || *p == '*')
    p++;
  else {
    int digits = 0;
    while (p < end && *p >= '0' && *p <= '9' && digits < 9) {
      p++;
      digits++;
    }
    if (!digits || p == end || (*p != '.' && *p != ')'))
      return NULL;
    p++;
  }
  if (p == end || (*p != ' ' && *p != '\t'))
    return NULL;
  do {
    p++;
  } while (p < end && (*p == ' ' || *p == '\t'));
  return p;
}

/* Apply the same unresolved-reference recovery only to text nodes. */
static int gfm_adapt_text(cmark_node *node) {
  const char *input = cmark_node_get_literal(node);
  if (!input || !strstr(input, "[^"))
    return 1;
  size_t length = strlen(input);
  cmark_strbuf first = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  cmark_strbuf result = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  for (size_t i = 0; i < length;) {
    size_t image = i + 2;
    if (input[i] == '[' && i + 1 < length && input[i + 1] == '^') {
      while (image < length && input[image] != ']' && input[image] != '\n' &&
             !(input[image] == '!' && image + 1 < length && input[image + 1] == '['))
        image++;
      if (image + 1 < length && input[image] == '!' && input[image + 1] == '[') {
        size_t p = image + 2;
        while (p < length && input[p] != ']' && input[p] != '\n')
          p++;
        if (p + 1 < length && input[p] == ']' && input[p + 1] == '(') {
          p += 2;
          while (p < length && input[p] != ')' && input[p] != '\n') {
            if (input[p] == '\\' && p + 1 < length && input[p + 1] != '\n')
              p += 2;
            else
              p++;
          }
          if (p + 1 < length && input[p] == ')' && input[p + 1] == ']') {
            cmark_strbuf_put(&first, (const unsigned char *)input + i, image - i);
            cmark_strbuf_putc(&first, ']');
            i = p + 2;
            continue;
          }
        }
      }
    }
    cmark_strbuf_putc(&first, (unsigned char)input[i++]);
  }
  const char *value = cmark_strbuf_cstr(&first);
  length = strlen(value);
  for (size_t i = 0; i < length;) {
    if (value[i] == '[' && i + 1 < length && value[i + 1] == '^') {
      size_t p = i + 2;
      while (p < length && value[p] != ']' && value[p] != '\n') {
        if (value[p] == '\\') {
          if (p + 1 == length || value[p + 1] == '\n')
            break;
          p += 2;
        } else
          p++;
      }
      if (p < length && value[p] == ']') {
        for (size_t j = i; j <= p; j++) {
          if (value[j] == '\\' && j < p &&
              strchr("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", value[j + 1]))
            j++;
          cmark_strbuf_putc(&result, (unsigned char)value[j]);
        }
        i = p + 1;
        continue;
      }
    }
    cmark_strbuf_putc(&result, (unsigned char)value[i++]);
  }
  int ok = strcmp(input, cmark_strbuf_cstr(&result)) == 0 ||
           cmark_node_set_literal(node, cmark_strbuf_cstr(&result));
  cmark_strbuf_free(&first);
  cmark_strbuf_free(&result);
  return ok;
}

/* Source columns are 1-based UTF-8 byte offsets, so no transcoding is needed. */
int gfm_adapt_nodes(cmark_node *root, const char *source, size_t length) {
  size_t count = 1;
  for (size_t i = 0; i < length; i++)
    if (source[i] == '\n')
      count++;
  if (count > SIZE_MAX / sizeof(size_t))
    return 0;
  size_t *lines = malloc(count * sizeof(*lines));
  if (!lines)
    return 0;
  size_t n = 0;
  lines[n++] = 0;
  for (size_t i = 0; i < length; i++)
    if (source[i] == '\n')
      lines[n++] = i + 1;
  cmark_iter *iter = cmark_iter_new(root);
  if (!iter) {
    free(lines);
    return 0;
  }
  int ok = 1;
  cmark_event_type event;
  while ((event = cmark_iter_next(iter)) != CMARK_EVENT_DONE) {
    if (event != CMARK_EVENT_ENTER)
      continue;
    cmark_node *node = cmark_iter_get_node(iter);
    const char *type = cmark_node_get_type_string(node);
    if (strcmp(type, "text") == 0) {
      if (!gfm_adapt_text(node)) {
        ok = 0;
        break;
      }
      continue;
    }
    if (strcmp(type, "item") != 0 && strcmp(type, "tasklist") != 0)
      continue;
    int line = cmark_node_get_start_line(node), column = cmark_node_get_start_column(node);
    if (line < 1 || (size_t)line > count || column < 1)
      continue;
    size_t begin = lines[line - 1], end = (size_t)line < count ? lines[line] - 1 : length;
    if ((size_t)(column - 1) > end - begin)
      continue;
    const char *marker = gfm_list_marker(source + begin + column - 1, source + end);
    if (!marker)
      continue;
    if (strcmp(type, "tasklist") == 0) {
      if (source + end - marker >= 3 && marker[0] == '[' &&
          (marker[1] == ' ' || marker[1] == 'x' || marker[1] == 'X') && marker[2] == ']')
        cmark_gfm_extensions_set_tasklist_item_checked(node, marker[1] != ' ');
      continue;
    }
    cmark_node *paragraph = cmark_node_first_child(node);
    if (!paragraph || strcmp(cmark_node_get_type_string(paragraph), "paragraph") != 0)
      continue;
    cmark_node *first = cmark_node_first_child(paragraph);
    if (!first || strcmp(cmark_node_get_type_string(first), "text") != 0)
      continue;
    const char *text = cmark_node_get_literal(first);
    if (!text || strlen(text) < 4 || text[0] != '[' ||
        (text[1] != ' ' && text[1] != 'x' && text[1] != 'X') || text[2] != ']' ||
        (text[3] != ' ' && text[3] != '\t'))
      continue;
    cmark_syntax_extension *extension = cmark_find_syntax_extension("tasklist");
    if (!extension || !cmark_node_set_syntax_extension(node, extension)) {
      ok = 0;
      break;
    }
    cmark_gfm_extensions_set_tasklist_item_checked(node, text[1] != ' ');
    if (!cmark_node_set_literal(first, text + 4)) {
      ok = 0;
      break;
    }
  }
  cmark_iter_free(iter);
  free(lines);
  return ok;
}
