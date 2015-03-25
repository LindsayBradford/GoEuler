// GoEuler - An implementation of Project Euler with Go.
// (c) 2015, Lindsay Bradford

package main

import (
  "goeuler/cmdline"
  "goeuler/config"
  "goeuler/daftlog"
  "goeuler/problem" 
)

func main() {

  var args = new(cmdline.Arguments)
  args.Process()
  
  daftlog.SetFormat(daftlog.FormatTimeOnly)
  
  daftlog.Print(cmdline.GetVersionString())
  daftlog.Print("Project Euler via Go, (c) 2015, Lindsay Bradford.")
  daftlog.Printf(
    "Problems Implemented: (%d/%d)", 
    config.LastImplementedEulerProblemID, 
    config.MaxEulerProblemID,
  )
  daftlog.Print("")
	
  process(
    Problems.BuildProblem(args.Problem),
  )
}

func process(p Problems.Problem) {
  daftlog.Printf("Problem %d: %s", p.GetID(), p.GetTitle())
	
  p.CalculateAnswer()
	
  daftlog.Printf("Answer  %d: %s", p.GetID(), p.GetAnswer())

  if p.IsAnswerVerified() {
    daftlog.Printf("Project Euler has verified answer %d as correct.", p.GetID())
  }
}

