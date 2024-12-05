/*
Implement a stack API using only a heap. A stack implements the following methods:

push(item), which adds an element to the stack
pop(), 			which removes and returns the most recently added element
						(or throws an error if there is nothing on the stack)
Recall that a heap has the following operations:

push(item), which adds a new key to the heap
pop(), 			which removes and returns the max value of the heap
* */

package stackbyheap

import (
	"container/heap"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// An Item is something we manage in priority queue
type Item struct {
	value    interface{}
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(value interface{}) {
	length := len(*pq)
	item := value.(*Item)
	item.index = length
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	length := len(*pq)
	item := old[length-1]
	old[length-1] = nil
	item.index = -1
	*pq = old[0 : length-1]

	return item
}

// Stack represents a stack imeplemented using a heap
type Stack struct {
	pq  PriorityQueue
	top int
}

func NewStack() *Stack {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	return &Stack{pq: pq}
}

func (s *Stack) Push(value interface{}) {
	heap.Push(&s.pq, &Item{
		value:    value,
		priority: s.top,
	})

	s.top++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.pq.Len() == 0 {
		return nil, errors.New("stack is empty")
	}

	item := heap.Pop(&s.pq).(*Item)
	s.top--

	return item.value, nil
}

func TestStackByHeap(t *testing.T) {
	assert := assert.New(t)

	stack := NewStack()

	assert.Zero(stack.pq.Len(), "stack must be empty initially")

	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	assert.Equal(3, stack.pq.Len(), "stack must be 3 items now")

	item, err := stack.Pop()
	assert.Nil(err, "error must be nil as stack has some items")
	assert.Equal("third", item, "the top elements should be `third` for now as it inserted last")
	assert.Equal(2, stack.pq.Len(), "stack must have 2 items now")
}
