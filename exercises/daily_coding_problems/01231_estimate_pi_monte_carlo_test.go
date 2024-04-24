package dailycodingproblems

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
* The area of a circle is defined as πr^2. Estimate π to 3 decimal places using a Monte Carlo method.
*
* Hint: The basic equation of a circle is x2 + y2 = r2.
* **/

func estimatePi(simulations int) float64 {
	circlePoints := 0
	squarePoints := simulations

	rand.Seed(time.Now().UnixNano())

	for index := 0; index < simulations; index++ {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1

		if ((x * x) + (y * y)) <= 1 {
			circlePoints++
		}
	}

	return float64(circlePoints) / float64(squarePoints) * 4
}

func TestEstimatePi(t *testing.T) {
	estimatedPi := estimatePi(100000000)
	fmt.Println(estimatedPi)
}
