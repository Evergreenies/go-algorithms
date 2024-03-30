package stacks_test

import (
	"testing"

	stackbylinkedlist "github.com/Evergreenies/go-algorithms/algos/stacks/stack_by_linked_list"
	"github.com/stretchr/testify/assert"
)

func TestStackByList(t *testing.T) {
	assert := assert.New(t)
	stack := new(stackbylinkedlist.Stack)

	assert.True(stack.IsEmpty(), "initially stack must be empty")
	assert.False(stack.IsFull(), "stack never full if it is implemented using linked list")
	assert.Equal(stack.Length(), 0, "initially stack length must be zero")
	assert.Equal(stack.Pop(), nil, "pop must return `nil` as stack is empty")
	assert.Nil(stack.Peek(), "peek must return `nil` as stack is empty")

	stack.Push(1)
	stack.Push("Go")
	stack.Push("Lang")

	assert.Equal(stack.Length(), 3, "stack must have 3 elements till now")
	assert.Equal(stack.Pop(), "Lang", "top was `Lang` in this case")
	assert.Equal(stack.Peek(), "Go", "stack top must be `Go` int this case")
}
