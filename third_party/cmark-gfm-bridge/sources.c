#define CMARK_GFM_STATIC_DEFINE 1
#define CMARK_GFM_EXTENSIONS_STATIC_DEFINE 1
#define match unity_extensions_autolink_match
#include "../../third_party/cmark-gfm/extensions/autolink.c"
#undef strncasecmp
#undef match
#include "../../third_party/cmark-gfm/extensions/core-extensions.c"
#include "../../third_party/cmark-gfm/extensions/ext_scanners.c"
#define match unity_extensions_strikethrough_match
#define can_contain unity_extensions_strikethrough_can_contain
#define commonmark_render unity_extensions_strikethrough_commonmark_render
#define get_type_string unity_extensions_strikethrough_get_type_string
#define html_render unity_extensions_strikethrough_html_render
#define latex_render unity_extensions_strikethrough_latex_render
#define man_render unity_extensions_strikethrough_man_render
#include "../../third_party/cmark-gfm/extensions/strikethrough.c"
#undef match
#undef can_contain
#undef commonmark_render
#undef get_type_string
#undef html_render
#undef latex_render
#undef man_render
#define can_contain unity_extensions_table_can_contain
#define commonmark_render unity_extensions_table_commonmark_render
#define get_type_string unity_extensions_table_get_type_string
#define html_render unity_extensions_table_html_render
#define latex_render unity_extensions_table_latex_render
#define man_render unity_extensions_table_man_render
#define contains_inlines unity_extensions_table_contains_inlines
#define matches unity_extensions_table_matches
#define xml_attr unity_extensions_table_xml_attr
#include "../../third_party/cmark-gfm/extensions/table.c"
#undef MAX_AUTOCOMPLETED_CELLS
#undef can_contain
#undef commonmark_render
#undef get_type_string
#undef html_render
#undef latex_render
#undef man_render
#undef contains_inlines
#undef matches
#undef xml_attr
#include "../../third_party/cmark-gfm/extensions/tagfilter.c"
#define can_contain unity_extensions_tasklist_can_contain
#define commonmark_render unity_extensions_tasklist_commonmark_render
#define get_type_string unity_extensions_tasklist_get_type_string
#define html_render unity_extensions_tasklist_html_render
#define matches unity_extensions_tasklist_matches
#define xml_attr unity_extensions_tasklist_xml_attr
#define parse_node_item_prefix unity_extensions_tasklist_parse_node_item_prefix
#include "../../third_party/cmark-gfm/extensions/tasklist.c"
#undef can_contain
#undef commonmark_render
#undef get_type_string
#undef html_render
#undef matches
#undef xml_attr
#undef parse_node_item_prefix
#include "../../third_party/cmark-gfm/src/arena.c"
#define contains_inlines unity_src_blocks_contains_inlines
#define parse_node_item_prefix unity_src_blocks_parse_node_item_prefix
#define S_is_line_end_char unity_src_blocks_S_is_line_end_char
#include "../../third_party/cmark-gfm/src/blocks.c"
#undef CODE_INDENT
#undef TAB_STOP
#undef MAX_LIST_DEPTH
#undef MIN
#undef peek_at
#undef contains_inlines
#undef parse_node_item_prefix
#undef S_is_line_end_char
#include "../../third_party/cmark-gfm/src/buffer.c"
#undef MIN
#include "../../third_party/cmark-gfm/src/cmark.c"
#include "../../third_party/cmark-gfm/src/cmark_ctype.c"
#define S_render_node unity_src_commonmark_S_render_node
#define outc unity_src_commonmark_outc
#include "../../third_party/cmark-gfm/src/commonmark.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef ENCODED_SIZE
#undef LISTMARKER_SIZE
#undef S_render_node
#undef outc
#include "../../third_party/cmark-gfm/src/footnotes.c"
#include "../../third_party/cmark-gfm/src/houdini_href_e.c"
#include "../../third_party/cmark-gfm/src/houdini_html_e.c"
#include "../../third_party/cmark-gfm/src/houdini_html_u.c"
#define S_render_node unity_src_html_S_render_node
#include "../../third_party/cmark-gfm/src/html.c"
#undef S_render_node
#define S_is_line_end_char unity_src_inlines_S_is_line_end_char
#include "../../third_party/cmark-gfm/src/inlines.c"
#undef make_str
#undef make_code
#undef make_raw_html
#undef make_linebreak
#undef make_softbreak
#undef make_emph
#undef make_strong
#undef MAXBACKTICKS
#undef FLAG_SKIP_HTML_CDATA
#undef FLAG_SKIP_HTML_DECLARATION
#undef FLAG_SKIP_HTML_PI
#undef FLAG_SKIP_HTML_COMMENT
#undef advance
#undef S_is_line_end_char
#include "../../third_party/cmark-gfm/src/iterator.c"
#define S_render_node unity_src_latex_S_render_node
#define outc unity_src_latex_outc
#include "../../third_party/cmark-gfm/src/latex.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef LIST_NUMBER_STRING_SIZE
#undef S_render_node
#undef outc
#include "../../third_party/cmark-gfm/src/linked_list.c"
#define S_render_node unity_src_man_S_render_node
#include "../../third_party/cmark-gfm/src/man.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef LIST_NUMBER_SIZE
#undef S_render_node
#include "../../third_party/cmark-gfm/src/map.c"
#include "../../third_party/cmark-gfm/src/node.c"
#undef NODE_MEM
#define S_render_node unity_src_plaintext_S_render_node
#define outc unity_src_plaintext_outc
#include "../../third_party/cmark-gfm/src/plaintext.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef LISTMARKER_SIZE
#undef S_render_node
#undef outc
#include "../../third_party/cmark-gfm/src/plugin.c"
#include "../../third_party/cmark-gfm/src/references.c"
#include "../../third_party/cmark-gfm/src/registry.c"
#include "../../third_party/cmark-gfm/src/render.c"
#include "../../third_party/cmark-gfm/src/scanners.c"
#include "../../third_party/cmark-gfm/src/syntax_extension.c"
#include "../../third_party/cmark-gfm/src/utf8.c"
#undef bufpush
#define S_render_node unity_src_xml_S_render_node
#include "../../third_party/cmark-gfm/src/xml.c"


