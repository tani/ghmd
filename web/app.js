"use strict";

const source = document.getElementById("source");
const status = document.getElementById("status");
const preview = document.getElementById("preview");
const htmlOutput = document.getElementById("html-output");
const renderView = document.getElementById("render-view");
const lintView = document.getElementById("lint-view");
const diffView = document.getElementById("diff-view");
const lintSummary = document.getElementById("lint-summary");
const lintResults = document.getElementById("lint-results");
const diffSummary = document.getElementById("diff-summary");
const diffResults = document.getElementById("diff-results");
const inputStats = document.getElementById("input-stats");
const tabs = Array.from(document.querySelectorAll(".tab"));

let mode = "render";
let ready = false;
let timer = 0;
let runSerial = 0;

function setStatus(message, kind = "") {
  status.textContent = message;
  status.className = "status" + (kind ? " " + kind : "");
}

function setMode(next) {
  mode = next;
  for (const tab of tabs) {
    const active = tab.dataset.mode === next;
    tab.classList.toggle("active", active);
    tab.setAttribute("aria-selected", String(active));
  }
  renderView.hidden = next !== "render";
  lintView.hidden = next !== "lint";
  diffView.hidden = next !== "diff";
  schedule();
}

function updateStats() {
  const bytes = new Blob([source.value]).size;
  const lines = source.value ? source.value.split("\n").length : 0;
  inputStats.textContent = `${bytes} bytes · ${lines} lines`;
}

function clearList(list) {
  list.replaceChildren();
}

function addFinding(list, title, subtitle, body) {
  const item = document.createElement("li");
  item.className = "finding";
  const head = document.createElement("div");
  head.className = "finding-head";
  const label = document.createElement("span");
  label.className = "finding-title";
  label.textContent = title;
  head.append(label);
  if (subtitle) {
    const meta = document.createElement("span");
    meta.className = "finding-meta";
    meta.textContent = subtitle;
    head.append(meta);
  }
  item.append(head);
  if (body) {
    const pre = document.createElement("pre");
    pre.textContent = body;
    item.append(pre);
  }
  list.append(item);
}

function renderPreview(html) {
  htmlOutput.textContent = html;
  preview.srcdoc = `<!doctype html>
<html><head><meta charset="utf-8"><style>
body{font:16px/1.55 -apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif;color:#1f2328;margin:0;padding:24px;background:#fff}
h1,h2,h3{border-bottom:1px solid #d8dee4;padding-bottom:.3em}
a{color:#0969da}code,pre{font-family:ui-monospace,SFMono-Regular,Menlo,monospace}
pre{background:#f6f8fa;padding:16px;overflow:auto;border-radius:6px}
code{background:#f6f8fa;padding:.2em .35em;border-radius:4px}
blockquote{margin:0;padding:0 1em;color:#57606a;border-left:.25em solid #d0d7de}
table{border-collapse:collapse}th,td{border:1px solid #d0d7de;padding:6px 13px}
.task-list-item{list-style:none}input{margin-right:.4em}
math-renderer{font-size:1.05em}
</style></head><body>${html}</body></html>`;
}

function handleError(message) {
  setStatus(message, "error");
}

function run() {
  if (!ready) return;
  const serial = ++runSerial;
  const input = source.value;
  const started = performance.now();
  setStatus("Running…");
  let result;
  try {
    result = JSON.parse(globalThis.ghmdRun(mode, input));
  } catch (error) {
    handleError(String(error));
    return;
  }
  if (serial !== runSerial) return;
  if (result.error) {
    handleError(result.error);
    return;
  }
  const elapsed = Math.round(performance.now() - started);
  if (mode === "render") {
    renderPreview(result.html || "");
    setStatus(`Rendered in ${elapsed} ms · ${(result.html || "").length} bytes`, "ok");
    return;
  }
  if (mode === "lint") {
    const findings = result.lint || [];
    clearList(lintResults);
    lintSummary.textContent = findings.length
      ? `${findings.length} finding${findings.length === 1 ? "" : "s"}`
      : "No findings.";
    for (const finding of findings) {
      addFinding(
        lintResults,
        `${finding.rule}`,
        `${finding.line}:${finding.column} · ${finding.severity}`,
        finding.message,
      );
    }
    setStatus(`Linted in ${elapsed} ms`, "ok");
    return;
  }
  const findings = result.diff || [];
  clearList(diffResults);
  diffSummary.textContent = findings.length
    ? `${findings.length} structural difference${findings.length === 1 ? "" : "s"}`
    : "No structural differences.";
  for (const finding of findings) {
    const body = `cmark-gfm: ${JSON.stringify(finding.want)}\nmd4c:     ${JSON.stringify(finding.got)}`;
    addFinding(diffResults, finding.path, "", body);
  }
  setStatus(`Compared in ${elapsed} ms`, "ok");
}

function schedule() {
  window.clearTimeout(timer);
  timer = window.setTimeout(run, 220);
}

for (const tab of tabs) {
  tab.addEventListener("click", () => setMode(tab.dataset.mode));
}
source.addEventListener("input", () => {
  updateStats();
  schedule();
});
source.addEventListener("keydown", (event) => {
  if ((event.metaKey || event.ctrlKey) && event.key === "Enter") {
    event.preventDefault();
    run();
  }
});

globalThis.ghmdReady = () => {
  ready = true;
  setStatus("Ready");
  run();
};

updateStats();
setMode("render");

const go = new Go();
WebAssembly.instantiateStreaming(fetch("ghmd.wasm"), go.importObject)
  .then((result) => go.run(result.instance))
  .catch((error) => {
    console.error(error);
    handleError(`Could not load ghmd.wasm: ${error.message || error}`);
  });
