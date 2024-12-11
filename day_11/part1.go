package main

import (
	"strconv"
	"strings"

	"github.com/chtozamm/adventofcode-2024/utils"
)

func partOne(input []byte) int {
	stones := utils.StringsToInts(strings.Fields(string(input)))

	for i := 0; i < 25; i++ {
		res := make([]int, 0)
		for _, stone := range stones {
			stoneStr := strconv.Itoa(stone)

			switch {
			case stone == 0:
				res = append(res, 1)
			case len(stoneStr)%2 == 0:
				left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
				right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
				res = append(res, left, right)
			default:
				res = append(res, stone*2024)
			}
		}
		stones = res
	}

	return len(stones)
}
