/* Math transformations over the serialized, sanitized DOM tree. */
typedef struct { char *outer, *inner; unsigned char trim_first, trim_last; } gfm_math_edit;

static void gfm_math_text(cmark_strbuf *out, gfm_html_tree *tree, gfm_math_edit *edits, size_t id) {
  size_t current = id;
  while (current != GFM_NONE) {
    gfm_html_node *node = &tree->nodes[current];
    if (node->element && !gfm_node_is(tree, current, "template") && node->first != GFM_NONE) {
      current = node->first;
      continue;
    }
    if (!node->element) {
      cmark_strbuf decoded = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      houdini_unescape_html_f(&decoded, tree->source + node->start, node->end - node->start);
      size_t start = edits[current].trim_first && decoded.size ? 1 : 0;
      size_t end = decoded.size;
      if (edits[current].trim_last && end > start) end--;
      cmark_strbuf_put(out, decoded.ptr + start, end - start);
      cmark_strbuf_free(&decoded);
    }
    while (current != id && tree->nodes[current].next == GFM_NONE) current = tree->nodes[current].parent;
    if (current == id) break;
    current = tree->nodes[current].next;
  }
}
static void gfm_math_escape(cmark_strbuf *out, const unsigned char *text, size_t length, int protect_angles) {
  for (size_t i = 0; i < length; i++) {
    if (text[i] == '&') cmark_strbuf_puts(out, "&amp;");
    else if (text[i] == '<') cmark_strbuf_puts(out, protect_angles ? "&amp;lt;" : "&lt;");
    else if (text[i] == '>') cmark_strbuf_puts(out, protect_angles ? "&amp;gt;" : "&gt;");
    else if (text[i] == 0xc2 && i + 1 < length && text[i + 1] == 0xa0) { cmark_strbuf_puts(out, "&nbsp;"); i++; }
    else cmark_strbuf_putc(out, text[i]);
  }
}
static void gfm_math_element(cmark_strbuf *out, const unsigned char *text, size_t length, int display, int protect_angles, const char *token) {
  cmark_strbuf_puts(out, display ? "<math-renderer class=\"js-display-math\" style=\"display: block\" data-run-id=\"" : "<math-renderer class=\"js-inline-math\" style=\"display: inline-block\" data-run-id=\"");
  cmark_strbuf_puts(out, token); cmark_strbuf_puts(out, "\">");
  gfm_math_escape(out, text, length, protect_angles);
  cmark_strbuf_puts(out, "</math-renderer>");
}
static int gfm_math_closest(gfm_html_tree *tree, size_t id, const char *names) {
  for (size_t p = id; p != GFM_NONE; p = tree->nodes[p].parent) {
    const char *name = names;
    while (*name) { if (gfm_node_is(tree, p, name)) return 1; name += strlen(name) + 1; }
  }
  return 0;
}
static int gfm_math_blocked(gfm_html_tree *tree, size_t id) {
  for (size_t p = tree->nodes[id].parent; p != GFM_NONE; p = tree->nodes[p].parent) {
    if (gfm_node_is(tree, p, "details")) return 0;
    if (gfm_node_is(tree,p,"b") || gfm_node_is(tree,p,"strong") || gfm_node_is(tree,p,"em") || gfm_node_is(tree,p,"a") || gfm_node_is(tree,p,"code") || gfm_node_is(tree,p,"pre") || gfm_node_is(tree,p,"section")) return 1;
  }
  return 0;
}
static int gfm_math_target(gfm_html_tree *tree, size_t id) {
  static const char *names[] = {"p","li","h1","h2","h3","h4","h5","h6","td","th","strong","em","i","del","span","summary"};
  for (size_t i=0;i<sizeof(names)/sizeof(names[0]);i++) if (gfm_node_is(tree,id,names[i])) return 1;
  return 0;
}
static int gfm_math_replaced_ancestor(gfm_html_tree *tree, gfm_math_edit *edits, size_t id) {
  for (size_t p=tree->nodes[id].parent;p!=GFM_NONE;p=tree->nodes[p].parent)
    if (edits[p].inner || edits[p].outer) return 1;
  return 0;
}
static char *gfm_math_finish(cmark_strbuf *buffer) { return (char *)cmark_strbuf_detach(buffer); }

