import { existsSync, mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { createRequire } from "node:module";
import { tmpdir } from "node:os";
import { dirname, join, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { isMainThread } from "node:worker_threads";
import { nodeApiWrappers } from "./napi.ts";
import type { CCLibrary, CCOptions } from "./index.ts";

const require = createRequire(import.meta.url);
// Keep optional peers out of TypeScript's required type dependencies as well.
interface ForeignLibrary {
  func(name: string, result?: string, args?: string[]): (...args: any[]) => any;
  unload(): void;
}
interface Koffi {
  load(path: string): ForeignLibrary;
  proto(result: string, args: string[]): unknown;
  pointer(type: unknown): unknown;
  register(callback: (...args: any[]) => any, type: unknown): bigint;
  unregister(pointer: bigint): void;
}
const typeNames = [
  "char",
  "i8",
  "u8",
  "i16",
  "u16",
  "i32",
  "u32",
  "i64",
  "u64",
  "double",
  "float",
  "bool",
  "ptr",
  "void",
  "cstring",
  "i64_fast",
  "u64_fast",
  "function",
  "napi_env",
  "napi_value",
  "buffer",
];
const types: Record<string, string> = {
  char: "char",
  i8: "int8_t",
  int8_t: "int8_t",
  u8: "uint8_t",
  uint8_t: "uint8_t",
  i16: "int16_t",
  int16_t: "int16_t",
  u16: "uint16_t",
  uint16_t: "uint16_t",
  int: "int32_t",
  c_int: "int32_t",
  i32: "int32_t",
  int32_t: "int32_t",
  c_uint: "uint32_t",
  u32: "uint32_t",
  uint32_t: "uint32_t",
  i64: "int64_t",
  int64_t: "int64_t",
  isize: "intptr_t",
  i64_fast: "int64_t",
  u64: "uint64_t",
  uint64_t: "uint64_t",
  usize: "uintptr_t",
  u64_fast: "uint64_t",
  double: "double",
  f64: "double",
  float: "float",
  f32: "float",
  bool: "bool",
  void: "void",
  ptr: "void *",
  pointer: "void *",
  "void*": "void *",
  "char*": "void *",
  buffer: "void *",
  function: "void *",
  callback: "void *",
  fn: "void *",
  cstring: "str",
};
const list = (value: string | readonly string[] | undefined): readonly string[] =>
  value === undefined ? [] : typeof value === "string" ? [value] : value;
const typeName = (type: string | number = "void"): string => {
  const name = typeof type === "number" ? typeNames[type] : type;
  if (!name || (!Object.hasOwn(types, name) && name !== "napi_env" && name !== "napi_value"))
    throw new TypeError(`Unsupported cc() ABI type: ${String(type)}`);
  return name;
};
const bigInteger = (type: string) =>
  ["i64", "int64_t", "u64", "uint64_t", "isize", "usize"].includes(type);

let compiler: ReturnType<typeof loadCompiler> | undefined;
function loadCompiler() {
  if (!isMainThread)
    throw new Error("The node-tinycc cc() backend requires the Node.js main thread");
  let root: string, koffi: Koffi;
  try {
    root = dirname(require.resolve("node-tinycc"));
    koffi = require("koffi");
  } catch (cause) {
    throw new Error("Node.js cc() requires optional peers: npm install node-tinycc koffi", {
      cause,
    });
  }
  if (process.platform === "win32" && !["x64", "ia32"].includes(process.arch))
    throw new Error(`node-tinycc has no bundled Windows compiler for ${process.arch}`);
  const windows = process.platform === "win32";
  const base = join(root, windows ? (process.arch === "x64" ? "win64" : "win32") : "posix");
  const binary = windows ? join(base, "libtcc.dll") : join(root, "build/Release/tcc.node");
  if (!existsSync(binary))
    throw new Error(`node-tinycc is not built: ${binary}. Install/build the optional peer first.`);
  // A NAN addon must be registered by Node before dlopen; otherwise its module
  // constructor leaves Node's pending-addon registration state occupied.
  if (!windows) require(binary);
  const library = koffi.load(binary);
  const api: Record<string, (...args: any[]) => any> = {};
  for (const declaration of [
    "void *tcc_new(void)",
    "void tcc_delete(void *state)",
    "void tcc_set_lib_path(void *state, const char *path)",
    "void tcc_set_error_func(void *state, void *context, void *callback)",
    "void tcc_set_options(void *state, const char *options)",
    "void tcc_define_symbol(void *state, const char *name, const char *value)",
    "int tcc_add_include_path(void *state, const char *path)",
    "int tcc_add_library_path(void *state, const char *path)",
    "int tcc_add_library(void *state, const char *name)",
    "int tcc_set_output_type(void *state, int type)",
    "int tcc_compile_string(void *state, const char *source)",
    "int tcc_output_file(void *state, const char *path)",
  ])
    api[declaration.match(/\b(tcc_\w+)\(/)![1]] = library.func(declaration);
  const header = readFileSync(join(base, windows ? "libtcc/libtcc.h" : "include/libtcc.h"), "utf8");
  const output = (name: string) => {
    const match = header.match(new RegExp(`^#define\\s+${name}\\s+(\\d+)\\b`, "m"));
    if (!match) throw new Error(`node-tinycc header is missing ${name}`);
    return Number(match[1]);
  };
  return {
    koffi,
    api,
    library,
    base,
    windows,
    libPath: windows ? base : join(base, "lib/tcc"),
    include: windows
      ? [join(base, "include")]
      : [join(base, "lib/tcc/include"), join(base, "include")],
    dll: output("TCC_OUTPUT_DLL"),
  };
}

function nodeApiHeader(): string {
  const candidates = [
    resolve(dirname(process.execPath), "../include/node/node_api.h"),
    join(dirname(require.resolve("koffi")), "vendor/node-api-headers/include/node_api.h"),
  ];
  const header = candidates.find(existsSync);
  if (!header) throw new Error("Node-API headers are missing from Node.js and the koffi package");
  return header;
}

function includeFile(path: string): string {
  if (/["\r\n\0]/.test(path)) throw new TypeError("Unsupported character in C include path");
  return `#include "${path.replaceAll("\\", "/")}"\n`;
}

/** The cc() subset used here: source files, primitive/pointer ABIs, and
 * napi_value functions with a leading napi_env and only napi_value arguments. */
export function cc(options: CCOptions): CCLibrary {
  const bindings = Object.entries(options.symbols).map(([name, signature]) => {
    if (!/^[A-Za-z_]\w*$/.test(name)) throw new TypeError(`Invalid C symbol: ${name}`);
    for (const key of Object.keys(signature))
      if (key !== "args" && key !== "returns")
        throw new TypeError(`Unsupported cc() symbol option: ${key}`);
    const args = (signature.args ?? []).map(typeName);
    const returns = typeName(signature.returns);
    const napi = returns.startsWith("napi_") || args.some((type) => type.startsWith("napi_"));
    if (
      napi &&
      (returns !== "napi_value" ||
        args[0] !== "napi_env" ||
        args.slice(1).some((type) => type !== "napi_value"))
    )
      throw new TypeError(
        `Unsupported Node-API signature for ${name}: expected napi_env, napi_value... -> napi_value`,
      );
    if (args.includes("void") || returns === "buffer")
      throw new TypeError(`Invalid cc() signature for ${name}`);
    return { name, args, returns, napi };
  });
  const napi = bindings.filter((binding) => binding.napi);
  const { koffi, api, ...config } = (compiler ??= loadCompiler());
  let state: any = null;
  const errors: string[] = [];
  const errorCallback = koffi.register(
    (_context: unknown, message: string) => errors.push(message),
    koffi.pointer(koffi.proto("void", ["void *", "str"])),
  );
  let directory: string | undefined, addonPath: string | undefined;
  const check = (status: number, action: string) => {
    if (status < 0)
      throw new Error(`cc(): ${action}${errors.length ? `\n${errors.join("\n")}` : ""}`);
  };
  const cleanup = () => {
    if (!directory) return;
    try {
      rmSync(directory, { recursive: true, force: true });
      directory = undefined;
    } catch {
      /* Loaded DLLs can remain locked on Windows. */
    }
  };
  try {
    state = api.tcc_new();
    if (!state) throw new Error("TinyCC could not allocate a compiler state");
    api.tcc_set_error_func(state, null, errorCallback);
    api.tcc_set_lib_path(state, config.libPath);
    check(
      api.tcc_add_library_path(state, config.windows ? join(config.base, "lib") : config.libPath),
      "library path",
    );
    for (const path of [...list(options.include), ...config.include])
      check(api.tcc_add_include_path(state, resolve(path)), `include path: ${path}`);
    if (options.flags) api.tcc_set_options(state, list(options.flags).join(" "));
    for (const [name, value] of Object.entries(options.define ?? {}))
      api.tcc_define_symbol(state, name, value);
    check(api.tcc_set_output_type(state, config.dll), "output type");
    let source = includeFile(
      options.source instanceof URL ? fileURLToPath(options.source) : resolve(options.source),
    );
    directory = mkdtempSync(join(tmpdir(), "gfm-cc-"));
    if (napi.length) {
      mkdirSync(join(directory, "node"));
      writeFileSync(join(directory, "node/node_api.h"), includeFile(nodeApiHeader()));
      check(api.tcc_add_include_path(state, directory), "Node-API include path");
      source += nodeApiWrappers(napi);
    }
    check(api.tcc_compile_string(state, source), "compilation failed");
    const defaults = config.windows ? [] : ["m"];
    for (const name of new Set([...defaults, ...list(options.library)]))
      check(api.tcc_add_library(state, name), `library: ${name}`);
    addonPath = join(directory, "module.node");
    check(api.tcc_output_file(state, addonPath), "linking failed");
    api.tcc_delete(state);
    state = null;
    const native = napi.length ? require(addonPath) : undefined;
    const library = koffi.load(addonPath);
    let closed = false;
    const symbols: CCLibrary["symbols"] = Object.create(null);
    try {
      for (const binding of bindings) {
        let call: (...args: any[]) => any;
        if (binding.napi) call = native![binding.name];
        else
          call = library.func(
            binding.name,
            types[binding.returns],
            binding.args.map((type) => types[type]),
          );
        symbols[binding.name] = (...args: any[]) => {
          if (closed) throw new TypeError("Cannot call a symbol from a closed cc() library");
          const value = call(...args);
          if (!binding.napi && bigInteger(binding.returns)) return BigInt(value);
          if (!binding.napi && types[binding.returns] === "void *") {
            if (value === null) return null;
            const pointer = Number(value);
            if (!Number.isSafeInteger(pointer))
              throw new RangeError("Native pointer exceeds JavaScript's safe integer range");
            return pointer;
          }
          return value;
        };
      }
    } catch (error) {
      library.unload();
      throw error;
    }
    cleanup();
    if (directory) process.once("exit", cleanup);
    return {
      symbols,
      close() {
        if (closed) return;
        closed = true;
        library.unload();
        if (addonPath) delete require.cache[addonPath];
        cleanup();
      },
    };
  } finally {
    if (state) api.tcc_delete(state);
    koffi.unregister(errorCallback);
    cleanup();
  }
}
