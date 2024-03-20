package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) // To analyze we will run on just 1 core
	fmt.Println("It started")

	go func() {
		for {
			
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("It ended")
	// In other versions of Go, the program will not end because the goroutine will be running forever and the main function will not be able to finish.
}
