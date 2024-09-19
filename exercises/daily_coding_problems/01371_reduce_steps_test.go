/*
* Consider the following scenario: there are N mice and N holes placed at integer points along a line.
* Given this, find a method that maps mice to holes such that the largest number of steps any mouse takes is minimized.
* Each move consists of moving one mouse one unit to the left or right, and only one mouse can fit inside each hole.
* For example, suppose the mice are positioned at [1, 4, 9, 15], and the holes are located at [10, -5, 0, 16].
* In this case, the best pairing would require us to send the mouse at 1 to the hole at -5, so our function
* should return 6.
* */

package dailycodingproblems

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MaxDistanceAnyMouseTravel struct{}

func (MaxDistanceAnyMouseTravel) bubbleSort(arr []int) {
	length := len(arr)

	for indexI := 0; indexI < length-1; indexI++ {
		for indexJ := 0; indexJ < length-indexI-1; indexJ++ {
			if arr[indexJ] > arr[indexJ+1] {
				arr[indexJ], arr[indexJ+1] = arr[indexJ+1], arr[indexJ]
			}
		}
	}
}

func (m MaxDistanceAnyMouseTravel) minMaxDistance(mice []int, holes []int) int {
	if len(mice) != len(holes) {
		return -1
	}

	// sort.Ints(mice)
	// sort.Ints(holes)
	m.bubbleSort(mice)
	m.bubbleSort(holes)

	maxDistance := 0
	for index := 0; index < len(mice); index++ {
		distance := int(math.Abs(float64(mice[index] - holes[index])))

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func TestMinMaxDistance(t *testing.T) {
	mice := []int{1, 4, 9, 15}
	holes := []int{10, -5, 0, 10}

	assert := assert.New(t)

	mx := MaxDistanceAnyMouseTravel{}
	assert.True(assert.Equal(mx.minMaxDistance(mice, holes), 6), "extected 6")
}
