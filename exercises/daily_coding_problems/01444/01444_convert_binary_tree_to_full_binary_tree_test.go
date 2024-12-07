/*
Recall that a full binary tree is one in which each node is either a leaf node, or
has two children. Given a binary tree, convert it to a full one by removing nodes
with only one child.

For example, given the following tree:

         0
      /     \
    1         2
  /            \
3                 4
  \             /   \
    5          6     7
You should convert it to:

     0
  /     \
5         4
        /   \
       6     7
* */

package binarytreetofulltree

import (
	"fmt"
	"testing"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// ConvertFullBinaryTree removes nodes with only one child
func ConvertFullBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// recursively convert left and right subtrees
	root.Left = ConvertFullBinaryTree(root.Left)
	root.Right = ConvertFullBinaryTree(root.Right)

	// if node has only only onde chinld return that children
	if root.Left == nil && root.Right != nil {
		return root.Right
	}

	if root.Right == nil && root.Left != nil {
		return root.Left
	}

	return root
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Value)
	PrintTree(root.Left)
	PrintTree(root.Right)
}

func TestFBT(t *testing.T) {
	// Construct the tree
	root := &TreeNode{Value: 0}
	root.Left = &TreeNode{Value: 1}
	root.Right = &TreeNode{Value: 2}
	root.Left.Left = &TreeNode{Value: 3}
	root.Left.Left.Right = &TreeNode{Value: 5}
	root.Right.Right = &TreeNode{Value: 4}
	root.Right.Right.Left = &TreeNode{Value: 6}
	root.Right.Right.Right = &TreeNode{Value: 7}
	fmt.Println("Original tree (preorder):")
	PrintTree(root)

	// Convert to a full binary tree
	fullRoot := ConvertFullBinaryTree(root)
	fmt.Println("\nFull binary tree (preorder):")
	PrintTree(fullRoot)
}
