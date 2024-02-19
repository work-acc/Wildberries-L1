package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	var str []string = []string{"cat", "cat", "dog", "cat", "tree"}
	var res []string
	var unic map[string]string = make(map[string]string)

	for _, value := range str {
		_, exist := unic[value]
		if !exist {
			unic[value] = ""
			res = append(res, value)
		}
	}

	fmt.Println(res)
}
