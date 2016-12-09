package objectlog

import (
	"log"
	"os"
)

/*
StandardLogger is adapter the default "*log.Logger", included in the Go language standard libraries.
 */
type (
	StandardLogger struct {
		logger *log.Logger
	}
)

// NewStandardLogger creates new *StandardLogger from the provided *log.Logger instance. If no instance
// is provided it defaults to a *log.Logger writing to STDERR
//
//	// using default
//	logToSTDERR := NewStandardLogger()
//
//	// using custom
//	fh, _ := os.OpenFile(..)
//	fhLog := log.New(fh, "", log.LstdFlags)
//	logToFile := NewStandardLogger(fhLog)
func NewStandardLogger(logger ...*log.Logger) *StandardLogger {
	if len(logger) == 0 {
		logger = []*log.Logger{log.New(os.Stderr, "", log.LstdFlags)}
	}

	return &StandardLogger{
		logger: logger[0],
	}
}

func (this *StandardLogger) Debug(msg string) {
	this.logger.Print("[DEBUG] " + msg)
}

func (this *StandardLogger) Info(msg string) {
	this.logger.Print("[INFO] " + msg)
}

func (this *StandardLogger) Warn(msg string) {
	this.logger.Print("[WARN] " + msg)
}

func (this *StandardLogger) Error(msg string) {
	this.logger.Print("[ERROR] " + msg)
}

func (this *StandardLogger) Fatal(msg string) {
	this.logger.Fatal(msg)
}
