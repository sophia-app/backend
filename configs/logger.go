package configs

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger is a custom logger that logs messages with different levels.
type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	write io.Writer
}

// NewLogger creates a new logger with the given prefix.
func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug: log.New(writer, fmt.Sprintf("(%s) DEBUG: ", logger.Prefix()), logger.Flags()),
		info:  log.New(writer, fmt.Sprintf("(%s) INFO: ", logger.Prefix()), logger.Flags()),
		warn:  log.New(writer, fmt.Sprintf("(%s) WARNING: ", logger.Prefix()), logger.Flags()),
		err:   log.New(writer, fmt.Sprintf("(%s) ERROR: ", logger.Prefix()), logger.Flags()),
		write: writer,
	}
}

// Debug logs a message with DEBUG level.
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Info logs a message with INFO level.
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Warn logs a message with WARNING level.
func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

// Error logs a message with ERROR level.
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Debugf logs a formatted message with DEBUG level.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// Infof logs a formatted message with INFO level.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// Warnf logs a formatted message with WARNING level.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

// Errorf logs a formatted message with ERROR level.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
