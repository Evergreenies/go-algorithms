package stacks

import (
	"errors"
)

// Define stuct for stack
type Stack struct {
	Top      int
	Capacity uint
	Arr      []interface{}
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
	stack.Top = -1
	stack.Capacity = capacity
	stack.Arr = make([]interface{}, capacity)

	return stack
}

// create new empty stack
func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

// stack is full when top is equal to the capacity
func (stack *Stack) IsFull() bool {
	return stack.Top == int(stack.Capacity)-1
}

// stack is empty when top is equals to -1
func (stack *Stack) IsEmpty() bool {
	return stack.Top == -1
}

// return size of the stack
func (stack *Stack) Size() uint {
	return uint(stack.Top + 1)
}

// push element to stack if it has space
func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}

	stack.Top++
	stack.Arr[stack.Top] = data

	return nil
}

// pop an element from top of the stack if it is not empty
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	temp := stack.Arr[stack.Top]
	stack.Arr[stack.Top] = make([]interface{}, 1)[0]
	stack.Top--

	return temp, nil
}

// return top element from arr
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return stack.Arr[stack.Top], nil
}

// drain removes all the element that are currently in stack
func (stack *Stack) Drain() {
	stack.Arr = nil
	stack.Top = -1
}
