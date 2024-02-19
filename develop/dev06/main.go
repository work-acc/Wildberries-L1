package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func main() {
	//Способ 1.
	go func() {
		fmt.Println("Gorutine started")
		time.Sleep(5 * time.Second)

		fmt.Println("Gorutine stoped")
	}()

	//Способ 2.
	stopWork := make(chan bool)
	go func() {
		for {
			select {
			case <-stopWork:
				fmt.Println("Gorutine stoped from channel")

				return
			default:
				fmt.Println("Gorutine working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	//Способ 3.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs

		fmt.Println("Gorutine stoped with signal: ", sig)
	}()

	time.Sleep(15 * time.Second)
	stopWork <- true

	fmt.Println("Working closed")

}
