package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(binarySearch(5, array))
}

func binarySearch(lookingFor int, arr []int) int {
	min := 0
	max := len(arr) - 1

	for min <= max {
		mid := min + (max-min)/2

		if arr[mid] == lookingFor {
			return mid
		} else if arr[mid] < lookingFor {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}
