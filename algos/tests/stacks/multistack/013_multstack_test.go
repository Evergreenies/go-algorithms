package multistack

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MultiStack struct {
	top1, top2 int
	capacity   int
	array      []int
}

func (stack *MultiStack) init(capacity int) *MultiStack {
	stack.top1 = -1
	stack.top2 = capacity
	stack.capacity = capacity
	stack.array = make([]int, capacity)

	return stack
}

func newMultiStack(capacity int) *MultiStack {
	return new(MultiStack).init(capacity)
}

func (stack *MultiStack) size(stackNum int) int {
	if stackNum == 1 {
		return stack.top1 + 1
	}

	return stack.capacity - stack.top2
}

func (stack *MultiStack) isFull() bool {
	return (stack.size(1) + stack.size(1)) == stack.capacity
}

func (stack *MultiStack) isEmpty(stackNum int) bool {
	if stackNum == 1 {
		return stack.top1 == -1
	}

	return stack.top2 == stack.capacity
}

func (stack *MultiStack) push(stackNum int, data int) error {
	if stack.isFull() {
		return errors.New("stack is full")
	}

	if stackNum == 1 {
		stack.top1++
		stack.array[stack.top1] = data
	} else {
		stack.top2 = stack.top2 - 1
		stack.array[stack.top2] = data
	}

	return nil
}

func (stack *MultiStack) pop(stackNum int) int {
	var result int
	if stack.isEmpty(stackNum) {
		return math.MinInt32
	}

	if stackNum == 1 {
		result = stack.array[stack.top1]
		stack.top1--
	} else {
		result = stack.array[stack.top2]
		stack.top2++
	}

	return result
}

func (stack *MultiStack) peek(stackNum int) int {
	var result int
	if stack.isEmpty(stackNum) {
		return math.MinInt32
	}

	if stackNum == 1 {
		result = stack.array[stack.top1]
	} else {
		result = stack.array[stack.top2]
	}

	return result
}

func (stack *MultiStack) drain() {
	stack.array = nil
	stack.top1 = -1
	stack.top2 = stack.capacity
}

func TestMultiStack(t *testing.T) {
	stack := newMultiStack(5)
	assert := assert.New(t)

	assert.True(stack.isEmpty(1), "initially stack1 must be empty")
	assert.False(stack.isFull(), "initially stack cannot be full as it is already empty")

	stack.push(1, 10)
	stack.push(1, 20)
	stack.push(1, 30)
	stack.push(2, 40)

	assert.Equal(stack.size(1), 3, "stack size must be 3")
	assert.Equal(stack.size(2), 1, "stack 2 size must be 1 for now")

	assert.Equal(stack.peek(1), 30, "stack 1 top is 30")
}