#include <stdlib.h>
#include <string.h>

typedef struct bridge_document {
  cmark_parser *parser;
  cmark_node *doc;
} bridge_document;

void bridge_register_extensions(void) { cmark_gfm_core_extensions_ensure_registered(); }

static void bridge_parser_options_default(void) {}

bridge_document *bridge_parse(const char *source, size_t length) {
  bridge_register_extensions();
  cmark_parser *parser = cmark_parser_new(CMARK_OPT_UNSAFE | CMARK_OPT_FOOTNOTES | CMARK_OPT_GITHUB_PRE_LANG | CMARK_OPT_HARDBREAKS);
  if (!parser)
    return NULL;
  const char *const names[] = {"table", "autolink", "tagfilter", "strikethrough", "tasklist"};
  for (size_t i = 0; i < sizeof(names) / sizeof(names[0]); i++) {
    cmark_syntax_extension *extension = cmark_find_syntax_extension(names[i]);
    if (!extension || !cmark_parser_attach_syntax_extension(parser, extension)) {
      cmark_parser_free(parser);
      return NULL;
    }
  }
  cmark_parser_feed(parser, source, length);
  cmark_node *doc = cmark_parser_finish(parser);
  if (!doc) {
    cmark_parser_free(parser);
    return NULL;
  }
  bridge_document *document = (bridge_document *)calloc(1, sizeof(*document));
  if (!document) {
    cmark_node_free(doc);
    cmark_parser_free(parser);
    return NULL;
  }
  document->parser = parser;
  document->doc = doc;
  return document;
}

void bridge_free(bridge_document *document) {
  if (!document)
    return;
  if (document->doc)
    cmark_node_free(document->doc);
  if (document->parser)
    cmark_parser_free(document->parser);
  free(document);
}

void *bridge_root(bridge_document *document) { return document->doc; }

char *bridge_render_html(bridge_document *document) {
  return cmark_render_html(document->doc, CMARK_OPT_UNSAFE | CMARK_OPT_FOOTNOTES |
                                           CMARK_OPT_GITHUB_PRE_LANG | CMARK_OPT_HARDBREAKS,
                          cmark_parser_get_syntax_extensions(document->parser));
}

void bridge_free_string(char *value) { free(value); }
const char *bridge_version(void) { return cmark_version_string(); }
const char *bridge_node_type(void *node) { return cmark_node_get_type_string((cmark_node *)node); }
void *bridge_first_child(void *node) { return cmark_node_first_child((cmark_node *)node); }
void *bridge_next(void *node) { return cmark_node_next((cmark_node *)node); }
void *bridge_parent(void *node) { return cmark_node_parent((cmark_node *)node); }
const char *bridge_node_literal(void *node) { return cmark_node_get_literal((cmark_node *)node); }
int bridge_set_literal(void *node, const char *literal, size_t length) {
  (void)length;
  return cmark_node_set_literal((cmark_node *)node, literal);
}
int bridge_start_line(void *node) { return cmark_node_get_start_line((cmark_node *)node); }
int bridge_start_column(void *node) { return cmark_node_get_start_column((cmark_node *)node); }
int bridge_end_line(void *node) { return cmark_node_get_end_line((cmark_node *)node); }
int bridge_end_column(void *node) { return cmark_node_get_end_column((cmark_node *)node); }
int bridge_set_extension(void *node, const char *name) {
  cmark_syntax_extension *extension = cmark_find_syntax_extension(name);
  if (!extension)
    return 0;
  return cmark_node_set_syntax_extension((cmark_node *)node, extension);
}
int bridge_set_task_checked(void *node, int checked) {
  return cmark_gfm_extensions_set_tasklist_item_checked((cmark_node *)node, checked);
}
int bridge_get_task_checked(void *node) {
  return cmark_gfm_extensions_get_tasklist_item_checked((cmark_node *)node);
}
