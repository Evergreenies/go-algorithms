/*
*
* Given a node in a binary search tree, return the next bigger element, also known as the inorder successor.
* For example, the inorder successor of 22 is 30.

	  10
	 /  \
	5    30
	    /  \
	  22    35

* You can assume each node has a parent pointer.
*/
package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	maxSuccessorOfTree struct{}

	treeNodeS struct {
		Value  int
		Left   *treeNodeS
		Right  *treeNodeS
		Parent *treeNodeS
	}
)

func (ms maxSuccessorOfTree) findMin(node *treeNodeS) *treeNodeS {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (ms maxSuccessorOfTree) inorderSuccessor(node *treeNodeS) *treeNodeS {
	if node.Right != nil {
		return ms.findMin(node.Right)
	}

	current := node
	for current.Parent != nil && current.Parent.Right == current {
		current = current.Parent
	}

	return current.Parent
}

func TestTreeMaxSuccessor(t *testing.T) {
	root := &treeNodeS{Value: 10}
	node5 := &treeNodeS{Value: 5}
	node30 := &treeNodeS{Value: 30}
	node22 := &treeNodeS{Value: 22}
	node35 := &treeNodeS{Value: 35}

	// Set up the tree structure
	root.Left = node5
	root.Right = node30
	node5.Parent = root
	node30.Parent = root
	node30.Left = node22
	node30.Right = node35
	node22.Parent = node30
	node35.Parent = node30

	ms := maxSuccessorOfTree{}
	mx := ms.inorderSuccessor(node22)

	assert := assert.New(t)

	assert.Equal(mx.Value, 30, fmt.Sprintf("exprected 30; got %v\n", mx.Value))
}
