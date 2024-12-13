/*
You are given an array of non-negative integers that represents a two-dimensional elevation
map where each element is unit-width wall and the integer is the height. Suppose it will
rain and all spots between two walls get filled up.

Compute how many units of water remain trapped on the map in O(N) time and O(1) space.

For example, given the input [2, 1, 2], we can hold 1 unit of water in the middle.

Given the input [3, 0, 1, 3, 0, 5], we can hold 3 units in the first index, 2 in the second,
and 3 in the fourth index (we cannot hold 5 since it would run off to the left), so we can
trap 8 units of water.
*
*/
package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type trappedWater struct{}

func (trappedWater) trappedWaterLogic(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax, totalWater := 0, 0, 0

	for left < right {
		if height[left] <= height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}

			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}

			right--
		}
	}

	return totalWater
}

func TestTrappedWaterLogic(t *testing.T) {
	assert := assert.New(t)
	tw := trappedWater{}

	assert.Equal(1, tw.trappedWaterLogic([]int{2, 1, 2}))
	assert.Equal(8, tw.trappedWaterLogic([]int{3, 0, 1, 3, 0, 5}))
	assert.Equal(0, tw.trappedWaterLogic([]int{5, 4, 3, 2, 1}))
}
