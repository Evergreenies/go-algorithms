/*
* Given a sorted list of integers of length N, determine if an element x is in the list without
* performing any multiplication, division, or bit-shift operations.
*
* Do this in O(log N) time
* */

package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type bs struct{}

func (bs) binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] == target {
			return true
		}

		if arr[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return false
}

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)
	b := bs{}

	assert.True(b.binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	assert.False(b.binarySearch([]int{1, 2, 3, 4, 6, 7, 8, 9}, 5))
}
