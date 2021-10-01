package log

import (
	"fmt"
	"os"
)

// severity is number in range 0-5 (see SetSeverity for more info)
var severity int

// Set severity for log. It's number in range 0-5, where 0 means
// all logs are disabled and 5 means all (even debug) logs will
// be printed.
func SetSeverity(s int) {
	severity = s
}

// the function log any message. The messages are printed to stderr,
// to keep the stdout clean for application output
func output(s int, format string, a ...interface{}) {
	// write only when severity of the message 's' is in range of
	// currently set 'severity'
	if s <= severity {
		fmt.Fprintf(os.Stderr, format, a...)
		fmt.Fprint(os.Stderr, "\n")
	}
}

// Errors are printed only when severity is set to 1, with 0,
// the errors are ignored and application need to handle it
func Error(err error) {
	output(1, "[ERROR]: %v", err)
}

// Warnings are printed for severity 2 and higher
func Warn(msg string, a ...interface{}) {
	output(2, "[WARN]:"+msg, a)
}

// Debug messages are printed for severity 5
func Debug(msg string, a ...interface{}) {
	output(5, "[DEBUG]:"+msg, a...)
}
