package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
}
