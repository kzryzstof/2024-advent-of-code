package main

import (
	"day_2/internal/algorithms"
	"fmt"
	"os"
	"time"

	"day_2/internal/io"
)

func main() {

	startTime := time.Now()

	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := getReader(inputFile)

	/* Reads all the reports */
	reports := reader.Read()

	safeReportsCount := algorithms.CountSafeReports(reports)

	/* Prints the results */
	fmt.Printf("Safe reports: %d\n", safeReportsCount)

	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}

func getReader(
	inputFile []string,
) *io.ReportsReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		fmt.Printf("Error parsing input file: %s\n", err)
		os.Exit(1)
	}

	return reader
}
