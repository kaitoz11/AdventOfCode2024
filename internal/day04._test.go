package internal_test

import (
	"aoc2024/internal"
	"testing"
)

func TestSolveDay04Part1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	expected := "18"
	result := internal.SolveDay04p1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay04Part2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	expected := "9"
	result := internal.SolveDay04p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay04Part2_1(t *testing.T) {
	input := `XXXX
XMXM
XXAX
XSXS`

	expected := "1"
	result := internal.SolveDay04p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
