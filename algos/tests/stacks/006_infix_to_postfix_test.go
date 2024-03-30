package stacks_test

import (
	"strings"
	"testing"

	customstack "github.com/Evergreenies/go-algorithms/algos/stacks/custom_stack"
	"github.com/stretchr/testify/assert"
)

func isOperator(char uint8) bool {
	return strings.ContainsAny(string(char), "+ & - & * & /")
}

func isOperand(char uint8) bool {
	return char >= '0' && char <= '9'
}

func getOperatorWeight(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/", "%":
		return 2
	}

	return -1
}

func hasHigherPrecedence(operator1 string, operator2 string) bool {
	return getOperatorWeight(operator1) >= getOperatorWeight(operator2)
}

func infixToPostfix(str string) string {
	stack := customstack.NewStack()
	postfix := ""
	length := len(str)

	for index := 0; index < length; index++ {
		char := string(str[index])

		if char == " " {
			continue
		}

		if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			for !stack.IsEmpty() {
				temp, _ := stack.Peek()
				st := temp.(string)

				if st == "(" {
					break
				}

				postfix += st
				stack.Pop()
			}
			stack.Pop()
		} else if !isOperator(str[index]) {
			index_j := index
			number := ""

			for ; index_j < length && isOperand(str[index_j]); index_j++ {
				number = number + string(str[index_j])
			}

			postfix += number
			index = index_j - 1

		} else {
			for !stack.IsEmpty() {
				temp, _ := stack.Peek()
				top := temp.(string)

				if top == "(" || hasHigherPrecedence(top, char) {
					break
				}
				postfix += top
				stack.Pop()
			}

			stack.Push(char)
		}
	}

	for !stack.IsEmpty() {
		temp, _ := stack.Peek()
		top := temp.(string)
		postfix += top
		stack.Pop()
	}

	return strings.TrimSpace(postfix)
}

func TestInfixToPostfix(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(infixToPostfix("(1+1)"), "this show return some expression")
	assert.Equal(infixToPostfix("(2-1)+2"), "21-2+", "this is valid expression")
	assert.Equal(infixToPostfix("(2-1)+(2+1)"), "21-21++", "this is valid expression")
}
