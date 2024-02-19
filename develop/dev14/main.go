package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	variable := []interface{}{100, "Redkin", true, make(chan bool)}

	for _, value := range variable {
		switch t := value.(type) {
		case int:
			fmt.Println("Type is an integer:", t)
		case string:
			fmt.Println("Type is a string:", t)
		case bool:
			fmt.Println("Type is a bool:", t)
		case chan bool:
			fmt.Println("Type is a channel:", t)
		default:
			fmt.Println("Type is unknown!")
		}
	}
}
