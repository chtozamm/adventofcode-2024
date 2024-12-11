package main

import (
	"testing"
)

type testCase struct {
	input    []byte
	expected []byte
}

func TestPartOne(t *testing.T) {
	input := []byte(`125 17`)
	expected := 55312

	if got := partOne(input); got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestPartTwo(t *testing.T) {
	input := []byte(`125 17`)
	blinks := 25
	expected := 55312

	if got := partTwo(input, blinks); got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}

}
