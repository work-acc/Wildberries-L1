package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	fmt.Println("Before the sleep")
	sleep(2000)
	fmt.Println("After sleeping")
}

func sleep(milliseconds int) {
	<-time.After(time.Duration(milliseconds) * time.Millisecond)
}
