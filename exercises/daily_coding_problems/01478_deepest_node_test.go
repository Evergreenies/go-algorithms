/*
* Given the root of a binary tree, return a deepest node. For example,
* in the following tree, return d.

    a
   / \
  b   c
 /
d
* */

package dailycodingproblems

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type treeNodeStrct1 struct {
	Val string

	Left  *treeNodeStrct1
	Right *treeNodeStrct1
}

func findDeepestNode(root *treeNodeStrct1) *treeNodeStrct1 {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)
	var deepest *treeNodeStrct1

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*treeNodeStrct1)
		deepest = node

		if node.Left != nil {
			queue.PushBack(node.Left)
		}

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return deepest
}

func TestFindDeepestNode(t *testing.T) {
	root := &treeNodeStrct1{Val: "a"}
	root.Left = &treeNodeStrct1{Val: "b"}
	root.Right = &treeNodeStrct1{Val: "c"}
	root.Left.Left = &treeNodeStrct1{Val: "d"}

	assert := assert.New(t)
	assert.NotNil(findDeepestNode(root))
}
