package main

import (
	"regexp"
	"strconv"
	"strings"
)

func partOne(input []byte) string {
	re := regexp.MustCompile(`\d+`)
	values := re.FindAllString(string(input), -1)

	A, _ := strconv.Atoi(values[0])
	B, _ := strconv.Atoi(values[1])
	C, _ := strconv.Atoi(values[2])
	program := make([]int, len(values[3:]))
	for i, val := range values[3:] {
		program[i], _ = strconv.Atoi(val)
	}

	pointer := 0
	output := strings.Builder{}

	for pointer < len(program) {
		opcode := program[pointer]
		operand := program[pointer+1]
		combooperand := operand
		if operand == 4 {
			combooperand = A
		} else if operand == 5 {
			combooperand = B
		} else if operand == 6 {
			combooperand = C
		}

		switch opcode {
		case 0:
			A >>= combooperand
		case 1:
			B ^= operand
		case 2:
			B = combooperand % 8
		case 3:
			if A != 0 {
				pointer = operand - 2
			}
		case 4:
			B ^= C
		case 5:
			if output.Len() != 0 {
				output.WriteString(",")
			}
			output.WriteString(strconv.Itoa(combooperand % 8))
		case 6:
			B = A >> combooperand
		case 7:
			C = A >> combooperand
		}
		pointer += 2
	}

	return output.String()
}
