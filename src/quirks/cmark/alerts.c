/* GitHub alert titles, emitted after HTML sanitization. */
static const char *const gfm_alert_names[] = {"note", "tip", "important", "warning", "caution"};
static const char *const gfm_alert_titles[] = {
    "<svg data-component=\"Octicon\" class=\"octicon octicon-info mr-2\" viewBox=\"0 0 16 16\" version=\"1.1\" width=\"16\" height=\"16\" aria-hidden=\"true\"><path d=\"M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8Zm8-6.5a6.5 6.5 0 1 0 0 13 6.5 6.5 0 0 0 0-13ZM6.5 7.75A.75.75 0 0 1 7.25 7h1a.75.75 0 0 1 .75.75v2.75h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5h.25v-2h-.25a.75.75 0 0 1-.75-.75ZM8 6a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z\"></path></svg>Note", "<svg data-component=\"Octicon\" class=\"octicon octicon-light-bulb mr-2\" viewBox=\"0 0 16 16\" version=\"1.1\" width=\"16\" height=\"16\" aria-hidden=\"true\"><path d=\"M8 1.5c-2.363 0-4 1.69-4 3.75 0 .984.424 1.625.984 2.304l.214.253c.223.264.47.556.673.848.284.411.537.896.621 1.49a.75.75 0 0 1-1.484.211c-.04-.282-.163-.547-.37-.847a8.456 8.456 0 0 0-.542-.68c-.084-.1-.173-.205-.268-.32C3.201 7.75 2.5 6.766 2.5 5.25 2.5 2.31 4.863 0 8 0s5.5 2.31 5.5 5.25c0 1.516-.701 2.5-1.328 3.259-.095.115-.184.22-.268.319-.207.245-.383.453-.541.681-.208.3-.33.565-.37.847a.751.751 0 0 1-1.485-.212c.084-.593.337-1.078.621-1.489.203-.292.45-.584.673-.848.075-.088.147-.173.213-.253.561-.679.985-1.32.985-2.304 0-2.06-1.637-3.75-4-3.75ZM5.75 12h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1 0-1.5ZM6 15.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75Z\"></path></svg>Tip", "<svg data-component=\"Octicon\" class=\"octicon octicon-report mr-2\" viewBox=\"0 0 16 16\" version=\"1.1\" width=\"16\" height=\"16\" aria-hidden=\"true\"><path d=\"M0 1.75C0 .784.784 0 1.75 0h12.5C15.216 0 16 .784 16 1.75v9.5A1.75 1.75 0 0 1 14.25 13H8.06l-2.573 2.573A1.458 1.458 0 0 1 3 14.543V13H1.75A1.75 1.75 0 0 1 0 11.25Zm1.75-.25a.25.25 0 0 0-.25.25v9.5c0 .138.112.25.25.25h2a.75.75 0 0 1 .75.75v2.19l2.72-2.72a.749.749 0 0 1 .53-.22h6.5a.25.25 0 0 0 .25-.25v-9.5a.25.25 0 0 0-.25-.25Zm7 2.25v2.5a.75.75 0 0 1-1.5 0v-2.5a.75.75 0 0 1 1.5 0ZM9 9a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z\"></path></svg>Important", "<svg data-component=\"Octicon\" class=\"octicon octicon-alert mr-2\" viewBox=\"0 0 16 16\" version=\"1.1\" width=\"16\" height=\"16\" aria-hidden=\"true\"><path d=\"M6.457 1.047c.659-1.234 2.427-1.234 3.086 0l6.082 11.378A1.75 1.75 0 0 1 14.082 15H1.918a1.75 1.75 0 0 1-1.543-2.575Zm1.763.707a.25.25 0 0 0-.44 0L1.698 13.132a.25.25 0 0 0 .22.368h12.164a.25.25 0 0 0 .22-.368Zm.53 3.996v2.5a.75.75 0 0 1-1.5 0v-2.5a.75.75 0 0 1 1.5 0ZM9 11a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z\"></path></svg>Warning", "<svg data-component=\"Octicon\" class=\"octicon octicon-stop mr-2\" viewBox=\"0 0 16 16\" version=\"1.1\" width=\"16\" height=\"16\" aria-hidden=\"true\"><path d=\"M4.47.22A.749.749 0 0 1 5 0h6c.199 0 .389.079.53.22l4.25 4.25c.141.14.22.331.22.53v6a.749.749 0 0 1-.22.53l-4.25 4.25A.749.749 0 0 1 11 16H5a.749.749 0 0 1-.53-.22L.22 11.53A.749.749 0 0 1 0 11V5c0-.199.079-.389.22-.53Zm.84 1.28L1.5 5.31v5.38l3.81 3.81h5.38l3.81-3.81V5.31L10.69 1.5ZM8 4a.75.75 0 0 1 .75.75v3.5a.75.75 0 0 1-1.5 0v-3.5A.75.75 0 0 1 8 4Zm0 8a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z\"></path></svg>Caution"};

