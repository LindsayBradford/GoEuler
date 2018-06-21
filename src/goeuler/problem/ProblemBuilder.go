// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

import (
	"log"
)

// BuildProblem creates a record that implements the
// Problem interface from the supplied ID number.

type ProblemId uint

type ProblemBuilder struct {
	logger *log.Logger
	problemId ProblemId
}

func (pb *ProblemBuilder) OfProblem(problemId uint) *ProblemBuilder {
	pb.problemId = ProblemId(problemId)
	return pb
}

func (pb *ProblemBuilder) WithLogger(logger *log.Logger) *ProblemBuilder {
	pb.logger = logger
	return pb
}

func (pb *ProblemBuilder) Build() Problem {
	var newProblem Problem
	switch pb.problemId {
	case 1:
		newProblem = new(problem1)
	case 2:
		newProblem = new(problem2)
	case 3:
		newProblem = new(problem3)
	default:
		newProblem = new(problemBase)
	}
	newProblem.Initialise(pb.logger)
	return newProblem
}
