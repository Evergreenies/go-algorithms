/*
Given a number represented by a list of digits, find the next greater permutation of a number,
in terms of lexicographic ordering. If there is not greater permutation possible, return the
permutation with the lowest value/ordering.

For example, the list [1,2,3] should return [1,3,2]. The list [1,3,2] should return [2,1,3].
The list [3,2,1] should return [1,2,3].

Can you perform the operation without allocating extra memory (disregarding the input memory)?
*/

package dailycodingproblems

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type nextPermutationStrct struct{}

func (nextPermutationStrct) reverse(arr []int) {
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		arr[left], arr[right] = arr[right], arr[left]
	}
}

func (np nextPermutationStrct) nextPermutation(arr []int) {
	index_i := len(arr) - 2
	for index_i >= 0 && arr[index_i] >= arr[index_i+1] {
		index_i--
	}

	if index_i >= 0 {
		index_j := len(arr) - 1
		for index_j >= 0 && arr[index_j] <= arr[index_i] {
			index_j--
		}

		arr[index_i], arr[index_j] = arr[index_j], arr[index_i]
	}

	np.reverse(arr[index_i+1:])
}

func TestNextPermutation(t *testing.T) {
	np := &nextPermutationStrct{}
	assert := assert.New(t)

	arr, result := []int{1, 2, 3}, []int{1, 3, 2}
	np.nextPermutation(arr)
	assert.True(reflect.DeepEqual(
		arr,
		result),
		fmt.Sprintf("expected %v; got %v", result, arr),
	)

	arr, result = []int{1, 3, 2}, []int{2, 1, 3}
	np.nextPermutation(arr)
	assert.True(reflect.DeepEqual(
		arr,
		result),
		fmt.Sprintf("expected %v; got %v", result, arr),
	)

	arr, result = []int{3, 2, 1}, []int{1, 2, 3}
	np.nextPermutation(arr)
	assert.True(reflect.DeepEqual(
		arr,
		result),
		fmt.Sprintf("expected %v; got %v", result, arr),
	)
}
