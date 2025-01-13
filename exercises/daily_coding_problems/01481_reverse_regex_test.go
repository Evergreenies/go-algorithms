/*
* Given a string and a set of delimiters, reverse the words in the string
* while maintaining the relative order of the delimiters.
*
* For example, given "hello/world:here", return "here/world:hello"
*
* Follow-up: Does your solution work for the following cases: "hello/world:here/", "hello//world:here"
* */

package dailycodingproblems

import (
	"regexp"
	"strings"
	"testing"
)

type reverseStringStrct struct{}

func (reverseStringStrct) reverseWords(s, delimiters string) string {
	pattern := "[" + regexp.QuoteMeta(delimiters) + "]+"
	re := regexp.MustCompile(pattern)

	words := re.Split(s, -1)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	var result strings.Builder
	delimiterList := re.FindAllString(s, -1)
	for i := 0; i < len(words); i++ {
		if i > 0 && i-1 < len(delimiterList) {
			result.WriteString(delimiterList[i-1])
		}

		result.WriteString(words[i])
	}

	if len(delimiterList) == len(words) {
		result.WriteString(delimiterList[len(delimiterList)-1])
	}

	return result.String()
}

func TestReverseWords(t *testing.T) {
	rs := reverseStringStrct{}
	t.Log(rs.reverseWords("hello/world:here", "/:"))
	t.Log(rs.reverseWords("hello/world:here/", "/:"))
	t.Log(rs.reverseWords("hello//world:here", "/:"))
}
