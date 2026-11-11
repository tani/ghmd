import { decodeHtml } from "../quirks/native-core";

// Experimental policy only: this is not a security-audited HTML sanitizer.
const htmlTags = new Set(
  `a abbr acronym address area article aside audio b bdi bdo big blink blockquote br button canvas caption center cite code col colgroup data datalist dd del details dfn dialog dir div dl dt em fieldset figcaption figure font footer form h1 h2 h3 h4 h5 h6 header hgroup hr i img input ins kbd label legend li main map mark marquee menu menuitem meter nav nobr ol optgroup option output p picture pre progress q rp rt ruby s samp search section select slot small source span strike strong sub summary sup table tbody td template textarea tfoot th thead time tr track tt u ul var video wbr`.split(
    " ",
  ),
);
const mathTags = new Set(
  `math menclose merror mfenced mfrac mglyph mi mlabeledtr mmultiscripts mn mo mover mpadded mphantom mroot mrow ms mspace msqrt mstyle msub msup msubsup mtable mtd mtext mtr munder munderover mprescripts`.split(
    " ",
  ),
);
const attributes = new Set(
  `accept action align alt autocapitalize autocomplete autoplay background bgcolor border capture cellpadding cellspacing checked cite class clear color cols colspan controls coords crossorigin datetime decoding default dir disabled download draggable enctype face for headers height hidden high href hreflang id inert inputmode ismap kind label lang list loading loop low max maxlength media method min minlength multiple muted name nowrap open optimum pattern placeholder playsinline poster preload readonly rel required reversed role rows rowspan scope selected shape size sizes slot span spellcheck src srclang start step summary tabindex title translate type usemap valign value width wrap accent accentunder bevelled close columnalign columnlines columnspacing columnspan denomalign depth display displaystyle encoding fence frame largeop length linethickness lquote lspace mathbackground mathcolor mathsize mathvariant maxsize minsize movablelimits notation numalign rowalign rowlines rowspacing rspace rquote scriptlevel scriptminsize scriptsizemultiplier separator separators stretchy subscriptshift superscriptshift symmetric voffset`.split(
    " ",
  ),
);
const dropContents = new Set(
  `script style iframe object embed noscript noembed noframes xmp plaintext annotation-xml title head`.split(
    " ",
  ),
);
const uriAttributes = new Set("action background cite href poster src usemap".split(" "));
const safeUri =
  /^(?:(?:(?:f|ht)tps?|mailto|tel|callto|sms|cid|xmpp|matrix):|[^a-z]|[a-z+.\-]+(?:[^a-z+.\-:]|$))/i;

/** Drop raw input tags; foreign-context handling may differ from parse5. */
export function removeRawInputs(html: string): string {
  if (!/<input(?:\s|\/?>)/i.test(html)) return html;
  return new HTMLRewriter()
    .on("input", {
      element(element) {
        element.removeAndKeepContent();
      },
    })
    .transform(html);
}

/** Filter the unsanitized HTML emitted by cmark with Bun's own tokenizer. */
export function sanitize(html: string, marker: string): string {
  return new HTMLRewriter()
    .onDocument({
      comments(comment) {
        comment.remove();
      },
    })
    .on("*", {
      element(element) {
        const tag = element.tagName.toLowerCase();
        const svg = element.namespaceURI === "http://www.w3.org/2000/svg";
        if (svg && tag !== "a") {
          element.removeAndKeepContent();
          return;
        }
        if (!htmlTags.has(tag) && !mathTags.has(tag)) {
          if (dropContents.has(tag)) element.remove();
          else element.removeAndKeepContent();
          return;
        }
        const hadHref = tag === "a" && element.hasAttribute("href");
        for (const [name, raw] of [...element.attributes]) {
          const attr = name.toLowerCase();
          if (
            (!attributes.has(attr) &&
              !/^data-[\w.:-]+$/.test(attr) &&
              !/^aria-[\w-]+$/.test(attr)) ||
            attr.startsWith("on") ||
            attr === "style" ||
            attr === "data-gfm-had-href"
          ) {
            element.removeAttribute(name);
            continue;
          }
          const value = decodeHtml(raw);
          if (
            (attr === "id" || attr === "name") &&
            /^(?:__proto__|constructor|prototype)$/i.test(value)
          ) {
            element.removeAttribute(name);
            continue;
          }
          if (uriAttributes.has(attr)) {
            const cleaned = value.replace(
              /[\u0000-\u0020\u00a0\u1680\u180e\u2000-\u2029\u205f\u3000]/g,
              "",
            );
            const dataImage = attr === "src" && tag === "img" && /^data:/i.test(cleaned);
            if (cleaned && !safeUri.test(cleaned) && !dataImage) element.removeAttribute(name);
          }
        }
        if (hadHref) element.setAttribute("data-gfm-had-href", marker);
      },
    })
    .transform(html);
}
