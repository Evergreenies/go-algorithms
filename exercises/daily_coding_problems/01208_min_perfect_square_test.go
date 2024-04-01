package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minPerfectSquare(nth int) int {
	dp := make([]int, nth+1)
	for index := range dp {
		dp[index] = int(1e9)
	}

	dp[0] = 0

	for index := 1; index*index <= nth; index++ {
		square := index * index
		if square > nth {
			break
		}

		for index_j := square; index_j <= nth; index_j++ {
			dp[index_j] = min(dp[index_j], (dp[index_j-square] + 1))
		}
	}

	return dp[nth]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func TestMinPerfectSquare(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(minPerfectSquare(4), 1, "1 is min perfect square of 4")
	assert.Equal(minPerfectSquare(17), 2, "2 is min perfect square of 17")
	assert.Equal(minPerfectSquare(18), 2, "2 is min perfect square of 18")
}
