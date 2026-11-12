/* Translation-unit entry point: the preprocessor includes unmodified upstream
 * files at compile time. Internal names are scoped per original C file, and
 * source-local macros are undefined between files to avoid collisions.
 * The upstream-source test verifies that every vendored C file is included;
 * behavioral tests check the combined build against saved references. */
#include "../../html/core.c"
#define CMARK_GFM_STATIC_DEFINE 1
#define CMARK_GFM_EXTENSIONS_STATIC_DEFINE 1
#define match unity_extensions_autolink_match
#include "../../../vendor/cmark-gfm/extensions/autolink.c"
#undef strncasecmp
#undef match
#include "../../../vendor/cmark-gfm/extensions/core-extensions.c"
#include "../../../vendor/cmark-gfm/extensions/ext_scanners.c"
#define match unity_extensions_strikethrough_match
#define can_contain unity_extensions_strikethrough_can_contain
#define commonmark_render unity_extensions_strikethrough_commonmark_render
#define get_type_string unity_extensions_strikethrough_get_type_string
#define html_render unity_extensions_strikethrough_html_render
#define latex_render unity_extensions_strikethrough_latex_render
#define man_render unity_extensions_strikethrough_man_render
#include "../../../vendor/cmark-gfm/extensions/strikethrough.c"
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
#include "../../../vendor/cmark-gfm/extensions/table.c"
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
#include "../../../vendor/cmark-gfm/extensions/tagfilter.c"
#define can_contain unity_extensions_tasklist_can_contain
#define commonmark_render unity_extensions_tasklist_commonmark_render
#define get_type_string unity_extensions_tasklist_get_type_string
#define html_render unity_extensions_tasklist_html_render
#define matches unity_extensions_tasklist_matches
#define xml_attr unity_extensions_tasklist_xml_attr
#define parse_node_item_prefix unity_extensions_tasklist_parse_node_item_prefix
#include "../../../vendor/cmark-gfm/extensions/tasklist.c"
#undef can_contain
#undef commonmark_render
#undef get_type_string
#undef html_render
#undef matches
#undef xml_attr
#undef parse_node_item_prefix
#include "../../../vendor/cmark-gfm/src/arena.c"
#define contains_inlines unity_src_blocks_contains_inlines
#define parse_node_item_prefix unity_src_blocks_parse_node_item_prefix
#define S_is_line_end_char unity_src_blocks_S_is_line_end_char
#include "../../../vendor/cmark-gfm/src/blocks.c"
#undef CODE_INDENT
#undef TAB_STOP
#undef MAX_LIST_DEPTH
#undef MIN
#undef peek_at
#undef contains_inlines
#undef parse_node_item_prefix
#undef S_is_line_end_char
#include "../../../vendor/cmark-gfm/src/buffer.c"
#undef MIN
#include "../../../vendor/cmark-gfm/src/cmark.c"
#include "../../../vendor/cmark-gfm/src/cmark_ctype.c"
#define S_render_node unity_src_commonmark_S_render_node
#define outc unity_src_commonmark_outc
#include "../../../vendor/cmark-gfm/src/commonmark.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef ENCODED_SIZE
#undef LISTMARKER_SIZE
#undef S_render_node
#undef outc
#include "../../../vendor/cmark-gfm/src/footnotes.c"
#include "../../../vendor/cmark-gfm/src/houdini_href_e.c"
#include "../../../vendor/cmark-gfm/src/houdini_html_e.c"
#include "../../../vendor/cmark-gfm/src/houdini_html_u.c"
#define S_render_node unity_src_html_S_render_node
#include "../../../vendor/cmark-gfm/src/html.c"
#undef S_render_node
#define S_is_line_end_char unity_src_inlines_S_is_line_end_char
#include "../../../vendor/cmark-gfm/src/inlines.c"
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
#include "../../../vendor/cmark-gfm/src/iterator.c"
#define S_render_node unity_src_latex_S_render_node
#define outc unity_src_latex_outc
#include "../../../vendor/cmark-gfm/src/latex.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef LIST_NUMBER_STRING_SIZE
#undef S_render_node
#undef outc
#include "../../../vendor/cmark-gfm/src/linked_list.c"
#define S_render_node unity_src_man_S_render_node
#include "../../../vendor/cmark-gfm/src/man.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef LIST_NUMBER_SIZE
#undef S_render_node
#include "../../../vendor/cmark-gfm/src/map.c"
#include "../../../vendor/cmark-gfm/src/node.c"
#undef NODE_MEM
#define S_render_node unity_src_plaintext_S_render_node
#define outc unity_src_plaintext_outc
#include "../../../vendor/cmark-gfm/src/plaintext.c"
#undef OUT
#undef LIT
#undef CR
#undef BLANKLINE
#undef LISTMARKER_SIZE
#undef S_render_node
#undef outc
#include "../../../vendor/cmark-gfm/src/plugin.c"
#include "../../../vendor/cmark-gfm/src/references.c"
#include "../../../vendor/cmark-gfm/src/registry.c"
#include "../../../vendor/cmark-gfm/src/render.c"
#include "../../../vendor/cmark-gfm/src/scanners.c"
#include "../../../vendor/cmark-gfm/src/syntax_extension.c"
#include "../../../vendor/cmark-gfm/src/utf8.c"
#undef bufpush
#define S_render_node unity_src_xml_S_render_node
#include "../../../vendor/cmark-gfm/src/xml.c"
#undef BUFFER_SIZE
#undef MAX_INDENT
#undef S_render_node
#include <stdlib.h>
#include "quirks.c"
#include "footnotes.c"
#include "math.c"
#include "unicode.c"
#include "html.c"
#include "runtime.c"
#include "alerts.c"
#include "list.c"
#include "math-dom.c"
#include "postprocess.c"

#include "stages.c"
#include "dom.c"
#include "api.c"
#include "document.c"
