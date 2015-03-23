// GoEuler - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package main

import (
  "Config"
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
  DaftLog.FormattedLogEntry(
    "Problems Implemented: (%d/%d)", 
    Config.LAST_IMPLEMENTED_PROBLEM_ID, 
    Config.MAX_EULER_PROBLEM_ID,
  )
  DaftLog.LogEntry("")
	
  process(
    BuildProblem(args.Problem),
  )
}

func process(thisProblem Problem) {
  DaftLog.FormattedLogEntry("Problem %d: %s", thisProblem.GetID(), thisProblem.GetTitle())
	
  thisProblem.CalculateAnswer()
	
  DaftLog.FormattedLogEntry("Answer  %d: %s", thisProblem.GetID(), thisProblem.GetAnswer())

  if thisProblem.IsAnswerVerified() {
    DaftLog.FormattedLogEntry("Project Euler has verified answer %d as correct.", thisProblem.GetID())
  }
}

