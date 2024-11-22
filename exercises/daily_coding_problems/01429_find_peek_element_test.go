/**
* Given a array that's sorted but rotated at some unknown pivot, in which all elements are distinct,
* find a "peak" element in O(log N) time.
*
* An element is considered a peak if it is greater than both its left and right neighbors. It is
* guaranteed that the first and last elements are lower than all others.
* */

package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type peekElementStrct struct{}

func (peekElementStrct) findPeekElement(arr []int) int {
	left, right := 0, len(arr)

	for left <= right {
		middle := int((left + right) / 2)

		if (arr[middle] > arr[middle-1] || middle == 0) && (middle == len(arr)-1 || arr[middle] > arr[middle+1]) {
			return arr[middle]
		}

		if arr[middle] < arr[middle+1] || middle < len(arr)-1 {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}

func TestFindPeekElement(t *testing.T) {
	p := &peekElementStrct{}
	ele := p.findPeekElement([]int{15, 20, 25, 5, 10})

	assert := assert.New(t)
	assert.Equal(25, ele, fmt.Sprintf("expected 25; got %v\n", ele))
}
