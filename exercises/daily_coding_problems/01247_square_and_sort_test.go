package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given a sorted list of integers, square the elements and give the output in sorted order.
*
* For example, given [-9, -2, 0, 2, 3], return [0, 4, 4, 9, 81].
* **/

func squareAndSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	if len(arr) == 1 {
		return []int{arr[0] * arr[0]}
	}

	for index, item := range arr {
		arr[index] = item * item
	}

	swapped := true
	for swapped {
		swapped = false
		for index := 0; index < len(arr)-1; index++ {
			if arr[index] > arr[index+1] {
				arr[index], arr[index+1] = arr[index+1], arr[index]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}

func squareAndSortOptimized(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	if len(arr) == 1 {
		return []int{arr[0] * arr[0]}
	}

	for index, item := range arr {
		arr[index] = item * item
	}

	for index_i := 0; index_i < len(arr)-1; index_i++ {
		swapped := false
		for index_j := 0; index_j < (len(arr) - index_i - 1); index_j++ {
			if arr[index_j] > arr[index_j+1] {
				arr[index_j], arr[index_j+1] = arr[index_j+1], arr[index_j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}

func TestSquareAndSort(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(squareAndSort([]int{-9, -2, 0, 2, 3}), []int{0, 4, 4, 9, 81}, "this should return sorted result")
	assert.Equal(squareAndSortOptimized([]int{-9, -2, 0, 2, 3}), []int{0, 4, 4, 9, 81}, "this should return sorted result")
}
