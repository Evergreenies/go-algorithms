package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Gray code is a binary code where each successive value differ in only one bit, as well as when
* wrapping around. Gray code is common in hardware so that we don't see temporary spurious values
* during transitions.
*
* Given a number of bits n, generate a possible gray code for it.
*
* For example, for n = 2, one gray code would be [00, 01, 11, 10].
* **/

func generateGrayCode(n int) []string {
	if n == 0 {
		return []string{}
	}

	if n == 1 {
		return []string{"0", "1"}
	}

	prevGrayCode := generateGrayCode(n - 1)
	grayCode := []string{}
	for _, item := range prevGrayCode {
		grayCode = append(grayCode, "0"+item)
	}

	for index := len(prevGrayCode) - 1; index >= 0; index-- {
		grayCode = append(grayCode, "1"+prevGrayCode[index])
	}

	return grayCode
}

func TestGenerateGrayCode(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(generateGrayCode(2), []string{"00", "01", "11", "10"}, "it must return 4 elements in array")
}
