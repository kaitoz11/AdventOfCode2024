package main

import (
	"aoc2024/internal"
	"aoc2024/pkg"
	"fmt"
	"os"
)

func main() {
	aoc := pkg.NewAocService(os.Getenv("AOC_SESSION"))
	input := aoc.GetInput(4)

	solution := internal.SolveDay04p1(input)
	fmt.Printf("Part 1 result: %s\n", solution)

	// solution = internal.SolveDay04p2(input)
	// fmt.Printf("Part 2 result: %s\n", solution)
}
