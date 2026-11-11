import { parseFragment } from "parse5";

/** Resolve image URLs through GitHub's Markdown API. Each unique canonical URL
 * is sent to GitHub; callers must opt in and should consider API rate limits. */
export function createGitHubImageProxy(
  options: { fetch?: typeof fetch; token?: string; timeoutMs?: number } = {},
) {
  const request = options.fetch ?? fetch;
  const cache = new Map<string, Promise<string>>();
  return (canonical: string): Promise<string> => {
    const existing = cache.get(canonical);
    if (existing) return existing;
    const pending = (async () => {
      const parsed = new URL(canonical);
      if (!["http:", "https:"].includes(parsed.protocol))
        throw new TypeError("Image URL must use HTTP(S)");
      const escaped = canonical
        .replaceAll("&", "&amp;")
        .replaceAll('"', "&quot;")
        .replaceAll("<", "&lt;")
        .replaceAll(">", "&gt;");
      const headers: Record<string, string> = {
        Accept: "application/vnd.github+json",
        "Content-Type": "application/json",
        "X-GitHub-Api-Version": "2022-11-28",
      };
      if (options.token) headers.Authorization = `Bearer ${options.token}`;
      const response = await request("https://api.github.com/markdown", {
        method: "POST",
        headers,
        body: JSON.stringify({ text: `<img src="${escaped}">`, mode: "gfm" }),
        signal: AbortSignal.timeout(options.timeoutMs ?? 15000),
      });
      if (!response.ok)
        throw new Error(`GitHub image proxy lookup failed: HTTP ${response.status}`);
      const html = await response.text();
      if (html.length > 1_000_000) throw new Error("GitHub image proxy response is too large");
      const tree: any = parseFragment(html);
      const find = (node: any): string | undefined => {
        if (node.tagName === "img") {
          const attrs = new Map<string, string>(
            (node.attrs ?? []).map((a: any) => [a.name, a.value]),
          );
          if (attrs.get("data-canonical-src") === canonical) {
            const src = attrs.get("src");
            if (src) {
              const url = new URL(src);
              if (url.protocol === "https:" && url.hostname === "camo.githubusercontent.com")
                return src;
            }
          }
        }
        for (const child of node.childNodes ?? []) {
          const result = find(child);
          if (result) return result;
        }
      };
      const url = find(tree);
      if (!url) throw new Error("GitHub did not return a signed image URL");
      return url;
    })();
    cache.set(canonical, pending);
    pending.catch(() => {
      if (cache.get(canonical) === pending) cache.delete(canonical);
    });
    return pending;
  };
}
