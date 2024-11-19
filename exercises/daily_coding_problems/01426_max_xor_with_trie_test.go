/**
* Given an array of integers, find the maximum XOR of any two elements.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	trieNode struct {
		left, right *trieNode
	}

	trieStrct struct{}
)

func (t *trieStrct) insert(num int, root *trieNode) {
	node := root
	for index := 31; index >= 0; index-- {
		if ((num >> index) & 1) == 0 {
			if node.left == nil {
				node.left = &trieNode{}
			}

			node = node.left
		} else {
			if node.right == nil {
				node.right = &trieNode{}
			}

			node = node.right
		}
	}
}

func (t *trieStrct) findMaximumXOR(arr []int) int {
	root := &trieNode{}
	maxXor := 0

	for _, num := range arr {
		t.insert(num, root)
	}

	for _, num := range arr {
		node := root
		currXor := 0

		for index := 31; index >= 0; index-- {
			if ((num >> index) & 1) == 0 {
				if node.right != nil {
					currXor += (1 << index)
					node = node.right
				} else {
					node = node.left
				}
			} else {
				if node.left != nil {
					currXor += (1 << index)
					node = node.left
				} else {
					node = node.right
				}
			}
		}

		if currXor > maxXor {
			maxXor = currXor
		}
	}

	return maxXor
}

func TestMaxXORWithTrie(t *testing.T) {
	tr := trieStrct{}
	mx := tr.findMaximumXOR([]int{3, 10, 5, 25, 2, 8})

	assert := assert.New(t)
	assert.Equal(28, mx, fmt.Sprintf("expected 28; got %v\n", mx))
}
