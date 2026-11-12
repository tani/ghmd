/** Input: output of upstream `perl single.pl html` at the pinned commit. */
import { createHash } from "node:crypto";
const path = process.argv[2];
if (!path) throw Error("Usage: bun scripts/prepare-lexbor.ts path/to/html-single.h");
const original = await Bun.file(path).text();
const provenance = await Bun.file(
  new URL("../vendor/lexbor/PROVENANCE.json", import.meta.url),
).json();
const hash = (value: string) => createHash("sha256").update(value).digest("hex");
if (hash(original) !== provenance.upstreamSha256)
  throw Error(
    "Input differs from the pinned upstream distribution; review PROVENANCE.json before upgrading",
  );
const definition = "lxb_html_tree_create_element_for_token(lxb_html_tree_t *tree,";
if (original.split(definition).length !== 3)
  throw Error("Expected one declaration and one definition");
const index = original.lastIndexOf(definition);
const generated =
  original.slice(0, index) +
  original
    .slice(index)
    .replace(definition, "gfm_lexbor_create_element_original(lxb_html_tree_t *tree,");
// The selected module includes float.h but references none of its limits.
if (/\b(?:FLT|DBL|LDBL)_[A-Z0-9_]+\b/.test(original))
  throw Error("Float limits now require a real compatibility header");
if (hash(generated) !== provenance.generatedSha256)
  throw Error("Generated distribution differs from PROVENANCE.json");
await Bun.write(new URL("../vendor/lexbor/lexbor.h", import.meta.url), generated);
console.log(JSON.stringify({ upstreamSha256: hash(original), generatedSha256: hash(generated) }));
