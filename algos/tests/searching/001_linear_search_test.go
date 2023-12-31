package searching_test

import (
	"testing"

	search "github.com/Evergreenies/go-algorithms/algos/searching"
)

func TestLinearSearch(t *testing.T) {

	arr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		arr[i] = i
	}

	var tests = []struct {
		input          []int
		target         int
		expected_index int
	}{
		{[]int{4, 2, 7, 1, 9, 5, 3}, 9, 4},
		{[]int{4, 2, 7, 1, 9, 5, 3}, 8, -1},
		{[]int{}, 9, -1},
		{[]int{9}, 9, 0},
		{[]int{6}, 9, -1},
		{[]int{2, 4, 6, 8, 2, 10}, 2, 0},
		{[]int{2, 4, 6, 8, 2, 10}, 9, -1},
		{arr, 999999, 999999},
		{[]int{-4, -2, -7, -1, -9, -5, -3}, -9, 4},
		{[]int{-4, -2, -7, -1, -9, -5, -3}, -8, -1},
	}

	for index, test := range tests {
		if actual_index := search.LinearSearch(test.input, test.target); actual_index != test.expected_index {
			t.Errorf("FAILED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		} else {
			t.Logf("PASSED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		}
	}
}
