# Advent of Code 2024 — Day 1 (Part 2)

This folder contains a Go solution for **Day 1 — Part 2** (often referred to as “Day 2” in this repo’s notes): compute the **similarity score** between two lists of location IDs.

## Problem recap (similarity score)

You’re given a file with two columns of integers (left list and right list). The similarity score is:

- For each number `x` in the **left** list, count how many times `x` appears in the **right** list.
- Add `x * countRight(x)` to the total.

Example:

```
3   4
4   3
2   5
1   3
3   9
3   3
```

Right list contains `3` three times, so each `3` in the left contributes `3 * 3 = 9`.

## Implementation overview

### Parsing (`internal/io/office_reader.go`)

- The input file is scanned line-by-line.
- Each line is split with `strings.Fields(line)`.
  - This splits on any run of whitespace and **drops empty parts**, which fits the puzzle input formatting.
- The two tokens are parsed with `strconv.Atoi` and appended to the left and right slices.

### Sorting (`internal/abstractions/sorted_location_ids.go`)

- Both slices are wrapped with `abstractions.NewList(ids)`.
- `NewList` sorts the slice in-place (`sort.Slice`), producing two sorted lists.

### Similarity score algorithm (`internal/algorithms/compute_similarity_score.go`)

The implementation leverages the fact that **both lists are sorted** and uses a two-pointer style scan:

- `locationIndex` walks the left list.
- `otherLocationIndex` walks the right list and **never moves backward**.

For each `locationId` from the left list:

1. Advance through the right list while `otherLocationId < locationId`.
2. When `otherLocationId == locationId`, count how many equal values are found (for the current scan position) and store that as `similarNumbers`.
3. If `otherLocationId > locationId`, stop early for this left value (because no match exists at the current right pointer position).
4. Add `uint64(locationId) * similarNumbers` to the running total.

The return type is `abstractions.SimilarityScore` (`uint64`).

Notes:
- The code checks both lists have the same length and returns an error otherwise.
- Because `otherLocationIndex` only moves forward, the scan is efficient on already-sorted inputs.

### Entry point (`cmd/main.go`)

- Reads the input file path from the first CLI argument.
- Parses and sorts both lists into an `Office`.
- Calls `algorithms.CompareSimilarityScore`.
- Prints the similarity score and execution time.

## How to run

From the `day_1` directory:

```bash
go run ./cmd ./input.txt
```

Expected output shape:

- `Similarity score between the two lists: <number>`
- `Execution time: <duration>`
