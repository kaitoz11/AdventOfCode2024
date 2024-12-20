package internal_test

import (
	"aoc2024/internal"
	"testing"
)

func TestSolveDay01(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected := "11"
	result := internal.SolveDay01(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
