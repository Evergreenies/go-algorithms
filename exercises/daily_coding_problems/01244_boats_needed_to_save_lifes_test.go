package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* An imminent hurricane threatens the coastal town of Codeville. If at most two people can
* fit in a rescue boat, and the maximum weight limit for a given boat is k, determine how
* many boats will be needed to save everyone.
*
* For example, given a population with weights [100, 200, 150, 80] and a boat limit of 200,
* the smallest number of boats required will be three.
* **/

func bubbleSortToSortWeights(arr []int) []int {
	swap := true
	for swap {
		swap = false

		for index := 0; index < len(arr)-1; index++ {
			if arr[index] < arr[index+1] {
				arr[index], arr[index+1] = arr[index+1], arr[index]
				swap = true
			}
		}
	}

	return arr
}

func numOfBoatsToRescue(weight []int, limit int) int {
	boats, left, right := 0, 0, len(weight)-1
	weight = bubbleSortToSortWeights(weight)

	for left < right {
		if (weight[left] + weight[right]) <= limit {
			left++
		}

		boats++
		right--
	}

	return boats
}

func TestSortWeights(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(bubbleSortToSortWeights([]int{100, 200, 150, 80}), []int{200, 150, 100, 80}, "arr must be sorted now")
}

func TestBoatstoRescue(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(numOfBoatsToRescue([]int{100, 200, 150, 80}, 200), 3, "we must have 3 boats to rescue all peoples")
}
