package dailycodingproblems

import (
	"slices"
	"testing"
)

/**
 * You are given an array of length N, where each element i represents the number
 * of ways we can produce i units of change. For example, [1, 0, 1, 1, 2]
 * would indicate that there is only one way to make 0, 2, or 3 units, and two
 * ways of making 4 units.
 * Given such an array, determine the denominations that must be in use. In the
 * case above, for example, there must be coins with value 2, 3, and 4.
 **/
func find_the_denominations(coins []int) []int {
	denominations := make([]int, 0)
	for index_i := 2; index_i < len(coins); index_i++ {
		if coins[index_i] > 0 {
			denominations = append(denominations, index_i)

			for index_j := 1; index_j < index_i; index_j++ {
				if coins[index_j] > 0 && coins[index_i-index_j] > 0 {
					if !slices.Contains(coins, index_j) {
						denominations = append(denominations, index_j)
					}

					if !slices.Contains(coins, index_i-index_j) {
						denominations = append(denominations, index_i-index_j)
					}
				}
			}
		}
	}

	return denominations
}

func TestFindDenominations(t *testing.T) {
	coins := []int{1, 0, 1, 1, 2}
	actualOutput := find_the_denominations(coins)
	if slices.Equal(actualOutput, []int{2, 3, 4}) {
		t.Logf("PASSED: \n\tActual: %v\n\tExpected: %v\n", actualOutput, []int{2, 3, 4})
	} else {
		t.Logf("FAILED: \n\tActual: %v\n\tExpected: %v\n", actualOutput, []int{2, 3, 4})
		t.Fail()
	}
}
