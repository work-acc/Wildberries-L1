package main

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type CounterWithAtomic struct {
	counter uint64
}

func (c *CounterWithAtomic) Increment() {
	atomic.AddUint64(&c.counter, 1)
}

func (c *CounterWithAtomic) Value() uint64 {
	return atomic.LoadUint64(&c.counter)
}

func main() {
	var counterWithAtomic CounterWithAtomic = CounterWithAtomic{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			counterWithAtomic.Increment()
		}()
	}

	wg.Wait()

	fmt.Printf("Counter = %v\n", counterWithAtomic.counter)
}
