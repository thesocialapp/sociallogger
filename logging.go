// backend/go/logging/logging.go

package logging

import (
	"log"
	"os"
)

// Logger provides a minimal logging interface
type Logger struct {
	*log.Logger
}

// NewLogger creates a new Logger
func NewLogger() *Logger {
	// Log to stdout
	l := log.New(os.Stdout, "[MYPROJECT] ", log.LstdFlags|log.Lshortfile)

	return &Logger{l}
}

// Printf logs with Printf from the standard logger
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

// Fatalf logs with Fatalf from the standard logger
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// HandleError logs an error with Printf
func (l *Logger) HandleError(err error) {
	if err != nil {
		l.Printf("Error: %v", err)
	}
}
