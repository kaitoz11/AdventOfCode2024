package internal

import (
	"fmt"
	"strings"
)

const (
	argSeparator  = ","
	bracketOpener = "("
	bracketCloser = ")"
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
	return 0, fmt.Errorf("Unknown operator: %s", o.Op)
}

func SolveDay03p1(input string) string {
	compilingInput := input
	OpStack := []operation{}
	// mulFunc := "mul("
	// callMul :=

	for {
		if compilingInput == "" {
			break
		}

		strings.Index(compilingInput, "mul(")
	}

	return ""
}
