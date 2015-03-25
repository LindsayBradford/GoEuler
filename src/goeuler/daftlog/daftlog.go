// (c) 2015, Lindsay Bradford

// Package daftlog. A logging package so simple, it's daft

package daftlog

import (
  "fmt"
  "time"
)

const (
	// FormatDateTime formats a log timestamp as per the fmt package pattern "yyyy-mm-dd hh:mm:ss"
  FormatDateTime string = "%04d-%02d-%02d %02d:%02d:%02d"  

	// FormatTimeOnly formats a log timestamp as per the fmt package pattern "hh:mm:ss"
  FormatTimeOnly string = "%02d:%02d:%02d"

	// FormatDateOnly formats a log timestamp as per the fmt package pattern "yyyy-mm-dd"
  FormatDateOnly string = "%04d-%02d-%02d"
  
  // FormatDefault defines the log timestamp to use as default (FormatDateTime)
  FormatDefault         = FormatDateTime
)

var logFormat string

// init ensures the daftlog is initialised with a default log format. 
func init() {
  logFormat = FormatDefault
}

// SetFormat allows the package user to choose from a number of different log formats.
func SetFormat(format string) {
  if format != FormatDateTime && format != FormatTimeOnly && format != FormatDateOnly {
    format = FormatDefault
  }

  logFormat = format
}

// Print creates a new log message, prepending the supplied argument with a formatted timestamp.
// Each call generates a new line in the log stream.
func Print(message string) {
  fmt.Printf(
    "%s | %s\n", now(), message,
  )
}

// Printf creates a new log message, prepending the supplied arguments with a formatted timestamp.
// The arguments supplied are formatted as per fmt.Sprintf before being added to the timestamp. 
// Each call generates a new line in the log stream.
func Printf(format string, a ...interface{}) {
  Print(
    fmt.Sprintf(
      format, a...,
    ),
  )
}

// now generates a formmated timestamp string of the current time as per the logFormat specified
// at initialisation of the package, or via the SetFormat funciton call.
func now() string {
  var now = time.Now()
  
  switch logFormat {
    case FormatDateTime:
      return fmt.Sprintf(
        FormatDateTime,
        now.Year(), now.Month(),  now.Day(), 
        now.Hour(), now.Minute(), now.Second(), 
      )
    case FormatTimeOnly:
      return fmt.Sprintf(
        FormatTimeOnly,
        now.Hour(), now.Minute(), now.Second(), 
       )
    case FormatDateOnly:
      return fmt.Sprintf(
        FormatDateOnly,
        now.Year(), now.Month(),  now.Day(), 
       )
    default:
      return fmt.Sprintf(
        FormatDefault,
        now.Year(), now.Month(),  now.Day(), 
        now.Hour(), now.Minute(), now.Second(), 
      )
  }
}