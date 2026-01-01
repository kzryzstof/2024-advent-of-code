# Advent of Code 2024 — Day 3 (Part 2)

This folder contains a Go solution for **Day 3: Mull It Over**.

The puzzle input is a corrupted “program memory” stream containing lots of garbage characters, plus some valid instructions.

## Problem recap

Valid instructions are:

- `mul(X,Y)` where `X` and `Y` are numbers written with digits
- `do()` which **enables** future `mul` instructions
- `don't()` which **disables** future `mul` instructions

Rules:

- Only the **most recent** `do()` / `don't()` instruction applies.
- At the beginning of the program, `mul` instructions are **enabled**.
- Invalid / corrupted sequences must be ignored.

Goal:

- Scan the corrupted memory left-to-right and sum the products of only the **enabled** `mul(X,Y)` instructions.

## Implementation overview

### 1) Read the full program memory (`internal/io/program_reader.go`)

The reader loads the entire file into a single string:

- `bufio.Scanner` reads each line
- lines are concatenated into one `lines` string

### 2) Tokenize with a single regex (do / don't / mul)

A single regular expression is used to find every relevant token in order:

- `do\(\)|don\'t\(\)|mul\((\d+),(\d+)\)`

This matches:

- `do()`
- `don't()` (note the literal apostrophe)
- `mul(<digits>,<digits>)` and captures the two operands

Because the regex is strict, it *does not* match invalid variants such as `mul[3,7]`, `mul ( 2 , 4 )`, or `mul(4*,...)`.

The code uses `FindAllStringSubmatchIndex` so we can:

- iterate matches left-to-right
- slice the original input string to extract the captured operand substrings

### 3) Single pass state machine: enabled/disabled

The reader maintains a boolean state:

- `isEnabled := true` initially

For each matched token:

- If it’s `do()`, set `isEnabled = true`.
- If it’s `don't()`, set `isEnabled = false`.
- Otherwise it must be a `mul(X,Y)` match:
  - if `isEnabled` is `false`, ignore it
  - if `isEnabled` is `true`, parse `X` and `Y` with `strconv.Atoi` and append a `mul` instruction

Only `mul` instructions that are enabled at the time they appear are kept.

### 4) Execute and sum (`internal/abstractions/program.go`)

- `Program.Execute()` iterates over the collected instructions.
- `Instruction.Execute()` multiplies operands for the `mul` operation.
- All products are summed into a single `int64` total.

The CLI prints:

- how many enabled `mul` instructions were collected
- the final sum
- execution time

## How to run

From the `day_3` directory:

```bash
go run ./cmd ./input.txt
```

Expected output shape:

- `Executing <n> instructions`
- `The total is <number>`
- `Execution time: <duration>`
