package github

import "golang.org/x/net/html"

const (
	htmlTags = "a abbr acronym address area article aside audio b bdi bdo big blink blockquote br button canvas caption center cite code col colgroup " +
		"content data datalist dd decorator del details dfn dialog dir div dl dt element em fieldset figcaption figure font footer form " +
		"h1 h2 h3 h4 h5 h6 header hgroup hr i img input ins kbd label legend li main map mark marquee menu menuitem meter nav nobr " +
		"ol optgroup option output p picture pre progress q rp rt ruby s samp search section select shadow slot small source spacer span " +
		"strike strong sub summary sup table tbody td template textarea tfoot th thead time tr track tt u ul var video wbr"
	mathTags = "math menclose merror mfenced mfrac mglyph mi mlabeledtr mmultiscripts mn mo mover mpadded mphantom mprescripts mroot mrow ms " +
		"mspace msqrt mstyle msub msubsup msup mtable mtd mtext mtr munder munderover"
	dropContents = "annotation annotation-xml base embed foreignobject frame frameset head iframe link meta noembed noframes noscript object " +
		"plaintext script style svg title xmp"
	htmlAttrs = "accept action align alt autocapitalize autocomplete autopictureinpicture autoplay background bgcolor border capture cellpadding cellspacing " +
		"checked cite class clear color cols colspan controls controlslist coords crossorigin datetime decoding default dir disabled disablepictureinpicture " +
		"disableremoteplayback download draggable enctype enterkeyhint exportparts face for headers height hidden high href hreflang id inert inputmode " +
		"integrity ismap kind label lang list loading loop low max maxlength media method min minlength multiple muted name nonce noshade novalidate nowrap " +
		"open optimum part pattern placeholder playsinline popover popovertarget popovertargetaction poster preload pubdate radiogroup readonly rel required " +
		"rev reversed role rows rowspan scope selected shape size sizes slot span spellcheck srclang start step src srcset summary tabindex title translate " +
		"type usemap valign value width wrap"
	mathAttrs = "accent accentunder bevelled close columnalign columnlines columnspacing columnspan denomalign depth display displaystyle fence frame largeop length " +
		"linethickness lquote lspace mathbackground mathcolor mathsize mathvariant maxsize minsize movablelimits notation numalign rowalign rowlines rowspacing " +
		"rspace rquote scriptlevel scriptminsize scriptsizemultiplier selection separator separators stretchy subscriptshift superscriptshift symmetric voffset"
	clobberNames = "__proto__ constructor prototype attributes childNodes children nodeName nodeType ownerDocument parentNode innerHTML outerHTML textContent " +
		"appendChild removeChild insertBefore replaceChild cloneNode documentElement body head forms images links anchors scripts cookie domain location " +
		"URL documentURI baseURI defaultView write writeln open close getElementById getElementsByName getElementsByTagName querySelector querySelectorAll " +
		"createElement createAttribute createTextNode createDocumentFragment importNode adoptNode implementation style dataset classList className id name " +
		"action method submit reset elements length target encoding enctype acceptCharset checkValidity reportValidity requestSubmit onload onerror"
	maxDOMBufferBytes = 64 * 1024 * 1024
)

func wordList(words, value string) bool {
	for len(words) != 0 {
		n := 0
		for n < len(words) && words[n] != ' ' {
			n++
		}
		if n == len(value) && words[:n] == value {
			return true
		}
		if n == len(words) {
			break
		}
		words = words[n+1:]
	}
	return false
}

func sanitizeEscape(out *Buffer, value string, attribute bool) {
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case '&':
			out.Put("&amp;")
		case '<':
			out.Put("&lt;")
		case '>':
			out.Put("&gt;")
		case '"':
			if attribute {
				out.Put("&quot;")
			} else {
				out.Putc(value[i])
			}
		default:
			out.Putc(value[i])
		}
	}
}

func urlSpacing(value string) int {
	if len(value) == 0 {
		return 0
	}
	if value[0] <= 0x20 || value[0] == 0x7f {
		return 1
	}
	if len(value) >= 2 && value[0] == 0xc2 && value[1] == 0xa0 {
		return 2
	}
	if len(value) >= 3 &&
		((value[0] == 0xe1 && ((value[1] == 0x9a && value[2] == 0x80) || (value[1] == 0xa0 && value[2] == 0x8e))) ||
			(value[0] == 0xe2 && ((value[1] == 0x80 && value[2] >= 0x80 && value[2] <= 0xa9) || (value[1] == 0x81 && value[2] == 0x9f))) ||
			(value[0] == 0xe3 && value[1] == 0x80 && value[2] == 0x80) ||
			(value[0] == 0xef && value[1] == 0xbb && value[2] == 0xbf)) {
		return 3
	}
	return 0
}

func safeURL(value string, imageData bool) bool {
	var scheme [64]byte
	used := 0
	for i := 0; i < len(value); {
		if space := urlSpacing(value[i:]); space != 0 {
			i += space
			continue
		}
		c := value[i]
		i++
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if used == 0 && (c < 'a' || c > 'z') {
			return true
		}
		if c == ':' {
			if wordList("http https ftp ftps mailto tel callto sms cid xmpp matrix", string(scheme[:used])) {
				return true
			}
			if imageData && string(scheme[:used]) == "data" {
				mime := value[i:]
				size := 0
				for size < len(mime) && mime[size] != ';' && mime[size] != ',' {
					size++
				}
				return wordList("image/png image/jpeg image/gif image/webp image/avif image/bmp image/x-icon image/svg+xml", mime[:size])
			}
			return false
		}
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '+' || c == '-' || c == '.') {
			return true
		}
		if used == len(scheme) {
			return false
		}
		scheme[used] = c
		used++
	}
	return true
}

