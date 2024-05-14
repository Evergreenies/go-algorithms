package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
*  Given a string, determine whether any permutation of it is a palindrome.
*
*  For example, carrace should return true, since it can be rearranged to form racecar,
*  which is a palindrome. daily should return false, since there's no rearrangement that
*  can form a palindrome.
* **/

func canPermuteToPallindrome(str string) bool {
	charCount := make(map[rune]int, 0)
	for _, char := range str {
		charCount[char]++
	}

	oddCount := 0
	for _, count := range charCount {
		oddCount += (count % 2)
	}

	return oddCount <= 1
}

func TestCanPermuteToPallindrome(t *testing.T) {
	assert := assert.New(t)
	assert.True(canPermuteToPallindrome("carrace"), "this must be a pollindromic combination")
	assert.False(canPermuteToPallindrome("daily"), "this must be a non pollindromic combination")
}
