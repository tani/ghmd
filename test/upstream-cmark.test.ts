import { describe, expect, test } from "bun:test";
import { createHash } from "node:crypto";
import { readdir } from "node:fs/promises";
import { join } from "node:path";
import manifest from "./fixtures/cmark-gfm-upstream.json";
import { nativeRender } from "../src/quirks/native";

const vendor = join(import.meta.dir, "../vendor/cmark-gfm");
describe("unmodified upstream cmark-gfm", () => {
  test("release source and recorded generated headers match their pinned hashes", async () => {
    expect(manifest.commit).toBe("587a12bb54d95ac37241377e6ddc93ea0e45439b");
    const actual: string[] = [];
    for (const dir of ["src", "extensions", "generated"])
      for (const entry of await readdir(join(vendor, dir), { withFileTypes: true }))
        if (entry.isFile()) actual.push(`${dir}/${entry.name}`);
    actual.push("COPYING");
    expect(actual.sort()).toEqual(Object.keys(manifest.files).sort());
    for (const [path, expected] of Object.entries(manifest.files)) {
      const bytes = await Bun.file(join(vendor, path)).arrayBuffer();
      expect(createHash("sha256").update(Buffer.from(bytes)).digest("hex")).toBe(expected);
    }
  });
  test("adjusts task states without changing literal code", () => {
    const html = nativeRender("- [ ] `[x]`\n- [x] yes\n\n`[^a\\]b]`").html;
    expect(html).toContain('<input type="checkbox" disabled="" /> <code>[x]</code>');
    expect(html).toContain('<input type="checkbox" checked="" disabled="" /> yes');
    expect(html).toContain("<code>[^a\\]b]</code>");
  });
  test("retains references whose label precedes an image", () => {
    const html = nativeRender("[^a![x](#)]\n\n[^a]: note").html;
    expect(html).toContain("data-footnote-ref");
    expect(html).toContain("<p>note ");
  });
  test("keeps code containing multiline footnote syntax literal", () => {
    const html = nativeRender("[^] and `[^a\nb]`\n\n[^a]: note").html;
    expect(html).toContain("<code>[^a b]</code>");
  });
});
