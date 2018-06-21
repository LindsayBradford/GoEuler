// GoEuler - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package main

import (
	"goeuler/cmdline"
	"goeuler/config"
	. "goeuler/problem"
	"log"
	"os"
)

var (
	args     = new(cmdline.Arguments)
	builder  = new(ProblemBuilder)
	basicLog = log.New(os.Stdout, "", 0)
	timedLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

func main() {
	args.Process()
	reportVersionAndProblemsImplemented()

	minProblemId, maxProblemId := resolveProblemRange()

	for problemId := minProblemId; problemId <= maxProblemId; problemId++ {
		problem := builder.
			OfProblem(problemId).
			WithLogger(timedLog).
			Build()

		  answer(problem)
	}
}

func reportVersionAndProblemsImplemented() {
	basicLog.Print(cmdline.GetVersionString())
	basicLog.Print("Project Euler via Go, (c) 2015, Lindsay Bradford.")
	basicLog.Printf(
		"Problems Implemented: (%d/%d)",
		config.LastImplementedEulerProblemID,
		config.MaxEulerProblemID,
	)
	basicLog.Println("")
}

func resolveProblemRange() (nimProblemId uint, maxProblemId uint) {
	if args.AllProblems {
		return 1, config.LastImplementedEulerProblemID
	}
	return args.Problem, args.Problem
}

func answer(p Problem) {
	timedLog.Printf("Problem %d: %s", p.GetID(), p.GetTitle())

	p.CalculateAnswer()

	timedLog.Printf("Problem %d answer: %s", p.GetID(), p.GetAnswer())

	if p.IsAnswerVerified() {
		timedLog.Printf("Problem %d as verified as correct by Project Euler.", p.GetID())
	}
}