static int gfm_math_apply(gfm_html_tree *tree, gfm_math_edit *edits) {
  for (size_t id=1;id<tree->count;id++) {
    if (tree->nodes[id].inert || gfm_math_replaced_ancestor(tree,edits,id) || !gfm_math_target(tree,id)) continue;
    if (gfm_math_blocked(tree,id) && !(gfm_node_is(tree,id,"strong") && !gfm_math_closest(tree,id,"b\0section\0a\0pre\0code\0\0"))) continue;
    int only_text_br=1, only_text_br_em=1, has_em=0;
    cmark_strbuf reconstructed = CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    for (size_t child=tree->nodes[id].first;child!=GFM_NONE;child=tree->nodes[child].next) {
      if (!tree->nodes[child].element) gfm_math_text(&reconstructed,tree,edits,child);
      else if (gfm_node_is(tree,child,"em")) { only_text_br=0; has_em=1; cmark_strbuf_putc(&reconstructed,'_'); gfm_math_text(&reconstructed,tree,edits,child); cmark_strbuf_putc(&reconstructed,'_'); }
      else if (!gfm_node_is(tree,child,"br")) { only_text_br=0; only_text_br_em=0; }
    }
    if (has_em && only_text_br_em && !gfm_math_closest(tree,id,"b\0\0")) {
      char *display=gfm_emphasis_display(reconstructed.ptr,reconstructed.size);
      if (display) {
        cmark_strbuf out=CMARK_BUF_INIT(cmark_get_default_mem_allocator());
        size_t display_length = strlen(display);
        gfm_math_element(&out,(unsigned char*)display,display_length,1,0,tree->token);
        free(display); edits[id].inner=gfm_math_finish(&out); cmark_strbuf_free(&reconstructed); continue;
      }
    }
    cmark_strbuf_free(&reconstructed);
    if (only_text_br && !gfm_node_is(tree,id,"li") && !gfm_math_closest(tree,id,"b\0details\0\0")) {
      cmark_strbuf text=CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      gfm_math_text(&text,tree,edits,id);
      gfm_math_matches *blocks=gfm_display_math(text.ptr,text.size);
      if (!blocks) { cmark_strbuf_free(&text); return 0; }
      if (blocks->count) {
        cmark_strbuf out=CMARK_BUF_INIT(cmark_get_default_mem_allocator());
        for(size_t i=0;i<blocks->count;i++) {
          size_t start=blocks->ranges[i*2],end=blocks->ranges[i*2+1];
          cmark_strbuf value=CMARK_BUF_INIT(cmark_get_default_mem_allocator());
          gfm_math_put_lines(&value,text.ptr+start,end-start);
          if(i)cmark_strbuf_putc(&out,'\n');
          gfm_math_element(&out,value.ptr,value.size,1,0,tree->token); cmark_strbuf_free(&value);
        }
        edits[id].inner=gfm_math_finish(&out);
      }
      gfm_inline_math_free(blocks); cmark_strbuf_free(&text);
      if(edits[id].inner)continue;
    }
    for(size_t code=tree->nodes[id].first;code!=GFM_NONE;code=tree->nodes[code].next) {
      if(!gfm_node_is(tree,code,"code"))continue;
      size_t before=tree->nodes[code].prev,after=tree->nodes[code].next;
      if(before==GFM_NONE||after==GFM_NONE||tree->nodes[before].element||tree->nodes[after].element)continue;
      cmark_strbuf a=CMARK_BUF_INIT(cmark_get_default_mem_allocator()), b=CMARK_BUF_INIT(cmark_get_default_mem_allocator()), c=CMARK_BUF_INIT(cmark_get_default_mem_allocator());
      gfm_math_text(&a,tree,edits,before);gfm_math_text(&b,tree,edits,after);gfm_math_text(&c,tree,edits,code);
      char *value=gfm_protected_math(a.ptr,a.size,b.ptr,b.size,c.ptr,c.size,gfm_node_is(tree,id,"em"));
      if(value){cmark_strbuf out=CMARK_BUF_INIT(cmark_get_default_mem_allocator());size_t value_length = strlen(value);gfm_math_element(&out,(unsigned char*)value,value_length,0,0,tree->token);free(value);edits[code].outer=gfm_math_finish(&out);edits[before].trim_last=1;edits[after].trim_first=1;}
      cmark_strbuf_free(&a);cmark_strbuf_free(&b);cmark_strbuf_free(&c);
    }
    if(gfm_node_is(tree,id,"em") || (gfm_node_is(tree,id,"strong") && gfm_math_closest(tree,id,"em\0\0")))continue;
    for(size_t child=tree->nodes[id].first;child!=GFM_NONE;child=tree->nodes[child].next){
      if(tree->nodes[child].element)continue;
      cmark_strbuf text=CMARK_BUF_INIT(cmark_get_default_mem_allocator());gfm_math_text(&text,tree,edits,child);
      gfm_math_matches *matches=gfm_inline_math(text.ptr,text.size);if(!matches){cmark_strbuf_free(&text);return 0;}
      if(matches->count){cmark_strbuf out=CMARK_BUF_INIT(cmark_get_default_mem_allocator());size_t previous=0;
        for(size_t i=0;i<matches->count;i++){size_t start=matches->ranges[i*2],end=matches->ranges[i*2+1];gfm_math_escape(&out,text.ptr+previous,start-previous,0);gfm_math_element(&out,text.ptr+start,end-start,0,1,tree->token);previous=end;}
        gfm_math_escape(&out,text.ptr+previous,text.size-previous,0);edits[child].outer=gfm_math_finish(&out);
      }
      gfm_inline_math_free(matches);cmark_strbuf_free(&text);
    }
  }
  for(size_t id=1;id<tree->count;id++){
    if(tree->nodes[id].inert||gfm_math_replaced_ancestor(tree,edits,id)||!gfm_node_is(tree,id,"pre")||!gfm_attribute_equals(tree,id,"lang","math")||gfm_math_blocked(tree,id))continue;
    cmark_strbuf text=CMARK_BUF_INIT(cmark_get_default_mem_allocator()),out=CMARK_BUF_INIT(cmark_get_default_mem_allocator());
    gfm_math_text(&text,tree,edits,id);char *value=gfm_fenced_math(text.ptr,text.size);if(!value){cmark_strbuf_free(&text);cmark_strbuf_free(&out);return 0;}
    size_t value_length = strlen(value);gfm_math_element(&out,(unsigned char*)value,value_length,1,0,tree->token);free(value);edits[id].outer=gfm_math_finish(&out);cmark_strbuf_free(&text);
  }
  return 1;
}

