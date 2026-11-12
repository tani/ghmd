/** Input: output of upstream `perl single.pl html` at the pinned commit. */
import { createHash } from "node:crypto";
const path = process.argv[2];
if (!path) throw Error("Usage: bun scripts/prepare-lexbor.ts path/to/html-single.h");
const original = await Bun.file(path).bytes();
const provenance = await Bun.file(
  new URL("../vendor/lexbor/PROVENANCE.json", import.meta.url),
).json();
const hash = createHash("sha256").update(original).digest("hex");
if (hash !== provenance.upstreamSha256)
  throw Error(
    "Input differs from the pinned upstream distribution; review PROVENANCE.json before upgrading",
  );
// The selected module includes float.h but references none of its limits.
if (/\b(?:FLT|DBL|LDBL)_[A-Z0-9_]+\b/.test(new TextDecoder().decode(original)))
  throw Error("Float limits now require a real compatibility header");
await Bun.write(new URL("../vendor/lexbor/lexbor.h", import.meta.url), original);
console.log(JSON.stringify({ upstreamSha256: hash }));
