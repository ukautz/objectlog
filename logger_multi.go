package objectlog

/*
MultiLogger broadcasts messages to multiple loggers, in case one is not sufficient. For example:
Log to local syslog and remote syslog and STDERR ..

	log1 := objectlog.NewStandardLogger() // writes to STDERR
	log2 := objectlog.NewBufferLogger()   // writes to buffer
	log3 := NewYourLogger()
	mult := objectlog.NewMultiLogger(log1, log2, log3)
	mult.LogDebug("Hello all") // writes to all three logers
 */
type (
	MultiLogger struct {
		loggers []ObjectLogger
	}
)

// NewMultiLogger creates new *MultiLogger instance
func NewMultiLogger(loggers ...ObjectLogger) *MultiLogger {
	if loggers == nil {
		loggers = make([]ObjectLogger, 0)
	}
	return &MultiLogger{
		loggers: loggers,
	}
}

// AddLogger adds another logger to the list
func (this *MultiLogger) AddLogger(logger ObjectLogger) *MultiLogger {
	this.loggers = append(this.loggers, logger)
	return this
}

// SetLoggers replaces all loggers with a new list.
func (this *MultiLogger) SetLoggers(loggers ...ObjectLogger) *MultiLogger {
	if loggers == nil {
		loggers = make([]ObjectLogger, 0)
	}
	this.loggers = loggers
	return this
}

// Loggers returns all registered loggers
func (this *MultiLogger) Loggers() []ObjectLogger {
	return this.loggers
}

// Debug writes message to all registered loggers
func (this *MultiLogger) Debug(msg string) {
	for _, logger := range this.loggers {
		logger.Debug(msg)
	}
}

// Info writes message to all registered loggers
func (this *MultiLogger) Info(msg string) {
	for _, logger := range this.loggers {
		logger.Info(msg)
	}
}

// Warn writes message to all registered loggers
func (this *MultiLogger) Warn(msg string) {
	for _, logger := range this.loggers {
		logger.Warn(msg)
	}
}

// Error writes message to all registered loggers
func (this *MultiLogger) Error(msg string) {
	for _, logger := range this.loggers {
		logger.Error(msg)
	}
}

// Fatal writes message to all registered loggers
func (this *MultiLogger) Fatal(msg string) {
	for _, logger := range this.loggers {
		logger.Fatal(msg)
	}
}
