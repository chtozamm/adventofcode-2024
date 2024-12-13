package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func partTwo(input []byte) int {
	tokens := 0
	machines := strings.Split(string(input), "\n\n")

	type Coordinates struct {
		X int
		Y int
	}

	type Machine struct {
		ButtonA Coordinates
		ButtonB Coordinates
		Prize   Coordinates
	}

	parseCoordinates := func(line string) Coordinates {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		x, _ := strconv.Atoi(matches[0])
		y, _ := strconv.Atoi(matches[1])
		return Coordinates{x, y}
	}

	for _, machine := range machines {
		components := strings.Split(machine, "\n")
		m := &Machine{
			ButtonA: parseCoordinates(components[0]),
			ButtonB: parseCoordinates(components[1]),
			Prize:   parseCoordinates(components[2]),
		}

		m.Prize.X += 10000000000000
		m.Prize.Y += 10000000000000

		/*
		   Assume that:
		   a = ButtonA, b = ButtonB, prize = Prize, s = countA, t = countB

		   System of linear equations:
		   aX * s + bX * t = prizeX
		   aY * s + bY * t = prizeY

		   Formulas to find counts:
		   s = (prizeX * bY) / (aX * bY - aY * bX)
		   t = (prize - aX * s) / bX
		*/

		countA := float64(m.Prize.X*m.ButtonB.Y-m.Prize.Y*m.ButtonB.X) / float64(m.ButtonA.X*m.ButtonB.Y-m.ButtonA.Y*m.ButtonB.X)
		countB := float64(float64(m.Prize.X)-float64(m.ButtonA.X)*countA) / float64(m.ButtonB.X)

		// Check if counts are integers
		if countA == math.Floor(countA) && countB == math.Floor(countB) {
			tokens += int(countA)*3 + int(countB)
		}
	}

	return tokens
}
