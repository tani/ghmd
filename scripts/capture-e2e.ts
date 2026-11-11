/** Capture only missing references; space requests and stop on rate limiting. */
import { createHash } from "node:crypto";
import { mkdir } from "node:fs/promises";
import cases from "./new-e2e-cases";
const token = process.env.GITHUB_PAT ?? process.env.GITHUB_TOKEN;
if (!token) throw new Error("Set GITHUB_PAT or GITHUB_TOKEN");
const delay = Number(process.env.GFM_CAPTURE_DELAY_MS ?? 650);
if (!Number.isFinite(delay) || delay < 250)
  throw new Error("GFM_CAPTURE_DELAY_MS must be at least 250");
const dir = "test/fixtures/e2e/cases";
const manifestPath = "test/fixtures/e2e/manifest.json";
const manifest = await Bun.file(manifestPath).json();
const sha = (s: string) => createHash("sha256").update(s).digest("hex");
await mkdir(dir, { recursive: true });
for (const item of cases) {
  const digest = sha(item.source);
  const entry = manifest.fixtures.find((f: any) => f.sourceSha256 === digest);
  if (entry) {
    if (!entry.cases.some((c: any) => c.id === item.id)) {
      entry.cases.push({ id: item.id, origin: item.origin, rules: item.rules });
      await Bun.write(manifestPath, JSON.stringify(manifest, null, 2) + "\n");
    }
    continue;
  }
  const response = await fetch("https://api.github.com/markdown", {
    method: "POST",
    headers: {
      Accept: "application/vnd.github+json",
      "Content-Type": "application/json",
      "X-GitHub-Api-Version": "2022-11-28",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({ text: item.source, mode: "gfm" }),
    signal: AbortSignal.timeout(30000),
  });
  if (!response.ok) {
    console.error(
      `Stopped at ${item.id}: HTTP ${response.status}; retry-after=${response.headers.get("retry-after") ?? "not provided"}; remaining=${response.headers.get("x-ratelimit-remaining") ?? "unknown"}`,
    );
    process.exitCode = 1;
    break;
  }
  const html = await response.text();
  await Bun.write(`${dir}/${item.id}.md`, item.source);
  await Bun.write(`${dir}/${item.id}.html`, html);
  manifest.fixtures.push({
    id: item.id,
    sourceSha256: digest,
    htmlSha256: sha(html),
    capturedAt: new Date().toISOString(),
    cases: [{ id: item.id, origin: item.origin, rules: item.rules }],
  });
  await Bun.write(manifestPath, JSON.stringify(manifest, null, 2) + "\n");
  console.log(`Captured ${item.id}`);
  await Bun.sleep(delay);
}
