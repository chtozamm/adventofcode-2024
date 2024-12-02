package main

import "testing"

type testCase struct {
	input    []byte
	expected int
}

func TestPartOne(t *testing.T) {
	example := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	tc := testCase{
		input:    []byte(example),
		expected: 2,
	}

	if got := partOne(tc.input); got != tc.expected {
		t.Errorf("got %v, expected %v", got, tc.expected)
	}
}

func TestPartTwo(t *testing.T) {
	example := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
11 6 4 3 1
3 6 4 3 1
1 3 6 7 9`

	tc := testCase{
		input:    []byte(example),
		expected: 6,
	}

	if got := partTwo(tc.input); got != tc.expected {
		t.Errorf("got %v, expected %v", got, tc.expected)
	}
}
