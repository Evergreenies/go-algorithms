package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given a string of parentheses, write a function to compute the minimum number of parentheses to be removed
* to make the string valid (i.e. each open parenthesis is eventually closed).
*
* For example, given the string "()())()", you should return 1. Given the string ")(", you should return 2,
* since we must remove all of them.
* **/

func makeBalancedParenthesis(str string) int {
	stack := make([]int, 0)
	for index, character := range str {
		if string(character) == "(" {
			stack = append(stack, index)
		} else if string(character) == ")" {
			if len(stack) > 0 && string(str[stack[len(stack)-1]]) == "(" {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, index)
			}
		}
	}

	return len(stack)
}

func TestMakeBalancedParenthesis(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(makeBalancedParenthesis("()())()"), 1, "()())() it has only one parentheses to remove to make it balanced")
	assert.Equal(makeBalancedParenthesis(")("), 2, "we need to remove both")
}
