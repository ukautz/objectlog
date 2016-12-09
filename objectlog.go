/*

Package objectlog provides a simple way to extend objects with logging capabilities.

Basics

A couple of lines of code show more than .. etc:

 	type FooThing struct {
 		*objectlog.ObjectLog
 		Name string
 	}

 	// ...

 	func NewFooThing(name string) *FooThing {
 		return &FooThing{
 			ObjectLog: objectlog.NewObjectLog().SetLogPrefix(name+": "),
 			Name: name,
 		}
 	}

 	// ...

	foo := NewFooThing("Mr. Foo")
	foo.LogDebug("Something")   //  "2039-12-24 23:59:59h Mr. Foo: Something"

It is neutral towards the used logger (Standard, Logrus, glog, ..), but it comes with some easy usable defaults.
*/
package objectlog

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	ObjectLogger interface {
		// Debug writes debug level log message
		Debug(msg string)

		// Info writes info level log message
		Info(msg string)

		// Warn writes warning level log message
		Warn(msg string)

		// Error writes error level log message
		Error(msg string)

		// Fatal writes fatal level log message AND exits
		Fatal(msg string)
	}

	// ObjectLogFormatter is function signature to format log message for log output
	ObjectLogFormatter func(level ObjectLogLevel, prefix, suffix, msg string, msgArgs []interface{}, logArgs map[string]interface{}) string

	// ObjectLogLevel represents log level
	ObjectLogLevel string

	// ObjectLog should be used to extend other structs, to provide logging methods
	ObjectLog struct {
		logger    ObjectLogger
		prefix    string
		suffix    string
		formatter ObjectLogFormatter
		args      map[string]interface{}
	}
)

const (
	OBJECT_LOG_LEVEL_DEBUG ObjectLogLevel = "debug"
	OBJECT_LOG_LEVEL_INFO  ObjectLogLevel = "info"
	OBJECT_LOG_LEVEL_WARN  ObjectLogLevel = "warn"
	OBJECT_LOG_LEVEL_ERROR ObjectLogLevel = "error"
	OBJECT_LOG_LEVEL_FATAL ObjectLogLevel = "fatal"
)

var (

	// DefaultFormatter formats default log message: `<prefix><message><suffix>( :: <log-arguments>)`.
	// Log arguments are rendered as JSON - only if they are not empty.
	DefaultFormatter = func(level ObjectLogLevel, prefix, suffix, msg string, msgArgs []interface{}, logArgs map[string]interface{}) string {
		logArgsStr := ""
		if len(logArgs) > 0 {
			raw, err := json.Marshal(logArgs)
			if err == nil {
				logArgsStr = " :: " + string(raw)
			}
		}
		return strings.Join([]string{
			prefix,
			fmt.Sprintf(msg, msgArgs...),
			suffix,
			logArgsStr,
		}, "")
	}

	// DefaultLogger is an instance of `*StandardObjectLogger` and can be globally overwritten.
	// It is used in `NewObjectLog`, if no logger is provided.
	DefaultLogger ObjectLogger = NewStandardLogger()
)

// NewObjectLog creates new ObjectLog instance using default formatter and provided logger. If no logger
// is provided, then `DefaultLogger` is used
func NewObjectLog(logger ...ObjectLogger) *ObjectLog {
	if len(logger) == 0 {
		logger = []ObjectLogger{DefaultLogger}
	}
	return &ObjectLog{
		logger:    logger[0],
		formatter: DefaultFormatter,
		args:      map[string]interface{}{},
	}
}

/*
------------------------------------
  CLONE
------------------------------------
*/

func (this *ObjectLog) LogCloneObjectLog() *ObjectLog {
	clone := NewObjectLog(this.logger)
	clone.formatter = this.formatter
	clone.prefix = this.prefix
	clone.suffix = this.suffix
	for k, v := range this.args {
		clone.args[k] = v
	}
	return clone
}

/*
------------------------------------
  PREFIX & SUFFIX
------------------------------------
*/

// SetLogPrefix sets a prefix for all log messages
//	obj.SetPrefix(obj.ID() + ": ")
func (this *ObjectLog) SetLogPrefix(prefix string) *ObjectLog {
	this.prefix = prefix
	return this
}

// LogPrefix returns the current prefix for log messages (can be empty string)
func (this *ObjectLog) LogPrefix() string {
	return this.prefix
}

// SetLogSuffix sets a suffix for all log messages
//	obj.SetSuffix(fmt.Sprintf(" (%s)", obj.ID())
func (this *ObjectLog) SetLogSuffix(suffix string) *ObjectLog {
	this.suffix = suffix
	return this
}

// LogSuffix returns the current suffix (can be empty string)
func (this *ObjectLog) LogSuffix() string {
	return this.suffix
}

/*
------------------------------------
  ARGS
------------------------------------
*/

// SetLogArgs defines all args which should be logged on every log message. Overwrites existing args!
//	obj.SetArgs(map[string]interface{}{"name": obj.Name()})
func (this *ObjectLog) SetLogArgs(args map[string]interface{}) *ObjectLog {
	if args == nil {
		args = map[string]interface{}{}
	}
	this.args = args
	return this
}

// SetLogArg sets a single log argument. Overwrites existing arg with the same name.
//	obj.SetArg("name", obj.Name())
func (this *ObjectLog) SetLogArg(key string, value interface{}) *ObjectLog {
	this.args[key] = value
	return this
}

// LogArgs returns all currently set log arguments.
func (this *ObjectLog) LogArgs() map[string]interface{} {
	return this.args
}

/*
------------------------------------
  LOGGER
------------------------------------
*/

// SetLogger replaces the current logger with another
func (this *ObjectLog) SetLogger(logger ObjectLogger) *ObjectLog {
	this.logger = logger
	return this
}

// Logger returns the currently configured logger
func (this *ObjectLog) Logger() ObjectLogger {
	return this.logger
}

/*
------------------------------------
  LOG METHODS
------------------------------------
*/

func (this *ObjectLog) build(level ObjectLogLevel, msg string, args ...interface{}) string {
	return this.formatter(level, this.prefix, this.suffix, msg, args, this.args)
}

// LogDebug writes the log message in DEBUG level
func (this *ObjectLog) LogDebug(msg string, args ...interface{}) {
	this.logger.Debug(this.build(OBJECT_LOG_LEVEL_DEBUG, msg, args...))
}

// LogInfo writes the log message in INFO level
func (this *ObjectLog) LogInfo(msg string, args ...interface{}) {
	this.logger.Info(this.build(OBJECT_LOG_LEVEL_INFO, msg, args...))
}

// LogWarn writes the log message in WARN level
func (this *ObjectLog) LogWarn(msg string, args ...interface{}) {
	this.logger.Warn(this.build(OBJECT_LOG_LEVEL_WARN, msg, args...))
}

// LogError writes the log message in ERROR level
func (this *ObjectLog) LogError(msg string, args ...interface{}) {
	this.logger.Error(this.build(OBJECT_LOG_LEVEL_ERROR, msg, args...))
}

// LogFatal writes the log message in FATAL level - and usually exits (depends on used `ObjectLogger`)
func (this *ObjectLog) LogFatal(msg string, args ...interface{}) {
	this.logger.Fatal(this.build(OBJECT_LOG_LEVEL_FATAL, msg, args...))
}
