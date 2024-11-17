/**
* You are given an array of nonnegative integers. Let's say you start at the beginning of
* the array and are trying to advance to the end. You can advance at most, the number of
* steps that you're currently on. Determine whether you can get to the end of the array.
*
* For example, given the array [1, 3, 1, 2, 0, 1], we can go from indices 0 -> 1 -> 3 -> 5,
* so return true.
*
* Given the array [1, 2, 1, 0, 0], we can't reach the end, so return false.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type canReachEndStrct struct{}

func (canReachEndStrct) canReachEnd(arr []int) bool {
	maxReach, target := 0, len(arr)-1

	for index, num := range arr {
		if index > maxReach {
			return false
		}

		maxReach = max(maxReach, index+num)
		if maxReach >= target {
			return true
		}
	}

	return false
}

func TestCanReachEndTest(t *testing.T) {
	assert := assert.New(t)

	cr := canReachEndStrct{}
	reached := cr.canReachEnd([]int{1, 3, 1, 2, 0, 1})
	assert.True(reached, fmt.Sprintf("expected true; got %v\n", reached))

	reached = cr.canReachEnd([]int{1, 2, 1, 0, 0})
	assert.False(reached, fmt.Sprintf("expected false; got %v\n", reached))
}
