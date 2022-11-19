package main

import "fmt"

func valueAdd5(count int) {
	count += 5
	fmt.Println("valueAdd5: ", count)
}

func pointerAdd5(count *int) {
	*count += 5
	fmt.Println("pointerAdd5: ", *count)
}

func main() {
	var count int

	valueAdd5(count)
	fmt.Println("Value of count after valueAdd5: ", count)

	pointerAdd5(&count)
	fmt.Println("Value of count after pointerAdd5: ", count)
}
