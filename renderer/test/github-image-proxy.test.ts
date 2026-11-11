import { expect, test } from "bun:test";
import { createGitHubImageProxy } from "../src/github-image-proxy";

test("looks up and caches signed URLs without trusting unrelated response elements", async () => {
  let calls = 0;
  const mock = async (_url: any, init: any) => {
    calls++;
    expect(JSON.parse(init.body)).toEqual({
      text: '<img src="https://example.com/a?x=1&amp;y=2">',
      mode: "gfm",
    });
    return new Response(
      '<img src="https://camo.githubusercontent.com/sig/encoded" data-canonical-src="https://example.com/a?x=1&amp;y=2">',
    );
  };
  const proxy = createGitHubImageProxy({ fetch: mock as typeof fetch });
  const urls = await Promise.all([
    proxy("https://example.com/a?x=1&y=2"),
    proxy("https://example.com/a?x=1&y=2"),
  ]);
  expect(urls).toEqual([
    "https://camo.githubusercontent.com/sig/encoded",
    "https://camo.githubusercontent.com/sig/encoded",
  ]);
  expect(calls).toBe(1);
});

test("does not cache failed lookups and rejects unexpected proxy hosts", async () => {
  let calls = 0;
  const proxy = createGitHubImageProxy({
    fetch: (async () => {
      calls++;
      return new Response(
        '<img src="https://evil.example/x" data-canonical-src="https://example.com/a">',
      );
    }) as typeof fetch,
  });
  await expect(proxy("https://example.com/a")).rejects.toThrow("did not return");
  await expect(proxy("https://example.com/a")).rejects.toThrow("did not return");
  expect(calls).toBe(2);
});
