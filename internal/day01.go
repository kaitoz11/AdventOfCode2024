package internal

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolveDay01PartTwo(input string) string {
	lines := strings.Split(input, "\n")
	leftList := make([]int, len(lines))
	rightCounterMap := make(map[int]int)

	for i, line := range lines {
		pair := strings.Split(line, "   ")
		if len(pair) == 2 {
			lVal, _ := strconv.Atoi(pair[0])
			rVal, _ := strconv.Atoi(pair[1])
			leftList[i] = lVal
			rightCounterMap[rVal]++
		}
	}

	result := 0

	for i := 0; i < len(leftList); i++ {
		result += leftList[i] * rightCounterMap[leftList[i]]
	}

	return strconv.Itoa(result)
}

func SolveDay01PartOne(input string) string {

	lines := strings.Split(input, "\n")
	leftList := make([]int, len(lines))
	rightList := make([]int, len(lines))
	for i, line := range lines {
		pair := strings.Split(line, "   ")
		if len(pair) == 2 {
			lVal, _ := strconv.Atoi(pair[0])
			rVal, _ := strconv.Atoi(pair[1])
			leftList[i] = lVal
			rightList[i] = rVal
		}
	}

	// fmt.Printf("len: %d %d\n", len(leftList), len(rightList))

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	result := 0
	for i := 0; i < len(lines); i++ {
		r := int(math.Abs(float64(leftList[i] - rightList[i])))
		// fmt.Printf("result: %d -> |%d - %d| = %d ->\n", result, leftList[i], rightList[i], r)
		result += r
	}

	return strconv.Itoa(result)
}