func customAttribute(name string) bool {
	prefix := 0
	if len(name) > 5 && (name[:5] == "data-" || name[:5] == "aria-") {
		prefix = 5
	}
	if prefix == 0 {
		return false
	}
	for i := prefix; i < len(name); i++ {
		c := name[i]
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-' || c == '_' || c == '.' || c >= 0x80) {
			return false
		}
	}
	return true
}

func safeSrcset(value string) bool {
	for i := 0; i < len(value); {
		for i < len(value) && (value[i] <= 32 || value[i] == ',') {
			i++
		}
		start := i
		for i < len(value) && value[i] > 32 && value[i] != ',' {
			i++
		}
		if i == start || !safeURL(value[start:i], false) {
			return false
		}
		if i < len(value) && value[i] == ',' {
			i++
			continue
		}
		for i < len(value) && value[i] <= 32 {
			i++
		}
		start = i
		for i < len(value) && value[i] != ',' {
			i++
		}
		end := i
		for end > start && value[end-1] <= 32 {
			end--
		}
		if end > start {
			suffix := value[end-1]
			if suffix != 'w' && suffix != 'x' || end == start+1 {
				return false
			}
			for n := start; n+1 < end; n++ {
				if (value[n] < '0' || value[n] > '9') && value[n] != '.' {
					return false
				}
			}
		}
		if i < len(value) {
			i++
		}
	}
	return true
}

func safeAttribute(node *html.Node, name, value string) bool {
	custom := customAttribute(name)
	if !custom && !wordList(htmlAttrs, name) && !(node.Namespace == "math" && wordList(mathAttrs, name)) {
		return false
	}
	if (name == "id" || name == "name") && wordList(clobberNames, value) {
		return false
	}
	if name == "srcset" {
		return safeSrcset(value)
	}
	if custom || wordList("alt class for id label name pattern placeholder role summary title value", name) {
		return true
	}
	return safeURL(value, node.Data == "img" && name == "src")
}

type sanitizePolicy struct {
	boundary      string
	legacyStopped bool
}

// sanitizeNode deliberately emits one stream instead of another parsed tree;
// this preserves the legacy comment boundaries used by the math renderer.
func sanitizeNode(out *Buffer, node *html.Node, depth int, policy *sanitizePolicy) bool {
	if len(out.Data) > maxDOMBufferBytes {
		return false
	}
	if depth > 512 {
		return false
	}
	switch node.Type {
	case html.TextNode:
		sanitizeEscape(out, node.Data, false)
		return len(out.Data) <= maxDOMBufferBytes
	case html.CommentNode:
		if policy.legacyStopped && policy.boundary != "" {
			out.Put("<!--")
			out.Put(policy.boundary)
			out.Put("-->")
		}
		policy.legacyStopped = true
		return len(out.Data) <= maxDOMBufferBytes
	case html.ElementNode:
		// Handled below.
	default:
		return true
	}

	isHTML := node.Namespace == ""
	isMath := node.Namespace == "math"
	if !isHTML && !isMath {
		policy.legacyStopped = true
		return true
	}
	allowed := isHTML && wordList(htmlTags, node.Data) || isMath && wordList(mathTags, node.Data)
	if !allowed {
		policy.legacyStopped = true
	}
	if !allowed && (!isHTML || wordList(dropContents, node.Data)) {
		return true
	}
	if allowed {
		out.Putc('<')
		out.Put(node.Data)
		for _, attr := range node.Attr {
			if attr.Namespace != "" {
				continue
			}
			value := attr.Val
			if attr.Key != "value" {
				for len(value) != 0 && value[0] <= 32 {
					value = value[1:]
				}
				for len(value) != 0 && value[len(value)-1] <= 32 {
					value = value[:len(value)-1]
				}
			}
			if !safeAttribute(node, attr.Key, value) {
				continue
			}
			out.Putc(' ')
			out.Put(attr.Key)
			out.Put("=\"")
			sanitizeEscape(out, value, true)
			out.Putc('"')
		}
		out.Putc('>')
	}

	childrenPolicy := policy
	if isHTML && node.Data == "template" {
		childrenPolicy = &sanitizePolicy{boundary: policy.boundary}
	}
	for child := domChildren(node); child != nil; child = child.NextSibling {
		if !sanitizeNode(out, child, depth+1, childrenPolicy) {
			return false
		}
	}
	if allowed && !(isHTML && wordList("area br col hr img input source track wbr", node.Data)) {
		out.Put("</")
		out.Put(node.Data)
		out.Putc('>')
	}
	return len(out.Data) <= maxDOMBufferBytes
}

func sanitizeDOM(root *html.Node, boundary string) (string, bool) {
	for i := range boundary {
		c := boundary[i]
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || c == '-') {
			return "", false
		}
	}
	var out Buffer
	policy := sanitizePolicy{boundary: boundary}
	for child := domChildren(root); child != nil; child = child.NextSibling {
		if !sanitizeNode(&out, child, 0, &policy) {
			return "", false
		}
	}
	if len(out.Data) > maxDOMBufferBytes {
		return "", false
	}
	return out.String(), true
}
