# Advent of Code 2024 — Day 3 (Part 1)

This folder contains a Go solution for **Day 3: Mull It Over**.

The puzzle input is a corrupted “program memory” that contains lots of garbage characters, plus some valid multiplication instructions of the form:

- `mul(X,Y)`
  - `X` and `Y` are positive integers written with digits
  - other similar-looking patterns (extra spaces, wrong brackets, punctuation inside, etc.) must be ignored

The goal is to scan the corrupted memory, extract only the valid `mul(...)` instructions, evaluate them, and return the **sum of all products**.

## Implementation overview

### 1) Read the full program memory (`internal/io/program_reader.go`)

The code reads the whole input file into a single string:

- `bufio.Scanner` iterates over every line
- each line is appended to one big `lines` string
- line breaks are effectively removed (which is fine because valid instructions can appear anywhere in the stream)

### 2) Extract valid instructions with a regex

We use a regular expression that matches only the exact valid syntax:

- `mul\((\d+),(\d+)\)`

Meaning:

- literal `mul(`
- capture 1+ digits for the left operand
- a literal comma
- capture 1+ digits for the right operand
- literal `)`

This intentionally **does not** match invalid variants like:

- `mul ( 2 , 4 )` (spaces)
- `mul[3,7]` (wrong brackets)
- `mul(4*,...` (non-digits)

The code calls `FindAllStringSubmatchIndex` to get the exact ranges for the full match and each captured group, then slices the input string to retrieve the operand substrings.

### 3) Parse operands and build instructions

For each match:

- parse the captured groups with `strconv.Atoi`
- create an `abstractions.Instruction{ Operation: "mul", LeftOperand: ..., RightOperand: ... }`

### 4) Execute and sum (`internal/abstractions/program.go`)

- `Program.Execute()` iterates over all extracted instructions
- `Instruction.Execute()` multiplies operands for the `mul` operation
- all results are summed into a single `int64` total

The CLI prints:

- how many `mul` instructions were extracted
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

