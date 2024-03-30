package stacks_test

import (
	"testing"

	customstack "github.com/Evergreenies/go-algorithms/algos/stacks/custom_stack"
	"github.com/stretchr/testify/assert"
)

func TestReverseStack(t *testing.T) {
	stack := customstack.NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	assert := assert.New(t)
	assert.Equal(stack.Size(), 3, "stack size must be 3")
	top, _ := stack.Peek()
	assert.Equal(top, 30, "before reverse peek element is 30")

	stack.ReverseStack()
	assert.Equal(stack.Size(), 3, "stack size must be 3 after reverse")
	top, _ = stack.Peek()
	assert.Equal(top, 10, "as stack reversed so top is 10 now")
}
