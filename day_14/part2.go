package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/chtozamm/adventofcode-2024/utils"
)

func partTwo(input []byte, width, height int) int {
	seconds := -1
	lines := strings.Split(string(input), "\n")

	type Point struct {
		x, y int
	}

	type Robot struct {
		Position Point
		Velocity Point
	}

	robots := make([]Robot, len(lines))
	robotsMap := make(map[int]Point) // map robot's index to a location
	re := regexp.MustCompile(`-?\d+`)

	for i, line := range lines {
		components := utils.StringsToInts(re.FindAllString(line, -1))
		robots[i] = Robot{
			Position: Point{components[0], components[1]},
			Velocity: Point{components[2], components[3]},
		}
		robotsMap[i] = robots[i].Position
	}

	// Assume that for arranging a picture of a Christmas tree
	// no robots can share the same position
	for second := 1; second < width*height; second++ {
		seen := make(map[Point]bool) // map to track overlaps
		overlapDetected := false

		for i := range robots {
			velocity := robots[i].Velocity
			newPosition := Point{
				x: (robotsMap[i].x + velocity.x + width) % width,
				y: (robotsMap[i].y + velocity.y + height) % height,
			}
			robotsMap[i] = newPosition

			// Check for overlaps
			if seen[newPosition] {
				overlapDetected = true
			} else {
				seen[newPosition] = true
			}
		}

		// If there are no overlaps, record elapsed seconds and break out of the loop
		if !overlapDetected {
			seconds = second

			// Bonus: print state with a picture of a Christmas tree 🎄
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if seen[Point{x, y}] {
						fmt.Print("#")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
			break
		}
	}

	return seconds
}
