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
  Licence bool
  Problem uint
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

 flag.BoolVar(
    &args.Licence, 
    "Licence", 
    false, 
    "Prints the copyright licence this software is released under.",
  )

  
  flag.UintVar(
    &args.Problem,
    "Problem",
    0,
    "Numeric Id of problem as per Projecteuler.net",
  )
  
  flag.Usage = usageMessage
  
  flag.Parse()
}

func usageMessage() {
  fmt.Printf("Help for %s\n", GetVersionString())
  fmt.Println("  --Help              Prints this help message.")
  fmt.Println("  --Version           Prints the version number of this utility.")
  fmt.Println("  --Licence           Prints the copyright licence this utility is release under.")
  fmt.Println("  --Problem <number>  Specifies problem ID to evaluate.")
  os.Exit(0)
}

func (args *Arguments) process() {

  if args.Licence == true {
    fmt.Println(
      getLicenceString(),
    )
    os.Exit(0)
  }

  if args.Version == true {
    fmt.Println(
      GetVersionString(),
    )
    os.Exit(0)
  }

  if args.Problem == 0 {
    usageMessage()
  }

  var mustExitWithError = false

  if args.Problem < Config.MIN_EULER_PROBLEM_ID || args.Problem > Config.MAX_EULER_PROBLEM_ID {
     fmt.Println("Error: Invalid Euler Project ID specified.")
     mustExitWithError = true
  }

  if args.Problem > Config.LAST_IMPLEMENTED_PROBLEM_ID {
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

func getLicenceString() string {
  licence := 
    "\nCopyright (c) 2015, Lindsay Bradford\n" + 
    "All rights reserved.\n\n" + 
    "Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:\n\n" + 
    "1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.\n\n" +
    "2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.\n\n" + 
    "3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.\n\n" +
    "THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS \"AS IS\" AND ANY EXPRESS OR IMPLIED " +
    "WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A " +
    "PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY " +
    "DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, " +
    "PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER " +
    "CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR " +
     "OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.\n"
     
  return licence
}