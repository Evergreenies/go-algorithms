package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Implement division of two positive integers without using the division, multiplication, or
* modulus operators. Return the quotient as an integer, ignoring the remainder.
* **/

func divisionWithoutOperator(dividend, divisor uint) int {
	if divisor == 0 {
		panic("Zero division error.")
	}

	if dividend == 0 {
		return 0
	}

	quotient := 0
	for dividend >= divisor {
		dividend = dividend - divisor
		quotient++
	}

	return quotient
}

func TestDivisionWithoutOperator(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(divisionWithoutOperator(15, 5), 3, "15/5 = 3")
	assert.Equal(divisionWithoutOperator(17, 5), 3, "17/5 = 3")
	assert.Equal(divisionWithoutOperator(5, 5), 1, "5/5 = 1")
	assert.Equal(divisionWithoutOperator(4, 5), 0, "4/5 = 0")
}
