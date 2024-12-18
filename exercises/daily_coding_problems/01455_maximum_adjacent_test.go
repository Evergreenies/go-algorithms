/*
You are given an array of arrays of integers, where each array corresponds to a row in
a triangle of numbers. For example, [[1], [2, 3], [1, 5, 1]] represents the triangle:

  1
 2 3
1 5 1
We define a path in the triangle to start at the top and go down one row at a time to
an adjacent value, eventually ending with an entry on the bottom row.

For example, 1 -> 3 -> 5. The weight of the path is the sum of the entries.

Write a program that returns the weight of the maximum weight path.
* */

package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type maxSumAdjacentStrct struct{}

func (m maxSumAdjacentStrct) maxSumAdjacent(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += m.max(triangle[row+1][col], triangle[row+1][col+1])
		}
	}

	return triangle[0][0]
}

func (maxSumAdjacentStrct) max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestMaxSumAdjacent(t *testing.T) {
	triangle := [][]int{
		{1},
		{2, 3},
		{1, 5, 1},
	}

	m := maxSumAdjacentStrct{}
	assert := assert.New(t)

	assert.Equal(9, m.maxSumAdjacent(triangle))
}
