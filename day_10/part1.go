package main

import (
	"bytes"
)

func partOne(input []byte) int {
	matrix := bytes.Split(input, []byte{'\n'})
	directions := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	trailsMap := make(map[[4]int]bool)

	var findTrails func(int, int, int, int, int)
	findTrails = func(currX, currY, currVal, headX, headY int) {
		// Check if the current value is a peak
		if currVal == 9 {
			trailsMap[[4]int{currX, currY, headX, headY}] = true
			return
		}

		for _, direction := range directions {
			dx, dy := direction[0], direction[1]
			newX, newY := currX+dx, currY+dy

			// Skip direction if the new coordinates are out of bounds
			if newY >= len(matrix) || newY < 0 ||
				newX >= len(matrix[0]) || newX < 0 {
				continue
			}

			nextVal := int(matrix[newY][newX] - '0')

			if nextVal == currVal+1 {
				findTrails(newX, newY, nextVal, headX, headY)
			}
		}
	}

	for y := range matrix {
		for x := range matrix[y] {
			val := int(matrix[y][x] - '0')
			if val == 0 {
				findTrails(x, y, val, x, y)
			}
		}
	}

	return len(trailsMap)
}
