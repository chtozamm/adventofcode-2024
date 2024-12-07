package main

import (
	"strconv"
	"strings"

	"github.com/chtozamm/adventofcode-2024/utils"
)

func partOne(input []byte) int {

	calibrationResult := 0
	equations := strings.Split(string(input), "\n")
	var isEquationPossible func(nums []int, target int) bool

	isEquationPossible = func(nums []int, target int) bool {
		if len(nums) == 1 {
			return nums[0] == target
		}

		// Check if multiplication is possible
		if target%nums[len(nums)-1] == 0 && isEquationPossible(nums[:len(nums)-1], target/nums[len(nums)-1]) {
			return true
		}

		// Check if addition is possible
		if target > nums[len(nums)-1] && isEquationPossible(nums[:len(nums)-1], target-nums[len(nums)-1]) {
			return true
		}

		return false
	}

	for _, equation := range equations {

		parts := strings.Split(equation, ":")
		target, _ := strconv.Atoi(parts[0])
		nums := utils.StringsToInts(strings.Fields(parts[1]))

		if isEquationPossible(nums, target) {
			calibrationResult += target
		}
	}

	return calibrationResult
}
