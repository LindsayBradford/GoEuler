// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

func BuildProblem(problemID uint) Problem {
  var newProblem Problem
  switch problemID {
    case 1:
      newProblem = new(Problem1) 
    case 2:
      newProblem = new(Problem2) 
    default:
      newProblem = new(ProblemBase)
  }
  newProblem.Initialise()
  return newProblem
}
