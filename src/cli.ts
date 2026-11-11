import { render } from "./index";
const args = process.argv.slice(2);
try {
  let runId: string | undefined, path: string | undefined;
  for (let i = 0; i < args.length; i++) {
    const arg = args[i];
    if (arg === "--run-id") {
      if (runId !== undefined || !args[i + 1]) throw new Error("--run-id requires one value");
      runId = args[++i];
    } else if ((arg.startsWith("-") && arg !== "-") || path !== undefined) {
      throw new Error("Usage: render [--run-id ID] [input.md|-]");
    } else path = arg;
  }
  const source = path && path !== "-" ? await Bun.file(path).text() : await Bun.stdin.text();
  process.stdout.write(await render(source, { runId }));
} catch (error) {
  console.error(String(error));
  process.exitCode = 1;
}
