package main

import (
	"testing"
)

type testCase struct {
	input    []byte
	expected int
}

func TestPartOne(t *testing.T) {
	t.Run("my test", func(t *testing.T) {
		input := []byte(`###############
#.......#.....#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S.E#.....#...#
###############`)
		expected := 2
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("first", func(t *testing.T) {
		input := []byte(`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`)

		expected := 7036
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
	t.Run("second", func(t *testing.T) {
		input := []byte(`#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`)

		expected := 11048
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}
