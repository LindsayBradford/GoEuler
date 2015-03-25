// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

// BuildProblem creates a record that implements the
// Problem interface from the supplied ID number.
func BuildProblem(problemID uint) Problem {
	var newProblem Problem
	switch problemID {
	case 1:
		newProblem = new(problem1)
	case 2:
		newProblem = new(problem2)
	case 3:
		newProblem = new(problem3)
	default:
		newProblem = new(problemBase)
	}
	newProblem.Initialise()
	return newProblem
}
