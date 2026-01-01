# Advent of Code 2024 — Day 1 (Part 1)

This folder contains a Go solution for the Day 1 puzzle where you’re given two columns of **location IDs** (left list and right list) and must compute the **total distance** between the lists.

## Problem recap

- Each input line contains two integers: the left location ID and the right location ID.
- To compare the lists, you:
  1. sort the left list ascending
  2. sort the right list ascending
  3. pair items by index (smallest with smallest, 2nd smallest with 2nd smallest, …)
  4. sum `abs(left[i] - right[i])` over all pairs

## Implementation overview

### Parsing (`internal/io/office_reader.go`)

- The input file is scanned line-by-line.
- Each line is split with `strings.Fields(line)`.
  - This splits on any run of whitespace and **automatically drops empty parts**, which is handy because the puzzle input uses variable spacing.
- The two resulting tokens are parsed with `strconv.Atoi` and appended to the left and right slices.

### Data model + sorting (`internal/abstractions/*`)

- The two slices are wrapped in `SortedLocationIds` via `abstractions.NewList(ids)`.
- `NewList` sorts the slice in place using `sort.Slice`, ensuring both lists are ordered before comparison.

### Distance computation (`internal/algorithms/compare_lists.go`)

- First, the code verifies both lists have the same length.
- Then it iterates from `0..n-1` and accumulates:

  `total += abs(float64(left[i]) - float64(right[i]))`

- The result is returned as `abstractions.Distance` (`uint64`).

### Entry point (`cmd/main.go`)

- Reads the input file path from the first CLI argument.
- Builds the in-memory `Office` (two sorted lists).
- Calls `algorithms.CompareLists` and prints:
  - the computed distance
  - execution time

## How to run

From the `day_1` directory:

```bash
go run ./cmd ./input.txt
```

(You can replace `./input.txt` with any file in the same two-column format.)

