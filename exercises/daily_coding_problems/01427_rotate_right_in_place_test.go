/**
* Given an array and a number k that's smaller than the length of the array, rotate the
* array to the right k elements in-place.
* */

package dailycodingproblems

import "testing"

type rotateRight struct{}

func (r rotateRight) reverse(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]

		start++
		end--
	}
}

func (r rotateRight) rotateRightInPlace(arr []int, k int) []int {
	length := len(arr)
	k %= length

	r.reverse(arr, 0, length-1)
	r.reverse(arr, 0, k-1)
	r.reverse(arr, k, length-1)

	return arr
}

func TestRotateRightInPlace(t *testing.T) {
	r := &rotateRight{}
	arr := r.rotateRightInPlace([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	t.Logf("%v\n", arr)
}
