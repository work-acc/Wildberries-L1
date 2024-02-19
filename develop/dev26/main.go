package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// 	aabcd — false

func main() {
	var result bool
	result = CountLetters("Abcde")
	fmt.Println("result = ", result)
}

func CountLetters(str string) bool {
	str = strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {

				return false
			}
		}
	}

	return true
}
