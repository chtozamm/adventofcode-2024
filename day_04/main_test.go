package main

import "testing"

type testCase struct {
	input    []byte
	expected int
}

func TestPartOne(t *testing.T) {
	example := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	tc := testCase{
		input:    []byte(example),
		expected: 18,
	}

	if got := partOne(tc.input); got != tc.expected {
		t.Errorf("got %v, expected %v", got, tc.expected)
	}
}

func TestPartTwo(t *testing.T) {
	example := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	tc := testCase{
		input:    []byte(example),
		expected: 9,
	}

	if got := partTwo(tc.input); got != tc.expected {
		t.Errorf("got %v, expected %v", got, tc.expected)
	}
}
