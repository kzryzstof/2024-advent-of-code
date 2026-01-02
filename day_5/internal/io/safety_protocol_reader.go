package io

import (
	"bufio"
	"day_5/internal/abstractions"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SafetyProtocolReader struct {
	inputFile *os.File
}

func NewReader(
	filePath string,
) (*SafetyProtocolReader, error) {

	inputFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}

	return &SafetyProtocolReader{
		inputFile,
	}, nil
}

func (r *SafetyProtocolReader) Read() (*abstractions.SafetyProtocol, *abstractions.Updates) {

	scanner := bufio.NewScanner(r.inputFile)

	safetyProtocol := r.extractSafetyProtocol(scanner)

	updates := r.extractUpdates(scanner)

	return safetyProtocol, updates
}

func (r *SafetyProtocolReader) extractSafetyProtocol(
	scanner *bufio.Scanner,
) *abstractions.SafetyProtocol {

	orderingRules := make([]abstractions.OrderingRule, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		rules := strings.Split(line, "|")

		fromPage, _ := strconv.ParseInt(rules[0], 10, 64)
		toPage, _ := strconv.ParseInt(rules[1], 10, 64)

		orderingRules = append(
			orderingRules,
			abstractions.OrderingRule{
				From: abstractions.PageNumber(fromPage),
				To:   abstractions.PageNumber(toPage),
			},
		)
	}

	return abstractions.NewSafetyProtocol(orderingRules)
}

func (r *SafetyProtocolReader) extractUpdates(
	scanner *bufio.Scanner,
) *abstractions.Updates {

	updates := make([]*abstractions.Update, 0)

	updateId := 1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		numbers := strings.Split(line, ",")
		pageNumbers := make([]abstractions.PageNumber, len(numbers))

		for index, pageNumber := range numbers {
			number, _ := strconv.ParseInt(pageNumber, 10, 64)
			pageNumbers[index] = abstractions.PageNumber(number)
		}

		updates = append(
			updates,
			abstractions.NewUpdate(
				abstractions.UpdateId(updateId),
				pageNumbers,
			),
		)

		updateId++
	}

	return abstractions.NewUpdates(updates)
}
