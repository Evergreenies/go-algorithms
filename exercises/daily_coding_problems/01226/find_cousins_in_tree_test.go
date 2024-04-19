package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Two nodes in a binary tree can be called cousins if they are on the same level of the tree but have
* different parents. For example, in the following diagram 4 and 6 are cousins.
*       1
*      / \
*     2   3
*    / \   \
*   4   5   6
*
* Given a binary tree and a particular node, find all cousins of that node.
* **/

type TreeNode struct {
	value any
	left  *TreeNode
	right *TreeNode
}

func findCousins(root *TreeNode, cousinsOfNode *TreeNode) []interface{} {
	level := findLevelOfTargetNode(root, cousinsOfNode, 1)
	cousins := make([]interface{}, 0)

	dfs(root, cousinsOfNode, nil, 0, level, &cousins)

	return cousins
}

func findLevelOfTargetNode(node *TreeNode, cousinsOfNode *TreeNode, level int) int {
	if node == nil {
		return -1
	}

	if node == cousinsOfNode {
		return level
	}

	leftLevel := findLevelOfTargetNode(node.left, cousinsOfNode, level+1)
	if leftLevel != -1 {
		return leftLevel
	}

	return findLevelOfTargetNode(node.right, cousinsOfNode, level+1)
}

func dfs(node *TreeNode, cousinsOfNode *TreeNode, parent *TreeNode, currentLevel int, level int, cousins *[]interface{}) {
	if node == nil {
		return
	}

	if (currentLevel == level) && (node != parent) {
		*cousins = append(*cousins, node.value)
	}

	dfs(node.left, cousinsOfNode, node, currentLevel+1, level, cousins)
	dfs(node.right, cousinsOfNode, node, currentLevel+1, level, cousins)
}

func TestFindCousins(t *testing.T) {
	assert := assert.New(t)

	root := &TreeNode{
		value: 1,
		left: &TreeNode{
			value: 2,
			right: &TreeNode{value: 4, left: nil, right: nil},
		},
		right: &TreeNode{
			value: 3,
			left: &TreeNode{
				value: 5,
				left:  &TreeNode{value: 6, left: nil, right: nil},
			},
			right: nil,
		},
	}

	cousinsOfTarget := findCousins(root, root.left.right)

	assert.Equal(cousinsOfTarget, []interface{}{6}, "must be 6 as cousin only")
}
