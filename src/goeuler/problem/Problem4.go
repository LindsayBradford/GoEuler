// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

import (
	"log"
	"fmt"
)

type problem4 struct {
	problemBase
}

func (proble *problem4) Initialise(logger *log.Logger) {
	proble.logger = logger
	proble.id = 4
	proble.title = "Largest palindrome product"
	proble.description = "A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99." +
		"Find the largest palindrome made from the product of two 3-digit numbers."
	proble.answer = "Answer not calculated yet."

}

func (problem *problem4) CalculateAnswer() {
	const largestThreeDigitNumber uint = 999
	const smallestThreeDigitNumber uint = 100

	type numberIdentifier uint

	const (
		ONE numberIdentifier = 1
		TWO numberIdentifier = 2
	)

	largestFoundPalindrome := uint(0)
	numbers                := make(map[numberIdentifier]uint)

	for numbers[ONE] = largestThreeDigitNumber; numbers[ONE] >= smallestThreeDigitNumber; numbers[ONE]-- {
		for numbers[TWO] = largestThreeDigitNumber; numbers[TWO] >= smallestThreeDigitNumber; numbers[TWO]-- {
			productOfNumbers := numbers[ONE] * numbers[TWO]
			if isPalindrome(productOfNumbers) && productOfNumbers > largestFoundPalindrome {
				largestFoundPalindrome = productOfNumbers
				problem.logger.Printf("Problem %d - New largest palindrome found: %d (%d * %d)\n",
					problem.id, largestFoundPalindrome, numbers[ONE], numbers[TWO])
			}
		}
	}
	problem.answer = fmt.Sprintf("Problem %d, Answer: %d\n",problem.id, largestFoundPalindrome)
	problem.answerVerified = true
}

func isPalindrome(number uint) bool {
	var digitsAsSlice []uint
	numberWithDigitsToHarvest := number

	for digitsAllHarvested := false; !digitsAllHarvested; digitsAllHarvested = checkIfAllDigitsHarvested(numberWithDigitsToHarvest) {
		newDigit := numberWithDigitsToHarvest % 10
		numberWithDigitsToHarvest = numberWithDigitsToHarvest / 10
		digitsAsSlice = append(digitsAsSlice, newDigit)
	}

	for firstIndex, lastIndex := 0, len(digitsAsSlice)-1; firstIndex < lastIndex; firstIndex++ {
		if digitsAsSlice[firstIndex] != digitsAsSlice[lastIndex] {
			return false
		}
		lastIndex--
	}
	return true
}

func checkIfAllDigitsHarvested(numberWithDigitsToHarvest uint) bool {
	return numberWithDigitsToHarvest == 0
}
