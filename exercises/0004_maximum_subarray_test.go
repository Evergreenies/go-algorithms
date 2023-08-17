package exercises

import (
	"math"
	"testing"
)

func maxSumSubArray(arr []float64) float64 {
	if len(arr) <= 0 {
		return -1
	}

	var max_sum, current_sum float64 = arr[0], 0

	for _, element := range arr {
		current_sum = math.Max(float64(element), current_sum+float64(element))
		max_sum = math.Max(max_sum, current_sum)
	}

	return max_sum
}

func TestMaxSumSubArray(t *testing.T) {

	var tests = []struct {
		input    []float64
		target   float64
		expected float64
	}{
		{[]float64{}, 9, -1},
		{[]float64{9}, 9, 9},
		{[]float64{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, 6},
		{[]float64{-2, -3, -1, -5}, -1, -1},
	}

	for index, test := range tests {
		if actual := maxSumSubArray(test.input); actual == test.expected {
			t.Logf("PASSED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected, actual)
		} else {
			t.Errorf("FAILED: Test Case: %v \n Expected: %v \n Actual: %v \n", index, test.expected, actual)
		}
	}
}
