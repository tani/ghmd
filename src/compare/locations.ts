import { parseFragment } from "parse5";
import {
  missingTextFragments,
  summarizeHtml,
  type SemanticKind,
  type SemanticNode,
  type SemanticSummary,
} from "./index";
import { parseWithSourceMap } from "./source-map";

type Point = { line: number; column: number; offset: number };
export type SourceRange = { start: Point; end: Point; precision: "exact" | "block" };
export type LocatedFinding = {
  category: SemanticKind;
  evidence?: unknown;
  locations: SourceRange[];
  status: "exact" | "block" | "ambiguous" | "unavailable";
  note?: string;
};
type Analysis = { categories: SemanticKind[]; github: SemanticSummary; vscode: SemanticSummary };
type Token = {
  type: string;
  tag: string;
  nesting: number;
  map: [number, number] | null;
  content: string;
  info: string;
  children: Token[] | null;
  attrGet(name: string): string | null;
};

/** Attach verified enclosing Markdown blocks to VS Code's semantic nodes. */
export function attachNodeLocations(
  source: string,
  nodes: SemanticNode[],
  parsed: ReturnType<typeof parseWithSourceMap>,
): SemanticNode[] {
  const starts = [0];
  for (let i = 0; i < source.length; i++) if (source[i] === "\n") starts.push(i + 1);
  const point = (offset: number): Point => {
    let line = 0;
    while (line + 1 < starts.length && starts[line + 1] <= offset) line++;
    return { line: line + 1, column: [...source.slice(starts[line], offset)].length + 1, offset };
  };
  const candidates = new Map<string, { ranges: SourceRange[]; count: number }>();
  const add = (node: SemanticNode, map: [number, number]) => {
    const begin = starts[map[0]];
    if (begin === undefined) return;
    const end = map[1] < starts.length ? starts[map[1]] : source.length;
    const key = JSON.stringify(node);
    const entry = candidates.get(key) ?? { ranges: [], count: 0 };
    entry.count++;
    if (!entry.ranges.some((r) => r.start.offset === begin && r.end.offset === end))
      entry.ranges.push({ start: point(begin), end: point(end), precision: "block" });
    candidates.set(key, entry);
  };
  for (const block of parsed.blocks)
    for (const node of summarizeHtml(block.html).nodes) add(node, block.map);
  // Fences are rendered by custom rules and do not carry source-map markers.
  for (const token of parsed.tokens as Token[]) {
    if (!token.map || !["fence", "code_block"].includes(token.type)) continue;
    const math = token.type === "fence" && token.info.trim().split(/\s+/)[0] === "math";
    const node: SemanticNode = math
      ? { kind: "math", display: true, expression: token.content.replace(/\n$/, "") }
      : { kind: "code", block: true, text: token.content.replace(/\n$/, "") };
    add(node, token.map);
  }
  return nodes.map((node) => {
    const entry = candidates.get(JSON.stringify(node));
    return {
      ...node,
      source: {
        status: !entry ? "unavailable" : entry.count > 1 ? "ambiguous" : "block",
        locations: entry?.ranges ?? [],
      },
    };
  });
}

