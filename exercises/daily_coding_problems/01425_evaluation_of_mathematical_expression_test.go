/**
* Suppose an arithmetic expression is given as a binary tree. Each leaf is an integer and each
* internal node is one of '+', '−', '∗', or '/'.
*
* Given the root to such a tree, write a function to evaluate it.
*
* For example, given the following tree:
    *
   / \
  +    +
 / \  / \
3  2  4  5

* You should return 45, as it is (3 + 2) * (4 + 5).
* */

package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	evaluateArithmeticExpressionStrct struct{}
	Node                              struct {
		value       string
		left, right *Node
	}
)

func (ev evaluateArithmeticExpressionStrct) evaluate(root *Node) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		var num int
		fmt.Sscanf(root.value, "%d", &num)

		return num
	}

	leftValue := ev.evaluate(root.left)
	rightValue := ev.evaluate(root.right)

	switch root.value {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	case "*":
		return leftValue * rightValue
	case "/":
		if rightValue == 0 {
			panic("division by zero error")
		}

		return leftValue / rightValue
	default:
		panic("invalid operator")
	}
}

func TestEvaluateArithmeticExpression(t *testing.T) {
	root := &Node{
		value: "*",
		left: &Node{
			value: "+",
			left:  &Node{value: "3"},
			right: &Node{value: "2"},
		},
		right: &Node{
			value: "+",
			left:  &Node{value: "4"},
			right: &Node{value: "5"},
		},
	}

	ev := evaluateArithmeticExpressionStrct{}
	result := ev.evaluate(root)

	assert := assert.New(t)
	assert.Equal(45, result, fmt.Sprintf("expected 45; got %v\n", result))
}
