package exercises

import (
	"testing"

	"golang.org/x/exp/slices"
)

/*
Problem Statement: Given an array of integers arr and an integer target,
return indices of the two numbers such that they add up to the target.
You may assume that each input would have exactly one solution, and you
may not use the same element twice.

Example:

	Input: arr = [2, 7, 11, 15], target = 9
	Output: [0, 1]
	Explanation: arr[0] + arr[1] equals 9, so the indices are 0 and 1.
*/
func twoSum(arr []int, target int) []int {
	var hash = map[int]int{}

	for index, element := range arr {
		diff := target - element

		if value, ok := hash[diff]; ok {
			return []int{value, index}
		}

		hash[element] = index
	}

	return []int{-1, -1}
}

func TestInterpolationSearch(t *testing.T) {

	var tests = []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{}, 9, []int{-1, -1}},
		{[]int{9}, 9, []int{-1, -1}},
		{[]int{6}, 9, []int{-1, -1}},
		{[]int{2, 4, 6, 8, 10}, 2, []int{-1, -1}},
		{[]int{2, 4, 6, 8, 10}, 6, []int{0, 1}},
		{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 10, []int{1, 2}},
		{[]int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}, 21, []int{2, 3}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, []int{3, 5}},
	}

	for index, test := range tests {
		if actual := twoSum(test.input, test.target); slices.Equal(actual[:], test.expected[:]) {
			t.Logf("PASSED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected, actual)
		} else {
			t.Errorf("FAILED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected, actual)
		}
	}
}
