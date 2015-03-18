// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

import (
  "fmt"
  "math"
)

type Problem1 struct {
  ProblemBase
}

func (this *Problem1) Initialise() {
  this.id = 1
  this.title = "Multiples of 3 and 5."
  this.description = "If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. " + 
                     " The sum of these multiples is 23." + 
                     " Find the sum of all the multiples of 3 or 5 below 1000."
  this.answer = "Answer not calculated yet."                    
}

func (this *Problem1) CalculateAnswer() {
  var factorSum uint = 0
  
  for i := 1; i < 1000; i++ {
    var iAsFloat = float64(i)
    if math.Mod(iAsFloat,3) == 0 || math.Mod(iAsFloat,5) == 0 {
      factorSum = factorSum + uint(i)
    }
  }
  
  this.answer = fmt.Sprintf("%d", factorSum)
  this.verified = true
}
