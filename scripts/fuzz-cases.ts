/** Deterministic, independently reproducible combinations of risky constructs. */
export function fuzzCase(seed: number, index: number): string {
  let state = (seed ^ Math.imul(index + 1, 0x9e3779b1)) >>> 0;
  if (state === 0) state = 0x6d2b79f5;
  const random = (length: number) => {
    state ^= state << 13;
    state ^= state >>> 17;
    state ^= state << 5;
    return (state >>> 0) % length;
  };
  const math = [
    "$x$",
    "a$x$",
    "$`a < b`$",
    "$$\nx_1 + y\n$$",
    "$a <b > c$",
    "`$x$`",
    "$$\n\\text{α}\\\\\nβ\n$$",
  ];
  const unicode = ["日本語", "ΟΣ", "İẞ", "😀e\u0301", "中文**強調**", "a\u00a0b", "\u0345Σ"];
  const html = [
    "<b>unclosed",
    '<a href="javascript:alert(1)">x',
    "<span>$x$<!--x-->$y$</span>",
    "<details><summary>x</summary>",
    "<table><tr><td>one|two",
    "<template><code>$x$</code></template>",
    "<input type=checkbox checked>",
  ];
  const marker = ["-", "+", "1."];
  const depth = 8 + random(6);
  const list = Array.from(
    { length: depth },
    (_, i) =>
      `${"  ".repeat(i)}${marker[random(marker.length)]} ${i === depth - 1 && random(2) ? "[x] " : ""}${i === depth - 1 ? math[random(math.length)] : unicode[random(unicode.length)]}`,
  ).join("\n");
  const blocks = [
    math[random(math.length)],
    unicode[random(unicode.length)],
    html[random(html.length)],
    list,
  ];
  // Different boundaries and nesting force features to interact rather than
  // merely repeating the independent reference fixtures.
  for (let i = blocks.length - 1; i > 0; i--) {
    const j = random(i + 1);
    [blocks[i], blocks[j]] = [blocks[j], blocks[i]];
  }
  const joiner = ["\n\n", "\n", "\n\n> "][random(3)];
  return blocks.join(joiner) + "\n";
}
