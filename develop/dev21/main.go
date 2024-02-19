package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

type Human interface {
	Requirement()
}

type LittleBro struct{}

func (adaptee *LittleBro) AdaptedRequirement() {
	fmt.Println("I want to eat")
}

type HumanAdapter struct {
	*LittleBro
}

func (adapter *HumanAdapter) Requirement() {
	adapter.AdaptedRequirement()
}

func NewAdapter(adaptee *LittleBro) Human {
	return &HumanAdapter{adaptee}
}

func main() {
	fmt.Println("My brother said:")
	adapter := NewAdapter(&LittleBro{})
	adapter.Requirement()
}
