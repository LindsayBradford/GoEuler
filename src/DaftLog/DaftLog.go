// Daft Log - A log so simple, it's daft
// (c) 2015, Lindsay Bradford

package DaftLog

import (
  "time"
  "fmt"
)

const (
	DATE_TIME_FORMAT string = "%04d-%02d-%02d %02d:%02d:%02d"  // "yyyy-mm-dd hh:mm:ss"
  TIME_ONLY_FORMAT string = "%02d:%02d:%02d"                 // "hh:mm:ss"
  DATE_ONLY_FORMAT string = "%04d-%02d-%02d"                 // "yyyy-mm-dd"
)

const DEFAULT_FOTMAT = DATE_TIME_FORMAT

var timestampFormat string

func init() {
	timestampFormat = DEFAULT_FOTMAT
}

func SetFormat(format string) {
	if format != DATE_TIME_FORMAT && format != TIME_ONLY_FORMAT && format != DATE_ONLY_FORMAT {
		format = DEFAULT_FOTMAT
	}
	
	timestampFormat = format
}

func FormattedLogEntry(format string, a ...interface{}) {
  LogEntry(
  	fmt.Sprintf(
  		format, a...,
  	),
  )
}

func LogEntry(message string) {
  fmt.Printf(
  	"%s | %s\n", nowTimestamp(), message,
  )
}

func nowTimestamp() string {
  var now = time.Now()
  
  switch timestampFormat {
  	case DATE_TIME_FORMAT:
		  return fmt.Sprintf(
    		DATE_TIME_FORMAT,
    		now.Year(), now.Month(),  now.Day(), 
    		now.Hour(), now.Minute(), now.Second(), 
    	)
    case TIME_ONLY_FORMAT:
	    return fmt.Sprintf(
  	    TIME_ONLY_FORMAT,
    	  now.Hour(), now.Minute(), now.Second(), 
   	  )
    case DATE_ONLY_FORMAT:
	    return fmt.Sprintf(
  	    DATE_ONLY_FORMAT,
    		now.Year(), now.Month(),  now.Day(), 
   	  )
	  default:
  	  return fmt.Sprintf(
    		DATE_TIME_FORMAT,
    		now.Year(), now.Month(),  now.Day(), 
    		now.Hour(), now.Minute(), now.Second(), 
    	)
  }
} 