/**
*  Given a string s and an integer k, break up the string into multiple lines such that each line
*  has a length of k or less. You must break it up so that words don't break across lines. Each
*  line has to have the maximum possible amount of words. If there's no way to break the text up,
*  then return null.
*
*  You can assume that there are no spaces at the ends of the string and that there is exactly one
*  space between each word.
*
*  For example, given the string "the quick brown fox jumps over the lazy dog" and k = 10, you should
*  return: ["the quick", "brown fox", "jumps over", "the lazy", "dog"]. No string in the list has a
*  length of more than 10.
* */
package dailycodingproblems

import (
	"strings"
	"testing"
)

type maxSplit struct{}

func (maxSplit) maxSplitWords(str string, k int) []string {
	words, lines := strings.Fields(str), []string{}

	currentLine := words[0]
	for _, word := range words {
		if len(currentLine)+1+len(word) <= k {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	lines = append(lines, currentLine)

	return lines
}

func TestMaxSplitWords(t *testing.T) {
	m := maxSplit{}
	arr := m.maxSplitWords("the quick brown fox jumps over the lazy dog", 10)
	t.Logf("expected 10; got %v\n", arr)
}
