package dynamicstack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stack struct {
	top      int
	capacity uint
	arr      []interface{}
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
	stack.top = -1
	stack.arr = make([]interface{}, capacity)
	stack.capacity = capacity

	return stack
}

func NewStack(capacity uint) *Stack {
	return new(Stack).Init(capacity)
}

func (stack *Stack) Size() int {
	return stack.top + 1
}

func (stack *Stack) IsFull() bool {
	return stack.Size() == int(stack.capacity)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Push(element interface{}) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}

	stack.top++
	stack.arr[stack.top] = element

	return nil
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	top := stack.arr[stack.top]
	stack.arr[stack.top] = make([]interface{}, 0)[0]
	stack.top--

	if stack.Size() <= int(stack.capacity)/2 {
		stack.Resize()
	}

	return top, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return stack.arr[stack.top], nil
}

func (stack *Stack) Drain() {
	stack.top = -1
	stack.arr = make([]interface{}, stack.capacity)
}

func (stack *Stack) Resize() {
	if stack.IsFull() {
		stack.capacity *= 2
	} else {
		stack.capacity /= 2
	}

	target := make([]interface{}, stack.capacity)
	copy(target, stack.arr[:stack.top+1])
	stack.arr = target
}

func TestDynamicStack(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack(5)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	err := stack.Push(6)
	if err != nil {
		assert.True(stack.IsFull(), "stack must be full now")
		stack.Resize()
		assert.Equal(int(stack.capacity), 10, "stack size should be doubled now")
	}
}
