import { describe, expect, test } from "bun:test";
import { readdir, readFile } from "node:fs/promises";
import { join, resolve } from "node:path";

const root = resolve(import.meta.dir, "..");
describe("in-process C compilation", () => {
  test("preprocessor entry point covers every upstream C translation unit", async () => {
    const text = await readFile(join(root, "src/cmark-unity.c"), "utf8");
    const found = [
      ...text.matchAll(/^#include "\.\.\/vendor\/cmark-gfm\/(src|extensions)\/([^"/]+\.c)"$/gm),
    ].map((m) => `${m[1]}/${m[2]}`);
    const expected: string[] = [];
    for (const dir of ["src", "extensions"])
      for (const name of await readdir(join(root, "vendor/cmark-gfm", dir)))
        if (name.endsWith(".c") && name !== "main.c") expected.push(`${dir}/${name}`);
    expect(found.sort()).toEqual(expected.sort());
  });
  test("fresh import works without a compiler on PATH or subprocess APIs", () => {
    const script = `
   Bun.spawn=()=>{throw Error('unexpected subprocess')};
   Bun.spawnSync=()=>{throw Error('unexpected subprocess')};
   const {render}=await import('./src/index.ts');
   const values=await Promise.all([render(''),render('日本語'),render('a\\0b')]);
   if(JSON.stringify(values)!==JSON.stringify(['','<p>日本語</p>','<p>a�b</p>']))throw Error('unexpected output');
  `;
    const result = Bun.spawnSync([process.execPath, "-e", script], {
      cwd: root,
      env: { ...process.env, PATH: "/nonexistent", CC: "/nonexistent/cc" },
      stdout: "pipe",
      stderr: "pipe",
    });
    expect(result.exitCode).toBe(0);
    expect(result.stderr.toString()).toBe("");
  });
});
