package dynamicstack

import (
	"errors"
)

type Stack struct {
	Top      int
	Capacity uint
	Arr      []interface{}
}

type StructImpl interface {
	Init(capacity uint) *Stack
	IsFull() bool
	IsEmpty() bool
	Size() int
	Drain()
	Push(element interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Resize()
}

func (stack *Stack) Init(capacity uint) *Stack {
	stack.Top = -1
	stack.Arr = make([]interface{}, capacity)
	stack.Capacity = capacity

	return stack
}

func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

func (stack *Stack) Size() int {
	return stack.Top + 1
}

func (stack *Stack) IsFull() bool {
	return stack.Size() == int(stack.Capacity)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Push(element interface{}) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}

	stack.Top++
	stack.Arr[stack.Top] = element

	return nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	Top := stack.Arr[stack.Top]
	stack.Arr[stack.Top] = make([]interface{}, 0)[0]
	stack.Top--

	if stack.Size() <= int(stack.Capacity)/2 {
		stack.Resize()
	}

	return Top, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return stack.Arr[stack.Top], nil
}

func (stack *Stack) Drain() {
	stack.Top = -1
	stack.Arr = make([]interface{}, stack.Capacity)
}

func (stack *Stack) Resize() {
	if stack.IsFull() {
		stack.Capacity *= 2
	} else {
		stack.Capacity /= 2
	}

	target := make([]interface{}, stack.Capacity)
	copy(target, stack.Arr[:stack.Top+1])
	stack.Arr = target
}
