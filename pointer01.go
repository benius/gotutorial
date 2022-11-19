package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int

	count2 := new(int)

	countTemp := 5
	count3 := &countTemp

	startTime := &time.Time{}

	fmt.Printf("count1: %#v\ncount2: %#v\ncount3: %#v\nstartTime: %#v", count1, count2, count3, startTime)

	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	} else {
		fmt.Print("count1 is nil\n")
	}

	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}

	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}

	if startTime != nil {
		fmt.Printf("startTime: %#v, display value: %#v\n", *startTime, startTime.String())
	}
}
