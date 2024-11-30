/**
* Given an array of numbers, find the length of the longest increasing subsequence in the array.
* The subsequence does not necessarily have to be contiguous.
*
* For example, given the array [0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15], the longest
* increasing subsequence has length 6: it is 0, 2, 6, 9, 11, 15.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type lisStruct struct{}

func (lisStruct) longestIncreasingSubsequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	// initialize an array to keep track of LIS
	lis := make([]int, len(arr))
	for index := range lis {
		lis[index] = 1
	}

	for index := 0; index < len(arr); index++ {
		for j := 0; j < index; j++ {
			if arr[index] > arr[j] && lis[index] < lis[j]+1 {
				lis[index] = lis[j] + 1
			}
		}
	}

	maxLength := 0
	for _, length := range lis {
		if length > maxLength {
			maxLength = length
		}
	}

	fmt.Println("LIS Index: ", lis)

	return maxLength
}

func TestLIS(t *testing.T) {
	arr := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	lis := &lisStruct{}
	length := lis.longestIncreasingSubsequence(arr)

	assert := assert.New(t)
	assert.Equal(6, length, fmt.Sprintf("expected 6; got %d", length))
}