/* Apply the alert conversion in document order, as the DOM loop does. */
static int gfm_adjust_alerts(gfm_html_tree *tree) {
  int changed = 0;
  for (size_t id = 1; id < tree->count; id++) {
    if (tree->nodes[id].inert || tree->nodes[id].removed || !gfm_node_is(tree, id, "blockquote"))
      continue;
    int nested = 0;
    for (size_t parent = tree->nodes[id].parent; parent != GFM_NONE;
         parent = tree->nodes[parent].parent)
      if (gfm_node_is(tree, parent, "blockquote") && !tree->nodes[parent].alert) {
        nested = 1;
        break;
      }
    if (nested)
      continue;
    size_t paragraph = tree->nodes[id].first;
    while (paragraph != GFM_NONE && !tree->nodes[paragraph].element)
      paragraph = tree->nodes[paragraph].next;
    if (paragraph == GFM_NONE || !gfm_node_is(tree, paragraph, "p"))
      continue;
    size_t marker = tree->nodes[paragraph].first;
    if (marker == GFM_NONE || tree->nodes[marker].element || tree->nodes[marker].synthetic)
      continue;
    size_t br = tree->nodes[marker].next;
    if (br == GFM_NONE || !gfm_node_is(tree, br, "br"))
      continue;
    cmark_strbuf text = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    houdini_unescape_html_f(&text, tree->source + tree->nodes[marker].start,
                            tree->nodes[marker].end - tree->nodes[marker].start);
    size_t kind = 0;
    static const char *const markers[] = {"[!NOTE]", "[!TIP]", "[!IMPORTANT]", "[!WARNING]",
                                          "[!CAUTION]"};
    while (kind < sizeof(markers) / sizeof(markers[0]) &&
           (text.size != strlen(markers[kind]) || memcmp(text.ptr, markers[kind], text.size)))
      kind++;
    cmark_strbuf_free(&text);
    if (kind == sizeof(markers) / sizeof(markers[0]))
      continue;
    gfm_detach_node(tree, br);
    tree->nodes[br].removed = 1;
    gfm_detach_node(tree, marker);
    tree->nodes[marker].removed = 1;
    size_t first = tree->nodes[paragraph].first;
    if (first != GFM_NONE && !tree->nodes[first].element && !tree->nodes[first].synthetic &&
        tree->nodes[first].start < tree->nodes[first].end &&
        tree->source[tree->nodes[first].start] == '\n')
      tree->nodes[first].start++;
    first = tree->nodes[id].first;
    if (first != GFM_NONE && !tree->nodes[first].element &&
        !gfm_text_has_visible_content(tree, first)) {
      gfm_detach_node(tree, first);
      tree->nodes[first].removed = 1;
    }
    tree->nodes[id].alert = kind + 1;
    changed = 1;
  }
  return changed;
}
