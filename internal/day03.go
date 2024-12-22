package internal

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	argSeparator  = ','
	bracketOpener = '('
	bracketCloser = ')'
)

type operator string

const (
	mul operator = "mul"
)

type operation struct {
	Op   operator
	Val1 int
	Val2 int
}

func (o operation) Do() (int, error) {
	switch o.Op {
	case mul:
		return o.Val1 * o.Val2, nil
	}
	return 0, fmt.Errorf("unknown operator: %s", o.Op)
}

func getNextNumInput(compilingInput string, endingCondition rune)(int,int,error){
	index := 0
	num := 0
	for i, c := range compilingInput{
		index = i
		if c < 48 || c > 57 {
			if c == endingCondition{
				break
			}else{
				return num,index, errors.New("invalid syntax for number argument")
			}
		}
		num = num*10+ int(c - '0')
	}
	return num, index, nil
}

func SolveDay03p1(input string) string {
	compilingInput := input
	result := 0
	for {
		if compilingInput == "" {
			break
		}
		op := string(mul) + string(bracketOpener)
		candidateIndex := strings.Index(compilingInput, op)
		if candidateIndex == -1 {
			break
		}
		compilingInput = compilingInput[candidateIndex+len(op):]
		num1, endIndex, err := getNextNumInput(compilingInput, argSeparator)
		if err != nil{
			// fmt.Println(err)
			continue
		}
		compilingInput = compilingInput[endIndex+1:]
		
		num2, endIndex, err := getNextNumInput(compilingInput, bracketCloser)
		if err != nil{
			// fmt.Println(err)
			continue
		}
		compilingInput = compilingInput[endIndex+1:]

		calculator := operation{
			Op: mul,
			Val1: num1,
			Val2: num2,
		}
		r, err:= calculator.Do()
		if err != nil {
			// fmt.Println(err)
			break
		}
		// fmt.Printf("---\ncompiling:\n%v\nnum1: %v num2: %v\n",compilingInput, num1, num2)
		
		result += r
	}

	return strconv.Itoa(result)
}
