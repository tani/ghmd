# Runtime C compiler adapter

`cc()` selects Bun's built-in `bun:ffi` implementation on Bun and a Node.js
adapter using the optional `node-tinycc` and `koffi` peers on Node. Both native
entry points in this project use this module, so the renderer and HTML parser
share the same C sources across runtimes.

`package.json` selects the backend through its conditional `imports` entry for
`#ffi`. The `bun` condition appears before `node`, with a Node fallback under
`default`. `index.ts` imports this alias statically; no runtime version check or
dynamic import is needed, and the unselected backend is not loaded.

## Use

Node.js 24 or later can run the repository's TypeScript directly from a checkout.
Install and build the optional peers before using its native renderer:

```sh
npm install node-tinycc koffi
node src/quirks/cli.ts input.md
npm run test:node
```

Both packages appear in `peerDependencies`, with `optional: true` in
`peerDependenciesMeta`. Bun uses its bundled compiler and does not load either
peer. `node-tinycc` requires native build tools at installation time; its current
POSIX installer also clones TinyCC over SSH. Follow its
[upstream installation source](https://github.com/jerch/node-tinycc/blob/master/install_tcc.sh)
when configuring those build prerequisites. Compilation by this adapter runs
in-process and does not invoke a compiler executable or download anything.

```ts
import { cc } from "gfm-bun-render/ffi";

const library = cc({
  source: new URL("./add.c", import.meta.url),
  symbols: { add: { args: ["int", "int"], returns: "int" } },
});
try {
  console.log(library.symbols.add(2, 3));
} finally {
  library.close();
}
```

The Node adapter supports a source-file path or file URL, `include`, `library`,
`define`, `flags`, and named `symbols`. It handles primitive numeric types,
booleans, buffers, raw pointers, C strings, and Bun's numeric type identifiers
for those types. Ordinary `i64`/`u64` results become BigInts; the `*_fast` types
follow Koffi's number-or-BigInt conversion. Pointer results are safe JavaScript
integers or `null`; C strings become JavaScript strings.

Node-API signatures must return `napi_value` and take a leading `napi_env`
followed only by `napi_value` arguments. Calls retain Bun 1.4.4's ignored first
argument slot: `symbol(undefined, ...values)`. The adapter generates real
Node-API callbacks, preserving JavaScript values, string ownership, and native
exceptions. Unsupported ABI types and mixed Node-API/primitive signatures are
rejected. This is a `cc()` subset, not an implementation of all of `bun:ffi`;
it does not expose `CString`, `JSCallback`, BunFile input, or automatic buffer
length arguments.

## Implementation and lifetime

| File | Responsibility |
| --- | --- |
| [index.ts](index.ts) | Shared option/result types and the public facade over `#ffi`. |
| [bun.ts](bun.ts) | Re-exports Bun's built-in `cc()` for the `bun` condition. |
| [node.ts](node.ts) | Optional-peer loading, TinyCC compilation, Koffi bindings, diagnostics, and cleanup. |
| [napi.ts](napi.ts) | Generates Node-API callbacks and module registration for declared env/value signatures. |

The Node adapter loads the compiler distributed by `node-tinycc`: its POSIX
native addon or its bundled Windows DLL. Koffi calls the underlying libtcc API
directly, allowing compiler states to be explicitly freed without using
node-tinycc's older FFI wrappers. The shared-library output constant is read from
the matching libtcc header because compiler versions assign different values.

TinyCC writes a library in a private temporary directory, then the compiler
state is freed. Koffi binds ordinary functions; Node loads generated Node-API
exports when needed. Node-API headers come from the Node installation or Koffi's
bundled headers. Temporary files are removed after loading where the OS permits,
with another cleanup attempt at close or process exit when files remain locked.

`close()` is idempotent, releases the Koffi handle, and rejects subsequent calls
even through retained symbol functions. Node retains loaded Node-API addons for
the environment's lifetime; closing the facade does not unload those addons.
Compilation and calls are synchronous. The Node backend requires the main
thread because node-tinycc's POSIX addon uses legacy module registration;
worker compilation is rejected before that addon is loaded.

The Node path has been verified on Node 24.16.0 / Linux ARM64 with
node-tinycc 0.1.0 and Koffi 3.3.7. Other platforms remain unverified. This private
repository exposes TypeScript source; distributing it under `node_modules` for
Node requires a separate transpilation step.

## Checks

`npm run test:node` uses Node's test runner for ABI conversion, error recovery,
ownership boundaries, optional-peer errors, HTML provenance, comparison, the
standalone C API, and the same saved GitHub API fixtures used by Bun. All
reference comparisons are offline. `bun run test` exercises the original Bun
compiler path without loading the Node-only peers.
