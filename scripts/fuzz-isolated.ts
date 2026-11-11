import { join } from "node:path";

/** A native crash must not terminate the driver or expose the API token to the worker. */
export async function renderIsolated(
  source: string,
  options: { worker?: string; timeoutMs?: number } = {},
) {
  const worker = Bun.spawn(
    [process.execPath, "--no-env-file", options.worker ?? join(import.meta.dir, "fuzz-worker.ts")],
    {
      stdin: "pipe",
      stdout: "pipe",
      stderr: "pipe",
      env: Object.fromEntries(
        Object.entries(process.env).filter(
          ([key]) => !["GITHUB_PAT", "GITHUB_TOKEN"].includes(key),
        ),
      ),
    },
  );
  worker.stdin.write(source);
  worker.stdin.end();
  let timedOut = false;
  const timer = setTimeout(() => {
    timedOut = true;
    worker.kill();
  }, options.timeoutMs ?? 30_000);
  try {
    const [html, stderr, exitCode] = await Promise.all([
      new Response(worker.stdout).text(),
      new Response(worker.stderr).text(),
      worker.exited,
    ]);
    return { html, stderr, exitCode, timedOut };
  } finally {
    clearTimeout(timer);
  }
}
