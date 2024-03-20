package main

import (
	"fmt"
	"time"
)

func contador(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go contador("goroutine 1")
	go contador("goroutine 2")
	
	time.Sleep(11 * time.Second)

	fmt.Println("Fim do programa")
}
