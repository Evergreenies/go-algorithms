/*
* A tree is symmetric if its data and shape remain unchanged when it is reflected about the root node.
* The following tree is an example:

        4
      / | \
    3   5   3
  /           \
9              9

* Given a k-ary tree, determine whether it is symmetric.
* */

package DCP01485

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Data     int
	Children []*Node
}

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}

	var areMirror func(n1, n2 *Node) bool
	areMirror = func(n1, n2 *Node) bool {
		if n1 == nil && n2 == nil {
			return true
		}

		if n1 == nil || n2 == nil || n1.Data != n2.Data {
			return false
		}

		if len(n1.Children) != len(n2.Children) {
			return false
		}

		for index := 0; index < len(n1.Children); index++ {
			if !areMirror(n1.Children[index], n2.Children[len(n2.Children)-1-index]) {
				return false
			}
		}

		return true
	}

	for index := 0; index < len(root.Children)/2; index++ {
		if !areMirror(root.Children[index], root.Children[len(root.Children)-1-index]) {
			return false
		}
	}

	return true
}

func TestIsSymmetric(t *testing.T) {
	root := &Node{Data: 4}
	child1 := &Node{Data: 3}
	child2 := &Node{Data: 5}
	child3 := &Node{Data: 3}

	root.Children = []*Node{child1, child2, child3}
	child1.Children = []*Node{{Data: 9}}
	child3.Children = []*Node{{Data: 9}}

	assert := assert.New(t)
	assert.True(isSymmetric(root))
}
