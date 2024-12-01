/*
*
You are given a histogram consisting of rectangles of different heights. These heights are
represented in an input list, such that [1, 3, 2, 5] corresponds to the following diagram:

	    x
	    x
	x   x
	x x x

x x x x

Determine the area of the largest rectangle that can be formed only from the bars of the
histogram. For the diagram above, for example, this would be six, representing the 2 x 3
area at the bottom right.
*
*/
package dailycodingproblems

import "testing"

type largestRectangleAeraStrct struct{}

func (largestRectangleAeraStrct) largestRectangleAera(arr []int) int {
	length, stack := len(arr), []int{}
	maxArea, index := 0, 0

	for index < length {
		if len(stack) == 0 || arr[index] >= arr[stack[len(stack)-1]] {
			stack = append(stack, index)
			index++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := arr[top]
			width := index

			if len(stack) != 0 {
				width = index - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, height+width)
		}
	}

	if len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		height := arr[top]
		width := index

		if len(stack) != 0 {
			width = index - stack[len(stack)-1] - 1
		}
		maxArea = max(maxArea, height+width)
	}

	return maxArea
}

func TestLargestRectangleArea(t *testing.T) {
	rgst := largestRectangleAeraStrct{}
	val := rgst.largestRectangleAera([]int{1, 3, 2, 5})
	t.Logf("Largest rectangle area: %d", val)
}
