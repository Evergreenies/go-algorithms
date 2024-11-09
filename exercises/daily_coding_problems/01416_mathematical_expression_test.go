/**
* Given a string consisting of parentheses, single digits, and positive and negative signs,
* convert the string into a mathematical expression to obtain the answer.
*
* Don't use eval or a similar built-in parser.
*
* For example, given '-1 + (2 + 3)', you should return 4.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mathematicalExpression struct{}

func (mathematicalExpression) mathematicalExpressionEvaluation(expression string) int {
	stack, result, num, sign := []int{}, 0, 0, 1

	for index := 0; index < len(expression); index++ {
		char := expression[index]

		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else if char == '+' {
			result += sign * num
			num, sign = 0, 1
		} else if char == '-' {
			result += sign * num
			num, sign = 0, -1
		} else if char == '(' {
			stack = append(stack, result)
			stack = append(stack, sign)
			result, sign = 0, 1
		} else if char == ')' {
			result += sign * num
			num = 0

			sign = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			prevResult := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result = prevResult + sign*result
		}
	}

	result += sign * num

	return result
}

func TestMathematicalExpression(t *testing.T) {
	mt := mathematicalExpression{}
	result := mt.mathematicalExpressionEvaluation("-1+(2+3)")
	t.Logf("sum of -1+(2+3) = %v \n", result)

	assert := assert.New(t)
	assert.Equal(4, result, fmt.Sprintf("expected 4; got %v", result))
}
