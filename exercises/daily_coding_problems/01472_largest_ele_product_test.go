/*
* Given a list of integers, return the largest product that can be made by multiplying any three integers.

For example, if the list is [-10, -10, 5, 2], we should return 500, since that's -10 * -10 * 5.

You can assume the list has at least three integers.
* */

package dailycodingproblems

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type largestEleProdStrct struct{}

func (largestEleProdStrct) largestProductOfThree(arr []int) int {
	sort.Ints(arr)
	length := len(arr)

	product1 := arr[length-1] * arr[length-2] * arr[length-3]
	product2 := arr[0] * arr[1] * arr[length-1]

	if product1 > product2 {
		return product1
	}

	return product2
}

func TestLargestProductOfThree(t *testing.T) {
	ls := largestEleProdStrct{}
	assert := assert.New(t)

	assert.Equal(500, ls.largestProductOfThree([]int{-10, -10, 5, 2}))
}
