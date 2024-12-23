/*
* A fixed point in an array is an element whose value is equal to its index. Given a sorted
* array of distinct elements, return a fixed point, if one exists. Otherwise, return False.
*
* For example, given [-6, 0, 2, 40], you should return 2. Given [1, 5, 7, 8], you should return False.
* */

package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fixedIndexArrStrct struct{}

func (fixedIndexArrStrct) fixedIndexArr(arr []int) bool {
	for index, value := range arr {
		if index == value {
			return true
		}
	}

	return false
}

func TestFixedIndexArr(t *testing.T) {
	assert := assert.New(t)
	ia := fixedIndexArrStrct{}

	assert.True(ia.fixedIndexArr([]int{-6, 0, 2, 40}))
	assert.False(ia.fixedIndexArr([]int{1, 5, 7, 8}))
}
