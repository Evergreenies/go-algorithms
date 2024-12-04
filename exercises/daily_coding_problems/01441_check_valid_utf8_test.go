/*
UTF-8 is a character encoding that maps each symbol to one, two, three, or four bytes.

For example, the Euro sign, €, corresponds to the three bytes 11100010 10000010 10101100.
The rules for mapping characters are as follows:

For a single-byte character, the first bit must be zero.
For an n-byte character, the first byte starts with n ones and a zero. The other n - 1 bytes
all start with 10.
Visually, this can be represented as follows.

	Bytes   |           Byte format

-----------------------------------------------

	1     | 0xxxxxxx
	2     | 110xxxxx 10xxxxxx
	3     | 1110xxxx 10xxxxxx 10xxxxxx
	4     | 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx

Write a program that takes in an array of integers representing byte values, and returns whether
it is a valid UTF-8 encoding.
*/

package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type validateUTF8Strct struct{}

func (validateUTF8Strct) isValidUTF8Encoding(data []int) bool {
	length := len(data)
	for index := 0; index < length; {
		// check number of bytes for current character
		if data[index]>>7 == 0 {
			// 1-byte character
			index++
		} else if data[index]>>5 == 0b110 {
			// 2-bytes characters
			if index+1 >= length || data[index+1]>>6 != 0b10 {
				return false
			}

			index += 2
		} else if data[index]>>4 == 0b1110 {
			// 3-bytes characters
			if index+2 >= length || data[index+1]>>6 != 0b10 || data[index+2]>>6 != 0b10 {
				return false
			}

			index += 3
		} else if data[index]>>3 == 0b11110 {
			if index+3 >= length || data[index+1]>>6 != 0b10 || data[index+2]>>6 != 0b10 || data[index+3]>>6 != 0b10 {
				return false
			}

			index += 4
		} else {
			return false
		}
	}

	return true
}

func TestIsValidUTF8Encoding(t *testing.T) {
	assert := assert.New(t)

	data1 := []int{226, 130, 172}
	data2 := []int{235, 140, 4}

	utf := validateUTF8Strct{}
	assert.True(utf.isValidUTF8Encoding(data1))
	assert.False(utf.isValidUTF8Encoding(data2))
}
