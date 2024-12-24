/*
* You are given N identical eggs and access to a building with k floors. Your task is to
* find the lowest floor that will cause an egg to break, if dropped from that floor. Once
* an egg breaks, it cannot be dropped again. If an egg breaks when dropped from the xth
* floor, you can assume it will also break when dropped from any floor greater than x.

Write an algorithm that finds the minimum number of trial drops it will take, in the worst
case, to identify this floor.

For example, if N = 1 and k = 5, we will need to try dropping the egg at every floor,
beginning with the first, until we reach the fifth floor, so our solution will be 5.
*
*/
package dailycodingproblems

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type eggDropStrct struct{}

func (e eggDropStrct) eggDrop(n, k int) int {
	dp := make([][]int, n+1)
	for index := range dp {
		dp[index] = make([]int, k+1)
	}

	for j := 1; j <= k; j++ {
		dp[1][j] = j
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = math.MaxInt32

			for x := 1; x <= j; x++ {
				res := 1 + e.max(dp[i-1][x-1], dp[i][j-x])
				if res < dp[i][j] {
					dp[i][j] = res
				}
			}
		}
	}

	return dp[n][k]
}

func (eggDropStrct) max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func TestEggDrop(t *testing.T) {
	assert := assert.New(t)
	e := eggDropStrct{}

	assert.Equal(4, e.eggDrop(2, 10))
	// assert.Equal(5, e.eggDrop(1, 5))
}
