package main

import (
	"testing"
)

type testCase struct {
	input    []byte
	expected string
}

func TestPartOne(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		input := []byte(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)
		expected := "4,6,3,5,6,3,5,2,1,0"
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}
