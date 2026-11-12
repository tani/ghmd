import { renderDocument } from "./native-core";

export interface RenderOptions {
  /** Use a specific GitHub-style 32-character lowercase hexadecimal render ID. */
  runId?: string;
}

/** Render a GitHub-style HTML fragment in the calling Bun process. */
export async function render(source: string, options: RenderOptions = {}): Promise<string> {
  const runId =
    options.runId === undefined ? crypto.randomUUID().replaceAll("-", "") : options.runId;
  return renderDocument(source, runId, crypto.randomUUID());
}
export default render;
