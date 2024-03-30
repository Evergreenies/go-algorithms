package stacks_test

import (
	"testing"

	stacks "github.com/Evergreenies/go-algorithms/algos/stacks/dynamic_stack"
	"github.com/stretchr/testify/assert"
)

func TestDynamicStack(t *testing.T) {
	assert := assert.New(t)
	stack := stacks.NewStack(5)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	err := stack.Push(6)
	if err != nil {
		assert.True(stack.IsFull(), "stack must be full now")
		stack.Resize()
		assert.Equal(int(stack.Capacity), 10, "stack size should be doubled now")
	}
}
