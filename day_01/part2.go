package main

import (
	"strconv"
	"strings"
)

func partTwo(input []byte) int {
	lines := strings.Split(string(input), "\n")
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))
	similarityScore := 0

	overlapCountMap := make(map[int]int)

	for i, line := range lines {
		tmp := strings.Fields(line)
		left, _ := strconv.Atoi(tmp[0])
		right, _ := strconv.Atoi(tmp[1])
		leftList[i] = left
		rightList[i] = right
	}

	for _, val := range rightList {
		overlapCountMap[val]++
	}

	for i := 0; i < len(leftList); i++ {
		similarityScore += leftList[i] * overlapCountMap[leftList[i]]
	}

	return similarityScore
}
