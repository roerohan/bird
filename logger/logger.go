package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Log is a struct to define the
// message and the function used to
// log the message of that type
type Log struct {
	Message string
	Func    func(string)
}

// Info is used for logging information
// or updates
func Info(msg string) {
	color.Cyan(fmt.Sprintf("\r[bird] %-50s ", msg))
}

// Error is used to log error messages
func Error(msg string) {
	color.Red(fmt.Sprintf("\r[bird] %-50s ", msg))
}

// Fatal is used to log error messages
// and exit the process
func Fatal(msg string) {
	color.HiRed(fmt.Sprintf("\r[bird] %-50s ", msg))
	os.Exit(1)
}

// Print exports a call to
// fmt.Print limited to type
// string
func Print(msg string) {
	fmt.Print(msg)
}

// Welcome displays the welcome message for bird
func Welcome() {
	color.Green(`
_\|      __       |/_
_-  \_  _/"->   _/  -_
-_    ''(   )'-'    _-
'=.__.=-(   )-=.__.='
        |/-\|
_ _ _ _ Y   Y _ _bird
	`)
}

// Start listens on a channel
// and executes log functions on
// that particular channel
func Start(logs chan Log) {
	for l := range logs {
		l.Func(l.Message)
	}
}
