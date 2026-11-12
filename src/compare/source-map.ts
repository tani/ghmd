import { createRequire } from "node:module";
import { parseFragment, serialize } from "../html/index.ts";
import { createMarkdownEngine } from "../vscode/index.ts";

const sourceMap = createRequire(import.meta.url)("markdown-it-source-map").default;

export interface MappedBlock {
  map: [number, number];
  /** HTML inside the source-mapped top-level element. */
  html: string;
}

/** Parse once and correlate trusted plugin-generated HTML markers with tokens. */
export function parseWithSourceMap(source: string): { tokens: any[]; blocks: MappedBlock[] } {
  const engine = createMarkdownEngine();
  engine.use(sourceMap);
  const marker = `data-gfm-map-${crypto.randomUUID().replaceAll("-", "")}`;
  const renderToken = engine.renderer.renderToken.bind(engine.renderer);
  engine.renderer.renderToken = (tokens, index, options) => {
    const token = tokens[index];
    if (token.level === 0 && token.map && token.type.endsWith("_open"))
      token.attrSet(marker, String(index));
    return renderToken(tokens, index, options);
  };
  const environment = {};
  const tokens = engine.parse(source, environment);
  const html = engine.renderer.render(tokens, engine.options, environment);
  const tree: any = parseFragment(html);
  const blocks: MappedBlock[] = [];
  const visited = new Set<number>();
  const visit = (node: any, parentHidden = false) => {
    const attrs = new Map<string, string>(
      (node.attrs ?? []).map((a: any) => [a.name, String(a.value)]),
    );
    const hidden =
      parentHidden ||
      node.tagName === "svg" ||
      attrs.get("aria-hidden") === "true" ||
      /(?:^|\s)sr-only(?:\s|$)/.test(attrs.get("class") ?? "");
    const value = attrs.get(marker),
      line = attrs.get("data-source-line");
    if (value !== undefined && /^\d+$/.test(value)) {
      const index = Number(value),
        token = tokens[index];
      if (
        token?.map &&
        token.level === 0 &&
        token.type.endsWith("_open") &&
        token.tag === node.tagName &&
        String(token.map[0] + 1) === line &&
        !visited.has(index)
      ) {
        visited.add(index);
        blocks.push({ map: [token.map[0], token.map[1]], html: hidden ? "" : serialize(node) });
      }
    }
    for (const child of node.childNodes ?? []) visit(child, hidden);
    if (node.content) visit(node.content, hidden);
  };
  visit(tree);
  blocks.sort((a, b) => a.map[0] - b.map[0] || a.map[1] - b.map[1]);
  return { tokens, blocks };
}
