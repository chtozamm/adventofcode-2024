package main

import "testing"

func TestSolution(t *testing.T) {
	input := []byte(`5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`)

	t.Run("part 1", func(t *testing.T) {
		expected := 22
		if got := partOne(input, 7, 12); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		expected := "6,1"
		if got := partTwo(input, 7, 13); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}
