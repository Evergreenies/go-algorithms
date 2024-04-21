package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given a list of possibly overlapping intervals, return a new list of intervals where all overlapping intervals have been merged.
*
* The input list is not necessarily ordered in any way.
*
* For example, given [(1, 3), (5, 8), (4, 10), (20, 25)], you should return [(1, 3), (4, 10), (20, 25)].
* **/

func sortIntervals(intervals [][]int) [][]int {
	length := len(intervals)

	isSwapped := true
	for isSwapped {
		isSwapped = false

		for index := 1; index < (length - 1); index++ {
			if intervals[index][0] > intervals[index+1][0] {
				intervals[index], intervals[index+1] = intervals[index+1], intervals[index]

				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}

	return intervals
}

func maxInterval(interval1, interval2 int) int {
	if interval1 > interval2 {
		return interval1
	}

	return interval2
}

func mergeIntervals(intervals [][]int) [][]int {
	intervals = sortIntervals(intervals)
	mergedIntervals := [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		if mergedIntervals[len(mergedIntervals)-1][1] >= interval[0] {
			mergedIntervals[len(mergedIntervals)-1][1] = maxInterval(mergedIntervals[len(mergedIntervals)-1][1], interval[1])
		} else {
			mergedIntervals = append(mergedIntervals, interval)
		}
	}

	return mergedIntervals
}

func TestMergeIntervals(t *testing.T) {
	assert := assert.New(t)
	intervals := [][]int{
		{1, 3},
		{5, 8},
		{4, 10},
		{20, 25},
	}
	mergedIntervals := mergeIntervals(intervals)

	assert.Equal(mergedIntervals, [][]int{{1, 3}, {4, 10}, {20, 25}}, "it must have return valid intervals")
}