char *gfm_math_dom(const unsigned char *source,size_t length,const char *token,const char *boundary) {
  gfm_html_tree tree={.source=source,.length=length,.token=token,.text_boundary=boundary};
  if(!gfm_parse_serialized_html(&tree))return NULL;
  gfm_math_edit *edits=calloc(tree.count,sizeof(*edits));if(!edits){free(tree.nodes);return NULL;}
  int ok=gfm_math_apply(&tree,edits),changed=tree.had_text_boundary;
  cmark_strbuf out=CMARK_BUF_INIT(cmark_get_default_mem_allocator());
  if(ok){size_t id=tree.nodes[0].first;while(id!=GFM_NONE){gfm_html_node *node=&tree.nodes[id];
    if(edits[id].outer){cmark_strbuf_puts(&out,edits[id].outer);changed=1;}
    else if(!node->element){if(edits[id].trim_first||edits[id].trim_last){cmark_strbuf text=CMARK_BUF_INIT(cmark_get_default_mem_allocator());gfm_math_text(&text,&tree,edits,id);gfm_math_escape(&out,text.ptr,text.size,0);cmark_strbuf_free(&text);changed=1;}else cmark_strbuf_put(&out,source+node->start,node->end-node->start);}
    else cmark_strbuf_put(&out,source+node->start,node->open_end-node->start);
    if(!edits[id].outer && edits[id].inner){cmark_strbuf_puts(&out,edits[id].inner);changed=1;}
    if(!edits[id].outer&&!edits[id].inner&&node->first!=GFM_NONE){id=node->first;continue;}
    while(id!=0){node=&tree.nodes[id];if(node->element&&!edits[id].outer&&node->end>node->close_start)cmark_strbuf_put(&out,source+node->close_start,node->end-node->close_start);if(node->next!=GFM_NONE){id=node->next;break;}id=node->parent;}if(id==0)break;
  }}
  for(size_t i=0;i<tree.count;i++){free(edits[i].inner);free(edits[i].outer);}free(edits);free(tree.nodes);
  if(!ok){cmark_strbuf_free(&out);return NULL;}
  if(!changed){cmark_strbuf_free(&out);cmark_strbuf_init(cmark_get_default_mem_allocator(),&out,0);cmark_strbuf_put(&out,source,length);}
  return gfm_math_finish(&out);
}
