/**
* Find the maximum of two numbers without using any if-else statements, branching, or direct comparisons.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"
)

type comparisons struct{}

func (comparisons) comparison(a, b int) int {
	diff := a - b
	sign := (diff >> 31) & 1
	return a*(1-sign) + b*sign
}

func TestComparision(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{10, 10, 10},
		{10, 20, 20},
	}

	c := comparisons{}
	for index, test := range tests {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			result := c.comparison(test.a, test.b)
			if result != test.expected {
				t.Errorf("TEST %v: expected %v; got %v\n", t.Name(), test.expected, result)
			} else {
				t.Logf("TEST %v: PASSED\n", t.Name())
			}
		})
	}
}
