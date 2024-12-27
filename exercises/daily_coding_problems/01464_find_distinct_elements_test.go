/*
* Given an array of integers in which two elements appear exactly once and all other elements appear
* exactly twice, find the two elements that appear only once.

For example, given the array [2, 4, 6, 8, 10, 2, 6, 10], return 4 and 8. The order does not matter.

Follow-up: Can you do this in linear time and constant space?
* */

package dailycodingproblems

import "testing"

type distinctNum struct{}

func (distinctNum) findTwoDistinctNums(arr []int) (int, int) {
	xorResult := 0
	for _, num := range arr {
		xorResult ^= num
	}

	setBit := xorResult & -xorResult

	num1, num2 := 0, 0
	for _, num := range arr {
		if num&setBit == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return num1, num2
}

func TestFindTwoDistinctNums(t *testing.T) {
	un := distinctNum{}
	num1, num2 := un.findTwoDistinctNums([]int{2, 4, 6, 8, 10, 2, 6, 10})
	t.Logf("The two unique numbers are: %d and %d\n", num1, num2)
}
