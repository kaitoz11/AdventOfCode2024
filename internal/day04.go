package internal

import (
	"slices"
	"strconv"
	"strings"
)

func SolveDay04p1(input string) string {
	magicWord := "XMAS"
	inputMatrix := strings.Split(input, "\n")
	height := len(inputMatrix)
	width := len(inputMatrix[0])
	result := 0

	// horizontal, vertical, diagonal, reverse each
	// verticalList := []string{}
	// horizontalList := []string{}
	// diagonalUpList := []string{}
	// diagonalDownList := []string{}

	// horizontal
	for _, row := range inputMatrix {
		reversingRow := []byte(row)
		slices.Reverse(reversingRow)
		reversedRow := string(reversingRow)
		// horizontalList = append(horizontalList, row)
		// horizontalList = append(horizontalList, reversedRow)
		result += strings.Count(row, magicWord)
		result += strings.Count(reversedRow, magicWord)
	}

	// vertical
	for i := 0; i < width; i++ {
		vertical := ""

		for _, row := range inputMatrix {
			vertical += string(row[i])
		}

		r := []byte(vertical)
		slices.Reverse(r)
		reversed := string(r)
		// verticalList = append(verticalList, vertical)
		// verticalList = append(verticalList, string(r))

		result += strings.Count(vertical, magicWord)
		result += strings.Count(reversed, magicWord)
	}

	// diagonal
	iPointer := 0
	jPointer := 0
	c := 0
	for {
		if c == width+height {
			break
		}
		iIndex := 0
		if iPointer < height {
			iIndex = iPointer
		}
		jIndex := jPointer

		diagonalDown := ""
		diagonalUp := ""

		isBad := iIndex > height-4 || jIndex > width-4
		// badToDown := iIndex > height-4 || jIndex > width-4
		// badToUp := iIndex > height-4 || jIndex > width-4

		for i, j := iIndex, jIndex; j < width && i < height; j++ {
			// diagonalDownList = append(diagonalDownList, inputMatrix[iIndex][j])
			// diagonalUpList = append(diagonalUpList, inputMatrix[height-iIndex][j])

			if !isBad {
				diagonalUp += string(inputMatrix[height-1-i][j])
				diagonalDown += string(inputMatrix[i][j])
			}

			i++
		}
		// fmt.Printf("---\nup: %v\ni: %v j: %v\n", diagonalUp, iIndex, jIndex)
		// fmt.Println(diagnalUp)
		result += strings.Count(diagonalUp, magicWord)
		r := []byte(diagonalUp)
		slices.Reverse(r)
		reversed := string(r)
		result += strings.Count(reversed, magicWord)

		result += strings.Count(diagonalDown, magicWord)
		r = []byte(diagonalDown)
		slices.Reverse(r)
		reversed = string(r)
		result += strings.Count(reversed, magicWord)

		if iPointer < height {
			// go down
			iPointer++
			if iPointer == height {
				// go right
				jPointer++
			}
		} else {
			// go right
			jPointer++
		}
		c++
	}

	return strconv.Itoa(result)
}
