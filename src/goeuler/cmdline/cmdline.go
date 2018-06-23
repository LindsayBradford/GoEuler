// GoEuler.go - experimental implementation
// Command-Line Argument processing package
// (c) 2014, Lindsay Bradford

// Package cmdline offers command-line processing for the goeuler code-base.
package cmdline

import (
	"flag"
	"fmt"
	"os"

	"goeuler/config"
)

// Arguments is a struct holding the arguments received from the command-line when goeuler is run.
type Arguments struct {
	Version     bool
	Licence     bool
	Problem     uint
	AllProblems bool
	Concurrent  bool
}

// Process defines the command-line arguments applicable to goeuler, then processes them for usage by the application.
func (args *Arguments) Process() {
	args.define()
	args.process()
}

// define sets up the command-line flags that goeuler recognises via the flag package, and parses them into args.
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

	flag.BoolVar(
		&args.AllProblems,
		"AllProblems",
		false,
		"Answer all problems for which there is code.",
	)

	flag.UintVar(
		&args.Problem,
		"Problem",
		0,
		"Numeric Id of problem as per Projecteuler.net",
	)

	flag.BoolVar(
		&args.Concurrent,
		"Concurrent",
		false,
		"Solves problems concurrently (may be faster)",
	)

	flag.Usage = usageMessage

	flag.Parse()
}

// usageMessage prints a user-readable message of the command-line arguments supported by goeuler.
func usageMessage() {
	fmt.Printf("Help for %s\n", GetVersionString())
	fmt.Println("  --Help              Prints this help message.")
	fmt.Println("  --Version           Prints the version number of this utility.")
	fmt.Println("  --Licence           Prints the copyright licence this utility is release under.")
	fmt.Println("  --Problem <number>  Specifies problem ID to evaluate.")
	fmt.Println("  --AllProblems       Evaluates all problems for which there is code (overrides --Problem)")
	fmt.Println("  --Concurrent        Sovlves problems concurrently (instead of the sequential default)")
	os.Exit(0)
}

// process interprets the command-line arguments parsed by the define method, and responds to the user as appropriate.
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

	if args.Problem == 0 && args.AllProblems == false {
		usageMessage()
	}

	var mustExitWithError = false

	if args.AllProblems == false && (args.Problem < config.MinEulerProblemID || args.Problem > config.MaxEulerProblemID) {
		fmt.Println("Error: Invalid Euler Project ID specified.")
		mustExitWithError = true
	}

	if args.Problem > config.LastImplementedEulerProblemID {
		fmt.Printf("Error: Project ID specified has not been implemented yet. Currently at ID: %d\n", config.LastImplementedEulerProblemID)
		mustExitWithError = true
	}

	if mustExitWithError {
		os.Exit(1)
	}
}

// GetVersionString returns a user-readable string identifying the running application name and version number.
func GetVersionString() string {
	return fmt.Sprintf("%s version %s", os.Args[0], config.Version)
}

// getLicenceString returns a user-readable string describing the copyright licence the software is released under.
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
