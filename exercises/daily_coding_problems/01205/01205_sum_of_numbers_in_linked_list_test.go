package dailycodingproblems

import (
	"errors"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stack struct {
	top      int
	capacity uint
	arr      []interface{}
}

type StackIntf interface {
	Init(capacity uint) *Stack
	IsFull() bool
	IsEmpty() bool
	Push(element interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	Drain()
	Size() int
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

func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

func (stack *Stack) Size() int {
	return stack.top
}

func (stack *Stack) Push(element int) error {
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

	element := stack.arr[stack.top]
	stack.arr[stack.top] = make([]interface{}, 1)[0]
	stack.top--

	return element, nil
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

func prepareNumber(number int, digits int) int {
	for digits >= 0 {
		number *= 10
		digits--
	}

	return number
}

func getNumberFromLinkedList(arr []int) int {
	stack := NewStack(uint(len(arr)))
	for _, element := range arr {
		stack.Push(element)
	}

	_sum := 0
	for stack.Size() >= 0 {
		number, _ := stack.Pop()
		num, _ := number.(int)
		_sum += prepareNumber(num, stack.Size())
	}

	return _sum
}

func sumOfTwoLinkedLists(linkedList1 []int, linkedList2 []int) []int {
	_sum := getNumberFromLinkedList(linkedList1) + getNumberFromLinkedList(linkedList2)

	capacity := uint(math.Log10(float64(_sum)))
	stack := NewStack(capacity + 1)
	for _sum > 0 {
		divisor := int(math.Pow10(int(math.Log10(float64(_sum)))))
		digit := int(_sum / divisor)
		reminder := _sum % divisor
		stack.Push(digit)
		_sum = reminder
	}

	result := make([]int, 0)
	for stack.Size() >= 0 {
		popedElement, _ := stack.Pop()
		element, _ := popedElement.(int)
		result = append(result, element)
	}

	return result
}

func TestSumOfLinkedList(t *testing.T) {
	assert := assert.New(t)
	stack := NewStack(5)

	assert.True(stack.IsEmpty(), "stack must be empty at start")
	assert.False(stack.IsFull(), "stack must be empty at start")
	assert.Equal(stack.top, -1, "top must be -1 when stack is empty")

	stack.Push(1)
	stack.Push(2)
	assert.False(stack.IsEmpty(), "stack shold not be empty as it has two element in it")
	assert.False(stack.IsFull(), "stack is not full as it has only two elements in it and capacity is 5")

	top, _ := stack.Pop()
	assert.Equal(top, 2, "top element was 2")
	assert.False(stack.IsEmpty(), "stack still has 1 element in it")

	peek, _ := stack.Peek()
	assert.Equal(peek, 1, "peek must be 1 as it is top")
	assert.Equal(stack.top, 0, "stack top must be at 0 position")

	assert.Equal(getNumberFromLinkedList([]int{1, 2, 3, 4, 5}), 54321, "it must return actual number from provided list")
	assert.Equal(sumOfTwoLinkedLists([]int{9, 9}, []int{5, 2}), []int{4, 2, 1}, "it should return sum of two linked lists")
}
