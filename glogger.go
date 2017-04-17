package glogger

import (
	"fmt"
	"time"
)

const (
	Debug = iota
	Info
	Warning
	Error
)

//Logs a message to the console.
func LogMessage(level int, message string) {
	if level == Debug {
		fmt.Println( time.Now().String() + " | DEBUG: " + message)
	}
	if level == Info {
		fmt.Println( time.Now().String() + " | INFO: " + message)
	}
	if level == Warning {
		fmt.Println(time.Now().String() + " | WARNING: " + message)
	}
	if level == Error {
		fmt.Println(time.Now().String() + " | DEBUG: " + message)
	}
}
