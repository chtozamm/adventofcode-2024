package main

import "bytes"

func partTwo(input []byte) int {
	totalPrice := 0
	lines := bytes.Split(input, []byte("\n"))
	height := len(lines)
	width := len(lines[0])

	seen := make(map[[2]int]bool)
	grid := make(map[[2]int]byte)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[[2]int{x, y}] = lines[y][x]
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	diagonals := [][2]int{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}}

	var findPrice func([2]int) (int, int)
	findPrice = func(start [2]int) (int, int) {
		seen[start] = true
		startValue := grid[start]

		area := 0
		sides := 0
		toVisit := [][2]int{start}

		for len(toVisit) > 0 {
			current := toVisit[0]
			toVisit = toVisit[1:]
			currX, currY := current[0], current[1]

			area++

			// Check for neighbors
			for _, direction := range directions {
				newX := currX + direction[0]
				newY := currY + direction[1]

				nextPoint := [2]int{newX, newY}

				if value, exists := grid[nextPoint]; exists && value == startValue {
					if !seen[nextPoint] {
						toVisit = append(toVisit, nextPoint)
						seen[nextPoint] = true
					}
				}
			}

			// Check for sides
			for _, diagonal := range diagonals {
				newX := currX + diagonal[0]
				newY := currY + diagonal[1]

				// Check for convex corners
				if grid[[2]int{newX, currY}] != startValue &&
					grid[[2]int{currX, newY}] != startValue {
					sides++
				}

				// Check for concave corners
				if grid[[2]int{newX, currY}] == startValue &&
					grid[[2]int{currX, newY}] == startValue &&
					grid[[2]int{newX, newY}] != startValue {
					sides++
				}
			}
		}

		return area, sides
	}

	for y := range lines {
		for x := range lines[y] {
			if seen[[2]int{x, y}] {
				continue
			}

			area, sides := findPrice([2]int{x, y})
			totalPrice += area * sides
		}
	}

	return totalPrice
}
