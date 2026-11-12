import { expect, test } from "bun:test";
import { readdir } from "node:fs/promises";
import { join } from "node:path";
import { render } from "../src/quirks";
import { renderDocument } from "../src/quirks/native-core";

const directory = join(import.meta.dir, "fixtures/native-pipeline");
for (const filename of (await readdir(directory))
  .filter((name) => name.endsWith(".md") && name !== "README.md")
  .sort()) {
  test(`complete C pipeline preserves the former renderer: ${filename}`, async () => {
    const source = await Bun.file(join(directory, filename)).text();
    const expected = await Bun.file(join(directory, filename.slice(0, -3) + ".html")).text();
    expect(await render(source, { runId: "a".repeat(32) })).toBe(expected);
  });
}

test("complete renderer enforces UTF-8/depth limits inside C and recovers after failures", async () => {
  const boundary = "01234567-89ab-cdef-0123-456789abcdef",
    runId = "a".repeat(32);
  expect(() => renderDocument("あ".repeat(333334), runId, boundary)).toThrow(RangeError);
  await expect(render("<div>".repeat(520))).rejects.toBeInstanceOf(RangeError);
  await expect(render("<template>".repeat(520) + "<input>")).rejects.toBeInstanceOf(RangeError);
  expect(await render("ok")).toBe("<p>ok</p>");
});

test("native identifier validation rejects truncation, embedded NULs and unsafe boundaries", () => {
  const valid = "01234567-89ab-cdef-0123-456789abcdef";
  for (const runId of [
    "A".repeat(32),
    "a".repeat(31),
    "a".repeat(32) + "\0",
    "a".repeat(31) + "\0",
  ])
    expect(() => renderDocument("x", runId, valid)).toThrow(TypeError);
  for (const boundary of ["a".repeat(36), valid.slice(0, -1) + ">", valid + "\0"])
    expect(() => renderDocument("x", "a".repeat(32), boundary)).toThrow(TypeError);
});
