package main

import "fmt"

func main() {
	arr1 := [...]string{
		"ready",
		"get",
		"go",
		"to",
	}

	seq := fmt.Sprintln(arr1[1], arr1[0], arr1[3], arr1[2])
	fmt.Print(seq)
}
