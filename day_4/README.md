# Advent of Code 2024 — Day 4 (Part 2)

This folder contains a Go solution for **Day 4: Ceres Search**.

Part 2 changes the task: instead of searching for the word `XMAS` in straight lines, you must find an **X-MAS** shape:

- two occurrences of `MAS` that form an `X`
- each `MAS` can be read forward (`M-A-S`) or backward (`S-A-M`)

Visually, an X-MAS is centered on an `A` and uses the four diagonal neighbors:

- top-left, top-right, bottom-right, bottom-left

## Implementation overview

### 1) Parse the grid (`internal/io/grid_reader.go`)

- The input file is scanned line-by-line using a `bufio.Scanner`.
- Each line becomes one row of the grid.
- The grid is stored in memory as a 2D slice of letters: `[][]int32`.

### 2) Count X-MAS patterns (`internal/abstractions/grid.go`)

The solver counts patterns with `Grid.CountXmasPatterns()`.

High-level idea:

- An X-MAS must be centered on an `A`.
- Around that `A`, the four diagonal cells must contain `M` and `S` in one of the valid configurations that corresponds to:
  - `MAS` on one diagonal and `MAS` (or `SAM`) on the other diagonal.

#### Directions

The four diagonal directions are represented as vectors:

- `UpLeft  = (-1, -1)`
- `UpRight = (-1,  1)`
- `DownRight = (1,  1)`
- `DownLeft  = (1, -1)`

They are stored in the `Directions` slice in that order.

#### Pattern matching

For each cell `(row, col)` that contains `A`, we look at the four diagonal letters around it and build an implicit 4-character “signature” in this order:

`UpLeft, UpRight, DownRight, DownLeft`

The implementation checks the signature against 4 allowed patterns:

- `MSSM`
- `MMSS`
- `SMMS`
- `SSMM`

These are exactly the cases where:

- one diagonal is `M-A-S` and the other diagonal is also `M-A-S`,
  or either/both are reversed as `S-A-M`.

As soon as one pattern matches for an `A` center, we count **one** X-MAS and move on.

### 3) Bounds checking

Each diagonal lookup verifies the target coordinate is inside the grid. If any required diagonal cell is out of bounds, that candidate center cannot form an X-MAS.

## How to run

From the `day_4` directory:

```bash
go run ./cmd ./input.txt
```

Expected output shape:

- `The 'XMAS' pattern appeared <n> times`
- `Execution time: <duration>`
