package logger

import (
	"fmt"
	"os"
)

const (
	// InfoFormat is the default info logging format
	InfoFormat = "INFO: %v\n"
	// DebugFormat is the default debug logging format
	DebugFormat = "DEBUG: %v\n"
	// ErrorFormat is the default error logging format
	ErrorFormat = "ERROR: %v\n"
)

var (
	// InfoLogEnabled controls whether Infof prints log
	InfoLogEnabled = true
	// DebugLogEnabled controls whether Infof prints log
	DebugLogEnabled = true
	// ErrorLogEnabled controls whether Infof prints log
	ErrorLogEnabled = true
)

func fPrintf(file *os.File, metaFormat string, format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(file, fmt.Sprintf(metaFormat, fmt.Sprintf(format, args...)))
}

// Infof prints info log to the default out stream. It returns bytes written and any error encountered.
func Infof(format string, args ...interface{}) (int, error) {
	if !InfoLogEnabled {
		return 0, nil
	}
	return fPrintf(os.Stdout, InfoFormat, format, args...)
}

// Debugf prints debug log to the default out stream. It returns bytes written and any error encountered.
func Debugf(format string, args ...interface{}) (int, error) {
	if !DebugLogEnabled {
		return 0, nil
	}
	return fPrintf(os.Stderr, DebugFormat, format, args...)
}

// Errorf prints error log to the default out stream. It returns bytes written and any error encountered.
func Errorf(format string, args ...interface{}) (int, error) {
	if !ErrorLogEnabled {
		return 0, nil
	}
	return fPrintf(os.Stderr, ErrorFormat, format, args...)
}

// Logger is a struct that contains logging functions
type Logger struct {
	InfoFormat      string
	DebugFormat     string
	ErrorFormat     string
	InfoOut         *os.File
	DebugOut        *os.File
	ErrorOut        *os.File
	InfoLogEnabled  bool
	DebugLogEnabled bool
	ErrorLogEnabled bool
}

// CreateLogger returns a new logger struct with default formatting strings and output streams
func CreateLogger() *Logger {
	l := new(Logger)
	l.InfoFormat = InfoFormat
	l.DebugFormat = DebugFormat
	l.ErrorFormat = ErrorFormat
	l.InfoOut = os.Stdout
	l.DebugOut = os.Stderr
	l.ErrorOut = os.Stderr
	l.InfoLogEnabled = true
	l.DebugLogEnabled = true
	l.ErrorLogEnabled = true
	return l
}

// Enter prints debug log that records region entering
func (r *Logger) Enter(region string) (int, error) {
	return r.Debugf("-> Entering %v", region)
}

// Exit prints debug log that records region exiting
func (r *Logger) Exit(region string) (int, error) {
	return r.Debugf("<- Exitting %v", region)
}

// Infof ...
func (r *Logger) Infof(format string, args ...interface{}) (int, error) {
	if !r.InfoLogEnabled {
		return 0, nil
	}
	return fPrintf(r.InfoOut, r.InfoFormat, format, args...)
}

// Debugf ...
func (r *Logger) Debugf(format string, args ...interface{}) (int, error) {
	if !r.DebugLogEnabled {
		return 0, nil
	}
	return fPrintf(r.DebugOut, r.DebugFormat, format, args...)
}

// Errorf ...
func (r *Logger) Errorf(format string, args ...interface{}) (int, error) {
	if !r.ErrorLogEnabled {
		return 0, nil
	}
	return fPrintf(r.ErrorOut, r.ErrorFormat, format, args...)
}
