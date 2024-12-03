package main

import (
	"regexp"
	"strconv"
)

func partOne(input []byte) int {
	result := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(string(input), -1)

	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		result += a * b
	}

	return result
}
