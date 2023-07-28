package exercises

import "testing"

/**
* Exercise 2: Peak Index in a Mountain Array
Problem Statement: Given an array `arr` that is guaranteed to be a mountain array,
find and return the peak index of the array.

Example:
	Input: arr = [0, 1, 0]
	Output: 1
	Explanation: The peak element is 1, and its index is 1.
*/

func findPick(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		return 0
	}

	left, right := 0, len(arr)-1

	for left < right {
		mid := int((left + right) / 2)

		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
func TestFindPick(t *testing.T) {
	var tests = []struct {
		input          []int
		expected_index int
	}{
		{[]int{1, 3, 5, 6}, 3},
		{[]int{1, 3}, 1},
		{[]int{}, -1},
		{[]int{9}, 0},
		{[]int{6}, 0},
		{[]int{2, 4, 6, 8, 10}, 4},
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 8, 10, 5, 2}, 2},
	}

	for index, test := range tests {
		if actual_index := findPick(test.input); actual_index != test.expected_index {
			t.Errorf("FAILED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		} else {
			t.Logf("PASSED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		}
	}
}
