package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Write a program to merge two binary trees. Each node in the new tree should hold a value equal to the sum
* of the values of the corresponding nodes of the input trees.
*
* If only one input tree has a node in a given position, the corresponding node in the new tree should match
* that input node.
* **/

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func newNode(value1, value2 int) *TreeNode {
	return &TreeNode{
		value: value1 + value2,
	}
}

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}

	newNd := newNode(root1.value, root2.value)
	newNd.left = mergeTrees(root1.left, root2.left)
	newNd.right = mergeTrees(root1.right, root2.right)

	return newNd
}

func printTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	nodeValues := make([]int, 0)
	for len(queue) > 0 {
		level := len(queue)

		for index := 0; index < level; index++ {
			currentNode := queue[0]
			queue = queue[1:]
			nodeValues = append(nodeValues, currentNode.value)

			fmt.Print(currentNode.value, " -> ")
			if currentNode.left != nil {
				queue = append(queue, currentNode.left)
			}

			if currentNode.right != nil {
				queue = append(queue, currentNode.right)
			}
		}
	}

	fmt.Println("nil")

	return nodeValues
}

func TestTreeSum(t *testing.T) {
	assert := assert.New(t)

	root1 := &TreeNode{
		value: 10,
		left:  &TreeNode{20, nil, nil},
		right: &TreeNode{30, nil, nil},
	}
	root2 := &TreeNode{
		value: 40,
		left:  &TreeNode{50, nil, nil},
		right: &TreeNode{60, nil, nil},
	}

	mergedRoot := mergeTrees(root1, root2)
	result := printTree(mergedRoot)
	fmt.Println(result)

	assert.Equal(result, []int{50, 70, 90}, "this tree must have some result in it")
}
