/*
*	In front of you is a row of N coins, with values v1, v1, ..., vn.
*	You are asked to play the following game. You and an opponent take turns choosing either
*	the first or last coin from the row, removing it from the row, and receiving the value
*	of the coin.
*
*	Write a program that returns the maximum amount of money you can win with certainty, if
*	you move first, assuming your opponent plays optimally.
* */

package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type maxCoinGame struct{}

func (mc maxCoinGame) maxCoinsWin(arr []int) int {
	length := len(arr)
	dp := make([][]int, length)

	// initialize dp table
	for index := range dp {
		dp[index] = make([]int, length)
	}

	for gap := 0; gap < length; gap++ {
		for row, col := 0, gap; col < length; row, col = row+1, col+1 {
			x, y, z := 0, 0, 0
			if row+2 <= col {
				x = dp[row+1][col]
			}
			if row+1 <= col-1 {
				y = dp[row+1][col-1]
			}
			if row <= col-2 {
				z = dp[row][col-2]
			}

			dp[row][col] = mc.max(arr[row]+min(x, y), arr[col]+min(y, z))
		}
	}

	return dp[0][length-1]
}

func (maxCoinGame) max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func (maxCoinGame) min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func TestMaxCoinWin(t *testing.T) {
	mc := maxCoinGame{}
	assert := assert.New(t)

	assert.Equal(28, mc.maxCoinsWin([]int{2, 4, 6, 8, 10, 2, 6, 10}))
}
