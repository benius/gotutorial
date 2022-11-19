package main

import "fmt"

func main() {
	poem := "青裙玉面初相識，九月茶花滿路開。"

	fmt.Printf("Poem: %s\n", poem)
	fmt.Printf("Length in characters: %d\n", len([]rune(poem)))
	fmt.Printf("Length in bytes: %d\n", len(poem))

	fmt.Println("------------------------------------")
	fmt.Println("cnt	index	rune	char	bytes")
	fmt.Println("------------------------------------")
	cnt := 0
	for index, char := range poem {
		cnt++
		fmt.Printf("%-2d\t%-2d\t%U\t%2c\t%X\n", cnt, index, char, char, []byte(string(char)))
	}
}
