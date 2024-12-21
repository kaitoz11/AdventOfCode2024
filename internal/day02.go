package internal

import (
	"strconv"
	"strings"
)

func SolveDay02p1(input string) string {
	safeCount := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := strings.Fields(line)
		isSafe := true
		isIncreasing := 1
		for i := 1; i < len(nums); i++ {
			curNum, _ := strconv.Atoi(nums[i])
			prevNum, _ := strconv.Atoi(nums[i-1])

			if i == 1 && curNum < prevNum {
				isIncreasing = -1
			}

			delta := (curNum - prevNum) * isIncreasing

			if delta < 1 || delta > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}
