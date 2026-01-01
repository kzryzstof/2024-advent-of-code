# 2024 Advent of Code

This repository contains my solutions for the **Advent of Code 2024** programming puzzles, implemented in Go (a language I’m currently learning).

Each day lives in its own folder (for example `day_1/`) with its own `go.mod`, a small command-line entrypoint under `cmd/`, and an `internal` package that holds the core logic.

## Status / structure

At the moment, the following days are implemented:

- `day_1/` – **Day 1: Historian Hysteria (location ID reconciliation)**
  - Part 1: sort both lists and sum the pairwise absolute differences ("total distance")
  - Part 2: compute the similarity score by counting how often each left ID appears in the right list

- `day_2/` – **Day 2: Red-Nosed Reports (reactor safety reports)**
  - Part 1: count reports that are strictly increasing or strictly decreasing, with adjacent deltas in `[1..3]`
  - Part 2: same rules, but a report also counts as safe if removing **one** level makes it safe ("Problem Dampener")

- `day_3/` – **Day 3: Mull It Over (corrupted memory parser)**
  - Part 1: scan the corrupted program memory for valid `mul(X,Y)` instructions and sum all multiplication results
  - Part 2: also process `do()` / `don't()` to enable or disable future multiplications, then sum only the enabled `mul` results

See each day’s README (`day_1/README.md`, `day_2/README.md`, `day_3/README.md`) for details about the approach and how to run it.

## Running a day

To run a given day, `cd` into the corresponding folder and either use `make` (if present):

```bash
cd day_1
make run ARGS="input.txt"
```

or call `go run` directly on the command:

```bash
cd day_1
go run ./cmd input.txt
```
