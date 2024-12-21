package internal_test

import (
	"aoc2024/internal"
	"testing"
)

func TestSolveDay02(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	expected := "2"
	result := internal.SolveDay02p1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
