/*
* Given a array of numbers representing the stock prices of a company in chronological order,
* write a function that calculates the maximum profit you could have made from buying and selling
* that stock once. You must buy before you can sell it.
*
* For example, given [9, 11, 8, 5, 7, 10], you should return 5, since you could buy the stock at
* 5 dollars and sell it at 10 dollars.
 */

package dailycodingproblems

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stockProfit struct{}

func (stockProfit) maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit, minPrice := 0, math.MaxInt64
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	sp := stockProfit{}
	assert.Equal(t, 5, sp.maxProfit([]int{9, 11, 8, 5, 7, 10}))
}
