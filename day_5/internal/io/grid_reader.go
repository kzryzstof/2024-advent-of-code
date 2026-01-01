package io

import (
	"bufio"
	"day_5/internal/abstractions"
	"fmt"
	"os"
)

type GridReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*GridReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &GridReader{
		inputFile,
	}, nil
}

func (r *GridReader) Read() *abstractions.Grid {

	scanner := bufio.NewScanner(r.inputFile)

	/* Reads all the file content in-memory */
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	letters := make([][]int32, 0)

	for row, line := range lines {
		letters = append(letters, make([]int32, len(line)))
		for col, letter := range line {
			letters[row][col] = letter
		}
	}

	return abstractions.NewGrid(
		letters,
	)
}
