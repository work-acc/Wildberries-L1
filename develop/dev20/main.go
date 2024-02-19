package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	var result string
	result = reverseWords("show dog sun")

	fmt.Println("result = ", result)
}

func reverseWords(str string) (res string) {
	words := strings.Split(str, " ")
	for i := len(words) - 1; i >= 0; i-- {
		if i == 0 {
			res += words[i]

			continue
		}

		res += words[i] + " "
	}

	return res
}
