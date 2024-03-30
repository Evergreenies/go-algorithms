package stacks_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type minStack struct {
	elementStack []int
	minimumStack []int
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func newMinStack() minStack {
	return minStack{}
}

func (stack *minStack) push(data int) {
	stack.elementStack = append(stack.elementStack, data)

	if len(stack.minimumStack) == 0 {
		stack.minimumStack = append(stack.minimumStack, data)
	} else {
		minimum := min(stack.minimumStack[len(stack.minimumStack)-1], data)
		stack.minimumStack = append(stack.minimumStack, minimum)
	}
}

func (stack *minStack) pop() int {
	if len(stack.elementStack) > 0 {
		popped := stack.elementStack[len(stack.elementStack)-1]
		stack.elementStack = stack.elementStack[:len(stack.elementStack)-1]
		stack.minimumStack = stack.minimumStack[:len(stack.minimumStack)-1]

		return popped
	}

	return math.MaxInt32
}

func (stack *minStack) peek() int {
	if len(stack.elementStack) > 0 {
		return stack.elementStack[len(stack.elementStack)-1]
	}

	return 0
}

func (stack *minStack) size() int {
	return len(stack.elementStack)
}

func (stack *minStack) getMinimum() int {
	if len(stack.minimumStack) > 0 {
		return stack.minimumStack[len(stack.minimumStack)-1]
	}

	return 0
}

func (stack *minStack) isEmpty() bool {
	return len(stack.elementStack) == 0
}

func (stack *minStack) clear() {
	stack.elementStack = nil
	stack.minimumStack = nil
}

func TestMinStack(t *testing.T) {
	stack := newMinStack()

	stack.push(-1)
	stack.push(2)
	stack.push(0)
	stack.push(-4)

	assert := assert.New(t)

	assert.Equal(stack.size(), 4, "stack size must be 4 as there are 4 elements in stack")
	assert.Equal(stack.getMinimum(), -4, "minimum of stack must be -4 for now")
	assert.Equal(stack.peek(), -4, "last pushed element was -4")
	assert.Equal(stack.pop(), -4, "as last element was -4 sot pop must return -4")
	assert.Equal(stack.getMinimum(), -1, "now minimum must be -1")

	stack.clear()
	assert.True(stack.isEmpty(), "stack must be empty now")
}
