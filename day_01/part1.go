package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func partOne(input []byte) int {
	lines := strings.Split(string(input), "\n")
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))
	distances := 0

	for i, line := range lines {
		tmp := strings.Fields(line)
		left, _ := strconv.Atoi(tmp[0])
		right, _ := strconv.Atoi(tmp[1])
		leftList[i] = left
		rightList[i] = right
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for i := 0; i < len(leftList); i++ {
		distances += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return distances
}
