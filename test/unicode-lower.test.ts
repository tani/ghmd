import { expect, test } from "bun:test";
import { unicodeLower } from "../src/quirks/native-core";
import { render } from "../src/quirks";

test("matches JavaScript lowercase for every Unicode scalar value", () => {
  let chunk = "";
  for (let cp = 0; cp <= 0x10ffff; cp++) {
    if (cp >= 0xd800 && cp <= 0xdfff) continue;
    // Separators prevent neighbouring scalars from affecting Greek sigma.
    chunk += String.fromCodePoint(cp) + "!";
    if (chunk.length >= 4096) {
      expect(unicodeLower(chunk)).toBe(chunk.toLowerCase());
      chunk = "";
    }
  }
  if (chunk) expect(unicodeLower(chunk)).toBe(chunk.toLowerCase());
});

test("handles contextual sigma, combining marks, supplementary letters and NUL", () => {
  for (const text of [
    "ΟΣ",
    "ΟΣΑ",
    "ΣΣ",
    "AΣ\u0301",
    "AΣ\u0301B",
    "\u0345Σ",
    "AΣ\u0345",
    "İẞ𐐀𐐧",
    "\0ABC\0",
    "",
  ]) {
    expect(unicodeLower(text)).toBe(text.toLowerCase());
  }
});

test("lowercases footnote IDs inside native HTML rendering", async () => {
  const runId = "0123456789abcdef0123456789abcdef";
  for (const id of ["ΟΣ", "ΟΣΑ", "AΣ\u0301", "AΣ\u0301B", "İẞ𐐀"]) {
    const html = await render(`<a id="${id}" data-footnote-ref>x</a>`, { runId });
    expect(html).toContain(`id="user-content-${id.toLowerCase()}-${runId}"`);
  }
});
