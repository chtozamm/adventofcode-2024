package main

import "testing"

func TestSolution(t *testing.T) {
	input := []byte(`#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`)

	t.Run("part 1", func(t *testing.T) {
		expected := 3
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}
