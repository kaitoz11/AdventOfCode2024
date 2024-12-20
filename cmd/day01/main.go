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

	solution := internal.SolveDay01(input)
	fmt.Printf("result: %s", solution) // 1941353
}
