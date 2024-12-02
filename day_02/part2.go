package main

import (
	"strings"

	"github.com/chtozamm/adventofcode-2024/utils"
)

func partTwo(input []byte) int {
	reports := strings.Split(string(input), "\n")
	safeReportsCount := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReportsCount++
		}
	}

	return safeReportsCount
}

// isReportSafe checks if removing any single level from the report results in a safe report.
func isReportSafe(report string) bool {
	levels := utils.StringsToInts(strings.Fields(report))

	for i := range levels {

		// Check if the modified slice (without the i-th element) is within bounds
		modified := make([]int, 0, len(levels)-1)
		modified = append(modified, levels[:i]...)
		modified = append(modified, levels[i+1:]...)

		if isWithinBounds(modified) {
			return true
		}
	}

	return false
}

// isWithingBounds returns true if the difference between adjacent levels
// falls into the specified bounds.
func isWithinBounds(levels []int) bool {

	// Note that using positive or negative ranges of bounds ensures
	// that all levels are either in ascending or descending order.
	for _, boundPair := range [][2]int{{0, 4}, {-4, 0}} {
		low, high := boundPair[0], boundPair[1]
		prev := levels[0]
		withinBounds := true

		for _, curr := range levels[1:] {
			diff := curr - prev
			if !(low < diff && diff < high) {
				withinBounds = false
				break
			}
			prev = curr
		}

		if withinBounds {
			return true
		}
	}

	return false
}
