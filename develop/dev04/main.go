package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	dataChan := make(chan int)

	go writeToChannel(dataChan)

	numWorkers := 5

	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(dataChan, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	<-sigChan

	close(dataChan)

	wg.Wait()
	fmt.Println("All workers have finished. Exiting...")
}

func writeToChannel(ch chan int) {
	i := 0
	for {
		ch <- i
		i++
	}
}

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Println("Worker:", data)
	}
}
