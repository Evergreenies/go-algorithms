/*
* Write a program that checks whether an integer is a palindrome.
*
* For example, 121 is a palindrome, as well as 888. 678 is not a palindrome.
*
* Do not convert the integer into a string.
* */

package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPallindrom(num int) bool {
	if num < 0 {
		return false
	}

	original := num
	reversedNum := 0

	for num > 0 {
		rem := num % 10
		reversedNum = reversedNum*10 + rem
		num = num / 10
	}

	return original == reversedNum
}

func TestIsPallindrom(t *testing.T) {
	assert := assert.New(t)

	assert.True(isPallindrom(121))
	assert.False(isPallindrom(123))
}
