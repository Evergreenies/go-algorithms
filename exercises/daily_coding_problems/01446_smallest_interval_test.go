/*
Given a set of closed intervals, find the smallest set of numbers that covers all the intervals.
If there are multiple smallest sets, return any of them.

For example, given the intervals [0, 3], [2, 6], [3, 4], [6, 9], one set of numbers that covers
all these intervals is {3, 6}.
* */

package dailycodingproblems

import (
	"sort"
	"testing"
)

// interval represents a closed interval [start, end]
type interval struct {
	start, end int
}

// findMinSetOfIntervals finds the smallest set of numbers that covers all
func findMinSetOfIntervals(intervals []interval) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	var result []int
	current := intervals[0].end
	result = append(result, current)

	for _, interval := range intervals {
		if interval.start > current {
			current = interval.end
			result = append(result, current)
		}
	}

	return result
}

func TestFindMinSetOfIntervals(t *testing.T) {
	intervals := []interval{
		{0, 3},
		{2, 6},
		{3, 4},
		{6, 9},
	}
	result := findMinSetOfIntervals(intervals)
	t.Logf("expected [3 9]; got %v\n", result)
}
