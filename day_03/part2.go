package main

import (
	"regexp"
	"strconv"
)

func partTwo(input []byte) int {
	result := 0
	instructions := []string{}
	enabled := true

	// Compile all necessary regex patterns
	instructionPattern := regexp.MustCompile(`(mul\(\d+,\d+\)|don't\(\)|do\(\))`)
	numPattern := regexp.MustCompile(`\d+`)

	// Find all instructions in one pass
	matches := instructionPattern.FindAllStringIndex(string(input), -1)
	for _, loc := range matches {
		instruction := string(input[loc[0]:loc[1]])
		instructions = append(instructions, instruction)
	}

	for _, instruction := range instructions {
		switch instruction {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				nums := numPattern.FindAllString(instruction, -1)
				a, _ := strconv.Atoi(nums[0])
				b, _ := strconv.Atoi(nums[1])
				result += a * b
			}
		}
	}

	return result
}
