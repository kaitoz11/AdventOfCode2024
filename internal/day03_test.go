package internal_test

import (
	"aoc2024/internal"
	"testing"
)

func TestSolveDay03Part1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	expected := "161"
	result := internal.SolveDay03p1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay03Part2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	expected := "48"
	result := internal.SolveDay03p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}