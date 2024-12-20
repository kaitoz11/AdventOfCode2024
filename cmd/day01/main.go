package main

import (
	"aoc2024/internal"
	"aoc2024/pkg"
	"fmt"
	"os"
)

func main() {
	aoc := pkg.NewAocService(os.Getenv("AOC_SESSION"))
	input := aoc.GetInput(1)

	solution := internal.SolveDay01PartOne(input)
	fmt.Printf("Part 1 result: %s\n", solution) // 1941353

	solution = internal.SolveDay01PartTwo(input)
	fmt.Printf("Part 2 result: %s\n", solution)
}
