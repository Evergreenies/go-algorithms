/*
* Given a number in the form of a list of digits, return all possible permutations.
*
* For example, given [1,2,3], return [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]].
* */

package dailycodingproblems

import "testing"

type permutationsStrct struct{}

func (permutationsStrct) permutations(arr []int) [][]int {
	var result [][]int
	var backtrack func(first int)

	backtrack = func(first int) {
		if first == len(arr) {
			perm := make([]int, len(arr))
			copy(perm, arr)
			result = append(result, perm)

			return
		}

		for index := first; index < len(arr); index++ {
			arr[first], arr[index] = arr[index], arr[first]
			backtrack(first + 1)
			arr[first], arr[index] = arr[index], arr[first]
		}
	}

	backtrack(0)

	return result
}

func TestPermutations(t *testing.T) {
	p := permutationsStrct{}
	t.Log(p.permutations([]int{1, 2, 3}))
}
