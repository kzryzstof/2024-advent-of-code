# Advent of Code 2024 — Day 2 (Part 1)

This folder contains a Go solution for **Day 2: Red-Nosed Reports**.

The input is a set of *reports*, one per line. Each report is a list of integer **levels** separated by spaces.

The goal for Part 1 is to count how many reports are **safe**.

## Safety rules

A report is *safe* if **both** are true:

1. The levels are **strictly** monotonically increasing *or* **strictly** monotonically decreasing.
2. For every adjacent pair, the absolute difference is between **1 and 3** (inclusive).

That means equal adjacent numbers (delta = 0) are unsafe, and large jumps (delta > 3) are unsafe.

## Implementation overview

### Parsing input (`internal/io/reports_reader.go`)

- The file is scanned line-by-line with a `bufio.Scanner`.
- Each line is split using `strings.Fields(line)`.
  - This splits on runs of whitespace and drops empty fields.
- Each field is parsed with `strconv.Atoi` and stored as an `abstractions.Level`.
- A `Report` is created with an id (line number) and the slice of levels.

### Report model and safety check (`internal/abstractions/report.go`)

Safety is determined by `Report.GetStatus()`.

The key idea is to track the **sign of the deltas** between adjacent levels:

- For each adjacent pair, compute `delta = level - previousLevel`.
- Compute `abs(delta)` and reject the report if it’s outside `[MinDelta..MaxDelta]`.
  - `MinDelta = 1`, `MaxDelta = 3`
- Convert the delta into a sign:
  - negative (or zero) → `-1`
  - positive → `+1`

Then:

- The first valid delta sets the expected direction (`deltaSign`).
- Every subsequent delta must have the **same sign**, otherwise the report switches direction and is unsafe.

If all pairs pass these checks, the report is marked `StatusSafe`.

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

