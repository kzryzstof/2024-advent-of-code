# Day 5: Print Queue

This module solves **Advent of Code 2024 — Day 5 (Part 1)**.

The input is split in two sections:

1. **Ordering rules** (`X|Y`): if an update contains both pages `X` and `Y`, then `X` must appear **before** `Y`.
2. **Updates** (comma-separated page numbers): each line is one print job.

The goal is to:

- Identify which updates are **already ordered correctly**.
- For each valid update, take its **middle page number**.
- Return the **sum of those middle page numbers**.

---

## Solution overview

### 1) Parse the input into a protocol + updates

Implemented in `internal/io/safety_protocol_reader.go`:

- Read ordering rules until the first blank line.
- Read updates after the blank line.

Data structures:

- `OrderingRule{From, To}` (`internal/abstractions/ordering_rule.go`)
- `SafetyProtocol` holding all rules + an optimized representation (`internal/abstractions/safety_protocol.go`)
- `Update` + `Updates` (`internal/abstractions/update.go`, `internal/abstractions/updates.go`)

---

### 2) Pre-compute an optimized rule map

A direct “check every rule for every update” approach would be slow and noisy.
Instead, `SafetyProtocol` builds an adjacency-style map once:

- `optimizedRules[from] = []to`

So for any page `from`, we can quickly know which pages must appear **after** it.

This is done in `getOptimizedRules()` in `internal/abstractions/safety_protocol.go`.

---

### 3) Validate an update

Implemented in `Update.IsOrderValid()` (`internal/abstractions/update.go`).

For each page in the update, we check that none of the pages that appear **before it** violate the protocol.

Concretely:

- For a candidate page `p`, we look at the pages that appear earlier in the update (`numbersBefore`).
- If any earlier page is found in `optimizedRules[p]`, it means we are printing a page **before `p`** that should be printed **after `p`**.
- That breaks a rule, so the update is invalid.

This check is done by `SafetyProtocol.IsValid(pageNumber, numbersBefore)`.

Notes:

- Rules involving pages that don’t exist in the update are naturally ignored.
- The code iterates from the end to the beginning, but the logic is the same regardless of direction.

---

### 4) Sum the middle page of valid updates

Implemented in `Updates.CheckUpdates()` (`internal/abstractions/updates.go`):

- Skip invalid updates.
- For valid ones, compute the middle index:

  - `middle := (len(pages)-1)/2` (integer division)

- Add that page number to the running sum.

The final value printed is:

- `The sum of middle pages is <N>`

---

## How to run

From the `day_5` folder:

```bash
make run ARGS=./input.txt
```

Or directly:

```bash
go run ./cmd ./input.txt
```

---

## Files of interest

- `cmd/main.go`: wiring + timing + printing the result.
- `internal/io/safety_protocol_reader.go`: parses rules and updates.
- `internal/abstractions/safety_protocol.go`: protocol representation + optimized rule map.
- `internal/abstractions/update.go`: order validation + middle page extraction.
- `internal/abstractions/updates.go`: filters valid updates and sums middle pages.

