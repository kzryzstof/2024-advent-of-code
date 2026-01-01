package main

import (
	"fmt"
	"os"

	"day_1/internal/io"
)

func main() {
	inputFile := os.Args[1:]
	fmt.Println(inputFile)

	/* 	Initializes the parser */
	instructionsParser := getReader(inputFile)

	/* Reads all the lists from the office */
	office := instructionsParser.Read()

	/* Prints the results */
	fmt.Printf("List: %d locations. Other list: %d locations\n", office.List.Len(), office.OtherList.Len())
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
