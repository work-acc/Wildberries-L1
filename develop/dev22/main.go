package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := new(big.Int)
	a.SetString("9223372036854775807", 10)
	b := new(big.Int)
	b.SetString("9223372036854775807", 10)
	result := new(big.Int)

	multiplyResult := result.Mul(a, b)
	fmt.Println("Multiplication:", multiplyResult)

	divideResult := result.Div(a, b)
	fmt.Println("Division:", divideResult)

	sumResult := result.Add(a, b)
	fmt.Println("Addition:", sumResult)

	subtractResult := result.Sub(a, b)
	fmt.Println("Subtraction:", subtractResult)
}
