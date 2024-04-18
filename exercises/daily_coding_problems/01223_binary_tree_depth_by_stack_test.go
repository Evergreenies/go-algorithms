package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* You are given a binary tree in a peculiar string representation. Each node is written in
* the form (lr), where l corresponds to the left child and r corresponds to the right child.
*
* If either l or r is null, it will be represented as a zero. Otherwise, it will be represented
* by a new (lr) pair.
*
* Here are a few examples:
*
* A root node with no children: (00)
* A root node with two children: ((00)(00))
* An unbalanced tree with three consecutive left children: ((((00)0)0)0)
*
* Given this representation, determine the depth of the tree.
* **/

func getMaxOfStack(depth int, stack []string) int {
	stackLen := len(stack)
	if depth > stackLen {
		return depth
	}

	return stackLen
}

func binaryTreeMaxDepth(str string) int {
	stack, depth := make([]string, 0), 0
	for _, char := range str {
		strChar := fmt.Sprintf("%s", string(char))
		if strChar == "(" {
			stack = append(stack, strChar)
			depth = getMaxOfStack(depth, stack)
		} else if strChar == ")" {
			stack = stack[:len(stack)-1]
		}
	}

	return depth
}

func TestMaxDepthOfTree(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(binaryTreeMaxDepth("(00)"), 1, "(00) this has just one node as root node")
	assert.Equal(binaryTreeMaxDepth("((00)(00))"), 2, "((00)(00)) this must have 2 nodes only so depth would be 2")
	assert.Equal(binaryTreeMaxDepth("((((00)0)0)0)"), 4, "((((00)0)0)0) it have depth of 4 only")
}
