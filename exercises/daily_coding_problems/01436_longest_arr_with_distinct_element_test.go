/**
* Given an array of elements, return the length of the longest subarray where all its
* elements are distinct.
*
* For example, given the array [5, 1, 3, 5, 2, 3, 4, 1], return 5 as the longest subarray
* of distinct elements is [5, 2, 3, 4, 1].
* */
package dailycodingproblems

import (
	"fmt"
	"testing"
)

type longestDistinctSubarrayStrct struct{}

// longestDistinctSubarray
// returns the length of the longest subarray with distinct items
func (longestDistinctSubarrayStrct) longestDistinctSubarray(arr []int) int {
	elementIndex := make(map[int]int)
	start, maxLength := 0, 0

	for end := 0; end < len(arr); end++ {
		// check if element already in the map
		// and its index is within current window
		if idx, ok := elementIndex[arr[end]]; ok && idx >= start {
			start = idx + 1
		}

		// update the elmenets index in the map
		elementIndex[arr[end]] = end

		// calculate length of current window and update maxLength
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func TestLongestDistinctSubarray(t *testing.T) {
	ldsa := longestDistinctSubarrayStrct{}
	fmt.Println(
		"Length of longest distinct subarray: ",
		ldsa.longestDistinctSubarray([]int{5, 3, 1, 5, 2, 3, 4, 1}),
	)
}
