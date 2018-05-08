// Package glogger provides logging functions with a debug mode switch
// to enable or disable debug messages from printing. Aka verbose mode.
package glogger

import (
	"log"
)

// Define the log level of the message.
const (
	Debug = iota
	Info
	Warning
	Error
)

var debugMode = false

// LogMessage Logs a message to the console.
func LogMessage(level int, message string) {
	if level == Debug && debugMode {
		log.Println("DEBUG: " + message)
	}
	if level == Info {
		log.Println("INFO: " + message)
	}
	if level == Warning {
		log.Println("WARNING: " + message)
	}
	if level == Error {
		log.Fatalln("ERROR: " + message)
	}
}

// SetDebugMode enables or disables debug messages from being printed to the console.
func SetDebugMode(dMode bool) {
	debugMode = dMode
}
