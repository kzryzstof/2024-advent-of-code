package main

import (
	"fmt"
	"os"
	"time"

	"day_3/internal/io"
)

func main() {

	startTime := time.Now()

	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the reader */
	reader := getReader(inputFile)

	/* Reads all the reports */
	program := reader.Read()

	fmt.Printf("Executing %d instructions\n", program.GetInstructionsCount())

	total := program.Execute()

	/* Prints the results */
	fmt.Printf("The total is %d\n", total)

	fmt.Printf("Execution time: %v\n", time.Since(startTime))
}

func getReader(
	inputFile []string,
) *io.ProgramReader {
	reader, err := io.NewReader(inputFile[0])

	if err != nil {
		fmt.Printf("Error parsing input file: %s\n", err)
		os.Exit(1)
	}

	return reader
}
