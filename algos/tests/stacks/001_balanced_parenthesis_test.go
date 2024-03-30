package stacks_test

import (
	"testing"

	customstack "github.com/Evergreenies/go-algorithms/algos/stacks/custom_stack"
	"github.com/stretchr/testify/assert"
)

type pair struct {
	open  rune
	close rune
}

var pairs = map[rune]pair{
	'(': {'(', ')'},
	'{': {'{', '}'},
	'[': {'[', ']'},
	'<': {'<', '>'},
	')': {'(', ')'},
	'}': {'{', '}'},
	']': {'[', ']'},
	'>': {'<', '>'},
}

func isParenthesisBalanced(str string) bool {
	stack := customstack.NewStack()

	for _, symbol := range str {
		pr, ok := pairs[symbol]

		if symbol == pr.open {
			stack.Push(symbol)
		} else if ok {

			if stack.IsEmpty() {
				return false
			}

			top, _ := stack.Peek()
			if top == pr.open {
				if symbol != pr.close {
					return false
				}

				stack.Pop()
			}
		}
	}

	return stack.IsEmpty()
}

func TestBalancedParenthesis(t *testing.T) {
	assert := assert.New(t)

	assert.True(isParenthesisBalanced("()"), "() is balances parenthesis")
	assert.True(isParenthesisBalanced("() (() [()])"), "() (() [()]) is balances parenthesis")
	assert.True(isParenthesisBalanced("(a+b)+(c-d)"), "(a+b)+(c-d) this is valid expression")
	assert.True(isParenthesisBalanced("((a+b)+[c-d])"), "((a+b)+[c-d]) this must be a valid expression")
	assert.Equal(isParenthesisBalanced("((a+b)+[c-d]"), false, "((a+b)+[c-d] this is not a valid expression")
	assert.Equal(isParenthesisBalanced("((a+b)+[c-d})"), false, "((a+b)+[c-d}) this is not a valid expression")
	assert.Equal(isParenthesisBalanced("[}"), false, "[c-d} is invalid syntax")
}
