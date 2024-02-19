package main

import (
	"fmt"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// В данном фрагменте кода есть потенциальная проблема с производительностью и использованием памяти из-за создания огромной строки
//с помощью createHugeString и копирования ее части в переменную justString.
// Проблема в том, что при создании огромной строки мы используем много памяти, а затем копируем только ее часть в justString,
//что также требует дополнительной памяти.
// Чтобы избежать этой проблемы, мы можем напрямую присвоить подстроку переменной justString без создания огромной строки.
//Это сэкономит память и повысит производительность.

var justString string

func someFunc() string {
	str := createHugeString(1 << 10)

	return str[:100]
}

func createHugeString(size int) string {
	var bytes []byte = make([]byte, size)
	for i := 0; i < 100; i++ {
		bytes[i] = 'A'
	}
	return string(bytes)
}

func main() {
	tempStr := someFunc()
	justString = string([]byte(tempStr))
	tempStr = ""

	fmt.Println(justString)
}
