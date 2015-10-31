package logger

import (
	"log"

	"github.com/mgutz/ansi"
)

// Fatal ...
func Fatal(s string) {
	log.Fatal(ansi.Color("FATAL: ", ansi.Red), ansi.Color(s, ansi.Red))
}

// Error ...
func Error(s string) {
	log.Fatal(ansi.Color("ERROR: ", ansi.Red), ansi.Color(s, ansi.Red))
}

// Warn ...
func Warn(s string) {
	log.Fatal(ansi.Color("WARN: ", ansi.Red), ansi.Color(s, ansi.Red))
}

// Info ...
func Info(s string) {
	log.Fatal(ansi.Color("INFO: ", ansi.Red), ansi.Color(s, ansi.Red))
}

// Debug ...
func Debug(s string) {
	log.Fatal(ansi.Color("DEBUG: ", ansi.Red), ansi.Color(s, ansi.Red))
}

// Trace ...
func Trace(s string) {
	log.Fatal(ansi.Color("TRACE: ", ansi.Red), ansi.Color(s, ansi.Red))
}
