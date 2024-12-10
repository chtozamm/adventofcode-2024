package main

import "testing"

type testCase struct {
	input    []byte
	expected int
}

func TestPartOne(t *testing.T) {

	t.Run("small input", func(t *testing.T) {
		input := []byte(`0123
1234
8765
9876`)
		expected := 1

		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("bigger input", func(t *testing.T) {
		input := []byte(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)
		expected := 36

		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}

func TestPartTwo(t *testing.T) {
	input := []byte(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)
	expected := 81

	if got := partTwo(input); got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}

}
