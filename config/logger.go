package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		error:   log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(msg ...interface{}) {
	l.debug.Println(msg...)
}

func (l *Logger) Warning(msg ...interface{}) {
	l.warning.Println(msg...)
}

func (l *Logger) Info(msg ...interface{}) {
	l.info.Println(msg...)
}

func (l *Logger) Error(msg ...interface{}) {
	l.error.Println(msg...)
}

func (l *Logger) Debugf(format string, msg ...interface{}) {
	l.debug.Printf(format, msg...)
}

func (l *Logger) Warningf(format string, msg ...interface{}) {
	l.warning.Printf(format, msg...)
}

func (l *Logger) Infof(format string, msg ...interface{}) {
	l.info.Printf(format, msg...)
}

func (l *Logger) Errorf(format string, msg ...interface{}) {
	l.error.Printf(format, msg...)
}
