package main

import (
	"fmt"
	"strconv"
	"time"
)

func findConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	//debug, logLevel, startUpTime := findConfig()
	var debug, logLevel, startUpTime = findConfig()
	fmt.Printf("Debug: %s; log level: %s; start up time: %v", strconv.FormatBool(debug), logLevel, startUpTime)
}
