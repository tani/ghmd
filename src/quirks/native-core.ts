import { cc } from "bun:ffi";
import { join } from "node:path";

const root = join(import.meta.dir, "../..");
const { symbols } = cc({
  source: join(root, "src/quirks/cmark/bun.c"),
  include: [
    join(root, "vendor/lexbor/compat"),
    ...["src", "extensions", "generated"].map((dir) => join(root, "vendor/cmark-gfm", dir)),
  ],
  symbols: {
    gfm_bun_render: {
      args: ["napi_env", "napi_value", "napi_value", "napi_value"],
      returns: "napi_value",
    },
  },
});

// Bun 1.4.4 keeps an argument slot for the automatically supplied napi_env.
export const renderDocument = (source: string, runId: string, boundary: string): string =>
  symbols.gfm_bun_render(undefined, source, runId, boundary) as string;
