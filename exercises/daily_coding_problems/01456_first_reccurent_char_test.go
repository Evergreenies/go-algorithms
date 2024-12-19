/*
* Given a string, return the first recurring character in it, or null if there is no recurring character.
*
* For example, given the string "acbbac", return "b". Given the string "abcdef", return null.
* */
package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type rCStrct struct{}

func (rCStrct) recurrentCharacter(str string) string {
	dict := map[rune]int{}
	for _, char := range str {
		dict[char] += 1
		if dict[char] > 1 {
			return string(char)
		}
	}

	return ""
}

func TestRecurrentCharacter(t *testing.T) {
	assert := assert.New(t)
	rc := rCStrct{}

	assert.Equal("b", rc.recurrentCharacter("acbbac"))
	assert.Equal("", rc.recurrentCharacter("acbdef"))
}
