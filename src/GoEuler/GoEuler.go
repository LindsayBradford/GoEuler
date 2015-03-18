// GoEuler - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package main

import (
  "CommandLine"
	"DaftLog"
  . "GoEuler/Problems" 
)

func main() {

  var args = new(CommandLine.Arguments)
  args.Process()
  
  DaftLog.SetFormat(DaftLog.TIME_ONLY_FORMAT)
  
	DaftLog.LogEntry(CommandLine.GetVersionString())
	DaftLog.LogEntry("Project Euler via Go, (c) 2015, Lindsay Bradford.")
	
	process(
  	BuildProblem(args.ProblemID),
	)
}

func process(thisProblem Problem) {
	thisProblem.Initialise()

	DaftLog.FormattedLogEntry("Problem %d: %s", thisProblem.GetID(), thisProblem.GetTitle())
	
	thisProblem.CalculateAnswer()
	
	DaftLog.FormattedLogEntry("Answer  %d: %s", thisProblem.GetID(), thisProblem.GetAnswer())
	
	if thisProblem.IsVerified() {
		DaftLog.FormattedLogEntry("Project Euler has verified answer %d as correct.", thisProblem.GetID())
	}
}

