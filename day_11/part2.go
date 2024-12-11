package main

import (
	"strconv"
	"strings"

	"github.com/chtozamm/adventofcode-2024/utils"
)

func partTwo(input []byte, blinks int) int {
	stones := utils.StringsToInts(strings.Fields(string(input)))
	res := 0

	cache := make(map[[2]int]int)

	var count func(int, int) int
	count = func(stone, blinksLeft int) int {
		if blinksLeft == 0 {
			return 1
		}

		if val, ok := cache[[2]int{stone, blinksLeft}]; ok {
			return val
		}

		stoneStr := strconv.Itoa(stone)

		switch {
		case stone == 0:
			val := count(1, blinksLeft-1)
			cache[[2]int{stone, blinksLeft}] = val
			return val
		case len(stoneStr)%2 == 0:
			left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
			right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
			val := count(left, blinksLeft-1) + count(right, blinksLeft-1)
			cache[[2]int{stone, blinksLeft}] = val
			return val
		default:
			val := count(stone*2024, blinksLeft-1)
			cache[[2]int{stone, blinksLeft}] = val
			return val
		}
	}

	for _, stone := range stones {
		res += count(stone, blinks)
	}

	return res
}
