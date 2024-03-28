package exercises

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Rotate an array of n elements to the right by k steps. For example,
 * with n= 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
* **/

func reverse_arr(arr []int, left int, right int) ([]int, error) {
	if arr == nil || len(arr) == 1 {
		return nil, errors.New("array looks empty or have single element so no need to reverse")
	}

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return arr, nil
}

func rotate_arr(arr []int, kth int) ([]int, error) {
	if arr == nil || len(arr) == 0 || kth < 0 {
		return nil, errors.New("illegal arguments")
	}

	if len(arr) == 1 {
		return arr, nil
	}

	if kth > len(arr) {
		kth = kth % len(arr)
	}

	pivot := len(arr) - kth
	arr, _ = reverse_arr(arr, 0, pivot-1)
	arr, _ = reverse_arr(arr, pivot, len(arr)-1)
	arr, _ = reverse_arr(arr, 0, len(arr)-1)

	return arr, nil
}

func TestRotateArray(t *testing.T) {
	assrt := assert.New(t)
	tests := []struct {
		input []int
		kth   int

		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{}, 3, nil},
		{[]int{1}, 3, []int{1}},
	}

	for index, test := range tests {
		actual, _ := rotate_arr(test.input, test.kth)

		assrt.Equal(test.expected, actual, fmt.Sprintf("Case: %d\n\tActual: %v\n\tExpected: %v\n", index+1, actual, test.expected))
	}
}
