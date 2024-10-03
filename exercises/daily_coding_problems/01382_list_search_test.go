/*
* Describe an algorithm to compute the longest increasing subsequence of an array of numbers in O(n log n) time.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"
)

type lisSearch struct{}

type lisSearchInf interface {
	binarySearch([]int, int) int
	lengthofLIS([]int) (int, []int)
}

func (lisSearch) binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := int((left + right) / 2)
		if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}

func (lis lisSearch) lengthofLIS(arr []int) (int, []int) {
	tails := []int{}
	for _, item := range arr {
		pos := lis.binarySearch(tails, item)
		if pos < len(tails) {
			tails[pos] = item
		} else {
			tails = append(tails, item)
		}
	}

	return len(tails), tails
}

func TestLengthOfLIS(t *testing.T) {
	var lis lisSearchInf
	lisA := lisSearch{}
	lis = lisA

	size, arr := lis.lengthofLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Printf("%v\n", arr)
	if size != 4 {
		t.Errorf("expected 4, got %d", size)
	}
}
