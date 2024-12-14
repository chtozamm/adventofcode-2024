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

	fmt.Println("Part 1:", partOne(input, 101, 103))
	fmt.Println("Part 2:", partTwo(input, 101, 103))
}
