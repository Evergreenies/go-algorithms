/**
* Given a list of elements, find the majority element, which appears more than half the time
* (> floor(len(lst) / 2.0)).
*
* You can assume that such element exists.
* For example, given [1, 2, 1, 1, 3, 4, 0], return 1
* */

package dailycodingproblems

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type boyersMooreStrct struct{}

func (boyersMooreStrct) count(candidate int, arr []int) int {
	counter := 0
	for _, item := range arr {
		if item == candidate {
			counter++
		}
	}

	return counter
}

func (bm boyersMooreStrct) findMajorityElements(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {
		return arr[0]
	}

	counter, candidate := 0, arr[0]
	for _, item := range arr {
		if counter == 0 {
			candidate = item
		}

		if item == candidate {
			counter++
		} else {
			counter--
		}
	}

	if bm.count(candidate, arr) > int(math.Floor(float64(len(arr)/2))) {
		return candidate
	}

	return -1
}

func TestBoyersMooreVotingAlgo(t *testing.T) {
	assert := assert.New(t)

	sm := new(boyersMooreStrct)
	candidate := sm.findMajorityElements([]int{1, 2, 1, 1, 3, 4, 0})
	assert.Equal(-1, candidate, fmt.Sprintf("expected -1; got %v\n", candidate))

	candidate = sm.findMajorityElements([]int{1, 2, 1, 1, 3, 4, 1})
	assert.Equal(1, candidate, fmt.Sprintf("expected 1; got %v\n", candidate))
}
