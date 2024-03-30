package stacks_test

import (
	"testing"

	stacks "github.com/Evergreenies/go-algorithms/algos/stacks/simple_stack"
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
	stack := stacks.NewStack(1)

	for _, symbol := range str {
		if pr, ok := pairs[symbol]; ok {
			if symbol == pr.open {
				stack.Push(symbol)
			} else {
				if stack.IsEmpty() {
					return false
				}

				top, _ := stack.Pop()
				if top == pr.open {
					if symbol != pr.close {
						return false
					}
				}
			}
		}
	}

	return stack.IsEmpty()
}

func TestBalancedParenthesis(t *testing.T) {
	assert := assert.New(t)

	assert.True(isParenthesisBalanced("()"), "() is balances parenthesis")
	// assert.True(isParenthesisBalanced("() (() [()])"), "() (() [()]) is balances parenthesis")
	// assert.True(isParenthesisBalanced("(a+b)+(c-d)"), "(a+b)+(c-d) this is valid expression")
	// assert.True(isParenthesisBalanced("((a+b)+[c-d])"), "((a+b)+[c-d]) this must be a valid expression")
	// assert.Equal(isParenthesisBalanced("((a+b)+[c-d]"), false, "((a+b)+[c-d] this is not a valid expression")
	// assert.Equal(isParenthesisBalanced("((a+b)+[c-d})"), false, "((a+b)+[c-d}) this is not a valid expression")
}
