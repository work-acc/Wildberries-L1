package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	var slice []int = []int{1, 2, 3, 4, 5}

	i := 2
	if i >= 0 && i < len(slice) {
		slice[i] = slice[len(slice)-1]
		slice = slice[:len(slice)-1]

		fmt.Println("Slice after deleting the i-th element", slice)
	} else {
		fmt.Println("The index is out of range")
	}
}
