package main

import "fmt"

func main() {
	var (
		arr1 [10]int
		arr2 = [...]int{9: 0}
		arr3 = [10]int{1, 9: 10, 4: 5}
		arr4 = [10]int{9: 10, 4: 5, 1}
	)

	fmt.Println("[10]int == [...]int {9: 0}                            :", arr1 == arr2)
	fmt.Println("[10]int == [10]int {1, 9: 10, 4: 5}                   :", arr1 == arr3)
	fmt.Println("[10]int {1, 9: 10, 4: 5} == [10]int {9: 10, 4: 5, 1}  :", arr3 == arr4)
	fmt.Println("[10]int {1, 9: 10, 4: 5} =", arr3)
	fmt.Println("[10]int {9: 10, 4: 5, 1} =", arr4)
}
