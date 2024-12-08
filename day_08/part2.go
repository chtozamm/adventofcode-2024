package main

import (
	"strings"
)

func partTwo(input []byte) int {
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
					// Add discovered antinode to the map
					antinodesMap[[2]int{col, row}] = true

					// Calculate displacement between frequencies
					dx := col - x
					dy := row - y

					// Find antinodes ahead of the discovered one
					for n := 1; withinBounds(col+n*dx, row+n*dy); n++ {
						newX := col + n*dx
						newY := row + n*dy
						antinodesMap[[2]int{newX, newY}] = true
					}

					// Find antinodes behind of the discovered one
					for n := 1; withinBounds(col-n*dx, row-n*dy); n++ {
						newX := col - n*dx
						newY := row - n*dy
						antinodesMap[[2]int{newX, newY}] = true
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
