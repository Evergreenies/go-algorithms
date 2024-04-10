package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Write an algorithm that finds the total number of set bits in all integers between 1 and N.
* **/

func brianKerninghanAlgo(num int) int {
	count := 0
	for num != 0 {
		count += num & 1
		num = num >> 1
	}

	return count
}

func countSetBits(n int) int {
	total_set_bits := 0
	for num := 0; num <= n; num++ {
		total_set_bits += brianKerninghanAlgo(num)
	}

	return total_set_bits
}

func TestCountSetBits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(countSetBits(5), 7, "total set bit count for 5 must be 7")
	assert.Equal(countSetBits(0), 0, "total set bit count for 0 must be 0")
	assert.Equal(countSetBits(7), 12, "total set bit count for 7 must be 12")
}
