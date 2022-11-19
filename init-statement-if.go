package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("the integer must be a positive value")
	} else if input > 100 {
		return errors.New("the integer must be smaller than 100")
	} else if input%7 == 0 {
		return errors.New("the integer cannot be a multiple of 7")
	} else {
		return nil
	}
}

func main() {
	// Get input from user's entry
	var input int

	fmt.Println("Enter a integer to test:")
	_, inputErr := fmt.Scanln(&input)
	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}

	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even.")
	} else {
		fmt.Println(input, "is odd.")
	}
}
