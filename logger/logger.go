package logger

import (
	"os"
	"github.com/fatih/color"
)

// Info is used for logging information
// or updates
func Info(msg string) {
	color.Cyan("[bird] " + msg)
}

// Error is used to log error messages
func Error(msg string) {
	color.Red("[bird] " + msg)
}

// Fatal is used to log error messages
// and exit the process
func Fatal(msg string) {
	color.HiRed("[bird] " + msg)
	os.Exit(1)
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
