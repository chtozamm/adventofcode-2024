package main

import (
	"strings"
)

func partOne(input []byte) int {
	possibleDesigns := 0
	parts := strings.Split(string(input), "\n\n")
	towels := make(map[string]bool)

	for _, towel := range strings.Fields(strings.ReplaceAll(parts[0], ",", "")) {
		towels[towel] = true
	}

	designs := strings.Split(parts[1], "\n")

	designMemo := make(map[string]bool)

	var isPossible func(string) bool
	isPossible = func(design string) bool {
		if possible, exists := designMemo[design]; exists {
			return possible
		}
		n := len(design)
		possiblePrefix := make([]bool, n+1)
		possiblePrefix[0] = true

		for tail := 1; tail <= n; tail++ {
			for head := 0; head < tail; head++ {
				if possiblePrefix[head] && towels[design[head:tail]] {
					possiblePrefix[tail] = true
					break
				}
			}
		}

		designMemo[design] = possiblePrefix[n]

		return possiblePrefix[n]
	}

	for _, design := range designs {
		if isPossible(design) {
			possibleDesigns++
		}
	}

	return possibleDesigns
}
