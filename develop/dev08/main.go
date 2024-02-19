package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var bit int64 = 5

	bit = install(bit, 3, true)
	fmt.Printf("After setting the 3rd bit to 1: %d\n", bit)

	bit = install(bit, 5, false)
	fmt.Printf("After setting the 5th bit to 0: %d\n", bit)
}

func install(bit int64, i uint, value bool) int64 {
	if value {
		return bit | (1 << i)
	} else {
		return bit &^ (1 << i)
	}
}
