package io

import (
	"bufio"
	"day_3/internal/abstractions"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ProgramReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*ProgramReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &ProgramReader{
		inputFile,
	}, nil
}

func (r *ProgramReader) Read() *abstractions.Program {

	scanner := bufio.NewScanner(r.inputFile)

	/* Reads all the file content */
	lines := ""

	for scanner.Scan() {
		lines += scanner.Text()
	}

	instructions := make([]abstractions.Instruction, 0)

	re := regexp.MustCompile(`do\(\)|don\'t\(\)|mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatchIndex(lines, -1)

	isEnabled := true

	for _, m := range matches {

		// m[0],m[1] = full match range
		// m[2],m[3] = group 1 range (left number)
		// m[4],m[5] = group 2 range (right number)
		operation := lines[m[0]:m[1]]

		if operation == "do()" {
			isEnabled = true
			continue
		} else if operation == "don't()" {
			isEnabled = false
			continue
		}

		if !isEnabled {
			continue
		}

		left := lines[m[2]:m[3]]
		right := lines[m[4]:m[5]]

		leftOperand, err := strconv.Atoi(left)

		if err != nil {
			fmt.Printf("Error converting left operand '%s' to int: %v\n", left, err)
			panic(err)
		}

		rightOperand, err := strconv.Atoi(right)

		if err != nil {
			fmt.Printf("Error converting right operand '%s' to int: %v\n", left, err)
			panic(err)
		}

		instructions = append(instructions, abstractions.Instruction{
			Operation:    "mul",
			LeftOperand:  int64(leftOperand),
			RightOperand: int64(rightOperand),
		})
	}

	return abstractions.NewProgram(
		instructions,
	)
}
