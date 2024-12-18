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

	// Trim trailing newlines
	input = bytes.TrimRight(input, "\n")

	fmt.Println("Part 1:", partOne(input, 71, 1024))
	fmt.Println("Part 2:", partTwo(input, 71, 1025))
}
