package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// multiple lines short variable declarations
	//debug := false
	//logLevel := "trace"
	//startUpTime := time.Now()

	// single line short variable declaration
	debug, logLevel, startUpTime := false, "trace", time.Now()

	fmt.Printf("Debug: %s; log level: %s; start up time: %v", strconv.FormatBool(debug), logLevel, startUpTime)
}
