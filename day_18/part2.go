package main

import (
	"math"
	"strings"
)

func partTwo(input []byte, size int, start int) string {
	lines := strings.Split(string(input), "\n")
	var coordinates string

	for i := start; i < len(lines); i++ {
		result := partOne(input, size, i)
		if result == math.MaxInt32 {
			break
		}
		coordinates = lines[i]
	}

	return coordinates
}
