package main

import "testing"

func TestSolution(t *testing.T) {
	input := []byte(`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`)

	t.Run("part 1", func(t *testing.T) {
		expected := 6
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 16
		if got := partTwo(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}
