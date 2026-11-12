import { expect, test } from "bun:test";
import { renderDocument } from "../src/quirks/native-core";

const token = "a".repeat(32),
  boundary = "01234567-89ab-cdef-0123-456789abcdef";

test("native render validates required strings and preserves Unicode and NUL handling", () => {
  expect(renderDocument("", token, boundary)).toBe("");
  expect(renderDocument("日本語\0😀", token, boundary)).toBe("<p>日本語�😀</p>");
  for (const value of [null, undefined, 1, {}, Buffer.from("x")]) {
    expect(() => renderDocument(value as any, token, boundary)).toThrow(TypeError);
    expect(() => renderDocument("x", value as any, boundary)).toThrow(TypeError);
    expect(() => renderDocument("x", token, value as any)).toThrow(TypeError);
  }
});

test("complete render validates identifiers inside C", () => {
  expect(() => renderDocument("x", "bad", boundary)).toThrow("runId must be 32");
  expect(() => renderDocument("x", token, "bad")).toThrow("Invalid render boundary");
});

test("production bindings expose only rendering and the HTML parser operations in use", async () => {
  expect(Object.keys(await import("../src/quirks/native-core"))).toEqual(["renderDocument"]);
  expect(Object.keys(await import("../src/html")).sort()).toEqual(["parseFragment", "serialize"]);
});
