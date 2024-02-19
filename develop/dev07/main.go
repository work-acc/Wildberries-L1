package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

func main() {
	var (
		mu    sync.Mutex
		myMap = make(map[string]int)
		wg    sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			key := fmt.Sprintf("key%d", index)

			mu.Lock()
			myMap[key] = index
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
