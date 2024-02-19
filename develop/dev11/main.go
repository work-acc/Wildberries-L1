package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	var result []int
	result = intersection([]int{1, 2, 3, 4}, []int{2, 3, 3})

	fmt.Println("result = ", result)
}

func intersection(first []int, second []int) (res []int) {
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				res = append(res, first[i])

				break
			}
		}
	}

	return res
}
