package main

import (
	"errors"
	"fmt"
	"unicode"
)

func checkPassword(passwd string) (bool, error) {
	passwdRune := []rune(passwd)

	if len(passwdRune) < 8 {
		return false, errors.New("密碼至少需要8個字元")
	}

	foundUpper := false
	foundLower := false
	foundNumber := false
	foundSymbol := false

	for _, char := range passwdRune {
		if !foundUpper && unicode.IsUpper(char) {
			foundUpper = true
		}

		if !foundLower && unicode.IsLower(char) {
			foundLower = true
		}

		if !foundNumber && unicode.IsNumber(char) {
			foundNumber = true
		}

		if !foundSymbol && (unicode.IsPunct(char) || unicode.IsSymbol(char)) {
			foundSymbol = true
		}
	}

	err := errors.New("密碼不符合規定：\n")

	if !foundUpper {
		err = fmt.Errorf("%w密碼需包含大寫字元\n", err)
	}

	if !foundLower {
		err = fmt.Errorf("%w密碼需包含小寫字元\n", err)
	}

	if !foundNumber {
		err = fmt.Errorf("%w密碼需包含數字\n", err)
	}

	if !foundSymbol {
		err = fmt.Errorf("%w密碼需包含標點符號或者特殊字元", err)
	}

	return foundUpper && foundLower && foundNumber && foundSymbol, err
}

func main() {
	// Get input from user's entry
	var input string

	fmt.Println("Enter a password to test:")
	_, inputErr := fmt.Scanln(&input)
	if inputErr != nil {
		fmt.Println(inputErr)
		return
	}

	if passed, err := checkPassword(input); passed {
		fmt.Println("密碼驗證成功")
	} else {
		fmt.Println(err)
	}
}
