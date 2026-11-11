/** Scope hints from Bun's tokenizer, not an HTML5 tree reconstruction. */
const formatting = new Set("a b big code del em font i nobr s small strike strong tt u".split(" "));
const blocks = new Set(
  "address article aside blockquote div dl dt dd h1 h2 h3 h4 h5 h6 li main ol p pre section table tbody td th thead tr ul".split(
    " ",
  ),
);
// These contexts can suppress subsequent native math. Closing them without
// reproducing the browser's formatting reconstruction can change rendering.
const mathBlocking = new Set("a b code em strong".split(" "));

export function annotateImplicitClosures(
  html: string,
  nonce: string,
  aggressive = false,
): { html: string; restore: (output: string) => string } {
  const synthetic = new Set<string>();
  const result = new HTMLRewriter()
    .on("*", {
      element(element) {
        const tag = element.tagName.toLowerCase();
        if (!formatting.has(tag)) return;
        element.onEndTag((end) => {
          const closing = end.name.toLowerCase();
          if (tag === closing) return;
          if (tag === "a" && blocks.has(closing)) {
            // GitHub ends reconstructed links at their enclosing block boundary.
            end.before("</a>", { html: true });
          } else if (formatting.has(closing) && (aggressive || !mathBlocking.has(tag))) {
            // Bound the native analysis scope while preserving the original
            // malformed markup for the browser's formatting reconstruction.
            const hint = `</${tag} data-gfm-scope-${nonce}>`;
            synthetic.add(hint);
            end.before(hint, { html: true });
          }
        });
      },
    })
    .transform(html);
  return {
    html: result,
    restore(output) {
      for (const hint of synthetic) output = output.replaceAll(hint, "");
      return output;
    },
  };
}
