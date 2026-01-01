package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"day_2/internal/abstractions"
)

const (
	DefaultInstructionsSliceCapacity = 1000
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

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		report, err := extractReport(line)

		if err != nil {
			os.Exit(1)
		}

		reports = append(
			reports,
			*report,
		)
	}

	return reports
}

func extractReport(
	line string,
) (*abstractions.Report, error) {

	levelValues := strings.Fields(line)

	levels := make([]abstractions.Level, len(levelValues))

	for _, levelValue := range levelValues {
		level, err := strconv.Atoi(levelValue)

		if err != nil {
			fmt.Printf("Error converting level '%s' to int: %v\n", levelValue[0], err)
			return nil, err
		}

		levels = append(levels, abstractions.Level(level))
	}

	return &abstractions.Report{
		Levels: levels,
	}, nil
}
