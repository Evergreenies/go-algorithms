/*
Mastermind is a two-player game in which the first player attempts to guess the secret
code of the second. In this version, the code may be any six-digit number with all distinct digits.

Each turn the first player guesses some number, and the second player responds by saying
how many digits in this number correctly matched their location in the secret code.
For example, if the secret code were 123456, then a guess of 175286 would score two,
since 1 and 6 were correctly placed.

Write an algorithm which, given a sequence of guesses and their scores, determines whether
there exists some secret code that could have produced them.

For example, for the following scores you should return True, since they correspond to the secret
code 123456:

{175286: 2, 293416: 3, 654321: 0}
However, it is impossible for any key to result in the following scores, so in this case you should
return False:

{123456: 4, 345678: 4, 567890: 4}
* */

package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sixDigitSecretGame struct{}

// helper function to calculate the score for the guess given a secret code
func (sixDigitSecretGame) calculateScore(secret, guess string) int {
	score := 0
	for index := 0; index < 6; index++ {
		if secret[index] == guess[index] {
			score++
		}
	}

	return score
}

// helper function to check if a string has a all unique digits
func (sixDigitSecretGame) isUniqueDigit(str string) bool {
	seen := make(map[rune]bool)
	for _, ch := range str {
		if seen[ch] {
			return false
		}

		seen[ch] = true
	}

	return true
}

// function to generate all six digits numbers with distinct digits
func (sx sixDigitSecretGame) generateAllCandidates() []string {
	candidates := []string{}
	for index := 0; index < 1000000; index++ {
		str := fmt.Sprintf("%06d", index)
		if sx.isUniqueDigit(str) {
			candidates = append(candidates, str)
		}
	}

	return candidates
}

// isValidSecreCode determines if there is a valid secret code for given guesses
func (sx sixDigitSecretGame) isValidSecreCode(guesses map[string]int) bool {
	candidates := sx.generateAllCandidates()
	for _, secret := range candidates {
		valid := true
		for guess, score := range guesses {
			if sx.calculateScore(secret, guess) != score {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}

func TestSixDigitsGuess(t *testing.T) {
	sx := sixDigitSecretGame{}
	assert := assert.New(t)

	// Example 1
	guesses1 := map[string]int{"175286": 2, "293416": 3, "654321": 0}
	assert.True(sx.isValidSecreCode(guesses1))

	// Example 2
	guesses2 := map[string]int{"123456": 4, "345678": 4, "567890": 4}
	assert.False(sx.isValidSecreCode(guesses2))
}
