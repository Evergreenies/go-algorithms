package stacks_test

import (
	"testing"

	customstack "github.com/Evergreenies/go-algorithms/algos/stacks/custom_stack"
	"github.com/stretchr/testify/assert"
)

func ispallindromeStack(str string) bool {
	stack := customstack.NewStack()
	index, length := 0, len(str)

	for index < length/2 {
		stack.Push(str[index])
		index++
	}

	if length%2 == 1 {
		index++
	}

	for index < len(str) {
		top, _ := stack.Peek()
		if stack.IsEmpty() || str[index] != top {
			return false
		}
		index++
		stack.Pop()
	}

	return true
}

func TestPallindromeStack(t *testing.T) {
	assert := assert.New(t)
	assert.True(ispallindromeStack("aba"), "this is the pallindrome string")
	assert.True(ispallindromeStack("aaaabbbbaaaa"), "this must be a valid pallindrome string")
	assert.False(ispallindromeStack("asasasas"), "this must not be pallindrome string")
}
