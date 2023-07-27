package exercises

import "testing"

/**
* Search Insert Position
Problem Statement: Given a sorted array of distinct integers and a target value,
return the index where the target is found. If the target is not found, return
the index where it would be if it were inserted in order.

Example:
	Input: 	arr = [1,3,5,6], target = 5
	Output: 2

Explanation: The target 5 is found at index 2.
*/

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := int((left + right) / 2)

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func searchIndexPosition(arr []int, target int) int {
	return binarySearch(arr, target)
}

func TestSearchIndexPosition(t *testing.T) {
	var tests = []struct {
		input          []int
		target         int
		expected_index int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3}, 2, 1},
		{[]int{1, 3}, 4, 2},
		{[]int{1, 3}, 0, 0},
		{[]int{}, 9, 0},
		{[]int{9}, 9, 0},
		{[]int{6}, 9, 1},
		{[]int{2, 4, 6, 8, 10}, 2, 0},
		{[]int{2, 4, 6, 8, 10}, 9, 4},
		{[]int{0, 0, 0, 1, 3, 5, 7, 9, 11, 13}, 7, 6},
		{[]int{0, 0, 0, 1, 3, 5, 7, 9, 11, 13}, 17, 10},
		{[]int{0, 0, 0, 1, 3, 5, 7, 9, 11, 13}, -1, 0},
	}

	for index, test := range tests {
		if actual_index := searchIndexPosition(test.input, test.target); actual_index != test.expected_index {
			t.Errorf("FAILED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		} else {
			t.Logf("PASSED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		}
	}
}
