package main

import (
	"day_1/internal/algorithms"
	"fmt"
	"os"
	"time"

	"day_1/internal/io"
)

func main() {

	startTime := time.Now()

	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	instructionsParser := getReader(inputFile)

	/* Reads all the lists from the office */
	office := instructionsParser.Read()

	distance, err := algorithms.CompareLists(office)

	if err != nil {
		fmt.Printf("Error comparing lists: %s\n", err)
		os.Exit(1)
	}

	/* Prints the results */
	fmt.Printf("Distance between the two lists: %d\n", distance)

	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}

func getReader(
	inputFile []string,
) *io.OfficeReader {
	instructionsReader, err := io.NewReader(inputFile[0])

	if err != nil {
		fmt.Printf("Error parsing input file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Parser initialized: %v\n", instructionsReader)
	return instructionsReader
}
