// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

import (
	"fmt"
	"log"
	"math"
)

type problem1 struct {
	problemBase
}

func (p1 *problem1) Initialise(logger *log.Logger) {
	p1.logger = logger
	p1.id = 1
	p1.title = "Multiples of 3 and 5."
	p1.description = "If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. " +
		" The sum of these multiples is 23." +
		" Find the sum of all the multiples of 3 or 5 below 1000."
	p1.answer = "Answer not calculated yet."
}

func (p1 *problem1) CalculateAnswer() {
	var sumOfMultiples uint

	for i := 3; i < 1000; i++ {
		var iAsFloat = float64(i)
		if math.Mod(iAsFloat, 3) == 0 || math.Mod(iAsFloat, 5) == 0 {
			sumOfMultiples = sumOfMultiples + uint(i)
		}
	}

	p1.answer = fmt.Sprintf("%d", sumOfMultiples)
	p1.answerVerified = true
}
