package main

import (
	"strings"
)

func partOne(input []byte) int {
	matrix := strings.Split(string(input), "\n")
	height := len(matrix)
	width := len(matrix[0])
	visited := make(map[[2]int]bool)

	directions := [][2]int{
		{-1, 0}, // Up
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
	}

	var position [2]int

	// Find starting position
	for i, row := range matrix {
		for j, col := range row {
			if col == '^' {
				position = [2]int{i, j}
				visited[position] = true
			}
		}
	}

	direction := 0 // Start facing up

	for {
		nextRow := position[0] + directions[direction][0]
		nextCol := position[1] + directions[direction][1]

		// Check if the next position is within bounds
		if nextRow < 0 || nextRow >= height || nextCol < 0 || nextCol >= width {
			break
		}

		// Check if an obstacle is encountered
		if matrix[nextRow][nextCol] == '#' {
			// Change direction clockwise
			direction = (direction + 1) % 4
		} else {
			// Move to the next position
			position = [2]int{nextRow, nextCol}
			visited[position] = true
		}
	}

	return len(visited)
}
