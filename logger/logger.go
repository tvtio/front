package logger

import (
	"log"
	"strings"

	"github.com/mgutz/ansi"
)

// Fatal ...
func Fatal(s ...string) {
	log.Fatal(ansi.Color("FATAL: ", "red"), strings.Join(s, " "))
}

// Error ...
func Error(s ...string) {
	log.Println(ansi.Color("ERROR: ", "red"), strings.Join(s, " "))
}

// Warn ...
func Warn(s ...string) {
	log.Println(ansi.Color("WARN: ", "red"), strings.Join(s, " "))
}

// Info ...
func Info(s ...string) {
	log.Println(ansi.Color("INFO: ", "yellow"), strings.Join(s, " "))
}

// Debug ...
func Debug(s ...string) {
	log.Println(ansi.Color("DEBUG: ", "green"), strings.Join(s, " "))
}

// Trace ...
func Trace(s ...string) {
	log.Println(ansi.Color("TRACE: ", "green"), strings.Join(s, " "))
}
