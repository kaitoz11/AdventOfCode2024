package internal

import (
	"fmt"
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

func SolveDay02p2(input string) string {
	safeCount := 0

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := strings.Fields(line)

		isSafe := true
		pairList := make([]int, len(nums)-1)
		increaseCount := 0
		decreaseCount := 0
		for i := 0; i < len(nums)-1; i++ {
			curNum, _ := strconv.Atoi(nums[i])
			nextNum, _ := strconv.Atoi(nums[i+1])

			delta := nextNum - curNum
			pairList[i] = delta

			if delta > 0 {
				increaseCount++
			} else {
				decreaseCount++
			}
		}

		isIncreasing := increaseCount > decreaseCount

		hasChance := true
		for i := 0; i < len(pairList); i++ {
			isSameTrend := pairList[i] > 0 == isIncreasing
			isInRange := func(num int) bool {
				return (num >= 1 && num <= 3) || (num <= -1 && num >= -3)

			}

			if !isInRange(pairList[i]) || !isSameTrend {
				if !hasChance {
					isSafe = false
					break
				}

				if i == 0 || i == len(pairList)-1 {
					hasChance = false
					i++
					continue
				}

				newPair := pairList[i] + pairList[i+1]
				if !isInRange(newPair) || (newPair > 0 != isIncreasing) {
					isSafe = false
					break
				}
				hasChance = false
				i++
			}
		}

		if isSafe {
			safeCount++
			fmt.Println(line)
		}

		fmt.Printf("---\nline: %s\n isSafe: %v\n%v\n", line, isSafe, pairList)
	}

	return strconv.Itoa(safeCount)
}
