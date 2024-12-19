package main

import "strings"

func partTwo(input []byte) int {
	possibleDesigns := 0
	parts := strings.Split(string(input), "\n\n")
	towels := make(map[string]bool)
	for _, towel := range strings.Fields(strings.ReplaceAll(parts[0], ",", "")) {
		towels[towel] = true
	}
	designs := strings.Split(parts[1], "\n")

	designMemo := make(map[string]int)

	var countPossibilities func(string) int
	countPossibilities = func(design string) int {
		if count, exists := designMemo[design]; exists {
			return count
		}

		n := len(design)
		possibleWays := make([]int, n+1)
		possibleWays[0] = 1

		for tail := 1; tail <= n; tail++ {
			for head := 0; head < tail; head++ {
				if possibleWays[head] > 0 && towels[design[head:tail]] {
					possibleWays[tail] += possibleWays[head]
				}
			}
		}

		designMemo[design] = possibleWays[n]

		return possibleWays[n]
	}

	for _, design := range designs {
		possibleDesigns += countPossibilities(design)
	}

	return possibleDesigns
}
