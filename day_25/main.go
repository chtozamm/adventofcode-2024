package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	input = bytes.TrimRight(input, "\n")

	fmt.Println("Part 1:", partOne(input))
	// fmt.Println("Part 2:", partTwo(input))
}
