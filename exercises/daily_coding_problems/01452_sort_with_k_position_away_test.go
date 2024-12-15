/*
You are given a list of N numbers, in which each number is located at most k places away
from its sorted position. For example, if k = 1, a given element at index 4 might end up
at indices 3, 4, or 5.

Come up with an algorithm that sorts this list in O(N log k) time.
* */

package dailycodingproblems

import (
	"container/heap"
	"testing"
)

type meanHeap []int

func (h meanHeap) Len() int                 { return len(h) }
func (h meanHeap) Less(idx1, idx2 int) bool { return h[idx1] < h[idx2] }
func (h meanHeap) Swap(idx1, idx2 int)      { h[idx1], h[idx2] = h[idx2], h[idx1] }
func (h *meanHeap) Push(x interface{})      { *h = append(*h, x.(int)) }

func (h *meanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func sortKSortedArray(arr []int, k int) []int {
	h := &meanHeap{}
	heap.Init(h)

	result := make([]int, 0, len(arr))
	for index := 0; index < k+1 && index < len(arr); index++ {
		heap.Push(h, arr[index])
	}

	for index := k + 1; index < len(arr); index++ {
		heap.Push(h, arr[index])

		result = append(result, heap.Pop(h).(int))
	}

	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(int))
	}

	return result
}

func TestSortKSortedArray(t *testing.T) {
	t.Log(sortKSortedArray([]int{3, 2, 1, 5, 4, 7, 6, 5}, 2))
}
