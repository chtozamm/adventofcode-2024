package main

import (
	"strings"
)

func partTwo(input []byte) int {
	matrix := strings.Split(string(input), "\n")
	height := len(matrix)
	width := len(matrix[0])
	visited := make(map[[2]int]bool)
	loopsCount := 0

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
		nextPos := [2]int{nextRow, nextCol}

		// Check if the next position is within bounds
		if nextRow < 0 || nextRow >= height || nextCol < 0 || nextCol >= width {
			break
		}

		// Check if an obstacle is encountered
		if matrix[nextRow][nextCol] == '#' {
			// Change direction clockwise
			direction = (direction + 1) % 4
		} else {
			// Try to put an obstruction on the way
			if !visited[nextPos] && isStuck((direction+1)%4, position, directions, height, width, matrix, nextPos) {
				loopsCount++
			}

			// Move to the next position
			position = nextPos
			visited[position] = true
		}
	}

	return loopsCount
}

func isStuck(direction int, position [2]int, directions [][2]int, height, width int, matrix []string, obstruction [2]int) bool {
	visited := make(map[[3]int]bool) // Store coordinates and direction

	for {
		nextRow := position[0] + directions[direction][0]
		nextCol := position[1] + directions[direction][1]

		// Check if the next position is within bounds
		if nextRow < 0 || nextRow >= height || nextCol < 0 || nextCol >= width {
			return false
		}

		// Check if the position has been visited
		if visited[[3]int{nextRow, nextCol, direction}] {
			return true
		}

		// Check if either natural or added obstacle is encountered
		if matrix[nextRow][nextCol] == '#' || (nextRow == obstruction[0] && nextCol == obstruction[1]) {
			// Change direction clockwise
			direction = (direction + 1) % 4
		} else {
			// Move to the next position
			position = [2]int{nextRow, nextCol}
			visited[[3]int{nextRow, nextCol, direction}] = true
		}
	}
}
