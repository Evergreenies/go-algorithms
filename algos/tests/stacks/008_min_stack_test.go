package stacks_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MinStack struct {
	elementStack []int
	minimumStack []int
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(data int) {
	stack.elementStack = append(stack.elementStack, data)

	if len(stack.minimumStack) == 0 {
		stack.minimumStack = append(stack.minimumStack, data)
	} else {
		minimum := min(stack.minimumStack[len(stack.minimumStack)-1], data)
		stack.minimumStack = append(stack.minimumStack, minimum)
	}
}

func (stack *MinStack) Pop() int {
	if len(stack.elementStack) > 0 {
		popped := stack.elementStack[len(stack.elementStack)-1]
		stack.elementStack = stack.elementStack[:len(stack.elementStack)-1]
		stack.minimumStack = stack.minimumStack[:len(stack.minimumStack)-1]

		return popped
	}

	return math.MaxInt32
}

func (stack *MinStack) Peek() int {
	if len(stack.elementStack) > 0 {
		return stack.elementStack[len(stack.elementStack)-1]
	}

	return 0
}

func (stack *MinStack) Size() int {
	return len(stack.elementStack)
}

func (stack *MinStack) GetMinimum() int {
	if len(stack.minimumStack) > 0 {
		return stack.minimumStack[len(stack.minimumStack)-1]
	}

	return 0
}

func (stack *MinStack) IsEmpty() bool {
	return len(stack.elementStack) == 0
}

func (stack *MinStack) Clear() {
	stack.elementStack = nil
	stack.minimumStack = nil
}

func TestMinStack(t *testing.T) {
	stack := NewMinStack()

	stack.Push(-1)
	stack.Push(2)
	stack.Push(0)
	stack.Push(-4)

	assert := assert.New(t)

	assert.Equal(stack.Size(), 4, "stack size must be 4 as there are 4 elements in stack")
	assert.Equal(stack.GetMinimum(), -4, "minimum of stack must be -4 for now")
	assert.Equal(stack.Peek(), -4, "last pushed element was -4")
	assert.Equal(stack.Pop(), -4, "as last element was -4 sot pop must return -4")
	assert.Equal(stack.GetMinimum(), -1, "now minimum must be -1")

	stack.Clear()
	assert.True(stack.IsEmpty(), "stack must be empty now")
}
