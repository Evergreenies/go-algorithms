/*
* A regular number in mathematics is defined as one which evenly divides some power of 60.
* Equivalently, we can say that a regular number is one whose only prime divisors are 2, 3, and 5.
*
* These numbers have had many applications, from helping ancient Babylonians keep time to tuning
* instruments according to the diatonic scale.
*
* Given an integer N, write a program that returns, in order, the first N regular numbers.
* */
package minHeapPkg

import (
	"container/heap"
	"testing"
)

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	length := len(old)
	item := old[length-1]
	*h = old[0 : length-1]

	return item
}

func generateRegularNumbers(n int) []int {
	if n < 1 {
		return []int{}
	}

	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, 1)

	seen := map[int]bool{1: true}
	primes := []int{2, 3, 5}
	result := []int{}

	for len(result) < n {
		current := heap.Pop(h).(int)
		result = append(result, current)

		for _, prime := range primes {
			newNum := current * prime
			if !seen[newNum] {
				heap.Push(h, newNum)
				seen[newNum] = true
			}
		}
	}

	return result
}

func TestGenerateRegularNumbers(t *testing.T) {
	t.Log(generateRegularNumbers(10))
}
