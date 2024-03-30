package customstack

import "errors"

type Stack struct {
	Top int
	Arr []interface{}
}

type StackIntf interface {
	Init() *Stack
	IsEmpty() bool
	Size() int
	Push(data interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}

func (stack *Stack) Init() *Stack {
	stack.Top = -1
	stack.Arr = make([]interface{}, 0)

	return stack
}

func NewStack() *Stack {
	return new(Stack).Init()
}

func (stack *Stack) Size() int {
	return len(stack.Arr)
}

func (stack *Stack) IsEmpty() bool {
	return stack.Top == -1
}

func (stack *Stack) Push(data interface{}) {
	// stack never become full as it is custom implementation
	stack.Arr = append(stack.Arr, data)
	stack.Top++
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	top := stack.Arr[stack.Size()-1]
	stack.Arr = stack.Arr[:stack.Size()-1]
	stack.Top--

	return top, nil
}

func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return stack.Arr[stack.Top], nil
}
