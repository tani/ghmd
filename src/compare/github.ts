/** Request one GitHub Markdown API reference. Call only when explicitly requested. */
export async function fetchGitHubHtml(
  source: string,
  options: { token?: string; fetch?: typeof fetch } = {},
): Promise<string> {
  const headers: Record<string, string> = {
    Accept: "application/vnd.github+json",
    "Content-Type": "application/json",
    "X-GitHub-Api-Version": "2022-11-28",
  };
  if (options.token) headers.Authorization = `Bearer ${options.token}`;
  const response = await (options.fetch ?? fetch)("https://api.github.com/markdown", {
    method: "POST",
    headers,
    body: JSON.stringify({ text: source, mode: "gfm" }),
    signal: AbortSignal.timeout(30_000),
  });
  if (!response.ok) throw new Error(`GitHub Markdown API returned HTTP ${response.status}`);
  return response.text();
}
