package glogger

import (
	"fmt"
	"time"
)

//Define the log level of the message.
const (
	Debug = iota
	Info
	Warning
	Error
)

var debugMode = false

//LogMessage Logs a message to the console.
func LogMessage(level int, message string) {
	if level == Debug && debugMode {
		fmt.Println(time.Now().Format(time.RFC822) + " | DEBUG: " + message)
	}
	if level == Info {
		fmt.Println(time.Now().Format(time.RFC822) + " | INFO: " + message)
	}
	if level == Warning {
		fmt.Println(time.Now().Format(time.RFC822) + " | WARNING: " + message)
	}
	if level == Error {
		fmt.Println(time.Now().Format(time.RFC822) + " | DEBUG: " + message)
	}
}

//SetDebugMode enables or disables debug messages from being printed to the console.
func SetDebugMode(dMode bool) {
	debugMode = dMode
}
