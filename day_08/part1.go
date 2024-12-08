package main

import (
	"strings"
)

func partOne(input []byte) int {
	// Map to store unique locations (x, y) of antinodes
	antinodesMap := make(map[[2]int]bool)

	matrix := strings.Split(string(input), "\n")

	height := len(matrix)
	width := len(matrix[0])

	withinBounds := func(x, y int) bool {
		return x >= 0 && x < width && y >= 0 && y < height
	}

	discoverAntinodes := func(x, y int, node rune) {
		for row := y; row < height; row++ {
			for col := 0; col < width; col++ {
				// Continue from the first item to the right of the node
				if row == y && col <= x {
					continue
				}

				if matrix[row][col] == byte(node) {
					// Calculate displacement between frequencies
					dx := col - x
					dy := row - y

					// Go forward from the discovered node's location
					if withinBounds(col+dx, row+dy) {
						antinodesMap[[2]int{col + dx, row + dy}] = true
					}

					// Go backwards from the *given* node's location
					if withinBounds(x-dx, y-dy) {
						antinodesMap[[2]int{x - dx, y - dy}] = true
					}
				}
			}
		}
	}

	// Find frequencies within the matrix and search for antinodes
	for y := range matrix {
		for x, char := range matrix[y] {
			if char != '.' {
				discoverAntinodes(x, y, char)
			}
		}
	}

	// Return the amount of antinodes
	return len(antinodesMap)
}
