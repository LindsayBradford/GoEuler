// GoEuler Problems package - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package Problems

type Problem interface {
  Initialise()

  GetID() uint
  GetTitle() string
  GetDescription() string
  GetAnswer() string
  
  CalculateAnswer()
  IsAnswerVerified() bool
} 

type ProblemBase struct {
  Problem
  id              uint
  title           string
  description     string
  answer          string
  answerVerified  bool 
}

func (this *ProblemBase) Initialise() {
  this.id = 0
  this.title = "No problem title supplied."
  this.description = "No problem description supplied."
  this.answer = "No answer generated."
  this.answerVerified = false
}

func (this *ProblemBase) GetID() uint {
  return this.id
}

func (this *ProblemBase) GetTitle() string {
  return this.title
}

func (this *ProblemBase) GetDescription() string {
  return this.description
}

func (this *ProblemBase) GetAnswer() string {
  return this.answer
}

func (this *ProblemBase) CalculateAnswer() {
  this.answer = "Some answer supplied."
}

func (this *ProblemBase) IsAnswerVerified() bool {
  return this.answerVerified	
}