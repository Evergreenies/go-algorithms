package stacks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define stuct for stack
type Stack struct {
	top      int
	capacity uint
	arr      []interface{}
}

// defining interface
type StackIntf interface {
	Init(capacity uint) *Stack
	IsFull() bool
	IsEmpty() bool
	Size() uint
	Push(data interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Drain()
}

// Initialize a stack
func (stack *Stack) Init(capacity uint) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.arr = make([]interface{}, capacity)

	return stack
}

// create new empty stack
func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// stack is full when top is equal to the capacity
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

// stack is empty when top is equals to -1
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

// return size of the stack
func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

// push element to stack if it has space
func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}

	stack.top++
	stack.arr[stack.top] = data

	return nil
}

// pop an element from top of the stack if it is not empty
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	temp := stack.arr[stack.top]
	stack.arr[stack.top] = make([]interface{}, 1)[0]
	stack.top--

	return temp, nil
}

// return top element from arr
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return stack.arr[stack.top], nil
}

// drain removes all the element that are currently in stack
func (stack *Stack) Drain() {
	stack.arr = nil
	stack.top = -1
}

func TestStackArray(t *testing.T) {
	stack := NewStack(5)
	assert := assert.New(t)

	expected := make([]interface{}, 5)
	assert.Equal(
		stack.arr,
		expected,
		"stack must be empty at first",
	)
	assert.Equal(stack.IsEmpty(), true, "stack is empty")

	stack.Push(1)
	stack.Push(2)
	expected = make([]interface{}, 5)
	expected[0] = 1
	expected[1] = 2
	assert.Equal(stack.arr, expected, "stack must have 2 element only and rest of three are empty")

	stack.Pop()
	expected = make([]interface{}, 5)
	expected[0] = 1
	assert.Equal(stack.arr, expected, "as poped top stack must have one element only")

	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	err := stack.Push(5)
	assert.ErrorContains(err, "stack is full", "stack must be full now")
	assert.Equal(stack.IsFull(), true, "stack is full")
	assert.Equal(stack.IsEmpty(), false, "stack is not empty")
}
