package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go sendData(ch)

	timeout := 5 * time.Second
	timer := time.NewTimer(timeout)

	for {
		select {
		case data, ok := <-ch:
			if ok {
				fmt.Println("Received:", data)
			} else {
				fmt.Println("The channel is closed")

				return
			}
		case <-timer.C:
			fmt.Println("Time is up")

			return
		}
	}
}
