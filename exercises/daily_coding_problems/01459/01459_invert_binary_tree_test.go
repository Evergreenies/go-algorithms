/*
* Invert a binary tree.
*
* For example, given the following tree:

    a
   / \
  b   c
 / \  /
d   e f

should become:

  a
 / \
 c  b
 \  / \
  f e  d
* */

package invert_binary_tree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	val         string
	left, right *TreeNode
}

func invertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.left, root.right = root.right, root.left

	invertBinaryTree(root.left)
	invertBinaryTree(root.right)

	return root
}

func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}

	printInOrder(root.left)
	fmt.Print(root.val, " -> ")
	printInOrder(root.right)
}

func TestInvertBinaryTree(t *testing.T) {
	// Create the binary tree
	root := &TreeNode{val: "a"}
	root.left = &TreeNode{val: "b"}
	root.right = &TreeNode{val: "c"}
	root.left.left = &TreeNode{val: "d"}
	root.left.right = &TreeNode{val: "e"}
	root.right.left = &TreeNode{val: "f"}

	// Print original tree
	t.Log("Original tree in-order: ")
	printInOrder(root)
	t.Log()

	// Invert the tree
	invertedRoot := invertBinaryTree(root)
	// Print inverted tree
	t.Log("Inverted tree in-order: ")
	printInOrder(invertedRoot)
	t.Log()
}
