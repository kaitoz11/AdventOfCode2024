package internal

import (
	"slices"
	"strconv"
	"strings"
)

func SolveDay05p1(input string) string {
	result := 0
	inp := strings.Split(input, "\n\n")
	rules := strings.Split(inp[0], "\n")
	// TODO:  value -> X|Y <- key
	ruleMapper := make(map[string][]int)
	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		ruleKey := ruleParts[1]
		ruleValues, _ := strconv.Atoi(ruleParts[0])
		ruleMapper[ruleKey] = append(ruleMapper[ruleKey], ruleValues)
	}

	updates := strings.Split(inp[1], "\n")
	for _, update := range updates {
		pages := strings.Split(update, ",")
		watchList := []int{}
		validatedPages := []int{}
		isCorrect := true
		for _, page := range pages {
			pageInt, _ := strconv.Atoi(page)

			// Check if page is in watchList
			_, found := slices.BinarySearch(watchList, pageInt)
			if found {
				isCorrect = false
				break
			}

			applyingRules := ruleMapper[page]
			if len(applyingRules) > 0 {
				watchList = append(watchList, applyingRules...)
				// I know this is overkill,
				// but I just wanna try out binary search from slices package
				slices.Sort(watchList)
			}

			validatedPages = append(validatedPages, pageInt)
		}
		if isCorrect {
			// fmt.Printf("%d: %v - %v\n", updateIndex, validatedPages, validatedPages[len(validatedPages)/2])
			result += validatedPages[len(validatedPages)/2]
		}

	}

	return strconv.Itoa(result)
}

func SolveDay05p2(input string) string {
	result := 0
	inp := strings.Split(input, "\n\n")
	rules := strings.Split(inp[0], "\n")

	ruleMapper := make(map[int][]int)
	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		ruleKey, _ := strconv.Atoi(ruleParts[1])

		ruleValues, _ := strconv.Atoi(ruleParts[0])
		ruleMapper[ruleKey] = append(ruleMapper[ruleKey], ruleValues)
	}

	// I was thinking about using graph to solve this...
	// but its 1am and I'm tired
	// at this point im hacking the solution

	updates := strings.Split(inp[1], "\n")
	for _, update := range updates {
		pages := strings.Split(update, ",")
		watchList := []int{}
		stupidList := []int{}
		watchingPositionTracker := make(map[int]int)
		validatedPages := []int{}

		for keyPage, page := range ruleMapper {
			watchList = append(watchList, keyPage)
			watchList = append(watchList, page...)
			// I know this is overkill,
			// but I just wanna try out binary search from slices package
			slices.Sort(watchList)
		}

		for pageIndex, page := range pages {
			pageInt, _ := strconv.Atoi(page)

			validatedPages = append(validatedPages, pageInt)
			_, found := slices.BinarySearch(watchList, pageInt)

			if found {
				// if has rules, add to watchList
				watchingPositionTracker[pageInt] = pageIndex
				stupidList = append(stupidList, pageInt)

			}

		}

		// fmt.Printf("%d: %v - %v\n", updateIndex, validatedPages, validatedPages[len(validatedPages)/2])
		// fmt.Printf("-: %v\n", watchingPositionTracker)
		swapped := false
		for i := 0; i < len(stupidList); i++ {
			for j := i + 1; j < len(stupidList); j++ {
				shouldSwap := slices.Contains(ruleMapper[stupidList[i]], stupidList[j])
				if shouldSwap {
					stupidList[i], stupidList[j] = stupidList[j], stupidList[i]
					// swap
					indexI := watchingPositionTracker[stupidList[i]]
					indexJ := watchingPositionTracker[stupidList[j]]
					watchingPositionTracker[stupidList[i]] = indexJ
					watchingPositionTracker[stupidList[j]] = indexI
					// fmt.Printf("- swapping %d and %d\n", indexI, indexJ)
					validatedPages[indexI], validatedPages[indexJ] = validatedPages[indexJ], validatedPages[indexI]
					swapped = true
				}
			}
			// fmt.Println(stupidList[i])
		}
		// fmt.Printf("sorted - %d: %v - %v\n", updateIndex, validatedPages, validatedPages[len(validatedPages)/2])
		if swapped {
			// fmt.Printf("->>%v\n", validatedPages[len(validatedPages)/2])
			result += validatedPages[len(validatedPages)/2]
		}

	}

	return strconv.Itoa(result)
}
