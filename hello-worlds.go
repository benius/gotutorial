package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	// declare with type and the initial value
	//debug       bool      = false
	//logLevel    string    = "info"
	//startUpTime time.Time = time.Now()

	// declare with the initial value, but no type
	debug       = true
	logLevel    = "warn"
	startUpTime = time.Now()
)

func main() {
	fmt.Println("Debug:", debug, ";log level:", logLevel, ";startup time:", startUpTime)
	fmt.Printf("Debug: %s; log level: %s; startup time: %v", strconv.FormatBool(debug), logLevel, startUpTime)
}
