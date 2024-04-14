package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Implement the function fib(n), which returns the nth number in the Fibonacci sequence, using only O(1) space.
* **/

func fib(num int) int {
	if num <= 1 {
		return num
	}

	old_sum, last_sum := 0, 1
	for index := 2; index < num+1; index++ {
		current_sum := old_sum + last_sum
		old_sum = last_sum
		last_sum = current_sum
	}

	return last_sum
}

func TestNthFib(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(fib(5), 5, "5th fib is must be 5")
}
