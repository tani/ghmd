# Style guide regression audit

This audit compares the renderer with GitHub's GFM HTML response, including cases where GitHub's behavior is surprising. A mismatch is reported as `fail`; it is not automatically classified as a GitHub bug. Batch outcomes apply to the whole batch, not each member separately.

After `bun install --frozen-lockfile` and `bun run setup`:

```sh
# Uses saved responses for the authored synthetic cases; makes no network calls.
bun run audit:styleguide

# Obtain current GitHub responses for the synthetic cases, without replacing snapshots.
bun run audit:styleguide --live

# Include the 24 original acid-test files from the exact pinned commit.
git clone https://github.com/nschloe/markdown-math-acid-test.git /tmp/gfm-math-acid-test
git -C /tmp/gfm-math-acid-test checkout 1a672974525509f26da28abbae230999b49e2f3d
bun --env-file=../.env scripts/styleguide-audit.ts --live --acid-dir /tmp/gfm-math-acid-test --acid-only
```

The fixed upstream revision contains no license file, so the original inputs are not redistributed here. The audit verifies all 24 input hashes against `acid-manifest.json`; a different revision or altered file is rejected. Case 06 is tested as written and as two separately extracted variants. Other safe-to-group inputs are joined with identifiable headings, so their results are **whole-document comparisons**, not proofs that every member behaves identically in isolation. Case 20 is isolated because unclosed HTML can affect subsequent inputs. The synthetic footnote cases are isolated because footnote identifiers have document scope.

Live mode sends the synthetic and (when requested) public upstream input text to GitHub. It batches requests, stops further requests on HTTP 403 or 429, and accepts `GITHUB_PAT` (or `GITHUB_TOKEN`) for authenticated API access. When the `.env` file is at the workspace root, run Bun with `--env-file=../.env` from the renderer directory. `--capture --live` explicitly refreshes only the synthetic snapshots. `--only ID` runs one bundle; `--acid-only` excludes the synthetic batches. `--reference-report PATH` reuses expected responses from a previous JSON report after verifying input hashes, allowing offline rechecks without additional API calls. The command exits nonzero on mismatches or missing references. A JSON report with exact synthetic input, expected and observed HTML, source hashes, versions, date, and results is written to `/tmp/gfm-styleguide-report.json`; set `STYLEGUIDE_REPORT` to change that path. For upstream inputs, the report records file names and hashes and the original source remains in the pinned checkout.

The HTML comparison fixes per-render IDs and reuses signed image URLs found in the reference response. It does not test Camo signing. The audit does not claim to test browser layout, Mermaid rendering, mobile, visual math, export, unavailable fonts, or streaming; the report lists these as not-run or unsupported. Successful HTML comparisons are not visual or semantic validation. See `GFM_STYLEGUIDE.md` Section 8 for the complete policy and limitations.

## Known differences as Markdown/HTML pairs

`../fixtures/known-differences/` contains the exact synthetic batch inputs and their saved GitHub HTML responses for synthetic mismatches identified in the earlier audit. Three additional pairs isolate task-list, nested-alert and Mermaid sections extracted from those saved batch responses; GitHub was not separately queried for the isolated sections, so their HTML is a batch-derived reference. The `.html` files are GitHub's expected output, not this renderer's output. The batch boundaries are part of the input and must remain in place: further splitting a batch may change footnote, HTML or math behavior. `manifest.json` records the case IDs, capture context and hashes.

`bun test test/known-differences.test.ts` checks fixture integrity and asserts byte-for-byte equality with each saved response. The audit command reports any new mismatch as `fail` and exits nonzero. The upstream acid-test inputs and the HTML derived from them are not redistributed because the pinned repository contains no license. Authenticated responses are retained only in the ignored local `.cache/github-acid-reference.json`. The `math-boundaries` pair is an independently authored synthetic fixture captured from GitHub in one request.

## Finding the failing case

The offline `bun test test/styleguide-cases.test.ts` suite has a named test for each synthetic case and rule ID. It renders each original batch once and checks the section for that case against the saved GitHub response. These results are labeled **batch context**: they retain the surrounding document, and are not claims about rendering the case alone. The original full-batch comparison remains in the audit.

```sh
# Run one named synthetic case (the test still renders its original batch).
bun test test/styleguide-cases.test.ts -t tasks-markers-nesting

# Diagnose one case in its original batch; no network is used.
bun scripts/styleguide-audit.ts --only tasks-markers-nesting

# Diagnose one upstream case in its batch using an existing local report.
bun scripts/styleguide-audit.ts --acid-dir /tmp/gfm-math-acid-test \
  --reference-report .cache/github-acid-reference.json \
  --only 12-inline-math-succeeded-by-nonwhitespace.md

# Only when an independently rendered result is needed: one API request.
STYLEGUIDE_REPORT=.cache/isolated-task.json bun --env-file=../.env \
  scripts/styleguide-audit.ts --live --isolate tasks-markers-nesting

# Repeat that independent comparison without another API request.
bun scripts/styleguide-audit.ts --isolate tasks-markers-nesting \
  --reference-report .cache/isolated-task.json
```

`--isolate` also accepts an upstream filename when `--acid-dir` is supplied. Its result is labeled **whole document** and the input is the exact individual file, rather than an excerpt from a batch. Live mode sends that selected input to GitHub. It does not silently issue 24 requests; isolate only the cases that require it. The report includes each case's status, comparison scope, and first differing line, column and values. If a boundary is missing, duplicated or reordered, case attribution is marked `not-run` and the full document must be inspected. A `--only` selection still checks the entire original batch, so another case's failure can make the command exit nonzero.
