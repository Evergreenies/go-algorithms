/**
* Given an integer list where each number represents the number of hops you can make, determine whether
* you can reach to the last index starting at index 0.
*
* For example, [2, 0, 1, 0] returns True while [1, 1, 0, 1] returns False.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"
)

type canReachEnd struct{}

func (canReachEnd) canReachEd(arr []int) bool {
	max_reachable := 0

	for index, hop := range arr {
		if index > max_reachable {
			return false
		}

		max_reachable = max(max_reachable, index+hop)
	}

	return max_reachable >= len(arr)-1
}

func TestCanReachEnd(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{2, 0, 1, 0}, true},
		{[]int{1, 1, 0, 1}, false},
	}

	cx := canReachEnd{}
	for index, test := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			result := cx.canReachEd(test.input)
			if result != test.expected {
				t.Errorf("expected %v; got %v", test.expected, result)
			}
		})
	}
}
