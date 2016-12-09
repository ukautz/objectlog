package logrus

import (
	lr "github.com/Sirupsen/logrus"
)

type (
	LogrusObjectLogger struct {
		logger *lr.Logger
	}
)

func NewLogrusObjectLogger(logger *lr.Logger) *LogrusObjectLogger {
	if logger == nil {
		logger = lr.New()
	}
	return &LogrusObjectLogger{
		logger: logger,
	}
}

func (this *LogrusObjectLogger) Debug(msg string) {
	this.logger.Debug(msg)
}

func (this *LogrusObjectLogger) Info(msg string) {
	this.logger.Info(msg)
}

func (this *LogrusObjectLogger) Warn(msg string) {
	this.logger.Warn(msg)
}

func (this *LogrusObjectLogger) Error(msg string) {
	this.logger.Error(msg)
}

func (this *LogrusObjectLogger) Fatal(msg string) {
	this.logger.Fatal(msg)
}
