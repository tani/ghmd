import { expect, test } from "bun:test";
import { fetchGitHubHtml } from "../src/compare/github";

test("live reference sends one GFM request only when invoked", async () => {
  const requests: { url: string; init: RequestInit }[] = [];
  const fetcher = async (url: RequestInfo | URL, init?: RequestInit) => {
    requests.push({ url: String(url), init: init! });
    return new Response("<p>ok</p>");
  };
  expect(
    await fetchGitHubHtml("hello", { token: "example-token", fetch: fetcher as typeof fetch }),
  ).toBe("<p>ok</p>");
  expect(requests).toHaveLength(1);
  expect(requests[0].url).toBe("https://api.github.com/markdown");
  expect(JSON.parse(String(requests[0].init.body))).toEqual({ text: "hello", mode: "gfm" });
  expect((requests[0].init.headers as Record<string, string>).Authorization).toBe(
    "Bearer example-token",
  );
});
test("live reference reports HTTP errors without exposing the token", async () => {
  const fetcher = async () => new Response("failed", { status: 403 });
  await expect(
    fetchGitHubHtml("hello", { token: "example-token", fetch: fetcher as typeof fetch }),
  ).rejects.toThrow("HTTP 403");
});
