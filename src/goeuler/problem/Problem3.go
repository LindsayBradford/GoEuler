// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

import (
//  "fmt"
//  "math"
)

type problem3 struct {
	problemBase
}

func (p3 *problem3) Initialise() {
	p3.id = 3
	p3.title = "Largest prime factor."
	p3.description = " The prime factors of 13195 are 5, 7, 13 and 29.\n" +
		" What is the largest prime factor of the number 600851475143 ?"
	p3.answer = "Answer not calculated yet."
}

func (p3 *problem3) CalculateAnswer() {

	// this.answer = fmt.Sprintf("%d", sumOfEvenFibonacciNumbers)
	// this.answerVerified = true
}
