/* HTML/MathML allowlist data is adapted from DOMPurify 3.4.1 (Apache-2.0).
 * Copyright 2025-2026 Dr.-Ing. Mario Heiderich, Cure53. See NOTICE and
 * POLICY-LICENSE in this directory. The traversal, URL validation, policy
 * enforcement, serialization and compatibility handling are implemented here.
 *
 * Renderer policy over a Lexbor HTML5 tree. Unknown markup is unwrapped;
 * executable/inert foreign containers and their contents are discarded.
 * No HTML tokenization, entity decoding or tree repair is implemented here. */
static int h_equal(const unsigned char *s, size_t length, const char *word) {
  return strlen(word) == length && !memcmp(s, word, length);
}
static int h_list(const char *words, const unsigned char *s, size_t length) {
  for (const char *p = words; *p;) {
    const char *end = strchr(p, ' ');
    size_t n = end ? (size_t)(end - p) : strlen(p);
    if (n == length && !memcmp(s, p, n)) return 1;
    p += n; if (*p) p++;
  }
  return 0;
}
static const char h_html_tags[] =
  "a abbr acronym address area article aside audio b bdi bdo big blink blockquote br button canvas caption center cite code col colgroup "
  "content data datalist dd decorator del details dfn dialog dir div dl dt element em fieldset figcaption figure font footer form "
  "h1 h2 h3 h4 h5 h6 header hgroup hr i img input ins kbd label legend li main map mark marquee menu menuitem meter nav nobr "
  "ol optgroup option output p picture pre progress q rp rt ruby s samp search section select shadow slot small source spacer span "
  "strike strong sub summary sup table tbody td template textarea tfoot th thead time tr track tt u ul var video wbr";
static const char h_math_tags[] =
  "math menclose merror mfenced mfrac mglyph mi mlabeledtr mmultiscripts mn mo mover mpadded mphantom mprescripts mroot mrow ms "
  "mspace msqrt mstyle msub msubsup msup mtable mtd mtext mtr munder munderover";
static const char h_drop_contents[] =
  "annotation annotation-xml base embed foreignobject frame frameset head iframe link meta noembed noframes noscript object "
  "plaintext script style svg title xmp";
static const char h_html_attrs[] =
  "accept action align alt autocapitalize autocomplete autopictureinpicture autoplay background bgcolor border capture cellpadding cellspacing "
  "checked cite class clear color cols colspan controls controlslist coords crossorigin datetime decoding default dir disabled disablepictureinpicture "
  "disableremoteplayback download draggable enctype enterkeyhint exportparts face for headers height hidden high href hreflang id inert inputmode "
  "integrity ismap kind label lang list loading loop low max maxlength media method min minlength multiple muted name nonce noshade novalidate nowrap "
  "open optimum part pattern placeholder playsinline popover popovertarget popovertargetaction poster preload pubdate radiogroup readonly rel required "
  "rev reversed role rows rowspan scope selected shape size sizes slot span spellcheck srclang start step src srcset summary tabindex title translate "
  "type usemap valign value width wrap";
static const char h_math_attrs[] =
  "accent accentunder bevelled close columnalign columnlines columnspacing columnspan denomalign depth display displaystyle fence frame largeop length "
  "linethickness lquote lspace mathbackground mathcolor mathsize mathvariant maxsize minsize movablelimits notation numalign rowalign rowlines rowspacing "
  "rspace rquote scriptlevel scriptminsize scriptsizemultiplier selection separator separators stretchy subscriptshift superscriptshift symmetric voffset";
static const char h_clobber_names[] =
  "__proto__ constructor prototype attributes childNodes children nodeName nodeType ownerDocument parentNode innerHTML outerHTML textContent "
  "appendChild removeChild insertBefore replaceChild cloneNode documentElement body head forms images links anchors scripts cookie domain location "
  "URL documentURI baseURI defaultView write writeln open close getElementById getElementsByName getElementsByTagName querySelector querySelectorAll "
  "createElement createAttribute createTextNode createDocumentFragment importNode adoptNode implementation style dataset classList className id name "
  "action method submit reset elements length target encoding enctype acceptCharset checkValidity reportValidity requestSubmit onload onerror";
