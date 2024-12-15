package main

import (
	"bytes"
)

func partOne(input []byte) int {
	sum := 0
	parts := bytes.Split(input, []byte("\n\n"))

	warehouseMap := bytes.Split(parts[0], []byte("\n"))
	moves := bytes.ReplaceAll(parts[1], []byte("\n"), []byte{})

	type Point struct {
		x, y int
	}

	var robot Point

	// Find robot's starting position
	for y := range warehouseMap {
		for x := range warehouseMap[y] {
			if warehouseMap[y][x] == '@' {
				robot = Point{x, y}
				break
			}
		}
	}

	directions := map[byte]Point{
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
	}

	for _, move := range moves {
		dir := directions[move]
		currX, currY := robot.x, robot.y
		boxes := []Point{}
		shouldMove := true

		// Move in the specified direction until hitting a wall or an empty space
		for {
			currX += dir.x
			currY += dir.y

			char := warehouseMap[currY][currX]

			if char == '#' {
				shouldMove = false
				break
			} else if char == 'O' {
				boxes = append(boxes, Point{currX, currY})
			} else {
				break
			}
		}

		if shouldMove {
			// Move boxes ahead
			for _, box := range boxes {
				warehouseMap[box.y+dir.y][box.x+dir.x] = 'O'
			}

			// Move robot ahead
			warehouseMap[robot.y][robot.x] = '.'
			warehouseMap[robot.y+dir.y][robot.x+dir.x] = '@'
			robot.x += dir.x
			robot.y += dir.y
		}
	}

	// Calculate sum of all boxes' GPS coordinates
	for y, row := range warehouseMap {
		for x, col := range row {
			if col == 'O' {
				sum += 100*y + x
			}
		}
	}

	return sum
}
