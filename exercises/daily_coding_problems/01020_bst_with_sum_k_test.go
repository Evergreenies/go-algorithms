/**
* Given the root of a binary search tree, and a target K, return two nodes in the tree whose sum equals K.
* For example, given the following tree and K of 20

    10
   /   \
 5      15
       /  \
     11    15

* Return the nodes 5 and 15.
* */

package dailycodingproblems

import (
	"fmt"
	"testing"
)

type bstSumUpToK struct{}

type bstTreeNode struct {
	value       int
	left, right *bstTreeNode
}

func (b bstSumUpToK) inoderTraversal(node *bstTreeNode, values *[]int) {
	if node == nil {
		return
	}

	b.inoderTraversal(node.left, values)
	*values = append(*values, node.value)
	b.inoderTraversal(node.right, values)
}

func (b bstSumUpToK) findBSTSumUpToK(node *bstTreeNode, k int) (int, int, bool) {
	values := []int{}
	b.inoderTraversal(node, &values)

	fmt.Println(len(values))

	left, right := 0, len(values)-1
	for left < right {
		current_sum := values[left] + values[right]
		if current_sum == k {
			return values[left], values[right], true
		} else if current_sum < k {
			left++
		} else {
			right--
		}
	}

	return 0, 0, false
}

func TestBSTSumUpToK(t *testing.T) {
	root := &bstTreeNode{value: 10}
	root.left = &bstTreeNode{value: 5}
	root.right = &bstTreeNode{value: 15}
	root.right.left = &bstTreeNode{value: 11}
	root.right.right = &bstTreeNode{value: 15}

	bst := bstSumUpToK{}
	left, right, matched := bst.findBSTSumUpToK(root, 20)
	if !matched {
		t.Errorf("expected (%d, %d), got (%d, %d)\n", 5, 15, left, right)
	}
}
