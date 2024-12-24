package internal_test

import (
	"aoc2024/internal"
	"testing"
)

func TestSolveDay06Part1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected := "41"
	result := internal.SolveDay06p1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay06Part1_1(t *testing.T) {
	input := `...>......`

	expected := "7"
	result := internal.SolveDay06p1(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay06Part2(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expected := "6"
	result := internal.SolveDay06p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay06Part2_1(t *testing.T) {
	input := `.#..#.....
........#.
....^.....
.......#..`

	expected := "2"
	result := internal.SolveDay06p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay06Part2_2(t *testing.T) {
	input := `.......v..
.#........
........#.
#.........
..........`

	expected := "1"
	result := internal.SolveDay06p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay06Part2_3(t *testing.T) {
	input := `..v.......
..........
.#.#......
..#.......`
	expected := "1"
	result := internal.SolveDay06p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestSolveDay06Part2_4(t *testing.T) {
	input := `..........
.#...#....
....#.....
..^.......`
	expected := "1"
	result := internal.SolveDay06p2(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
