import MarkdownIt from "markdown-it";
import { createRequire } from "node:module";

// The plugin publishes a CommonJS object with an explicit default export.
const math = createRequire(import.meta.url)("@vscode/markdown-it-katex").default;

export interface RenderOptions {
  /** VS Code's markdown.preview.linkify setting. Default: true. */
  linkify?: boolean;
  /** VS Code's markdown.preview.breaks setting. Default: false. */
  breaks?: boolean;
  /** VS Code's markdown.math.enabled setting. Default: true. */
  math?: boolean;
}

/** Build the same Markdown engine for rendering and source mapping. */
export function createMarkdownEngine(options: RenderOptions = {}) {
  const md = new MarkdownIt({
    html: true,
    linkify: options.linkify ?? true,
    breaks: options.breaks ?? false,
    typographer: false,
  });
  md.linkify.set({ fuzzyLink: false });
  if (options.math !== false)
    md.use(math, { enableFencedBlocks: true, globalGroup: true, macros: {} });
  return md;
}

/** Approximate VS Code's built-in Markdown preview. The output is not sanitized. */
export async function render(source: string, options: RenderOptions = {}): Promise<string> {
  if (typeof source !== "string") throw new TypeError("render expects a string");
  if (Buffer.byteLength(source) > 1_000_000) throw new RangeError("Markdown exceeds 1 MB");
  const md = createMarkdownEngine(options);
  const result = md.render(source);
  if (Buffer.byteLength(result) > 32 * 1024 * 1024) throw new RangeError("HTML exceeds 32 MB");
  return result;
}
export default render;
