package stacks_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type optMinStack struct {
	elementStack []int
	minimumStack []int
}

func minOpt(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func (stack *optMinStack) push(data int) {
	lastTop := stack.peek()
	stack.elementStack = append(stack.elementStack, data)

	if len(stack.minimumStack) == 0 || lastTop >= data {
		stack.minimumStack = append(stack.minimumStack, data)
	}
}

func (stack *optMinStack) pop() int {
	if len(stack.elementStack) > 0 {
		popped := stack.elementStack[len(stack.elementStack)-1]
		stack.elementStack = stack.elementStack[:len(stack.elementStack)-1]

		if stack.peek() == popped {
			stack.minimumStack = stack.minimumStack[:len(stack.minimumStack)-1]
		}

		return popped
	}

	return math.MinInt32
}

func (stack *optMinStack) peek() int {
	if len(stack.elementStack) > 0 {
		return stack.elementStack[len(stack.elementStack)-1]
	}

	return 0
}

func (stack *optMinStack) size() int {
	return len(stack.elementStack)
}

func (stack *optMinStack) isEmpty() bool {
	return len(stack.elementStack) == 0
}

func (stack *optMinStack) clear() {
	stack.elementStack = nil
	stack.minimumStack = nil
}

func (stack *optMinStack) getMin() int {
	if len(stack.minimumStack) == 0 {
		return 0
	}

	return stack.minimumStack[len(stack.minimumStack)-1]
}

func TestOptMinStack(t *testing.T) {
	assert := assert.New(t)
	stack := optMinStack{}

	assert.True(stack.isEmpty(), "stack must be empty initially")
	assert.Equal(stack.size(), 0, "stack size must be zero initially")

	stack.push(1)
	stack.push(2)

	assert.Equal(stack.size(), 2, "stack size must be 2 now")
	assert.Equal(stack.peek(), 2, "stack peek must be 2 for now")
	assert.Equal(stack.pop(), 2, "last stack top element was 2")

	stack.push(3)
	assert.Equal(stack.getMin(), 1, "minimum is just 1 now")

	stack.clear()
	assert.True(stack.isEmpty(), "stack became empty on clear")
}
