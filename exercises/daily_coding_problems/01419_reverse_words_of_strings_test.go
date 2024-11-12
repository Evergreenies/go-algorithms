/**
* Given a string of words delimited by spaces, reverse the words in string. For example,
* given "hello world here", return "here world hello"
*
* Follow-up: given a mutable string representation, can you perform this operation in-place?
* */
package dailycodingproblems

import (
	"testing"
)

type (
	reverseWordsPfStringsStr struct{}
	reverseWordsPfStringsInf interface {
		reverse(str []byte, start, end int)
		reverseWords(str []byte) string
	}
)

func (r reverseWordsPfStringsStr) reverseWords(str []byte) string {
	length := len(str)
	r.reverse(str, 0, length-1)

	start := 0
	for index := 0; index < length; index++ {
		if str[index] == ' ' || index == length {
			r.reverse(str, start, index-1)
			start = index + 1
		}
	}

	return string(str)
}

func (r reverseWordsPfStringsStr) reverse(str []byte, start, end int) {
	for start < end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}
}

func TestReverseWordsOfString(t *testing.T) {
	str := "hello world here"
	var rInf reverseWordsPfStringsInf
	rStr := reverseWordsPfStringsStr{}

	rInf = rStr
	str1 := rInf.reverseWords([]byte(str))
	t.Logf("STRING: %v\n", str1)
}
