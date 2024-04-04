package dailycodingproblems

type Stack struct {
	arr []int
	top int
}

type StackIntf interface {
	init() *Stack
	size() int
	isEmpty() bool

	push(data int)
	pop() int
	peek() int
}

func (stack *Stack) initStack() *Stack {
	return &Stack{
		arr: make([]int, 0),
		top: -1,
	}
}

func newStack() *Stack {
	return new(Stack).initStack()
}

func (stack *Stack) size() int {
	return stack.top + 1
}

func (stack *Stack) isEmpty() bool {
	return stack.top == -1
}

func (stack *Stack) push(data int) {
	stack.arr = append(stack.arr, data)
	stack.top++
}

func (stack *Stack) pop() int {
	top := stack.arr[stack.size()-1]
	stack.arr = stack.arr[:stack.size()-1]
	stack.top--

	return top
}

func (stack *Stack) peek() int {
	return stack.arr[stack.size()-1]
}
