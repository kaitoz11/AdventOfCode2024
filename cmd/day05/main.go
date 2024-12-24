package main

import (
	"aoc2024/internal"
	"aoc2024/pkg"
	"fmt"
	"os"
)

func main() {
	aoc := pkg.NewAocService(os.Getenv("AOC_SESSION"))
	input := aoc.GetInput(5)

	solution := internal.SolveDay05p1(input)
	fmt.Printf("Part 1 result: %s\n", solution) // 5991

	solution = internal.SolveDay05p2(input)
	fmt.Printf("Part 2 result: %s\n", solution) // 5479
}
