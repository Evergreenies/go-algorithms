package dailycodingproblems

import (
	"fmt"
	"math"
	"testing"
)

/**
* mplement a stack that has the following methods:
*
* push(val), which pushes an element onto the stack
* pop(), which pops off and returns the topmost element of the stack.
* If there are no elements in the stack, then it should throw an error or return null.
* max(), which returns the maximum value in the stack currently.
* If there are no elements in the stack, then it should throw an error or return null.
*
* Each method should run in constant time.
* **/
type Stack struct {
	stack    []int
	top      int
	size     int
	maxValue int
}

func (stack *Stack) push(element int) {
	stack.stack = append(stack.stack, element)
	stack.top = element
	stack.size++

	if element > stack.maxValue {
		stack.maxValue = element
	}
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.size == 0 {
		return nil, fmt.Errorf("Cannot remove element from an empty stack.")
	}

	element := stack.stack[stack.size-1]
	stack.stack = stack.stack[:stack.size-1]
	stack.size--

	if stack.size == 0 {
		stack.maxValue = int(math.Inf(-1))
		stack.top = -1

		return stack.top, nil
	}

	stack.top = stack.stack[stack.size-1]
	if element >= stack.maxValue {
		maxValue := int(math.Inf(-1))
		for _, value := range stack.stack {
			if value > maxValue {
				maxValue = value
			}
		}
		stack.maxValue = maxValue
	}

	return stack.top, nil
}

func (stack *Stack) max() int {
	return stack.maxValue
}

func (stack *Stack) display(t *testing.T) {
	t.Logf("\n\nindex | stack\n")
	t.Logf("----- | -----\n")

	for index, element := range stack.stack {
		t.Logf("%v    | %v\n", index+1, element)
	}

	if stack.size != 0 {
		t.Logf("----- | -----\n")
	}

	t.Logf("SIZE: %v\n", stack.size)
	t.Logf("TOP:  %v\n\n", stack.top)
}

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		t.Logf("PASSED: \n\tActual: %v\n\tExpected: %v\n", actual, expected)
	} else {
		t.Logf("FAILED: \n\tActual: %v\n\tExpected: %v\n", actual, expected)
		t.Fail()
	}
}

func TestStack(t *testing.T) {
	stack := &Stack{
		stack:    make([]int, 0),
		top:      -1,
		maxValue: int(math.Inf(-1)),
	}

	stack.display(t)
	stack.push(1)
	stack.push(2)
	stack.push(3)
	t.Log(stack.maxValue)
	stack.display(t)

	element, err := stack.pop()
	if err != nil {
		t.Fatal(err)
	}
	assert(t, element, 2)
	stack.display(t)

	assert(t, stack.max(), 2)
}
