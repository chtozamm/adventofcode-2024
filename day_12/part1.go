package main

import "bytes"

func partOne(input []byte) int {
	totalPrice := 0
	matrix := bytes.Split(input, []byte("\n"))
	height := len(matrix)
	width := len(matrix[0])
	directions := [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	seen := make(map[[2]int]bool)

	var findPrice func(int, int) (int, int)
	findPrice = func(startX, startY int) (int, int) {
		if seen[[2]int{startX, startY}] {
			return 0, 0
		}
		seen[[2]int{startX, startY}] = true

		area := 1
		perimeter := 0

		for _, direction := range directions {
			newX := startX + direction[0]
			newY := startY + direction[1]

			if newX < 0 || newX >= width || newY < 0 || newY >= height {
				perimeter++
				continue
			}

			if matrix[newY][newX] != matrix[startY][startX] {
				perimeter++
				continue
			}

			addArea, addPerimeter := findPrice(newX, newY)
			area += addArea
			perimeter += addPerimeter
		}

		return area, perimeter
	}

	for y := range matrix {
		for x := range matrix[y] {
			if seen[[2]int{x, y}] {
				continue
			}

			area, perimeter := findPrice(x, y)
			totalPrice += area * perimeter
		}
	}

	return totalPrice
}
