// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

// Package Problems contains a Euler Problem interface and
// the structures that implement it.
package Problems

import (
	"log"
)

// Problem is an interface that defines those functions that
// allow calling code to successfully interact  with a
// Euler problem withouth having to know the problem specifics.
type Problem interface {
	Initialise(*log.Logger)

	GetID() ProblemId
	GetTitle() string
	GetDescription() string
	GetAnswer() string

	CalculateAnswer()
	IsAnswerVerified() bool
}

// ProblemBase defines a structure that specific Euler
// problem structures are expected to reuse via type extension.
type problemBase struct {
	Problem
	logger         *log.Logger
	id             ProblemId
	title          string
	description    string
	answer         string
	answerVerified bool
}

// The function Initialise() will initialise the content of a ProblemBase
// structure with reasonable, generic content.
func (base *problemBase) Initialise(logger *log.Logger) {
	base.logger = logger
	base.id = 0
	base.title = "No problem title supplied."
	base.description = "No problem description supplied."
	base.answer = "No answer generated."
	base.answerVerified = false
}

// GetID returns the id number (as supplied on the Euler project) of the
// relevant Euler problem.
func (base *problemBase) GetID() ProblemId {
	return base.id
}

// GetTitle returns the title text (as supplied on the Euler project) of the
// relevant Euler problem.
func (base *problemBase) GetTitle() string {
	return base.title
}

// GetTitle returns the long-form description (as supplied on the Euler project) of the
// relevant Euler problem.
func (base *problemBase) GetDescription() string {
	return base.description
}

// CalculateAnswer generates an answer to the relavante Euler problem. Unless confirmed by
// the Euler project, the answer is not necessarilly correct.
func (base *problemBase) CalculateAnswer() {
	base.answer = "Some base answer supplied."
}

// GetAnswer returns the answer generated in attempting to solve the relevante Euler problem.
// The answer is undefined unless the CalculateAnswer function has already been called.
func (base *problemBase) GetAnswer() string {
	return base.answer
}

// InsAnswerVerified returns true only if the result generated from calling CalculateAnswer has
// been verified  by the Euler project as being correct, false otherwise.
func (base *problemBase) IsAnswerVerified() bool {
	return base.answerVerified
}
