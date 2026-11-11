import { cc, CString, toArrayBuffer } from "bun:ffi";
import { join } from "node:path";
import { removeRawInputs } from "./raw-html";
const root = join(import.meta.dir, "../..");
const lib = cc({
  source: join(root, "src/quirks/cmark-unity.c"),
  include: ["src", "extensions", "generated"].map((dir) => join(root, "vendor/cmark-gfm", dir)),
  symbols: {
    cmark_version_string: { args: [], returns: "cstring" },
    cmark_gfm_core_extensions_ensure_registered: { args: [], returns: "void" },
    gfm_document_new: { args: ["buffer", "u64", "i32"], returns: "ptr" },
    gfm_document_raw_count: { args: ["ptr"], returns: "u64" },
    gfm_document_raw_node: { args: ["ptr", "u64"], returns: "ptr" },
    gfm_document_html: { args: ["ptr"], returns: "ptr" },
    gfm_document_xml: { args: ["ptr"], returns: "ptr" },
    gfm_document_free: { args: ["ptr"], returns: "void" },
    gfm_finalize_serialized_html: { args: ["buffer", "u64", "buffer", "buffer"], returns: "ptr" },
    gfm_math_dom: { args: ["buffer", "u64", "buffer", "buffer"], returns: "ptr" },
    gfm_unicode_lower: { args: ["buffer", "u64"], returns: "ptr" },
    gfm_unicode_lower_length: { args: ["ptr"], returns: "u64" },
    gfm_unicode_lower_data: { args: ["ptr"], returns: "ptr" },
    cmark_node_get_literal: { args: ["ptr"], returns: "cstring" },
    cmark_node_set_literal: { args: ["ptr", "buffer"], returns: "i32" },
    gfm_free: { args: ["ptr"], returns: "void" },
  },
});
const s = lib.symbols;
s.cmark_gfm_core_extensions_ensure_registered();
export const nativeVersion = String(s.cmark_version_string());
function adaptRawHtml(document: any) {
  const count = Number(s.gfm_document_raw_count(document));
  for (let i = 0; i < count; i++) {
    const node = s.gfm_document_raw_node(document, i);
    const original = String(s.cmark_node_get_literal(node));
    const filtered = removeRawInputs(original);
    if (filtered !== original && !s.cmark_node_set_literal(node, Buffer.from(filtered + "\0")))
      throw Error("Cannot remove raw input elements");
  }
}
// Explicit byte length handles embedded NUL and non-ASCII input correctly.
export function nativeRender(source: string): { html: string; astXML: string } {
  const options = (1 << 17) | (1 << 13) | (1 << 11) | (1 << 2); // Raw HTML is sanitized later.
  const bytes = Buffer.from(source, "utf8");
  const document = s.gfm_document_new(bytes, bytes.byteLength, options);
  if (!document) throw Error("Cannot parse Markdown");
  const owned = (p: any) => {
    if (!p) throw Error("cmark render allocation failed");
    try {
      return String(new CString(p));
    } finally {
      s.gfm_free(p);
    }
  };
  try {
    adaptRawHtml(document);
    return {
      html: owned(s.gfm_document_html(document)),
      astXML: owned(s.gfm_document_xml(document)),
    };
  } finally {
    s.gfm_document_free(document);
  }
}

/** Apply table, list, code and math-spacing adjustments to serialized, sanitized HTML. */
export function finalizeHtml(
  html: string,
  token = "00000000000000000000000000000000",
  tableMarker = "",
): string {
  const bytes = Buffer.from(html, "utf8");
  const result = s.gfm_finalize_serialized_html(
    bytes,
    bytes.byteLength,
    Buffer.from(token + "\0"),
    Buffer.from(tableMarker + "\0"),
  );
  if (!result) throw Error("Cannot finalize HTML");
  try {
    return String(new CString(result));
  } finally {
    s.gfm_free(result);
  }
}

function consumeNativeString(pointer: any): string | undefined {
  if (!pointer) return;
  try {
    return String(new CString(pointer));
  } finally {
    s.gfm_free(pointer);
  }
}
/** Unicode 17 default lowercase, including context-sensitive Greek sigma. */
export function unicodeLower(text: string): string {
  const input = Buffer.from(text, "utf8");
  const result = s.gfm_unicode_lower(input, input.byteLength);
  if (!result) throw Error("Cannot lowercase Unicode text");
  try {
    const size = Number(s.gfm_unicode_lower_length(result));
    return size
      ? new TextDecoder().decode(toArrayBuffer(s.gfm_unicode_lower_data(result), 0, size))
      : "";
  } finally {
    s.gfm_free(result);
  }
}

/** Transform math in sanitized HTML. */
export function renderMathHtml(html: string, token: string, boundary: string): string {
  const bytes = Buffer.from(html, "utf8");
  const result = consumeNativeString(
    s.gfm_math_dom(
      bytes,
      bytes.byteLength,
      Buffer.from(token + "\0"),
      Buffer.from(boundary + "\0"),
    ),
  );
  if (result === undefined) throw Error("Cannot transform math HTML");
  return result;
}
