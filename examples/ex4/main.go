package main

import (
	"fmt"
	"time"
)

func main() {
	hello := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		hello <- "Hello, world!"
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("default")
	}

	time.Sleep(5 * time.Second)
}
