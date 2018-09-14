package main

import (
	"flag"
	"fmt"

	sort "github.com/nqbao/learn-go/externalmergesort"
)

func main() {
	var command = flag.String("command", "", "Which command to run")
	var input = flag.String("input", "", "Input file")
	var count = flag.Int("count", 0, "How many number to generate")
	var chunkSize = flag.Float64("chunk", 1.0, "Chunk size")
	flag.Parse()

	if *command == "" || *input == "" {
		fmt.Printf("Usage: sort -command [command] -input [input_file]\n")
	}

	// generate the test data to a file
	if *command == "generate" {
		sort.GenerateData(input, *count)
	} else if *command == "validate" {
		sort.ValidateData(input)
	} else if *command == "sort" {
		sort.ExternalMergeSort(input, *chunkSize)
	}
}
