package dailycodingproblems

import (
	"log"
	"testing"
)

/**
* Given n numbers, find the greatest common denominator between them.
* For example, given the numbers [42, 56, 14], return 14.
* **/

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func find_gcd(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		return arr[0]
	}

	a := arr[0]
	for _, b := range arr[1:] {
		a = gcd(a, b)
	}

	return a
}

func assert1204(t *testing.T, testNum int, actual interface{}, expected interface{}) {
	if actual == expected {
		log.Printf("PASSED: Case: %d\n\tActual: %v, \n\tExpected: %v\n", testNum, actual, expected)
	} else {
		t.Errorf("FAILED: Case: %d\n\tActual: %v,\n\tExpected: %v\n", testNum, actual, expected)
		t.Fail()
	}
}

func TestFindGCD(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{42, 56, 14}, 14},
		{[]int{}, -1},
		{[]int{1}, 1},
		{[]int{0, 28}, 28},
		{[]int{17, 23}, 1},
	}

	for index, test := range tests {
		assert1204(t, index+1, find_gcd(test.input), test.expected)
	}
}
