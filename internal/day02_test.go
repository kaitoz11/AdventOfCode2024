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

func TestSolveDay02p1(t *testing.T) {
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

func TestSolveDay02p2_1(t *testing.T) {
	input := "10 13 12 13 14\n10 11 12 13 14 1\n31 28 27 26 28 24\n58 57 58 56 54"
	expected := "4"
	result := internal.SolveDay02p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay02p2(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	expected := "4"
	result := internal.SolveDay02p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = `11 6 4 2 1`
	expected = "1"
	result = internal.SolveDay02p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	input = `1 2 1 3 5 11`
	expected = "0"
	result = internal.SolveDay02p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	input = `10 11 12 13 14 1`
	expected = "1"
	result = internal.SolveDay02p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	input = `1 10 12 13 14 15`
	expected = "1"
	result = internal.SolveDay02p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay02p2_2(t *testing.T) {
	input := `11 6 4 2 1`

	expected := "1"
	result := internal.SolveDay02p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
