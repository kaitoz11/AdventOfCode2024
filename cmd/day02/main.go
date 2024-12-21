package main

import (
	"aoc2024/internal"
	"aoc2024/pkg"
	"fmt"
	"os"
)

func main() {
	aoc := pkg.NewAocService(os.Getenv("AOC_SESSION"))
	input := aoc.GetInput(2)

	solution := internal.SolveDay02p1(input)
	fmt.Printf("Part 1 result: %s", solution) // 236
}
