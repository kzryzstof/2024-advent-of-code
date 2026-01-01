package main

import (
	"fmt"
	"os"
	"time"

	"day_4/internal/io"
)

func main() {

	startTime := time.Now()

	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := getReader(inputFile)

	/* Reads all the reports */
	grid := reader.Read()

	count := grid.CountWord("XMAS")
	
	/* Prints the results */
	fmt.Printf("The word 'XMAS' appeared %d times", count)

	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}

func getReader(
	inputFile []string,
) *io.GridReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		fmt.Printf("Error parsing input file: %s\n", err)
		os.Exit(1)
	}

	return reader
}
