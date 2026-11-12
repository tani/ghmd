# Native pipeline regressions

Each Markdown/HTML pair preserves the output of the previous renderer with a fixed `runId` of 32 `a` characters. These are migration
regressions, not additional GitHub API captures. The existing API captures remain
under `../e2e/` and are unchanged.

The cases cover paragraph closing, table/template insertion, nested links,
attribute identity after relocation, text boundaries, Unicode math, raw input
filtering and untrusted tag names. Tests compare exact strings, including
serialization and intentionally malformed nesting.
