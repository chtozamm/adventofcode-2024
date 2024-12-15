package main

import (
	"bytes"
	"slices"
)

func partTwo(input []byte) int {
	sum := 0
	parts := bytes.Split(input, []byte("\n\n"))

	expansion := map[byte][]byte{
		'#': []byte("##"),
		'O': []byte("[]"),
		'.': []byte(".."),
		'@': []byte("@."),
	}

	lines := bytes.Split(parts[0], []byte("\n"))
	warehouseMap := make([][]byte, len(lines))

	for i, line := range lines {
		expandedLevel := make([]byte, 0, len(line)*2)
		for _, char := range line {
			expandedLevel = append(expandedLevel, expansion[char]...)
		}
		warehouseMap[i] = expandedLevel
	}

	moves := bytes.ReplaceAll(parts[1], []byte("\n"), []byte{})

	type Point struct {
		x, y int
	}

	var robot Point

	// Find robot's starting position
	for y, row := range warehouseMap {
		for x := 0; x < len(row); x++ {
			if row[x] == '@' {
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
		boxes := []Point{{robot.x, robot.y}}
		shouldMove := true

		for i := 0; i < len(boxes); i++ {
			currX := boxes[i].x + dir.x
			currY := boxes[i].y + dir.y

			// Skip iteration if the box is already in the slice
			if slices.Index(boxes, Point{currX, currY}) != -1 {
				continue
			}

			char := warehouseMap[currY][currX]

			if char == '#' {
				shouldMove = false
				break
			} else if char == '[' {
				boxes = append(boxes, Point{currX, currY}, Point{currX + 1, currY})
			} else if char == ']' {
				boxes = append(boxes, Point{currX, currY}, Point{currX - 1, currY})
			}
		}

		if shouldMove {
			// Create a snapshot of the warehouse map
			warehouseCopy := make([][]byte, len(warehouseMap))
			for i := range warehouseCopy {
				warehouseCopy[i] = make([]byte, len(warehouseMap[i]))
				copy(warehouseCopy[i], warehouseMap[i])
			}

			// Vacate current positions of the boxes
			for _, box := range boxes[1:] {
				warehouseMap[box.y][box.x] = '.'
			}

			// Move boxes ahead
			for _, box := range boxes[1:] {
				warehouseMap[box.y+dir.y][box.x+dir.x] = warehouseCopy[box.y][box.x]
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
			if col == '[' {
				sum += 100*y + x
			}
		}
	}

	return sum
}
