package logger

import (
	"context"
	"io"
	"log"
	"sync"
)

const (
	// InfoPrefix is the default info logging prefix
	InfoPrefix = "[INFO]"
	// DebugPrefix is the default debug logging prefix
	DebugPrefix = "[DEBUG]"
	// ErrorPrefix is the default error logging prefix
	ErrorPrefix = "[ERROR]"
)

type key int

const loggerKey key = 0

type mutexBool struct {
	v bool
	m sync.Mutex
}

func (r *mutexBool) Set(nv bool) {
	r.m.Lock()
	defer r.m.Unlock()
	r.v = nv
}

func (r *mutexBool) Get() bool {
	r.m.Lock()
	defer r.m.Unlock()
	return r.v
}

// Logger is a struct that contains logging functions
type Logger struct {
	lInfo  *log.Logger
	lDebug *log.Logger
	lError *log.Logger

	infoLogEnabled  *mutexBool
	debugLogEnabled *mutexBool
	errorLogEnabled *mutexBool
}

// New creates a new logger
func New(out io.Writer, err io.Writer, prefix string, flag int) *Logger {
	l := new(Logger)

	l.lInfo = log.New(out, prefix+" "+InfoPrefix, flag)
	l.lDebug = log.New(err, prefix+" "+DebugPrefix, flag)
	l.lError = log.New(err, prefix+" "+ErrorPrefix, flag)

	l.infoLogEnabled.Set(true)
	l.debugLogEnabled.Set(true)
	l.errorLogEnabled.Set(true)
	return l
}

// NewContext returns a new Context that carries a provided Logger value
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext extracts a Logger from a Context
func FromContext(ctx context.Context) (Logger, bool) {
	logger, ok := ctx.Value(loggerKey).(Logger)
	return logger, ok
}

// ToggleInfo ...
func (r *Logger) ToggleInfo(shouldPrint bool) {
	r.infoLogEnabled.Set(shouldPrint)
}

// ToggleDebug ...
func (r *Logger) ToggleDebug(shouldPrint bool) {
	r.debugLogEnabled.Set(shouldPrint)
}

// ToggleError ...
func (r *Logger) ToggleError(shouldPrint bool) {
	r.errorLogEnabled.Set(shouldPrint)
}

// Enter prints debug log that records region entering
func (r *Logger) Enter(region string) {
	r.Debugf("-> Entering %v", region)
}

// Exit prints debug log that records region exiting
func (r *Logger) Exit(region string) {
	r.Debugf("<- Exitting %v", region)
}

// Infof ...
func (r *Logger) Infof(format string, args ...interface{}) {
	if !r.infoLogEnabled.Get() {
		return
	}
	r.lInfo.Printf(format, args...)
}

// Debugf ...
func (r *Logger) Debugf(format string, args ...interface{}) {
	if !r.debugLogEnabled.Get() {
		return
	}
	r.lDebug.Printf(format, args...)
}

// Errorf ...
func (r *Logger) Errorf(format string, args ...interface{}) {
	if !r.errorLogEnabled.Get() {
		return
	}
	r.lError.Printf(format, args...)
}
