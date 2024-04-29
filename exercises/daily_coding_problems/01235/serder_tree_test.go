package dailycodingproblems

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given the root to a binary tree, implement serialize(root), which serializes the tree into a string,
* and deserialize(s), which deserializes the string back into the tree.
* **/

type Node struct {
	value string
	left  *Node
	right *Node
}

func serialize(root *Node) string {
	result := []string{}

	var preoderDFS func(node *Node)
	preoderDFS = func(node *Node) {
		if node == nil {
			result = append(result, "#")

			return
		}

		result = append(result, node.value)
		preoderDFS(node.left)
		preoderDFS(node.right)
	}

	preoderDFS(root)

	return strings.Join(result, ",")
}

func deserialize(str string) *Node {
	nodes := strings.Split(str, ",")

	var buildTree func(index int) *Node
	buildTree = func(index int) *Node {
		if nodes[index] == "#" {
			index++
			return nil
		}

		root := &Node{value: nodes[index]}
		index++
		root.left = buildTree(index)
		index++
		root.right = buildTree(index)

		return root
	}

	return buildTree(0)
}

func TestSerdeTree(t *testing.T) {
	assert := assert.New(t)

	root := &Node{value: "root", left: &Node{value: "left", left: &Node{value: "left.left"}}, right: &Node{value: "right"}}
	assert.Equal(deserialize(serialize(root)).left.left.value, "left.left", "this must deserializes the tree")
}
