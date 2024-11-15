/**
* Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
* For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
*
* Follow-up: Can you do this in O(N) time and constant space?
* */

package dailycodingproblems

import "testing"

type sumOfLargestNonAdjacentStrct struct{}

func (sumOfLargestNonAdjacentStrct) sumOfLargestNonAdjacent(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return max(0, arr[0])
	}

	include, exclude := max(0, arr[0]), 0
	for _, num := range arr[1:] {
		newExclude := max(include, exclude)
		include = exclude + num
		exclude = newExclude
	}

	return max(include, exclude)
}

func TestSumOfLargestNonAdjacent(t *testing.T) {
	sm := sumOfLargestNonAdjacentStrct{}
	t.Logf("expected 13; got %v\n", sm.sumOfLargestNonAdjacent([]int{2, 4, 6, 2, 5}))
	t.Logf("expected 10; got %v\n", sm.sumOfLargestNonAdjacent([]int{5, 1, 1, 5}))
}
