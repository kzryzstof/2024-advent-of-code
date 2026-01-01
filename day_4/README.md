# Advent of Code 2024 — Day 4 (Part 1)

This folder contains a Go solution for **Day 4: Ceres Search**.

The input is a rectangular grid of uppercase letters (a word search). The goal is to count **all** occurrences of the word **`XMAS`**, where matches may be:

- horizontal, vertical, or diagonal
- forwards or backwards
- overlapping

## Implementation overview

### 1) Parse the grid (`internal/io/grid_reader.go`)

- The input file is scanned line-by-line using a `bufio.Scanner`.
- Each line becomes one row of the grid.
- The grid is stored as a 2D slice of `int32` (`[][]int32`) holding the rune values for each character.

### 2) Search for the word in all 8 directions (`internal/abstractions/grid.go`)

The search is implemented in `Grid.CountWord(word string)`:

1. Iterate over every cell `(row, col)`.
2. Only consider starting positions whose letter is `'X'` (the first character of `XMAS`).
3. For each `'X'`, try all 8 directions:

- Right, Left
- Up, Down
- UpRight, UpLeft
- DownRight, DownLeft

These directions are expressed as `Vector{X, Y}` step offsets and stored in the `Directions` slice.

### 3) Check a candidate match with bounds + character checks

`Grid.IsWordPresent(word, fromRow, fromCol, direction)` validates a match by:

- stepping 1, 2, 3 cells away from the start (for the remaining letters of `XMAS`)
- rejecting immediately if the step goes out of bounds
- comparing the grid letter at each stepped cell against `word[letterIndex]`

If all characters match, the word is present in that direction and the total counter is incremented.

This approach naturally handles:

- reversed matches (because directions include both forward and backward vectors)
- overlaps (each match is counted independently)

## How to run

From the `day_4` directory:

```bash
go run ./cmd ./input.txt
```

Expected output shape:

- `The word 'XMAS' appeared <n> times`
- `Execution time: <duration>`