static void h_escape(h_buffer *out, const unsigned char *value, size_t length, int attribute) {
  for (size_t i = 0; i < length; i++) {
    unsigned char c = value[i];
    if (c == '&') h_s(out, "&amp;");
    else if (c == '<') h_s(out, "&lt;");
    else if (c == '>') h_s(out, "&gt;");
    else if (attribute && c == '"') h_s(out, "&quot;");
    else h_put(out, value + i, 1);
  }
}
/* URL checks use decoded values. Remove ASCII controls and Unicode spacing
 * before recognizing a scheme, including entity-obfuscated protocol names. */
static size_t h_spacing(const unsigned char *s, size_t length) {
  if (!length) return 0;
  if (s[0] <= 0x20 || s[0] == 0x7f) return 1;
  if (length >= 2 && s[0] == 0xc2 && s[1] == 0xa0) return 2;
  if (length >= 3 && ((s[0] == 0xe1 && ((s[1] == 0x9a && s[2] == 0x80) || (s[1] == 0xa0 && s[2] == 0x8e))) ||
      (s[0] == 0xe2 && ((s[1] == 0x80 && s[2] >= 0x80 && s[2] <= 0xa9) || (s[1] == 0x81 && s[2] == 0x9f))) ||
      (s[0] == 0xe3 && s[1] == 0x80 && s[2] == 0x80) || (s[0] == 0xef && s[1] == 0xbb && s[2] == 0xbf))) return 3;
  return 0;
}
static int h_url(const unsigned char *value, size_t length, int image_data) {
  /* A scheme fits this bounded buffer; oversized/invalid schemes fail closed. */
  unsigned char scheme[64]; size_t used = 0;
  for (size_t i = 0; i < length;) {
    size_t space = h_spacing(value + i, length - i);
    if (space) { i += space; continue; }
    unsigned char c = value[i++];
    if (c >= 'A' && c <= 'Z') c += 32;
    if (!used && (c < 'a' || c > 'z')) return 1;
    if (c == ':') {
      if (h_list("http https ftp ftps mailto tel callto sms cid xmpp matrix", scheme, used)) return 1;
      if (image_data && h_equal(scheme, used, "data")) {
        const unsigned char *mime = value + i; size_t size = 0;
        while (i + size < length && mime[size] != ';' && mime[size] != ',') size++;
        return h_list("image/png image/jpeg image/gif image/webp image/avif image/bmp image/x-icon image/svg+xml", mime, size);
      }
      return 0;
    }
    if (!((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '+' || c == '-' || c == '.')) return 1;
    if (used == sizeof(scheme)) return 0;
    scheme[used++] = c;
  }
  return 1;
}
static int h_custom_attr(const unsigned char *name, size_t length) {
  size_t prefix = length > 5 && !memcmp(name, "data-", 5) ? 5 : length > 5 && !memcmp(name, "aria-", 5) ? 5 : 0;
  if (!prefix) return 0;
  for (size_t i = prefix; i < length; i++) {
    unsigned char c = name[i];
    if (!((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.' || c >= 0x80)) return 0;
  }
  return 1;
}
/* srcset has comma-separated candidates. Data URLs and malformed descriptors
 * are deliberately rejected here; ordinary relative/http images are retained. */
static int h_srcset(const unsigned char *value, size_t length) {
  size_t i = 0;
  while (i < length) {
    while (i < length && (value[i] <= 32 || value[i] == ',')) i++;
    size_t start = i;
    while (i < length && value[i] > 32 && value[i] != ',') i++;
    if (i == start || !h_url(value + start, i - start, 0)) return 0;
    if (i < length && value[i] == ',') { i++; continue; }
    while (i < length && value[i] <= 32) i++;
    start = i;
    while (i < length && value[i] != ',') i++;
    size_t end = i; while (end > start && value[end - 1] <= 32) end--;
    if (end > start) {
      unsigned char suffix = value[end - 1];
      if (suffix != 'w' && suffix != 'x') return 0;
      if (end == start + 1) return 0;
      for (size_t n = start; n + 1 < end; n++) if ((value[n] < '0' || value[n] > '9') && value[n] != '.') return 0;
    }
    if (i < length) i++;
  }
  return 1;
}
static int h_attribute(lxb_dom_node_t *node, const unsigned char *name, size_t size, const unsigned char *value, size_t length) {
  int custom = h_custom_attr(name, size);
  if (!custom && !h_list(h_html_attrs, name, size) && !(node->ns == LXB_NS_MATH && h_list(h_math_attrs, name, size))) return 0;
  if ((h_equal(name, size, "id") || h_equal(name, size, "name")) && h_list(h_clobber_names, value, length)) return 0;
  if (h_equal(name, size, "srcset")) return h_srcset(value, length);
  if (custom || h_list("alt class for id label name pattern placeholder role summary title value", name, size)) return 1;
  return h_url(value, length, node->local_name == LXB_TAG_IMG && h_equal(name, size, "src"));
}
/* The former DOM iterator stopped when its current node was removed. Later
 * comments therefore separated math text runs. Preserve only those harmless
 * boundaries with a renderer-generated nonce, while sanitizing EVERY node. */
typedef struct { const char *boundary; int legacy_stopped; } h_policy;
static void h_sanitize_node(h_buffer *out, lxb_dom_node_t *node, unsigned depth, h_policy *policy) {
  if (out->failed) return;
  if (depth > 512) { out->failed = H_DEPTH_EXCEEDED; return; }
  if (node->type == LXB_DOM_NODE_TYPE_TEXT) {
    lexbor_str_t *text = &((lxb_dom_character_data_t *)node)->data;
    h_escape(out, text->data, text->length, 0); return;
  }
  if (node->type == LXB_DOM_NODE_TYPE_COMMENT) {
    if (policy->legacy_stopped && policy->boundary[0]) { h_s(out, "<!--"); h_s(out, policy->boundary); h_s(out, "-->"); }
    policy->legacy_stopped = 1; return;
  }
  if (node->type != LXB_DOM_NODE_TYPE_ELEMENT) return;
  size_t length; const unsigned char *name = lxb_dom_element_local_name(lxb_dom_interface_element(node), &length);
  /* Foreign containers are not unwrapped: doing so changes HTML integration
   * points and can turn inert text into active markup on a later parse. */
  if (node->ns != LXB_NS_HTML && node->ns != LXB_NS_MATH) { policy->legacy_stopped = 1; return; }
  int allowed = node->ns == LXB_NS_HTML ? h_list(h_html_tags, name, length) : h_list(h_math_tags, name, length);
  if (!allowed) policy->legacy_stopped = 1;
  if (!allowed && (node->ns != LXB_NS_HTML || h_list(h_drop_contents, name, length))) return;
  if (allowed) {
    h_s(out, "<"); h_put(out, name, length);
    for (lxb_dom_attr_t *attr = lxb_dom_interface_element(node)->first_attr; attr; attr = attr->next) {
      if (attr->node.ns != LXB_NS__UNDEF && attr->node.ns != LXB_NS_HTML) continue;
      size_t size, value_length; const unsigned char *key = lxb_dom_attr_local_name(attr, &size);
      const unsigned char *value = lxb_dom_attr_value(attr, &value_length);
      /* Keep input values verbatim. Other attributes follow HTML sanitizer
       * trimming semantics; protocol validation is stricter than trimming. */
      if (!h_equal(key, size, "value")) {
        while (value_length && value[0] <= 32) { value++; value_length--; }
        while (value_length && value[value_length - 1] <= 32) value_length--;
      }
      if (!h_attribute(node, key, size, value, value_length)) continue;
      h_s(out, " "); h_put(out, key, size); h_s(out, "=\""); h_escape(out, value, value_length, 1); h_s(out, "\"");
    }
    h_s(out, ">");
  }
  h_policy template_policy = {policy->boundary, 0};
  h_policy *children_policy = node->ns == LXB_NS_HTML && node->local_name == LXB_TAG_TEMPLATE ? &template_policy : policy;
  for (lxb_dom_node_t *child = h_children(node); child; child = child->next) h_sanitize_node(out, child, depth + 1, children_policy);
  if (allowed && !(node->ns == LXB_NS_HTML && h_list("area br col hr img input source track wbr", name, length))) {
    h_s(out, "</"); h_put(out, name, length); h_s(out, ">");
  }
}
static void h_sanitize(h_buffer *out, h_document *document, lxb_dom_node_t *root, const char *boundary) {
  (void)document;
  /* Only UUID-like nonces may be emitted as internal comments. */
  for (const char *p = boundary; *p; p++) if (!((*p >= '0' && *p <= '9') || (*p >= 'a' && *p <= 'f') || *p == '-')) { out->failed = 1; return; }
  h_policy policy = {boundary, 0};
  for (lxb_dom_node_t *child = root->first_child; child; child = child->next) h_sanitize_node(out, child, 0, &policy);
}
