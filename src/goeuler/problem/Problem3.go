// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

import (
	"fmt"
	"log"
	"math"
)

type problem3 struct {
	problemBase
}

func (p3 *problem3) Initialise(logger *log.Logger) {
	p3.logger = logger
	p3.id = 3
	p3.title = "Largest prime factor."
	p3.description = " The prime factors of 13195 are 5, 7, 13 and 29.\n" +
		" What is the largest prime factor of the number 600851475143 ?"
	p3.answer = "Answer not calculated yet."
}

func (p3 *problem3) CalculateAnswer() {

	const BaseNumber = 600851475143

	var largestPossiblePrimeFactor = uint(math.Sqrt(BaseNumber)) // largest factor of a number is its square root.

	// A prime is never even, find the largest odd to start on.
	if math.Mod(float64(largestPossiblePrimeFactor), 2) == 0 {
		largestPossiblePrimeFactor -= 1
	}

	p3.logger.Printf("Problem 3 - Largest possible odd factor of %d = %d", BaseNumber, largestPossiblePrimeFactor)

	var primeNotFound = true

	for primeNotFound {
		if math.Mod(BaseNumber, float64(largestPossiblePrimeFactor)) == 0 { // found an odd factor.

			var resultOfPrimeTest string

			if isOddPrime(float64(largestPossiblePrimeFactor)) { // is this odd factor a prime?
				resultOfPrimeTest = fmt.Sprintf("and %d is a prime!", largestPossiblePrimeFactor)
				primeNotFound = false // signal that that we're done searching.
			} else {
				resultOfPrimeTest = fmt.Sprintf("but %d is not a prime!", largestPossiblePrimeFactor)
			}

			p3.logger.Printf("Problem 3 - Found odd factor at %d, %s", largestPossiblePrimeFactor, resultOfPrimeTest)
		}

		if primeNotFound {
			largestPossiblePrimeFactor -= 2 // skip to next smaller odd number as possible factor.
		}
	}
	p3.answer = fmt.Sprintf("%d", largestPossiblePrimeFactor)

	p3.answerVerified = true
}

// isOddPrime tests to see if the odd number (assumed) provided is also a prime.
func isOddPrime(testNumber float64) bool {

	var largestFactor = int(math.Sqrt(testNumber)) // largest factor of a number is its square root.

	for i := 3; i < largestFactor; i++ {
		if math.Mod(testNumber, float64(i)) == 0 {
			return false
		}
	}
	return true
}
