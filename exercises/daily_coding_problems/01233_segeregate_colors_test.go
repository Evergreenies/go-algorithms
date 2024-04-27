package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given an array of strictly the characters 'R', 'G', and 'B', segregate the values of the array
* so that all the Rs come first, the Gs come second, and the Bs come last. You can only swap elements of the array.
*
* Do this in linear time and in-place.
*
* For example, given the array ['G', 'B', 'R', 'R', 'B', 'R', 'G'], it should become ['R', 'R', 'R', 'G', 'G', 'B', 'B'].
* **/

func segeregateColors(arr []string) []string {
	left, middle, right := 0, 0, len(arr)-1

	for middle <= right {
		if arr[middle] == "R" {
			arr[left], arr[middle] = arr[middle], arr[left]
			left += 1
			middle += 1
		} else if arr[middle] == "B" {
			arr[middle], arr[right] = arr[right], arr[middle]
			right -= 1
		} else {
			middle += 1
		}
	}

	return arr
}

func TestSegeregateColors(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(segeregateColors([]string{"G", "B", "R", "R", "B", "R", "G"}), []string{"R", "R", "R", "G", "G", "B", "B"}, "this should return correct order")
}
