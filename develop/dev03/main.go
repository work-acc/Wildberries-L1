package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := sum(numbers)

	fmt.Println(sum)
}

func sum(numbers []int) (sum int) {
	var wg sync.WaitGroup

	for _, value := range numbers {
		wg.Add(1)

		go func(value int) {
			defer wg.Done()

			sum += value * value
		}(value)
	}

	wg.Wait()

	return sum
}
