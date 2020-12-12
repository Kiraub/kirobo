package logger

import (
	"fmt"
	"os"
)

const (
	infoFormat  = "INFO: %v\n"
	errorFormat = "ERROR: %v\n"
	debugFormat = "DEBUG: %v\n"
)

var (
	// InfoLogEnabled controls whether Infof prints log
	InfoLogEnabled = true
	// DebugLogEnabled controls whether Infof prints log
	DebugLogEnabled = true
	// ErrorLogEnabled controls whether Infof prints log
	ErrorLogEnabled = true
)

func fprintf(file *os.File, metaFormat string, format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(file, fmt.Sprintf(metaFormat, fmt.Sprintf(format, args...)))
}

// Infof prints info log to the default out stream. It returns bytes written and any error encountered.
func Infof(format string, args ...interface{}) (int, error) {
	if !InfoLogEnabled {
		return 0, nil
	}
	return fprintf(os.Stdout, infoFormat, format, args...)
}

// Debugf prints debug log to the default out stream. It returns bytes written and any error encountered.
func Debugf(format string, args ...interface{}) (int, error) {
	if !DebugLogEnabled {
		return 0, nil
	}
	return fprintf(os.Stderr, debugFormat, format, args...)
}

// Errorf prints error log to the default out stream. It returns bytes written and any error encountered.
func Errorf(format string, args ...interface{}) (int, error) {
	if !ErrorLogEnabled {
		return 0, nil
	}
	return fprintf(os.Stderr, errorFormat, format, args...)
}
