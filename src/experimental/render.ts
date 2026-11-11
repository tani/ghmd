/** Bun-only HTML experiment. Not a security-audited replacement for the default renderer. */
import {
  finalizeHtml,
  flattenTemplates,
  renderMarkdown,
  prepareDetails,
  prepareMathTables,
  processLinks,
  removeBoundaries,
  renderMathHtml,
} from "../quirks/native-core";
import { removeRawInputs, sanitize } from "./sanitize";
import { annotateImplicitClosures } from "./scopes";
import type { RenderOptions } from "../quirks/index";

export interface ExperimentalOptions extends RenderOptions {
  strategy?: "stream" | "closure-events" | "closure-events-aggressive";
}
export async function render(source: string, options: ExperimentalOptions = {}): Promise<string> {
  if (typeof source !== "string") throw new TypeError("render expects a string");
  if (options.runId !== undefined && !/^[a-f0-9]{32}$/.test(options.runId))
    throw new TypeError("runId must be 32 lowercase hexadecimal characters");
  if (Buffer.byteLength(source) > 1_000_000) throw new RangeError("Markdown exceeds 1 MB");
  const mark = crypto.randomUUID();
  const token = options.runId ?? crypto.randomUUID().replaceAll("-", "");
  const tableMarker = `data-gfm-table-${mark.replaceAll("-", "")}`;
  const html = renderMarkdown(source, removeRawInputs);
  const screened = sanitize(html, mark);
  const scoped =
    options.strategy === "stream"
      ? { html: screened, restore: (value: string) => value }
      : annotateImplicitClosures(
          screened,
          mark.replaceAll("-", ""),
          options.strategy === "closure-events-aggressive",
        );
  const flattened = flattenTemplates(scoped.html, mark);
  const prepared = prepareDetails(flattened);
  const marked = prepareMathTables(prepared, token, mark, tableMarker);
  const math = renderMathHtml(marked, token, mark);
  const linked = processLinks(removeBoundaries(math, mark), mark);
  const output = scoped.restore(finalizeHtml(linked, token, tableMarker)).replace(/\n$/, "");
  if (Buffer.byteLength(output) > 32 * 1024 * 1024)
    throw new RangeError("Rendered HTML exceeds 32 MB");
  return output;
}
