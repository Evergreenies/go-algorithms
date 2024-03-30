package stacks_test

import (
	"strconv"
	"testing"

	customstack "github.com/Evergreenies/go-algorithms/algos/stacks/custom_stack"
	"github.com/stretchr/testify/assert"
)

func isDigit(char uint8) bool {
	return char >= 48 && char <= 57
}

func postfixEvaluation(expression string) int {
	bs := 1
	ls := 1

	stack := customstack.NewStack()
	r := 0
	op := 0

	for index := 0; index < len(expression); index++ {
		if expression[index] == ' ' {
			continue
		}

		if expression[index] == '(' {
			stack.Push(bs)
			bs = bs * ls
			ls = 1
		} else if expression[index] == ')' {
			top, _ := stack.Pop()
			bs = top.(int)
			ls = 1
		} else if expression[index] == '-' || expression[index] == '+' {
			if expression[index] == '-' {
				ls = -1
			}
		} else {
			if isDigit(expression[index]) {
				end := index
				for index_j := index + 1; index_j < len(expression); index_j++ {
					if isDigit(expression[index_j]) {
						end = index_j
					} else {
						break
					}
				}

				op, _ = strconv.Atoi(expression[index : end+1])
				r += ls * bs * op
				index = end
				ls = 1
			}
		}
	}

	return r
}

func TestPostfixEvaluation(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(postfixEvaluation("1+1"), 2, "it must return 2 as result")
	assert.Equal(postfixEvaluation("2-1+2"), 3, "it must return 3 as a result")
}
