package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given a set of distinct positive integers, find the largest subset such that
* every pair of elements in the subset (i, j) satisfies
* either i % j = 0 or j % i = 0.
*
* For example, given the set [3, 5, 10, 20, 21], you should return [5, 10, 20].
* Given [1, 3, 6, 24], return [1, 3, 6, 24].
* **/

func insertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	for index := 1; index < len(arr); index++ {
		currentElement := arr[index]
		prevIndex := index - 1

		for prevIndex >= 0 && currentElement < arr[prevIndex] {
			arr[prevIndex+1] = arr[prevIndex]
			prevIndex--
		}

		arr[prevIndex+1] = currentElement
	}

	return arr
}

func getMaxKey(key1, key2 int) int {
	if key1 > key2 {
		return key1
	}

	return key2
}

func largestDivisibleSubset(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	arr = insertionSort(arr)
	result := make([]int, 0)
	dp := make(map[int][]int, 0)

	for _, num := range arr {
		isDivisible := true
		for _, element := range result {
			if (num%element) != 0 && (element%num) != 0 {
				isDivisible = false
				dp[len(result)] = result
				result = []int{num}

				break
			}
		}

		if isDivisible {
			result = append(result, num)
		}
	}

	dp[len(result)] = result
	maxKey := 0

	for key := range dp {
		maxKey = getMaxKey(maxKey, key)
	}

	return dp[maxKey]
}

func TestLargestDivisibleSubset(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(largestDivisibleSubset([]int{3, 5, 10, 20, 21}), []int{5, 10, 20}, "this should match case")
	assert.Equal(largestDivisibleSubset([]int{1, 3, 6, 24}), []int{1, 3, 6, 24}, "this should match case")
}
