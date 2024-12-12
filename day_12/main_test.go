package main

import (
	"testing"
)

type testCase struct {
	input    []byte
	expected int
}

func TestPartOne(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		input := []byte(`AAAA
BBCD
BBCC
EEEC`)
		expected := 140
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("second", func(t *testing.T) {
		input := []byte(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`)
		expected := 772
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("third", func(t *testing.T) {
		input := []byte(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)
		expected := 1930
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		input := []byte(`AAAA
BBCD
BBCC
EEEC`)
		expected := 80
		if got := partTwo(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("second", func(t *testing.T) {
		input := []byte(`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`)
		expected := 236
		if got := partTwo(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("third", func(t *testing.T) {
		input := []byte(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`)
		expected := 368
		if got := partTwo(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("fourth", func(t *testing.T) {
		input := []byte(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`)
		expected := 1206
		if got := partTwo(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}
