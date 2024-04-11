package dailycodingproblems

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Implement three stacks using list
* **/

type MultiStack struct {
	list             [][]any
	number_of_stacks int
}

type MultiStackIntf interface {
	init(int) *MultiStack

	size(int) int
	isEmpty(int) bool
	isStackExceed(int) bool

	push(int, any)
	pop(int) (any, error)
	peek(int) (any, error)
}

func newMultiStack(number_of_stacks int) *MultiStack {
	return new(MultiStack).init(number_of_stacks)
}

func (stack *MultiStack) init(number_of_stacks int) *MultiStack {
	return &MultiStack{
		list:             make([][]any, number_of_stacks),
		number_of_stacks: number_of_stacks,
	}
}

func (stack *MultiStack) isStackExceed(number_of_stacks int) bool {
	return !(number_of_stacks > 0 && number_of_stacks <= stack.number_of_stacks)
}

func (stack *MultiStack) size(stack_number int) int {
	if stack.isStackExceed(stack_number) {
		panic(fmt.Sprintf("we have just %v stack, but entered %v\n", stack.number_of_stacks, stack_number))
	}

	return len(stack.list[stack_number-1])
}

func (stack *MultiStack) isEmpty(stack_number int) bool {
	if stack.isStackExceed(stack_number) {
		panic(fmt.Sprintf("we have just %v stack, but entered %v\n", stack.number_of_stacks, stack_number))
	}

	return len(stack.list[stack_number-1]) == 0
}

func (stack *MultiStack) push(stack_number int, item any) {
	if stack.isStackExceed(stack_number) {
		panic(fmt.Sprintf("we have just %v stack, but entered %v\n", stack.number_of_stacks, stack_number))
	}

	stack.list[stack_number-1] = append(stack.list[stack_number-1], item)
}

func (stack *MultiStack) pop(stack_number int) (any, error) {
	if stack.isStackExceed(stack_number) {
		panic(fmt.Sprintf("we have just %v stack, but entered %v\n", stack.number_of_stacks, stack_number))
	}

	if stack.isEmpty(stack_number) {
		return nil, errors.New("stack is empty")
	}

	top := stack.list[stack_number-1][len(stack.list[stack_number-1])-1]
	stack.list[stack_number-1] = stack.list[stack_number-1][:len(stack.list[stack_number-1])-1]

	return top, nil
}

func (stack *MultiStack) peek(stack_number int) (any, error) {
	if stack.isStackExceed(stack_number) {
		panic(fmt.Sprintf("we have just %v stack, but entered %v\n", stack.number_of_stacks, stack_number))
	}

	if stack.isEmpty(stack_number) {
		return nil, errors.New("stack is empty")
	}

	top := stack.list[stack_number-1][len(stack.list[stack_number-1])-1]
	return top, nil
}

func TestThreeStackInAList(t *testing.T) {
	assert := assert.New(t)
	stack := newMultiStack(3)

	assert.True(stack.isEmpty(1), "stack 1 must be empty at start")
	assert.Equal(stack.size(1), 0, "stack 1 size must be 0 at start")

	_, err := stack.pop(1)
	assert.NotNil(err, "error must not be nil as stack is empty")

	stack.push(1, 10)
	stack.push(1, 20)
	stack.push(2, 30)
	stack.push(3, 40)

	top, err := stack.peek(3)
	assert.Nil(err, "stack has elements sp it would not be nil at start")
	assert.Equal(top, 40, "stack 3 top is 40 for now")

	top, err = stack.pop(3)
	assert.Nil(err, "stack has elements sp it would not be nil at start")
	assert.Equal(top, 40, "stack 3 top is 40 for now")
}
