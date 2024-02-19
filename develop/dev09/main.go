package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	var numbers []int = []int{1, 2, 3, 4, 5}
	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for {
			x, open := <-input
			if !open {
				close(output)

				return
			}

			output <- x * 2
		}
	}()

	go func() {
		defer wg.Done()

		for {
			x, open := <-output
			if open {
				fmt.Println(x)
			} else {
				return
			}
		}
	}()

	for _, value := range numbers {
		input <- value
	}

	close(input)
	wg.Wait()
}
