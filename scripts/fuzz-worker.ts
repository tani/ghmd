import { render } from "../src/quirks";
const source = await Bun.stdin.text();
try {
  const html = await render(source, { runId: "0123456789abcdef0123456789abcdef" });
  process.stdout.write(html);
} catch (error) {
  console.error(error instanceof Error ? `${error.name}: ${error.message}` : "Render failed");
  process.exitCode = 1;
}
