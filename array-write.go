package main

import "fmt"

func main() {
	arr1 := [...]string{
		"Get",
		"ready",
		"to",
		"go",
	}

	arr1[0] = "It's"
	arr1[1] = "time"

	seq := fmt.Sprintln(arr1[0], arr1[1], arr1[2], arr1[3])
	fmt.Print(seq)
}
