package stacks_test

import (
	"testing"

	stacks "github.com/Evergreenies/go-algorithms/algos/stacks/simple_stack"
	"github.com/stretchr/testify/assert"
)

func TestStackArray(t *testing.T) {
	stack := stacks.NewStack(5)
	assert := assert.New(t)

	expected := make([]interface{}, 5)
	assert.Equal(
		stack.Arr,
		expected,
		"stack must be empty at first",
	)
	assert.Equal(stack.IsEmpty(), true, "stack is empty")

	stack.Push(1)
	stack.Push(2)
	expected = make([]interface{}, 5)
	expected[0] = 1
	expected[1] = 2
	assert.Equal(stack.Arr, expected, "stack must have 2 element only and rest of three are empty")

	stack.Pop()
	expected = make([]interface{}, 5)
	expected[0] = 1
	assert.Equal(stack.Arr, expected, "as poped top stack must have one element only")

	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	err := stack.Push(5)
	assert.ErrorContains(err, "stack is full", "stack must be full now")
	assert.Equal(stack.IsFull(), true, "stack is full")
	assert.Equal(stack.IsEmpty(), false, "stack is not empty")
}
