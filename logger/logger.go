package logger

import "github.com/inconshreveable/log15"

// This logger wrapper is convenient for unit testing.
// It makes it possible to skip defining calls for mock.

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Logger --case=underscore

// Logger defines logger interface.
type Logger interface {
	DebugKV(message string, kvs ...interface{})
	InfoKV(message string, kvs ...interface{})
	WarnKV(message string, kvs ...interface{})
	ErrorKV(message string, kvs ...interface{})
	FatalKV(message string, kvs ...interface{})
	// AddContext adds a scope key-value pairs to the logger and returnes scoped logger
	AddContext(ctx ...interface{}) Logger
}

type logger struct {
	logger log15.Logger
}

// NewEmptyLogger is for units tests where you don't want to define logger mock calls.
func NewEmptyLogger() Logger {
	return &logger{}
}

// NewLogger returns a new instance of Logger.
// ctxKvs adds a scope key-value pairs to the logger and returnes scoped logger.
func NewLogger(ctxKvs ...interface{}) Logger {
	return &logger{
		logger: log15.New(ctxKvs...),
	}
}

func (l *logger) DebugKV(message string, kvs ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Debug(message, kvs...)
}

func (l *logger) InfoKV(message string, kvs ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Info(message, kvs...)
}

func (l *logger) WarnKV(message string, kvs ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Warn(message, kvs...)
}

func (l *logger) ErrorKV(message string, kvs ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Error(message, kvs...)
}

func (l *logger) FatalKV(message string, kvs ...interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.Crit(message, kvs...)
}

func (l *logger) AddContext(ctxKvs ...interface{}) Logger {
	if l.logger == nil {
		return &logger{}
	}

	return &logger{
		logger: l.logger.New(ctxKvs...),
	}
}
