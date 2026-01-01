# 2024 Advent of Code

This repository contains my solutions for the **Advent of Code 2024** programming puzzles, implemented in Go (a language I’m currently learning).

Each day lives in its own folder (for example `day_1/`) with its own `go.mod`, a small command-line entrypoint under `cmd/`, and an `internal` package that holds the core logic.

## Status / structure

At the moment, only **Day 1** is implemented:

- `day_1/` – **Day 1: Historian Hysteria (location ID reconciliation)**
  - Part 1: sort both lists and sum the pairwise absolute differences ("total distance")
  - Part 2: compute the similarity score by counting how often each left ID appears in the right list

See `day_1/README.md` for details about the approach and how to run it.

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
