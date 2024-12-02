package main

import (
	"strings"

	"github.com/chtozamm/adventofcode-2024/utils"
)

func partOne(input []byte) int {
	reports := strings.Split(string(input), "\n")
	safeReportsCount := 0

outer:
	for _, report := range reports {
		levels := utils.StringsToInts(strings.Fields(report))

		ascending := levels[0] < levels[1]
		prev := levels[0]
		for _, curr := range levels[1:] {
			diff := utils.AbsoluteDifference(curr, prev)
			if ascending && curr < prev ||
				!ascending && curr > prev ||
				diff < 1 || diff > 3 {
				continue outer
			}
			prev = curr
		}
		safeReportsCount++
	}

	return safeReportsCount
}
