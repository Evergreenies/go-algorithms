package stacks_test

import (
	"testing"

	customstack "github.com/Evergreenies/go-algorithms/algos/stacks/custom_stack"
	"github.com/stretchr/testify/assert"
)

func TestCustomStack(t *testing.T) {
	assert := assert.New(t)
	stack := customstack.NewStack()

	assert.True(stack.IsEmpty(), "stack must be empty initially")
	assert.Equal(stack.Size(), 0, "stack size must be zero at start")
	assert.Equal(stack.Arr, make([]interface{}, 0), "stack arr must be empty at start")

	stack.Push(1)
	stack.Push(2)
	assert.Equal(stack.Size(), 2, "stack size must be 2 now")
	top, _ := stack.Peek()
	assert.Equal(top, 2, "peek of stack is 2 for now")

	top, _ = stack.Pop()
	assert.Equal(top, 2, "last popped emement is 2")
}
