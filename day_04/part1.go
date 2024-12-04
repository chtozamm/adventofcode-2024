package main

import "strings"

func partOne(input []byte) int {
	xmasCount := 0
	lines := strings.Split(string(input), "\n")

	rows := len(lines)
	cols := len(lines[0])

	const (
		X = 'X'
		M = 'M'
		A = 'A'
		S = 'S'
	)

	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	checkDirections := func(row, col, dr, dc int) bool {
		if row+3*dr < 0 || row+3*dr >= rows || col+3*dc < 0 || col+3*dc >= cols {
			return false
		}

		return lines[row][col] == X &&
			lines[row+dr][col+dc] == M &&
			lines[row+2*dr][col+2*dc] == A &&
			lines[row+3*dr][col+3*dc] == S
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				if checkDirections(row, col, dir[0], dir[1]) {
					xmasCount++
				}
			}
		}
	}

	return xmasCount
}
