import { expect, test } from "bun:test";
import { normalizeGithubImages } from "./support/github-images";

const proxy = "https://camo.githubusercontent.com/signature/encoded";
const original = "https://example.com/image.png?a=1&b=2";
const img = `<img src="${proxy}" alt="x" data-canonical-src="https://example.com/image.png?a=1&amp;b=2" style="max-width: 100%;">`;

test("restores the original URL and removes only the generated proxy wrapper", () => {
  const html = `<p><a target="_blank" rel="noopener noreferrer nofollow" href="${proxy}">${img}</a></p>`;
  expect(normalizeGithubImages(html)).toBe(
    '<p><img src="https://example.com/image.png?a=1&amp;b=2" alt="x"></p>',
  );
});

test("preserves a user-authored link around a proxied image", () => {
  const html = `<p><a href="https://example.com/page" rel="nofollow">${img}</a></p>`;
  expect(normalizeGithubImages(html)).toBe(
    '<p><a href="https://example.com/page" rel="nofollow"><img src="https://example.com/image.png?a=1&amp;b=2" alt="x"></a></p>',
  );
});

test("leaves unrelated images, wrappers and additional image styles intact", () => {
  const noCanonical = `<a href="${proxy}"><img src="${proxy}"></a>`;
  expect(normalizeGithubImages(noCanonical)).toBe(noCanonical);
  const extraAttr = `<a target="_blank" rel="noopener noreferrer nofollow" href="${proxy}" title="kept">${img}</a>`;
  expect(normalizeGithubImages(extraAttr)).toContain('title="kept"');
  const style = `<img src="${proxy}" data-canonical-src="${original.replaceAll("&", "&amp;")}" style="width: 20px;">`;
  expect(normalizeGithubImages(style)).toContain('style="width: 20px;"');
});
