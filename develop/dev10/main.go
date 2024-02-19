package main

import (
	"fmt"
	"sync"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	var wg sync.WaitGroup
	for _, value := range temperatures {
		wg.Add(1)

		go func(value float64) {
			wg.Done()

			div := int(value) / 10
			key := div * 10
			groups[key] = append(groups[key], value)
		}(value)
	}

	wg.Wait()

	for key, value := range groups {
		fmt.Printf("%v: %v\n", key, value)
	}
}
