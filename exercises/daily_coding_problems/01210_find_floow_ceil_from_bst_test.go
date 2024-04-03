package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given a binary search tree, find the floor and ceiling of a given integer.
* The floor is the highest element in the tree less than or equal to an integer,
* while the ceiling is the lowest element in the tree greater than or equal to an integer.
*
* If either value does not exist, return None.
* **/

type Node1 struct {
	value int
	left  *Node1
	right *Node1
}

type MyResultSet1 struct {
	floor *int
	ceil  *int
}

func findFloorCeilFromBST(root *Node1, target int) MyResultSet1 {
	var floor, ceil int
	currentNode := root

	for currentNode != nil {
		if currentNode.value == target {
			return MyResultSet1{&currentNode.value, &currentNode.value}
		}

		if target < currentNode.value {
			ceil = currentNode.value
			currentNode = currentNode.left
		} else {
			floor = currentNode.value
			currentNode = currentNode.right
		}
	}

	return MyResultSet1{&floor, &ceil}
}

func TestFindFloorCeilFromBST(t *testing.T) {
	root := &Node1{
		value: 10,
		left:  &Node1{5, &Node1{value: 3}, &Node1{value: 7}},
		right: &Node1{value: 15, right: &Node1{value: 18}},
	}

	assert := assert.New(t)

	var floor, ceil int
	floor = 10
	ceil = 15
	assert.Equal(findFloorCeilFromBST(root, 12), MyResultSet1{&floor, &ceil}, "this should return floor as 10 and ceil as 15")

	floor = 18
	ceil = 0
	assert.Equal(findFloorCeilFromBST(root, 20), MyResultSet1{&floor, &ceil}, "this should return floor as 18 and ceil as `nil`")
}
