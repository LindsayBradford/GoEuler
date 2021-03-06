// GoEuler - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package main

import (
	"goeuler/cmdline"
	"goeuler/config"
	. "goeuler/problem"
	"log"
	"os"
	"sync"
	"time"
)

var (
	args     = new(cmdline.Arguments)
	builder  = new(ProblemBuilder)
	basicLog = log.New(os.Stdout, "", 0)
	timedLog = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)
)

func main() {
	reportVersionAndProblemsImplemented()

	args.Process()
	minProblemId, maxProblemId := resolveProblemRange()

	startTime := time.Now()
	if args.Concurrent {
		answerConcurrently(minProblemId, maxProblemId)
	} else {
		answerSequentially(minProblemId, maxProblemId)
	}
	endTime := time.Now()

	timedLog.Printf("Problems (%d through %d) took %v to answer\n", minProblemId, maxProblemId, endTime.Sub(startTime))
}

func answerConcurrently(minProblemId uint, maxProblemId uint) {
	var wg sync.WaitGroup

	for problemId := minProblemId; problemId <= maxProblemId; problemId++ {
		wg.Add(1)

		go func(problem uint) {
			defer wg.Done()
			answer(buildProblem(problem))
		}(problemId)
	}

	wg.Wait()
}

func answerSequentially(minProblemId uint, maxProblemId uint) {
	for problemId := minProblemId; problemId <= maxProblemId; problemId++ {
		answer(buildProblem(problemId))
	}
}

func buildProblem(problem uint) Problem {
	return builder.
		OfProblem(problem).
		WithLogger(timedLog).
		Build()
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
		timedLog.Printf("Problem %d is verified as correct by Project Euler.", p.GetID())
	}
}
