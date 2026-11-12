/* Include the upstream amalgamation without modifying its bytes.
 *
 * A plain name macro would also rename every call, bypassing source tracking.
 * In the pinned upstream source, declarations/definitions start their first
 * argument with lxb_html_tree_t, while calls pass tree. Token pasting selects
 * the original implementation for declarations/definitions and the wrapper
 * for calls. The declaration branch also declares the wrapper with its full
 * types before any call sites. Other argument forms fail to compile.
 *
 * The wrapper's own name stays unexpanded while this macro is being expanded.
 * All adapter macros are removed immediately after the include. */
#define H_LEXBOR_CREATE_lxb_html_tree_t                                                            \
  lxb_html_tree_create_element_for_token(lxb_html_tree_t *tree, lxb_html_token_t *token,           \
                                         lxb_ns_id_t ns);                                          \
  lxb_html_element_t *gfm_lexbor_create_element_original(lxb_html_tree_t
#define H_LEXBOR_CREATE_tree lxb_html_tree_create_element_for_token(tree
#define lxb_html_tree_create_element_for_token(first, token, ns)                                   \
  H_LEXBOR_CREATE_##first, token, ns)
#include "../../vendor/lexbor/lexbor.h"
#undef lxb_html_tree_create_element_for_token
#undef H_LEXBOR_CREATE_tree
#undef H_LEXBOR_CREATE_lxb_html_tree_t
