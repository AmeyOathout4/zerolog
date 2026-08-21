package logging

import (
	"log"
	"strings"
)

// LogLevel represents the log level
type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	FATAL LogLevel = "FATAL"
)

// FilterLog filters logs based on the configured log level
func FilterLog(level LogLevel, message string) {
	if level == "" {
		log.Println("Invalid log level")
		return
	}
	
	switch level {
	case DEBUG, INFO, WARN, ERROR, FATAL:
		log.Println(message)
	default:
		log.Println("Invalid log level")
	}
}