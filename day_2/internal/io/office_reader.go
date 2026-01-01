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

type OfficeReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*OfficeReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &OfficeReader{
		inputFile,
	}, nil
}

func (r *OfficeReader) Read() *abstractions.Office {

	scanner := bufio.NewScanner(r.inputFile)

	leftLocationIds := make([]abstractions.LocationId, 0)
	rightLocationIds := make([]abstractions.LocationId, 0)

	for scanner.Scan() {

		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		leftLocationId, rightLocationId, err := extractLocationIds(line)

		if err != nil {
			os.Exit(1)
		}

		leftLocationIds = append(
			leftLocationIds,
			leftLocationId)

		rightLocationIds = append(
			rightLocationIds,
			rightLocationId)
	}

	return abstractions.NewOffice(
		abstractions.NewList(leftLocationIds),
		abstractions.NewList(rightLocationIds),
	)
}

func extractLocationIds(
	line string,
) (abstractions.LocationId, abstractions.LocationId, error) {

	locationIds := strings.Fields(line)

	leftLocationId, err := strconv.Atoi(locationIds[0])

	if err != nil {
		fmt.Printf("Error converting location id '%s' to int: %v\n", locationIds[0], err)
		return 0, 0, err
	}

	rightLocationId, err := strconv.Atoi(locationIds[1])

	if err != nil {
		fmt.Printf("Error converting location id '%s' to int: %v\n", locationIds[0], err)
		return 0, 0, err
	}

	return abstractions.LocationId(leftLocationId),
		abstractions.LocationId(rightLocationId),
		nil
}
