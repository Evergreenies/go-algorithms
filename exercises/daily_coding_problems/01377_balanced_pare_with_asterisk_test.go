/*
* You're given a string consisting solely of (, ), and *. * can represent either a (, ), or an empty string.
* Determine whether the parentheses are balanced.
*
* For example, (()* and (*) are balanced. )*( is not balanced.
* */

package dailycodingproblems

import (
	"testing"
)

type balancedPareWithAstersik struct{}

func (balancedPareWithAstersik) check_valid_string(str string) bool {
	low := 0  // min possible open parentheses count
	high := 0 // max possible open parentheses count

	for _, char := range str {
		switch char {
		case '(':
			low++
			high++
		case ')':
			low--
			high--
		default:
			low--
			high++
		}

		if high < 0 {
			return false
		}

		low = max(low, 0)

	}

	return low == 0
}

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"(()*", true},
		{"(*)", true},
		{")*(", false},
		{"((*)", true},
		{"((**)(", false},
		{"", true},
		{"**", true},
		{"(())", true},
		{"((*))", true},
		{"(*)))", false},
	}

	bp := balancedPareWithAstersik{}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := bp.check_valid_string(test.input)
			if result != test.expected {
				t.Errorf("expected %v; got %v", test.expected, result)
			}
		})
	}
}
