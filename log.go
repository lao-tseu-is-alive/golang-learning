package main

import (
	"github.com/mgutz/ansi"
	"log"
	"os"
)

var loggerInfo = log.New(os.Stdout, "INFO: ", log.Lshortfile)
var loggerTrace = log.New(os.Stdout, "TRACE: ", log.Lshortfile)
var loggerWarning = log.New(os.Stderr, "WARNING: ", log.Lshortfile)
var loggerError = log.New(os.Stderr, "ERROR: ", log.Lshortfile)

func Info(message string) {
	blue := ansi.ColorFunc("cyan+")
	loggerInfo.Output(2, blue(message))
}

func Trace(message string) {
	magenta := ansi.ColorFunc("magenta+b")
	loggerTrace.Output(2, magenta(message))
}

func Warn(message string) {
	magenta := ansi.ColorFunc("red+b+green")
	loggerWarning.Output(2, magenta(message))
}

func Err(message string) {
	magenta := ansi.ColorFunc("red+b+green")
	loggerError.Output(2, magenta(message))
}
