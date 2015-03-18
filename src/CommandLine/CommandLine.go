// GoEuler.go - experimental implementation
// Command-Line Argument processing pakage
// (c) 2014, Lindsay Bradford

package CommandLine

import (
  "Config"
  "flag"
  "fmt"
  "os"
)

type Arguments struct {
  Version bool
  ProblemID uint
}

func (args *Arguments) Process() {
  args.define()
  args.process()
}

func (args *Arguments) define() {

 flag.BoolVar(
    &args.Version, 
    "Version", 
    false, 
    "Prints the version number of this utility and exits.",
  )
  
  flag.UintVar(
    &args.ProblemID,
    "ProblemID",
    0,
    "ID of problem as per Projecteuler.net",
  )
  
  flag.Usage = usageMessage
  
  flag.Parse()
}

func usageMessage() {
  fmt.Printf("Help for %s\n", GetVersionString())
  fmt.Println("  --Help                Prints this help message.")
  fmt.Println("  --Version             Prints the version number of this utility.")
  fmt.Println("  --ProblemID <number>  Specifies problem ID to evaluate.")
  os.Exit(0)
}

func (args *Arguments) process() {

  if args.Version == true {
    fmt.Println(
      GetVersionString(),
    )
    os.Exit(0)
  }

  var mustExitWithError = false

  if args.ProblemID < Config.MIN_EULER_PROBLEM_ID || args.ProblemID > Config.MAX_EULER_PROBLEM_ID {
     fmt.Println("Error: Invalid Euler Project ID specified.")
     mustExitWithError = true
  }

  if args.ProblemID > Config.LAST_IMPLEMENTED_PROBLEM_ID {
     fmt.Printf("Error: Project ID specified has not been implemented yet. Currently at ID: %d\n", Config.LAST_IMPLEMENTED_PROBLEM_ID)
     mustExitWithError = true
  }

  if mustExitWithError {
    os.Exit(1)
  }
}

func GetVersionString() string {
  return fmt.Sprintf("%s version %s", os.Args[0], Config.VERSION)
}