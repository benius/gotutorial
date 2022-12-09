package main

import "fmt"

func createSequenceArray() [10]int {
	var intArray = [10]int{0: 1}

	for i := 1; i < len(intArray); i++ {
		intArray[i] = intArray[i-1] + 1
	}

	return intArray
}

func powerOf2(inputArray [10]int) [10]int {
	// copy by value
	arrayCopy := inputArray

	for i := 0; i < len(arrayCopy); i++ {
		arrayCopy[i] = arrayCopy[i] * arrayCopy[i]
	}

	return arrayCopy
}

func main() {
	seqArray := createSequenceArray()

	msg := ""
	for i := 0; i < len(seqArray); i++ {
		msg += fmt.Sprintf("%v: %v\n", i, seqArray[i])
	}

	fmt.Printf("Sequence array elements:\n%v\n", msg)

	powerArray := powerOf2(seqArray)

	msg = ""
	for i := 0; i < len(powerArray); i++ {
		msg += fmt.Sprintf("%v: %v\n", i, powerArray[i])
	}

	fmt.Printf("Power of 2 array elements:\n%v\n", msg)
}
