package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	FirstName string
	LastName  string
}

type Action struct {
	Human
	age int
}

func main() {
	act := Action{
		Human: Human{FirstName: "Vladislav", LastName: "Redkin"},
		age:   23,
	}

	fmt.Println(act.FullName())
}

func (h Human) FullName() string {
	return h.FirstName + " " + h.LastName
}
