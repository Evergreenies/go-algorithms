package stackbylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	top  *Node
	size int
}

type StackIntf interface {
	isFull() bool
	isEmpty() bool
	length() int
	push(data interface{})
	pop() (data interface{})
	peek() (data interface{})
}

func (stack *Stack) isFull() bool {
	return false // never full as it is list implementation
}

func (stack *Stack) isEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) length() int {
	return stack.size
}

func (stack *Stack) push(data interface{}) {
	stack.top = &Node{data, stack.top}
	stack.size++
}

func (stack *Stack) pop() (data interface{}) {
	if stack.length() <= 0 {
		return nil
	}

	data, stack.top = stack.top.data, stack.top.next
	stack.size--

	return
}

func (stack *Stack) peek() (data interface{}) {
	if stack.length() <= 0 {
		return nil
	}

	data = stack.top.data
	return
}

func TestStackByList(t *testing.T) {
	assert := assert.New(t)
	stack := new(Stack)

	assert.True(stack.isEmpty(), "initially stack must be empty")
	assert.False(stack.isFull(), "stack never full if it is implemented using linked list")
	assert.Equal(stack.length(), 0, "initially stack length must be zero")
	assert.Equal(stack.pop(), nil, "pop must return `nil` as stack is empty")
	assert.Nil(stack.peek(), "peek must return `nil` as stack is empty")

	stack.push(1)
	stack.push("Go")
	stack.push("Lang")

	assert.Equal(stack.length(), 3, "stack must have 3 elements till now")
	assert.Equal(stack.pop(), "Lang", "top was `Lang` in this case")
	assert.Equal(stack.peek(), "Go", "stack top must be `Go` int this case")
}
