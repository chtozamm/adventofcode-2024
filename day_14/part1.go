package main

import (
	"regexp"
	"strings"

	"github.com/chtozamm/adventofcode-2024/utils"
)

func partOne(input []byte, width, height int) int {
	safetyFactor := 1
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

	// Simulate 100 seconds
	for second := 0; second < 100; second++ {
		for i := range robots {
			velocity := robots[i].Velocity
			newX := (robotsMap[i].x + velocity.x + width) % width
			newY := (robotsMap[i].y + velocity.y + height) % height
			robotsMap[i] = Point{newX, newY}
		}
	}

	robotsPerQuadrant := make([]int, 4)

	// Populate robotsPerQuadrant
	for _, point := range robotsMap {
		if point.x == width/2 || point.y == height/2 {
			continue
		}

		quadrant := 0
		if point.x > width/2 {
			quadrant += 1
		}
		if point.y > height/2 {
			quadrant += 2
		}
		robotsPerQuadrant[quadrant]++
	}

	for _, n := range robotsPerQuadrant {
		safetyFactor *= n
	}

	return safetyFactor
}
