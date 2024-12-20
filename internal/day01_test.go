package internal_test

import (
	"aoc2024/internal"
	"testing"
)

func TestSolveDay01PartOne(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := "11"
	result := internal.SolveDay01PartOne(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay01PartTwo(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := "31"
	result := internal.SolveDay01PartTwo(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
