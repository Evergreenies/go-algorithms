/*
* Given an integer N, construct all possible binary search trees
* with N nodes.
* */

package dailycodingproblems

import (
	"testing"
)

type treeNodeStrct2 struct {
	Val         int
	Left, Right *treeNodeStrct2
}

type binaryTreeStrct2 struct {
	t *testing.T
}

func (b binaryTreeStrct2) generateSubtree(start, end int) []*treeNodeStrct2 {
	if start > end {
		return []*treeNodeStrct2{nil}
	}

	var allTrees []*treeNodeStrct2
	for index := start; index <= end; index++ {
		leftTrees := b.generateSubtree(start, index-1)
		rightTrees := b.generateSubtree(index+1, end)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currentTree := &treeNodeStrct2{Val: index}

				currentTree.Left = left
				currentTree.Right = right

				allTrees = append(allTrees, currentTree)
			}
		}
	}

	return allTrees
}

func (b binaryTreeStrct2) generateTrees(n int) []*treeNodeStrct2 {
	if n == 0 {
		return []*treeNodeStrct2{}
	}

	return b.generateSubtree(1, n)
}

func (b binaryTreeStrct2) printPreorder(root *treeNodeStrct2) {
	if root == nil {
		return
	}

	b.t.Logf("%d ", root.Val)

	b.printPreorder(root.Left)
	b.printPreorder(root.Right)
}

func TestGenerateTrees(t *testing.T) {
	b := &binaryTreeStrct2{t: t}
	n := 3

	trees := b.generateTrees(n)

	t.Logf("All possible BST's with %d nodes: (total number of trees %d)\n", n, len(trees))
	for _, tree := range trees {
		b.printPreorder(tree)
		t.Log()
	}
}
