package dailycodingproblems

/**
 * Given a string, find the length of the smallest window that contains
 * every distinct character. Characters may appear more than once in the window.
 * For example, given "jiujitsu", you should return 5, corresponding to
 * the final five letters.
**/
import (
	"testing"
)

func smallestWindow(str string) int {
	charCount := make(map[rune]int)
	uniqueChars := len(getUniquChars(str))
	left, right := 0, 0
	minLength := len(str)

	for right < len(str) {
		currentChar := rune(str[right])
		charCount[currentChar]++

		for len(charCount) == uniqueChars && all(charCount) {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}

			leftChar := rune(str[left])
			charCount[leftChar]--
			if charCount[leftChar] == 0 {
				delete(charCount, leftChar)
			}

			left++
		}

		right++
	}

	return minLength
}

func getUniquChars(str string) map[rune]bool {
	unique := make(map[rune]bool)

	for _, char := range str {
		unique[rune(char)] = true
	}

	return unique
}

func all(mp map[rune]int) bool {
	for _, value := range mp {
		if value <= 0 {
			return false
		}
	}

	return true
}

func TestSmallestWindow(t *testing.T) {
	tests := []struct {
		str           string
		expectedCount int
	}{
		{str: "jiujitsu", expectedCount: 5},
		{str: "a", expectedCount: 1},
		{str: "", expectedCount: 0},
		{str: "aabbccabaabbccc", expectedCount: 3},
	}

	for index, test := range tests {
		if actualCount := smallestWindow(test.str); actualCount != test.expectedCount {
			t.Errorf("FAILED: Test case: %v\n\tActual: %v\n\tExpected: %v\n", index, actualCount, test.expectedCount)
		} else {
			t.Logf("PASSED: Test case %v\n\tActual: %v\n\tExpected: %v\n", index, actualCount, test.expectedCount)
		}
	}
}
