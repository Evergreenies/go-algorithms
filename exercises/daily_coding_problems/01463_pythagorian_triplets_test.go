/*
* Given an array of integers, determine whether it contains a Pythagorean triplet. Recall that
* a Pythogorean triplet (a, b, c) is defined by the equation a2+ b2= c2.
* */
package dailycodingproblems

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type pythagorianTripletsStrct struct{}

func (pythagorianTripletsStrct) pythagorianTriplets(arr []int) bool {
	length := len(arr)

	for index := 0; index < length; index++ {
		arr[index] = arr[index] * arr[index]
	}

	sort.Ints(arr)

	for index := length - 1; index >= 2; index-- {
		left, right := 0, length-1

		for left < right {
			_sum := arr[left] + arr[right]
			if _sum == arr[index] {
				return true
			} else if _sum < arr[index] {
				left++
			} else {
				right--
			}
		}
	}

	return false
}

func TestPythagorianTriplets(t *testing.T) {
	pt := pythagorianTripletsStrct{}
	assert := assert.New(t)

	assert.True(pt.pythagorianTriplets([]int{3, 1, 4, 6, 5}))
	assert.False(pt.pythagorianTriplets([]int{10, 4, 6, 12, 5}))
}
