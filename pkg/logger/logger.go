package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Info(message string) {
	l.Println("INFO: " + message)
}

func (l *Logger) Error(message string) {
	l.Println("ERROR: " + message)
}

func (l *Logger) Warn(message string) {
	l.Println("WARN: " + message)
}
