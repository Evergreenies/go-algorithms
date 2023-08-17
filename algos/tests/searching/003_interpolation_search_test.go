package searching_test

import (
	"testing"

	search "github.com/Evergreenies/go-algorithms/algos/searching"
)

func TestInterpolationSearch(t *testing.T) {

	var tests = []struct {
		input          []int
		target         int
		expected_index int
	}{
		{[]int{}, 9, -1},
		{[]int{9}, 9, 0},
		{[]int{6}, 9, -1},
		{[]int{2, 4, 6, 8, 10}, 2, 0},
		{[]int{2, 4, 6, 8, 10}, 9, -1},
		{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, 10, 4},
		{[]int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}, 21, 6},
		{[]int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}, 27, -1},
		{[]int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, 64, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, -1},
	}

	for index, test := range tests {
		if actual_index := search.InterpolationSearch(test.input, test.target); actual_index != test.expected_index {
			t.Errorf("FAILED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		} else {
			t.Logf("PASSED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected_index, actual_index)
		}
	}
}
