/*
* Given an arithmetic expression in Reverse Polish Notation, write a program to evaluate it.
* The expression is given as a list of numbers and operands. For example: [5, 3, '+']
* should return 5 + 3 = 8.
*
* For example, [15, 7, 1, 1, '+', '-', '/', 3, '*', 2, 1, 1, '+', '+', '-'] should return 5,
* since it is equivalent to ((15 / (7 - (1 + 1))) * 3) - (2 + (1 + 1)) = 5.
*
* You can assume the given expression is always valid.
* */

package dailycodingproblems

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type evalReversePolishNotationStrct struct{}

func (evalReversePolishNotationStrct) evalReversePolishNotation(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		stack_len := len(stack)
		switch token {
		case "+":
			b, a := stack[stack_len-1], stack[stack_len-2]
			stack = stack[:stack_len-2]
			stack = append(stack, a+b)
		case "-":
			b, a := stack[stack_len-1], stack[stack_len-2]
			stack = stack[:stack_len-2]
			stack = append(stack, a-b)
		case "*":
			b, a := stack[stack_len-1], stack[stack_len-2]
			stack = stack[:stack_len-2]
			stack = append(stack, a*b)
		case "/":
			b, a := stack[stack_len-1], stack[stack_len-2]
			stack = stack[:stack_len-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func TestEvalReversePolishNotation(t *testing.T) {
	assert := assert.New(t)
	ep := evalReversePolishNotationStrct{}

	assert.Equal(8, ep.evalReversePolishNotation([]string{"5", "3", "+"}))
	assert.Equal(5, ep.evalReversePolishNotation([]string{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"}))
}
