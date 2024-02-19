package main

import (
	"fmt"
)

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	input := "главрыба — абырвалг"
	reversed := reverseString(input)

	fmt.Println("Reversed:", reversed)
}

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)
	lastIndex := length - 1

	for i := 0; i < lastIndex; i++ {
		runes[i], runes[lastIndex] = runes[lastIndex], runes[i]
		lastIndex--
	}

	return string(runes)
}
