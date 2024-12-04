package main

import "strings"

func partTwo(input []byte) int {
	xmasCount := 0
	lines := strings.Split(string(input), "\n")

	rows := len(lines)
	cols := len(lines[0])

	const (
		A = 'A'
		M = 'M'
		S = 'S'
	)

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if lines[row][col] == A {
				topLeft := lines[row-1][col-1]
				topRight := lines[row-1][col+1]
				bottomLeft := lines[row+1][col-1]
				bottomRight := lines[row+1][col+1]

				if (topLeft == M && topRight == S && bottomLeft == M && bottomRight == S) ||
					(topLeft == S && topRight == M && bottomLeft == S && bottomRight == M) ||
					(topLeft == S && topRight == S && bottomLeft == M && bottomRight == M) ||
					(topLeft == M && topRight == M && bottomLeft == S && bottomRight == S) {
					xmasCount++
				}
			}
		}
	}

	return xmasCount
}
