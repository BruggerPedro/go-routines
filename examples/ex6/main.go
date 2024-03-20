package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", workerId, " Msg: ", res)
		time.Sleep(time.Second)
	}
}

func main() {

	msg := make(chan int)

	// The power of Go is that we can create as many workers as we want and the main function will not be blocked.
	// The main function will not be blocked because the workers are running in parallel.
	// The go process consumes around 2k per thread unlike if we were working with threads on the OS.
	go worker(1, msg)
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)
	go worker(5, msg)
	go worker(6, msg)
	go worker(7, msg)
	go worker(8, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}
}