/** Map output differences to source; ambiguous or untraceable matches are never guessed. */
export function locateDifferences(
  source: string,
  analysis: Analysis,
  sourceMap?: ReturnType<typeof parseWithSourceMap>,
): LocatedFinding[] {
  if (!analysis.categories.length) return [];
  const parsed = sourceMap ?? parseWithSourceMap(source);
  const tokens = parsed.tokens as Token[];
  const mappedBlocks = parsed.blocks.map((item) => ({
    ...item,
    summary: summarizeHtml(item.html),
  }));
  // The plugin maps only top-level rendered elements. Restrict matching inside
  // those verified scopes; retain token fallback for unmapped constructs.
  const sourceMapAllows = (map: Token["map"], category: SemanticKind, value: unknown) => {
    if (!map) return true;
    const enclosing = mappedBlocks.filter((item) => item.map[0] <= map[0] && item.map[1] >= map[1]);
    return (
      !enclosing.length ||
      enclosing.some((item) =>
        item.summary[category].some(
          (candidate) => JSON.stringify(candidate) === JSON.stringify(value),
        ),
      )
    );
  };
  const starts = [0];
  for (let i = 0; i < source.length; i++) if (source[i] === "\n") starts.push(i + 1);
  const point = (offset: number): Point => {
    let lo = 0,
      hi = starts.length;
    while (lo + 1 < hi) {
      const mid = (lo + hi) >> 1;
      if (starts[mid] <= offset) lo = mid;
      else hi = mid;
    }
    return { line: lo + 1, column: [...source.slice(starts[lo], offset)].length + 1, offset };
  };
  const range = (start: number, end: number, precision: SourceRange["precision"]): SourceRange => ({
    start: point(start),
    end: point(end),
    precision,
  });
  const block = (map: Token["map"]): SourceRange[] => {
    if (!map) return [];
    const begin = starts[map[0]],
      end = map[1] < starts.length ? starts[map[1]] : source.length;
    return begin === undefined ? [] : [range(begin, end, "block")];
  };
  const occurrenceCache = new Map<string, SourceRange[]>();
  const occurrences = (needle: string, map: Token["map"]): SourceRange[] => {
    if (!needle || !map) return [];
    const begin = starts[map[0]],
      end = map[1] < starts.length ? starts[map[1]] : source.length;
    if (begin === undefined) return [];
    const cacheKey = `${begin}:${end}:${needle}`;
    const cached = occurrenceCache.get(cacheKey);
    if (cached) return cached;
    const found: SourceRange[] = [];
    for (
      let pos = source.indexOf(needle, begin);
      pos >= 0 && pos + needle.length <= end && found.length < 64;
      pos = source.indexOf(needle, pos + 1)
    )
      found.push(range(pos, pos + needle.length, "exact"));
    occurrenceCache.set(cacheKey, found);
    return found;
  };
  const unescapedOccurrences = (needle: string, map: Token["map"]): SourceRange[] => {
    if (!needle || !map) return [];
    const begin = starts[map[0]],
      end = map[1] < starts.length ? starts[map[1]] : source.length;
    if (begin === undefined) return [];
    let normalized = "";
    const offsets: number[] = [],
      ends: number[] = [];
    for (let i = begin; i < end; i++) {
      const start = i;
      if (
        source[i] === "\\" &&
        i + 1 < end &&
        /[!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~]/.test(source[i + 1])
      )
        i++;
      normalized += source[i];
      offsets.push(start);
      ends.push(i + 1);
    }
    const found: SourceRange[] = [];
    for (
      let i = normalized.indexOf(needle);
      i >= 0 && found.length < 64;
      i = normalized.indexOf(needle, i + 1)
    ) {
      found.push(range(offsets[i], ends[i + needle.length - 1], "exact"));
    }
    return found;
  };
  const missing = (expected: unknown[], actual: unknown[]): unknown[] => {
    const counts = new Map<string, number>();
    for (const item of actual) {
      const key = JSON.stringify(item);
      counts.set(key, (counts.get(key) ?? 0) + 1);
    }
    return expected.filter((item) => {
      const key = JSON.stringify(item),
        n = counts.get(key) ?? 0;
      if (!n) return true;
      counts.set(key, n - 1);
      return false;
    });
  };
  const finalize = (
    category: SemanticKind,
    evidence: unknown,
    matches: SourceRange[],
    note?: string,
  ): LocatedFinding => {
    const all = [
      ...new Map(matches.map((r) => [`${r.start.offset}:${r.end.offset}`, r])).values(),
    ].sort((a, b) => a.start.offset - b.start.offset || a.end.offset - b.end.offset);
    const unique = all.slice(0, 128);
    const status = !unique.length
      ? "unavailable"
      : unique.length > 1
        ? "ambiguous"
        : unique[0].precision;
    return {
      category,
      ...(evidence === undefined ? {} : { evidence }),
      locations: unique,
      status,
      ...(note || all.length > 128 || unique.length >= 64
        ? {
            note: [
              note,
              unique.length >= 64
                ? "Candidate lists are capped; further matches may be omitted."
                : "",
            ]
              .filter(Boolean)
              .join(" "),
          }
        : {}),
    };
  };
  const inline = tokens.filter((t) => t.type === "inline" && t.map);
  const findings: LocatedFinding[] = [];
  const uniqueMissing = (expected: unknown[], actual: unknown[]) => [
    ...new Map(missing(expected, actual).map((value) => [JSON.stringify(value), value])).values(),
  ];
  for (const category of analysis.categories) {
    if (category === "math") {
      for (const value of uniqueMissing(analysis.vscode.math, analysis.github.math) as {
        display: boolean;
        expression: string;
      }[]) {
        const candidates: SourceRange[] = [];
        const seenScopes = new Set<string>();
        for (const token of tokens) {
          if (
            (token.type.startsWith("math_") ||
              (token.type === "fence" && token.info.trim().split(/\s+/)[0] === "math")) &&
            token.map
          ) {
            if (
              value.display &&
              token.content.replace(/\n$/, "") === value.expression &&
              sourceMapAllows(token.map, category, value)
            )
              candidates.push(...block(token.map));
          }
          if (token.type !== "inline" || !token.map) continue;
          if (!sourceMapAllows(token.map, category, value)) continue;
          for (const child of token.children ?? [])
            if (child.type === "math_inline") {
              const inner = child.content,
                normalized =
                  inner.length > 2 && inner.startsWith("`") && inner.endsWith("`")
                    ? inner.slice(1, -1)
                    : inner;
              if (!value.display && normalized === value.expression) {
                const scope = `${token.map[0]}:${token.map[1]}:${inner}`;
                if (seenScopes.has(scope)) continue;
                seenScopes.add(scope);
                const exact = occurrences(`$${inner}$`, token.map);
                candidates.push(...(exact.length ? exact : block(token.map)));
              }
            }
        }
        findings.push(
          finalize(
            category,
            value,
            candidates,
            "Ranges identify matching Markdown expressions; repeated expressions may be ambiguous.",
          ),
        );
      }
      continue;
    }
    if (category === "links" || category === "images") {
      for (const value of uniqueMissing(analysis.vscode[category], analysis.github[category])) {
        const target =
          category === "links" ? String(value) : String((value as { src: string }).src);
        const candidates: SourceRange[] = [];
        const inspectHtml = (token: Token, map: Token["map"]) => {
          if (!map || !token.content) return;
          const tree: any = parseFragment(token.content, { sourceCodeLocationInfo: true });
          const visit = (node: any) => {
            const name = category === "links" ? "href" : "src";
            const tag = category === "links" ? "a" : "img";
            const attr = node.attrs?.find((item: any) => item.name === name);
            const location = node.sourceCodeLocation?.attrs?.[name];
            if (node.tagName === tag && attr?.value === target && location) {
              const raw = token.content.slice(location.startOffset, location.endOffset);
              const exact = occurrences(raw, map);
              candidates.push(...(exact.length ? exact : block(map)));
            }
            for (const child of node.childNodes ?? []) visit(child);
            if (node.content) visit(node.content);
          };
          visit(tree);
        };
        for (const token of inline)
          for (const child of token.children ?? []) {
            if (!sourceMapAllows(token.map, category, value)) continue;
            if (child.type === "html_inline") inspectHtml(child, token.map);
            const childTarget =
              category === "links" && child.type === "link_open"
                ? child.attrGet("href")
                : category === "images" && child.type === "image"
                  ? child.attrGet("src")
                  : null;
            if (childTarget !== target) continue;
            const exact = occurrences(target, token.map);
            candidates.push(...(exact.length ? exact : block(token.map)));
          }
        for (const token of tokens)
          if (token.type === "html_block" && sourceMapAllows(token.map, category, value))
            inspectHtml(token, token.map);
        findings.push(
          finalize(
            category,
            value,
            candidates,
            "Exact ranges identify the destination or HTML attribute, not necessarily the whole link or image.",
          ),
        );
      }
      continue;
    }
    if (category === "code") {
      for (const value of uniqueMissing(analysis.vscode.code, analysis.github.code) as {
        block: boolean;
        text: string;
      }[]) {
        const candidates: SourceRange[] = [];
        if (value.block)
          for (const token of tokens) {
            if (
              ["fence", "code_block"].includes(token.type) &&
              token.content.replace(/\n$/, "") === value.text &&
              sourceMapAllows(token.map, category, value)
            )
              candidates.push(...block(token.map));
          }
        else
          for (const token of inline)
            for (const child of token.children ?? [])
              if (
                child.type === "code_inline" &&
                child.content === value.text &&
                sourceMapAllows(token.map, category, value)
              ) {
                const exact = occurrences(`\`${child.content}\``, token.map);
                candidates.push(...(exact.length ? exact : block(token.map)));
              }
        findings.push(finalize(category, value, candidates));
      }
      continue;
    }
    if (category === "text") {
      const words = (value: unknown) => String(value ?? "").match(/[\p{L}\p{N}_]+/gu) ?? [];
      const fragments = missingTextFragments(analysis.vscode.text, analysis.github.text);
      const different = [
        ...new Set(
          missing(fragments.flatMap(words), analysis.github.text.flatMap(words)) as string[],
        ),
      ];
      const candidates: SourceRange[] = [];
      for (const word of different.slice(0, 32)) {
        for (const token of inline)
          for (const child of token.children ?? []) {
            if (child.type !== "text" || !words(child.content).includes(word)) continue;
            for (const match of occurrences(word, token.map)) {
              const before = source[match.start.offset - 1] ?? "",
                after = source[match.end.offset] ?? "";
              if (!/[\p{L}\p{N}_]/u.test(before) && !/[\p{L}\p{N}_]/u.test(after))
                candidates.push(match);
            }
          }
      }
      if (!candidates.length) {
        const right = fragments[0] ?? "";
        const score = (candidate: string) => {
          let prefix = 0,
            suffix = 0;
          while (
            prefix < candidate.length &&
            prefix < right.length &&
            candidate[prefix] === right[prefix]
          )
            prefix++;
          while (
            suffix < candidate.length - prefix &&
            suffix < right.length - prefix &&
            candidate[candidate.length - 1 - suffix] === right[right.length - 1 - suffix]
          )
            suffix++;
          return prefix + suffix;
        };
        // A single VS Code text node may span several GitHub blocks. Compare
        // nearby chunks only for this unmatched node, not a flattened document.
        const chunks = analysis.github.text.map(String);
        const windows = chunks.flatMap((_, i) =>
          [1, 2, 3]
            .filter((n) => i + n <= chunks.length)
            .map((n) => chunks.slice(i, i + n).join(" ")),
        );
        const left = windows.sort((a, b) => score(b) - score(a))[0] ?? "";
        let prefix = 0;
        while (prefix < left.length && prefix < right.length && left[prefix] === right[prefix])
          prefix++;
        let suffix = 0;
        while (
          suffix < left.length - prefix &&
          suffix < right.length - prefix &&
          left[left.length - 1 - suffix] === right[right.length - 1 - suffix]
        )
          suffix++;
        const deltas = [
          right.slice(prefix, right.length - suffix).trim(),
          left.slice(prefix, left.length - suffix).trim(),
        ];
        const scopes = [
          ...inline.map((token) => token.map),
          ...tokens.filter((token) => token.type === "html_block").map((token) => token.map),
        ];
        for (const delta of deltas) {
          if (!delta) continue;
          for (const map of scopes) candidates.push(...occurrences(delta, map));
          if (!candidates.length)
            for (const map of scopes) candidates.push(...unescapedOccurrences(delta, map));
          if (candidates.length) break;
        }
      }
      findings.push(
        finalize(
          category,
          different.slice(0, 32),
          candidates,
          `Ranges identify matching text fragments or words, including decoded Markdown escapes; they are candidates, not proven causes.${different.length > 32 ? " Only the first 32 distinct words were examined." : ""}${!candidates.length && analysis.categories.includes("math") ? " The same document also has math findings; their locations may help investigation, but the relationship is unconfirmed." : ""}`,
        ),
      );
      continue;
    }
    if (category === "tasks") {
      const missingStates = missing(analysis.vscode.tasks, analysis.github.tasks) as boolean[];
      const wanted = new Set(missingStates.length ? missingStates : analysis.vscode.tasks);
      const candidates: SourceRange[] = [];
      const inspectHtml = (token: Token, map: Token["map"]) => {
        if (!map || !token.content) return;
        const tree: any = parseFragment(token.content, { sourceCodeLocationInfo: true });
        const visit = (node: any) => {
          if (
            node.tagName === "input" &&
            node.attrs?.some((a: any) => a.name === "type" && a.value.toLowerCase() === "checkbox")
          ) {
            const checked = node.attrs.some((a: any) => a.name === "checked");
            const location = node.sourceCodeLocation?.startTag;
            if (wanted.has(checked) && location) {
              const markup = token.content.slice(location.startOffset, location.endOffset);
              const exact = occurrences(markup, map);
              candidates.push(...(exact.length ? exact : block(map)));
            }
          }
          for (const child of node.childNodes ?? []) visit(child);
          if (node.content) visit(node.content);
        };
        visit(tree);
      };
      for (const token of tokens) {
        if (token.type === "html_block") inspectHtml(token, token.map);
        if (token.type === "inline")
          for (const child of token.children ?? []) {
            if (child.type === "html_inline") inspectHtml(child, token.map);
          }
      }
      findings.push(
        finalize(
          category,
          missingStates,
          candidates,
          "Ranges identify matching source-authored checkbox elements; other task differences may be untraceable.",
        ),
      );
      continue;
    }
    if (category === "structure" || category === "formatting") {
      const names =
        category === "structure"
          ? new Set(["ul", "ol", "li", "table", "thead", "tbody", "tr", "th", "td", "br"])
          : new Set(["strong", "em", "del", "h1", "h2", "h3", "h4", "h5", "h6", "blockquote"]);
      const events: { name: string; map: Token["map"] }[] = [];
      const stack: { tag: string; map: Token["map"] }[] = [];
      const add = (token: Token, inherited: Token["map"]) => {
        if (token.type === "softbreak") return;
        const tag =
          token.tag === "s"
            ? "del"
            : token.tag === "b"
              ? "strong"
              : token.tag === "i"
                ? "em"
                : token.tag;
        if (!names.has(tag)) return;
        if (token.nesting === 1) {
          stack.push({ tag, map: token.map ?? inherited });
          events.push({ name: `<${tag}>`, map: token.map ?? inherited });
        } else if (token.nesting === -1) {
          const previous = stack.pop();
          events.push({ name: `</${tag}>`, map: previous?.tag === tag ? previous.map : inherited });
        } else if (tag === "br") events.push({ name: "<br>", map: token.map ?? inherited });
      };
      for (const token of tokens) {
        if (token.type === "inline")
          for (const child of token.children ?? []) add(child, token.map);
        else add(token, token.map);
      }
      const namesFromOutput = analysis.vscode[category] as string[];
      if (JSON.stringify(events.map((e) => e.name)) !== JSON.stringify(namesFromOutput)) {
        findings.push(
          finalize(
            category,
            undefined,
            [],
            "The parser tokens could not be aligned reliably with the rendered structure.",
          ),
        );
        continue;
      }
      let index = 0;
      for (const name of analysis.github[category] as string[])
        if (name === namesFromOutput[index]) index++;
      const event = events[index];
      findings.push(
        finalize(
          category,
          event?.name,
          event ? block(event.map) : [],
          "The range is the enclosing Markdown block, not an exact cause.",
        ),
      );
      continue;
    }
    findings.push(
      finalize(
        category,
        undefined,
        [],
        "This document-wide difference cannot be reliably mapped to a source range.",
      ),
    );
  }
  return findings;
}
