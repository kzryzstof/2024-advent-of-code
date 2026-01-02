package main

import (
	"fmt"
	"os"
	"time"

	"day_5/internal/io"
)

func main() {

	startTime := time.Now()

	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := getReader(inputFile)

	/* Reads all the reports */
	safetyProtocol, updates := reader.Read()

	middlePagesCount := updates.CheckUpdates(safetyProtocol)

	/* Prints the results */
	fmt.Printf("The sum of middle pages is %d\n", middlePagesCount)

	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}

func getReader(
	inputFile []string,
) *io.SafetyProtocolReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		fmt.Printf("Error parsing input file: %s\n", err)
		os.Exit(1)
	}

	return reader
}
