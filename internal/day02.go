package internal

import (
	"fmt"
	"math"
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

		unsafeIndex := getUnsafeIndex(pairList)
		if unsafeIndex == -1 {
			// safe by dafualt
			safeCount++
			continue
		}

		cutHead := getUnsafeIndex(pairList[1:])
		safeToCut := unsafeIndex <= 1 && cutHead == -1
		cutTail := getUnsafeIndex(pairList[:len(pairList)-1])
		safeToCut = safeToCut || (unsafeIndex == len(pairList)-1 && cutTail == -1)
		if safeToCut {
			// safe when remove head or tail
			safeCount++
			continue
		}

		// check remove neighbor
		if unsafeIndex == 0 || getUnsafeIndex(removeNeighbor(pairList, unsafeIndex, true)) > -1 {
			if unsafeIndex == len(pairList)-1 || getUnsafeIndex(removeNeighbor(pairList, unsafeIndex, false)) > -1 {
				// not safe
				continue
			}
		}
		safeCount++

	}
	return strconv.Itoa(safeCount)
}

func removeNeighbor(pairList []int, index int, isLeft bool) []int {
	newDelta := pairList[index]
	if isLeft {
		newDelta += pairList[index-1]
	} else {
		newDelta += pairList[index+1]
	}
	newList := make([]int, 0, len(pairList)-1)
	for i := 0; i < len(pairList); i++ {
		if i == index {
			newList = append(newList, newDelta)
			continue
		}
		if isLeft && i == index-1 {
			continue
		} else if !isLeft && i == index+1 {
			continue
		}
		newList = append(newList, pairList[i])
	}
	return newList
}

// if safe return -1, else return the index of first unsafe pair
func getUnsafeIndex(pairList []int) int {
	firstAbsDelta := math.Abs(float64(pairList[0]))
	if firstAbsDelta > 3 || firstAbsDelta < 1 {
		return 0
	}

	for i := 1; i < len(pairList); i++ {
		absDelta := math.Abs(float64(pairList[i]))
		if pairList[i]*pairList[i-1] < 0 {
			return i
		}

		if absDelta > 3 || absDelta < 1 {
			return i
		}
	}
	return -1
}

func SolveDay02p2_imdonewiththis(input string) string {
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

		isInRange := func(num int) bool {
			return (num >= 1 && num <= 3) || (num <= -1 && num >= -3)
		}

		getUnsafeIndex := func(list []int) int {
			if !isInRange(list[0]) {
				return 0
			}
			for i := 1; i < len(list); i++ {
				if !isInRange(list[i]) {
					return i
				}

				if list[i]*list[i-1] < 0 {
					return i
				}
			}
			return -1
		}

		unsafeIndex := getUnsafeIndex(pairList)
		if unsafeIndex != -1 {
			cutHead := getUnsafeIndex(pairList[1:])
			safeToCut := unsafeIndex <= 1 && cutHead == -1
			cutTail := getUnsafeIndex(pairList[:len(pairList)-1])
			safeToCut = safeToCut || (unsafeIndex == len(pairList)-1 && cutTail == -1)
			if !safeToCut {
				isIncreasing := increaseCount > decreaseCount
				hasChance := true
				for i := 0; i < len(pairList); i++ {

					isTrend := pairList[i] > 0 == isIncreasing

					isSameSign := true
					if i == 0 {
						isSameSign = pairList[i]*pairList[i+1] > 0
					} else if i == len(pairList)-1 {
						isSameSign = pairList[i]*pairList[i-1] > 0
					}

					if !isSameSign || !isInRange(pairList[i]) || !isTrend {
						if !hasChance {
							isSafe = false
							break
						}
						isValidLeft := false
						isValidRight := false

						if i == 1 {
							newLeft := pairList[i] + pairList[i-1]
							isValidLeft = isInRange(newLeft) && newLeft > 0 == isIncreasing
						}
						if i != len(pairList)-1 {
							newRight := pairList[i] + pairList[i+1]
							isValidRight = isInRange(newRight) && newRight > 0 == isIncreasing
						}

						if isValidLeft {
							hasChance = false
							i++
							continue
						} else if isValidRight {

							hasChance = false
							i++
							continue
						} else if i == 0 && !isValidRight {
							continue
						}
						isSafe = false
						break
					}
				}
			} else {
				isSafe = true
			}
		}
		// isIncreasing := increaseCount > decreaseCount
		// hasChance := true
		// for i := 0; i < len(pairList); i++ {
		// 	isSameTrend := pairList[i] > 0 == isIncreasing

		// 	sameSign := true
		// 	if i == 0 {
		// 		sameSign = pairList[i]*pairList[i+1] > 0
		// 	}

		// 	if !isInRange(pairList[i]) || !isSameTrend || !sameSign {
		// 		if !hasChance {
		// 			isSafe = false
		// 			break
		// 		}

		// 		if i >= len(pairList)-1 {
		// 			break
		// 		}

		// 		newPair := pairList[i] + pairList[i+1]
		// 		if !isInRange(newPair) || (newPair > 0 != isIncreasing) {
		// 			if i == 0 {
		// 				if isInRange(pairList[1]) && !isInRange(pairList[0]) {
		// 					hasChance = false
		// 				}
		// 				continue
		// 			}
		// 			isSafe = false
		// 			break
		// 		}
		// 		hasChance = false
		// 		i++
		// 	}
		// }

		if isSafe {
			safeCount++
			// fmt.Println(line)
		} else {
			// fmt.Println(line)
		}

		fmt.Printf("---\nline: %s\n isSafe: %v\n%v\n", line, isSafe, pairList)
	}

	return strconv.Itoa(safeCount)
}
