/*
* Given a positive integer n, find the smallest number of squared integers which sum to n.
*
* For example, given n = 13, return 2 since 13 = 32 + 22 = 9 + 4.
* Given n = 27, return 3 since 27 = 32 + 32 + 32 = 9 + 9 + 9.
* */

package dailycodingproblems

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type minSquaredStrct struct{}

func (m minSquaredStrct) numSquares(n int) int {
	dp := make([]int, n+1)
	for index := range dp {
		dp[index] = math.MaxInt32
	}
	dp[0] = 0

	for index := 1; index <= n; index++ {
		for j := 1; j*j <= index; j++ {
			dp[index] = m.min(dp[index], dp[index-j*j]+1)
		}
	}

	return dp[n]
}

func (minSquaredStrct) min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func TestNumSquares(t *testing.T) {
	m := minSquaredStrct{}
	assert := assert.New(t)

	assert.Equal(2, m.numSquares(13))
	assert.Equal(3, m.numSquares(27))
}
