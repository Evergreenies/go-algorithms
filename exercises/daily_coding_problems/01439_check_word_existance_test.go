/*
*
Given a 2D matrix of characters and a target word, write a function that returns whether
the word can be found in the matrix by going left-to-right, or up-to-down.

For example, given the following matrix:

[['F', 'A', 'C', 'I'],

	['O', 'B', 'Q', 'P'],
	['A', 'N', 'O', 'B'],
	['M', 'A', 'S', 'S']]

and the target word 'FOAM', you should return true, since it's the leftmost column.
Similarly, given the target word 'MASS', you should return true, since it's the last row.
*
*/
package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type wordExistenceStrct struct{}

func (wordExistenceStrct) wordExistence(matrix [][]rune, word string) bool {
	wordLength := len(word)
	if wordLength == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	// check each row for the word (horizontal search)
	for row := 0; row < rows; row++ {
		for col := 0; col <= cols-wordLength; col++ {
			if string(matrix[row][col:col+wordLength]) == word {
				return true
			}
		}
	}

	// check each column for word (vertical search)
	for col := 0; col < cols; col++ {
		for row := 0; row <= rows-wordLength; row++ {
			match := true

			for k := 0; k < wordLength; k++ {
				if matrix[row+k][col] != rune(word[k]) {
					match = false
					break
				}
			}

			if match {
				return true
			}
		}
	}

	return false
}

func TestWordExistence(t *testing.T) {
	wrdStrct := wordExistenceStrct{}
	matrix := [][]rune{
		{'F', 'A', 'C', 'I'},
		{'O', 'B', 'Q', 'P'},
		{'A', 'N', 'O', 'B'},
		{'M', 'A', 'S', 'S'},
	}

	assert := assert.New(t)
	assert.Equal(true, wrdStrct.wordExistence(matrix, "FOAM"))
	assert.Equal(true, wrdStrct.wordExistence(matrix, "MASS"))
	assert.Equal(false, wrdStrct.wordExistence(matrix, "FACE"))
}
