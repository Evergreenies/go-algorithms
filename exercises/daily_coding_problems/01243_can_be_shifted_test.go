package dailycodingproblems

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given two strings A and B, return whether or not A can be shifted some number of times to get B.
*
* For example, if A is abcde and B is cdeab, return true. If A is abc and B is acb, return false.
* **/

func canBeShifted(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	return strings.Contains(str1+str1, str2)
}

func TestCanBeShifted(t *testing.T) {
	assert := assert.New(t)
	assert.True(canBeShifted("abcde", "cdeab"), "this can be shifted")
	assert.False(canBeShifted("abc", "acb"), "this cannot be shifted")
}
