package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given the root of a binary tree, find the most frequent subtree sum. The subtree sum of a node is the
* sum of all values under a node, including the node itself.
* **/

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func dfs(node *TreeNode, subtreeSumCounts *map[int]int) int {
	if node == nil {
		return 0
	}

	subtreeSum := node.value + dfs(node.left, subtreeSumCounts) + dfs(node.right, subtreeSumCounts)
	_sum, ok := (*subtreeSumCounts)[subtreeSum]
	if ok {
		(*subtreeSumCounts)[subtreeSum] = _sum + 1
	} else {
		(*subtreeSumCounts)[subtreeSum] = 1
	}

	return subtreeSum
}

func findMostFrequentSum(root *TreeNode) []int {
	subtreeSumCounts := make(map[int]int, 0)
	dfs(root, &subtreeSumCounts)

	maxCount := 0
	mostFrequentSum := 0
	for key, count := range subtreeSumCounts {
		if count > maxCount {
			maxCount = count
			mostFrequentSum = key
		}
	}

	return []int{mostFrequentSum, maxCount}
}

func TestMostFrequentSum(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		value: 5,
		left:  &TreeNode{1, nil, nil},
		right: &TreeNode{
			value: 9,
			left:  &TreeNode{5, nil, nil},
			right: &TreeNode{2, nil, nil},
		},
	}

	mostFrequentSum := findMostFrequentSum(root)
	assert.Equal(mostFrequentSum[0], 1, "it must have only one most frequent sum")

	root = &TreeNode{
		value: 5,
		left:  &TreeNode{2, nil, nil},
		right: &TreeNode{-5, nil, nil},
	}

	mostFrequentSum = findMostFrequentSum(root)
	assert.Equal(mostFrequentSum[0], 2, "it must have 2 as most frequent sum")
}
