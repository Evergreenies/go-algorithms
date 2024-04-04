package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given a stack of N elements, interleave the first half of the stack with
* the second half reversed using only one other queue. This should be done
* in-place.
* Recall that you can only push or pop from a stack, and enqueue or dequeue
* from a queue.
*
* For example, if the stack is [1, 2, 3, 4, 5], it should become [1, 5, 2, 4, 3].
* If the stack is [1, 2, 3, 4], it should become [1, 4, 2, 3].
*
* Hint: Try working backwards from the end state.
* **/

func interleavedStack(stack *Stack) *Stack {
	start, stop := 1, stack.size()
	queue := newQueue()

	for start != stop {
		for stack.size() > start {
			queue.append(stack.pop())
		}

		for queue.size > 0 {
			stack.push(queue.popLeft())
		}

		start++
	}
	fmt.Printf("ARR: %v\n", stack.arr)

	return stack
}

func TestInterleavedStack(t *testing.T) {
	assert := assert.New(t)

	stack := newStack()
	assert.True(stack.isEmpty(), "stack must be empty for now")

	stack.push(1)
	assert.Equal(stack.peek(), 1, "stack has only one element `1`")
	assert.Equal(stack.pop(), 1, "stack top was 1")

	queue := newQueue()
	assert.True(queue.isEmpty(), "queue must be empty at start")

	queue.append(1)
	assert.Equal(queue.size, 1, "queue size must be 1 for now")
	assert.Equal(queue.popLeft(), 1, "popLeft should return 1")

	stack = newStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	assert.Equal(interleavedStack(stack).arr, []int{1, 5, 2, 4, 3}, "it must return interleaved stack")
}
