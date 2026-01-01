# Advent of Code 2024 — Day 2 (Part 2)

This folder contains a Go solution for **Day 2: Red-Nosed Reports**.

The input is a set of *reports*, one report per line. Each report is a list of integer **levels** separated by spaces.

The goal here is to count how many reports are **safe** when the reactor’s **Problem Dampener** is enabled.

## Safety rules (with the Problem Dampener)

A report counts as *safe* if either:

- it is already safe under the original rules, **or**
- removing **exactly one** level (any position) makes it safe.

A report is safe under the original rules if:

1. The levels are all **strictly** monotonically increasing or all **strictly** monotonically decreasing.
2. For every adjacent pair, the absolute difference is between **1 and 3** (inclusive).

So:
- equal adjacent numbers (`delta = 0`) are not allowed,
- jumps larger than 3 are not allowed,
- switching from increasing to decreasing (or vice-versa) is not allowed.

## Implementation overview

### Parsing input (`internal/io/reports_reader.go`)

- The file is scanned line-by-line with a `bufio.Scanner`.
- Each line is split using `strings.Fields(line)` (splits on whitespace and drops empty fields).
- Each value is parsed with `strconv.Atoi` and stored as an `abstractions.Level`.
- A `Report` is created with an id (line number) and the slice of levels.

### Report model and safety check (`internal/abstractions/report.go`)

Safety is determined by `Report.GetStatus()`.

It works in two phases:

1) **Check the report as-is**

- `GetStatus()` calls an internal helper `getStatus(levels)`.
- `getStatus` determines whether the sequence is strictly monotonic and each adjacent delta is within bounds.
- If the report is safe, we return `StatusSafe` immediately.

2) **If unsafe, try the Problem Dampener (remove one level)**

- If the original report is unsafe, `GetStatus()` tries removing one level at every index:
  - build a new `alteredLevels` slice that skips `removedIndex`
  - run `getStatus(alteredLevels)`
  - if any removal yields `StatusSafe`, the report counts as safe

If none of the single-level removals produce a safe report, it remains `StatusUnsafe`.

Complexity note: this is a simple brute-force approach: for a report of length `n`, we may test up to `n` altered reports, each taking `O(n)` to check, so worst-case `O(n²)` per report. For typical AoC input sizes this is fast enough and keeps the code straightforward.

### Counting safe reports (`internal/algorithms/count_safe_reports.go`)

- Iterate over all reports.
- Call `report.GetStatus()`.
- Count how many are `StatusSafe`.

### Entry point (`cmd/main.go`)

- Reads the input file path from the first CLI argument.
- Parses reports from the file.
- Counts safe reports and prints the result.

## How to run

From the `day_2` directory:

```bash
go run ./cmd ./input.txt
```

Expected output shape:

- `Safe reports: <number>`
- `Execution time: <duration>`
