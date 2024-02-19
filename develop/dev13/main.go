package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	num1 := 5
	num2 := 3

	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2

	fmt.Println(num1, num2)
}
