package stacks_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPallindrome(input string) bool {
	for index := 0; index < len(input)/2; index++ {
		if input[index] != input[len(input)-index-1] {
			return false
		}
	}

	return true
}

func TestPallindrome(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(isPallindrome("aba"), true, "it is pallindrome string")
	assert.True(isPallindrome("abba"), "this is a pallindrome string")
	assert.False(isPallindrome("abb"), "this must not be a pallindrome string")
}
