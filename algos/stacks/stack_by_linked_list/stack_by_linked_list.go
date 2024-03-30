package stackbylinkedlist

type Node struct {
	Data interface{}
	Next *Node
}

type Stack struct {
	Top  *Node
	Size int
}

type StackIntf interface {
	IsFull() bool
	IsEmpty() bool
	Length() int
	Push(Data interface{})
	Pop() (Data interface{})
	Peek() (Data interface{})
}

func (stack *Stack) IsFull() bool {
	return false // never full as it is list implementation
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size == 0
}

func (stack *Stack) Length() int {
	return stack.Size
}

func (stack *Stack) Push(Data interface{}) {
	stack.Top = &Node{Data, stack.Top}
	stack.Size++
}

func (stack *Stack) Pop() (Data interface{}) {
	if stack.Length() <= 0 {
		return nil
	}

	Data, stack.Top = stack.Top.Data, stack.Top.Next
	stack.Size--

	return
}

func (stack *Stack) Peek() (data interface{}) {
	if stack.Length() <= 0 {
		return nil
	}

	data = stack.Top.Data
	return
}
