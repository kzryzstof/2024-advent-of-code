package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"day_2/internal/abstractions"
)

type ReportsReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*ReportsReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &ReportsReader{
		inputFile,
	}, nil
}

func (r *ReportsReader) Read() []abstractions.Report {

	scanner := bufio.NewScanner(r.inputFile)

	reports := make([]abstractions.Report, 0)
	lineNumber := uint(1)

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		report, err := extractReport(lineNumber, line)

		if err != nil {
			os.Exit(1)
		}

		reports = append(
			reports,
			*report,
		)

		lineNumber++
	}

	return reports
}

func extractReport(
	lineNumber uint,
	line string,
) (*abstractions.Report, error) {

	levelValues := strings.Fields(line)

	levels := make([]abstractions.Level, len(levelValues))

	for index, levelValue := range levelValues {
		level, err := strconv.Atoi(levelValue)

		if err != nil {
			fmt.Printf("Error converting level '%s' to int: %v\n", levelValue, err)
			return nil, err
		}

		levels[index] = abstractions.Level(level)
	}

	return abstractions.NewReport(
		lineNumber,
		levels,
	), nil
}
