package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Get input from user's entry
	var input string

	fmt.Println("Enter int/uint8 to calculate memory usage:")
	_, inputErr := fmt.Scanln(&input)
	if inputErr != nil {
		fmt.Println(inputErr)
		return
	} else if input != "int" && input != "uint8" {
		fmt.Println("Only int and uint8 are valid!")
		return
	}

	if input == "int" {
		var intList []int

		for i := 0; i < 10000000; i++ {
			intList = append(intList, 100)
		}

		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		fmt.Printf("Total allocated memory in heap = %v MiB\n", mem.TotalAlloc/1024/1024)
	} else {
		var intList []uint8

		for i := 0; i < 10000000; i++ {
			intList = append(intList, 100)
		}

		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		fmt.Printf("Total allocated memory in heap = %v MiB\n", mem.TotalAlloc/1024/1024)
	}

}
